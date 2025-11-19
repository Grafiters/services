package coaching

import (
	"riskmanagement/lib"
	models "riskmanagement/models/coaching"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type CoachingMapPesertaDefinition interface {
	WithTrx(trxHandle *gorm.DB) CoachingMapPesertaRepository
	Store(request *models.CoachingMapPeserta, tx *gorm.DB) (responses *models.CoachingMapPeserta, err error)
	GetByIDCoaching(id int64) (responses []models.CoachingMapPesertaResponse, err error)
	DeleteByID(id int64, tx *gorm.DB) (err error)
	GetPesertaReport(id int64) (responses []models.PesertaReportResponse, err error)
}

type CoachingMapPesertaRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewCoachingMapPesertaRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) CoachingMapPesertaDefinition {
	return CoachingMapPesertaRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// DeleteByID implements CoachingMapPesertaDefinition
func (CoachingMapPesertaRepository) DeleteByID(id int64, tx *gorm.DB) (err error) {
	return tx.Where("id = ?", id).Delete(&models.CoachingMapPeserta{}).Error
}

// GetByIDCoaching implements CoachingMapPesertaDefinition
func (repo CoachingMapPesertaRepository) GetByIDCoaching(id int64) (responses []models.CoachingMapPesertaResponse, err error) {
	rows, err := repo.db.DB.Raw(`SELECT * FROM coaching_map_peserta WHERE id_coaching = ?`, id).Rows()

	defer rows.Close()
	var peserta models.CoachingMapPesertaResponse
	for rows.Next() {
		repo.db.DB.ScanRows(rows, &peserta)
		responses = append(responses, peserta)
	}

	return responses, err
}

// Store implements CoachingMapPesertaDefinition
func (repo CoachingMapPesertaRepository) Store(request *models.CoachingMapPeserta, tx *gorm.DB) (responses *models.CoachingMapPeserta, err error) {
	return request, tx.Save(&request).Error

}

// WithTrx implements CoachingMapPesertaDefinition
func (repo CoachingMapPesertaRepository) WithTrx(trxHandle *gorm.DB) CoachingMapPesertaRepository {
	if trxHandle == nil {
		repo.logger.Zap.Error("transaction Database not found in gin context")
		return repo
	}

	repo.db.DB = trxHandle
	return repo
}

// GetPesertaReport implements CoachingMapPesertaDefinition
func (repo CoachingMapPesertaRepository) GetPesertaReport(id int64) (responses []models.PesertaReportResponse, err error) {
	rows, err := repo.db.DB.Raw(`SELECT nama_peserta 'peserta' FROM coaching_map_peserta WHERE id_coaching = ?`, id).Rows()

	defer rows.Close()
	var peserta models.PesertaReportResponse

	for rows.Next() {
		repo.db.DB.ScanRows(rows, &peserta)
		responses = append(responses, peserta)
	}

	return responses, err
}
