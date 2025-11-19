package briefing

import (
	"riskmanagement/lib"
	models "riskmanagement/models/briefing"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type BriefingMapPesertaDefinition interface {
	WithTrx(trxHandle *gorm.DB) BriefingMapPesertaRepository
	Store(requests *models.BriefingMapPeserta, tx *gorm.DB) (responses *models.BriefingMapPeserta, err error)
	GetByIDBriefing(id int64) (responses []models.BriefingMapPesertaResponse, err error)
	GetPesertaReport(id int64) (responses []models.PesertaReportResponse, err error)
	DeleteByID(id int64, tx *gorm.DB) (err error)
}

type BriefingMapPesertaRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewBriefingMapPesertaRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) BriefingMapPesertaDefinition {
	return BriefingMapPesertaRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// DeleteByID implements BriefingMapPesertaDefinition
func (repo BriefingMapPesertaRepository) DeleteByID(id int64, tx *gorm.DB) (err error) {
	return tx.Where("id = ?", id).Delete(&models.BriefingMapPeserta{}).Error
}

// GetByIDBriefing implements BriefingMapPesertaDefinition
func (repo BriefingMapPesertaRepository) GetByIDBriefing(id int64) (responses []models.BriefingMapPesertaResponse, err error) {
	rows, err := repo.db.DB.Raw(`SELECT * FROM briefing_map_peserta WHERE id_briefing = ?`, id).Rows()

	defer rows.Close()
	var peserta models.BriefingMapPesertaResponse

	for rows.Next() {
		repo.db.DB.ScanRows(rows, &peserta)
		responses = append(responses, peserta)
	}

	return responses, err
}

// Store implements BriefingMapPesertaDefinition
func (repo BriefingMapPesertaRepository) Store(requests *models.BriefingMapPeserta, tx *gorm.DB) (responses *models.BriefingMapPeserta, err error) {
	return requests, tx.Save(&requests).Error
}

// WithTrx implements BriefingMapPesertaDefinition
func (repo BriefingMapPesertaRepository) WithTrx(trxHandle *gorm.DB) BriefingMapPesertaRepository {
	if trxHandle == nil {
		repo.logger.Zap.Error("transaction Database not found in gin context")
		return repo
	}

	repo.db.DB = trxHandle
	return repo
}

// GetPesertaReport implements BriefingMapPesertaDefinition
func (repo BriefingMapPesertaRepository) GetPesertaReport(id int64) (responses []models.PesertaReportResponse, err error) {
	rows, err := repo.db.DB.Raw(`SELECT nama_peserta 'peserta' FROM briefing_map_peserta WHERE id_briefing = ?`, id).Rows()

	defer rows.Close()
	var peserta models.PesertaReportResponse

	for rows.Next() {
		repo.db.DB.ScanRows(rows, &peserta)
		responses = append(responses, peserta)
	}

	return responses, err
}
