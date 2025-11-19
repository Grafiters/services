package verifikasirealisasi

import (
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/verifikasirealisasi"
	"strings"
	"time"

	"gitlab.com/golang-package-library/logger"

	"gorm.io/gorm"
)

type VerifikasiRealisasiDefinition interface {
	// Verifikasi Realisasi
	GetData(request models.VerifikasiRealisasiFilterRequest) (response []models.VerifikasiRealisasiList, totalData int64, err error)
	GetDetailVerifikasi(id int64) (response models.VerifikasiRealisasiDetailResponse, err error)
	StoreVerifikasi(request *models.VerifikasiRealisasi, tx *gorm.DB) (response *models.VerifikasiRealisasi, err error)
	UpdateVerifikasi(request *models.VerifikasiRealisasiUpdate, include []string, tx *gorm.DB) (response bool, err error)
	DeleteVerifikasi(request *models.VerifikasiRealisasiUpdateDelete, include []string, tx *gorm.DB) (response bool, err error)

	// SampleDataRealisasi
	GetDataRealisasiById(id int64) (response []models.SampleDataRealisasiResponse, err error)
	SaveDataRealisasi(request *models.SampleDataRealisasi, tx *gorm.DB) (response bool, err error)

	// CriteriaVerifikasi
	GetKriteriaById(id int64) (response []models.RealisasiKreditKriteriaResponse, err error)
	SaveDataCriteria(request *models.RealisasiKreditKriteria, tx *gorm.DB) (response bool, err error)

	// LampiranVerifikasi
	GetFileById(id int64) (response []models.VerifikasiRealisasiFilesResponse, err error)
	SaveDataFile(request *models.VerifikasiRealisasiFilesRequest, tx *gorm.DB) (response bool, err error)

	// No Pelaporan
	GetNoPelaporan(request *models.NoPalaporanVerifikasiRealisasiRequest) (responses string, err error)
}

type VerifikasiRealisasiRepository struct {
	db      lib.Database
	logger  logger.Logger
	timeout time.Duration
}

func NewVerfikasiRealisasiRepository(
	db lib.Database,
	logger logger.Logger,
) VerifikasiRealisasiDefinition {
	return VerifikasiRealisasiRepository{
		db:      db,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// GetData implements VerifikasiRealisasiDefinition.
func (v VerifikasiRealisasiRepository) GetData(request models.VerifikasiRealisasiFilterRequest) (response []models.VerifikasiRealisasiList, totalData int64, err error) {

	query := v.db.DB.Table("verifikasi_realisasi_kredit vrk").
		Select(`
			ROW_NUMBER() OVER (ORDER BY vrk.id DESC) AS 'no',
			vrk.id,
			vrk.no_pelaporan,
			concat(vrk.BRANCH," - ", vrk.BRDESC) 'unit_kerja',
			vrk.activity_name 'aktifitas',
			CASE
				WHEN vrk.indikasi_fraud = '1' THEN 'YA'
				ELSE 'Tidak'
			END 'indikasi_fraud',
			vrk.action 'status_verif',
			vrk.action_validasi 'status_fraud',
			vdr.status_verifikasi 'sudah_verifikasi'
		`).
		Where("vrk.deleted = 0").
		Joins("LEFT JOIN verifikasi_data_realisasi vdr ON vdr.verifikasi_id = vrk.id").
		Order(`vrk.id DESC`).
		Group(`vrk.id`)

	// konsisi Filter Tobe add

	// if request.Branches != "" {
	// 	branches := strings.Split(request.Branches, ",")
	// 	query = query.Where(`v.BRANCH in (?)`, branches)
	// }

	if request.REGION != "all" && request.REGION != "" {
		query = query.Where("vrk.REGION = ?", request.REGION)
	}

	if request.MAINBR != "all" && request.MAINBR != "" {
		mainbrs := strings.Split(request.MAINBR, ",")
		query = query.Where("vrk.MAINBR in (?)", mainbrs)
	}

	if request.BRANCH != "all" && request.BRANCH != "" {
		branches := strings.Split(request.BRANCH, ",")
		if len(branches) > 1 {
			fmt.Println("...")
			query = query.Where("vrk.BRANCH in (?)", branches)
		} else {
			fmt.Println(",,,")
			query = query.Where("vrk.BRANCH = ?", request.BRANCH)
		}
	}

	if request.Kostl != "" {
		query = query.Where(`vrk.created_id = ?`, request.Pernr)
	}

	if request.UnitKerja != "" {
		query = query.Where("vrk.unit_kerja = ?", request.UnitKerja)
	}

	if request.CriteriaID != "" {
		query = query.Where("JSON_CONTAINS(vrk.kriteria_data, ?)", request.CriteriaID)
	}

	if request.Segment != "" && request.Segment != "all" {
		query = query.Where(`JSON_UNQUOTE(JSON_EXTRACT(LOWER(vdr.data_realisasi), '$.segment')) = ?`, strings.ToLower(request.Segment))
	}

	if request.ProductID != "" && request.ProductID != "all" {
		query = query.Where("vrk.product_id = ?", request.ProductID)
	}

	if request.SudahVerifikasi != "all" && request.SudahVerifikasi != "" {
		query = query.Where("vdr.status_verifikasi = ?", request.SudahVerifikasi)
	}

	if request.Efektif != "all" && request.Efektif != "" {
		query = query.Where("vrk.butuh_perbaikan = ?", request.Efektif)
	}

	if request.IndikasiFraud != "all" && request.IndikasiFraud != "" {
		query = query.Where("vrk.indikasi_fraud = ?", request.IndikasiFraud)
	}

	if err = query.Count(&totalData).Error; err != nil {
		v.logger.Zap.Error("Error counting records:", err)
		return
	}

	if request.Limit != 0 {
		query = query.Limit(request.Limit)
	}

	if request.Offset != 0 {
		query = query.Offset(request.Offset)
	}

	err = query.Scan(&response).Error

	if err != nil {
		v.logger.Zap.Error(err)
		return nil, totalData, err
	}

	return response, totalData, err
}

// GetDetailVerifikasi implements VerifikasiRealisasiDefinition.
func (v VerifikasiRealisasiRepository) GetDetailVerifikasi(id int64) (response models.VerifikasiRealisasiDetailResponse, err error) {
	db := v.db.DB.Table(`verifikasi_realisasi_kredit`).Where(`id = ?`, id)

	err = db.Scan(&response).Error

	if err != nil {
		v.logger.Zap.Error(err.Error())
		return response, err
	}

	return response, err
}

// StoreVerifikasi implements VerifikasiRealisasiDefinition.
func (v VerifikasiRealisasiRepository) StoreVerifikasi(request *models.VerifikasiRealisasi, tx *gorm.DB) (response *models.VerifikasiRealisasi, err error) {
	return request, tx.Save(&request).Error
}

// UpdateVerifikasi implements VerifikasiRealisasiDefinition.
func (v VerifikasiRealisasiRepository) UpdateVerifikasi(request *models.VerifikasiRealisasiUpdate, include []string, tx *gorm.DB) (response bool, err error) {
	return true, tx.Table("verifikasi_realisasi_kredit").Save(&request).Error
}

// DeleteVerifikasi implements VerifikasiRealisasiDefinition.
func (v VerifikasiRealisasiRepository) DeleteVerifikasi(request *models.VerifikasiRealisasiUpdateDelete, include []string, tx *gorm.DB) (response bool, err error) {
	err = tx.Table(`verifikasi_realisasi_kredit`).Save(&request).Error
	if err != nil {
		return false, err
	}
	return true, err
}

// GetDataRealisasiById implements VerifikasiRealisasiDefinition.
func (v VerifikasiRealisasiRepository) GetDataRealisasiById(id int64) (response []models.SampleDataRealisasiResponse, err error) {
	db := v.db.DB.Model(&response).Where(`verifikasi_id = ?`, id)

	err = db.Scan(&response).Error

	if err != nil {
		v.logger.Zap.Error(err.Error())
		return nil, err
	}

	return response, err
}

// StoreDataRealisasi implements VerifikasiRealisasiDefinition.
func (v VerifikasiRealisasiRepository) SaveDataRealisasi(request *models.SampleDataRealisasi, tx *gorm.DB) (response bool, err error) {
	return true, tx.Save(&request).Error
}

// GetKriteriaById implements VerifikasiRealisasiDefinition.
func (v VerifikasiRealisasiRepository) GetKriteriaById(id int64) (response []models.RealisasiKreditKriteriaResponse, err error) {
	db := v.db.DB.Model(&response).Where(`verifikasi_id = ?`, id)

	err = db.Scan(&response).Error

	if err != nil {
		v.logger.Zap.Error(err.Error())
		return nil, err
	}

	return response, err
}

// StoreDataCriteria implements VerifikasiRealisasiDefinition.
func (v VerifikasiRealisasiRepository) SaveDataCriteria(request *models.RealisasiKreditKriteria, tx *gorm.DB) (response bool, err error) {
	return true, tx.Save(&request).Error
}

// GetFileById implements VerifikasiRealisasiDefinition.
func (v VerifikasiRealisasiRepository) GetFileById(id int64) (response []models.VerifikasiRealisasiFilesResponse, err error) {
	// panic("unimplemented")
	db := v.db.DB.Table("verifikasi_realisasi_lampiran vrl").
		Select(`
				vrl.id,
				vrl.verifikasi_id,
				vrl.files_id,
				f.filename,
				f.path`).
		Joins("JOIN files f ON vrl.files_id = f.id").
		Where(`verifikasi_id = ?`, id)

	err = db.Scan(&response).Error

	if err != nil {
		v.logger.Zap.Error(err.Error())
		return nil, err
	}
	return response, err
}

// StoreDataFile implements VerifikasiRealisasiDefinition.
func (v VerifikasiRealisasiRepository) SaveDataFile(request *models.VerifikasiRealisasiFilesRequest, tx *gorm.DB) (response bool, err error) {
	return true, tx.Save(&request).Error
}

// GetNoPelaporan implements VerifikasiRealisasiDefinition.
func (v VerifikasiRealisasiRepository) GetNoPelaporan(request *models.NoPalaporanVerifikasiRealisasiRequest) (responses string, err error) {
	kode := "VRK-"
	today := lib.GetTimeNow("date2")

	if request.ORGEH != "" {
		kode += request.ORGEH + "-" + today
	}

	query := v.db.DB.Table("verifikasi_realisasi_kredit").
		Select(`RIGHT(CONCAT("0000",(count(*) + 1)), 4) 'no_pelaporan'`).
		Where(`no_pelaporan like ?`, fmt.Sprintf("%s%%", kode))

	err = query.Scan(&responses).Error

	no_pelaporan := kode + "-" + responses

	if err != nil {
		v.logger.Zap.Error(err)
		return responses, err
	}

	return no_pelaporan, err
}
