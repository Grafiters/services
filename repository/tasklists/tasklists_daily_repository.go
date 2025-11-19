package tasklists

import (
	"riskmanagement/lib"
	models "riskmanagement/models/tasklists"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type TasklistsDailyDefinition interface {
	GetTasklist(request *models.TasklistCheckRequest) (response models.TasklistCheckResponse, err error)
	StoreDaily(request *models.TasklistDailyStore, tx *gorm.DB) (response *models.TasklistDailyStore, err error)
	UpdateDaily(request *models.ProgresUpdateRequest, tx *gorm.DB) (response *models.ProgresUpdateRequest, err error)
}

type TasklistsDailyRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewTasklistsDailyRepository(db lib.Database, dbRaw lib.Databases, logger logger.Logger) TasklistsDailyDefinition {
	return TasklistsDailyRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

func (r TasklistsDailyRepository) GetTasklist(request *models.TasklistCheckRequest) (response models.TasklistCheckResponse, err error) {
	err = r.db.DB.Raw(`
		SELECT * FROM tasklists t JOIN tasklists_uker tu ON tu.tasklist_id = t.id
		WHERE t.risk_issue_id = ? AND t.activity_id = ? AND t.product_id = ? AND tu.branch = ?`, request.RiskIssueID, request.ActivityID, request.ProductID, request.BRANCH).Find(&response).Error

	if err != nil {
		r.logger.Zap.Error(err)
		return response, err
	}

	return response, err
}

func (r TasklistsDailyRepository) StoreDaily(request *models.TasklistDailyStore, tx *gorm.DB) (response *models.TasklistDailyStore, err error) {
	return request, tx.Save(&request).Error
}

func (r TasklistsDailyRepository) UpdateDaily(request *models.ProgresUpdateRequest, tx *gorm.DB) (response *models.ProgresUpdateRequest, err error) {
	return request, tx.Save(&request).Error
}
