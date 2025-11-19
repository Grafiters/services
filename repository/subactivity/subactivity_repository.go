package subactivity

import (
	"database/sql"
	"fmt"
	"math"
	"riskmanagement/lib"
	models "riskmanagement/models/subactivity"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type SubActivityDefinition interface {
	GetAll() (responses []models.SubActivityResponses, err error)
	GetAllWithPagination(request *models.Paginate) (responses []models.SubActivityResponses, totalRows int, totalData int, err error)
	GetOne(id int64) (responses models.SubActivityResponse, err error)
	GetLastID(id int64) (responses []models.SubActivityResponse, err error)
	Store(request *models.SubActivityRequest) (responses bool, err error)
	Update(request *models.SubActivityRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) SubActivityRepository
	GetKodeSubActivity(requests models.KodeSubActivityRequest) (responses []models.KodeSubActivityResponse, err error)
}

type SubActivityRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func checkCount(rows *sql.Rows) (count int) {
	for rows.Next() {
		err := rows.Scan(&count)
		checkErr(err)
	}
	return count
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// GetAll implements SubActivityDefinition
func (subactivity SubActivityRepository) GetAll() (responses []models.SubActivityResponses, err error) {
	// return responses, subactivity.db.DB.Find(&responses).Error
	rows, err := subactivity.db.DB.Raw(`
		SELECT 
			sub.id 'id',
			sub.activity_id 'activity_id',
			act.name 'activity_name',
			sub.kode_sub_activity 'kode_sub_activity',
			sub.name 'name_sub_activity',
			sub.created_at 'created_at',
			sub.updated_at 'updated_at'
		FROM sub_activity sub
		JOIN activity act ON act.kode_activity = sub.activity_id
	`).Rows()

	defer rows.Close()

	var subAct models.SubActivityResponses

	for rows.Next() {
		subactivity.db.DB.ScanRows(rows, &subAct)
		responses = append(responses, subAct)
	}

	return responses, err
}

func NewSubActivityRepository(db lib.Database, dbRaw lib.Databases, logger logger.Logger) SubActivityDefinition {
	return SubActivityRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// GetLastID implements SubActivityDefinition
func (subactivity SubActivityRepository) GetLastID(id int64) (responses []models.SubActivityResponse, err error) {
	return responses, subactivity.db.DB.Where("activity_id = ?", id).Find(&responses).Error
}

// Delete implements SubActivityDefinition
func (subactivity SubActivityRepository) Delete(id int64) (err error) {
	return subactivity.db.DB.Where("id = ?", id).Delete(&models.SubActivityResponse{}).Error
}

// GetAll implements SubActivityDefinition
// func (subactivity SubActivityRepository) GetAll() (responses []models.SubActivityResponse, err error) {
// 	return responses, subactivity.db.DB.Find(&responses).Error
// }

// GetAll implements SubActivityDefinition
func (subactivity SubActivityRepository) GetAllWithPagination(request *models.Paginate) (responses []models.SubActivityResponses, totalRows int, totalData int, err error) {
	// return responses, subactivity.db.DB.Find(&responses).Error
	rows, err := subactivity.db.DB.Raw(`
		SELECT 
			sub.id 'id',
			sub.activity_id 'activity_id',
			act.name 'activity_name',
			sub.kode_sub_activity 'kode_sub_activity',
			sub.name 'name_sub_activity',
			sub.created_at 'created_at',
			sub.updated_at 'updated_at'
		FROM sub_activity sub
		LEFT JOIN activity act ON act.kode_activity = sub.activity_id ORDER BY sub.id ASC LIMIT ? OFFSET ?
	`, request.Limit, request.Offset).Rows()
	if err != nil {
		return responses, totalRows, totalData, err
	}
	defer rows.Close()

	var subAct models.SubActivityResponses

	for rows.Next() {
		subactivity.db.DB.ScanRows(rows, &subAct)
		responses = append(responses, subAct)
	}

	paginateQuery := `SELECT 
						count(*) 
					FROM sub_activity sub 
					LEFT JOIN activity act ON act.id = sub.activity_id`
	err = subactivity.dbRaw.DB.QueryRow(paginateQuery).Scan(&totalData)
	if err != nil {
		return responses, totalRows, totalData, err
	}

	if totalData > 0 {
		totalRows = int(math.Ceil((float64(totalData)) / float64(request.Limit)))
	}

	return responses, totalRows, totalData, err
}

// GetOne implements SubActivityDefinition
func (subactivity SubActivityRepository) GetOne(id int64) (responses models.SubActivityResponse, err error) {
	return responses, subactivity.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements SubActivityDefinition
func (subactivity SubActivityRepository) Store(request *models.SubActivityRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	err = subactivity.db.DB.Save(&models.SubActivityRequest{
		ActivityID:      request.ActivityID,
		KodeSubActivity: request.KodeSubActivity,
		Name:            request.Name,
		CreatedAt:       &timeNow,
	}).Error
	if err != nil {
		return false, err
	}

	return true, err
}

// Update implements SubActivityDefinition
func (subactivity SubActivityRepository) Update(request *models.SubActivityRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	err = subactivity.db.DB.Save(&models.SubActivityRequest{
		ID:              request.ID,
		KodeSubActivity: request.KodeSubActivity,
		ActivityID:      request.ActivityID,
		Name:            request.Name,
		CreatedAt:       request.CreatedAt,
		UpdatedAt:       &timeNow,
	}).Error
	if err != nil {
		return false, err
	}

	return true, err
}

// WithTrx implements SubActivityDefinition
func (subactivity SubActivityRepository) WithTrx(trxHandle *gorm.DB) SubActivityRepository {
	if trxHandle == nil {
		subactivity.logger.Zap.Error("transaction Database not found in gin context")
		return subactivity
	}

	subactivity.db.DB = trxHandle
	return subactivity
}

// GetKodeSubActivity implements SubActivityDefinition
func (subActivity SubActivityRepository) GetKodeSubActivity(requests models.KodeSubActivityRequest) (responses []models.KodeSubActivityResponse, err error) {
	kode := ""
	if requests.KodeActivity != "" {
		kode += requests.KodeActivity + "."
	}

	key := fmt.Sprintf("%s%%", kode)

	// query := `SELECT (count(*) + 1) 'kode_sub_activity' FROM sub_activity WHERE kode_sub_activity like '` + kode + `%'`
	rowsCheck, err := subActivity.dbRaw.DB.Query("SELECT COUNT(*) FROM sub_activity WHERE kode_sub_activity LIKE ?", key)
	checkErr(err)

	if checkCount(rowsCheck) == 0 {
		fmt.Println("masih 0")
		response := models.KodeSubActivityResponse{
			KodeSubActivity: "1",
		}
		responses = append(responses, response)
	} else {
		fmt.Println("dah gak 0")

		query := `SELECT (t.kode_sub_activity + 1) 'kode_sub_activity' FROM (
						SELECT
							CAST(SUBSTRING_INDEX(kode_sub_activity ,'.', -1) AS DECIMAL) 'kode_sub_activity'
						FROM sub_activity
						WHERE kode_sub_activity LIKE ?) AS t
					ORDER BY t.kode_sub_activity DESC LIMIT 1`

		//SELECT (count(*) + 1) 'kode_sub_activity' FROM sub_activity WHERE kode_sub_activity like "1.%";

		subActivity.logger.Zap.Info(query)
		rows, err := subActivity.dbRaw.DB.Query(query, key)
		defer rows.Close()

		subActivity.logger.Zap.Info("rows ", rows)
		if err != nil {
			return responses, err
		}

		response := models.KodeSubActivityResponse{}
		for rows.Next() {
			_ = rows.Scan(
				&response.KodeSubActivity,
			)

			responses = append(responses, response)
		}

		if err = rows.Err(); err != nil {
			return responses, err
		}
	}

	return responses, err
}
