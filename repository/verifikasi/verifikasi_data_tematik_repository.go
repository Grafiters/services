package verifikasi

import (
	"riskmanagement/lib"
	models "riskmanagement/models/verifikasi"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type VerifikasiDataTematikDefinition interface {
	WithTrx(trxHandle *gorm.DB) VerifikasiDataTematikRepository
	GetDataTematik(id int64) (responses []models.VerifikasiDataTematikResponse, err error)
	Store(request *models.VerifikasiDataTematik, tx *gorm.DB) (responses *models.VerifikasiDataTematik, err error)
	Delete(id int64, tx *gorm.DB) (err error)
}

type VerifikasiDataTematikRepository struct {
	db     lib.Database
	logger logger.Logger
	// timeout time.Duration
}

func NewVerifikasiDataTematikRepository(
	db lib.Database,
	logger logger.Logger,
) VerifikasiDataTematikDefinition {
	return VerifikasiDataTematikRepository{
		db:     db,
		logger: logger,
	}
}

// Delete implements VerifikasiDataTematikDefinition.
func (v VerifikasiDataTematikRepository) Delete(id int64, tx *gorm.DB) (err error) {
	panic("unimplemented")
}

// GetDataTematik implements VerifikasiDataTematikDefinition.
func (v VerifikasiDataTematikRepository) GetDataTematik(id int64) (responses []models.VerifikasiDataTematikResponse, err error) {
	db := v.db.DB.Model(&responses).Where(`verifikasi_id = ?`, id)

	err = db.Scan(&responses).Error

	return responses, err
}

// Store implements VerifikasiDataTematikDefinition.
func (v VerifikasiDataTematikRepository) Store(request *models.VerifikasiDataTematik, tx *gorm.DB) (responses *models.VerifikasiDataTematik, err error) {
	return request, tx.Save(&request).Error
}

// WithTrx implements VerifikasiDataTematikDefinition.
func (repo VerifikasiDataTematikRepository) WithTrx(trxHandle *gorm.DB) VerifikasiDataTematikRepository {
	if trxHandle == nil {
		repo.logger.Zap.Error("transaction Database not found in gin context.")
		return repo
	}

	repo.db.DB = trxHandle
	return repo
}
