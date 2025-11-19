package linibisnislv3

import (
	"math"
	"riskmanagement/lib"
	models "riskmanagement/models/linibisnislv3"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type LiniBisnisLv3Definition interface {
	GetAll() (responses []models.LiniBisnisLv3Response, err error)
	GetAllWithPaginate(request *models.Paginate) (responses []models.LiniBisnisLv3Response, totalRows int, totalData int, err error)
	GetOne(id int64) (responses models.LiniBisnisLv3Response, err error)
	Store(request *models.LiniBisnisLv3Request) (responses bool, err error)
	Update(request *models.LiniBisnisLv3Request) (responses bool, err error)
	Delete(id int64) (err error)
	GetKodeLiniBisnis() (responses []models.KodeLiniBisnis, err error)
	WithTrx(trxHandle *gorm.DB) LiniBisnisLv3Repository
	GetLBByID(request *models.KodeLB2) (responses []models.LiniBisnisLv3Response, err error)
}

type LiniBisnisLv3Repository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewLiniBisnisLv3Repository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) LiniBisnisLv3Definition {
	return LiniBisnisLv3Repository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: 0,
	}
}

// Delete implements LiniBisnisLv3Definition
func (lb3 LiniBisnisLv3Repository) Delete(id int64) (err error) {
	return lb3.db.DB.Where("id = ?", id).Delete(&models.LiniBisnisLv3Response{}).Error
}

// GetAllWithPaginate implements LiniBisnisLv3Definition
func (lb3 LiniBisnisLv3Repository) GetAllWithPaginate(request *models.Paginate) (responses []models.LiniBisnisLv3Response, totalRows int, totalData int, err error) {
	rows, err := lb3.db.DB.Raw(`
	SELECT
		lbl.id 'id',
		lbl.id_lini_bisnis_lv2 'id_lini_bisnis_lv2',
		lbl.kode_lini_bisnis_lv3 'kode_lini_bisnis_lv3',
		lbl.lini_bisnis_lv3 'lini_bisnis_lv3',
		lbl.deskripsi 'deskripsi',
		lbl.status 'status',
		lbl.created_at 'created_at',
		lbl.updated_at 'updated_at'
	FROM lini_bisnis_lv3 lbl ORDER BY lbl.id LIMIT ? OFFSET ?
	`, request.Limit, request.Offset).Rows()
	if err != nil {
		return responses, totalRows, totalData, err
	}

	defer rows.Close()

	var linibisnislv1 models.LiniBisnisLv3Response
	for rows.Next() {
		lb3.db.DB.ScanRows(rows, &linibisnislv1)
		responses = append(responses, linibisnislv1)
	}

	paginateQuery := `SELECT COUNT(*) FROM lini_bisnis_lv3`
	err = lb3.dbRaw.DB.QueryRow(paginateQuery).Scan(&totalData)
	if err != nil {
		return responses, totalRows, totalData, err
	}

	if totalData > 0 {
		totalRows = int(math.Ceil(float64(totalData) / float64(request.Limit)))
	}
	return responses, totalRows, totalData, err
}

// GetAll implements LiniBisnisLv3Definition
func (lb3 LiniBisnisLv3Repository) GetAll() (responses []models.LiniBisnisLv3Response, err error) {
	return responses, lb3.db.DB.Find(&responses).Error
}

// GetKodeLiniBisnis implements LiniBisnisLv3Definition
func (lb3 LiniBisnisLv3Repository) GetKodeLiniBisnis() (responses []models.KodeLiniBisnis, err error) {
	query := `SELECT 
				RIGHT(CONCAT("0000",(T.kode_lini_bisnis + 1)), 4)
			FROM(
				SELECT
					CAST(SUBSTRING_INDEX(lb3.kode_lini_bisnis_lv3,'.', -1) as DECIMAL) kode_lini_bisnis
				FROM lini_bisnis_lv3 lb3 
				ORDER BY lb3.id DESC LIMIT 1) 
			AS T`

	lb3.logger.Zap.Info(query)
	rows, err := lb3.dbRaw.DB.Query(query)
	defer rows.Close()

	lb3.logger.Zap.Info("rows ", rows)
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

// GetOne implements LiniBisnisLv3Definition
func (lb3 LiniBisnisLv3Repository) GetOne(id int64) (responses models.LiniBisnisLv3Response, err error) {
	return responses, lb3.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements LiniBisnisLv3Definition
func (lb3 LiniBisnisLv3Repository) Store(request *models.LiniBisnisLv3Request) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	err = lb3.db.DB.Save(&models.LiniBisnisLv3Request{
		IDLiniBisnisLv2:   request.IDLiniBisnisLv2,
		KodeLiniBisnisLv3: request.KodeLiniBisnisLv3,
		LiniBisnisLv3:     request.LiniBisnisLv3,
		Deskripsi:         request.Deskripsi,
		Status:            request.Status,
		CreatedAt:         &timeNow,
	}).Error
	if err != nil {
		return false, err
	}

	return true, err
}

// Update implements LiniBisnisLv3Definition
func (lb3 LiniBisnisLv3Repository) Update(request *models.LiniBisnisLv3Request) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	err = lb3.db.DB.Save(&models.LiniBisnisLv3Request{
		ID:                request.ID,
		IDLiniBisnisLv2:   request.IDLiniBisnisLv2,
		KodeLiniBisnisLv3: request.KodeLiniBisnisLv3,
		LiniBisnisLv3:     request.LiniBisnisLv3,
		Deskripsi:         request.Deskripsi,
		Status:            request.Status,
		CreatedAt:         request.CreatedAt,
		UpdatedAt:         &timeNow,
	}).Error
	if err != nil {
		return false, err
	}

	return true, err
}

// WithTrx implements LiniBisnisLv3Definition
func (lb3 LiniBisnisLv3Repository) WithTrx(trxHandle *gorm.DB) LiniBisnisLv3Repository {
	if trxHandle == nil {
		lb3.logger.Zap.Error("transaction Database not found in gin context")
		return lb3
	}

	lb3.db.DB = trxHandle
	return lb3
}

// GetLBByID implements LiniBisnisLv2Definition
func (lb3 LiniBisnisLv3Repository) GetLBByID(request *models.KodeLB2) (responses []models.LiniBisnisLv3Response, err error) {
	if request.KodeLB != "" {
		// where := "WHERE id_lini_bisnis_lv2 = '" + request.KodeLB + "'"
		where := "WHERE id_lini_bisnis_lv2 = ?"

		query := `SELECT * FROM lini_bisnis_lv3 ` + where + ` AND status != 0`

		lb3.logger.Zap.Info(query)
		rows, err := lb3.dbRaw.DB.Query(query, request.KodeLB)
		defer rows.Close()

		lb3.logger.Zap.Info("rows =>", rows)
		if err != nil {
			return responses, err
		}

		response := models.LiniBisnisLv3Response{}
		for rows.Next() {
			_ = rows.Scan(
				&response.ID,
				&response.IDLiniBisnisLv2,
				&response.KodeLiniBisnisLv3,
				&response.LiniBisnisLv3,
				&response.Deskripsi,
				&response.Status,
				&response.CreatedAt,
				&response.UpdatedAt,
			)

			responses = append(responses, response)
		}

		if err = rows.Err(); err != nil {
			return responses, err
		}

	}
	return responses, err
}
