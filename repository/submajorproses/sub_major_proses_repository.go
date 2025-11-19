package submajorproses

import (
	"math"
	"riskmanagement/lib"
	models "riskmanagement/models/submajorproses"
	"time"

	"database/sql"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type SubMajorProsesDefinition interface {
	GetAll() (responses []models.SubMajorProsesResponse, err error)
	GetAllWithPaginate(request *models.Paginate) (responses []models.SubMajorProsesResponse, totalRows int, totalData int, err error)
	GetOne(id int64) (responses models.SubMajorProsesResponse, err error)
	Store(request *models.SubMajorProsesRequest) (responses bool, err error)
	Update(request *models.SubMajorProsesRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) SubMajorProsesRepository
	GetDataByID(request *models.KodeMajor) (responses []models.SubMajorProsesResponse, err error)
}

type SubMajorProsesRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewSubMajorProsesRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) SubMajorProsesDefinition {
	return SubMajorProsesRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: 0,
	}
}

// Delete implements SubMajorProsesDefinition
func (submajorproses SubMajorProsesRepository) Delete(id int64) (err error) {
	return submajorproses.db.DB.Where("id = ?", id).Delete(&models.SubMajorProsesResponse{}).Error
}

// GetAll implements SubMajorProsesDefinition
func (submajorproses SubMajorProsesRepository) GetAll() (responses []models.SubMajorProsesResponse, err error) {
	query := "SELECT a.id, a.id_major_proses, b.major_proses, a.kode_sub_major_proses, a.sub_major_proses, a.deskripsi, a.status, a.created_at, a.updated_at FROM sub_major_proses a JOIN major_proses b ON a.id_major_proses = b.kode_major_proses"

	submajorproses.logger.Zap.Info(query)
	rows, err := submajorproses.dbRaw.DB.Query(query)
	defer rows.Close()

	submajorproses.logger.Zap.Info("rows ", rows)
	if err != nil {
		return responses, err
	}

	response := models.SubMajorProsesResponse{}
	for rows.Next() {
		_ = rows.Scan(
			&response.ID,
			&response.IDMajorProses,
			&response.MajorProses,
			&response.KodeSubMajorProses,
			&response.SubMajorProses,
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
	// return responses, submajorproses.db.DB.Find(&responses).Error
}

// GetAllWithPaginate implements SubMajorProsesDefinition
func (subMajor SubMajorProsesRepository) GetAllWithPaginate(request *models.Paginate) (responses []models.SubMajorProsesResponse, totalRows int, totalData int, err error) {
	rows, err := subMajor.db.DB.Raw(`
		SELECT
			a.id,
			a.id_major_proses, 
			b.major_proses, 
			a.kode_sub_major_proses, 
			a.sub_major_proses, 
			a.deskripsi, 
			a.status, 
			a.created_at, 
			a.updated_at 
		FROM sub_major_proses a 
		LEFT OUTER JOIN major_proses b ON a.id_major_proses = b.kode_major_proses ORDER BY a.id LIMIT ? OFFSET ?
	`, request.Limit, request.Offset).Rows()
	if err != nil {
		return responses, totalRows, totalData, err
	}
	defer rows.Close()

	var subMajorProses models.SubMajorProsesResponse

	for rows.Next() {
		subMajor.db.DB.ScanRows(rows, &subMajorProses)
		responses = append(responses, subMajorProses)
	}

	paginateQuery := `SELECT COUNT(*) FROM sub_major_proses`
	err = subMajor.dbRaw.DB.QueryRow(paginateQuery).Scan(&totalData)
	if err != nil {
		return responses, totalRows, totalData, err
	}

	if totalData > 0 {
		totalRows = int(math.Ceil((float64(totalData)) / float64(request.Limit)))
	}

	return responses, totalRows, totalData, err
}

// GetOne implements SubMajorProsesDefinition
func (submajorproses SubMajorProsesRepository) GetOne(id int64) (responses models.SubMajorProsesResponse, err error) {
	return responses, submajorproses.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements SubMajorProsesDefinition
func (submajorproses SubMajorProsesRepository) Store(request *models.SubMajorProsesRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	rowsCheckCode, err := submajorproses.dbRaw.DB.Query("SELECT COUNT(*) as count FROM sub_major_proses WHERE kode_sub_major_proses = ?", request.KodeSubMajorProses)
	// fmt.Println("Total count:",)
	checkErr(err)

	if checkCount(rowsCheckCode) == 0 {
		err = submajorproses.db.DB.Save(&models.SubMajorProsesRequest{
			IDMajorProses:      request.IDMajorProses,
			KodeSubMajorProses: request.KodeSubMajorProses,
			SubMajorProses:     request.SubMajorProses,
			Deskripsi:          request.Deskripsi,
			Status:             request.Status,
			CreatedAt:          &timeNow,
			UpdatedAt:          &timeNow,
		}).Error

		return true, err
	}

	return false, err
}

// Update implements SubMajorProsesDefinition
func (submajorproses SubMajorProsesRepository) Update(request *models.SubMajorProsesRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	err = submajorproses.db.DB.Save(&models.SubMajorProsesRequest{
		ID:                 request.ID,
		IDMajorProses:      request.IDMajorProses,
		KodeSubMajorProses: request.KodeSubMajorProses,
		SubMajorProses:     request.SubMajorProses,
		Deskripsi:          request.Deskripsi,
		Status:             request.Status,
		CreatedAt:          request.CreatedAt,
		UpdatedAt:          &timeNow,
	}).Error
	if err != nil {
		return false, err
	}

	return true, err
}

// WithTrx implements SubMajorProsesDefinition
func (submajorproses SubMajorProsesRepository) WithTrx(trxHandle *gorm.DB) SubMajorProsesRepository {
	if trxHandle == nil {
		submajorproses.logger.Zap.Error("transaction Database not found in gin context")
		return submajorproses
	}

	submajorproses.db.DB = trxHandle
	return submajorproses
}

// GetDataByID implements SubMajorProsesDefinition
func (submajorproses SubMajorProsesRepository) GetDataByID(request *models.KodeMajor) (responses []models.SubMajorProsesResponse, err error) {
	if request.KodeMajor != "" {
		where := " WHERE id_major_proses = ?"

		query := `SELECT * FROM sub_major_proses ` + where + ` AND status != 0`

		submajorproses.logger.Zap.Info(query)
		rows, err := submajorproses.dbRaw.DB.Query(query, request.KodeMajor)
		defer rows.Close()

		submajorproses.logger.Zap.Info("rows =>", rows)

		if err != nil {
			return responses, err
		}

		response := models.SubMajorProsesResponse{}
		for rows.Next() {
			_ = rows.Scan(
				&response.ID,
				&response.IDMajorProses,
				&response.KodeSubMajorProses,
				&response.SubMajorProses,
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
