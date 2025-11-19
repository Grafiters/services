package tasklists

import (
	"riskmanagement/lib"
	models "riskmanagement/models/tasklists"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type TasklistsUkerDefinition interface {
	Store(request *models.TasklistsUker, tx *gorm.DB) (responses *models.TasklistsUker, err error)
	Delete(request *models.TasklistsUker, tx *gorm.DB) (err error)
}

type TasklistsUkerRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewTasklistsUkerRepository(
	db 		lib.Database,
	dbRaw 	lib.Databases,
	logger 	logger.Logger,
) TasklistsUkerDefinition {
	return TasklistsUkerRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

func (ta TasklistsUkerRepository) Store(request *models.TasklistsUker, tx *gorm.DB) (responses *models.TasklistsUker, err error){
	return request, tx.Create(&request).Error
}

func (ta TasklistsUkerRepository) Delete(request *models.TasklistsUker, tx *gorm.DB) (err error){
	return tx.Where("tasklist_id", request.TasklistID).Delete(&request).Error
}