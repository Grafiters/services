package majorproses

import (
	"database/sql"
	"fmt"
	"math"
	"riskmanagement/lib"
	models "riskmanagement/models/majorproses"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type MajorProsesDefinition interface {
	GetAll() (responses []models.MajorProsesResponse, err error)
	GetOne(id int64) (responses models.MajorProsesResponse, err error)
	GetAllWithPaginate(request *models.Paginate) (responses []models.MajorProsesResponse, totalRows int, totalData int, err error)
	Store(request *models.MajorProsesRequest) (responses bool, err error)
	Update(request *models.MajorProsesRequest) (responses bool, err error)
	Delete(id int64) (err error)
	GetKodeMajorProses(request models.KodeMegaProses) (responses []models.KodeMajorProses, err error)
	WithTrx(trxHandle *gorm.DB) MajorProsesRepository
	GetMajorByMegaProses(request *models.KodeMegaProses) (responses []models.MajorProsesResponse, err error)
}

type MajorProsesRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewMajorProsesRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) MajorProsesDefinition {
	return MajorProsesRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: 0,
	}
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

// Delete implements MajorprosesDefinition
func (majorproses MajorProsesRepository) Delete(id int64) (err error) {
	return majorproses.db.DB.Where("id = ?", id).Delete(&models.MajorProsesResponse{}).Error
}

// GetAll implements MajorprosesDefinition
func (majorproses MajorProsesRepository) GetAll() (responses []models.MajorProsesResponse, err error) {
	query := `SELECT 
				a.id, 
				a.id_mega_proses, 
				b.mega_proses as mega_proses_name, 
				a.kode_major_proses, 
				a.major_proses, 
				a.deskripsi, a.status, 
				a.created_at, a.updated_at 
			FROM major_proses a JOIN mega_proses b ON a.id_mega_proses = b.kode_mega_proses
			WHERE a.status = 1`

	majorproses.logger.Zap.Info(query)
	rows, err := majorproses.dbRaw.DB.Query(query)
	defer rows.Close()

	majorproses.logger.Zap.Info("rows ", rows)
	if err != nil {
		return responses, err
	}

	response := models.MajorProsesResponse{}
	for rows.Next() {
		_ = rows.Scan(
			&response.ID,
			&response.IDMegaProses,
			&response.MegaProsesName,
			&response.KodeMajorProses,
			&response.MajorProses,
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

	return responses, err
	// return responses, majorproses.db.DB.Find(&responses).Error
}

// GetAllWithPaginate implements MajorProsesDefinition
func (major MajorProsesRepository) GetAllWithPaginate(request *models.Paginate) (responses []models.MajorProsesResponse, totalRows int, totalData int, err error) {
	rows, err := major.db.DB.Raw(`
	SELECT 
		a.id, 
		a.id_mega_proses, 
		b.mega_proses as mega_proses_name, 
		a.kode_major_proses, 
		a.major_proses, 
		a.deskripsi, 
		a.status,
		a.created_at,
		a.updated_at 
	FROM major_proses a 
	LEFT OUTER JOIN mega_proses b ON  a.id_mega_proses = b.kode_mega_proses ORDER BY a.id LIMIT ? OFFSET ?
	`, request.Limit, request.Offset).Rows()
	if err != nil {
		return responses, totalRows, totalData, err
	}

	defer rows.Close()

	var majorProses models.MajorProsesResponse

	for rows.Next() {
		major.db.DB.ScanRows(rows, &majorProses)
		responses = append(responses, majorProses)
	}

	paginateQuery := `SELECT COUNT(*) FROM major_proses`
	err = major.dbRaw.DB.QueryRow(paginateQuery).Scan(&totalData)
	if err != nil {
		return responses, totalRows, totalData, err
	}

	if totalData > 0 {
		totalRows = int(math.Ceil(float64(totalData) / float64(request.Limit)))
	}
	return responses, totalRows, totalData, err
}

// GetKodeMajorProses implements MajorprosesDefinition
func (majorproses MajorProsesRepository) GetKodeMajorProses(request models.KodeMegaProses) (responses []models.KodeMajorProses, err error) {
	kode := ""
	if request.KodeMegaProses != "" {
		kode += request.KodeMegaProses + "."
	}

	key := fmt.Sprintf("%s%%", kode)

	rowsCheck, err := majorproses.dbRaw.DB.Query(`SELECT COUNT(*) FROM major_proses WHERE kode_major_proses  LIKE ?`, key)
	checkErr(err)

	if checkCount(rowsCheck) == 0 {
		response := models.KodeMajorProses{
			KodeMajorProses: "1",
		}
		responses = append(responses, response)
	} else {
		query := `SELECT (t.kode_major_proses + 1) 'kode_major_proses' FROM (
			SELECT
				CAST(SUBSTRING_INDEX(kode_major_proses ,'.', -1) AS DECIMAL) 'kode_major_proses'
			FROM major_proses
			WHERE kode_major_proses LIKE ?) AS t
		ORDER BY t.kode_major_proses DESC LIMIT 1`

		majorproses.logger.Zap.Info(query)
		rows, err := majorproses.dbRaw.DB.Query(query, key)
		defer rows.Close()

		majorproses.logger.Zap.Info("rows ", rows)
		if err != nil {
			return responses, err
		}

		response := models.KodeMajorProses{}
		for rows.Next() {
			_ = rows.Scan(
				&response.KodeMajorProses,
			)

			responses = append(responses, response)
		}

		if err = rows.Err(); err != nil {
			return responses, err
		}
	}

	return responses, err
}

// GetOne implements MajorprosesDefinition
func (majorproses MajorProsesRepository) GetOne(id int64) (responses models.MajorProsesResponse, err error) {
	return responses, majorproses.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements MajorprosesDefinition
func (majorproses MajorProsesRepository) Store(request *models.MajorProsesRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	err = majorproses.db.DB.Save(&models.MajorProsesRequest{
		IDMegaProses:    request.IDMegaProses,
		KodeMajorProses: request.KodeMajorProses,
		MajorProses:     request.MajorProses,
		Deskripsi:       request.Deskripsi,
		Status:          request.Status,
		CreatedAt:       &timeNow,
	}).Error
	if err != nil {
		return false, err
	}

	return true, err
}

// Update implements MajorprosesDefinition
func (majorproses MajorProsesRepository) Update(request *models.MajorProsesRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	err = majorproses.db.DB.Save(&models.MajorProsesRequest{
		ID:              request.ID,
		IDMegaProses:    request.IDMegaProses,
		KodeMajorProses: request.KodeMajorProses,
		MajorProses:     request.MajorProses,
		Deskripsi:       request.Deskripsi,
		Status:          request.Status,
		CreatedAt:       request.CreatedAt,
		UpdatedAt:       &timeNow,
	}).Error
	if err != nil {
		return false, err
	}

	return true, err
}

// WithTrx implements MajorprosesDefinition
func (majorproses MajorProsesRepository) WithTrx(trxHandle *gorm.DB) MajorProsesRepository {
	if trxHandle == nil {
		majorproses.logger.Zap.Error("transaction Database not found in gin context")
		return majorproses
	}

	majorproses.db.DB = trxHandle
	return majorproses
}

// GetMajorByMegaProses implements MajorProsesDefinition
func (majorproses MajorProsesRepository) GetMajorByMegaProses(request *models.KodeMegaProses) (responses []models.MajorProsesResponse, err error) {
	if request.KodeMegaProses != "" {
		// where := " WHERE id_mega_proses = '" + request.KodeMegaProses + "'"
		where := " WHERE id_mega_proses = ?"

		query := `SELECT * FROM major_proses ` + where + ` AND status != 0`

		majorproses.logger.Zap.Info(query)
		rows, err := majorproses.dbRaw.DB.Query(query, request.KodeMegaProses)
		defer rows.Close()

		majorproses.logger.Zap.Info("rows =>", rows)
		if err != nil {
			return responses, err
		}

		response := models.MajorProsesResponse{}
		for rows.Next() {
			_ = rows.Scan(
				&response.ID,
				&response.IDMegaProses,
				&response.KodeMajorProses,
				&response.MajorProses,
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
