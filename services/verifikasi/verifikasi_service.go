package verifikasi

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"riskmanagement/lib"
	modelTematik "riskmanagement/models/data_tematik"
	requestFile "riskmanagement/models/files"
	models "riskmanagement/models/verifikasi"
	datatematik "riskmanagement/repository/data_tematik"
	fileRepo "riskmanagement/repository/files"
	verifikasi "riskmanagement/repository/verifikasi"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"gitlab.com/golang-package-library/logger"
	minio "gitlab.com/golang-package-library/minio"
	"gorm.io/gorm"
)

// var (
// 	timeNow = lib.GetTimeNow("timestime")
// )

type VerifikasiDefinition interface {
	WithTrx(trxHandle *gorm.DB) VerifikasiService
	GetAll() (responses []models.VerifikasiResponse, err error)
	GetListData() (responses []models.VerifikasiList, err error)
	GetOne(id int64) (responses models.VerifikasiResponseGetOne, status bool, err error)
	GetDataWithPagination(request models.VerifikasiPagination) (responses []models.VerifikasiList, pagination lib.Pagination, err error)
	FilterVerifikasi(request models.VerifikasiFilterRequest) (responses []models.VerifikasiList, pagination lib.Pagination, err error)
	StoreDraft(request models.VerifikasiRequest) (status bool, err error)
	Delete(request *models.VerifikasiRequestUpdateMaintain) (response bool, err error)
	KonfirmSave(request *models.VerifikasiUpdateMaintain) (response bool, err error)
	DeleteLampiranVerifikasi(request *models.VerifikasiFileRequest) (status bool, err error)
	UpdateAllVerifikasi(request *models.VerifikasiRequestMaintain) (status bool, message string, err error)
	GetNoPelaporan(request models.NoPalaporanRequest) (responses []models.NoPelaporanResponse, err error)
	GetLastID() (responses []models.VerifikasiLastIDResponse, err error)
	FilterReport(request models.VerifikasiFilterReport) (responses []models.VerifikasiReportResponse, pagination lib.Pagination, err error)
	StoreSimpan(request models.VerifikasiRequest) (status bool, message string, err error)
	DeleteRiskControl(request models.VerifikasiRiskControl) (response bool, err error)

	// Versioning 1.0.0.1 by panji 31/08/2023
	DeleteAnomaliByID(request *models.VerifikasiAnomaliData) (response bool, err error)

	// VerifikasiReportFilter(request models.VerifikasiFilterReportRequest) (responses []models.VerifikasiFilterReportResponse, totalRows int, err error)
	VerifikasiReportFilter(request models.VerifikasiFilterReportRequest) (responses []models.VerifikasiFilterReportResponse, totalRows int64, err error)
	VerifikasiReportFilterComplete(request models.VerifikasiFilterReportRequest) (responses []models.VerifikasiFilterReportCompleteResponse, totalRows int64, err error)

	VerifikasiReportDetail(request models.VerifikasiReportDetailRequest) (responsesDetail models.VerifikasiReportDetailResponse, err error)
	RiskControlByVerificationId(request models.DataRiskControlRequest) (responses []models.DataRiskIndicatorResponse, totalRows int64, err error)

	VerifikasiReportWithWeaknessOnlyFilter(request models.VerifikasiFilterReportRequest) (responses []models.VerifikasiFilterReportWeaknessOnlyResponse, totalRows int64, err error)
	VerifikasiReportWithNonWeaknessOnlyFilter(request models.VerifikasiFilterReportRequest) (responses []models.VerifikasiFilterReportNonWeaknessOnlyResponse, totalRows int64, err error)

	GetRiskIndicatorAsMateri(request models.VerifikasiFilterReportRequest) (responses []models.GetRiskIndicatorAsMateriResponse, err error)

	VerificationReportByUkerFilter(request models.VerificationFilterReportByUkerRequest) (responses []models.VerificationFilterReportByUkerResponse, totalRows int, err error)
	VerificationReportFilterByUkerComplete(request models.VerificationFilterReportByUkerRequest) (responses []models.VerificationFilterByUkerReportCompleteResponse, totalRows int64, err error)

	VerifikasiReportByFraudIndicatorFilter(request models.VerificationFilterReportByUkerRequest) (responses []models.VerificationFilterReportByFraudIndicatorResponse, totalRows int64, err error)
	VerificationReportFilterByFraudIndicatorComplete(request models.VerificationFilterReportByUkerRequest) (responses []models.VerifikasiFilterReportCompleteResponse, totalRows int, err error)

	//add 23 Feb 2023 By Panji
	VerifikasiReportMateriList(request models.VerifikasiMateriRequest) (responses []models.VerifikasiDetailMateriResponse, err error)

	VerificationReportUkerByAllActivity(request models.VerificationFilterReportByUkerRequest) (responses models.VerifikasiReportAllUker, totalRows int, err error)
	VerificationReportUkerByAllActivityComplete(request models.VerificationFilterReportByUkerRequest) (responses []map[string]interface{}, totalRows int, err error)

	VerificationReportUkerByAllActivityCompleteWithRiskIssue(request models.VerificationFilterReportByUkerRequest) (responses []map[string]interface{}, totalRows int, err error)
	VerifikasiReportList(request models.VerifikasiReportListRequest) (responses []models.VerifikasiReportListResponse, totalRows int, err error)
	RptRekapitulasiBCV(request models.RptRekapitulasiBCVRequest) (responses []models.RptRekapitulasiBCVResponse, totalRows int, err error)

	//RptRekomendasiRisk
	RptRekomendasiRisk(request models.RptRekomendasiRiskRequest) (responses []models.RptRekomendasiRiskResponse, totalRows int, err error)

	//ValidasiVerifikasi
	ValidasiVerifikasi(request models.ValidasiVerifikasiRequest) (responses []models.ValidasiVerifikasiResponse, totalRows int, err error)
	AcceptValidasi(request *models.AcceptValidasiRequest) (responses bool, err error)
	UpdateStatusVerifikasi(request *models.UpdateStatusVerifikasi) (responses bool, err error)
	RejectValidasi(request *models.RejectValidasiRequest) (response bool, err error)

	//RTL Indikasi Fraud
	GetRtlIndikasiFraud(request models.ReqRtlIndikasiFraud) (responses models.RtlIndikasiFraudResponse, totalRows int, err error)
	ValidasiVerifikasiDetailData(request models.VerifikasiReportDetailRequest) (responsesDetail models.ValidasiVerifikasiDetailedResponse, err error)

	// #Batch 3
	GetRekomendasiTindakLanjut(request models.RTLRequest) (responses []models.RTLResponses, err error)
	DeletePenyebabKejadian(id int64) (response bool, err error)
	VerifikasiSummaryRpt(request models.SummaryVerifikasiRequest) (responses []models.SummaryVerifikasiResponse, totalRows int, err error)
	VerifikasiFrekuensiRpt(request models.FrekuensiVerifikasiRequest) (responses []models.FrekuensiVerifikasiResponse, totalRows int, err error)
}

type VerifikasiService struct {
	db                         lib.Database
	minio                      minio.Minio
	logger                     logger.Logger
	verifikasiRepo             verifikasi.VerifikasiDefinition
	verifikasiAnomali          verifikasi.VerifikasiAnomaliDefinition
	verifikasiFile             verifikasi.VerifikasiFilesDefinition
	verifikasiPIC              verifikasi.VerifikasiPICDefinition
	verifikasiRiskControl      verifikasi.VerifikasiRiskControlDefinition
	verifikasiAnomaliKRID      verifikasi.VerifikasiAnomaliDataKRIDDefinition
	verifikasiQuestioner       verifikasi.VerifikasiQuestionnerDefinition
	verifikasiPenyababKejadian verifikasi.VerifikasiPenyababKejadianDefinition
	verifikasiUsulanPerbaikan  verifikasi.VerifikasiUsulanPerbaikanDefinition
	verifikasiDataTematik      verifikasi.VerifikasiDataTematikDefinition
	fileRepo                   fileRepo.FilesDefinition
	dataTematikRepo            datatematik.DataTematikDefinition
}

func NewVerifikasiService(
	db lib.Database,
	minio minio.Minio,
	logger logger.Logger,
	verifikasiRepo verifikasi.VerifikasiDefinition,
	verifikasiAnomali verifikasi.VerifikasiAnomaliDefinition,
	verifikasiFile verifikasi.VerifikasiFilesDefinition,
	verifikasiPIC verifikasi.VerifikasiPICDefinition,
	verifikasiRiskControl verifikasi.VerifikasiRiskControlDefinition,
	verifikasiAnomaliKRID verifikasi.VerifikasiAnomaliDataKRIDDefinition,
	verifikasiQuestioner verifikasi.VerifikasiQuestionnerDefinition,
	verifikasiPenyababKejadian verifikasi.VerifikasiPenyababKejadianDefinition,
	verifikasiUsulanPerbaikan verifikasi.VerifikasiUsulanPerbaikanDefinition,
	verifikasiDataTematik verifikasi.VerifikasiDataTematikDefinition,
	fileRepo fileRepo.FilesDefinition,
	dataTematikRepo datatematik.DataTematikDefinition,
) VerifikasiDefinition {
	return VerifikasiService{
		db:                         db,
		minio:                      minio,
		logger:                     logger,
		verifikasiRepo:             verifikasiRepo,
		verifikasiAnomali:          verifikasiAnomali,
		verifikasiFile:             verifikasiFile,
		verifikasiPIC:              verifikasiPIC,
		verifikasiRiskControl:      verifikasiRiskControl,
		verifikasiAnomaliKRID:      verifikasiAnomaliKRID,
		fileRepo:                   fileRepo,
		verifikasiQuestioner:       verifikasiQuestioner,
		verifikasiPenyababKejadian: verifikasiPenyababKejadian,
		verifikasiUsulanPerbaikan:  verifikasiUsulanPerbaikan,
		verifikasiDataTematik:      verifikasiDataTematik,
		dataTematikRepo:            dataTematikRepo,
	}
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// DeleteRiskControl implements VerifikasiDefinition
func (verif VerifikasiService) DeleteRiskControl(request models.VerifikasiRiskControl) (response bool, err error) {
	err = verif.verifikasiRiskControl.Delete(request.ID)

	return true, err
}

// Delete implements VerifikasiDefinition
func (verifikasi VerifikasiService) Delete(request *models.VerifikasiRequestUpdateMaintain) (response bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	tx := verifikasi.db.DB.Begin()

	getOneVerifikasi, exist, err := verifikasi.GetOne(request.ID)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		tx.Rollback()
		return false, err
	}

	dataVerif, _ := verifikasi.verifikasiRepo.GetOne(request.ID)

	if dataVerif.SumberData == "Tematik" {
		dataTematik, err := verifikasi.verifikasiDataTematik.GetDataTematik(request.ID)
		if err != nil {
			verifikasi.logger.Zap.Error(err)
			tx.Rollback()
			return false, err
		}

		for _, tematik := range dataTematik {
			namaTable := fmt.Sprintf("lampiran_rap_%s_%s", strconv.FormatInt(dataVerif.RiskIssueID, 10), strconv.FormatInt(dataVerif.RiskIndicatorID, 10))
			fmt.Println("table lampiran =>", namaTable)

			var data map[string]interface{}
			err := json.Unmarshal([]byte(tematik.ColumnsData), &data)
			if err != nil {
				verifikasi.logger.Zap.Error("Error unmarshaling JSON: %s", err)
			}

			// fmt.Println("MARSAL DATA =>", reflect.TypeOf(data["id"]))

			// id, _ := strconv.ParseInt(data["id"], 10, 64)

			status, err := verifikasi.dataTematikRepo.UpdateStatusDataSample(&modelTematik.RequestUpdate{
				NamaTable:    namaTable,
				Id:           int64(data["id"].(float64)),
				Status:       "-",
				NoVerifikasi: "",
			})

			fmt.Println("status =>", status)

			if err != nil {
				verifikasi.logger.Zap.Error(err)
				tx.Rollback()
				return false, err
			}
		}

	}

	UpdateDataVerifikasi := &models.VerifikasiUpdateDelete{
		ID:            request.ID,
		LastMakerID:   request.LastMakerID,
		LastMakerDesc: request.LastMakerDesc,
		LastMakerDate: &timeNow,
		Status:        "02b", //selesai
		Action:        "UpdateDelete",
		Deleted:       true,
		UpdatedAt:     &timeNow,
	}

	fmt.Println(UpdateDataVerifikasi)

	include := []string{
		"last_maker_id",
		"last_maker_desc",
		"last_maker_date",
		"deleted",
		"status",
		"action",
		"updated_at",
	}

	_, err = verifikasi.verifikasiRepo.Delete(UpdateDataVerifikasi, include, tx)
	if err != nil {
		tx.Rollback()
		verifikasi.logger.Zap.Error(err)
		return false, err
	}

	if exist {
		fmt.Println("getOneVerif", getOneVerifikasi)
		tx.Commit()
		return true, err
	}

	return false, err
}

// GetAll implements VerifikasiDefinition
func (verifikasi VerifikasiService) GetAll() (responses []models.VerifikasiResponse, err error) {
	return verifikasi.verifikasiRepo.GetAll()
}

// GetListData implements VerifikasiDefinition
func (verifikasi VerifikasiService) GetListData() (responses []models.VerifikasiList, err error) {
	return verifikasi.verifikasiRepo.GetListData()
}

// GetOne implements VerifikasiDefinition
func (verifikasi VerifikasiService) GetOne(id int64) (responses models.VerifikasiResponseGetOne, status bool, err error) {
	dataVerif, err := verifikasi.verifikasiRepo.GetOne(id)
	fmt.Println(dataVerif)

	if dataVerif.ID != 0 {
		fmt.Println("bukan 0")

		data_anomali, err := verifikasi.verifikasiAnomali.GetOneByVerifikasi(dataVerif.ID)
		if err != nil {
			verifikasi.logger.Zap.Error(err)
			return responses, false, err
		}

		questionner, err := verifikasi.verifikasiQuestioner.GetOneByVerifikasi(dataVerif.ID)
		if err != nil {
			verifikasi.logger.Zap.Error(err)
			return responses, false, err
		}

		files, err := verifikasi.verifikasiFile.GetOneFileByID(dataVerif.ID)
		if err != nil {
			verifikasi.logger.Zap.Error(err)
			return responses, false, err
		}

		pic_tindak_lanjut, err := verifikasi.verifikasiPIC.GetOneByPIC(dataVerif.ID)
		if err != nil {
			verifikasi.logger.Zap.Error(err)
			return responses, false, err
		}

		risk_control, err := verifikasi.verifikasiRiskControl.GetOneDataByID(dataVerif.ID)
		if err != nil {
			verifikasi.logger.Zap.Error(err)
			return responses, false, err
		}

		data_anomali_krid, err := verifikasi.verifikasiAnomaliKRID.GetOneByVerifikasi(dataVerif.ID)
		if err != nil {
			verifikasi.logger.Zap.Error(err)
			return responses, false, err
		}

		penyebabKejadian, err := verifikasi.verifikasiPenyababKejadian.GetData(dataVerif.ID)
		if err != nil {
			verifikasi.logger.Zap.Error(err)
			return responses, false, err
		}

		usulan, err := verifikasi.verifikasiUsulanPerbaikan.GetData(dataVerif.ID)
		if err != nil {
			verifikasi.logger.Zap.Error(err)
			return responses, false, err
		}

		dataTematik, err := verifikasi.verifikasiDataTematik.GetDataTematik(dataVerif.ID)
		if err != nil {
			verifikasi.logger.Zap.Error(err)
			return responses, false, err
		}

		responses = models.VerifikasiResponseGetOne{
			ID:                        dataVerif.ID,
			NoPelaporan:               dataVerif.NoPelaporan,
			REGION:                    dataVerif.REGION,
			RGDESC:                    dataVerif.RGDESC,
			MAINBR:                    dataVerif.MAINBR,
			MBDESC:                    dataVerif.MBDESC,
			BRANCH:                    dataVerif.BRANCH,
			BRDESC:                    dataVerif.BRDESC,
			ActivityID:                dataVerif.ActivityID,
			SubActivityID:             dataVerif.SubActivityID,
			ProductID:                 dataVerif.ProductID,
			RiskIssueID:               dataVerif.RiskIssueID,
			RiskIssue:                 dataVerif.RiskIssue,
			RiskIndicatorID:           dataVerif.RiskIndicatorID,
			RiskIndicator:             dataVerif.RiskIndicator,
			SumberData:                dataVerif.SumberData,
			ApplicationID:             dataVerif.ApplicationID,
			HasilVerifikasi:           dataVerif.HasilVerifikasi,
			KunjunganNasabah:          dataVerif.KunjunganNasabah,
			Perbaikan:                 dataVerif.Perbaikan,
			IndikasiFraud:             dataVerif.IndikasiFraud,
			TerdapatKerugianFinansial: dataVerif.TerdapatKerugianFinansial,
			JenisKerugianFinansial:    dataVerif.JenisKerugianFinansial,
			JumlahPerkiraanKerugian:   dataVerif.JumlahPerkiraanKerugian,
			JenisKerugianNonFinansial: dataVerif.JenisKerugianNonFinansial,
			JenisRekomendasi:          dataVerif.JenisRekomendasi,
			RekomendasiTindakLanjut:   dataVerif.RekomendasiTindakLanjut,
			RencanaTindakLanjut:       dataVerif.RencanaTindakLanjut,
			RiskTypeID:                dataVerif.RiskTypeID,
			TanggalDitemukan:          dataVerif.TanggalDitemukan,
			TanggalMulaiRTL:           dataVerif.TanggalMulaiRTL,
			TanggalTargetSelesai:      dataVerif.TanggalTargetSelesai,
			MakerID:                   dataVerif.MakerID,
			MakerDesc:                 dataVerif.MakerDesc,
			MakerDate:                 dataVerif.MakerDate,
			LastMakerID:               dataVerif.LastMakerID,
			LastMakerDesc:             dataVerif.LastMakerDesc,
			LastMakerDate:             dataVerif.LastMakerDate,
			Status:                    dataVerif.Status,
			Action:                    dataVerif.Action,
			Deleted:                   dataVerif.Deleted,
			DataAnomali:               data_anomali,
			DataAnomaliKRID:           data_anomali_krid,
			PICTindakLanjut:           pic_tindak_lanjut,
			Questionner:               questionner,
			Files:                     files,
			RiskControl:               risk_control,
			PenyababKejadian:          penyebabKejadian,
			AdaUsulanPerbaikan:        dataVerif.AdaUsulanPerbaikan,
			UsulanPerbaikan:           usulan,
			UpdatedAt:                 dataVerif.UpdatedAt,
			CreatedAt:                 dataVerif.CreatedAt,
			SampleDataTeamatik:        dataTematik,
		}

		return responses, true, err
	}

	return responses, false, err

}

// Store implements VerifikasiDefinition
func (verifikasi VerifikasiService) StoreDraft(request models.VerifikasiRequest) (status bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	tx := verifikasi.db.DB.Begin()

	//input data verifikasi
	reqVerif := &models.Verifikasi{
		NoPelaporan:   request.NoPelaporan,
		REGION:        request.REGION,
		RGDESC:        request.RGDESC,
		MAINBR:        request.MAINBR,
		MBDESC:        request.MBDESC,
		BRANCH:        request.BRANCH,
		BRDESC:        request.BRDESC,
		ActivityID:    request.ActivityID,
		SubActivityID: request.SubActivityID,
		ProductID:     request.ProductID,
		RiskIssueID:   request.RiskIssueID,
		RiskIssue:     request.RiskIssue,
		// RiskIssueOther:            request.RiskIssueOther,
		RiskIndicatorID: request.RiskIndicatorID,
		RiskIndicator:   request.RiskIndicator,
		// RiskIndicatorOther:        request.RiskIndicatorOther,
		SumberData:                request.SumberData,
		ApplicationID:             request.ApplicationID,
		HasilVerifikasi:           request.HasilVerifikasi,
		KunjunganNasabah:          *request.KunjunganNasabah,
		Perbaikan:                 request.Perbaikan,
		IndikasiFraud:             request.IndikasiFraud,
		JenisKerugianFinansial:    request.JenisKerugianFinansial,
		JumlahPerkiraanKerugian:   request.JumlahPerkiraanKerugian,
		JenisKerugianNonFinansial: request.JenisKerugianNonFinansial,
		RekomendasiTindakLanjut:   request.RekomendasiTindakLanjut,
		RencanaTindakLanjut:       request.RencanaTindakLanjut,
		RiskTypeID:                request.RiskTypeID,
		TanggalDitemukan:          request.TanggalDitemukan,
		TanggalMulaiRTL:           request.TanggalMulaiRTL,
		TanggalTargetSelesai:      request.TanggalTargetSelesai,
		MakerID:                   request.MakerID,
		MakerDesc:                 request.MakerDesc,
		MakerDate:                 &timeNow,
		LastMakerID:               request.LastMakerID,
		LastMakerDesc:             request.LastMakerDesc,
		LastMakerDate:             &timeNow,
		Status:                    "01a",
		Action:                    "Draft",
		Deleted:                   false,
		CreatedAt:                 &timeNow,
	}

	dataVerif, err := verifikasi.verifikasiRepo.Store(reqVerif, tx)

	if err != nil {
		tx.Rollback()
		verifikasi.logger.Zap.Error(err)
		return false, err
	}
	// fmt.Println("data verifikasi : ", dataVerif)
	//end data verifikasi

	//Begin Input data anomali
	if request.SumberData == "KRID" {
		if len(request.DataAnomaliKRID) != 0 {
			for _, value := range request.DataAnomaliKRID {
				// objectString := []byte(value.Object)

				fmt.Println("object Stringg =======>s", value)
				if value.Periode != "" && value.Object != "" {
					_, err = verifikasi.verifikasiAnomaliKRID.Store(&models.VerifikasiAnomaliDataKRID{
						VerifikasiID: dataVerif.ID,
						Periode:      value.Periode,
						Object:       value.Object,
						Status:       false,
					}, tx)

					if err != nil {
						tx.Rollback()
						verifikasi.logger.Zap.Error(err)
						return false, err
					}
				}

			}
		} else {
			tx.Rollback()
			verifikasi.logger.Zap.Error(err)
			return false, err
		}
	} else {
		if len(request.DataAnomali) != 0 {
			for _, value := range request.DataAnomali {
				fmt.Println("object Stringg =======>s", value)

				_, err = verifikasi.verifikasiAnomali.Store(&models.VerifikasiAnomaliData{
					VerifikasiID:    dataVerif.ID,
					TanggalKejadian: value.TanggalKejadian,
					NomorRekening:   value.NomorRekening,
					Nominal:         value.Nominal,
					Keterangan:      value.Keterangan,
				}, tx)

				if err != nil {
					tx.Rollback()
					verifikasi.logger.Zap.Error(err)
					return false, err
				}
			}
		} else {
			tx.Rollback()
			verifikasi.logger.Zap.Error(err)
			return false, err
		}
	}
	//End Input data anomali

	// Add Input Questionner by Panji 30 07 2023
	if len(request.Questionner) != 0 {
		for _, value := range request.Questionner {
			_, err = verifikasi.verifikasiQuestioner.Store(&models.VerifikasiQuestionner{
				// ID:           va,
				VerifikasiID: dataVerif.ID,
				Questionner:  value.Questionner,
				DataSumber:   value.DataSumber,
				Checker:      value.Checker,
				Signer:       value.Signer,
				JenisFraud:   value.JenisFraud,
			}, tx)

			if err != nil {
				tx.Rollback()
				verifikasi.logger.Zap.Error(err)
				return false, err
			}
		}
	} else {
		tx.Rollback()
		verifikasi.logger.Zap.Error(err)
		return false, err
	}

	// if request.Perbaikan == true {
	//Begin Input Kelemahan Kontrol
	if len(request.RiskControl) != 0 {
		for _, value := range request.RiskControl {
			_, err = verifikasi.verifikasiRiskControl.Store(&models.VerifikasiRiskControl{
				VerifikasiId:  dataVerif.ID,
				RiskControlID: value.RiskControlID,
				RiskControl:   value.RiskControl,
			}, tx)

			if err != nil {
				tx.Rollback()
				verifikasi.logger.Zap.Error(err)
				return false, err
			}
		}
	} else {
		tx.Rollback()
		verifikasi.logger.Zap.Error(err)
		return false, err
	}
	//End Input Kelemahan Kontrol

	//Begin Input data PIC
	if len(request.PICTindakLanjut) != 0 {
		for _, value := range request.PICTindakLanjut {
			_, err = verifikasi.verifikasiPIC.Store(&models.VerifikasiPICTindakLanjut{
				VerifikasiID:          dataVerif.ID,
				PICID:                 value.PICID,
				PICDetail:             value.PICDetail,
				TanggalTindakLanjut:   value.TanggalTindakLanjut,
				DeskripsiTindakLanjut: value.DeskripsiTindakLanjut,
				Status:                "0",
			}, tx)

			if err != nil {
				tx.Rollback()
				verifikasi.logger.Zap.Error(err)
				return false, err
			}
		}
	} else {
		tx.Rollback()
		verifikasi.logger.Zap.Error(err)
		return false, err
	}
	//End Input data PIC

	//Begin Input Lampiran
	bucket := os.Getenv("BUCKET_NAME")

	if len(request.Files) != 0 {
		for _, value := range request.Files {
			UUID := uuid.New()
			var destinationPath string
			if value.Filename != "" {
				bucketExist := verifikasi.minio.BucketExist(verifikasi.minio.Client(), bucket)

				pathSplit := strings.Split(value.Path, "/")
				sourcePath := fmt.Sprint(value.Path)
				destinationPath = pathSplit[1] + "/" +
					lib.GetTimeNow("year") + "/" +
					lib.GetTimeNow("month") + "/" +
					lib.GetTimeNow("day") + "/" +
					UUID.String() + "/" +
					// pathSplit[2] + "/" +
					value.Filename

				// newPath := "verifikasi/" +
				// 	lib.GetTimeNow("year") + "/" +
				// 	lib.GetTimeNow("month") + "/" +
				// 	lib.GetTimeNow("day")

				// destinationPath = newPath + "/" + value.Filename

				if bucketExist {
					fmt.Println("Exist")
					fmt.Println(bucket)
					fmt.Println(sourcePath)
					fmt.Println(destinationPath)
					uploaded := verifikasi.minio.CopyObject(verifikasi.minio.Client(), bucket, sourcePath, bucket, destinationPath)
					// uploaded := verifikasi.minio.PutObject(verifikasi.minio.MinioClient, bucket, destinationPath, sourcePath)

					fmt.Println(uploaded)
				} else {
					fmt.Println("Not Exist")
					fmt.Println(bucket)
					fmt.Println(sourcePath)
					fmt.Println(destinationPath)
					verifikasi.minio.MakeBucket(verifikasi.minio.Client(), bucket, "")
					uploaded := verifikasi.minio.CopyObject(verifikasi.minio.Client(), bucket, sourcePath, bucket, destinationPath)
					// uploaded := verifikasi.minio.PutObject(verifikasi.minio.MinioClient, bucket, destinationPath, sourcePath)

					fmt.Println(uploaded)
				}
			}

			files, err := verifikasi.fileRepo.Store(&requestFile.Files{
				Filename:  value.Filename,
				Path:      destinationPath,
				Extension: value.Extension,
				Size:      value.Size,
				CreatedAt: &timeNow,
			}, tx)

			if err != nil {
				tx.Rollback()
				verifikasi.logger.Zap.Error(err)
				return false, err
			}

			_, err = verifikasi.verifikasiFile.Store(&models.VerifikasiFiles{
				VerifikasiID: dataVerif.ID,
				FilesID:      files.ID,
				// CreatedAt:    &timeNow,
			}, tx)

			if err != nil {
				tx.Rollback()
				verifikasi.logger.Zap.Error(err)
				return false, err
			}
		}
	} else {
		tx.Rollback()
		verifikasi.logger.Zap.Error(err)
		return false, err
	}

	//End Input Lampiran
	// }

	tx.Commit()
	return true, err
}

// WithTrx implements VerifikasiDefinition
func (verifikasi VerifikasiService) WithTrx(trxHandle *gorm.DB) VerifikasiService {
	verifikasi.verifikasiRepo = verifikasi.verifikasiRepo.WithTrx(trxHandle)
	return verifikasi
}

// DeleteLampiranVerifikasi implements VerifikasiDefinition
func (verifikasi VerifikasiService) DeleteLampiranVerifikasi(request *models.VerifikasiFileRequest) (status bool, err error) {
	bucket := os.Getenv("BUCKET_NAME")

	ok := verifikasi.minio.RemoveObject(verifikasi.minio.Client(), bucket, request.Path)

	if !ok {
		return false, err
	} else {
		tx := verifikasi.db.DB.Begin()
		err = verifikasi.fileRepo.Delete(request.FilesID, tx)
		if err != nil {
			tx.Rollback()
			verifikasi.logger.Zap.Error(err)
			return false, err
		}

		err = verifikasi.verifikasiRepo.DeleteLampiranVerifikasi(request.VerifikasiLampiranID, tx)
		if err != nil {
			tx.Rollback()
			verifikasi.logger.Zap.Error(err)
			return false, err
		}

		tx.Commit()

	}

	return true, err

}

// KonfirmSave implements VerifikasiDefinition
func (verifikasi VerifikasiService) KonfirmSave(request *models.VerifikasiUpdateMaintain) (response bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	tx := verifikasi.db.DB.Begin()

	getOneVerifikasi, exist, err := verifikasi.GetOne(request.ID)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		tx.Rollback()
		return false, err
	}

	UpdateDataVerifikasi := &models.VerifikasiUpdateMaintain{
		ID:            request.ID,
		LastMakerID:   request.LastMakerID,
		LastMakerDesc: request.LastMakerDesc,
		LastMakerDate: &timeNow,
		Status:        "02b", //selesai
		Action:        "Selesai",
		UpdatedAt:     &timeNow,
	}

	include := []string{
		"last_maker_id",
		"last_maker_desc",
		"last_maker_date",
		"status",
		"action",
		"updated_at",
	}

	_, err = verifikasi.verifikasiRepo.KonfirmSave(UpdateDataVerifikasi, include, tx)
	if err != nil {
		tx.Rollback()
		verifikasi.logger.Zap.Error(err)
		return false, err
	}

	if exist {
		fmt.Println("getOneVerif", getOneVerifikasi)
		tx.Commit()
		return true, err
	}
	return false, err

}

// UpdateAllVerifikasi implements VerifikasiDefinition
func (verifikasi VerifikasiService) UpdateAllVerifikasi(request *models.VerifikasiRequestMaintain) (status bool, message string, err error) {
	timeNow := lib.GetTimeNow("timestime")

	tx := verifikasi.db.DB.Begin()

	fmt.Println("Hai saya disini !!")

	updateVerifikasi := &models.VerifikasiUpdateAll{
		ID:                        request.ID,
		NoPelaporan:               request.NoPelaporan,
		REGION:                    request.REGION,
		RGDESC:                    request.RGDESC,
		MAINBR:                    request.MAINBR,
		MBDESC:                    request.MBDESC,
		BRANCH:                    request.BRANCH,
		BRDESC:                    request.BRDESC,
		ActivityID:                request.ActivityID,
		SubActivityID:             request.SubActivityID,
		ProductID:                 request.ProductID,
		RiskIssueID:               request.RiskIssueID,
		RiskIssue:                 request.RiskIssue,
		RiskIndicatorID:           request.RiskIndicatorID,
		RiskIndicator:             request.RiskIndicator,
		SumberData:                request.SumberData,
		ApplicationID:             request.ApplicationID,
		HasilVerifikasi:           request.HasilVerifikasi,
		KunjunganNasabah:          request.KunjunganNasabah,
		Perbaikan:                 request.Perbaikan,
		IndikasiFraud:             request.IndikasiFraud,
		TerdapatKerugianFinansial: request.TerdapatKerugianFinansial,
		JenisKerugianFinansial:    request.JenisKerugianFinansial,
		JumlahPerkiraanKerugian:   request.JumlahPerkiraanKerugian,
		JenisKerugianNonFinansial: request.JenisKerugianNonFinansial,
		JenisRekomendasi:          request.JenisRekomendasi,
		RekomendasiTindakLanjut:   request.RekomendasiTindakLanjut,
		RencanaTindakLanjut:       request.RencanaTindakLanjut,
		RiskTypeID:                request.RiskTypeID,
		AdaUsulanPerbaikan:        request.AdaUsulanPerbaikan,
		TanggalDitemukan:          request.TanggalDitemukan,
		TanggalMulaiRTL:           request.TanggalMulaiRTL,
		TanggalTargetSelesai:      request.TanggalTargetSelesai,
		LastMakerID:               request.LastMakerID,
		LastMakerDesc:             request.LastMakerDesc,
		LastMakerDate:             &timeNow,
		Status:                    request.Status,
		Action:                    request.Action,
		StatusIndikasiFraud:       request.StatusIndikasiFraud,
		ActionIndikasiFraud:       request.ActionIndikasiFraud,
		UpdatedAt:                 &timeNow,
	}

	include := []string{
		"no_pelaporan",
		"REGION",
		"RGDESC",
		"MAINBR",
		"MBDESC",
		"BRANCH",
		"BRDESC",
		"activity_id",
		"sub_activity_id",
		"product_id",
		"risk_issue_id",
		"risk_issue",
		"risk_indicator_id",
		"risk_indicator",
		"sumber_data",
		"application_id",
		"hasil_verifikasi",
		"kunjungan_nasabah",
		"perbaikan",
		"indikasi_fraud",
		"terdapat_kerugian_finansial",
		"jenis_kerugian_finansial",
		"jumlah_kerugian_finansial",
		"jenis_rekomedasi",
		"rekomendasi_tindak_lanjut",
		"rencana_tindak_lanjut",
		"risk_type_id",
		"ada_usulan_perbaikan",
		"tanggal_ditemukan",
		"tanggal_mulai_rtl",
		"tanggal_target_selesai",
		"last_maker_id",
		"last_maker_desc",
		"last_maker_date",
		"status",
		"action",
		"status_indikasi_fraud",
		"action_indikasi_fraud",
		"updated_at",
	}

	_, err = verifikasi.verifikasiRepo.UpdateAllVerifikasi(updateVerifikasi, include, tx)

	if err != nil {
		tx.Rollback()
		verifikasi.logger.Zap.Error(err)
		message = "Error, transaction database !"
		return false, message, err
	}

	//Update & add Data Anomali
	// versioning 1.0.0.1 By Panji 31/08/2023
	if request.SumberData == "KRID" {
		if len(request.DataAnomaliKRID) != 0 {
			fmt.Println(" ======> KRID")
			for _, value := range request.DataAnomaliKRID {
				if value.Periode != "" && value.Object != "" {
					updateAnomali := &models.VerifikasiAnomaliDataKRID{
						ID:           value.ID,
						VerifikasiID: request.ID,
						Periode:      value.Periode,
						Object:       value.Object,
						Status:       false,
					}

					_, err = verifikasi.verifikasiAnomaliKRID.Store(updateAnomali, tx)

					if err != nil {
						tx.Rollback()
						verifikasi.logger.Zap.Error(err)
						message = "Error, transaction database !"
						return false, message, err
					}
				}
			}
		} else {
			tx.Rollback()
			verifikasi.logger.Zap.Error(err)
			message = "Error, Data Sample Null !"
			return false, message, err
		}
	} else if request.SumberData == "Non KRID" {
		if len(request.DataAnomali) != 0 {
			fmt.Println(" ======> Non KRID")
			for _, value := range request.DataAnomali {
				updateAnomali := &models.VerifikasiAnomaliData{
					ID:              value.ID,
					VerifikasiID:    request.ID,
					TanggalKejadian: value.TanggalKejadian,
					NomorRekening:   value.NomorRekening,
					Nominal:         value.Nominal,
					Keterangan:      value.Keterangan,
				}

				fmt.Println("update data =>", updateAnomali)

				_, err = verifikasi.verifikasiAnomali.Update(updateAnomali, tx)

				if err != nil {
					tx.Rollback()
					verifikasi.logger.Zap.Error(err)
					message = "Error, transaction database !"
					return false, message, err
				}
			}
		} else {
			tx.Rollback()
			verifikasi.logger.Zap.Error(err)
			message = "Error, Data Sample Null !"
			return false, message, err
		}
	} else if request.SumberData == "Tematik" {
		if len(request.SampleDataTeamatik) != 0 {
			for _, value := range request.SampleDataTeamatik {
				updateTematik := &models.VerifikasiDataTematik{
					ID:           value.ID,
					VerifikasiId: request.ID,
					Periode:      value.Periode,
					Columns:      value.Columns,
					ColumnsData:  value.ColumnsData,
					Status:       false,
				}

				_, err = verifikasi.verifikasiDataTematik.Store(updateTematik, tx)
				if err != nil {
					tx.Rollback()
					verifikasi.logger.Zap.Error(err)
					message = "Error, transaction database !"
					return false, message, err
				}
			}
		} else {
			if err != nil {
				tx.Rollback()
				verifikasi.logger.Zap.Error(err)
				message = "Error, Data Sample Null !"
				return false, message, err
			}
		}
	}

	//Update & add Data Anomali

	//#update & add Anomali

	// Update & add Questionner
	if len(request.Questionner) != 0 {
		for _, value := range request.Questionner {
			_, err = verifikasi.verifikasiQuestioner.Store(&models.VerifikasiQuestionner{
				ID:                value.ID,
				VerifikasiID:      request.ID,
				Questionner:       value.Questionner,
				DataSumber:        value.DataSumber,
				Checker:           value.Checker,
				Signer:            value.Signer,
				ApprovalOrd:       value.ApprovalOrd,
				JenisFraud:        value.JenisFraud,
				StatusValidasiRmc: value.StatusValidasiRmc,
			}, tx)

			if err != nil {
				tx.Rollback()
				verifikasi.logger.Zap.Error(err)
				message = "Error, transaction database !"
				return false, message, err
			}
		}
	} else {
		tx.Rollback()
		verifikasi.logger.Zap.Error(err)
		message = "Error, Questioner Null !"
		return false, message, err
	}

	// if request.Perbaikan == true {

	//Update & add Risk Control
	if len(request.RiskControl) != 0 {
		for _, value := range request.RiskControl {
			updateRiskControl := &models.VerifikasiRiskControl{
				ID:            value.ID,
				VerifikasiId:  request.ID,
				RiskControlID: value.RiskControlID,
				RiskControl:   value.RiskControl,
			}
			_, err = verifikasi.verifikasiRiskControl.Store(updateRiskControl, tx)

			if err != nil {
				tx.Rollback()
				verifikasi.logger.Zap.Error(err)
				message = "Error, transaction database !"
				return false, message, err
			}

		}
	} else {
		if err != nil {
			tx.Rollback()
			verifikasi.logger.Zap.Error(err)
			message = "Error, Data Risk Control Null !"
			return false, message, err
		}
	}
	//#Update & add Risk Control

	//Update & add Data PIC Tindak Lanjut
	if len(request.PICTindakLanjut) != 0 {
		for _, value := range request.PICTindakLanjut {
			updatePIC := &models.VerifikasiPICTindakLanjut{
				ID:                    value.ID,
				VerifikasiID:          request.ID,
				PICID:                 value.PICID,
				PICDetail:             value.PICDetail,
				TanggalTindakLanjut:   value.TanggalTindakLanjut,
				DeskripsiTindakLanjut: value.DeskripsiTindakLanjut,
				Status:                value.Status,
			}

			_, err = verifikasi.verifikasiPIC.Store(updatePIC, tx)

			if err != nil {
				tx.Rollback()
				verifikasi.logger.Zap.Error(err)
				message = "Error, transaction database !"
				return false, message, err
			}

		}
	} else {
		if err != nil {
			tx.Rollback()
			verifikasi.logger.Zap.Error(err)
			message = "Error, transaction database !"
			return false, message, err
		}
	}
	//#Update & add Data PIC Tindak Lanjut

	//#Update Lampiran
	bucket := os.Getenv("BUCKET_NAME")

	if len(request.Files) != 0 {

		err := verifikasi.verifikasiFile.DeleteFilesByID(request.ID, tx)
		if err != nil {
			tx.Rollback()
			verifikasi.logger.Zap.Error(err)
			message = "Error, transaction database !"
			return false, message, err
		}

		for _, value := range request.Files {
			UUID := uuid.New()
			var destinationPath string
			allowedExtensions := []string{".jpg", ".jpeg", ".pdf", ".doc", ".docx", ".xls", ".xlsx"}

			extension := strings.ToLower(filepath.Ext(value.Filename))
			fmt.Println("Extension =>", extension)
			if value.Filename != "" {
				if !contains(allowedExtensions, extension) {
					tx.Rollback()
					verifikasi.logger.Zap.Error(err)
					message = "Error, Invalid file extension !"
					return false, message, err
				}

				bucketExist := verifikasi.minio.BucketExist(verifikasi.minio.Client(), bucket)

				pathSplit := strings.Split(value.Path, "/")
				sourcePath := fmt.Sprintf(value.Path)
				destinationPath = pathSplit[1] + "/" +
					lib.GetTimeNow("year") + "/" +
					lib.GetTimeNow("month") + "/" +
					lib.GetTimeNow("day") + "/" +
					UUID.String() + "/" +
					// pathSplit[2] + "/" +
					value.Filename

				if pathSplit[0] == "tmp" {
					verifikasi.logger.Zap.Info("==========> New Files")

					var uploaded bool
					if bucketExist {
						fmt.Println("Exist")
						fmt.Println(bucket)
						fmt.Println(sourcePath)
						fmt.Println(destinationPath)
						// uploaded := verifikasi.minio.PutObject(verifikasi.minio.MinioClient, bucket, destinationPath, sourcePath)
						uploaded := verifikasi.minio.CopyObject(verifikasi.minio.Client(), bucket, sourcePath, bucket, destinationPath)
						fmt.Println(uploaded)
					} else {
						fmt.Println("Not Exist")
						fmt.Println(bucket)
						fmt.Println(sourcePath)
						fmt.Println(destinationPath)
						verifikasi.minio.MakeBucket(verifikasi.minio.Client(), bucket, "")
						uploaded := verifikasi.minio.CopyObject(verifikasi.minio.Client(), bucket, sourcePath, bucket, destinationPath)
						// uploaded := verifikasi.minio.PutObject(verifikasi.minio.MinioClient, bucket, destinationPath, sourcePath)
						fmt.Println(uploaded)
					}

					if !uploaded {
						tx.Rollback()
						verifikasi.logger.Zap.Error(err)
						message = "Error, Temporary path document not found !"
						return false, message, err
					}

					files, err := verifikasi.fileRepo.Store(&requestFile.Files{
						ID:        value.ID,
						Filename:  value.Filename,
						Path:      destinationPath,
						Extension: value.Extension,
						Size:      value.Size,
						UpdatedAt: &timeNow,
					}, tx)

					if err != nil {
						tx.Rollback()
						verifikasi.logger.Zap.Error(err)
						message = "Error, transaction database !"
						return false, message, err
					}

					_, err = verifikasi.verifikasiFile.Store(&models.VerifikasiFiles{
						VerifikasiID: request.ID,
						FilesID:      files.ID,
					}, tx)

					if err != nil {
						tx.Rollback()
						verifikasi.logger.Zap.Error(err)
						message = "Error, transaction database !"
						return false, message, err
					}
				} else {
					verifikasi.logger.Zap.Info("==========> Old Files")

					files, err := verifikasi.fileRepo.Store(&requestFile.Files{
						ID:        value.ID,
						Filename:  value.Filename,
						Path:      value.Path,
						Extension: value.Extension,
						Size:      value.Size,
						UpdatedAt: &timeNow,
					}, tx)

					if err != nil {
						tx.Rollback()
						verifikasi.logger.Zap.Error(err)
						message = "Error, transaction database !"
						return false, message, err
					}

					_, err = verifikasi.verifikasiFile.Store(&models.VerifikasiFiles{
						VerifikasiID: request.ID,
						FilesID:      files.ID,
					}, tx)

					if err != nil {
						tx.Rollback()
						verifikasi.logger.Zap.Error(err)
						message = "Error, transaction database !"
						return false, message, err
					}
				}
			}
		}
	} else {
		tx.Rollback()
		verifikasi.logger.Zap.Error(err)
		message = "Error, Data Lampiran Null !"
		return false, message, err
	}

	//#UpdateLampiran
	// }

	// Batch 3 Add by Panji 04/04/2024
	// Start Penyabab Kejadian
	if len(request.PenyababKejadian) != 0 {
		for _, value := range request.PenyababKejadian {
			_, err = verifikasi.verifikasiPenyababKejadian.Store(&models.VerifikasiPenyababKejadian{
				ID:                    value.ID,
				VerifikasiID:          request.ID,
				IDPenyebabKejadian:    value.IDPenyebabKejadian,
				IDSubPenyebabKejadian: value.IDSubPenyebabKejadian,
			}, tx)

			if err != nil {
				tx.Rollback()
				verifikasi.logger.Zap.Error(err)
				message = "Error, transaction database !"
				return false, message, err
			}
		}
	} else {
		tx.Rollback()
		verifikasi.logger.Zap.Error(err)
		message = "Error, Data Penyebab Kejadian Null !"
		return false, message, err
	}
	// End Of Penyebab Kejadian

	// Start of Usulan Perbaikan
	// fmt.Println("data usulan =>", request.UsulanPerbaikan)
	if len(request.UsulanPerbaikan) != 0 {
		for _, value := range request.UsulanPerbaikan {
			fmt.Println("perbaikin =>", value.Usulan)

			_, err = verifikasi.verifikasiUsulanPerbaikan.Store(&models.VerifikasiUsulanPerbaikan{
				ID:           value.ID,
				VerifikasiID: request.ID,
				Usulan:       value.Usulan,
				Deskripsi:    value.Deskripsi,
				Aplikasi:     value.Aplikasi,
			}, tx)
		}
	} else {
		tx.Rollback()
		verifikasi.logger.Zap.Error(err)
		message = "Error, transaction database !"
		return false, message, err
	}
	// End if Usulan Perbaikan

	tx.Commit()
	message = "Success, Input Data Berhasil !"
	return false, message, err
}

// FilterVerifikasi implements VerifikasiDefinition
func (verifikasi VerifikasiService) FilterVerifikasi(request models.VerifikasiFilterRequest) (responses []models.VerifikasiList, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort
	dataVerif, totalRows, totalData, err := verifikasi.verifikasiRepo.FilterVerifikasi(&request)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataVerif {
		responses = append(responses, models.VerifikasiList{
			ID:            response.ID,
			NoPelaporan:   response.NoPelaporan,
			UnitKerja:     response.UnitKerja,
			Aktifitas:     response.Aktifitas,
			IndikasiFraud: response.IndikasiFraud,
			StatusRtl:     response.StatusRtl,
			StatusVerif:   response.StatusVerif,
			StatusFraud:   response.StatusFraud,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)

	return responses, pagination, err
}

// GetDataWithPagination implements VerifikasiDefinition
func (verifikasi VerifikasiService) GetDataWithPagination(request models.VerifikasiPagination) (responses []models.VerifikasiList, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort
	dataVerif, totalRows, totalData, err := verifikasi.verifikasiRepo.GetDataWithPagination(&request)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataVerif {
		responses = append(responses, models.VerifikasiList{
			ID:            response.ID,
			NoPelaporan:   response.NoPelaporan,
			UnitKerja:     response.UnitKerja,
			Aktifitas:     response.Aktifitas,
			IndikasiFraud: response.IndikasiFraud,
			StatusVerif:   response.StatusVerif,
			StatusRtl:     response.StatusRtl,
			StatusFraud:   response.StatusFraud,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)

	return responses, pagination, err
}

// GetNoPelaporan implements VerifikasiDefinition
func (verifikasi VerifikasiService) GetNoPelaporan(request models.NoPalaporanRequest) (responses []models.NoPelaporanResponse, err error) {
	dataVerif, err := verifikasi.verifikasiRepo.GetNoPelaporan(&request)

	if err != nil {
		verifikasi.logger.Zap.Error(err)
		return responses, err
	}

	for _, response := range dataVerif {
		responses = append(responses, models.NoPelaporanResponse{
			ORGEH:       request.ORGEH,
			NoPelaporan: response.NoPelaporan.String,
		})
	}

	return responses, err
}

// GetLastID implements VerifikasiDefinition
func (verifikasi VerifikasiService) GetLastID() (responses []models.VerifikasiLastIDResponse, err error) {
	dataVerif, err := verifikasi.verifikasiRepo.GetLastID()
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		return responses, err
	}

	for _, response := range dataVerif {
		responses = append(responses, models.VerifikasiLastIDResponse{
			ID: response.ID.Int64,
		})
	}

	return responses, err
}

// FilterReport implements VerifikasiDefinition
func (verifikasi VerifikasiService) FilterReport(request models.VerifikasiFilterReport) (responses []models.VerifikasiReportResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort
	dataReport, totalRows, totalData, err := verifikasi.verifikasiRepo.FilterReport(&request)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataReport {
		responses = append(responses, models.VerifikasiReportResponse{
			ID:          response.ID.Int64,
			Tanggal:     response.Tanggal.String,
			KodeBranch:  response.KodeBranch.String,
			Aktifitas:   response.Aktifitas.String,
			Produk:      response.Produk.String,
			RiskIssue:   response.RiskIssue.String,
			JudulMateri: response.JudulMateri.String,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err
}

// StoreSimpan implements VerifikasiDefinition
func (verifikasi VerifikasiService) StoreSimpan(request models.VerifikasiRequest) (status bool, message string, err error) {
	timeNow := lib.GetTimeNow("timestime")

	tx := verifikasi.db.DB.Begin()

	//input data verifikasi
	reqVerif := &models.Verifikasi{
		ID:                        request.ID,
		NoPelaporan:               request.NoPelaporan,
		REGION:                    request.REGION,
		RGDESC:                    request.RGDESC,
		MAINBR:                    request.MAINBR,
		MBDESC:                    request.MBDESC,
		BRANCH:                    request.BRANCH,
		BRDESC:                    request.BRDESC,
		ActivityID:                request.ActivityID,
		SubActivityID:             request.SubActivityID,
		ProductID:                 request.ProductID,
		RiskIssueID:               request.RiskIssueID,
		RiskIssue:                 request.RiskIssue,
		RiskIndicatorID:           request.RiskIndicatorID,
		RiskIndicator:             request.RiskIndicator,
		SumberData:                request.SumberData,
		ApplicationID:             request.ApplicationID,
		HasilVerifikasi:           request.HasilVerifikasi,
		KunjunganNasabah:          *request.KunjunganNasabah,
		Perbaikan:                 request.Perbaikan,
		IndikasiFraud:             request.IndikasiFraud,
		TerdapatKerugianFinansial: request.TerdapatKerugianFinansial,
		JenisKerugianFinansial:    request.JenisKerugianFinansial,
		JumlahPerkiraanKerugian:   request.JumlahPerkiraanKerugian,
		JenisKerugianNonFinansial: request.JenisKerugianNonFinansial,
		JenisRekomendasi:          request.JenisRekomendasi,
		RekomendasiTindakLanjut:   request.RekomendasiTindakLanjut,
		RencanaTindakLanjut:       request.RencanaTindakLanjut,
		RiskTypeID:                request.RiskTypeID,
		AdaUsulanPerbaikan:        request.AdaUsulanPerbaikan,
		TanggalDitemukan:          request.TanggalDitemukan,
		TanggalMulaiRTL:           request.TanggalMulaiRTL,
		TanggalTargetSelesai:      request.TanggalTargetSelesai,
		MakerID:                   request.MakerID,
		MakerDesc:                 request.MakerDesc,
		MakerDate:                 &timeNow,
		LastMakerID:               request.LastMakerID,
		LastMakerDesc:             request.LastMakerDesc,
		LastMakerDate:             &timeNow,
		Status:                    request.Status,
		Action:                    request.Action,
		StatusIndikasiFraud:       request.StatusIndikasiFraud,
		ActionIndikasiFraud:       request.ActionIndikasiFraud,
		Deleted:                   false,
		CreatedAt:                 &timeNow,
	}

	dataVerif, err := verifikasi.verifikasiRepo.Store(reqVerif, tx)

	if err != nil {
		tx.Rollback()
		verifikasi.logger.Zap.Error(err)
		message = "Error, transaction database !"
		return false, message, err
	}

	// fmt.Println("data verifikasi : ", dataVerif)
	//end data verifikasi

	//Begin Input data anomali
	if request.SumberData == "KRID" {
		if len(request.DataAnomaliKRID) != 0 {
			for _, value := range request.DataAnomaliKRID {
				// objectString := []byte(value.Object)

				fmt.Println("object Stringg =======>", value)
				if value.Periode != "" && value.Object != "" {
					_, err = verifikasi.verifikasiAnomaliKRID.Store(&models.VerifikasiAnomaliDataKRID{
						VerifikasiID: dataVerif.ID,
						Periode:      value.Periode,
						Object:       value.Object,
						Status:       false,
					}, tx)

					if err != nil {
						tx.Rollback()
						verifikasi.logger.Zap.Error(err)
						message = "Error, transaction database !"
						return false, message, err
					}
				}

			}
		} else {
			tx.Rollback()
			verifikasi.logger.Zap.Error(err)
			message = "Error, Data Sample Null !"
			return false, message, err
		}
	} else if request.SumberData == "Non KRID" {
		if len(request.DataAnomali) != 0 {
			for _, value := range request.DataAnomali {
				_, err = verifikasi.verifikasiAnomali.Store(&models.VerifikasiAnomaliData{
					VerifikasiID:    dataVerif.ID,
					TanggalKejadian: value.TanggalKejadian,
					NomorRekening:   value.NomorRekening,
					Nominal:         value.Nominal,
					Keterangan:      value.Keterangan,
				}, tx)

				if err != nil {
					tx.Rollback()
					verifikasi.logger.Zap.Error(err)
					message = "Error, transaction database !"
					return false, message, err
				}
			}
		} else {
			tx.Rollback()
			verifikasi.logger.Zap.Error(err)
			message = "Error, Data Sample Null!"
			return false, message, err
		}
	} else if request.SumberData == "Tematik" {
		if len(request.SampleDataTeamatik) != 0 {
			for _, value := range request.SampleDataTeamatik {
				if value.Periode != "" && value.ColumnsData != "" {
					_, err = verifikasi.verifikasiDataTematik.Store(&models.VerifikasiDataTematik{
						VerifikasiId: dataVerif.ID,
						Periode:      value.Periode,
						Columns:      value.Columns,
						ColumnsData:  value.ColumnsData,
						Status:       false,
					}, tx)
				}

				if err != nil {
					tx.Rollback()
					verifikasi.logger.Zap.Error(err)
					message = "Error, transaction database !"
					return false, message, err
				}
			}
		} else {
			tx.Rollback()
			verifikasi.logger.Zap.Error(err)
			message = "Error, Data Sample Null !"
			return false, message, err
		}
	}
	//End Input data anomali

	// Add Input Questionner by Panji 30 07 2023
	if len(request.Questionner) != 0 {
		for _, value := range request.Questionner {
			_, err = verifikasi.verifikasiQuestioner.Store(&models.VerifikasiQuestionner{
				// ID:           va,
				VerifikasiID: dataVerif.ID,
				Questionner:  value.Questionner,
				DataSumber:   value.DataSumber,
				Checker:      value.Checker,
				Signer:       value.Signer,
				ApprovalOrd:  value.ApprovalOrd,
				JenisFraud:   value.JenisFraud,
			}, tx)

			if err != nil {
				tx.Rollback()
				verifikasi.logger.Zap.Error(err)
				message = "Error, transaction database !"
				return false, message, err
			}
		}
	} else {
		tx.Rollback()
		verifikasi.logger.Zap.Error(err)
		message = "Error, Questionare null !"
		return false, message, err
	}

	// if request.Perbaikan == true {
	//Begin Input Kelemahan Kontrol
	if len(request.RiskControl) != 0 {
		for _, value := range request.RiskControl {
			_, err = verifikasi.verifikasiRiskControl.Store(&models.VerifikasiRiskControl{
				VerifikasiId:  dataVerif.ID,
				RiskControlID: value.RiskControlID,
				RiskControl:   value.RiskControl,
			}, tx)

			if err != nil {
				tx.Rollback()
				verifikasi.logger.Zap.Error(err)
				message = "Error, transaction database !"
				return false, message, err
			}
		}
	} else {
		tx.Rollback()
		verifikasi.logger.Zap.Error(err)
		message = "Error, Data Risk Control Null !"
		return false, message, err
	}
	//End Input Kelemahan Kontrol

	//Begin Input data PIC
	if len(request.PICTindakLanjut) != 0 {
		for _, value := range request.PICTindakLanjut {
			_, err = verifikasi.verifikasiPIC.Store(&models.VerifikasiPICTindakLanjut{
				VerifikasiID:          dataVerif.ID,
				PICID:                 value.PICID,
				PICDetail:             value.PICDetail,
				TanggalTindakLanjut:   value.TanggalTindakLanjut,
				DeskripsiTindakLanjut: value.DeskripsiTindakLanjut,
				Status:                value.Status,
			}, tx)

			if err != nil {
				tx.Rollback()
				verifikasi.logger.Zap.Error(err)
				message = "Error, transaction database !"
				return false, message, err
			}
		}
	} else {
		tx.Rollback()
		verifikasi.logger.Zap.Error(err)
		message = "Error, Data PIC Tindak Lanju Null!"
		return false, message, err
	}
	//End Input data PIC

	//Begin Input Lampiran
	bucket := os.Getenv("BUCKET_NAME")

	if len(request.Files) != 0 {
		for _, value := range request.Files {
			UUID := uuid.New()
			var destinationPath string
			allowedExtensions := []string{".jpg", ".jpeg", ".pdf", ".doc", ".docx", ".xls", ".xlsx"}

			extension := strings.ToLower(filepath.Ext(value.Filename))
			fmt.Println("Extension =>", extension)
			if value.Filename != "" {

				if !contains(allowedExtensions, extension) {
					tx.Rollback()
					verifikasi.logger.Zap.Error(err)
					message = "Error, Invalid file extension"
					return false, message, err
				}

				bucketExist := verifikasi.minio.BucketExist(verifikasi.minio.Client(), bucket)

				pathSplit := strings.Split(value.Path, "/")
				sourcePath := fmt.Sprint(value.Path)
				destinationPath = pathSplit[1] + "/" +
					lib.GetTimeNow("year") + "/" +
					lib.GetTimeNow("month") + "/" +
					lib.GetTimeNow("day") + "/" +
					UUID.String() + "/" +
					// pathSplit[2] + "/" +
					value.Filename

				// newPath := "verifikasi/" +
				// 	lib.GetTimeNow("year") + "/" +
				// 	lib.GetTimeNow("month") + "/" +
				// 	lib.GetTimeNow("day")

				// destinationPath = newPath + "/" + value.Filename
				var uploaded bool
				if bucketExist {
					fmt.Println("Exist")
					fmt.Println(bucket)
					fmt.Println(sourcePath)
					fmt.Println(destinationPath)
					uploaded = verifikasi.minio.CopyObject(verifikasi.minio.Client(), bucket, sourcePath, bucket, destinationPath)
					// uploaded := verifikasi.minio.PutObject(verifikasi.minio.MinioClient, bucket, destinationPath, sourcePath)

					fmt.Println(uploaded)
				} else {
					fmt.Println("Not Exist")
					fmt.Println(bucket)
					fmt.Println(sourcePath)
					fmt.Println(destinationPath)
					verifikasi.minio.MakeBucket(verifikasi.minio.Client(), bucket, "")
					uploaded = verifikasi.minio.CopyObject(verifikasi.minio.Client(), bucket, sourcePath, bucket, destinationPath)
					// uploaded := verifikasi.minio.PutObject(verifikasi.minio.MinioClient, bucket, destinationPath, sourcePath)

					fmt.Println(uploaded)
				}

				if !uploaded {
					tx.Rollback()
					verifikasi.logger.Zap.Error(err)
					message = "Error, Temporary path document not found !"
					return false, message, err
				}
			}

			files, err := verifikasi.fileRepo.Store(&requestFile.Files{
				Filename:  value.Filename,
				Path:      destinationPath,
				Extension: value.Extension,
				Size:      value.Size,
				CreatedAt: &timeNow,
			}, tx)

			if err != nil {
				tx.Rollback()
				verifikasi.logger.Zap.Error(err)
				message = "Error, transaction database !"
				return false, message, err
			}

			_, err = verifikasi.verifikasiFile.Store(&models.VerifikasiFiles{
				VerifikasiID: dataVerif.ID,
				FilesID:      files.ID,
				// CreatedAt:    &timeNow,
			}, tx)

			if err != nil {
				tx.Rollback()
				verifikasi.logger.Zap.Error(err)
				message = "Error, transaction database !"
				return false, message, err
			}
		}
	} else {
		tx.Rollback()
		verifikasi.logger.Zap.Error(err)
		message = "Error, Data Lampiran Kosong !"
		return false, message, err
	}

	//End Input Lampiran
	// }

	// Batch 3 Add by Panji 04/04/2024
	// Start Penyabab Kejadian
	if len(request.PenyababKejadian) != 0 {
		for _, value := range request.PenyababKejadian {
			_, err = verifikasi.verifikasiPenyababKejadian.Store(&models.VerifikasiPenyababKejadian{
				VerifikasiID:          dataVerif.ID,
				IDPenyebabKejadian:    value.IDPenyebabKejadian,
				IDSubPenyebabKejadian: value.IDSubPenyebabKejadian,
			}, tx)

			if err != nil {
				tx.Rollback()
				verifikasi.logger.Zap.Error(err)
				message = "Error, transaction database !"
				return false, message, err
			}
		}
	} else {
		tx.Rollback()
		verifikasi.logger.Zap.Error(err)
		message = "Error, Data Penyebab Kejadian Null !"
		return false, message, err
	}
	// End Of Penyebab Kejadian

	// Start of Usulan Perbaikan
	// fmt.Println("data usulan =>", request.UsulanPerbaikan)
	if len(request.UsulanPerbaikan) != 0 {
		for _, value := range request.UsulanPerbaikan {
			fmt.Println("perbaikin =>", value.Usulan)

			_, err = verifikasi.verifikasiUsulanPerbaikan.Store(&models.VerifikasiUsulanPerbaikan{
				VerifikasiID: dataVerif.ID,
				Usulan:       value.Usulan,
				Deskripsi:    value.Deskripsi,
				Aplikasi:     value.Aplikasi,
			}, tx)
		}
	} else {
		tx.Rollback()
		verifikasi.logger.Zap.Error(err)
		message = "Error, transaction database !"
		return false, message, err
	}
	// End if Usulan Perbaikan

	tx.Commit()
	message = "Success, Input Data Berhasil !"
	return false, message, err
}

// report filter
func (verifikasi VerifikasiService) VerifikasiReportFilter(request models.VerifikasiFilterReportRequest) (responses []models.VerifikasiFilterReportResponse, totalRows int64, err error) {
	dataVerifikasi, totalRows, err := verifikasi.verifikasiRepo.VerifikasiReportFilter(&request)

	if err != nil {
		verifikasi.logger.Zap.Error(err)
		return responses, totalRows, err
	}

	// generate total
	sumOfTotalWeakness := int64(0)
	sumOfTotalNonWeakness := int64(0)
	sumOfGrandTotal := int64(0)

	fmt.Println("dataVerifikasi ==============")
	fmt.Println(dataVerifikasi)

	for _, dataSum := range dataVerifikasi {
		sumOfTotalWeakness += dataSum.TotalWeakness
		sumOfTotalNonWeakness += dataSum.TotalNonWeakness
		sumOfGrandTotal += dataSum.GrandTotal
	}

	// assign total to all data
	for _, response := range dataVerifikasi {
		PercentWeakness := float64(0)
		PercentNonWeakness := float64(0)
		PercentGrandTotal := float64(0)

		if sumOfTotalWeakness != 0 {
			PercentWeakness = (float64(response.TotalWeakness) / float64(sumOfTotalWeakness) * 100)
			fmt.Println("Service of SumOfTotalWeakness =>", sumOfTotalWeakness)
		}

		if sumOfTotalNonWeakness != 0 {
			PercentNonWeakness = (float64(response.TotalNonWeakness) / float64(sumOfTotalNonWeakness) * 100)
			fmt.Println("Service of SumOfTotalWeakness =>", sumOfTotalWeakness)
		}

		if sumOfGrandTotal != 0 {
			PercentGrandTotal = (float64(response.GrandTotal) / float64(sumOfGrandTotal) * 100)
		}

		responses = append(responses, models.VerifikasiFilterReportResponse{
			Id:                 response.Id,
			Code:               response.Code,
			Name:               response.Name,
			TotalWeakness:      response.TotalWeakness,
			PercentWeakness:    PercentWeakness,
			TotalNonWeakness:   response.TotalNonWeakness,
			PercentNonWeakness: PercentNonWeakness,
			GrandTotal:         response.GrandTotal,
			PercentGrandTotal:  PercentGrandTotal,
		})
	}

	fmt.Println("================== Responses")
	fmt.Println(responses)

	return responses, totalRows, err
}

func (verifikasi VerifikasiService) VerifikasiReportWithWeaknessOnlyFilter(request models.VerifikasiFilterReportRequest) (responses []models.VerifikasiFilterReportWeaknessOnlyResponse, totalRows int64, err error) {
	dataVerifikasi, totalRows, err := verifikasi.verifikasiRepo.VerifikasiReportWithWeaknessOnlyFilter(&request)

	if err != nil {
		verifikasi.logger.Zap.Error(err)
		return responses, totalRows, err
	}

	// generate total
	sumOfTotalWeakness := int64(0)

	fmt.Println("dataVerifikasi ==============")
	fmt.Println(dataVerifikasi)

	for _, dataSum := range dataVerifikasi {
		sumOfTotalWeakness += dataSum.TotalWeakness
	}

	// assign total to all data
	for _, response := range dataVerifikasi {
		PercentWeakness := float64(0)

		if sumOfTotalWeakness != 0 {
			PercentWeakness = (float64(response.TotalWeakness) / float64(sumOfTotalWeakness) * 100)
		}

		responses = append(responses, models.VerifikasiFilterReportWeaknessOnlyResponse{
			Id:              response.Id,
			Code:            response.Code,
			Name:            response.Name,
			TotalWeakness:   response.TotalWeakness,
			PercentWeakness: PercentWeakness,
		})
	}

	fmt.Println("================== Responses")
	fmt.Println(responses)

	return responses, totalRows, err
}

func (verifikasi VerifikasiService) VerifikasiReportWithNonWeaknessOnlyFilter(request models.VerifikasiFilterReportRequest) (responses []models.VerifikasiFilterReportNonWeaknessOnlyResponse, totalRows int64, err error) {
	dataVerifikasi, totalRows, err := verifikasi.verifikasiRepo.VerifikasiReportFilter(&request)

	if err != nil {
		verifikasi.logger.Zap.Error(err)
		return responses, totalRows, err
	}

	// generate total
	sumOfTotalNonWeakness := int64(0)

	fmt.Println("dataVerifikasi ==============")
	fmt.Println(dataVerifikasi)

	for _, dataSum := range dataVerifikasi {
		sumOfTotalNonWeakness += dataSum.TotalNonWeakness
	}

	// assign total to all data
	for _, response := range dataVerifikasi {
		PercentNonWeakness := float64(0)

		if sumOfTotalNonWeakness != 0 {
			PercentNonWeakness = (float64(response.TotalNonWeakness) / float64(sumOfTotalNonWeakness) * 100)
		}

		responses = append(responses, models.VerifikasiFilterReportNonWeaknessOnlyResponse{
			Id:                 response.Id,
			Code:               response.Code,
			Name:               response.Name,
			TotalNonWeakness:   response.TotalNonWeakness,
			PercentNonWeakness: PercentNonWeakness,
		})
	}

	fmt.Println("================== Responses")
	fmt.Println(responses)

	return responses, totalRows, err
}

func (verifikasi VerifikasiService) VerifikasiReportFilterComplete(request models.VerifikasiFilterReportRequest) (responses []models.VerifikasiFilterReportCompleteResponse, totalRows int64, err error) {
	dataVerifikasi, totalRows, err := verifikasi.verifikasiRepo.VerifikasiReportFilterComplete(&request)

	if err != nil {
		verifikasi.logger.Zap.Error(err)
		return responses, totalRows, err
	}

	// assign total to all data
	for _, response := range dataVerifikasi {
		responses = append(responses, models.VerifikasiFilterReportCompleteResponse{
			Id:           response.Id,
			Date:         response.Date,
			BRANCH:       response.BRANCH,
			BRDESC:       response.BRDESC,
			ActivityName: response.ActivityName,
			ProductName:  response.ProductName,
			RiskIssue:    response.RiskIssue,
			JudulMateri:  response.JudulMateri,
		})
	}

	fmt.Println("================== Responses")
	fmt.Println(responses)

	return responses, totalRows, err
}

func (verifikasi VerifikasiService) VerifikasiReportDetail(request models.VerifikasiReportDetailRequest) (responses models.VerifikasiReportDetailResponse, err error) {
	dataDetail, err := verifikasi.verifikasiRepo.VerifikasiReportDetail(&request)

	if err != nil {
		verifikasi.logger.Zap.Error(err)
		return responses, err
	}

	data_anomali, err := verifikasi.verifikasiAnomali.GetOneByVerifikasi(dataDetail.Id.Int64)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		return responses, err
	}

	data_anomali_krid, err := verifikasi.verifikasiAnomaliKRID.GetOneByVerifikasi(dataDetail.Id.Int64)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		return responses, err
	}

	list_files, err := verifikasi.verifikasiFile.GetOneFileByID(dataDetail.Id.Int64)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		return responses, err
	}

	fmt.Println("data anomali==>", data_anomali_krid)

	// assign data anomali
	// responseDataAnomali := []models.DataAnomali{}
	// for _, dataAnomali := range dataAnomali {
	// 	responseDataAnomali = append(responseDataAnomali, models.DataAnomali{
	// 		Date:       dataAnomali.Date.String,
	// 		NoRek:      dataAnomali.NoRek.String,
	// 		Nominal:    dataAnomali.Nominal.String,
	// 		Keterangan: dataAnomali.Keterangan.String,
	// 	})
	// }

	// assign to response
	responses = models.VerifikasiReportDetailResponse{
		Id:                   dataDetail.Id.Int64,
		NoPelaporan:          dataDetail.NoPelaporan.String,
		BRANCH:               dataDetail.BRANCH.String,
		BRDESC:               dataDetail.BRDESC.String,
		MAINBR:               dataDetail.MAINBR.String,
		MBDESC:               dataDetail.MBDESC.String,
		REGION:               dataDetail.REGION.String,
		RGDESC:               dataDetail.RGDESC.String,
		ActivityName:         dataDetail.ActivityName.String,
		SubActivityName:      dataDetail.SubActivityName.String,
		ProductName:          dataDetail.ProductName.String,
		RiskIssue:            dataDetail.RiskIssue.String,
		RiskIndicator:        dataDetail.RiskIndicator.String,
		DataAnomali:          data_anomali,
		DataAnomaliKRID:      data_anomali_krid,
		Files:                list_files,
		IncidentCauseCode:    dataDetail.IncidentCauseCode.String,
		IncidentCauseName:    dataDetail.IncidentCauseName.String,
		SubIncidentCauseCode: dataDetail.SubIncidentCauseCode.String,
		SubIncidentCauseName: dataDetail.SubIncidentCauseName.String,
		VerificationResult:   dataDetail.VerificationResult.String,
		DataSource:           dataDetail.DataSource.String,
		Perbaikan:            dataDetail.Perbaikan.Bool,
		IndikasiFraud:        dataDetail.IndikasiFraud.Bool,
	}

	fmt.Println("")
	fmt.Println("responses service")
	fmt.Println(responses.IncidentCauseName)
	fmt.Println("")

	return responses, err
}

func (verifikasi VerifikasiService) RiskControlByVerificationId(request models.DataRiskControlRequest) (responses []models.DataRiskIndicatorResponse, totalRows int64, err error) {
	data, totalRows, err := verifikasi.verifikasiRepo.RiskControlByVerificationId(&request)
	totalAllDataRisk := int64(0)

	if err != nil {
		verifikasi.logger.Zap.Error(err)
		return responses, totalRows, err
	}

	// count total
	for _, response := range data {
		totalAllDataRisk += response.Total
	}

	// assign data response
	for _, response := range data {
		percentCount := (float64(response.Total) / float64(totalAllDataRisk) * 100)
		if percentCount == 0 {
			percentCount = 0
		}
		responses = append(responses, models.DataRiskIndicatorResponse{
			RiskControl: response.RiskControl,
			Total:       response.Total,
			Percent:     percentCount,
		})
	}

	return responses, totalRows, err
}

func (verifikasi VerifikasiService) GetRiskIndicatorAsMateri(request models.VerifikasiFilterReportRequest) (responses []models.GetRiskIndicatorAsMateriResponse, err error) {
	data, err := verifikasi.verifikasiRepo.GetRiskIndicatorAsMateri(&request)

	if err != nil {
		verifikasi.logger.Zap.Error(err)
		return responses, err
	}

	// assign data response
	for _, response := range data {
		responses = append(responses, models.GetRiskIndicatorAsMateriResponse{
			ID:   response.ID.String,
			Code: response.Code.String,
			Name: response.Name.String,
		})
	}

	return responses, err
}

func (verifikasi VerifikasiService) VerificationReportByUkerFilter(request models.VerificationFilterReportByUkerRequest) (responses []models.VerificationFilterReportByUkerResponse, totalRows int, err error) {
	data, SumData, totalRows, err := verifikasi.verifikasiRepo.VerificationReportByUkerFilter(&request)

	if err != nil {
		verifikasi.logger.Zap.Error(err)
		return responses, totalRows, err
	}

	// assign data response
	for _, response := range data {
		PERCENTVERIFICATION := float64(0)
		PERCENTWEAKNESS := float64(0)
		PERCENTPERBAIKANONPROGRESS := float64(0)

		// TOTALVERIFICATION := response.TOTALVERIFICATION.Int64 + response.TOTALBRC.Int64
		TOTALVERIFICATION := SumData
		if TOTALVERIFICATION != 0 {
			PERCENTVERIFICATION = (float64(response.TOTALVERIFICATION.Int64) / float64(TOTALVERIFICATION)) * 100
		}

		TOTALWEAKNESS := response.TOTALNONWEAKNESS.Int64 + response.TOTALNWEAKNESS.Int64
		if TOTALWEAKNESS != 0 {
			PERCENTWEAKNESS = (float64(response.TOTALNWEAKNESS.Int64) / float64(TOTALWEAKNESS)) * 100
		}

		// TOTALPERBAIKANONPROGRESS := response.TOTALPERBAIKANONPROGRESS.Int64 + response.TOTALPERBAIKANDONE.Int64
		TOTALPERBAIKANONPROGRESS := response.TOTALNWEAKNESS.Int64
		if TOTALPERBAIKANONPROGRESS != 0 {
			PERCENTPERBAIKANONPROGRESS = (float64(response.TOTALPERBAIKANONPROGRESS.Int64) / float64(TOTALPERBAIKANONPROGRESS)) * 100
		}

		// PERCENTWEAKNESS := float64( (response.TOTALNWEAKNESS.Int64 / (response.TOTALNONWEAKNESS.Int64 + response.TOTALNWEAKNESS.Int64)) * 100 )
		// PERCENTPERBAIKANONPROGRESS := float64( (response.TOTALPERBAIKANONPROGRESS.Int64 / (response.TOTALPERBAIKANONPROGRESS.Int64 + response.TOTALPERBAIKANDONE.Int64)) * 100 )

		responses = append(responses, models.VerificationFilterReportByUkerResponse{
			REGION: response.REGION.String,
			RGDESC: response.RGDESC.String,
			MAINBR: response.MAINBR.String,
			MBDESC: response.MBDESC.String,
			BRANCH: response.BRANCH.String,
			BRDESC: response.BRDESC.String,

			TOTALVERIFICATION:   response.TOTALVERIFICATION.Int64,
			TOTALBRC:            response.TOTALBRC.Int64,
			PERCENTVERIFICATION: PERCENTVERIFICATION,

			TOTALNONWEAKNESS: response.TOTALNONWEAKNESS.Int64,
			TOTALNWEAKNESS:   response.TOTALNWEAKNESS.Int64,
			PERCENTWEAKNESS:  PERCENTWEAKNESS,

			TOTALPERBAIKANONPROGRESS:   response.TOTALPERBAIKANONPROGRESS.Int64,
			TOTALPERBAIKANDONE:         response.TOTALPERBAIKANDONE.Int64,
			PERCENTPERBAIKANONPROGRESS: PERCENTPERBAIKANONPROGRESS,
		})
	}

	fmt.Println("response service =====")
	fmt.Println(responses)

	return responses, totalRows, err
}

func (verifikasi VerifikasiService) VerificationReportFilterByUkerComplete(request models.VerificationFilterReportByUkerRequest) (responses []models.VerificationFilterByUkerReportCompleteResponse, totalRows int64, err error) {
	data, totalRows, err := verifikasi.verifikasiRepo.VerificationReportFilterByUkerComplete(&request)

	if err != nil {
		verifikasi.logger.Zap.Error(err)
		return responses, totalRows, err
	}

	// assign data response
	for _, response := range data {
		responses = append(responses, models.VerificationFilterByUkerReportCompleteResponse{
			Id:               response.Id,
			Date:             response.Date,
			BRANCH:           response.BRANCH,
			BRDESC:           response.BRDESC,
			Activity:         response.Activity,
			Product:          response.Product,
			RiskIssue:        response.RiskIssue,
			Materi:           response.Materi,
			IsRequiredFixing: response.IsRequiredFixing,
			FixingStatus:     response.FixingStatus,
		})
	}

	return responses, totalRows, err
}

func (verifikasi VerifikasiService) VerifikasiReportByFraudIndicatorFilter(request models.VerificationFilterReportByUkerRequest) (responses []models.VerificationFilterReportByFraudIndicatorResponse, totalRows int64, err error) {
	data, totalRows, err := verifikasi.verifikasiRepo.VerifikasiReportByFraudIndicatorFilter(&request)

	if err != nil {
		verifikasi.logger.Zap.Error(err)
		return responses, totalRows, err
	}

	// assign data response
	for _, response := range data {
		responses = append(responses, models.VerificationFilterReportByFraudIndicatorResponse{
			REGION:     response.REGION,
			RGDESC:     response.RGDESC,
			MAINBR:     response.MAINBR,
			MBDESC:     response.MBDESC,
			BRANCH:     response.BRANCH,
			BRDESC:     response.BRDESC,
			TOTALFRAUD: response.TOTALFRAUD,
		})
	}

	return responses, totalRows, err
}

func (verifikasi VerifikasiService) VerificationReportFilterByFraudIndicatorComplete(request models.VerificationFilterReportByUkerRequest) (responses []models.VerifikasiFilterReportCompleteResponse, totalRows int, err error) {
	data, totalRows, err := verifikasi.verifikasiRepo.VerificationReportFilterByFraudIndicatorComplete(&request)

	if err != nil {
		verifikasi.logger.Zap.Error(err)
		return responses, totalRows, err
	}

	// assign data response
	for _, response := range data {
		responses = append(responses, models.VerifikasiFilterReportCompleteResponse{
			Id:           response.Id,
			Date:         response.Date,
			BRANCH:       response.BRANCH,
			BRDESC:       response.BRDESC,
			ActivityName: response.ActivityName,
			ProductName:  response.ProductName,
			RiskIssue:    response.RiskIssue,
			JudulMateri:  response.JudulMateri,
		})
	}

	return responses, totalRows, err
}

// add 23 Feb 2023 By Panji
func (verif VerifikasiService) VerifikasiReportMateriList(request models.VerifikasiMateriRequest) (responses []models.VerifikasiDetailMateriResponse, err error) {
	dataBriefing, _, _, err := verif.verifikasiRepo.VerifikasiReportMateriList(&request)

	if err != nil {
		verif.logger.Zap.Error(err)
		return responses, err
	}

	for _, response := range dataBriefing {
		responses = append(responses, models.VerifikasiDetailMateriResponse{
			ID:           response.ID.Int64,
			Filename:     response.Filename.String,
			NamaLampiran: response.NamaLampiran.String,
			Path:         response.Path.String,
		})
	}

	fmt.Println("response service")
	fmt.Println(responses)

	return responses, err
}

func (verif VerifikasiService) VerificationReportUkerByAllActivity(request models.VerificationFilterReportByUkerRequest) (responses models.VerifikasiReportAllUker, totalRows int, err error) {
	data, totalRows, err := verif.verifikasiRepo.VerificationReportUkerByAllActivity(&request)

	// assign activity list
	// add grand total in activitylist
	data.ActivityList = append(data.ActivityList, models.ActivityList{
		ID:           int64(0),
		KodeActivity: "",
		Name:         "GRAND TOTAL",
	})

	responses.ActivityList = data.ActivityList

	if err != nil {
		verif.logger.Zap.Error(err)
		return responses, totalRows, err
	}

	for _, item := range data.Data {
		grandTotal := 0.0

		for key, value := range item {
			if floatValue, ok := value.(float64); ok {
				grandTotal += floatValue
				item[key] = lib.ToFixedWithPercent(floatValue, 2)
			}
		}

		item["GRAND TOTAL"] = lib.ToFixedWithPercent(grandTotal, 2)
	}

	// fmt.Println("responses - service", data)

	return data, totalRows, err
}

func (verif VerifikasiService) VerificationReportUkerByAllActivityComplete(request models.VerificationFilterReportByUkerRequest) (responses []map[string]interface{}, totalRows int, err error) {
	data, totalRows, err := verif.verifikasiRepo.VerificationReportUkerByAllActivityComplete(&request)

	for _, rows := range data {
		num := 0.0
		percent := ""

		if rows.WEAKNESS != 0 {
			num = (float64(rows.WEAKNESS) / float64(rows.TOTAL)) * 100
			percent = lib.ToFixedWithPercent(num, 2)
		} else {
			percent = lib.ToFixedWithPercent(0, 2)
		}

		responseTemp := map[string]interface{}{
			"risk_issue": rows.RiskIssue,
			"percent":    percent,
		}

		responses = append(responses, responseTemp)
	}

	return responses, totalRows, err
}

func (verif VerifikasiService) VerificationReportUkerByAllActivityCompleteWithRiskIssue(request models.VerificationFilterReportByUkerRequest) (responses []map[string]interface{}, totalRows int, err error) {
	data, totalRows, err := verif.verifikasiRepo.VerificationReportUkerByAllActivityCompleteWithRiskIssue(&request)

	return data, totalRows, err
}

// VerifikasiReportList implements VerifikasiDefinition
func (verif VerifikasiService) VerifikasiReportList(request models.VerifikasiReportListRequest) (responses []models.VerifikasiReportListResponse, totalRows int, err error) {
	dataVerif, totalRows, err := verif.verifikasiRepo.VerifikasiReportList(&request)

	var presentase_perbaikan float64

	// var count int

	if err != nil {
		verif.logger.Zap.Error(err)
		return responses, totalRows, err
	}

	for _, response := range dataVerif {

		if response.ButuhPerbaikan == "Tidak" {
			presentase_perbaikan = (float64(response.StatusPerbaikanSelesai) / float64(response.JumlahDataYgDiverifikasi)) * 100
		} else {
			presentase_perbaikan = (float64(response.StatusPerbaikanSelesai) / float64(response.JumlahDataYgHarusDiperbaiki)) * 100
		}
		// presentase_perbaikan = (response.StatusPerbaikanSelesai / response.JumlahDataYgHarusDiperbaiki) * 100

		fmt.Println(response.ButuhPerbaikan, "=>", presentase_perbaikan)

		// fmt.Println("presentase_perbaikan ==>", presentase_perbaikan)

		responses = append(responses, models.VerifikasiReportListResponse{
			ID:                          response.ID,
			Periode:                     response.Periode,
			RGDESC:                      response.RGDESC,
			MBDESC:                      response.MBDESC,
			BRANCH:                      response.BRANCH,
			BRDESC:                      response.BRDESC,
			NoPelaporan:                 response.NoPelaporan,
			Aktifitas:                   response.Aktifitas,
			SubAktifitas:                response.SubAktifitas,
			InformasiLain:               response.InformasiLain,
			StatusPerbaikanKonsolidasi:  response.StatusPerbaikanKonsolidasi,
			Maker:                       response.Maker,
			RiskIssueCode:               response.RiskIssueCode,
			RiskIssue:                   response.RiskIssue,
			RiskIndicator:               response.RiskIndicator,
			RiskControl:                 response.RiskControl,
			HasilVerifikasi:             response.HasilVerifikasi,
			JumlahDataYgDiverifikasi:    response.JumlahDataYgDiverifikasi,
			ButuhPerbaikan:              response.ButuhPerbaikan,
			JumlahDataYgHarusDiperbaiki: response.JumlahDataYgHarusDiperbaiki,
			RTLUser:                     response.RTLUser,
			StatusPerbaikanSelesai:      response.StatusPerbaikanSelesai,
			StatusPerbaikanProses:       response.StatusPerbaikanProses,
			PresentasePerbaikan:         int(presentase_perbaikan),
			BatasWaktuPerbaikan:         response.BatasWaktuPerbaikan,
			IndikasiFraud:               response.IndikasiFraud,
			Filename:                    response.Filename,
			Filepath:                    response.Filepath,
		})
	}

	return responses, totalRows, err
}

// RptRekapitulasiBCV implements VerifikasiDefinition
func (verif VerifikasiService) RptRekapitulasiBCV(request models.RptRekapitulasiBCVRequest) (responses []models.RptRekapitulasiBCVResponse, totalRows int, err error) {
	dataRekap, totalRows, err := verif.verifikasiRepo.RptRakapitulasiBCV(&request)

	if err != nil {
		verif.logger.Zap.Error(err)
		return responses, totalRows, err
	}

	for _, response := range dataRekap {
		responses = append(responses, models.RptRekapitulasiBCVResponse{
			Pernr:   response.Pernr,
			BRC:     response.BRC,
			BRANCH:  response.BRANCH,
			BRDESC:  response.BRDESC,
			MBDESC:  response.MBDESC,
			RGDESC:  response.RGDESC,
			BDraft:  response.BDraft,
			BFinish: response.BFinish,
			BTotal:  response.BTotal,
			CDraft:  response.CDraft,
			CFinish: response.CFinish,
			CTotal:  response.CTotal,
			VDraft:  response.VDraft,
			VFinish: response.VFinish,
			VTotal:  response.VTotal,
		})
	}

	return responses, totalRows, err
}

// RptRekomendasiRisk implements VerifikasiDefinition
func (verif VerifikasiService) RptRekomendasiRisk(request models.RptRekomendasiRiskRequest) (responses []models.RptRekomendasiRiskResponse, totalRows int, err error) {
	rptRekomendasiRisk, totalRows, err := verif.verifikasiRepo.RptRekomendasiRisk(&request)

	if err != nil {
		verif.logger.Zap.Error(err)
		return responses, totalRows, err
	}

	if len(rptRekomendasiRisk) != 0 {
		for _, response := range rptRekomendasiRisk {
			responses = append(responses, models.RptRekomendasiRiskResponse{
				RiskEvent:     response.RiskEvent,
				RiskIndicator: response.RiskIndicator,
				Module:        response.Module,
				Count:         response.Count,
			})
		}
	}

	return responses, totalRows, err
}

// ValidasiVerifikasi implements VerifikasiDefinition
func (verif VerifikasiService) ValidasiVerifikasi(request models.ValidasiVerifikasiRequest) (responses []models.ValidasiVerifikasiResponse, totalRows int, err error) {
	dataValidasi, totalRows, err := verif.verifikasiRepo.ValidasiVerifikasi(&request)

	if err != nil {
		verif.logger.Zap.Error(err)
		return responses, totalRows, err
	}

	for _, response := range dataValidasi {
		responses = append(responses, models.ValidasiVerifikasiResponse{
			VerifikasiId: response.VerifikasiId,
			NoPelaporan:  response.NoPelaporan,
			UnitKerja:    response.UnitKerja,
			Aktivitas:    response.Aktivitas,
			RiskIssue:    response.RiskIssue,
			MakerDesc:    response.MakerDesc,
			ValidasiId:   response.ValidasiId,
			ValidatorRmc: response.ValidatorRmc,
			StatusRmc:    response.StatusRmc,
			ValidatorRrm: response.ValidatorRrm,
			StatusSigner: response.StatusSigner,
			ValidatorOrd: response.ValidatorOrd,
			StatusOrd:    response.StatusOrd,
		})
	}

	return responses, totalRows, err
}

// AcceptValidasi implements VerifikasiDefinition
func (vs VerifikasiService) AcceptValidasi(request *models.AcceptValidasiRequest) (responses bool, err error) {
	tx := vs.db.DB.Begin()

	_, err = vs.verifikasiQuestioner.AcceptValidasi(&models.AcceptValidasiRequest{
		ID:                              request.ID,
		StatusValidasiRmc:               request.StatusValidasiRmc,
		TindakLanjutIndikasiFraudRmc:    request.TindakLanjutIndikasiFraudRmc,
		TindakLanjutRmc:                 request.TindakLanjutRmc,
		CatatanRmc:                      request.CatatanRmc,
		StatusValidasiSigner:            request.StatusValidasiSigner,
		TindakLanjutIndikasiFraudSigner: request.TindakLanjutIndikasiFraudSigner,
		TindakLanjutSigner:              request.TindakLanjutSigner,
		CatatanSigner:                   request.CatatanSigner,
		StatusValidasiOrd:               request.StatusValidasiOrd,
		ValidasiIndikasiFraudOrd:        request.ValidasiIndikasiFraudOrd,
		TindakLanjutOrd:                 request.TindakLanjutOrd,
		CatatanOrd:                      request.CatatanOrd,
		ValidationBy:                    request.ValidationBy,
	}, tx)

	if err != nil {
		tx.Rollback()
		vs.logger.Zap.Error(err)
		return false, err
	}

	tx.Commit()
	return true, err
}

// UpdateStatusValidasi implements VerifikasiDefinition
func (vs VerifikasiService) UpdateStatusVerifikasi(request *models.UpdateStatusVerifikasi) (responses bool, err error) {
	tx := vs.db.DB.Begin()

	_, err = vs.verifikasiQuestioner.UpdateStatusVerifikasi(&models.UpdateStatusVerifikasi{
		ID:                  request.ID,
		IsTask:              request.IsTask,
		StatusIndikasiFraud: request.StatusIndikasiFraud,
		ActionIndikasiFraud: request.ActionIndikasiFraud,
	}, tx)

	if err != nil {
		tx.Rollback()
		vs.logger.Zap.Error(err)
		return false, err
	}

	tx.Commit()
	return true, err
}

// RejectValidasi implements VerifikasiDefinition
func (vs VerifikasiService) RejectValidasi(request *models.RejectValidasiRequest) (response bool, err error) {
	tx := vs.db.DB.Begin()

	_, err = vs.verifikasiQuestioner.RejectValidasi(&models.RejectValidasiRequest{
		ID:                   request.ID,
		StatusValidasiRmc:    request.StatusValidasiRmc,
		CatatanRmc:           request.CatatanRmc,
		StatusValidasiSigner: request.StatusValidasiSigner,
		CatatanSigner:        request.CatatanSigner,
		StatusValidasiOrd:    request.StatusValidasiOrd,
		CatatanOrd:           request.CatatanOrd,
		RejectBy:             request.RejectBy,
	}, tx)

	// fmt.Println("service Reject =>", request)

	if err != nil {
		tx.Rollback()
		vs.logger.Zap.Error(err)
		return false, err
	}

	tx.Commit()
	return true, err
}

func (verif VerifikasiService) GetRtlIndikasiFraud(request models.ReqRtlIndikasiFraud) (responses models.RtlIndikasiFraudResponse, totalRows int, err error) {
	responses, totalRows, err = verif.verifikasiRepo.GetRtlIndikasiFraud(&request)

	return responses, totalRows, err
}

// DeleteAnomaliByID implements VerifikasiDefinition
// Versioning 1.0.0.1 by panji 31/08/2023
func (verif VerifikasiService) DeleteAnomaliByID(request *models.VerifikasiAnomaliData) (response bool, err error) {
	err = verif.verifikasiAnomali.Delete(request.ID)

	return true, err
}

// ValidasiVerifikasiDetailData implements VerifikasiDefinition.
func (verifikasi VerifikasiService) ValidasiVerifikasiDetailData(request models.VerifikasiReportDetailRequest) (responses models.ValidasiVerifikasiDetailedResponse, err error) {
	dataVerif, err := verifikasi.verifikasiRepo.ValidasiVerifikasiDetailData(&request)
	fmt.Println(dataVerif)

	if dataVerif.Id != 0 {
		fmt.Println("bukan 0")

		data_anomali, err := verifikasi.verifikasiAnomali.GetOneByVerifikasi(dataVerif.Id)
		// questionner, err := verifikasi.verifikasiQuestioner.GetOneByVerifikasi(dataVerif.Id)
		if err != nil {
			verifikasi.logger.Zap.Error(err)
			return responses, err
		}

		files, err := verifikasi.verifikasiFile.GetOneFileByID(dataVerif.Id)
		if err != nil {
			verifikasi.logger.Zap.Error(err)
			return responses, err
		}

		pic_tindak_lanjut, err := verifikasi.verifikasiPIC.GetOneByPIC(dataVerif.Id)
		if err != nil {
			verifikasi.logger.Zap.Error(err)
			return responses, err
		}

		risk_control, err := verifikasi.verifikasiRiskControl.GetOneDataByID(dataVerif.Id)
		if err != nil {
			verifikasi.logger.Zap.Error(err)
			return responses, err
		}

		data_anomali_krid, err := verifikasi.verifikasiAnomaliKRID.GetOneByVerifikasi(dataVerif.Id)
		if err != nil {
			verifikasi.logger.Zap.Error(err)
			return responses, err
		}

		penyebabKejadian, err := verifikasi.verifikasiPenyababKejadian.GetDataDetail(dataVerif.Id)
		if err != nil {
			verifikasi.logger.Zap.Error(err)
			return responses, err
		}

		usulan, err := verifikasi.verifikasiUsulanPerbaikan.GetData(dataVerif.Id)
		if err != nil {
			verifikasi.logger.Zap.Error(err)
			return responses, err
		}

		rekemendasi_materi, err := verifikasi.VerifikasiReportMateriList(
			models.VerifikasiMateriRequest{
				Id: dataVerif.RekomendasiTindakLanjut,
			},
		)
		if err != nil {
			verifikasi.logger.Zap.Error(err)
			return responses, err
		}

		responses = models.ValidasiVerifikasiDetailedResponse{
			Id:                        dataVerif.Id,
			NoPelaporan:               dataVerif.NoPelaporan,
			BRANCH:                    dataVerif.BRANCH,
			BRDESC:                    dataVerif.BRDESC,
			MAINBR:                    dataVerif.MAINBR,
			MBDESC:                    dataVerif.MBDESC,
			REGION:                    dataVerif.REGION,
			RGDESC:                    dataVerif.RGDESC,
			ActivityName:              dataVerif.ActivityName,
			SubActivityName:           dataVerif.SubActivityName,
			ProductName:               dataVerif.ProductName,
			RiskIssue:                 dataVerif.RiskIssue,
			RiskIndicator:             dataVerif.RiskIndicator,
			VerificationResult:        dataVerif.VerificationResult,
			DataSource:                dataVerif.DataSource,
			Perbaikan:                 dataVerif.Perbaikan,
			IndikasiFraud:             dataVerif.IndikasiFraud,
			TerdapatKerugianFinansial: dataVerif.TerdapatKerugianFinansial,
			JenisKerugianFinansial:    dataVerif.JenisKerugianFinansial,
			JumlahPerkiraanKerugian:   dataVerif.JumlahPerkiraanKerugian,
			JenisRekomendasi:          dataVerif.JenisRekomendasi,
			RekomendasiTindakLanjut:   dataVerif.RekomendasiTindakLanjut,
			RencanaTindakLanjut:       dataVerif.RencanaTindakLanjut,
			RiskType:                  dataVerif.RiskType,
			TanggalDitemukan:          dataVerif.TanggalDitemukan,
			TanggalMulaiRTL:           dataVerif.TanggalMulaiRTL,
			TanggalTargetSelesai:      dataVerif.TanggalTargetSelesai,
			AdaUsulanPerbaikan:        dataVerif.AdaUsulanPerbaikan,
			VerifikasiMateriDetail:    rekemendasi_materi,
			DataAnomali:               data_anomali,
			DataAnomaliKRID:           data_anomali_krid,
			PICTindakLanjut:           pic_tindak_lanjut,
			Files:                     files,
			RiskControl:               risk_control,
			UsulanPerbaikan:           usulan,
			PenyababKejadian:          penyebabKejadian,
		}

		return responses, err
	}

	return responses, err
}

// GetRekomendasiTindakLanjut implements VerifikasiDefinition.
func (verif VerifikasiService) GetRekomendasiTindakLanjut(request models.RTLRequest) (responses []models.RTLResponses, err error) {
	dataRTL, err := verif.verifikasiRepo.GetRekomendasiTindakLanjut(&request)

	return dataRTL, err
}

// DeletePenyebabKejadian implements VerifikasiDefinition.
func (verif VerifikasiService) DeletePenyebabKejadian(id int64) (response bool, err error) {
	err = verif.verifikasiPenyababKejadian.Delete(id)

	return true, err
}

// VerifikasiSummaryRpt implements VerifikasiDefinition.
func (verif VerifikasiService) VerifikasiSummaryRpt(request models.SummaryVerifikasiRequest) (responses []models.SummaryVerifikasiResponse, totalRows int, err error) {
	summaryRpt, totalRows, err := verif.verifikasiRepo.VerifikasiSummaryRpt(&request)

	if err != nil {
		verif.logger.Zap.Error(err)
		return responses, totalRows, err
	}

	for _, value := range summaryRpt {
		responses = append(responses, models.SummaryVerifikasiResponse{
			Aktivitas:     value.Aktivitas,
			Produk:        value.Produk,
			RiskEvent:     value.RiskEvent,
			RiskIndicator: value.RiskIndicator,
			RiskControl:   value.RiskControl,
			Jumlah:        value.Jumlah,
		})
	}

	return responses, totalRows, err
}

func (verif VerifikasiService) VerifikasiFrekuensiRpt(request models.FrekuensiVerifikasiRequest) (responses []models.FrekuensiVerifikasiResponse, totalRows int, err error) {
	frekuensiRpt, totalRows, err := verif.verifikasiRepo.VerifikasiFrekuensiRpt(&request)

	if err != nil {
		verif.logger.Zap.Error(err)
		return responses, totalRows, err
	}

	for _, value := range frekuensiRpt {
		responses = append(responses, models.FrekuensiVerifikasiResponse{
			Aktivitas:     value.Aktivitas,
			Produk:        value.Produk,
			RiskEvent:     value.RiskEvent,
			RiskIndicator: value.RiskIndicator,
			RiskControl:   value.RiskControl,
			Jumlah:        value.Jumlah,
		})
	}

	return responses, totalRows, err
}
