package tasklists

import (
	"riskmanagement/lib"
	models "riskmanagement/models/tasklists"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type TasklistsDataAnomaliKRIDDefinition interface {
	Store(request *models.TasklistsAnomaliDataKRIDRequest, tx *gorm.DB) (responses *models.TasklistsAnomaliDataKRIDRequest, err error)
	Delete(request *models.TasklistsAnomaliDataKRIDDelete, tx *gorm.DB) (err error)
	GetDataAnomali(request models.TasklistDataAnomaliRequest) (response []models.TasklistDataAnomaliResponse, err error)
}

type TasklistsDataAnomaliKRIDRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewTasklistsDataAnomaliKRIDRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) TasklistsDataAnomaliKRIDDefinition {
	return TasklistsDataAnomaliKRIDRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

func (repo TasklistsDataAnomaliKRIDRepository) Store(request *models.TasklistsAnomaliDataKRIDRequest, tx *gorm.DB) (responses *models.TasklistsAnomaliDataKRIDRequest, err error){
	return request, tx.Create(&request).Error
}

func (repo TasklistsDataAnomaliKRIDRepository) Delete(request *models.TasklistsAnomaliDataKRIDDelete, tx *gorm.DB) (err error){
	return tx.Where("tasklist_id", request.TasklistID).Delete(&request).Error	
}

func (repo TasklistsDataAnomaliKRIDRepository) GetDataAnomali(request models.TasklistDataAnomaliRequest) (response []models.TasklistDataAnomaliResponse, err error){
	rows, err := repo.db.DB.Raw(`
				SELECT tasklists_data_anomali_krid.id AS "tasklist_id", tasklists_data_anomali_krid.object AS "object"
				FROM tasklists_data_anomali_krid
				WHERE tasklist_id = ?`, 
				request.TasklistID).Rows()

	defer rows.Scan()

	var tasklistsDataAnomali models.TasklistDataAnomaliResponse

	for rows.Next() {
		repo.db.DB.ScanRows(rows, &tasklistsDataAnomali)
		response = append(response, tasklistsDataAnomali)
	}
	return response, err	
}