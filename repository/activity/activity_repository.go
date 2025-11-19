package activity

import (
	"fmt"
	"math"
	"riskmanagement/lib"
	models "riskmanagement/models/activity"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type ActivityDefinition interface {
	GetAll() (responses []models.ActivityResponse, err error)
	GetAllWithPagination(request *models.Paginate) (response []models.ActivityResponse, totalRows int, totalData int, err error)
	GetOne(id int64) (responses models.ActivityResponse, err error)
	Store(request *models.ActivityRequest) (responses bool, err error)
	Update(request *models.ActivityRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) ActivityRepository
	GetKodeActivity() (responses []models.KodeActivity, err error)
}

type ActivityRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewActivityRepository(db lib.Database, dbRaw lib.Databases, logger logger.Logger) ActivityDefinition {
	return ActivityRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// Delete implements ActicityDefinition
func (activity ActivityRepository) Delete(id int64) (err error) {
	return activity.db.DB.Where("id = ?", id).Delete(&models.ActivityResponse{}).Error
}

// GetAll implements ActicityDefinition
func (activity ActivityRepository) GetAll() (responses []models.ActivityResponse, err error) {
	return responses, activity.db.DB.Find(&responses).Error
}

// GetAllWithPagination implements ActivityDefinition
func (activity ActivityRepository) GetAllWithPagination(request *models.Paginate) (response []models.ActivityResponse, totalRows int, totalData int, err error) {
	rows, err := activity.db.DB.Raw(`
				SELECT
					act.id 'id',
					act.kode_activity 'kode_activity',
					act.name 'name',
					act.create_at 'create_at',
					act.update_at 'update_at'
				FROM activity act ORDER BY act.id ASC LIMIT ? OFFSET ?`, request.Limit, request.Offset).Rows()
	if err != nil {
		return response, totalRows, totalData, err
	}

	defer rows.Close()

	defer rows.Scan()
	var aktifitas models.ActivityResponse

	for rows.Next() {
		activity.db.DB.ScanRows(rows, &aktifitas)
		response = append(response, aktifitas)
	}

	paginateQuery := `SELECT count(*) FROM activity`
	err = activity.dbRaw.DB.QueryRow(paginateQuery).Scan(&totalData)
	if err != nil {
		return response, totalRows, totalData, err
	}

	if totalData > 0 {
		totalRows = int(math.Ceil(float64(totalData) / float64(request.Limit)))
	}

	return response, totalRows, totalData, err
}

// GetOne implements ActicityDefinition
func (activity ActivityRepository) GetOne(id int64) (responses models.ActivityResponse, err error) {
	return responses, activity.db.DB.Where("kode_activity = ?", id).Find(&responses).Error

}

// Store implements ActicityDefinition
func (activity ActivityRepository) Store(request *models.ActivityRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	fmt.Println("repo = ", models.ActivityRequest{
		Name:         request.Name,
		KodeActivity: request.KodeActivity,
		CreateAt:     &timeNow,
	})

	err = activity.db.DB.Save(&models.ActivityRequest{
		Name:         request.Name,
		KodeActivity: request.KodeActivity,
		CreateAt:     &timeNow,
	}).Error
	if err != nil {
		return false, err
	}

	return true, err
}

// Update implements ActicityDefinition
func (activity ActivityRepository) Update(request *models.ActivityRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	err = activity.db.DB.Save(&models.ActivityRequest{
		ID:           request.ID,
		KodeActivity: request.KodeActivity,
		Name:         request.Name,
		CreateAt:     request.CreateAt,
		UpdateAt:     &timeNow,
	}).Error
	if err != nil {
		return false, err
	}

	return true, err
}

// WithTrx implements ActicityDefinition
func (activity ActivityRepository) WithTrx(trxHandle *gorm.DB) ActivityRepository {
	if trxHandle == nil {
		activity.logger.Zap.Error("transaction Database not found in gin context")
		return activity
	}
	activity.db.DB = trxHandle
	return activity
}

// getKodeActivity implements ActivityDefinition
func (activity ActivityRepository) GetKodeActivity() (responses []models.KodeActivity, err error) {
	// query := `SELECT RIGHT(CONCAT("00",(count(*) + 1)), 2) 'kode_activity' FROM activity`
	query := `SELECT 
					RIGHT(CONCAT("00",(t.kode_activity  + 1)), 2) 'kode_activity'
				FROM(
					SELECT 
						CAST(kode_activity AS DECIMAL) 'kode_activity'
					FROM activity 
				) AS t ORDER BY t.kode_activity DESC LIMIT 1`

	activity.logger.Zap.Info(query)
	rows, err := activity.dbRaw.DB.Query(query)
	defer rows.Close()

	activity.logger.Zap.Info("Rows ", rows)
	if err != nil {
		return responses, err
	}

	response := models.KodeActivity{}
	for rows.Next() {
		_ = rows.Scan(
			&response.KodeActivity,
		)

		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, err
	}

	return responses, err
}
