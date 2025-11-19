package tasklists

import (
	"riskmanagement/lib"
	models "riskmanagement/models/tasklists"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type TasklistsActivityDefinition interface {
	Store(request *models.TasklistsRiskIndicator, tx *gorm.DB) (responses *models.TasklistsRiskIndicator, err error)
	Delete(request *models.TasklistsRiskIndicator, tx *gorm.DB) (err error)
}

type TasklistsActivityRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewTasklistsActivityRepository(
	db 		lib.Database,
	dbRaw 	lib.Databases,
	logger 	logger.Logger,
) TasklistsActivityDefinition {
	return TasklistsActivityRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

func (ta TasklistsActivityRepository) Store(request *models.TasklistsRiskIndicator, tx *gorm.DB) (responses *models.TasklistsRiskIndicator, err error){
	return request, tx.Create(&request).Error
}

func (ta TasklistsActivityRepository) Delete(request *models.TasklistsRiskIndicator, tx *gorm.DB) (err error){
	return tx.Where("tasklists_id", request.TasklistsID).Delete(&request).Error
}