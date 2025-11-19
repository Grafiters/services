package megaproses

import (
	"math"
	"riskmanagement/lib"
	models "riskmanagement/models/megaproses"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type MegaProsesDefinition interface {
	GetAll() (responses []models.MegaProsesResponse, err error)
	GetAllWithPaginate(request *models.Paginate) (responses []models.MegaProsesResponse, totalRows int, totalData int, err error)
	GetOne(id int64) (responses models.MegaProsesResponse, err error)
	Store(request *models.MegaProsesRequest) (responses bool, err error)
	Update(request *models.MegaProsesRequest) (responses bool, err error)
	Delete(id int64) (err error)
	GetKodeMegaProses() (responses []models.KodeMegaProses, err error)
	WithTrx(trxHandle *gorm.DB) MegaprosesRepository
}

type MegaprosesRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewMegaprosesRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) MegaProsesDefinition {
	return MegaprosesRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: 0,
	}
}

// Delete implements MegaprosesDefinition
func (megaproses MegaprosesRepository) Delete(id int64) (err error) {
	return megaproses.db.DB.Where("id = ?", id).Delete(&models.MegaProsesResponse{}).Error
}

// GetAll implements MegaprosesDefinition
func (megaproses MegaprosesRepository) GetAll() (responses []models.MegaProsesResponse, err error) {
	return responses, megaproses.db.DB.Where("status = ?", 1).Find(&responses).Error
}

// GetAllWithPaginate implements MegaProsesDefinition
func (mega MegaprosesRepository) GetAllWithPaginate(request *models.Paginate) (responses []models.MegaProsesResponse, totalRows int, totalData int, err error) {
	rows, err := mega.db.DB.Raw(`
		SELECT
			mp.id 'id',
			mp.kode_mega_proses 'kode_mega_proses',
			mp.mega_proses 'mega_proses',
			mp.deskripsi 'deskripsi',
			mp.status 'status',
			mp.created_at 'created_at',
			mp.updated_at 'updated_at'
		FROM mega_proses mp ORDER BY mp.id ASC LIMIT ? OFFSET ?
	`, request.Limit, request.Offset).Rows()
	if err != nil {
		return responses, totalRows, totalData, err
	}

	defer rows.Close()

	var megaProses models.MegaProsesResponse

	for rows.Next() {
		mega.db.DB.ScanRows(rows, &megaProses)
		responses = append(responses, megaProses)
	}

	paginateQuery := `SELECT COUNT(*) FROM mega_proses`
	err = mega.dbRaw.DB.QueryRow(paginateQuery).Scan(&totalData)
	if err != nil {
		return responses, totalRows, totalData, err
	}

	if totalData > 0 {
		totalRows = int(math.Ceil(float64(totalData) / float64(request.Limit)))
	}
	return responses, totalRows, totalData, err
}

// GetKodeMegaProses implements MegaprosesDefinition
func (megaproses MegaprosesRepository) GetKodeMegaProses() (responses []models.KodeMegaProses, err error) {
	query := `SELECT (kode_mega_proses + 1) 'kode_mega_proses' FROM mega_proses ORDER BY id DESC LIMIT 1`

	megaproses.logger.Zap.Info(query)
	rows, err := megaproses.dbRaw.DB.Query(query)
	defer rows.Close()

	megaproses.logger.Zap.Info("rows ", rows)
	if err != nil {
		return responses, err
	}

	response := models.KodeMegaProses{}
	for rows.Next() {
		_ = rows.Scan(
			&response.KodeMegaProses,
		)

		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, err
	}

	return responses, err
}

// GetOne implements MegaprosesDefinition
func (megaproses MegaprosesRepository) GetOne(id int64) (responses models.MegaProsesResponse, err error) {
	return responses, megaproses.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements MegaprosesDefinition
func (megaproses MegaprosesRepository) Store(request *models.MegaProsesRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	err = megaproses.db.DB.Save(&models.MegaProsesRequest{
		KodeMegaProses: request.KodeMegaProses,
		MegaProses:     request.MegaProses,
		Deskripsi:      request.Deskripsi,
		Status:         request.Status,
		CreatedAt:      &timeNow,
	}).Error
	if err != nil {
		return false, err
	}

	return true, err
}

// Update implements MegaprosesDefinition
func (megaproses MegaprosesRepository) Update(request *models.MegaProsesRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	err = megaproses.db.DB.Save(&models.MegaProsesRequest{
		ID:             request.ID,
		KodeMegaProses: request.KodeMegaProses,
		MegaProses:     request.MegaProses,
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

// WithTrx implements MegaprosesDefinition
func (megaproses MegaprosesRepository) WithTrx(trxHandle *gorm.DB) MegaprosesRepository {
	if trxHandle == nil {
		megaproses.logger.Zap.Error("transaction Database not found in gin context")
		return megaproses
	}

	megaproses.db.DB = trxHandle
	return megaproses
}
