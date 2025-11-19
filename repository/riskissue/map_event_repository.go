package riskissue

import (
	"riskmanagement/lib"
	models "riskmanagement/models/riskissue"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type MapEventDefinition interface {
	GetAll() (responses []models.MapEventResponse, err error)
	GetOne(id int64) (responses []models.MapEventResponse, err error)
	Store(request *models.MapEvent, tx *gorm.DB) (responses *models.MapEvent, err error)
	Update(request *models.MapEvent, tx *gorm.DB) (responses bool, err error)
	Delete(id int64) (err error)
	GetOneDataByID(id int64) (responses []models.MapEventResponseFinal, err error)
	DeleteDataByID(id int64, tx *gorm.DB) (err error)
	WithTrx(trxHandle *gorm.DB) MapEventRepository
}

type MapEventRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewMapEventRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) MapEventDefinition {
	return MapEventRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// Delete implements MapEventDefinition
func (mp MapEventRepository) Delete(id int64) (err error) {
	return mp.db.DB.Where("id = ?", id).Delete(&models.MapEventResponseFinal{}).Error
}

// DeleteDataByID implements MapEventDefinition
func (mp MapEventRepository) DeleteDataByID(id int64, tx *gorm.DB) (err error) {
	return tx.Where("id_risk_issue = ?", id).Delete(&models.MapEventResponseFinal{}).Error
}

// GetAll implements MapEventDefinition
func (mp MapEventRepository) GetAll() (responses []models.MapEventResponse, err error) {
	return responses, mp.db.DB.Find(&responses).Error
}

// GetOne implements MapEventDefinition
func (mp MapEventRepository) GetOne(id int64) (responses []models.MapEventResponse, err error) {
	return responses, mp.db.DB.Where("id = ?", id).Find(&responses).Error
}

// GetOneDataByID implements MapEventDefinition
func (mp MapEventRepository) GetOneDataByID(id int64) (responses []models.MapEventResponseFinal, err error) {
	rows, err := mp.db.DB.Raw(`
	SELECT
			me.id 'id',
			me.id_risk_issue 'id_risk_issue',
			me.event_type_lv1 'event_type_lv1',
			et1.event_type 'event_type_lv1_desc',
			me.event_type_lv2 'event_type_lv2',
			et2.event_type_lv2 'event_type_lv2_desc',
			me.event_type_lv3 'event_type_lv3',
			et3.event_type_lv3 'event_type_lv3_desc'
	FROM risk_issue_map_event me
	LEFT JOIN event_type_lv1 et1 ON et1.kode_event_type = me.event_type_lv1
	LEFT JOIN event_type_lv2 et2 ON et2.kode_event_type_lv2 = me.event_type_lv2
	LEFT JOIN event_type_lv3 et3 ON et3.kode_event_type_lv3 = me.event_type_lv3
	WHERE id_risk_issue = ?`, id).Rows()
	if err != nil {
		return responses, err
	}

	defer rows.Close()
	var DataMp models.MapEventResponseFinal

	for rows.Next() {
		mp.db.DB.ScanRows(rows, &DataMp)
		responses = append(responses, DataMp)
	}

	return responses, err
}

// Store implements MapEventDefinition
func (mp MapEventRepository) Store(request *models.MapEvent, tx *gorm.DB) (responses *models.MapEvent, err error) {
	return request, tx.Save(&request).Error
}

// Update implements MapEventDefinition
func (mp MapEventRepository) Update(request *models.MapEvent, tx *gorm.DB) (responses bool, err error) {
	return true, tx.Save(&request).Error
}

// WithTrx implements MapEventDefinition
func (mp MapEventRepository) WithTrx(trxHandle *gorm.DB) MapEventRepository {
	if trxHandle == nil {
		mp.logger.Zap.Error("transaction Database not found in gin context.")
		return mp
	}

	mp.db.DB = trxHandle
	return mp
}
