package linibisnislv2

import (
	"math"
	"riskmanagement/lib"
	models "riskmanagement/models/linibisnislv2"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type LiniBisnisLv2Definition interface {
	GetAll() (responses []models.LiniBisnisLv2Response, err error)
	GetAllWithPaginate(request *models.Paginate) (responses []models.LiniBisnisLv2Response, totalRows int, totalData int, err error)
	GetOne(id int64) (responses models.LiniBisnisLv2Response, err error)
	Store(request *models.LiniBisnisLv2Request) (responses bool, err error)
	Update(request *models.LiniBisnisLv2Request) (responses bool, err error)
	Delete(id int64) (err error)
	GetKodeLiniBisnis() (responses []models.KodeLiniBisnis, err error)
	WithTrx(trxHandle *gorm.DB) LiniBisnisLv2Repository
	GetLBByID(request *models.KodeLB1) (responses []models.LiniBisnisLv2Response, err error)
}

type LiniBisnisLv2Repository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewLiniBisnisLv2Repository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) LiniBisnisLv2Definition {
	return LiniBisnisLv2Repository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: 0,
	}
}

// Delete implements LiniBisnisLv2Definition
func (lb2 LiniBisnisLv2Repository) Delete(id int64) (err error) {
	return lb2.db.DB.Where("id = ?", id).Delete(&models.LiniBisnisLv2Response{}).Error
}

// GetAllWithPaginate implements EventTypeLv2Definition
func (lb1 LiniBisnisLv2Repository) GetAllWithPaginate(request *models.Paginate) (responses []models.LiniBisnisLv2Response, totalRows int, totalData int, err error) {
	rows, err := lb1.db.DB.Raw(`
	SELECT
		lbl.id 'id',
		lbl.id_lini_bisnis_lv1 'id_lini_bisnis_lv1',
		lbl.kode_lini_bisnis_lv2 'kode_lini_bisnis_lv2',
		lbl.lini_bisnis_lv2 'lini_bisnis_lv2',
		lbl.deskripsi 'deskripsi',
		lbl.status 'status',
		lbl.created_at 'created_at',
		lbl.updated_at 'updated_at'
	FROM lini_bisnis_lv2 lbl ORDER BY lbl.id LIMIT ? OFFSET ?
	`, request.Limit, request.Offset).Rows()
	if err != nil {
		return responses, totalRows, totalData, err
	}

	defer rows.Close()

	var linibisnislv1 models.LiniBisnisLv2Response

	for rows.Next() {
		lb1.db.DB.ScanRows(rows, &linibisnislv1)
		responses = append(responses, linibisnislv1)
	}

	rows.Close()

	paginateQuery := `SELECT COUNT(*) FROM lini_bisnis_lv2`
	err = lb1.dbRaw.DB.QueryRow(paginateQuery).Scan(&totalData)
	if err != nil {
		return responses, totalRows, totalData, err
	}

	if totalData > 0 {
		totalRows = int(math.Ceil(float64(totalData) / float64(request.Limit)))
	}
	return responses, totalRows, totalData, err
}

// GetAll implements LiniBisnisLv2Definition
func (lb2 LiniBisnisLv2Repository) GetAll() (responses []models.LiniBisnisLv2Response, err error) {
	return responses, lb2.db.DB.Where("status = ?", 1).Find(&responses).Error
}

// GetKodeLiniBisnis implements LiniBisnisLv2Definition
func (lb2 LiniBisnisLv2Repository) GetKodeLiniBisnis() (responses []models.KodeLiniBisnis, err error) {
	query := `SELECT 
				RIGHT(CONCAT("0000",(T.kode_lini_bisnis + 1)), 4)
			FROM(
				SELECT
					CAST(SUBSTRING_INDEX(lb2.kode_lini_bisnis_lv2,'.', -1) as DECIMAL) kode_lini_bisnis
				FROM lini_bisnis_lv2 lb2 
				ORDER BY lb2.id DESC LIMIT 1) 
			AS T`

	lb2.logger.Zap.Info(query)
	rows, err := lb2.dbRaw.DB.Query(query)
	defer rows.Close()

	lb2.logger.Zap.Info("rows ", rows)
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

// GetOne implements LiniBisnisLv2Definition
func (lb2 LiniBisnisLv2Repository) GetOne(id int64) (responses models.LiniBisnisLv2Response, err error) {
	return responses, lb2.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements LiniBisnisLv2Definition
func (lb2 LiniBisnisLv2Repository) Store(request *models.LiniBisnisLv2Request) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	err = lb2.db.DB.Save(&models.LiniBisnisLv2Request{
		IDLiniBisnisLv1:   request.IDLiniBisnisLv1,
		KodeLiniBisnisLv2: request.KodeLiniBisnisLv2,
		LiniBisnisLv2:     request.LiniBisnisLv2,
		Deskripsi:         request.Deskripsi,
		Status:            request.Status,
		CreatedAt:         &timeNow,
	}).Error
	if err != nil {
		return false, err
	}

	return true, err
}

// Update implements LiniBisnisLv2Definition
func (lb2 LiniBisnisLv2Repository) Update(request *models.LiniBisnisLv2Request) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	err = lb2.db.DB.Save(&models.LiniBisnisLv2Request{
		ID:                request.ID,
		IDLiniBisnisLv1:   request.IDLiniBisnisLv1,
		KodeLiniBisnisLv2: request.KodeLiniBisnisLv2,
		LiniBisnisLv2:     request.LiniBisnisLv2,
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

// WithTrx implements LiniBisnisLv2Definition
func (lb2 LiniBisnisLv2Repository) WithTrx(trxHandle *gorm.DB) LiniBisnisLv2Repository {
	if trxHandle == nil {
		lb2.logger.Zap.Error("transaction Database not found in gin context")
		return lb2
	}

	lb2.db.DB = trxHandle
	return lb2
}

// GetLBByID implements LiniBisnisLv2Definition
func (lb2 LiniBisnisLv2Repository) GetLBByID(request *models.KodeLB1) (responses []models.LiniBisnisLv2Response, err error) {
	if request.KodeLB != "" {
		// where := "WHERE id_lini_bisnis_lv1 = '" + request.KodeLB + "'"
		where := "WHERE id_lini_bisnis_lv1 = ?"

		query := `SELECT * FROM lini_bisnis_lv2 ` + where + ` AND status != 0`

		lb2.logger.Zap.Info(query)
		rows, err := lb2.dbRaw.DB.Query(query, request.KodeLB)
		defer rows.Close()

		lb2.logger.Zap.Info("rows =>", rows)
		if err != nil {
			return responses, err
		}

		response := models.LiniBisnisLv2Response{}
		for rows.Next() {
			_ = rows.Scan(
				&response.ID,
				&response.IDLiniBisnisLv1,
				&response.KodeLiniBisnisLv2,
				&response.LiniBisnisLv2,
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
