package penyebabkejadianlv3

import (
	"math"
	"riskmanagement/lib"
	models "riskmanagement/models/penyebabkejadianlv3"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type PenyebabKejadianLv3Definition interface {
	// GetAll() (responses []models.PenyebabKejadianLv3Response, err error)
	GetAll() (responses []models.PenyebabKejadianLv3Responses, err error)
	GetAllWithPaginate(request *models.Paginate) (responses []models.PenyebabKejadianLv3Responses, totalRows int, totalData int, err error)
	GetOne(id int64) (responses models.PenyebabKejadianLv3Response, err error)
	Store(request *models.PenyebabKejadianLv3Request) (responses bool, err error)
	Update(request *models.PenyebabKejadianLv3Request) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) PenyebabKejadianLv3Repository
	GetKodePenyebabKejadian() (responses []models.KodePenyebabKejadian, err error)
	GetKejadianByIDlv2(request *models.KodeSubKejadianRequest) (responses []models.PenyebabKejadianLv3Response, err error)
	GetKejadianByIDlv1(request *models.KodePenyebabKejadian) (responses []models.PenyebabKejadianLv3Response, err error)
	GetSubKejadian(id int64) (responses []models.PenyebabKejadianLv3Response, err error)
}

type PenyebabKejadianLv3Repository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewPenyebabKejadianLv3Repository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) PenyebabKejadianLv3Definition {
	return PenyebabKejadianLv3Repository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// Delete implements PenyebabKejadianLv3Definition
func (PK3 PenyebabKejadianLv3Repository) Delete(id int64) (err error) {
	return PK3.db.DB.Where("id = ?", id).Delete(&models.PenyebabKejadianLv3Response{}).Error
}

// GetAll implements PenyebabKejadianLv3Definition
//
//	func (PK3 PenyebabKejadianLv3Repository) GetAll() (responses []models.PenyebabKejadianLv3Response, err error) {
//		return responses, pK3.db.DB.Find(&responses).Error
//	}
func (PK3 PenyebabKejadianLv3Repository) GetAll() (responses []models.PenyebabKejadianLv3Responses, err error) {
	rows, err := PK3.db.DB.Raw(`
			SELECT
				pk3.id 'id',
				pk2.kode_sub_kejadian 'kode_sub_kejadian',
				pk2.kriteria_penyebab_kejadian 'penyebab_kejadian_lv2',
				pk3.kode_penyebab_kejadian_lv3 'kode_penyebab_kejadian_lv3',
				pk3.penyebab_kejadian_lv3 'penyebab_kejadian_lv3',
				pk3.deskripsi 'deskripsi',
				pk3.status 'status'
			FROM penyebab_kejadian_lv3 pk3
			LEFT OUTER JOIN penyebab_kejadian_lv2 pk2 on pk2.kode_sub_kejadian = pk3.kode_sub_kejadian
	`).Rows()

	defer rows.Close()

	var subInci models.PenyebabKejadianLv3Responses

	for rows.Next() {
		PK3.db.DB.ScanRows(rows, &subInci)
		responses = append(responses, subInci)
	}

	return responses, err
}

// GetAllWithPaginate implements PenyebabKejadianLv3Definition
func (PK3 PenyebabKejadianLv3Repository) GetAllWithPaginate(request *models.Paginate) (responses []models.PenyebabKejadianLv3Responses, totalRows int, totalData int, err error) {
	rows, err := PK3.db.DB.Raw(`
			SELECT
				pk3.id 'id',
				pk2.kode_sub_kejadian 'kode_sub_kejadian',
				pk2.kriteria_penyebab_kejadian 'penyebab_kejadian_lv2',
				pk3.kode_penyebab_kejadian_lv3 'kode_penyebab_kejadian_lv3',
				pk3.penyebab_kejadian_lv3 'penyebab_kejadian_lv3',
				pk3.deskripsi 'deskripsi',
				pk3.status 'status'
			FROM penyebab_kejadian_lv3 pk3
			LEFT OUTER JOIN penyebab_kejadian_lv2 pk2 on pk2.kode_sub_kejadian = pk3.kode_sub_kejadian ORDER BY pk3.id LIMIT ? OFFSET ?
	`, request.Limit, request.Offset).Rows()
	if err != nil {
		return responses, totalRows, totalData, err
	}

	defer rows.Close()

	var pkLevel3 models.PenyebabKejadianLv3Responses

	for rows.Next() {
		PK3.db.DB.ScanRows(rows, &pkLevel3)
		responses = append(responses, pkLevel3)
	}

	paginateQuery := `SELECT COUNT(*) FROM penyebab_kejadian_lv3`
	err = PK3.dbRaw.DB.QueryRow(paginateQuery).Scan(&totalData)
	if err != nil {
		return responses, totalRows, totalData, err
	}

	if totalData > 0 {
		totalRows = int(math.Ceil((float64(totalData)) / float64(request.Limit)))
	}
	return responses, totalRows, totalData, err
}

// GetOne implements PenyebabKejadianLv3Definition
func (PK3 PenyebabKejadianLv3Repository) GetOne(id int64) (responses models.PenyebabKejadianLv3Response, err error) {
	return responses, PK3.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements PenyebabKejadianLv3Definition
func (PK3 PenyebabKejadianLv3Repository) Store(request *models.PenyebabKejadianLv3Request) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	err = PK3.db.DB.Save(&models.PenyebabKejadianLv3Request{
		KodeSubKejadian:         request.KodeSubKejadian,
		KodePenyebabKejadianLv3: request.KodePenyebabKejadianLv3,
		PenyebabKejadianLv3:     request.PenyebabKejadianLv3,
		Deskripsi:               request.Deskripsi,
		Status:                  request.Status,
		CreatedAt:               &timeNow,
	}).Error
	if err != nil {
		return false, err
	}

	return true, err
}

// Update implements PenyebabKejadianLv3Definition
func (PK3 PenyebabKejadianLv3Repository) Update(request *models.PenyebabKejadianLv3Request) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	err = PK3.db.DB.Save(&models.PenyebabKejadianLv3Request{
		ID:                      request.ID,
		KodeSubKejadian:         request.KodeSubKejadian,
		KodePenyebabKejadianLv3: request.KodePenyebabKejadianLv3,
		PenyebabKejadianLv3:     request.PenyebabKejadianLv3,
		Deskripsi:               request.Deskripsi,
		Status:                  request.Status,
		CreatedAt:               request.CreatedAt,
		UpdatedAt:               &timeNow,
	}).Error
	if err != nil {
		return false, err
	}

	return true, err
}

// WithTrx implements PenyebabKejadianLv3Definition
func (PK3 PenyebabKejadianLv3Repository) WithTrx(trxHandle *gorm.DB) PenyebabKejadianLv3Repository {
	if trxHandle == nil {
		PK3.logger.Zap.Error("transaction Database not found in gin context")
		return PK3
	}

	PK3.db.DB = trxHandle
	return PK3
}

func (PK3 PenyebabKejadianLv3Repository) GetKodePenyebabKejadian() (responses []models.KodePenyebabKejadian, err error) {
	query := `SELECT 
				RIGHT(CONCAT("0000",(T.kode_penyebab_kejadian_lv3 + 1)), 4)
			FROM(
				SELECT
					CAST(SUBSTRING_INDEX(pkl.kode_penyebab_kejadian_lv3,'.', -1) as DECIMAL) 'kode_penyebab_kejadian_lv3'
				FROM penyebab_kejadian_lv3 pkl 
				ORDER BY pkl.id DESC LIMIT 1) 
			AS T`

	PK3.logger.Zap.Info(query)
	rows, err := PK3.dbRaw.DB.Query(query)
	defer rows.Close()

	PK3.logger.Zap.Info("rows", rows)
	for err != nil {
		return responses, err
	}

	response := models.KodePenyebabKejadian{}
	for rows.Next() {
		_ = rows.Scan(
			&response.KodePenyebabKejadian,
		)

		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, err
	}

	return responses, err
}

// GetKejadianByIDlv2 implements PenyebabKejadianLv3Definition
func (PK3 PenyebabKejadianLv3Repository) GetKejadianByIDlv2(request *models.KodeSubKejadianRequest) (responses []models.PenyebabKejadianLv3Response, err error) {
	if request.KodeSubKejadian != "" {
		where := " WHERE kode_sub_kejadian = ?"

		query := `SELECT * FROM penyebab_kejadian_lv3` + where + ` AND status != 0`

		PK3.logger.Zap.Info(query)
		rows, err := PK3.dbRaw.DB.Query(query, request.KodeSubKejadian)
		defer rows.Close()

		PK3.logger.Zap.Info("rows =>", rows)
		if err != nil {
			return responses, err
		}

		response := models.PenyebabKejadianLv3Response{}
		for rows.Next() {
			_ = rows.Scan(
				&response.ID,
				&response.KodeSubKejadian,
				&response.KodePenyebabKejadianLv3,
				&response.PenyebabKejadianLv3,
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

// GetKejadianByIDlv1 implements PenyebabKejadianLv3Definition
func (PK3 PenyebabKejadianLv3Repository) GetKejadianByIDlv1(request *models.KodePenyebabKejadian) (responses []models.PenyebabKejadianLv3Response, err error) {
	if request.KodePenyebabKejadian != "" {
		where := " WHERE pkl.kode_kejadian = ?"

		query := `SELECT
					pkl3.id 'id',
					pkl3.kode_sub_kejadian 'kode_sub_kejadian',
					pkl3.kode_penyebab_kejadian_lv3 'kode_penyebab_kejadian_lv3',
					pkl3.penyebab_kejadian_lv3 'penyebab_kejadian_lv3',
					pkl3.deskripsi 'deskripsi',
					pkl3.status 'status',
					pkl3.created_at 'created_at',
					pkl3.updated_at 'updated_at' 
				FROM penyebab_kejadian_lv3 pkl3 
				INNER JOIN penyebab_kejadian_lv2 pkl2 ON pkl2.kode_sub_kejadian = pkl3.kode_sub_kejadian 
				INNER JOIN penyebab_kejadian_lv1 pkl ON pkl2.kode_kejadian = pkl.kode_kejadian` + where + ` AND pkl3.status != 0`

		PK3.logger.Zap.Info(query)
		rows, err := PK3.dbRaw.DB.Query(query, request.KodePenyebabKejadian)
		defer rows.Close()

		PK3.logger.Zap.Info("rows =>", rows)
		if err != nil {
			return responses, err
		}

		response := models.PenyebabKejadianLv3Response{}
		for rows.Next() {
			_ = rows.Scan(
				&response.ID,
				&response.KodeSubKejadian,
				&response.KodePenyebabKejadianLv3,
				&response.PenyebabKejadianLv3,
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

// GetSubKejadian implements PenyebabKejadianLv3Definition.
func (PK3 PenyebabKejadianLv3Repository) GetSubKejadian(id int64) (responses []models.PenyebabKejadianLv3Response, err error) {
	query := PK3.db.DB.Table(`penyebab_kejadian_lv3 pkl3`).
		Select(`
			pkl3.id 'id',
			pkl3.kode_sub_kejadian 'kode_sub_kejadian',
			pkl3.kode_penyebab_kejadian_lv3 'kode_penyebab_kejadian_lv3',
			pkl3.penyebab_kejadian_lv3 'penyebab_kejadian_lv3',
			pkl3.deskripsi 'deskripsi',
			pkl3.status 'status',
			pkl3.created_at 'created_at',
			pkl3.updated_at 'updated_at' 
		`).
		Joins(`INNER JOIN penyebab_kejadian_lv2 pkl2 ON pkl2.kode_sub_kejadian = pkl3.kode_sub_kejadian`).
		Joins(`INNER JOIN penyebab_kejadian_lv1 pkl ON pkl2.kode_kejadian = pkl.kode_kejadian`).
		Where(`pkl3.status != 0`).
		Where(`pkl.id = ?`, id)

	err = query.Scan(&responses).Error

	return responses, err
}
