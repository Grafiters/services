package riskissue

import (
	"riskmanagement/lib"
	models "riskmanagement/models/riskissue"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type MapLiniBisnisDefinition interface {
	GetAll() (responses []models.MapLiniBisnisResponse, err error)
	GetOne(id int64) (responses []models.MapLiniBisnisResponse, err error)
	Store(request *models.MapLiniBisnis, tx *gorm.DB) (responses *models.MapLiniBisnis, err error)
	Update(request *models.MapLiniBisnis, tx *gorm.DB) (responses bool, err error)
	Delete(id int64) (err error)
	GetOneDataByID(id int64) (responses []models.MapLiniBisnisResponseFinal, err error)
	DeleteDataByID(id int64, tx *gorm.DB) (err error)
	WithTrx(trxHandle *gorm.DB) MapLiniBisnisRepository
}

type MapLiniBisnisRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewMapLiniBisnisRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) MapLiniBisnisDefinition {
	return MapLiniBisnisRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// Delete implements MapLiniBisnisDefinition
func (mp MapLiniBisnisRepository) Delete(id int64) (err error) {
	return mp.db.DB.Where("id = ?", id).Delete(&models.MapLiniBisnisResponseFinal{}).Error
}

// DeleteDataByID implements MapLiniBisnisDefinition
func (mp MapLiniBisnisRepository) DeleteDataByID(id int64, tx *gorm.DB) (err error) {
	return tx.Where("id_risk_issue = ?", id).Delete(&models.MapLiniBisnisResponseFinal{}).Error
}

// GetAll implements MapLiniBisnisDefinition
func (mp MapLiniBisnisRepository) GetAll() (responses []models.MapLiniBisnisResponse, err error) {
	return responses, mp.db.DB.Find(&responses).Error
}

// GetOne implements MapLiniBisnisDefinition
func (mp MapLiniBisnisRepository) GetOne(id int64) (responses []models.MapLiniBisnisResponse, err error) {
	return responses, mp.db.DB.Where("id = ?", id).Find(&responses).Error
}

// GetOneDataByID implements MapLiniBisnisDefinition
func (mp MapLiniBisnisRepository) GetOneDataByID(id int64) (responses []models.MapLiniBisnisResponseFinal, err error) {
	rows, err := mp.db.DB.Raw(`
	SELECT
		ml.id 'id',
		ml.id_risk_issue 'id_risk_issue',
		ml.lini_bisnis_lv1 'lini_bisnis_lv1',
		lb1.lini_bisnis1 'lini_bisnis_lv1_desc',
		ml.lini_bisnis_lv2 'lini_bisnis_lv2',
		lb2.lini_bisnis_lv2 'lini_bisnis_lv2_desc',
		ml.lini_bisnis_lv3 'lini_bisnis_lv3',
		lb3.lini_bisnis_lv3 'lini_bisnis_lv3_desc'
	FROM risk_issue_map_lini_bisnis ml
	LEFT JOIN lini_bisnis_lv1 lb1 ON lb1.kode_lini_bisnis = ml.lini_bisnis_lv1
	LEFT JOIN lini_bisnis_lv2 lb2 ON lb2.kode_lini_bisnis_lv2 = ml.lini_bisnis_lv2
	LEFT JOIN lini_bisnis_lv3 lb3 ON lb3.kode_lini_bisnis_lv3 = ml.lini_bisnis_lv3
	WHERE id_risk_issue = ?`, id).Rows()

	if err != nil {
		return responses, err
	}

	defer rows.Close()
	var DataMp models.MapLiniBisnisResponseFinal

	for rows.Next() {
		mp.db.DB.ScanRows(rows, &DataMp)
		responses = append(responses, DataMp)
	}

	return responses, err
}

// Store implements MapLiniBisnisDefinition
func (mp MapLiniBisnisRepository) Store(request *models.MapLiniBisnis, tx *gorm.DB) (responses *models.MapLiniBisnis, err error) {
	return request, tx.Save(&request).Error
}

// Update implements MapLiniBisnisDefinition
func (mp MapLiniBisnisRepository) Update(request *models.MapLiniBisnis, tx *gorm.DB) (responses bool, err error) {
	return true, tx.Save(&request).Error
}

// WithTrx implements MapLiniBisnisDefinition
func (mp MapLiniBisnisRepository) WithTrx(trxHandle *gorm.DB) MapLiniBisnisRepository {
	if trxHandle == nil {
		mp.logger.Zap.Error("transaction Database not found in gin context.")
		return mp
	}

	mp.db.DB = trxHandle
	return mp
}
