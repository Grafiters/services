package linibisnislv1

import (
	"math"
	"riskmanagement/lib"
	models "riskmanagement/models/linibisnislv1"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type LiniBisnisLv1Definition interface {
	GetAll() (responses []models.LiniBisnisLv1Response, err error)
	GetAllWithPaginate(request *models.Paginate) (responses []models.LiniBisnisLv1Response, totalRows int, totalData int, err error)
	GetOne(id int64) (responses models.LiniBisnisLv1Response, err error)
	Store(request *models.LiniBisnisLv1Request) (responses bool, err error)
	Update(request *models.LiniBisnisLv1Request) (responses bool, err error)
	Delete(id int64) (err error)
	GetKodeLiniBisnis() (responses []models.KodeLiniBisnis, err error)
	WithTrx(trxHandle *gorm.DB) LiniBisnisLv1Repository
}

type LiniBisnisLv1Repository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewLiniBisnisLv1Repository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) LiniBisnisLv1Definition {
	return LiniBisnisLv1Repository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: 0,
	}
}

// Delete implements LiniBisnisLv1Definition
func (lb1 LiniBisnisLv1Repository) Delete(id int64) (err error) {
	return lb1.db.DB.Where("id = ?", id).Delete(&models.LiniBisnisLv1Response{}).Error
}

// GetAllWithPaginate implements EventTypeLv2Definition
func (lb1 LiniBisnisLv1Repository) GetAllWithPaginate(request *models.Paginate) (responses []models.LiniBisnisLv1Response, totalRows int, totalData int, err error) {
	rows, err := lb1.db.DB.Raw(`
	SELECT
		lbl.id 'id',
		lbl.kode_lini_bisnis 'kode_lini_bisnis',
		lbl.lini_bisnis1 'lini_bisnis1',
		lbl.deskripsi 'deskripsi',
		lbl.status 'status',
		lbl.created_at 'created_at',
		lbl.updated_at 'updated_at'
	FROM lini_bisnis_lv1 lbl ORDER BY lbl.id LIMIT ? OFFSET ?
	`, request.Limit, request.Offset).Rows()
	if err != nil {
		return responses, totalRows, totalData, err
	}

	defer rows.Close()

	var linibisnislv1 models.LiniBisnisLv1Response

	for rows.Next() {
		lb1.db.DB.ScanRows(rows, &linibisnislv1)
		responses = append(responses, linibisnislv1)
	}

	paginateQuery := `SELECT COUNT(*) FROM lini_bisnis_lv1`
	err = lb1.dbRaw.DB.QueryRow(paginateQuery).Scan(&totalData)
	if err != nil {
		return responses, totalRows, totalData, err
	}

	rows.Close()

	if totalData > 0 {
		totalRows = int(math.Ceil(float64(totalData) / float64(request.Limit)))
	}
	return responses, totalRows, totalData, err
}

// GetAll implements LiniBisnisLv1Definition
func (lb1 LiniBisnisLv1Repository) GetAll() (responses []models.LiniBisnisLv1Response, err error) {
	return responses, lb1.db.DB.Where("status = ?", 1).Find(&responses).Error
}

// GetKodeLiniBisnis implements LiniBisnisLv1Definition
func (lb1 LiniBisnisLv1Repository) GetKodeLiniBisnis() (responses []models.KodeLiniBisnis, err error) {
	query := `SELECT 
				RIGHT(CONCAT("0000",(T.kode_lini_bisnis + 1)), 4)
			FROM(
				SELECT
					CAST(SUBSTRING_INDEX(lbl.kode_lini_bisnis,'.', -1) as DECIMAL) kode_lini_bisnis
				FROM lini_bisnis_lv1 lbl 
				ORDER BY lbl.id DESC LIMIT 1) 
			AS T`

	lb1.logger.Zap.Info(query)
	rows, err := lb1.dbRaw.DB.Query(query)
	defer rows.Close()

	lb1.logger.Zap.Info("rows ", rows)
	if err != nil {
		return responses, err
	}

	response := models.KodeLiniBisnis{}
	for rows.Next() {
		_ = rows.Scan(
			&response.KodeLiniBisnis,
		)

		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, err
	}

	return responses, err
}

// GetOne implements LiniBisnisLv1Definition
func (lb1 LiniBisnisLv1Repository) GetOne(id int64) (responses models.LiniBisnisLv1Response, err error) {
	return responses, lb1.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements LiniBisnisLv1Definition
func (lb1 LiniBisnisLv1Repository) Store(request *models.LiniBisnisLv1Request) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	err = lb1.db.DB.Save(&models.LiniBisnisLv1Request{
		KodeLiniBisnis: request.KodeLiniBisnis,
		LiniBisnis1:    request.LiniBisnis1,
		Deskripsi:      request.Deskripsi,
		Status:         request.Status,
		CreatedAt:      &timeNow,
	}).Error
	if err != nil {
		return false, err
	}

	return true, err
}

// Update implements LiniBisnisLv1Definition
func (lb1 LiniBisnisLv1Repository) Update(request *models.LiniBisnisLv1Request) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	err = lb1.db.DB.Save(&models.LiniBisnisLv1Request{
		ID:             request.ID,
		KodeLiniBisnis: request.KodeLiniBisnis,
		LiniBisnis1:    request.LiniBisnis1,
		Deskripsi:      request.Deskripsi,
		Status:         request.Status,
		CreatedAt:      request.CreatedAt,
		UpdatedAt:      &timeNow,
	}).Error
	if err != nil {
		return false, err
	}

	return true, err
}

// WithTrx implements LiniBisnisLv1Definition
func (lb1 LiniBisnisLv1Repository) WithTrx(trxHandle *gorm.DB) LiniBisnisLv1Repository {
	if trxHandle == nil {
		lb1.logger.Zap.Error("transaction Database not found in gin context")
		return lb1
	}

	lb1.db.DB = trxHandle
	return lb1
}
