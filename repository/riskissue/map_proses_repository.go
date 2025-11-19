package riskissue

import (
	"riskmanagement/lib"
	models "riskmanagement/models/riskissue"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type MapProsesDefinition interface {
	GetAll() (responses []models.MapProsesResponse, err error)
	GetOne(id int64) (responses []models.MapProsesResponse, err error)
	Store(request *models.MapProses, tx *gorm.DB) (responses *models.MapProses, err error)
	Update(request *models.MapProses, tx *gorm.DB) (responses bool, err error)
	Delete(id int64) (err error)
	GetOneDataByID(id int64) (responses []models.MapProsesResponseFinal, err error)
	DeleteDataByID(id int64, tx *gorm.DB) (err error)
	WithTrx(trxHandle *gorm.DB) MapProsesRepository
}

type MapProsesRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewMapProsesRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) MapProsesDefinition {
	return MapProsesRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// Delete implements MapProsesDefinition
func (mp MapProsesRepository) Delete(id int64) (err error) {
	return mp.db.DB.Where("id = ?", id).Delete(&models.MapProsesResponseFinal{}).Error
}

// DeleteDataByID implements MapProsesDefinition
func (mp MapProsesRepository) DeleteDataByID(id int64, tx *gorm.DB) (err error) {
	return tx.Where("id_risk_issue = ?", id).Delete(&models.MapProsesResponseFinal{}).Error
}

// GetAll implements MapProsesDefinition
func (mp MapProsesRepository) GetAll() (responses []models.MapProsesResponse, err error) {
	return responses, mp.db.DB.Find(&responses).Error
}

// GetOne implements MapProsesDefinition
func (mp MapProsesRepository) GetOne(id int64) (responses []models.MapProsesResponse, err error) {
	return responses, mp.db.DB.Where("id = ?", id).Find(&responses).Error
}

// GetOneDataByID implements MapProsesDefinition
func (mp MapProsesRepository) GetOneDataByID(id int64) (responses []models.MapProsesResponseFinal, err error) {
	rows, err := mp.db.DB.Raw(`
	SELECT
		mp.id 'id',
		mp.id_risk_issue 'id_risk_issue',
		mp.mega_proses 'mega_proses',
		mega.mega_proses 'mega_proses_desc',
		mp.major_proses 'major_proses',
		major.major_proses 'major_proses_desc',
		mp.sub_major_proses 'sub_major_proses',
		submajor.sub_major_proses 'sub_major_proses_desc'
	FROM risk_issue_map_proses mp
	LEFT JOIN mega_proses mega ON mega.kode_mega_proses = mp.mega_proses
	LEFT JOIN major_proses major ON major.kode_major_proses = mp.major_proses
	LEFT JOIN sub_major_proses submajor ON submajor.kode_sub_major_proses = mp.sub_major_proses
	WHERE id_risk_issue = ?`, id).Rows()

	if err != nil {
		return responses, err
	}

	defer rows.Close()
	var DataMp models.MapProsesResponseFinal

	for rows.Next() {
		mp.db.DB.ScanRows(rows, &DataMp)
		responses = append(responses, DataMp)
	}

	return responses, err
}

// Store implements MapProsesDefinition
func (mp MapProsesRepository) Store(request *models.MapProses, tx *gorm.DB) (responses *models.MapProses, err error) {
	return request, tx.Save(&request).Error
}

// Update implements MapProsesDefinition
func (mp MapProsesRepository) Update(request *models.MapProses, tx *gorm.DB) (responses bool, err error) {
	return true, tx.Save(&request).Error
}

// WithTrx implements MapProsesDefinition
func (mp MapProsesRepository) WithTrx(trxHandle *gorm.DB) MapProsesRepository {
	if trxHandle == nil {
		mp.logger.Zap.Error("transaction Database not found in gin context.")
		return mp
	}

	mp.db.DB = trxHandle
	return mp
}
