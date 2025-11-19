package download

import (
	"fmt"
	"os"
	"riskmanagement/lib"
	audittrailmodels "riskmanagement/models/audittrail"
	briefingmodels "riskmanagement/models/briefing"
	coachingmodels "riskmanagement/models/coaching"
	models "riskmanagement/models/download"
	riskindicatorIdmodels "riskmanagement/models/riskindicator"
	verifmodels "riskmanagement/models/verifikasi"
	download "riskmanagement/repository/download"
	audittrail "riskmanagement/services/audittrail"
	briefing "riskmanagement/services/briefing"
	coaching "riskmanagement/services/coaching"
	riskindicator "riskmanagement/services/riskindicator"
	riskissue "riskmanagement/services/riskissue"
	verifikasi "riskmanagement/services/verifikasi"
	"strings"

	activity "riskmanagement/services/activity"
	unitKerja "riskmanagement/services/unitkerja"

	activityModel "riskmanagement/models/activity"
	// riskIssueModel "riskmanagement/models/riskissue"
	"strconv"

	"encoding/json"
	// "strings"

	// "github.com/google/uuid"
	"github.com/google/uuid"
	"gitlab.com/golang-package-library/logger"
	minio "gitlab.com/golang-package-library/minio"

	// "gorm.io/gorm"
	generateExcels "riskmanagement/jobs/queue/generateExcels"
	// "encoding/json"

	// RptVerifikasiRealpin
	verifRealpinModels "riskmanagement/models/verifikasireportrealisasi"
)

type DownloadDefinition interface {
	Generate(request models.DownloadRequest) (responses models.DownloadResponse, err error)
	DownloadHandler(id string) (url string, filename string, err error)
	GetListDownload(request models.ListDownloadRequest) (response []models.ListDownloadResponse, pagination int64, err error)
	GetReportType() (response []models.ReportTypeResponse, err error)
	Retry(request models.RetryRequest) (response bool, err error)
	FetchOneRows(id int64) (responses models.ListDownloadResponse, err error)
}

type DownloadService struct {
	db                lib.Database
	minio             minio.Minio
	logger            logger.Logger
	verifikasiRepo    verifikasi.VerifikasiDefinition
	briefingRepo      briefing.BriefingDefinition
	coachingRepo      coaching.CoachingDefinition
	downloadRepo      download.DownloadDefinition
	unitKerjaRepo     unitKerja.UnitKerjaDefinition
	activityRepo      activity.ActivityDefinition
	riskissueRepo     riskissue.RiskIssueDefinition
	riskindicatorRepo riskindicator.RiskIndicatorDefinition
	audittrailRepo    audittrail.AuditTrailDefinition
}

func NewDownloadService(
	db lib.Database,
	minio minio.Minio,
	logger logger.Logger,
	verifikasiRepo verifikasi.VerifikasiDefinition,
	briefingRepo briefing.BriefingDefinition,
	coachingRepo coaching.CoachingDefinition,
	downloadRepo download.DownloadDefinition,
	unitKerjaRepo unitKerja.UnitKerjaDefinition,
	activityRepo activity.ActivityDefinition,
	riskissueRepo riskissue.RiskIssueDefinition,
	riskindicatorRepo riskindicator.RiskIndicatorDefinition,
	audittrailRepo audittrail.AuditTrailDefinition,
) DownloadDefinition {
	return DownloadService{
		db:                db,
		minio:             minio,
		logger:            logger,
		verifikasiRepo:    verifikasiRepo,
		briefingRepo:      briefingRepo,
		coachingRepo:      coachingRepo,
		downloadRepo:      downloadRepo,
		unitKerjaRepo:     unitKerjaRepo,
		activityRepo:      activityRepo,
		riskissueRepo:     riskissueRepo,
		riskindicatorRepo: riskindicatorRepo,
		audittrailRepo:    audittrailRepo,
	}
}

var (
	date   = lib.GetTimeNow("date")
	hour   = lib.GetTimeNow("hour")
	minute = lib.GetTimeNow("minutes")
	second = lib.GetTimeNow("second")
)

func (service DownloadService) Generate(request models.DownloadRequest) (responses models.DownloadResponse, err error) {
	fmt.Println("masuk-service")

	jsonParams := request.JSONPARAMS
	var jsonUmarshal []map[string]interface{}
	fileName := ""
	var generateInfo generateExcels.GenerateInfo
	var columnNames []string

	if request.RequestStatus == "loading" {
		fmt.Println("request.RequestStatus ====>", "loading")

		// cek id sent from request
		data, isExistErr := service.downloadRepo.CheckIsExist(&request)
		if data.ID != 0 && isExistErr == nil {
			//get data by report_id and jsonparams
			fmt.Println("exist")
			return data, err
		}
	} else if request.RequestStatus == "generate" {
		fmt.Println("request.RequestStatus ====>", "generate")

		//if else condition for id
		// 1 for verifikasi_all_activity
		if request.ReportId == 1 {
			var requestDownload verifmodels.VerifikasiReportDetailRequest

			// Convert JSON string to struct
			if err := json.Unmarshal([]byte(jsonParams), &requestDownload); err != nil {
				fmt.Println("Error - unmarshal from json string to json struct:", err)
				return responses, err
			}

			dataDetail, errDetail := service.verifikasiRepo.VerifikasiReportDetail(requestDownload)
			if errDetail != nil {
				return responses, errDetail
			}

			jsonData, errMarshal := json.Marshal(dataDetail)
			if errMarshal != nil {
				fmt.Println("Error:", errMarshal)
				return responses, errMarshal
			}

			var dataMap map[string]interface{}
			errUnmarshal := json.Unmarshal(jsonData, &dataMap)
			if errUnmarshal != nil {
				fmt.Println("Error:", errUnmarshal)
				return responses, errUnmarshal
			}

			jsonUmarshal = append(jsonUmarshal, dataMap)

			headers, errHeaders := GetHeaderExcel(service, request.ReportId, jsonParams)
			if err != nil {
				fmt.Println("Error:", errHeaders)
				return responses, errHeaders
			}

			columnNames = []string{
				"No Pelaporan",
				"Kode Uker",
				"Uker",
				"Aktivitas",
				"Sub Aktivitas",
				"Produk",
				"Risk Issue",
				"Risk Indicator",
				"Kode Insiden",
				"Penyebab Insiden",
				"Kode Sub Insiden",
				"Penyebab Sub Insiden",
				"Hasil Verifikasi",
				"Sumber Data",
			}
			no_file := "VERIFICATION-01"
			fileName = no_file + " - Detail-Verifikasi Semua Aktifitas" + " - " + dataDetail.NoPelaporan + " - " + dataDetail.RGDESC + " - " + dataDetail.MBDESC + " - " + dataDetail.BRDESC

			generateInfo = generateExcels.GenerateInfo{
				NoFile:     no_file,
				ReportId:   request.ReportId,
				JSONPARAMS: jsonParams,
				MAKERID:    request.PERNR,
				MAKERDESC:  "",
			}

			fmt.Println("generateInfo", generateInfo)

			// maker_desc TO BE DISCUSS

			// generateExcels.Start(&service.minio, &service.db, jsonUmarshal, columnNames, headers, generateInfo, fileName)
			generateExcels.Start(&service.minio, &service.db, columnNames, headers, generateInfo, fileName, "00")
		} else if request.ReportId == 2 {
			var requestDownload briefingmodels.BriefingReportListRequest

			// Convert JSON string to struct
			if err := json.Unmarshal([]byte(jsonParams), &requestDownload); err != nil {
				fmt.Println("Error - unmarshal from json string to json struct:", err)
				return responses, err
			}

			headers, errHeaders := GetHeaderExcel(service, request.ReportId, jsonParams)
			if err != nil {
				fmt.Println("Error:", errHeaders)
				return responses, errHeaders
			}

			columnNames, fileName, no_file, err := GeneratorTemplate(models.GeneratorTemplate{
				ReportId:  int64(request.ReportId),
				Pernr:     requestDownload.Pernr,
				StartDate: requestDownload.StartDate,
				EndDate:   requestDownload.EndDate,
				Timestamp: requestDownload.Timestime,
			})

			if err != nil {
				fmt.Println("Error :", err)
				return responses, err
			}

			generateInfo = generateExcels.GenerateInfo{
				NoFile:     no_file,
				ReportId:   request.ReportId,
				JSONPARAMS: jsonParams,
				MAKERID:    request.PERNR,
				MAKERDESC:  "",
			}

			// fmt.Println("generateInfo", generateInfo)

			// maker_desc TO BE DISCUSS

			// generateExcels.Start(&service.minio, &service.db, jsonUmarshal, columnNames, headers, generateInfo, fileName)
			generateExcels.Start(&service.minio, &service.db, columnNames, headers, generateInfo, fileName, "00")
		} else if request.ReportId == 3 {
			var requestDownload coachingmodels.CoachingReportListRequest

			// Convert JSON string to struct
			if err := json.Unmarshal([]byte(jsonParams), &requestDownload); err != nil {
				fmt.Println("Error - unmarshal from json string to json struct:", err)
				return responses, err
			}

			headers, errHeaders := GetHeaderExcel(service, request.ReportId, jsonParams)
			if err != nil {
				fmt.Println("Error:", errHeaders)
				return responses, errHeaders
			}

			columnNames, fileName, no_file, err := GeneratorTemplate(models.GeneratorTemplate{
				ReportId:  int64(request.ReportId),
				Pernr:     requestDownload.Pernr,
				StartDate: requestDownload.StartDate,
				EndDate:   requestDownload.EndDate,
				Timestamp: requestDownload.Timestime,
			})

			if err != nil {
				fmt.Println("Error :", err)
				return responses, err
			}

			generateInfo = generateExcels.GenerateInfo{
				NoFile:     no_file,
				ReportId:   request.ReportId,
				JSONPARAMS: jsonParams,
				MAKERID:    request.PERNR,
				MAKERDESC:  "",
			}

			// generateExcels.Start(&service.minio, &service.db, jsonUmarshal, columnNames, headers, generateInfo, fileName)
			generateExcels.Start(&service.minio, &service.db, columnNames, headers, generateInfo, fileName, "00")
		} else if request.ReportId == 4 {
			var requestDownload verifmodels.VerifikasiReportListRequest

			// Convert JSON string to struct
			if err := json.Unmarshal([]byte(jsonParams), &requestDownload); err != nil {
				fmt.Println("Error - unmarshal from json string to json struct:", err)
				return responses, err
			}

			headers, errHeaders := GetHeaderExcel(service, request.ReportId, jsonParams)
			if err != nil {
				fmt.Println("Error:", errHeaders)
				return responses, errHeaders
			}

			columnNames, fileName, no_file, err := GeneratorTemplate(models.GeneratorTemplate{
				ReportId:  int64(request.ReportId),
				Pernr:     requestDownload.Pernr,
				StartDate: requestDownload.StartDate,
				EndDate:   requestDownload.EndDate,
				Timestamp: requestDownload.Timestime,
			})

			if err != nil {
				fmt.Println("Error :", err)
				return responses, err
			}

			generateInfo = generateExcels.GenerateInfo{
				NoFile:     no_file,
				ReportId:   request.ReportId,
				JSONPARAMS: jsonParams,
				MAKERID:    request.PERNR,
				MAKERDESC:  "",
			}
			// maker_desc TO BE DISCUSS

			// generateExcels.Start(&service.minio, &service.db, jsonUmarshal, columnNames, headers, generateInfo, fileName)
			generateExcels.Start(&service.minio, &service.db, columnNames, headers, generateInfo, fileName, "00")
		} else if request.ReportId == 5 {
			var requestDownload audittrailmodels.FilterAudit

			// Convert JSON string to struct
			if err := json.Unmarshal([]byte(jsonParams), &requestDownload); err != nil {
				fmt.Println("Error - unmarshal from json string to json struct:", err)
				return responses, err
			}

			headers, errHeaders := GetHeaderExcel(service, request.ReportId, jsonParams)
			if err != nil {
				fmt.Println("Error:", errHeaders)
				return responses, errHeaders
			}

			// fmt.Println("headers =>", headers)
			columnNames, fileName, no_file, err := GeneratorTemplate(models.GeneratorTemplate{
				ReportId:  int64(request.ReportId),
				Pernr:     requestDownload.PERNR,
				StartDate: requestDownload.StartDate,
				EndDate:   requestDownload.EndDate,
				Timestamp: requestDownload.Timestime,
			})

			if err != nil {
				fmt.Println("Error :", err)
				return responses, err
			}

			generateInfo = generateExcels.GenerateInfo{
				NoFile:     no_file,
				ReportId:   request.ReportId,
				JSONPARAMS: jsonParams,
				MAKERID:    request.PERNR,
				MAKERDESC:  "",
			}

			fmt.Println("generateInfo", generateInfo)

			// maker_desc TO BE DISCUSS

			// generateExcels.Start(&service.minio, &service.db, jsonUmarshal, columnNames, headers, generateInfo, fileName)
			generateExcels.Start(&service.minio, &service.db, columnNames, headers, generateInfo, fileName, "00")
		} else if request.ReportId == 6 {
			var requestDownload riskindicatorIdmodels.RiskIndicatorGetOne

			if err := json.Unmarshal([]byte(jsonParams), &requestDownload); err != nil {
				fmt.Println("Error - unmarshal from json string to json struct:", err)
				return responses, err
			}

			headers, errHeaders := GetHeaderExcel(service, request.ReportId, jsonParams)
			if err != nil {
				fmt.Println("Error:", errHeaders)
				return responses, errHeaders
			}

			columnNames, fileName, no_file, err := GeneratorTemplate(models.GeneratorTemplate{
				ReportId: int64(request.ReportId),
				// Pernr:     requestDownload.PERNR,
				// StartDate: requestDownload.StartDate,
				// EndDate:   requestDownload.EndDate,
				// Timestamp: requestDownload.Timestime,
			})

			if err != nil {
				fmt.Println("Error :", err)
				return responses, err
			}

			generateInfo = generateExcels.GenerateInfo{
				NoFile:     no_file,
				ReportId:   request.ReportId,
				JSONPARAMS: jsonParams,
				MAKERID:    request.PERNR,
				MAKERDESC:  "",
			}

			fmt.Println("generateInfo", generateInfo)

			// generateExcels.Start(&service.minio, &service.db, jsonUmarshal, columnNames, headers, generateInfo, fileName)
			generateExcels.Start(&service.minio, &service.db, columnNames, headers, generateInfo, fileName, "00")
		} else if request.ReportId == 7 {
			var requestDownload verifmodels.RptRekapitulasiBCVRequest

			// Convert JSON string to struct
			if err := json.Unmarshal([]byte(jsonParams), &requestDownload); err != nil {
				fmt.Println("Error - unmarshal from json string to json struct:", err)
				return responses, err
			}

			headers, errHeaders := GetHeaderExcel(service, request.ReportId, jsonParams)
			if err != nil {
				fmt.Println("Error:", errHeaders)
				return responses, errHeaders
			}

			columnNames, fileName, no_file, err := GeneratorTemplate(models.GeneratorTemplate{
				ReportId:  int64(request.ReportId),
				Pernr:     requestDownload.BRC,
				StartDate: requestDownload.StartDate,
				EndDate:   requestDownload.EndDate,
				// Timestamp: requestDownload.Timestime,
			})

			if err != nil {
				fmt.Println("Error :", err)
				return responses, err
			}

			generateInfo = generateExcels.GenerateInfo{
				NoFile:     no_file,
				ReportId:   request.ReportId,
				JSONPARAMS: jsonParams,
				MAKERID:    request.PERNR,
				MAKERDESC:  "",
			}

			fmt.Println("generateInfo", generateInfo)

			// generateExcels.Start(&service.minio, &service.db, jsonUmarshal, columnNames, headers, generateInfo, fileName)
			generateExcels.Start(&service.minio, &service.db, columnNames, headers, generateInfo, fileName, "00")
		} else if request.ReportId == 8 {
			var requestDownload verifmodels.RptRekomendasiRiskRequest

			// Convert JSON string to struct
			if err := json.Unmarshal([]byte(jsonParams), &requestDownload); err != nil {
				fmt.Println("Error - unmarshal from json string to json struct:", err)
				return responses, err
			}

			headers, errHeaders := GetHeaderExcel(service, request.ReportId, jsonParams)
			if err != nil {
				fmt.Println("Error:", errHeaders)
				return responses, errHeaders
			}

			columnNames, fileName, no_file, err := GeneratorTemplate(models.GeneratorTemplate{
				ReportId:   int64(request.ReportId),
				Pernr:      requestDownload.JenisData,
				StartDate:  requestDownload.StartDate,
				EndDate:    requestDownload.EndDate,
				JsonParams: request.JSONPARAMS,
				// Timestamp: requestDownload.Timestime,
			})

			if err != nil {
				fmt.Println("Error :", err)
				return responses, err
			}

			generateInfo = generateExcels.GenerateInfo{
				NoFile:     no_file,
				ReportId:   request.ReportId,
				JSONPARAMS: jsonParams,
				MAKERID:    request.PERNR,
				MAKERDESC:  "",
			}

			fmt.Println("generateInfo", generateInfo)

			// generateExcels.Start(&service.minio, &service.db, jsonUmarshal, columnNames, headers, generateInfo, fileName)
			generateExcels.Start(&service.minio, &service.db, columnNames, headers, generateInfo, fileName, "00")
		} else if request.ReportId == 9 {
			fmt.Println("Rpt Verifikasi Realisasi Pinjaman List")
			var requestDownload verifRealpinModels.ReportRealisasiKreditListRequest

			if err := json.Unmarshal([]byte(jsonParams), &requestDownload); err != nil {
				fmt.Println("Error - unmarshal from json string to json struct:", err)
				return responses, err
			}

			headers, errHeaders := GetHeaderExcel(service, request.ReportId, jsonParams)
			if err != nil {
				fmt.Println("Error:", errHeaders)
				return responses, errHeaders
			}

			columnNames, fileName, no_file, err := GeneratorTemplate(models.GeneratorTemplate{
				ReportId:   int64(request.ReportId),
				Pernr:      requestDownload.Pernr,
				Timestamp:  requestDownload.Timestime,
				JsonParams: request.JSONPARAMS,
			})

			if err != nil {
				fmt.Println("Error :", err)
				return responses, err
			}

			generateInfo = generateExcels.GenerateInfo{
				NoFile:     no_file,
				ReportId:   request.ReportId,
				JSONPARAMS: jsonParams,
				MAKERID:    request.PERNR,
				MAKERDESC:  "",
			}

			generateExcels.Start(&service.minio, &service.db, columnNames, headers, generateInfo, fileName, "00")
		} else if request.ReportId == 10 {
			fmt.Println("Rpt Verifikasi Realisasi Pinjaman Summary")
			var requestDownload verifRealpinModels.ReportRealisasiKreditSummaryRequest

			if err := json.Unmarshal([]byte(jsonParams), &requestDownload); err != nil {
				fmt.Println("Error - unmarshal from json string to json struct:", err)
				return responses, err
			}

			headers, errHeaders := GetHeaderExcel(service, request.ReportId, jsonParams)
			if errHeaders != nil {
				fmt.Println("Error:", errHeaders)
				return responses, errHeaders
			}

			columnNames, fileName, no_file, err := GeneratorTemplate(models.GeneratorTemplate{
				ReportId:   int64(request.ReportId),
				Pernr:      requestDownload.Pernr,
				Timestamp:  requestDownload.Timestime,
				JsonParams: request.JSONPARAMS,
			})
			fmt.Println(fileName, no_file)
			if err != nil {
				fmt.Println("Error :", err)
				return responses, err
			}

			generateInfo = generateExcels.GenerateInfo{
				NoFile:     no_file,
				ReportId:   request.ReportId,
				JSONPARAMS: jsonParams,
				MAKERID:    request.PERNR,
				MAKERDESC:  "",
			}

			generateExcels.Start(&service.minio, &service.db, columnNames, headers, generateInfo, fileName, "00")
		}
	}

	return responses, err
}

func (service DownloadService) DownloadHandler(id string) (url string, filename string, err error) {
	fmt.Println("masuk service")

	result, errRes := service.downloadRepo.GetDownloadUrl(id)
	if errRes != nil {
		return "", "", errRes
	}

	filename = result.FILENAME
	bucket := os.Getenv("BUCKET_NAME")

	preSign := service.minio.SignUrl(service.minio.Client(), bucket, result.FILEPATH, filename)
	url = fmt.Sprint(preSign)

	fmt.Println("url", url)

	return url, filename, err
}

func GetHeaderExcel(service DownloadService, reportId int, jsonParams string) (headers map[string]string, err error) {
	if reportId == 2 {
		var requestDownload briefingmodels.BriefingReportListRequest

		// Convert JSON string to struct
		if err := json.Unmarshal([]byte(jsonParams), &requestDownload); err != nil {
			fmt.Println("Error - unmarshal from json string to json struct:", err)
			return headers, err
		}

		// variables
		noPelaporan := "-"
		status := requestDownload.Status
		kanwil := "Semua"
		kanca := "Semua"
		uker := "Semua"
		activity := "Semua"
		judulMateri := "Semua"
		periode := requestDownload.StartDate + " s.d " + requestDownload.EndDate

		// get kanwil name
		if requestDownload.REGION != "all" {
			kanwil, _ = service.unitKerjaRepo.GetRegionName(requestDownload.REGION)
		}
		// get kanca name
		if requestDownload.MAINBR != "all" {
			kanca, _ = service.unitKerjaRepo.GetMainbrName(requestDownload.MAINBR)
		}
		// get uker name
		if requestDownload.BRANCH != "all" {
			uker, _ = service.unitKerjaRepo.GetBranchName(requestDownload.BRANCH)
		}

		// conditions
		if requestDownload.NoPelaporan != "" {
			noPelaporan = requestDownload.NoPelaporan
		}

		if requestDownload.JudulMateri == "" {
			judulMateri = "-"
		} else if requestDownload.JudulMateri != "all" {
			judulMateri = requestDownload.JudulMateri
		}

		if requestDownload.ActivityID != "all" {
			// get activity name
			var dataActivity activityModel.ActivityResponse
			activityId, err := strconv.ParseInt(requestDownload.ActivityID, 10, 64)
			if err != nil {
				fmt.Println("Error:", err)
				return headers, err
			}

			dataActivity, err = service.activityRepo.GetOne(activityId)
			if err != nil {
				fmt.Println("Error:", err)
				return headers, err
			}

			activity = dataActivity.Name
		}

		headers = map[string]string{
			"Nomor Pelaporan: ": noPelaporan,
			"Kanwil: ":          kanwil,
			"Kanca: ":           kanca,
			"Unit Kerja: ":      uker,
			"Aktivitas: ":       activity,
			"Judul Materi: ":    judulMateri,
			"Status: ":          status,
			"Periode: ":         periode,
		}
	} else if reportId == 3 {
		var requestDownload coachingmodels.CoachingReportListRequest

		// Convert JSON string to struct
		if err := json.Unmarshal([]byte(jsonParams), &requestDownload); err != nil {
			fmt.Println("Error - unmarshal from json string to json struct:", err)
			return headers, err
		}

		// variables
		noPelaporan := "-"
		status := requestDownload.Status
		kanwil := "Semua"
		kanca := "Semua"
		uker := "Semua"
		activity := "Semua"
		riskEvent := "Semua"
		judulMateri := "Semua"
		periode := requestDownload.StartDate + " s.d " + requestDownload.EndDate

		// get kanwil name
		if requestDownload.REGION != "all" {
			kanwil, _ = service.unitKerjaRepo.GetRegionName(requestDownload.REGION)
		}
		// get kanca name
		if requestDownload.MAINBR != "all" {
			kanca, _ = service.unitKerjaRepo.GetMainbrName(requestDownload.MAINBR)
		}
		// get uker name
		if requestDownload.BRANCH != "all" {
			uker, _ = service.unitKerjaRepo.GetBranchName(requestDownload.BRANCH)
		}

		// conditions
		if requestDownload.NoPelaporan != "" {
			noPelaporan = requestDownload.NoPelaporan
		}

		if requestDownload.JudulMateri == "" {
			judulMateri = "-"
		} else if requestDownload.JudulMateri != "all" {
			judulMateri = requestDownload.JudulMateri
		}

		if requestDownload.ActivityID != "all" {
			// get activity name
			var dataActivity activityModel.ActivityResponse
			activityId, err := strconv.ParseInt(requestDownload.ActivityID, 10, 64)
			if err != nil {
				fmt.Println("Error:", err)
				return headers, err
			}

			dataActivity, err = service.activityRepo.GetOne(activityId)
			if err != nil {
				fmt.Println("Error:", err)
				return headers, err
			}

			activity = dataActivity.Name
		}

		if requestDownload.RiskIssueID != "all" {
			// get activity name
			id, err := strconv.ParseInt(requestDownload.RiskIssueID, 10, 64)
			if err != nil {
				fmt.Println("Error:", err)
				return headers, err
			}

			data, err := service.riskissueRepo.GetRiskEventName(id)
			if err != nil {
				fmt.Println("Error:", err)
				return headers, err
			}

			riskEvent = data
		}

		headers = map[string]string{
			"Nomor Pelaporan: ": noPelaporan,
			"Kanwil: ":          kanwil,
			"Kanca: ":           kanca,
			"Unit Kerja: ":      uker,
			"Aktivitas: ":       activity,
			"Risk Event: ":      riskEvent,
			"Judul Materi: ":    judulMateri,
			"Status: ":          status,
			"Periode: ":         periode,
		}
	} else if reportId == 4 {
		var requestDownload verifmodels.VerifikasiReportListRequest

		// Convert JSON string to struct
		if err := json.Unmarshal([]byte(jsonParams), &requestDownload); err != nil {
			fmt.Println("Error - unmarshal from json string to json struct:", err)
			return headers, err
		}

		// variables
		noPelaporan := "-"
		namaBrcUrc := "Semua"
		kanwil := "Semua"
		kanca := "Semua"
		uker := "Semua"
		riskEvent := "Semua"
		riskIndicator := "Semua"
		indikasiFraud := "Semua"
		status := "Semua"
		periode := requestDownload.StartDate + " s.d " + requestDownload.EndDate

		// get kanwil name
		if requestDownload.REGION != "all" {
			kanwil, _ = service.unitKerjaRepo.GetRegionName(requestDownload.REGION)
		}
		// get kanca name
		if requestDownload.MAINBR != "all" {
			kanca, _ = service.unitKerjaRepo.GetMainbrName(requestDownload.MAINBR)
		}
		// get uker name
		if requestDownload.BRANCH != "all" {
			uker, _ = service.unitKerjaRepo.GetBranchName(requestDownload.BRANCH)
		}

		// conditions
		if requestDownload.NoPelaporan != "" {
			noPelaporan = requestDownload.NoPelaporan
		}

		if requestDownload.BrcUrc != "" {
			namaBrcUrc = requestDownload.BrcUrc
		}

		if requestDownload.Status != "all" {
			status = requestDownload.Status
		}

		if requestDownload.IndikasiFraud == "1" {
			indikasiFraud = "Ya"
		} else if requestDownload.IndikasiFraud == "2" {
			indikasiFraud = "Tidak"
		}

		if requestDownload.RiskIssueID != "all" {
			id, err := strconv.ParseInt(requestDownload.RiskIssueID, 10, 64)
			if err != nil {
				fmt.Println("Error:", err)
				return headers, err
			}

			data, err := service.riskissueRepo.GetRiskEventName(id)
			if err != nil {
				fmt.Println("Error:", err)
				return headers, err
			}

			riskEvent = data
		}

		// if requestDownload.RiskIndicatorID != "all" {
		// 	var data riskindicatorIdmodels.RiskIndicatorGetOne
		// 	id, err := strconv.ParseInt(requestDownload.RiskIndicatorID, 10, 64)
		// 	if err != nil {
		// 		fmt.Println("Error:", err)
		// 		return headers, err
		// 	}

		// 	data, _, err = service.riskindicatorRepo.GetOne(id)
		// 	if err != nil {
		// 		fmt.Println("Error:", err)
		// 		return headers, err
		// 	}

		// 	riskIndicator = data.RiskIndicator
		// }

		if requestDownload.RiskIndicator != "all" {
			riskIndicator = requestDownload.RiskIndicator
		}

		headers = map[string]string{
			"Nomor Pelaporan: ": noPelaporan,
			"Nama BRC/URC: ":    namaBrcUrc,
			"Kanwil: ":          kanwil,
			"Kanca: ":           kanca,
			"Unit Kerja: ":      uker,
			"Risk Event: ":      riskEvent,
			"Risk Indicator: ":  riskIndicator,
			"Indikasi Fraud: ":  indikasiFraud,
			"Status: ":          status,
			"Periode: ":         periode,
		}
	} else if reportId == 5 {
		var requestDownload audittrailmodels.FilterAudit

		// Convert JSON string to struct
		if err := json.Unmarshal([]byte(jsonParams), &requestDownload); err != nil {
			fmt.Println("Error - unmarshal from json string to json struct:", err)
			return headers, err
		}

		// variables
		namaBrcUrc := "Semua"
		aktivitas := "Semua"
		kanwil := "Semua"
		kanca := "Semua"
		uker := "Semua"
		periode := requestDownload.StartDate + " s.d " + requestDownload.EndDate

		// get kanwil name
		if requestDownload.REGION != "all" {
			kanwil, _ = service.unitKerjaRepo.GetRegionName(requestDownload.REGION)
		}
		// get kanca name
		if requestDownload.MAINBR != "all" {
			kanca, _ = service.unitKerjaRepo.GetMainbrName(requestDownload.MAINBR)
		}
		// get uker name
		if requestDownload.BRANCH != "all" {
			uker, _ = service.unitKerjaRepo.GetBranchName(requestDownload.BRANCH)
		}

		if requestDownload.PERNR != "" {
			namaBrcUrc = requestDownload.PERNR
		}

		if requestDownload.Aktifitas != "all" {
			aktivitas = requestDownload.Aktifitas
		}

		headers = map[string]string{
			"Nama BRC/URC : ": namaBrcUrc,
			"Aktivitas : ":    aktivitas,
			"Kanwil : ":       kanwil,
			"Kanca : ":        kanca,
			"Unit Kerja : ":   uker,
			"Periode : ":      periode,
		}
	} else if reportId == 7 {
		var requestDownload verifmodels.RptRekapitulasiBCVRequest

		if err := json.Unmarshal([]byte(jsonParams), &requestDownload); err != nil {
			fmt.Println("Error - unmarshal from json string to json struct:", err)
			return headers, err
		}

		kanwil := "Semua"
		kanca := "Semua"
		uker := "Semua"
		namaBrcUrc := "Semua"
		periode := requestDownload.StartDate + " s.d " + requestDownload.EndDate

		// get kanwil name
		if requestDownload.REGION != "all" {
			kanwil, _ = service.unitKerjaRepo.GetRegionName(requestDownload.REGION)
		}
		// get kanca name
		if requestDownload.MAINBR != "all" {
			kanca, _ = service.unitKerjaRepo.GetMainbrName(requestDownload.MAINBR)
		}
		// get uker name
		if requestDownload.BRANCH != "all" {
			uker, _ = service.unitKerjaRepo.GetBranchName(requestDownload.BRANCH)
		}

		if requestDownload.BRC != "" {
			namaBrcUrc = requestDownload.BRC
		}

		headers = map[string]string{
			"Kanwil: ":       kanwil,
			"Kanca: ":        kanca,
			"Unit Kerja: ":   uker,
			"Nama BRC/URC: ": namaBrcUrc,
			"Periode: ":      periode,
		}
	} else if reportId == 8 {
		var requestDownload verifmodels.RptRekomendasiRiskRequest

		if err := json.Unmarshal([]byte(jsonParams), &requestDownload); err != nil {
			fmt.Println("Error - unmarshal from json string to json struct:", err)
			return headers, err
		}

		jenisData := requestDownload.JenisData
		periode := requestDownload.StartDate + " s.d " + requestDownload.EndDate

		headers = map[string]string{
			"Jenis Data: ": jenisData,
			"Periode: ":    periode,
		}
	} else if reportId == 9 {
		var requestDownload verifRealpinModels.ReportRealisasiKreditListRequest

		if err := json.Unmarshal([]byte(jsonParams), &requestDownload); err != nil {
			fmt.Println("Error - unmarshal from json string to json struct:", err)
			return headers, err
		}

		jenisReport := requestDownload.ReportType
		kanwil := "Semua"
		kanca := "Semua"
		uker := "Semua"

		if requestDownload.REGION != "all" {
			kanwil, _ = service.unitKerjaRepo.GetRegionName(requestDownload.REGION)
		}
		// get kanca name
		if requestDownload.MAINBR != "all" {
			kanca, _ = service.unitKerjaRepo.GetMainbrName(requestDownload.MAINBR)
		}
		// get uker name
		if requestDownload.BRANCH != "all" {
			uker, _ = service.unitKerjaRepo.GetBranchName(requestDownload.BRANCH)
		}

		headers = map[string]string{
			"Jenis Report : ": jenisReport,
			"Kanwil : ":       kanwil,
			"Kanca : ":        kanca,
			"Unit Kerja : ":   uker,
		}
	} else if reportId == 10 {
		var requestDownload verifRealpinModels.ReportRealisasiKreditSummaryRequest

		if err := json.Unmarshal([]byte(jsonParams), &requestDownload); err != nil {
			fmt.Println("Error - unmarshal from json string to json struct:", err)
			return headers, err
		}

		jenisReport := requestDownload.ReportType
		if len(jenisReport) >= 0 {
			jenisReport = strings.ToUpper(string(jenisReport[0])) + jenisReport[1:]
		}
		kanwil := "Semua"
		kanca := "Semua"
		uker := "Semua"

		if requestDownload.REGION != "all" {
			kanwil, _ = service.unitKerjaRepo.GetRegionName(requestDownload.REGION)
		}
		// get kanca name
		if requestDownload.MAINBR != "all" {
			kanca, _ = service.unitKerjaRepo.GetMainbrName(requestDownload.MAINBR)
		}
		// get uker name
		if requestDownload.BRANCH != "all" {
			uker, _ = service.unitKerjaRepo.GetBranchName(requestDownload.BRANCH)
		}

		headers = map[string]string{
			"Jenis Report : ": jenisReport,
			"Kanwil : ":       kanwil,
			"Kanca : ":        kanca,
			"Unit Kerja : ":   uker,
		}
	}

	return headers, err
}

// GetListDownload implements DownloadDefinition.
func (service DownloadService) GetListDownload(request models.ListDownloadRequest) (responses []models.ListDownloadResponse, pagination int64, err error) {
	list_download, pagination, err := service.downloadRepo.GetListDownload(&request)

	if err != nil {
		service.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, value := range list_download {
		responses = append(responses, models.ListDownloadResponse{
			No:          value.No,
			Id:          value.Id,
			ReportId:    value.ReportId,
			NamaLaporan: value.NamaLaporan,
			Kanwil:      value.Kanwil,
			Kanca:       value.Kanca,
			UnitKerja:   value.UnitKerja,
			PeriodeData: value.PeriodeData,
			Status:      value.Status,
			FileDesc:    value.FileDesc,
			Filename:    value.Filename,
			Filepath:    value.Filepath,
			JsonParams:  value.JsonParams,
		})
	}

	return responses, pagination, err
}

// GetReportType implements DownloadDefinition.
func (service DownloadService) GetReportType() (response []models.ReportTypeResponse, err error) {
	fmt.Println("Masuk service")

	list_data, err := service.downloadRepo.GetReportType()

	if err != nil {
		service.logger.Zap.Error(err)
		return response, err
	}

	for _, value := range list_data {
		response = append(response, models.ReportTypeResponse{
			Id:   value.Id,
			Name: value.Name,
		})
	}

	return response, err
}

// Retry implements DownloadDefinition.
func (service DownloadService) Retry(request models.RetryRequest) (response bool, err error) {
	if request.RetryCode == "03" {
		generateInfo := generateExcels.GenerateInfo{
			ID:         request.InsertId,
			JSONPARAMS: request.JsonParams,
			FILENAME:   request.Filename,
		}

		generateExcels.Start(&service.minio, &service.db, nil, nil, generateInfo, request.Filename, request.RetryCode)
	}

	if request.RetryCode == "01" {
		headers, errHeaders := GetHeaderExcel(service, request.ReportId, request.JsonParams)
		if err != nil {
			fmt.Println("Error:", errHeaders)
			return false, errHeaders
		}

		fmt.Println("Header =>", headers)

		ColumnName, FileName, _, err := GeneratorTemplate(models.GeneratorTemplate{
			ReportId:   int64(request.ReportId),
			Pernr:      request.Pernr,
			StartDate:  request.StartDate,
			EndDate:    request.EndDate,
			JsonParams: request.JsonParams,
		})

		if err != nil {
			fmt.Println(err)
			return false, err
		}

		fmt.Println("retry SERVICE =>", ColumnName)
		fmt.Println("FILENAME =>", FileName)

		generateInfo := generateExcels.GenerateInfo{
			ID:         request.InsertId,
			ReportId:   request.ReportId,
			JSONPARAMS: request.JsonParams,
			FILENAME:   FileName,
		}

		generateExcels.Start(&service.minio, &service.db, ColumnName, headers, generateInfo, FileName, request.RetryCode)
	}

	// data, err := service.downloadRepo.CheckRptStatus(int64(request.InsertId))

	return true, err
}

// FetchOneRows implements DownloadDefinition.
func (service DownloadService) FetchOneRows(id int64) (responses models.ListDownloadResponse, err error) {
	data, err := service.downloadRepo.CheckRptStatus(id)

	return data, err
}

func GeneratorTemplate(request models.GeneratorTemplate) (ColumnName []string, FileName string, NoFile string, err error) {
	if request.ReportId == 2 {
		ColumnName = []string{
			"No",
			"Kode Branch",
			"Unit Kerja",
			"Kanca",
			"Kanwil",
			"No Pelaporan",
			"Judul Materi",
			"Risk Event",
			"Rincian Materi",
			"Jumlah Peserta",
			"Jabatan Peserta",
			"Peserta",
			"Aktifitas",
			"Status",
			"Maker",
		}

		UUID := uuid.New()
		NoFile = "BRIEFING-01"
		FileName = NoFile + "-Report-List Briefing" + " - " + request.Pernr + " - " + request.StartDate + " - " + request.EndDate + " - " + UUID.String()
	} else if request.ReportId == 3 {
		ColumnName = []string{
			"No",
			"Kode Branch",
			"Unit Kerja",
			"Kanca",
			"Kanwil",
			"No Pelaporan",
			"Judul Materi",
			"Rincian Materi",
			"Jumlah Peserta",
			"Jabatan Peserta",
			"Peserta Coaching",
			"Aktifitas",
			"Sub Aktifitas",
			"Isu Risiko",
			"Risk Indicator",
			"Status",
			"Maker",
		}

		UUID := uuid.New()
		NoFile = "COACHING-01"
		FileName = NoFile + "-Report-List Coaching" + " - " + request.Pernr + " - " + request.StartDate + " - " + request.EndDate + " - " + UUID.String()
	} else if request.ReportId == 4 {
		ColumnName = []string{
			"No",
			"Kode Branch",
			"Unit Kerja",
			"Kanca",
			"Kanwil",
			"No Pelaporan",
			"Aktifitas",
			"Sub Aktifitas",
			"Informasi Lainnya",
			"Status Perbaikan (Konsolidasi)",
			"Maker",
			"ID Risk Event",
			"Risk Event Name",
			"Risk Indicator",
			"Risk Control",
			"Hasil Verifikasi",
			"Jumlah Data yg Diverifikasi",
			"Butuh Perbaikan",
			"Jumlah Data yg Harus di Perbaiki",
			"RTL Uker",
			"Status Perbaikan Selesai",
			"Status Perbaikan yg Masih Proses",
			"Persentase Perbaikan",
			"Batas Waktu Perbaikan",
			"Indikasi Fraud",
		}

		UUID := uuid.New()
		NoFile = "VERIFICATION-01"
		FileName = NoFile + "-Report-List Verifikasi" + " - " + request.Pernr + " - " + request.StartDate + " - " + request.EndDate + " - " + UUID.String()
	} else if request.ReportId == 5 {
		ColumnName = []string{
			"No",
			"Tanggal",
			"PN",
			"Nama BRC/URC",
			"Kanwil",
			"Kanca",
			"Unit Kerja",
			"No. Pelaporan",
			"Aktivitas",
			"IP Address",
			"Lokasi",
		}

		UUID := uuid.New()
		NoFile = "AUDITTRAIL-01"
		FileName = NoFile + "-Report-Audit Trail" + " - " + request.Pernr + " - " + request.StartDate + " - " + request.EndDate + " - " + UUID.String()
	} else if request.ReportId == 6 {
		ColumnName = []string{
			"Id",
			"Key_Risk_Indicator",
			"Aktivitas",
			"Produk",
			"Jenis_Indikator",
			"Indikasi_Risiko",
			"Deskripsi",
			"SLA_Verifikasi",
			"SLA_TL",
			"Risk_Awarness",
			"Data_Source",
			"KCK_1_MIN",
			"KCK_2_MIN",
			"KCK_3_MIN",
			"KCK_4_MIN",
			"KCK_5_MIN",
			"KC_1_MIN",
			"KC_2_MIN",
			"KC_3_MIN",
			"KC_4_MIN",
			"KC_5_MIN",
			"KCP_1_MIN",
			"KCP_2_MIN",
			"KCP_3_MIN",
			"KCP_4_MIN",
			"KCP_5_MIN",
			"UN_1_MIN",
			"UN_2_MIN",
			"UN_3_MIN",
			"UN_4_MIN",
			"UN_5_MIN",
			"KK_1_MIN",
			"KK_2_MIN",
			"KK_3_MIN",
			"KK_4_MIN",
			"KK_5_MIN",
			"KCK_1_MAX",
			"KCK_2_MAX",
			"KCK_3_MAX",
			"KCK_4_MAX",
			"KCK_5_MAX",
			"KC_1_MAX",
			"KC_2_MAX",
			"KC_3_MAX",
			"KC_4_MAX",
			"KC_5_MAX",
			"KCP_1_MAX",
			"KCP_2_MAX",
			"KCP_3_MAX",
			"KCP_4_MAX",
			"KCP_5_MAX",
			"UN_1_MAX",
			"UN_2_MAX",
			"UN_3_MAX",
			"UN_4_MAX",
			"UN_5_MAX",
			"KK_1_MAX",
			"KK_2_MAX",
			"KK_3_MAX",
			"KK_4_MAX",
			"KK_5_MAX",
			"Parameter",
			"Status Indikator",
			"is_aktif",
		}

		UUID := uuid.New()
		NoFile = "THRESHOLD-01"
		FileName = NoFile + "KRID_THRESHOLD_INDICATOR_" + UUID.String()
	} else if request.ReportId == 7 {
		ColumnName = []string{
			"No",
			"PN",
			"BRC",
			"KODE BRANCH",
			"UNIT KERJA",
			"KANCA",
			"KANWIL",
			"BRIEFING",
			"COACHING",
			"VERIFIKASI",
		}

		UUID := uuid.New()
		NoFile = "VERIFICATION-02"
		FileName = NoFile + "-Report-Rekapituasi BCV" + " - " + request.Pernr + " - " + request.StartDate + " - " + request.EndDate + " - " + UUID.String()
	} else if request.ReportId == 8 {
		if request.Pernr == "Risk Event" {
			ColumnName = []string{
				"No",
				"Risk Event",
				"Module",
				"Count",
			}
		} else {
			ColumnName = []string{
				"No",
				"Risk Indicator",
				"Module",
				"Count",
			}
		}

		UUID := uuid.New()
		NoFile = "VERIFICATION-03"
		FileName = NoFile + "-Report-Rekapituasi BCV" + " - " + request.Pernr + " - " + request.StartDate + " - " + request.EndDate + " - " + UUID.String()
	} else if request.ReportId == 9 {
		ColumnName = []string{
			"Nomor Pelaporan",
			"AKTIVITAS",
			"NO REKENING",
			"RESTRUK",
			"NAMA",
			"KANWIL",
			"KANCA",
			"UKER",
			"CIF",
			"TANGGAL REALISASI",
			"TANGGAL JATUH TEMPO",
			"SEGMEN",
			"PRODUK",
			"PN PEMRAKARSA",
			"NAMA PEMRAKARSA",
			"OUTSTANDING",
			"PLAFOND",
			"LOAN_TYPE",
			"PEMUTUS",
			"SUDAH DILAKUKAN VERIFIKASI",
			"EFEKTIF",
			"KRITERIA",
			"HASIL VERIFIKASI",
			"KUNJUNGAN NASABAH",
			"TANGGAL KUNJUNGAN",
			"PN BRC/URC",
			"NAMA BRC/URC",
		}

		UUID := uuid.New()
		NoFile = "VERIFICATION-04"
		FileName = NoFile + "-Report-VerifikasiRealPin-List " + " - " + request.Pernr + " - " + UUID.String()
	} else if request.ReportId == 10 {

		var requestParams verifRealpinModels.ReportRealisasiKreditSummaryRequest

		if err := json.Unmarshal([]byte(request.JsonParams), &requestParams); err != nil {
			fmt.Println("Error - unmarshal from json string to json struct:", err)
		}

		ColumnName = []string{
			"NO.",
		}

		if containsString(requestParams.GroupBy, "produk") {
			ColumnName = append(ColumnName, "PRODUK")
		}

		if containsString(requestParams.GroupBy, "pn-pemrakarsa") {
			ColumnName = append(ColumnName, "PN PEMRAKARSA")
		}

		if containsString(requestParams.GroupBy, "regional-office") {
			ColumnName = append(ColumnName, "KANWIL")
		}

		if containsString(requestParams.GroupBy, "branch-office") {
			ColumnName = append(ColumnName, "KANCA")
		}

		if containsString(requestParams.GroupBy, "unit-kerja") {
			ColumnName = append(ColumnName, "UKER")
		}

		if containsString(requestParams.GroupBy, "pn-brc-urc") {
			ColumnName = append(ColumnName, "PN BRC/URC", "NAMA BRC/URC")
		}

		if containsString(requestParams.GroupBy, "efektifitas") {
			ColumnName = append(ColumnName, "EFEKTIF", "TIDAK EFEKTIF")
		}

		if containsString(requestParams.GroupBy, "status-verifikasi") {
			ColumnName = append(ColumnName, "SUDAH DILAKUKAN VERIFIKASI")
		}

		if containsString(requestParams.GroupBy, "criteria") {
			ColumnName = append(ColumnName, "CRITERIA", "YA", "TIDAK")
		}

		UUID := uuid.New()
		NoFile = "VERIFICATION-04"
		FileName = NoFile + "-Report-VerifikasiRealPin-Summary " + " - " + request.Pernr + " - " + request.StartDate + " - " + request.EndDate + " - " + UUID.String()
	} else {
		return nil, "", "", fmt.Errorf("error, Report Template belum terdaftar")
	}

	return ColumnName, FileName, NoFile, nil
}

func containsString(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}
