package riskindicator

import (
	"bytes"
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"riskmanagement/dto"
	"riskmanagement/lib"
	modelActiv "riskmanagement/models/activity"
	modelProduct "riskmanagement/models/product"
	models "riskmanagement/models/riskindicator"
	activity "riskmanagement/repository/activity"
	product "riskmanagement/repository/product"
	riskindicator "riskmanagement/repository/riskindicator"
	"strconv"
	"time"

	fileModel "riskmanagement/models/filemanager"
	filemanager "riskmanagement/services/filemanager"
	"strings"

	"github.com/google/uuid"
	"github.com/jung-kurt/gofpdf"
	"github.com/xuri/excelize/v2"

	"gitlab.com/golang-package-library/logger"
	minio "gitlab.com/golang-package-library/minio"
)

type RiskIndicatorDefinition interface {
	// WithTrx(trxHandle *gorm.DB) RiskIndicatorService
	GetAll() (responses []models.RiskIndicatorResponse, err error)
	GetOne(id int64) (responses models.RiskIndicatorGetOne, status bool, err error)
	GetAllWithPaginate(request models.Paginate) (responses []models.RiskIndicatorResponse, pagination lib.Pagination, err error)
	Store(request models.RiskIndicatorRequest) (responses bool, err error)
	Update(requests *models.RiskIndicatorRequest) (responses bool, err error)
	DeleteFilesByID(id int64) (response bool, err error)
	SearchRiskIndicatorByIssue(request models.SearchRequest) (responses []models.RiskIndicatorResponsesFinal, pagination lib.Pagination, err error)
	GetRekomendasiMateri(id int64) (responses []models.RekomendasiMateri, err error)
	SearchRiskIndicatorKRID(request models.KeyRiskRequest) (responses []models.RiskIndicatorKRIDResponses, pagination lib.Pagination, err error)
	Delete(request *models.UpdateDelete) (response bool, err error)
	GetKode() (response []models.KodeResponse, err error)
	FilterRiskIndicator(request models.FilterRequest) (responses []models.RiskIndicatorResponse, pagination lib.Pagination, err error)
	SaveThreshold(request models.RiskIndicatorRequest) (responses bool, err error)
	GetMappingThrehold(id int64) (responses []models.ThresholdIndicatorResponse, err error)
	GetMappingRiskIssue(id int64) (responses models.RiskIndicatorGetOne, err error)
	GetIndicatorByAktivityProduct(request models.IndicatorRequest) (responses []models.IndikatorResponse, err error)

	// Batch3
	SearchRiskIndicatorTematik(request models.SearchRequest) (responses []models.IndicatorTematikResponse, err error)
	GetTematikData(request models.TematikDataRequest) (responses []byte, err error)
	// GetTematikData(request models.TematikDataRequest) (responses models.TematikDataResponse, err error)

	GetMateriIfFinish(request models.RequestMateriIfFinish) (responses []models.RekomendasiMateri, err error)
	Template() ([]byte, string, error)
	Preview(pernr string, data [][]string) (dto.PreviewFileImport[[17]string], error)
	ImportData(pernr string, data [][]string) error
	Download(pernr string, format string) ([]byte, string, error)
	UpdateStatus(id int64) (err error)
}

type RiskIndicatorService struct {
	db                lib.Database
	minio             minio.Minio
	dbRaw             lib.Databases
	logger            logger.Logger
	activityRepo      activity.ActivityDefinition
	product           product.ProductDefinition
	riskIndicatorRepo riskindicator.RiskIndicatorDefinition
	lampiran          riskindicator.LampiranIndicatorDefinition
	filemanager       filemanager.FileManagerDefinition
	MapThreshold      riskindicator.MapThresholdDefinition
	MapRiskIssue      riskindicator.MapRiskIssueDefinition
}

func NewRiskIndicatorService(
	db lib.Database,
	minio minio.Minio,
	dbRaw lib.Databases,
	logger logger.Logger,
	activityRepo activity.ActivityDefinition,
	product product.ProductDefinition,
	riskIndicatorRepo riskindicator.RiskIndicatorDefinition,
	lampiran riskindicator.LampiranIndicatorDefinition,
	filemanager filemanager.FileManagerDefinition,
	MapThreshold riskindicator.MapThresholdDefinition,
	MapRiskIssue riskindicator.MapRiskIssueDefinition,
) RiskIndicatorDefinition {
	return RiskIndicatorService{
		db:                db,
		minio:             minio,
		dbRaw:             dbRaw,
		logger:            logger,
		riskIndicatorRepo: riskIndicatorRepo,
		lampiran:          lampiran,
		filemanager:       filemanager,
		MapThreshold:      MapThreshold,
		MapRiskIssue:      MapRiskIssue,
		activityRepo:      activityRepo,
		product:           product,
	}
}

func checkCount(rows *sql.Rows) (count int) {
	for rows.Next() {
		err := rows.Scan(&count)
		checkErr(err)
	}
	return count
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// GetAll implements RiskIndicatorDefinition
func (riskIndicator RiskIndicatorService) GetAll() (responses []models.RiskIndicatorResponse, err error) {
	return riskIndicator.riskIndicatorRepo.GetAll()
}

// GetKode implements RiskIndicatorDefinition
func (riskIndicator RiskIndicatorService) GetKode() (response []models.KodeResponse, err error) {
	dataIndicator, err := riskIndicator.riskIndicatorRepo.GetKode()
	if err != nil {
		riskIndicator.logger.Zap.Error(err)
		return response, err
	}

	for _, indicator := range dataIndicator {
		response = append(response, models.KodeResponse{
			Kode: indicator.Kode.String,
		})
	}

	return response, err
}

// Delete implements RiskIndicatorDefinition
func (ri RiskIndicatorService) Delete(request *models.UpdateDelete) (response bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	tx := ri.db.DB.Begin()

	getOneData, exist, err := ri.GetOne(request.ID)
	if err != nil {
		ri.logger.Zap.Error(err)
		tx.Rollback()
		return false, err
	}

	updateDataIndicator := &models.UpdateDelete{
		ID:         request.ID,
		DeleteFlag: true,
		UpdatedAt:  &timeNow,
	}

	_, err = ri.riskIndicatorRepo.Delete(updateDataIndicator,
		[]string{
			"delete_flag",
			"updated_at",
		}, tx)

	if err != nil {
		tx.Rollback()
		ri.logger.Zap.Error(err)
		return false, err
	}

	if exist {
		fmt.Println("getOne", getOneData)
		tx.Commit()
		return true, err
	}

	return false, err
}

// GetAllWithPaginate implements RiskIndicatorDefinition
func (ri RiskIndicatorService) GetAllWithPaginate(request models.Paginate) (responses []models.RiskIndicatorResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)

	request.Offset = offset
	request.Order = order
	request.Sort = sort
	request.Limit = limit

	ri.logger.Zap.Debug(request)

	dataPgs, totalRows, totalData, err := ri.riskIndicatorRepo.GetAllWithPaginate(&request)
	if err != nil {
		ri.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataPgs {
		responses = append(responses, models.RiskIndicatorResponse{
			ID:                    response.ID,
			RiskIndicatorCode:     response.RiskIndicatorCode,
			RiskIndicator:         response.RiskIndicator,
			ActivityID:            response.ActivityID,
			ProductID:             response.ProductID,
			Deskripsi:             response.Deskripsi,
			Satuan:                response.Satuan,
			Sifat:                 response.Sifat,
			BusinessCycleActivity: response.BusinessCycleActivity,
			Batasan:               response.Batasan,
			Kondisi:               response.Kondisi,
			Type:                  response.Type,
			SLAVerifikasi:         response.SLAVerifikasi,
			SLATindakLanjut:       response.SLATindakLanjut,
			SumberData:            response.SumberData,
			SumberDataText:        response.SumberDataText,
			PeriodePemantauan:     response.PeriodePemantauan,
			Owner:                 response.Owner,
			KPI:                   response.KPI,
			StatusIndikator:       response.StatusIndikator,
			DataSourceAnomaly:     response.DataSourceAnomaly,
			Status:                response.Status,
			CreatedAt:             response.CreatedAt,
			UpdatedAt:             response.UpdatedAt,
		})
	}

	ri.logger.Zap.Debug(responses)
	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err
}

func (rc RiskIndicatorService) UpdateStatus(id int64) (err error) {
	var (
		status bool = true
	)
	data, err := rc.riskIndicatorRepo.GetOne(id)
	if err != nil {
		return err
	}

	if data.Status {
		status = false
	}

	err = rc.riskIndicatorRepo.UpdateStatus(id, status)

	return err
}

// GetOne implements RiskIndicatorDefinition
func (riskIndicator RiskIndicatorService) GetOne(id int64) (responses models.RiskIndicatorGetOne, status bool, err error) {
	dataRiskIndicator, err := riskIndicator.riskIndicatorRepo.GetOne(id)
	if dataRiskIndicator.ID != 0 {
		fmt.Println("bukan 0")
		data_lampiran, err := riskIndicator.lampiran.GetOneFileByID(dataRiskIndicator.ID)

		var minioLink []models.MinioLink
		var index int64
		for _, value := range data_lampiran {
			mini_link, err := riskIndicator.filemanager.GetFile(fileModel.FileManagerRequest{
				Subdir:   value.Path,
				Filename: value.Filename,
			})

			if err != nil {
				riskIndicator.logger.Zap.Error(err)
			}

			minioLink = append(minioLink, models.MinioLink{
				Index:    index,
				Filepath: mini_link.MinioPath,
			})

			index++
		}

		responses = models.RiskIndicatorGetOne{
			ID:                dataRiskIndicator.ID,
			RiskIndicatorCode: dataRiskIndicator.RiskIndicatorCode,
			RiskIndicator:     dataRiskIndicator.RiskIndicator,
			ActivityID:        dataRiskIndicator.ActivityID,
			ProductID:         dataRiskIndicator.ProductID,
			Deskripsi:         dataRiskIndicator.Deskripsi,
			Satuan:            dataRiskIndicator.Satuan,
			Sifat:             dataRiskIndicator.Sifat,
			SLAVerifikasi:     dataRiskIndicator.SLAVerifikasi,
			SLATindakLanjut:   dataRiskIndicator.SLATindakLanjut,
			SumberData:        dataRiskIndicator.SumberData,
			SumberDataText:    dataRiskIndicator.SumberDataText,
			PeriodePemantauan: dataRiskIndicator.PeriodePemantauan,
			Owner:             dataRiskIndicator.Owner,
			KPI:               dataRiskIndicator.KPI,
			StatusIndikator:   dataRiskIndicator.StatusIndikator,
			DataSourceAnomaly: dataRiskIndicator.DataSourceAnomaly,
			Status:            dataRiskIndicator.Status,
			LampiranIndicator: data_lampiran,
			MinioLink:         minioLink,
			CreatedAt:         dataRiskIndicator.CreatedAt,
			UpdatedAt:         dataRiskIndicator.UpdatedAt,
		}

		return responses, true, err
	}

	return responses, false, err
}

// Store implements RiskIndicatorDefinition
func (riskIndicator RiskIndicatorService) Store(request models.RiskIndicatorRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	rowsCheckCode, err := riskIndicator.dbRaw.DB.Query("SELECT COUNT(*) as count FROM risk_indicator ri WHERE ri.risk_indicator_code = ? AND ri.delete_flag = 0", request.RiskIndicatorCode)

	checkErr(err)

	if checkCount(rowsCheckCode) < 1 {
		fmt.Println("masook")
		fmt.Println("jumlah ==>", checkCount(rowsCheckCode))

		tx := riskIndicator.db.DB.Begin()

		reqRiskIndicator := &models.RiskIndicator{
			RiskIndicatorCode:     request.RiskIndicatorCode,
			RiskIndicator:         request.RiskIndicator,
			ActivityID:            request.ActivityID,
			ProductID:             request.ProductID,
			Deskripsi:             request.Deskripsi,
			Satuan:                request.Satuan,
			Sifat:                 request.Sifat,
			BusinessCycleActivity: request.BusinessCycleActivity,
			Batasan:               request.Batasan,
			Kondisi:               request.Kondisi,
			Type:                  request.Type,
			SLAVerifikasi:         request.SLAVerifikasi,
			SLATindakLanjut:       request.SLATindakLanjut,
			SumberData:            request.SumberData,
			SumberDataText:        request.SumberDataText,
			PeriodePemantauan:     request.PeriodePemantauan,
			Owner:                 request.Owner,
			KPI:                   request.KPI,
			StatusIndikator:       request.StatusIndikator,
			DataSourceAnomaly:     request.DataSourceAnomaly,
			Status:                request.Status,
			CreatedAt:             &timeNow,
		}

		dataRiskIndicator, err := riskIndicator.riskIndicatorRepo.Store(reqRiskIndicator, tx)
		fmt.Println("data => ", dataRiskIndicator)

		if err != nil {
			tx.Rollback()
			riskIndicator.logger.Zap.Error(err)
			return false, err
		}

		bucket := os.Getenv("BUCKET_NAME")
		if len(request.LampiranIndicator) != 0 {
			for _, value := range request.LampiranIndicator {
				switch value.JenisFile {
				case "Upload Document":
					fmt.Println("upload")
					UUID := uuid.New()

					var destinationPath string
					bucketExist := riskIndicator.minio.BucketExist(riskIndicator.minio.Client(), bucket)

					pathSplit := strings.Split(value.Path, "/")
					sourchePath := fmt.Sprint(value.Path)
					destinationPath = pathSplit[1] + "/" +
						lib.GetTimeNow("year") + "/" +
						lib.GetTimeNow("month") + "/" +
						lib.GetTimeNow("day") + "/" +
						UUID.String() + "/" +
						pathSplit[2] + "/" + value.Filename

					if bucketExist {
						fmt.Println("Exist")
						fmt.Println(bucket)
						fmt.Println(sourchePath)
						fmt.Println(destinationPath)
						uploaded := riskIndicator.minio.CopyObject(riskIndicator.minio.Client(), bucket, sourchePath, bucket, destinationPath)

						fmt.Println(uploaded)
					} else {
						fmt.Println("Not Exist")
						fmt.Println(bucket)
						fmt.Println(sourchePath)
						fmt.Println(destinationPath)
						riskIndicator.minio.MakeBucket(riskIndicator.minio.Client(), bucket, "")
						uploaded := riskIndicator.minio.CopyObject(riskIndicator.minio.Client(), bucket, sourchePath, bucket, destinationPath)

						fmt.Println(uploaded)
					}

					_, err = riskIndicator.lampiran.Store(&models.LampiranIndicator{
						ID:            value.ID,
						IDIndicator:   reqRiskIndicator.ID,
						NamaLampiran:  value.NamaLampiran,
						NomorLampiran: value.NomorLampiran,
						JenisFile:     value.JenisFile,
						Path:          destinationPath,
						Filename:      value.Filename,
					}, tx)

					if err != nil {
						tx.Rollback()
						riskIndicator.logger.Zap.Error(err)
						return false, err
					}
				case "Link Document":
					fmt.Println("link")
					_, err = riskIndicator.lampiran.Store(&models.LampiranIndicator{
						ID:            value.ID,
						IDIndicator:   reqRiskIndicator.ID,
						NamaLampiran:  value.NamaLampiran,
						NomorLampiran: value.NomorLampiran,
						JenisFile:     value.JenisFile,
						Path:          value.Path,
					}, tx)

					if err != nil {
						tx.Rollback()
						riskIndicator.logger.Zap.Error(err)
						return false, err
					}
				}
			}
		} else {
			if err != nil {
				tx.Rollback()
				riskIndicator.logger.Zap.Error(err)
				return false, err
			}
		}

		tx.Commit()

		return true, err

	}
	fmt.Println("gagal")
	// riskIndicator.logger.Zap.Error(err)
	return false, err

}

// Update implements RiskIndicatorDefinition
func (riskIndicator RiskIndicatorService) Update(requests *models.RiskIndicatorRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	tx := riskIndicator.db.DB.Begin()

	updateIndicator := &models.RiskIndicator{
		RiskIndicatorCode:     requests.RiskIndicatorCode,
		RiskIndicator:         requests.RiskIndicator,
		ActivityID:            requests.ActivityID,
		ProductID:             requests.ProductID,
		Deskripsi:             requests.Deskripsi,
		Satuan:                requests.Satuan,
		Sifat:                 requests.Sifat,
		BusinessCycleActivity: requests.BusinessCycleActivity,
		Batasan:               requests.Batasan,
		Kondisi:               requests.Kondisi,
		Type:                  requests.Type,
		SLAVerifikasi:         requests.SLAVerifikasi,
		SLATindakLanjut:       requests.SLATindakLanjut,
		SumberData:            requests.SumberData,
		SumberDataText:        requests.SumberDataText,
		PeriodePemantauan:     requests.PeriodePemantauan,
		Owner:                 requests.Owner,
		KPI:                   requests.KPI,
		StatusIndikator:       requests.StatusIndikator,
		DataSourceAnomaly:     requests.DataSourceAnomaly,
		Status:                requests.Status,
		UpdatedAt:             &timeNow,
	}

	include := []string{
		"risk_indicator_code",
		"risk_indicator",
		"activity_id",
		"product_id",
		"deskripsi",
		"satuan",
		"sifat",
		"sla_verifikasi",
		"sla_tindak_lanjut",
		"sumber_data",
		"sumber_data_text",
		"periode_pemantauan",
		"owner",
		"kpi",
		"status_indikator",
		"status",
		"updated_at",
	}

	_, err = riskIndicator.riskIndicatorRepo.Update(updateIndicator, include, tx)

	if err != nil {
		tx.Rollback()
		riskIndicator.logger.Zap.Error(err)
		return false, err
	}

	bucket := os.Getenv("BUCKET_NAME")

	if len(requests.LampiranIndicator) != 0 {
		err := riskIndicator.lampiran.DeleteFilesByIndicator(requests.ID, tx)

		if err != nil {
			tx.Rollback()
			riskIndicator.logger.Zap.Error(err)
			return false, err
		}

		for _, value := range requests.LampiranIndicator {
			switch value.JenisFile {
			case "Upload Document":
				fmt.Println("upload")
				UUID := uuid.New()

				var destinationPath string
				bucketExist := riskIndicator.minio.BucketExist(riskIndicator.minio.Client(), bucket)

				pathSplit := strings.Split(value.Path, "/")
				sourchePath := fmt.Sprint(value.Path)
				destinationPath = pathSplit[1] + "/" +
					lib.GetTimeNow("year") + "/" +
					lib.GetTimeNow("month") + "/" +
					lib.GetTimeNow("day") + "/" +
					UUID.String() + "/" +
					pathSplit[2] + "/" + value.Filename

				if pathSplit[0] == "tmp" {
					riskIndicator.logger.Zap.Info("======> New Files")

					if bucketExist {
						fmt.Println("Exist")
						fmt.Println(bucket)
						fmt.Println(sourchePath)
						fmt.Println(destinationPath)
						uploaded := riskIndicator.minio.CopyObject(riskIndicator.minio.Client(), bucket, sourchePath, bucket, destinationPath)

						fmt.Println(uploaded)
					} else {
						fmt.Println("Not Exist")
						fmt.Println(bucket)
						fmt.Println(sourchePath)
						fmt.Println(destinationPath)
						riskIndicator.minio.MakeBucket(riskIndicator.minio.Client(), bucket, "")
						uploaded := riskIndicator.minio.CopyObject(riskIndicator.minio.Client(), bucket, sourchePath, bucket, destinationPath)

						fmt.Println(uploaded)
					}

					_, err = riskIndicator.lampiran.Store(&models.LampiranIndicator{
						ID:            value.ID,
						IDIndicator:   requests.ID,
						NamaLampiran:  value.NamaLampiran,
						NomorLampiran: value.NomorLampiran,
						JenisFile:     value.JenisFile,
						Path:          destinationPath,
						Filename:      value.Filename,
					}, tx)

					if err != nil {
						tx.Rollback()
						riskIndicator.logger.Zap.Error(err)
						return false, err
					}
				} else {
					riskIndicator.logger.Zap.Info("======> Old Files")
					_, err = riskIndicator.lampiran.Store(&models.LampiranIndicator{
						ID:            value.ID,
						IDIndicator:   requests.ID,
						NamaLampiran:  value.NamaLampiran,
						NomorLampiran: value.NomorLampiran,
						JenisFile:     value.JenisFile,
						Path:          value.Path,
						Filename:      value.Filename,
					}, tx)

					if err != nil {
						tx.Rollback()
						riskIndicator.logger.Zap.Error(err)
						return false, err
					}
				}

			case "Link Document":
				fmt.Println("link")
				_, err = riskIndicator.lampiran.Store(&models.LampiranIndicator{
					ID:            value.ID,
					IDIndicator:   requests.ID,
					NamaLampiran:  value.NamaLampiran,
					NomorLampiran: value.NomorLampiran,
					JenisFile:     value.JenisFile,
					Path:          value.Path,
					Filename:      value.NamaLampiran,
				}, tx)

				if err != nil {
					tx.Rollback()
					riskIndicator.logger.Zap.Error(err)
					return false, err
				}
			}
		}
	} else {
		tx.Rollback()
		riskIndicator.logger.Zap.Error(err)
		return false, err
	}

	tx.Commit()
	return true, err
}

// DeleteFilesByID implements RiskIndicatorDefinition
func (riskIndicator RiskIndicatorService) DeleteFilesByID(id int64) (response bool, err error) {
	dataFiles, err := riskIndicator.lampiran.GetOne(id)

	switch dataFiles.JenisFile {
	case "Upload Document":
		fmt.Println("data => ", dataFiles.Path)
		bucket := os.Getenv("BUCKET_NAME")
		objectName := dataFiles.Path

		bucketExist := riskIndicator.minio.BucketExist(riskIndicator.minio.Client(), bucket)
		if bucketExist {
			remove := riskIndicator.minio.RemoveObject(riskIndicator.minio.Client(), bucket, objectName)
			if remove {
				riskIndicator.lampiran.Delete(dataFiles.ID)
				return true, err
			} else {
				return false, err
			}
		} else {
			return false, err
		}
	case "Link Document":
		riskIndicator.lampiran.Delete(dataFiles.ID)
	}

	return true, err
}

// SearchRiskIndicatorByIssue implements RiskIndicatorDefinition
func (LI RiskIndicatorService) SearchRiskIndicatorByIssue(request models.SearchRequest) (responses []models.RiskIndicatorResponsesFinal, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	dataIndicator, totalRows, totalData, err := LI.riskIndicatorRepo.SearchRiskIndicatorByIssue(&request)
	if err != nil {
		LI.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataIndicator {
		responses = append(responses, models.RiskIndicatorResponsesFinal{
			ID:                response.ID,
			RiskIndicatorCode: response.RiskIndicatorCode,
			RiskIndicator:     response.RiskIndicator,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err
}

// GetRekomendasiMateri implements RiskIndicatorDefinition
func (LI RiskIndicatorService) GetRekomendasiMateri(id int64) (responses []models.RekomendasiMateri, err error) {
	Materi, err := LI.riskIndicatorRepo.GetRekomendasiMateri(id)
	if err != nil {
		LI.logger.Zap.Error()
		return responses, err
	}

	for _, response := range Materi {
		responses = append(responses, models.RekomendasiMateri{
			ID:           response.ID.Int64,
			IDIndicator:  response.IDIndicator.Int64,
			NamaLampiran: response.NamaLampiran.String,
			Filename:     response.Filename.String,
			Path:         response.Path.String,
		})
	}

	return responses, err
}

// SearchRiskIndicatorKRID implements RiskIndicatorDefinition
func (LI RiskIndicatorService) SearchRiskIndicatorKRID(request models.KeyRiskRequest) (responses []models.RiskIndicatorKRIDResponses, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	dataIndicator, totalRows, totalData, err := LI.riskIndicatorRepo.SearchRiskIndicatorKRID(&request)
	if err != nil {
		LI.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataIndicator {
		responses = append(responses, models.RiskIndicatorKRIDResponses{
			ID:                   response.ID.Int64,
			KodeKeyRiskIndicator: response.KodeKeyRiskIndicator.Int64,
			KeyRiskIndicator:     response.KeyRiskIndicator.String,
			Aktifitas:            response.Aktifitas.String,
			Produk:               response.Produk.String,
			JenisIndicator:       response.JenisIndicator.String,
			IndikasiRisiko:       response.IndikasiRisiko.String,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err
}

// FilterRiskIndicator implements RiskIndicatorDefinition
func (ri RiskIndicatorService) FilterRiskIndicator(request models.FilterRequest) (responses []models.RiskIndicatorResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	dataRiskIndicator, totalRows, totalData, err := ri.riskIndicatorRepo.FilterRiskIndicator(&request)
	if err != nil {
		ri.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataRiskIndicator {
		responses = append(responses, models.RiskIndicatorResponse{
			ID:                response.ID,
			RiskIndicatorCode: response.RiskIndicatorCode,
			RiskIndicator:     response.RiskIndicator,
			ActivityID:        response.ActivityID,
			ProductID:         response.ProductID,
			Deskripsi:         response.Deskripsi,
			Satuan:            response.Satuan,
			Sifat:             response.Sifat,
			SLAVerifikasi:     response.SLAVerifikasi,
			SLATindakLanjut:   response.SLATindakLanjut,
			SumberData:        response.SumberData,
			SumberDataText:    response.SumberDataText,
			PeriodePemantauan: response.PeriodePemantauan,
			Owner:             response.Owner,
			KPI:               response.KPI,
			StatusIndikator:   response.StatusIndikator,
			Status:            response.Status,
			CreatedAt:         response.CreatedAt,
			UpdatedAt:         response.UpdatedAt,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err
}

// MapThreshold implements RiskIndicatorDefinition
func (ri RiskIndicatorService) SaveThreshold(request models.RiskIndicatorRequest) (responses bool, err error) {
	tx := ri.db.DB.Begin()

	if len(request.MapThreshold) != 0 {
		for _, value := range request.MapThreshold {
			_, err = ri.MapThreshold.SaveThreshold(&models.MapThreshold{
				ID:          value.ID,
				IDIndicator: request.ID,
				KCK_1_MIN:   value.KCK_1_MIN,
				KCK_2_MIN:   value.KCK_2_MIN,
				KCK_3_MIN:   value.KCK_3_MIN,
				KCK_4_MIN:   value.KCK_4_MIN,
				KCK_5_MIN:   value.KCK_5_MIN,
				KC_1_MIN:    value.KC_1_MIN,
				KC_2_MIN:    value.KC_2_MIN,
				KC_3_MIN:    value.KC_3_MIN,
				KC_4_MIN:    value.KC_4_MIN,
				KC_5_MIN:    value.KC_5_MIN,
				KCP_1_MIN:   value.KCP_1_MIN,
				KCP_2_MIN:   value.KCP_2_MIN,
				KCP_3_MIN:   value.KCP_3_MIN,
				KCP_4_MIN:   value.KCP_4_MIN,
				KCP_5_MIN:   value.KCP_5_MIN,
				UN_1_MIN:    value.UN_1_MIN,
				UN_2_MIN:    value.UN_2_MIN,
				UN_3_MIN:    value.UN_3_MIN,
				UN_4_MIN:    value.UN_4_MIN,
				UN_5_MIN:    value.UN_5_MIN,
				KK_1_MIN:    value.KK_1_MIN,
				KK_2_MIN:    value.KK_2_MIN,
				KK_3_MIN:    value.KK_3_MIN,
				KK_4_MIN:    value.KK_4_MIN,
				KK_5_MIN:    value.KK_5_MIN,
				KCK_1_MAX:   value.KCK_1_MAX,
				KCK_2_MAX:   value.KCK_2_MAX,
				KCK_3_MAX:   value.KCK_3_MAX,
				KCK_4_MAX:   value.KCK_4_MAX,
				KCK_5_MAX:   value.KCK_5_MAX,
				KC_1_MAX:    value.KC_1_MAX,
				KC_2_MAX:    value.KC_2_MAX,
				KC_3_MAX:    value.KC_3_MAX,
				KC_4_MAX:    value.KC_4_MAX,
				KC_5_MAX:    value.KC_5_MAX,
				KCP_1_MAX:   value.KCP_1_MAX,
				KCP_2_MAX:   value.KCP_2_MAX,
				KCP_3_MAX:   value.KCP_3_MAX,
				KCP_4_MAX:   value.KCP_4_MAX,
				KCP_5_MAX:   value.KCP_5_MAX,
				UN_1_MAX:    value.UN_1_MAX,
				UN_2_MAX:    value.UN_2_MAX,
				UN_3_MAX:    value.UN_3_MAX,
				UN_4_MAX:    value.UN_4_MAX,
				UN_5_MAX:    value.UN_5_MAX,
				KK_1_MAX:    value.KK_1_MAX,
				KK_2_MAX:    value.KK_2_MAX,
				KK_3_MAX:    value.KK_3_MAX,
				KK_4_MAX:    value.KK_4_MAX,
				KK_5_MAX:    value.KK_5_MAX,
			}, tx)

			if err != nil {
				tx.Rollback()
				ri.logger.Zap.Error(err)
				return false, err
			}
		}
	} else {
		tx.Rollback()
		ri.logger.Zap.Error(err)
		return false, err
	}

	tx.Commit()
	return true, err
}

// GetMappingThrehold implements RiskIndicatorDefinition
func (ri RiskIndicatorService) GetMappingThrehold(id int64) (responses []models.ThresholdIndicatorResponse, err error) {
	dataIndikator, err := ri.riskIndicatorRepo.GetDataThreshold(id)

	if dataIndikator[0].Index != 0 {

		for _, response := range dataIndikator {
			dataThreshold, err := ri.MapThreshold.GetThreshold(response.Index)
			if err != nil {
				ri.logger.Zap.Error(err)
				return responses, err
			}

			responses = append(responses, models.ThresholdIndicatorResponse{
				Index:            response.Index,
				Id:               response.Id,
				KeyRiskIndicator: response.KeyRiskIndicator,
				Aktivitas:        response.Aktivitas,
				Produk:           response.Produk,
				JenisIndikator:   response.JenisIndikator,
				IndikasiRisiko:   response.IndikasiRisiko,
				Deskripsi:        response.Deskripsi,
				SlaVerifikasi:    response.SlaVerifikasi,
				SlaTl:            response.SlaTl,
				RiskAwarness:     response.RiskAwarness,
				DataSource:       response.DataSource,
				Parameter:        response.Parameter,
				StatusIndikator:  response.StatusIndikator,
				IsAktif:          response.IsAktif,
				MapThreshold:     dataThreshold,
			})
		}

		return responses, err

	}

	return responses, err
}

// GetMappingRiskIssue implements RiskIndicatorDefinition
func (ri RiskIndicatorService) GetMappingRiskIssue(id int64) (responses models.RiskIndicatorGetOne, err error) {
	dataIndikator, err := ri.riskIndicatorRepo.GetOne(id)

	if dataIndikator.ID != 0 {
		dataRiskIssue, err := ri.MapRiskIssue.GetRiskIssue(dataIndikator.ID)

		responses = models.RiskIndicatorGetOne{
			ID:                dataIndikator.ID,
			RiskIndicatorCode: dataIndikator.RiskIndicatorCode,
			RiskIndicator:     dataIndikator.RiskIndicator,
			Satuan:            dataIndikator.Satuan,
			MapRiskIssue:      dataRiskIssue,
		}

		return responses, err
	}

	return responses, err
}

// GetIndicatorByAktivityProduct implements RiskIndicatorDefinition
func (ri RiskIndicatorService) GetIndicatorByAktivityProduct(request models.IndicatorRequest) (responses []models.IndikatorResponse, err error) {
	dataIndikator, err := ri.riskIndicatorRepo.GetIndicatorByAktivityProduct(&request)
	if err != nil {
		ri.logger.Zap.Error(err)
		return responses, err
	}

	for _, res := range dataIndikator {
		responses = append(responses, models.IndikatorResponse{
			ID:                res.ID,
			RiskIndicatorCode: res.RiskIndicatorCode,
			RiskIndicator:     res.RiskIndicator,
		})
	}

	return responses, err
}

// SearchRiskIndicatorTematik implements RiskIndicatorDefinition.
func (ri RiskIndicatorService) SearchRiskIndicatorTematik(request models.SearchRequest) (responses []models.IndicatorTematikResponse, err error) {
	dataIndicator, err := ri.riskIndicatorRepo.SearchRiskIndicatorTematik(&request)

	if err != nil {
		ri.logger.Zap.Error(err)
		return responses, err
	}

	for _, value := range dataIndicator {
		responses = append(responses, models.IndicatorTematikResponse{
			Id:            value.Id,
			RiskIndicator: value.RiskIndicator,
			NamaTable:     value.NamaTable,
		})
	}

	return responses, err
}

func (ri RiskIndicatorService) GetTematikData(request models.TematikDataRequest) (responses []byte, err error) {
	fmt.Println("masuk service")

	dataResponse, err := ri.riskIndicatorRepo.GetTematikData(&request)

	return dataResponse, err
}

/*
// GetTematikData implements RiskIndicatorDefinition.
func (ri RiskIndicatorService) GetTematikData(request models.TematikDataRequest) (responses models.TematikDataResponse, err error) {
	dataResponse, err := ri.riskIndicatorRepo.GetTematikData(&request)

	if err != nil {
		ri.logger.Zap.Error(err)
		return responses, err
	}

	var data []string

	for _, val := range dataResponse.Data {
		data = append(data, val)
	}

	responses = models.TematikDataResponse{
		Header: dataResponse.Header,
		Data:   data,
	}

	return responses, err
}
*/

// GetMateriIfFinish implements RiskIndicatorDefinition.
func (ri RiskIndicatorService) GetMateriIfFinish(request models.RequestMateriIfFinish) (responses []models.RekomendasiMateri, err error) {
	document, err := ri.riskIndicatorRepo.GetMateriIfFinish(&request)

	return document, err
}

func (ri RiskIndicatorService) Template() ([]byte, string, error) {
	f := excelize.NewFile()
	sheet := "Template"

	f.SetSheetName("Sheet1", sheet)

	sheetIndex, err := f.GetSheetIndex(sheet)
	if err != nil {
		return nil, "", fmt.Errorf("failed to get sheet index: %w", err)
	}
	f.SetActiveSheet(sheetIndex)

	headers := []string{
		"Data Source Anomaly",
		"Risk Indicator Code",
		"Risk Indicator",
		"Aktivitas",
		"Produk",
		"Deskripsi",
		"satuan",
		"Periode",
		"Batasan",
		"Business Cycle Aktivitas",
		"SLA Verifikasi",
		"SLA Tindak Lanjut",
		"KPI",
		"Tipe Data Anomali",
		"Kondisi",
		"Mapping Set Header",
		"Mapping Set Key Uker",
	}

	// tulis header
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		if err := f.SetCellValue(sheet, cell, h); err != nil {
			return nil, "", fmt.Errorf("failed to set cell: %w", err)
		}
	}

	exampleData := [][]string{
		{
			"Manual",
			"RI-001",
			"Jumlah Transaksi",
			"Pembayaran",
			"Produk A",
			"Indikator untuk jumlah transaksi harian",
			"Kali",
			"Bulanan",
			"Restricted",
			"Proses Bisnis A",
			"2",
			"5",
			"Meningkatkan efisiensi",
			"Posisi",
			"Tampilkan Semua & Verifikasi Terbaru",
			"",
			"",
		},
		{
			"Tematik",
			"RI-002",
			"Nilai Penyimpangan",
			"Pengiriman",
			"Produk B",
			"Indikator penyimpangan nilai harian",
			"Persen",
			"Mingguan",
			"Restricted",
			"Proses Bisnis B",
			"1",
			"3",
			"Menurunkan error operasional",
			"Posisi",
			"Tampilkan Semua & Verifikasi Terbaru",
			"Header-X;Header-Y",
			"Uker-010;Uker-011",
		},
	}

	for rowIndex, row := range exampleData {
		for colIndex, val := range row {
			cell, _ := excelize.CoordinatesToCellName(colIndex+1, rowIndex+2) // +2 karena header di baris 1
			if err := f.SetCellValue(sheet, cell, val); err != nil {
				return nil, "", fmt.Errorf("failed to set example data: %w", err)
			}
		}
	}

	// optional: set lebar kolom
	for i := 1; i <= len(headers); i++ {
		col, _ := excelize.ColumnNumberToName(i)
		f.SetColWidth(sheet, col, col, 25)
	}

	// contoh data
	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		return nil, "", fmt.Errorf("failed to write excel: %w", err)
	}

	return buf.Bytes(), "risk_indicator_template.xlsx", nil
}

func (ri RiskIndicatorService) Preview(pernr string, data [][]string) (dto.PreviewFileImport[[17]string], error) {
	indicator, err := ri.riskIndicatorRepo.GetAll()
	if err != nil {
		ri.logger.Zap.Error("Errored to query activity: %s", err)
		return dto.PreviewFileImport[[17]string]{}, err
	}

	indicatorMap := make(map[string]bool, len(indicator))
	for _, a := range indicator {
		indicatorMap[strings.ToLower(a.RiskIndicatorCode)] = true
	}

	activityList, err := ri.activityRepo.GetAll()
	if err != nil {
		ri.logger.Zap.Error("Errored to query activity: %s", err)
		return dto.PreviewFileImport[[17]string]{}, err
	}

	activityMap := make(map[string]modelActiv.ActivityResponse, len(activityList))
	for _, a := range activityList {
		activityMap[strings.ToLower(a.Name)] = a
	}

	productList, err := ri.product.GetAll()
	if err != nil {
		ri.logger.Zap.Error("Errored to query product: %s", err)
		return dto.PreviewFileImport[[17]string]{}, err
	}

	productMap := make(map[string]modelProduct.ProductResponse, len(productList))
	for _, p := range productList {
		productMap[strings.ToLower(p.Product)] = p
	}

	headers := []string{
		"Data Source Anomaly",
		"Risk Indicator Code",
		"Risk Indicator",
		"Aktivitas",
		"Produk",
		"Deskripsi",
		"satuan",
		"Periode",
		"Batasan",
		"Business Cycle Aktivitas",
		"SLA Verifikasi",
		"SLA Tindak Lanjut",
		"KPI",
		"Tipe Data Anomali",
		"Kondisi",
		"Mapping Set Header",
		"Mapping Set Key Uker",
	}

	previewFile := dto.PreviewFileImport[[17]string]{}
	body := []dto.PreviewFile[[17]string]{}

	for index, row := range data {
		if index == 0 {
			if len(row) < 17 {
				return dto.PreviewFileImport[[17]string]{}, fmt.Errorf("invalid header format risk indicator")
			}

			for i, v := range headers {
				if strings.TrimSpace(row[i]) != v {
					return dto.PreviewFileImport[[17]string]{}, fmt.Errorf("header kolom ke-%d invalid format, diharapkan '%s', diterima '%s'", i+1, v, row[i])
				}
			}

			previewFile.Header = [17]string{
				row[0],
				row[1],
				row[2],
				row[3],
				row[4],
				row[5],
				row[6],
				row[7],
				row[8],
				row[9],
				row[10],
				row[11],
				row[12],
				row[13],
				row[14],
				row[15],
				row[16],
			}

			continue
		}

		var (
			col [17]string
		)

		validation := ""

		activityName := strings.ToLower(row[3])
		productName := strings.ToLower(row[4])

		if _, ok := indicatorMap[strings.ToLower(row[1])]; ok {
			validation += fmt.Sprintf("Code Indicator sudah terdaftar: %s; ", row[0])
		}

		activityData, activityStat := activityMap[activityName]
		if !activityStat {
			validation += fmt.Sprintf("Activity tidak terdaftar: %s; ", activityName)
		}

		productData, productStat := productMap[productName]
		if !productStat {
			validation += fmt.Sprintf("Product tidak terdaftar: %s; ", productName)
		}

		var productActivityID string
		if productData.ActivityID != nil {
			productActivityID = *productData.ActivityID
		}

		activityIDStr := strconv.FormatInt(activityData.ID, 10)

		if productActivityID != activityIDStr {
			validation += fmt.Sprintf("Product tidak sesuai dengan Activity: %s; ", row[1])
		}

		for i := range 17 {
			if i < len(row) {
				col[i] = row[i]
			}
		}

		body = append(body, dto.PreviewFile[[17]string]{
			PerRow:     col,
			Validation: validation,
		})

	}
	previewFile.Body = body

	return previewFile, nil
}

func (ri RiskIndicatorService) ImportData(pernr string, data [][]string) error {
	indicator, err := ri.riskIndicatorRepo.GetAll()
	if err != nil {
		ri.logger.Zap.Error("Errored to query activity: %s", err)
		return err
	}

	indicatorMap := make(map[string]bool, len(indicator))
	for _, a := range indicator {
		indicatorMap[strings.ToLower(a.RiskIndicatorCode)] = true
	}

	activityList, err := ri.activityRepo.GetAll()
	if err != nil {
		ri.logger.Zap.Error("Errored to query activity: %s", err)
		return err
	}

	activityMap := make(map[string]modelActiv.ActivityResponse, len(activityList))
	for _, a := range activityList {
		activityMap[strings.ToLower(a.Name)] = a
	}

	productList, err := ri.product.GetAll()
	if err != nil {
		ri.logger.Zap.Error("Errored to query product: %s", err)
		return err
	}

	productMap := make(map[string]modelProduct.ProductResponse, len(productList))
	for _, p := range productList {
		productMap[strings.ToLower(p.Product)] = p
	}

	headers := []string{
		"Data Source Anomaly",
		"Risk Indicator Code",
		"Risk Indicator",
		"Aktivitas",
		"Produk",
		"Deskripsi",
		"satuan",
		"Periode",
		"Batasan",
		"Business Cycle Aktivitas",
		"SLA Verifikasi",
		"SLA Tindak Lanjut",
		"KPI",
		"Tipe Data Anomali",
		"Kondisi",
		"Mapping Set Header",
		"Mapping Set Key Uker",
	}

	newRecord := make([]models.RiskIndicator, 0)

	for index, row := range data {
		if index == 0 {
			if len(row) < 17 {
				return fmt.Errorf("invalid header format risk indicator")
			}

			for i, v := range headers {
				if strings.TrimSpace(row[i]) != v {
					return fmt.Errorf("header kolom ke-%d invalid format, diharapkan '%s', diterima '%s'", i+1, v, row[i])
				}
			}

			continue

		}

		var (
			indicatorExists      bool  = false
			activityExist        bool  = false
			productExist         bool  = false
			validActivityProduct bool  = false
			currentActivituID    int64 = 0
			currentProductID     int64 = 0
		)

		activityName := strings.ToLower(row[3])
		productName := strings.ToLower(row[4])

		if _, ok := indicatorMap[strings.ToLower(row[0])]; !ok {
			indicatorExists = true
		}

		activityData, activityStat := activityMap[activityName]
		if activityStat {
			currentActivituID = activityData.ID
			activityExist = true
		}

		productData, productStat := productMap[productName]
		if productStat {
			currentProductID = productData.ID
			productExist = true
		}

		var productActivityID string
		if productData.ActivityID != nil {
			productActivityID = *productData.ActivityID
		}

		activityIDStr := strconv.FormatInt(activityData.ID, 10)

		if productActivityID == activityIDStr {
			validActivityProduct = true
		}

		if indicatorExists && activityExist && productExist && validActivityProduct {
			newRecord = append(newRecord, models.RiskIndicator{
				RiskIndicatorCode:     row[1],
				RiskIndicator:         row[2],
				ActivityID:            currentActivituID,
				ProductID:             currentProductID,
				Deskripsi:             row[5],
				Satuan:                row[6],
				BusinessCycleActivity: row[9],
				Batasan:               row[8],
				Kondisi:               ConditionMap(row[14]),
				SLAVerifikasi:         lib.ToInt64(row[10]),
				SLATindakLanjut:       lib.ToInt64(row[11]),
				PeriodePemantauan:     row[7],
				KPI:                   row[12],
				Type:                  row[13],
				DataSourceAnomaly:     strings.ToLower(row[0]),
			})
		}
	}

	if len(newRecord) > 0 {
		tx := ri.db.DB.Begin()
		err := ri.riskIndicatorRepo.BulkCreateRiskIndicator(newRecord, tx)
		if err != nil {
			tx.Rollback()
			ri.logger.Zap.Error("cannot create risk control data: %s ", err)
			return err
		}
	}

	return nil
}

func (ri RiskIndicatorService) Download(pernr string, format string) ([]byte, string, error) {
	data, err := ri.riskIndicatorRepo.GetAll()
	if err != nil {
		ri.logger.Zap.Error("Errored when try to query indicator: %s", err)
		return nil, "", err
	}

	if len(data) == 0 {
		return nil, "", nil
	}

	switch format {
	case "csv":
		return exportCSV(data)
	case "xlsx":
		return exportExcel(data)
	case "pdf":
		return exportPDF(data)
	default:
		return nil, "", fmt.Errorf("unsupported format export file")
	}
}

func exportPDF(data []models.RiskIndicatorResponse) ([]byte, string, error) {
	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.SetAutoPageBreak(false, 10)
	pdf.SetMargins(10, 10, 10) // margin kiri, atas, kanan
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 14)
	pdf.CellFormat(0, 10, "Risk Control Report", "", 1, "C", false, 0, "")

	headers := []string{
		"Data Source Anomaly",
		"Risk Indicator Code",
		"Risk Indicator",
		"Aktivitas",
		"Produk",
		"Deskripsi",
		"satuan",
		"Periode",
		"Batasan",
		"Business Cycle Aktivitas",
		"SLA Verifikasi",
		"SLA Tindak Lanjut",
		"KPI",
		"Tipe Data Anomali",
		"Kondisi",
		"Created Time",
		"Updated Time",
		"Mapping Set Header",
		"Mapping Set Key Uker",
	}

	colWidths := []float64{20, 20, 25, 15, 15, 25, 15, 20, 20, 20, 20, 25, 20, 20, 20, 20, 20, 20, 20}
	printHeader := func() {
		pdf.SetFillColor(200, 200, 200) // abu-abu header
		pdf.SetFont("Arial", "B", 10)
		for i, h := range headers {
			pdf.CellFormat(colWidths[i], 8, h, "1", 0, "C", true, 0, "")
		}
		pdf.Ln(-1)
		pdf.SetFont("Arial", "", 9)
	}

	printHeader()

	_, pageHeight := pdf.GetPageSize()
	marginBottom := 15.0

	getRowHeight := func(row []string) float64 {
		maxHeight := 0.0
		lineHeight := 5.0
		for i, txt := range row {
			lines := pdf.SplitLines([]byte(txt), colWidths[i])
			h := float64(len(lines)) * lineHeight
			if h > maxHeight {
				maxHeight = h
			}
		}
		return maxHeight
	}

	for _, d := range data {
		createTime := lib.FormatDatePtr(d.CreatedAt)
		updateTime := lib.FormatDatePtr(d.UpdatedAt)

		row := []string{
			d.DataSourceAnomaly,                  // data_source_anomaly
			d.RiskIndicatorCode,                  // risk_indicator_code
			d.RiskIndicator,                      // risk_indicator
			strconv.Itoa(int(d.ActivityID)),      // aktivitas
			strconv.Itoa(int(d.ProductID)),       // produk
			d.Deskripsi,                          // deskripsi
			d.Satuan,                             // satuan
			d.PeriodePemantauan,                  // periode_data
			d.Batasan,                            // Batasan
			d.BusinessCycleActivity,              // Business Cycle Aktivitas
			strconv.Itoa(int(d.SLAVerifikasi)),   // sla_verifikasi
			strconv.Itoa(int(d.SLATindakLanjut)), // sla_tindak_lanjut
			d.KPI,                                // kpi
			d.Type,                               // Tipe Data Anomali
			ReverseCondition(d.Kondisi),          // Kondisi
			createTime,
			updateTime,
			"",
			"",
		}

		rowHeight := getRowHeight(row)
		xStart := pdf.GetX()
		yStart := pdf.GetY()

		// Check page break
		if yStart+rowHeight+marginBottom > pageHeight {
			pdf.AddPage()
			printHeader()
			xStart = pdf.GetX()
			yStart = pdf.GetY()
		}

		// Print each cell with MultiCell and border
		for i, txt := range row {
			x := pdf.GetX()
			y := pdf.GetY()

			pdf.Rect(x, y, colWidths[i], rowHeight, "D")
			pdf.MultiCell(colWidths[i], 5, txt, "", "L", false)
			pdf.SetXY(x+colWidths[i], yStart)
		}
		pdf.SetXY(xStart, yStart+rowHeight)
	}

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, "", fmt.Errorf("failed to generate PDF: %w", err)
	}

	fileName := fmt.Sprintf("risk_indicator_%s.pdf", time.Now().Format("20060102_150405"))

	return buf.Bytes(), fileName, nil
}

func exportExcel(data []models.RiskIndicatorResponse) ([]byte, string, error) {
	f := excelize.NewFile()
	sheet := "risk-indicator"

	f.SetSheetName("Sheet1", sheet)
	headers := []string{
		"Data Source Anomaly",
		"Risk Indicator Code",
		"Risk Indicator",
		"Aktivitas",
		"Produk",
		"Deskripsi",
		"satuan",
		"Periode",
		"Batasan",
		"Business Cycle Aktivitas",
		"SLA Verifikasi",
		"SLA Tindak Lanjut",
		"KPI",
		"Tipe Data Anomali",
		"Kondisi",
		"Created Time",
		"Updated Time",
		"Mapping Set Header",
		"Mapping Set Key Uker",
	}

	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		if err := f.SetCellValue(sheet, cell, h); err != nil {
			return nil, "", fmt.Errorf("failed to set cell: %w", err)
		}
	}

	for idx, v := range data {
		row := idx + 2
		createTime := lib.FormatDatePtr(v.CreatedAt)
		updateTime := lib.FormatDatePtr(v.UpdatedAt)

		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), v.DataSourceAnomaly)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), v.RiskIndicatorCode)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", row), v.RiskIndicator)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", row), strconv.FormatInt(v.ActivityID, 10))
		f.SetCellValue(sheet, fmt.Sprintf("E%d", row), strconv.FormatInt(v.ProductID, 10))
		f.SetCellValue(sheet, fmt.Sprintf("F%d", row), v.Deskripsi)
		f.SetCellValue(sheet, fmt.Sprintf("G%d", row), v.Satuan)
		f.SetCellValue(sheet, fmt.Sprintf("H%d", row), v.PeriodePemantauan)
		f.SetCellValue(sheet, fmt.Sprintf("I%d", row), v.Batasan)
		f.SetCellValue(sheet, fmt.Sprintf("J%d", row), v.BusinessCycleActivity)
		f.SetCellValue(sheet, fmt.Sprintf("K%d", row), strconv.FormatInt(v.SLAVerifikasi, 10))
		f.SetCellValue(sheet, fmt.Sprintf("L%d", row), strconv.FormatInt(v.SLATindakLanjut, 10))
		f.SetCellValue(sheet, fmt.Sprintf("M%d", row), v.KPI)
		f.SetCellValue(sheet, fmt.Sprintf("N%d", row), v.Type)
		f.SetCellValue(sheet, fmt.Sprintf("O%d", row), v.Kondisi)
		f.SetCellValue(sheet, fmt.Sprintf("P%d", row), createTime)
		f.SetCellValue(sheet, fmt.Sprintf("Q%d", row), updateTime)
		f.SetCellValue(sheet, fmt.Sprintf("R%d", row), "")
		f.SetCellValue(sheet, fmt.Sprintf("S%d", row), "")
	}

	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		return nil, "", fmt.Errorf("failed to write excel file: %w", err)
	}

	fileName := fmt.Sprintf("risk_indicator_%s.xlsx", time.Now().Format("20060102_150405"))

	return buf.Bytes(), fileName, nil
}

func exportCSV(data []models.RiskIndicatorResponse) ([]byte, string, error) {
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	headers := []string{
		"Data Source Anomaly",
		"Risk Indicator Code",
		"Risk Indicator",
		"Aktivitas",
		"Produk",
		"Deskripsi",
		"satuan",
		"Periode",
		"Batasan",
		"Business Cycle Aktivitas",
		"SLA Verifikasi",
		"SLA Tindak Lanjut",
		"KPI",
		"Tipe Data Anomali",
		"Kondisi",
		"Created Time",
		"Updated Time",
		"Mapping Set Header",
		"Mapping Set Key Uker",
	}
	if err := writer.Write(headers); err != nil {
		return nil, "", fmt.Errorf("failed to write csv header: %w", err)
	}

	for _, d := range data {
		createTime := lib.FormatDatePtr(d.CreatedAt)
		updateTime := lib.FormatDatePtr(d.UpdatedAt)

		row := []string{
			d.DataSourceAnomaly,                  // data_source_anomaly
			d.RiskIndicatorCode,                  // risk_indicator_code
			d.RiskIndicator,                      // risk_indicator
			strconv.Itoa(int(d.ActivityID)),      // aktivitas
			strconv.Itoa(int(d.ProductID)),       // produk
			d.Deskripsi,                          // deskripsi
			d.Satuan,                             // satuan
			d.PeriodePemantauan,                  // periode_data
			d.Batasan,                            // Batasan
			d.BusinessCycleActivity,              // Business Cycle Aktivitas
			strconv.Itoa(int(d.SLAVerifikasi)),   // sla_verifikasi
			strconv.Itoa(int(d.SLATindakLanjut)), // sla_tindak_lanjut
			d.KPI,                                // kpi
			d.Type,                               // Tipe Data Anomali
			ReverseCondition(d.Kondisi),          // Kondisi
			createTime,
			updateTime,
			"",
			"",
		}

		if err := writer.Write(row); err != nil {
			return nil, "", fmt.Errorf("failed to write csv row: %w", err)
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return nil, "", fmt.Errorf("failed to flush csv data: %w", err)
	}

	fileName := fmt.Sprintf("risk_indicator_%s.csv", time.Now().Format("20060102_150405"))
	return buf.Bytes(), fileName, nil
}

func ConditionMap(text string) string {
	normalized := strings.ToLower(text)
	normalized = strings.ReplaceAll(normalized, "&", "dan")
	normalized = strings.ReplaceAll(normalized, "  ", " ")
	normalized = strings.TrimSpace(normalized)

	label := map[string]string{
		"tampilkan semua dan verifikasi semua":     "tampilkanSemuaDanVerifikasiSemua",
		"tampilkan semua dan verifikasi terbaru":   "tampilkanSemuaDanVerifikasiTerbaru",
		"tampilkan terbaru dan verifikasi terbaru": "tampilkanTerbaruDanVerifikasiTerbaru",
	}

	condition, ok := label[normalized]
	if !ok {
		return ""
	}

	return condition
}

func ReverseCondition(text string) string {
	label := map[string]string{
		"tampilkanSemuaDanVerifikasiSemua":     "Tampilkan Semua dan Verifikasi Semua",
		"tampilkanSemuaDanVerifikasiTerbaru":   "Tampilkan Semua dan Verifikasi Terbaru",
		"tampilkanTerbaruDanVerifikasiTerbaru": "Tampilkan Terbaru dan Verifikasi Terbaru",
	}

	condition, ok := label[text]
	if !ok {
		return ""
	}

	return condition
}
