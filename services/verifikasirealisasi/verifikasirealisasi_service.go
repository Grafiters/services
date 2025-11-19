package verifikasirealisasi

import (
	"fmt"
	"os"
	"riskmanagement/lib"
	requestFile "riskmanagement/models/files"
	models "riskmanagement/models/verifikasirealisasi"
	fileRepo "riskmanagement/repository/files"
	verifikasiRealisasi "riskmanagement/repository/verifikasirealisasi"
	"strings"

	"github.com/google/uuid"
	"gitlab.com/golang-package-library/logger"
	minio "gitlab.com/golang-package-library/minio"
)

// var (
// 	timeNow = lib.GetTimeNow("timestime")
// 	UUID    = uuid.NewString()
// )

type VerifikasiRealisasiServiceDefinition interface {
	GetData(request models.VerifikasiRealisasiFilterRequest) (response []models.VerifikasiRealisasiList, totalData int64, err error)
	StoreVerifikasi(request *models.VerifikasiRealisasiRequest) (response bool, err error)
	GetOne(id int64) (response models.VerifikasiRealisasiResponse, status bool, err error)
	Delete(request *models.VerifikasiRealisasiRequest) (response bool, sampledata []models.SampleDataRealisasiResponse, err error)
	Update(request *models.VerifikasiRealisasiRequest) (status bool, err error)
	GetNoPelaporan(request *models.NoPalaporanVerifikasiRealisasiRequest) (response string, err error)
}

type VerifikasiRealisasiService struct {
	db                      lib.Database
	minio                   minio.Minio
	logger                  logger.Logger
	verifikasiRealisasiRepo verifikasiRealisasi.VerifikasiRealisasiDefinition
	fileRepo                fileRepo.FilesDefinition
}

func NewVerifikasiRealisasiService(
	db lib.Database,
	minio minio.Minio,
	logger logger.Logger,
	verifikasiRealisasiRepo verifikasiRealisasi.VerifikasiRealisasiDefinition,
	fileRepo fileRepo.FilesDefinition,
) VerifikasiRealisasiServiceDefinition {
	return VerifikasiRealisasiService{
		db:                      db,
		minio:                   minio,
		logger:                  logger,
		verifikasiRealisasiRepo: verifikasiRealisasiRepo,
		fileRepo:                fileRepo,
	}
}

// GetData implements VerifikasiRealisasiServiceDefinition.
func (v VerifikasiRealisasiService) GetData(request models.VerifikasiRealisasiFilterRequest) (response []models.VerifikasiRealisasiList, totalData int64, err error) {
	data_verifikasi, totalData, err := v.verifikasiRealisasiRepo.GetData(request)

	if err != nil {
		v.logger.Zap.Error(err.Error())
		return nil, totalData, err
	}

	for _, value := range data_verifikasi {
		response = append(response, models.VerifikasiRealisasiList{
			ID:            value.ID,
			No:            value.No,
			NoPelaporan:   value.NoPelaporan,
			UnitKerja:     value.UnitKerja,
			Aktifitas:     value.Aktifitas,
			IndikasiFraud: value.IndikasiFraud,
			StatusVerif:   value.StatusVerif,
			StatusFraud:   value.StatusFraud,
		})
	}

	return response, totalData, err
}

// StoreVerifikasi implements VerifikasiRealisasiServiceDefinition.
func (v VerifikasiRealisasiService) StoreVerifikasi(request *models.VerifikasiRealisasiRequest) (response bool, err error) {
	// panic("unimplemented")
	today := lib.GetTimeNow("timestime")
	UUID := uuid.New()
	tx := v.db.DB.Begin()

	reqVerifRealisasiKredit := &models.VerifikasiRealisasi{
		ID:               request.ID,
		NoPelaporan:      request.NoPelaporan,
		SumberData:       request.SumberData,
		REGION:           request.REGION,
		RGDESC:           request.RGDESC,
		MAINBR:           request.MAINBR,
		MBDESC:           request.MBDESC,
		BRANCH:           request.BRANCH,
		BRDESC:           request.BRDESC,
		ActivityID:       request.ActivityID,
		ActivityName:     request.ActivityName,
		ProductID:        request.ProductID,
		ProductName:      request.ProductName,
		SubActivityID:    request.SubActivityID,
		SubActivityName:  request.SubActivityName,
		RestruckFlag:     request.RestruckFlag,
		PeriodeData:      request.PeriodeData,
		KunjunganNasabah: request.KunjunganNasabah,
		TglKunjungan:     request.TglKunjungan,
		ButuhPerbaikan:   request.ButuhPerbaikan,
		IndikasiFraud:    request.IndikasiFraud,
		HasilVerifikasi:  request.HasilVerifikasi,
		Status:           request.Status,
		Action:           request.Action,
		Deleted:          false,
		StatusValidasi:   request.StatusValidasi,
		ActionValidasi:   request.ActionValidasi,
		CreatedAt:        today,
		CreatedID:        request.CreatedID,
		CreatedDesc:      request.CreatedDesc,
		UpdatedAt:        today,
		UpdatedBy:        request.CreatedID,
		UpdatedDesc:      request.CreatedDesc,
		KriteriaData:     request.KriteriaData,
		ListCriteria:     request.ListCriteria,
	}

	dataVerif, err := v.verifikasiRealisasiRepo.StoreVerifikasi(reqVerifRealisasiKredit, tx)

	// fmt.Println(dataVerif)

	if err != nil {
		tx.Rollback()
		v.logger.Zap.Error(err)
		return false, err
	}

	// Input Sample Data
	// fmt.Println("raw data =>", request.SampleData)
	statusVerif := false
	if request.Action == "Draft" {
		statusVerif = false
	} else if request.Action == "Selesai" {
		statusVerif = true
	}

	if len(request.SampleData) != 0 {
		for _, val := range request.SampleData {

			fmt.Println("value :", val)
			_, err = v.verifikasiRealisasiRepo.SaveDataRealisasi(&models.SampleDataRealisasi{
				// ID:               0,
				VerifikasiID:     dataVerif.ID,
				DataRealisasi:    val.DataRealisasi,
				StatusVerifikasi: statusVerif,
			}, tx)

			if err != nil {
				tx.Rollback()
				v.logger.Zap.Error(err)
				return false, err
			}
		}
	} else {
		tx.Rollback()
		v.logger.Zap.Error(err)
		return false, err
	}

	//Begin Input Lampiran
	bucket := os.Getenv("BUCKET_NAME")

	if len(request.Files) != 0 {
		for _, value := range request.Files {
			var destinationPath string
			if value.Filename != "" {
				bucketExist := v.minio.BucketExist(v.minio.Client(), bucket)

				pathSplit := strings.Split(value.Path, "/")
				sourcePath := fmt.Sprint(value.Path)
				destinationPath = pathSplit[1] + "/" +
					lib.GetTimeNow("year") + "/" +
					lib.GetTimeNow("month") + "/" +
					lib.GetTimeNow("day") + "/" +
					pathSplit[2] + "/" + UUID.String() + "/" + value.Filename

				if bucketExist {
					fmt.Println("Exist")
					fmt.Println(bucket)
					fmt.Println(sourcePath)
					fmt.Println(destinationPath)
					uploaded := v.minio.CopyObject(v.minio.Client(), bucket, sourcePath, bucket, destinationPath)

					fmt.Println(uploaded)
				} else {
					fmt.Println("Not Exist")
					fmt.Println(bucket)
					fmt.Println(sourcePath)
					fmt.Println(destinationPath)
					v.minio.MakeBucket(v.minio.Client(), bucket, "")
					uploaded := v.minio.CopyObject(v.minio.Client(), bucket, sourcePath, bucket, destinationPath)

					fmt.Println(uploaded)
				}
			}

			files, err := v.fileRepo.Store(&requestFile.Files{
				Filename:  value.Filename,
				Path:      destinationPath,
				Extension: value.Extension,
				Size:      value.Size,
				CreatedAt: &today,
			}, tx)

			if err != nil {
				tx.Rollback()
				v.logger.Zap.Error(err)
				return false, err
			}

			_, err = v.verifikasiRealisasiRepo.SaveDataFile(&models.VerifikasiRealisasiFilesRequest{
				VerifikasiID: dataVerif.ID,
				FilesID:      files.ID,
				// CreatedAt:    &timeNow,
			}, tx)

			if err != nil {
				tx.Rollback()
				v.logger.Zap.Error(err)
				return false, err
			}
		}
	} else {
		tx.Rollback()
		v.logger.Zap.Error(err)
		return false, err
	}

	tx.Commit()
	return true, err
}

// GetOne implements VerifikasiRealisasiServiceDefinition.
func (v VerifikasiRealisasiService) GetOne(id int64) (response models.VerifikasiRealisasiResponse, status bool, err error) {
	dataVerif, err := v.verifikasiRealisasiRepo.GetDetailVerifikasi(id)
	fmt.Println(dataVerif)

	if dataVerif.ID != 0 {
		files, err := v.verifikasiRealisasiRepo.GetFileById(dataVerif.ID)
		if err != nil {
			v.logger.Zap.Error(err)
			return response, false, err
		}

		sampledata, err := v.verifikasiRealisasiRepo.GetDataRealisasiById(dataVerif.ID)
		if err != nil {
			v.logger.Zap.Error(err)
			return response, false, err
		}

		// kriteria, err := v.verifikasiRealisasiRepo.GetKriteriaById(dataVerif.ID)
		// if err != nil {
		// 	v.logger.Zap.Error(err)
		// 	return response, false, err
		// }

		response = models.VerifikasiRealisasiResponse{
			ID:               dataVerif.ID,
			NoPelaporan:      dataVerif.NoPelaporan,
			SumberData:       dataVerif.SumberData,
			REGION:           dataVerif.REGION,
			RGDESC:           dataVerif.RGDESC,
			MAINBR:           dataVerif.MAINBR,
			MBDESC:           dataVerif.MBDESC,
			BRANCH:           dataVerif.BRANCH,
			BRDESC:           dataVerif.BRDESC,
			ActivityID:       dataVerif.ActivityID,
			ActivityName:     dataVerif.ActivityName,
			ProductID:        dataVerif.ProductID,
			ProductName:      dataVerif.ProductName,
			SubActivityID:    dataVerif.SubActivityID,
			SubActivityName:  dataVerif.SubActivityName,
			RestruckFlag:     dataVerif.RestruckFlag,
			PeriodeData:      dataVerif.PeriodeData,
			KunjunganNasabah: dataVerif.KunjunganNasabah,
			TglKunjungan:     dataVerif.TglKunjungan,
			ButuhPerbaikan:   dataVerif.ButuhPerbaikan,
			IndikasiFraud:    dataVerif.IndikasiFraud,
			HasilVerifikasi:  dataVerif.HasilVerifikasi,
			Status:           dataVerif.Status,
			Action:           dataVerif.Action,
			Deleted:          dataVerif.Deleted,
			StatusValidasi:   dataVerif.StatusValidasi,
			ActionValidasi:   dataVerif.ActionValidasi,
			CreatedAt:        dataVerif.CreatedAt,
			CreatedID:        dataVerif.CreatedID,
			CreatedDesc:      dataVerif.CreatedDesc,
			UpdatedAt:        dataVerif.UpdatedAt,
			UpdatedBy:        dataVerif.UpdatedBy,
			UpdatedDesc:      dataVerif.UpdatedDesc,
			SampleData:       sampledata,
			KriteriaData:     dataVerif.KriteriaData,
			ListCriteria:     dataVerif.ListCriteria,
			Files:            files,
		}
	}

	if err != nil {
		v.logger.Zap.Error(err)
		return response, false, err
	}

	return response, true, err
}

// Delete implements VerifikasiRealisasiServiceDefinition.
func (v VerifikasiRealisasiService) Delete(request *models.VerifikasiRealisasiRequest) (response bool, sampledata []models.SampleDataRealisasiResponse, err error) {
	today := lib.GetTimeNow("timestime")

	tx := v.db.DB.Begin()

	UpdateDataVerifikasiRealisasiKredit := &models.VerifikasiRealisasiUpdateDelete{
		ID:          request.ID,
		UpdatedBy:   request.UpdatedBy,
		UpdatedDesc: request.UpdatedDesc,
		Deleted:     true,
		UpdatedAt:   &today,
	}

	include := []string{
		"id",
		"updated_by",
		"updated_desc",
		"deleted",
		"updated_at",
	}

	_, err = v.verifikasiRealisasiRepo.DeleteVerifikasi(UpdateDataVerifikasiRealisasiKredit, include, tx)
	if err != nil {
		tx.Rollback()
		v.logger.Zap.Error(err)
		return false, sampledata, err
	}

	sampledata, err = v.verifikasiRealisasiRepo.GetDataRealisasiById(request.ID)
	if err != nil {
		v.logger.Zap.Error(err)
		return false, sampledata, err
	}

	tx.Commit()
	return true, sampledata, err
}

// Update implements VerifikasiRealisasiServiceDefinition.
func (v VerifikasiRealisasiService) Update(request *models.VerifikasiRealisasiRequest) (status bool, err error) {
	today := lib.GetTimeNow("timestime")
	UUID := uuid.New()
	tx := v.db.DB.Begin()

	// fmt.Println(".,.,.,./,/,")

	updateVerifikasiRealisasiKredit := &models.VerifikasiRealisasiUpdate{
		ID:               request.ID,
		NoPelaporan:      request.NoPelaporan,
		SumberData:       request.SumberData,
		REGION:           request.REGION,
		RGDESC:           request.RGDESC,
		MAINBR:           request.MAINBR,
		MBDESC:           request.MBDESC,
		BRANCH:           request.BRANCH,
		BRDESC:           request.BRDESC,
		ActivityID:       request.ActivityID,
		ActivityName:     request.ActivityName,
		ProductID:        request.ProductID,
		ProductName:      request.ProductName,
		SubActivityID:    request.SubActivityID,
		SubActivityName:  request.SubActivityName,
		RestruckFlag:     request.RestruckFlag,
		PeriodeData:      request.PeriodeData,
		KunjunganNasabah: request.KunjunganNasabah,
		TglKunjungan:     request.TglKunjungan,
		ButuhPerbaikan:   request.ButuhPerbaikan,
		IndikasiFraud:    request.IndikasiFraud,
		HasilVerifikasi:  request.HasilVerifikasi,
		Status:           request.Status,
		Action:           request.Action,
		Deleted:          request.Deleted,
		StatusValidasi:   request.StatusValidasi,
		ActionValidasi:   request.ActionValidasi,
		UpdatedAt:        today,
		UpdatedBy:        request.UpdatedBy,
		UpdatedDesc:      request.UpdatedDesc,
		KriteriaData:     request.KriteriaData,
	}

	include := []string{
		"id",
		"no_pelaporan",
		"sumber_data",
		"REGION",
		"RGDESC",
		"MAINBR",
		"MBDESC",
		"BRANCH",
		"BRDESC",
		"activity_id",
		"activity_name",
		"product_id",
		"product_name",
		"sub_activity_id",
		"restruck_flag",
		"periode_data",
		"kunjungan_nasabah",
		"tgl_kunjungan",
		"butuh_perbaikan",
		"indikasi_fraud",
		"hasil_verifikasi",
		"status",
		"action",
		"deleted",
		"status_validasi",
		"action_validasi",
		"updated_at",
		"updated_by",
		"updated_desc",
		"kriteria_data",
	}

	_, err = v.verifikasiRealisasiRepo.UpdateVerifikasi(updateVerifikasiRealisasiKredit, include, tx)

	if err != nil {
		v.logger.Zap.Error(err)
		tx.Rollback()
		return false, err
	}

	statusVerif := false
	if request.Action == "Draft" {
		statusVerif = false
	} else if request.Action == "Selesai" {
		statusVerif = true
	}

	if len(request.SampleData) != 0 {
		for _, val := range request.SampleData {
			_, err = v.verifikasiRealisasiRepo.SaveDataRealisasi(&models.SampleDataRealisasi{
				ID:               val.ID,
				VerifikasiID:     val.VerifikasiID,
				DataRealisasi:    val.DataRealisasi,
				StatusVerifikasi: statusVerif,
			}, tx)

			if err != nil {
				tx.Rollback()
				v.logger.Zap.Error(err)
				return false, err
			}
		}
	} else {
		tx.Rollback()
		v.logger.Zap.Error(err)
		return false, err
	}

	// if len(request.KriteriaData) != 0 {
	// 	for _, val := range request.KriteriaData {
	// 		_, err = v.verifikasiRealisasiRepo.SaveDataCriteria(&models.RealisasiKreditKriteria{
	// 			ID:           val.ID,
	// 			VerifikasiID: val.VerifikasiID,
	// 			KriteriaID:   val.KriteriaID,
	// 			Value:        val.Value,
	// 		}, tx)

	// 		if err != nil {
	// 			tx.Rollback()
	// 			v.logger.Zap.Error(err)
	// 			return false, err
	// 		}
	// 	}
	// } else {
	// 	tx.Rollback()
	// 	v.logger.Zap.Error(err)
	// 	return false, err
	// }

	bucket := os.Getenv("BUCKET_NAME")

	if len(request.Files) != 0 {
		for _, value := range request.Files {
			var destinationPath string
			if value.Filename != "" {
				bucketExist := v.minio.BucketExist(v.minio.Client(), bucket)

				pathSplit := strings.Split(value.Path, "/")
				sourcePath := fmt.Sprint(value.Path)
				destinationPath = pathSplit[1] + "/" +
					lib.GetTimeNow("year") + "/" +
					lib.GetTimeNow("month") + "/" +
					lib.GetTimeNow("day") + "/" +
					pathSplit[2] + "/" + UUID.String() + "/" + value.Filename

				if bucketExist {
					fmt.Println("Exist")
					fmt.Println(bucket)
					fmt.Println(sourcePath)
					fmt.Println(destinationPath)
					uploaded := v.minio.CopyObject(v.minio.Client(), bucket, sourcePath, bucket, destinationPath)

					fmt.Println(uploaded)
				} else {
					fmt.Println("Not Exist")
					fmt.Println(bucket)
					fmt.Println(sourcePath)
					fmt.Println(destinationPath)
					v.minio.MakeBucket(v.minio.Client(), bucket, "")
					uploaded := v.minio.CopyObject(v.minio.Client(), bucket, sourcePath, bucket, destinationPath)

					fmt.Println(uploaded)
				}

				files, err := v.fileRepo.Store(&requestFile.Files{
					Filename:  value.Filename,
					Path:      destinationPath,
					Extension: value.Extension,
					Size:      value.Size,
					CreatedAt: &today,
				}, tx)

				if err != nil {
					tx.Rollback()
					v.logger.Zap.Error(err)
					return false, err
				}

				_, err = v.verifikasiRealisasiRepo.SaveDataFile(&models.VerifikasiRealisasiFilesRequest{
					ID:           value.ID,
					VerifikasiID: request.ID,
					FilesID:      files.ID,
				}, tx)

				if err != nil {
					tx.Rollback()
					v.logger.Zap.Error(err)
					return false, err
				}
			}
		}
	} else {
		tx.Rollback()
		v.logger.Zap.Error(err)
		return false, err
	}

	tx.Commit()
	return true, err
}

// GetNoPelaporan implements VerifikasiRealisasiServiceDefinition.
func (v VerifikasiRealisasiService) GetNoPelaporan(request *models.NoPalaporanVerifikasiRealisasiRequest) (response string, err error) {
	no_pelaporan, err := v.verifikasiRealisasiRepo.GetNoPelaporan(request)

	if err != nil {
		v.logger.Zap.Error(err)
		return response, err
	}

	return no_pelaporan, err
}
