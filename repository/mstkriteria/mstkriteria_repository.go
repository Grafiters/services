package mstkriteria

import (
	"database/sql"
	"fmt"
	"math"
	"riskmanagement/lib"
	models "riskmanagement/models/mstkriteria"
	"strings"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

var (
	timeNow = lib.GetTimeNow("timestime")
)

type MstKriteriaDefinition interface {
	WithTrx(trxHandle *gorm.DB) MstKriteriaRepository
	GetAll(request models.FilterRequest) (responses []models.MstKriteriaResponse, err error)
	GetAllWithPaginate(request *models.FilterRequest) (responses []models.MstKriteriaResponse, totalRows int, totalData int, err error)
	GetKodeCriteria() (responses []models.KodeMstKriteria, err error)
	GetOne(id int64) (responses models.MstKriteriaResponse, err error)
	Store(request *models.MstKriteria, tx *gorm.DB) (responses *models.MstKriteria, err error)
	Update(request *models.MstKriteriaRequest, tx *gorm.DB) (responses bool, err error)
	Delete(id int64) (err error)

	// add by panji 18/12/2024
	GetCriteriaById(request models.CriteriaRequestById) (responses []models.MstKriteriaResponse, err error)
	GetCriteriaByPeriode(request models.PeriodeRequest) (responses []models.MstKriteriaHistoryResponses, err error)
	StoreHistory(request *models.MstKriteriaHistory, tx *gorm.DB) (status bool, err error)
}

type MstKriteriaRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewMstKriteriaRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) MstKriteriaDefinition {
	return MstKriteriaRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// GetAll implements MstKriteriaDefinition
func (mstKriteria MstKriteriaRepository) GetAll(request models.FilterRequest) (responses []models.MstKriteriaResponse, err error) {
	query := mstKriteria.db.DB.Where("(status = 1 AND active_date <= NOW() OR status = 0 AND inactive_date >= NOW())")
	if request.Restruck != "" {
		if request.Restruck == "1" {
			query.Where("restruck in ('1','0')")
		} else if request.Restruck == "0" {
			query.Where("restruck in (0)")
		}
	}

	err = query.Find(&responses).Error
	if err != nil {
		mstKriteria.logger.Zap.Error(err.Error())
	}
	return responses, err
}

// WithTrx implements MstKriteriaDefinition
func (mstKriteria MstKriteriaRepository) WithTrx(trxHandle *gorm.DB) MstKriteriaRepository {
	if trxHandle == nil {
		mstKriteria.logger.Zap.Error("transaction Database not found in gin context.")
		return mstKriteria
	}
	mstKriteria.db.DB = trxHandle
	return mstKriteria
}

// GetAllWithPaginate implements MstKriteriaDefinition
func (mstKriteria MstKriteriaRepository) GetAllWithPaginate(request *models.FilterRequest) (responses []models.MstKriteriaResponse, totalRows int, totalData int, err error) {
	var keyword string
	if request.Keyword != "" {
		keyword += `WHERE (
						mk.kode_criteria LIKE '%` + request.Keyword + `%'
						OR mk.criteria LIKE '%` + request.Keyword + `%'
					)`
	}

	rows, err := mstKriteria.db.DB.Raw(`
		SELECT
			mk.id 'id',
			mk.kode_criteria 'kode_criteria',
			mk.criteria 'criteria',
			mk.restruck 'restruck',
			mk.status 'status',
			mk.active_date 'active_date',
			mk.inactive_date 'inactive_date',
			mk.created_at 'created_at',
			mk.created_by 'created_by',
			mk.created_desc 'created_desc',
			mk.enabled_date 'enabled_date',
			mk.enabled_by 'enabled_by',
			mk.enabled_desc 'enabled_desc',
			mk.disabled_date 'disabled_date',
			mk.disabled_by 'disabled_by',
			mk.disabled_desc 'disabled_desc'
		FROM mst_kriteria mk `+keyword+` LIMIT ? OFFSET ?`, request.Limit, request.Offset).Rows()
	if err != nil {
		return responses, totalRows, totalData, err
	}

	defer rows.Close()

	var mstKriteriaResponse models.MstKriteriaResponse

	for rows.Next() {
		mstKriteria.db.DB.ScanRows(rows, &mstKriteriaResponse)
		responses = append(responses, mstKriteriaResponse)
	}

	paginateQuery := `SELECT COUNT(*) FROM mst_kriteria`
	err = mstKriteria.dbRaw.DB.QueryRow(paginateQuery).Scan(&totalData)
	if err != nil {
		return responses, totalRows, totalData, err
	}

	if totalData > 0 {
		totalRows = int(math.Ceil(float64(totalData) / float64(request.Limit)))
	}
	return responses, totalRows, totalData, err
}

// GetOne implements MstKriteriaDefinition
func (mstKriteria MstKriteriaRepository) GetOne(id int64) (responses models.MstKriteriaResponse, err error) {
	return responses, mstKriteria.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements MstKriteriaDefinition
func (mstKriteria MstKriteriaRepository) Store(request *models.MstKriteria, tx *gorm.DB) (responses *models.MstKriteria, err error) {
	return request, tx.Save(request).Error
}

// Update implements MstKriteriaDefinition
func (mstKriteria MstKriteriaRepository) Update(request *models.MstKriteriaRequest, tx *gorm.DB) (responses bool, err error) {
	err = tx.Save(&models.MstKriteriaRequest{
		ID:           request.ID,
		KodeCriteria: request.KodeCriteria,
		Criteria:     request.Criteria,
		Restruck:     request.Restruck,
		Status:       request.Status,
		ActiveDate:   request.ActiveDate,
		InactiveDate: request.InactiveDate,
		CreatedAt:    &timeNow,
		CreatedBy:    request.CreatedBy,
		CreatedDesc:  request.CreatedDesc,
		EnabledDate:  request.EnabledDate,
		EnabledBy:    request.EnabledBy,
		EnabledDesc:  request.EnabledDesc,
		DisabledDate: request.DisabledDate,
		DisabledBy:   request.DisabledBy,
		DisabledDesc: request.DisabledDesc,
	}).Error

	if err != nil {
		return false, err
	}

	return true, err
}

// Delete implements MstKriteriaDefinition
func (mstKriteria MstKriteriaRepository) Delete(id int64) (err error) {
	return mstKriteria.db.DB.Where("id = ?", id).Delete(&models.MstKriteriaResponse{}).Error
}

func (subActivity MstKriteriaRepository) GetKodeCriteria() (responses []models.KodeMstKriteria, err error) {
	var resultID int
	rowsCheck, err := subActivity.dbRaw.DB.Query("SELECT COUNT(*) FROM mst_kriteria")
	checkErr(err)

	if checkCount(rowsCheck) == 0 {
		resultID = 1
	} else {
		query := `SELECT (t.kode_criteria + 1) 'kode_criteria' FROM (
						SELECT
							CAST(SUBSTRING_INDEX(kode_criteria ,'.', -1) AS DECIMAL) 'kode_criteria'
						FROM mst_kriteria) AS t
				  ORDER BY t.kode_criteria DESC LIMIT 1`

		//SELECT (count(*) + 1) 'kode_sub_activity' FROM sub_activity WHERE kode_sub_activity like "1.%";

		subActivity.logger.Zap.Info(query)
		rows, err := subActivity.dbRaw.DB.Query(query)
		checkErr(err)

		defer rows.Close()

		subActivity.logger.Zap.Info("rows ", rows)
		if err != nil {
			return responses, err
		}

		for rows.Next() {
			_ = rows.Scan(
				&resultID,
			)
		}

		if err = rows.Err(); err != nil {
			return responses, err
		}
	}
	result := fmt.Sprintf("CRI.%03d", (resultID))
	response := models.KodeMstKriteria{
		KodeCriteria: result,
	}
	responses = append(responses, response)
	return responses, err
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

// GetCriteriaById implements MstKriteriaDefinition.
func (mstKriteria MstKriteriaRepository) GetCriteriaById(request models.CriteriaRequestById) (responses []models.MstKriteriaResponse, err error) {
	id_criteria := strings.Split(request.Id, ",")

	query := mstKriteria.db.DB.Where(`id in (?)`, id_criteria)

	if request.Restruck != "" {
		if request.Restruck == "1" {
			query.Where("restruck in ('1','0')")
		} else if request.Restruck == "0" {
			query.Where("restruck in (0)")
		}
	}

	err = query.Find(&responses).Error
	if err != nil {
		mstKriteria.logger.Zap.Error(err.Error())
	}
	return responses, err
}

// GetCriteriaByPeriode implements MstKriteriaDefinition.
func (mstKriteria MstKriteriaRepository) GetCriteriaByPeriode(request models.PeriodeRequest) (responses []models.MstKriteriaHistoryResponses, err error) {
	query := mstKriteria.db.DB.Table(`mst_kriteria_history`).
		Where(`DATE(active_date) >= ? AND DATE(active_date) <= ?`, request.TglAwal, request.TglAkhir).Group(`id_criteria`)

	err = query.Find(&responses).Error
	if err != nil {
		mstKriteria.logger.Zap.Error(err.Error())
	}

	return responses, err
}

// StoreHistory implements MstKriteriaDefinition.
func (mstKriteria MstKriteriaRepository) StoreHistory(request *models.MstKriteriaHistory, tx *gorm.DB) (status bool, err error) {
	return true, tx.Save(request).Error
}
