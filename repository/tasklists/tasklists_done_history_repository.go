package tasklists

import (
	"riskmanagement/lib"
	models "riskmanagement/models/tasklists"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type TasklistsDoneHistoryDefinition interface {
	Store(request *models.TasklistsDoneHistory, tx *gorm.DB) (responses *models.TasklistsDoneHistory, err error)
	Get(request models.TasklistsDoneHistoryCheckRequest) (responses models.TasklistsDoneHistory, err error)
}

type TasklistsDoneHistoryRepository struct {
	db     		lib.Database
	dbRaw 		lib.Databases
	logger 		logger.Logger
	timeout 	time.Duration
}

func NewTasklistsDoneHistoryRepository(db lib.Database, dbRaw lib.Databases, logger logger.Logger) TasklistsDoneHistoryDefinition {
	return TasklistsDoneHistoryRepository{
		db:    	 db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

func (r TasklistsDoneHistoryRepository) Store(request *models.TasklistsDoneHistory, tx *gorm.DB) (responses *models.TasklistsDoneHistory, err error){
	return responses, tx.Create(&request).Error
}

func (r TasklistsDoneHistoryRepository) Get(request models.TasklistsDoneHistoryCheckRequest) (responses models.TasklistsDoneHistory, err error){
	err = r.db.DB.Raw(`
		SELECT * FROM tasklists_done_history WHERE pernr = ? and tasklist_id = ? and date(created_at) = ?`, request.PERNR, request.TasklistID, request.Date).Find(&responses).Error

	if err != nil {
		r.logger.Zap.Error(err)
		return responses, err
	}
	return responses, err	
}