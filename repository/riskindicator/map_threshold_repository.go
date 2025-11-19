package riskindicator

import (
	"riskmanagement/lib"
	models "riskmanagement/models/riskindicator"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type MapThresholdDefinition interface {
	GetThreshold(id int64) (responses []models.MapThresholdResponse, err error)
	SaveThreshold(request *models.MapThreshold, tx *gorm.DB) (responses *models.MapThreshold, err error)
	WithTrx(trxHandle *gorm.DB) MapThresholdRepository
}

type MapThresholdRepository struct {
	db      lib.Database
	logger  logger.Logger
	timeout time.Duration
}

func NewMapThresholdRepository(
	db lib.Database,
	logger logger.Logger,
) MapThresholdDefinition {
	return MapThresholdRepository{
		db:      db,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// GetThreshold implements MapThresholdDefinition
func (mt MapThresholdRepository) GetThreshold(id int64) (responses []models.MapThresholdResponse, err error) {
	return responses, mt.db.DB.Where("id_indicator = ?", id).Find(&responses).Error
}

// SaveThreshold implements MapThresholdDefinition
func (mt MapThresholdRepository) SaveThreshold(request *models.MapThreshold, tx *gorm.DB) (responses *models.MapThreshold, err error) {
	return request, tx.Save(&request).Error
}

// WithTrx implements MapThresholdDefinition
func (mt MapThresholdRepository) WithTrx(trxHandle *gorm.DB) MapThresholdRepository {
	if trxHandle == nil {
		mt.logger.Zap.Error("transaction Database not found in gin context.")
		return mt
	}

	mt.db.DB = trxHandle
	return mt
}
