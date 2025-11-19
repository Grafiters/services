package verifikasi

import (
	"riskmanagement/lib"
	models "riskmanagement/models/verifikasi"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type VerifikasiUsulanPerbaikanDefinition interface {
	GetData(id int64) (responses []models.VerifikasiUsulanPerbaikanResponse, err error)
	Store(request *models.VerifikasiUsulanPerbaikan, tx *gorm.DB) (responses *models.VerifikasiUsulanPerbaikan, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) VerifikasiUsulanPerbaikanRepository
}

type VerifikasiUsulanPerbaikanRepository struct {
	db     lib.Database
	logger logger.Logger
}

func NewVerifikasiUsulanPerbaikanRepository(
	db lib.Database,
	logger logger.Logger,
) VerifikasiUsulanPerbaikanDefinition {
	return VerifikasiUsulanPerbaikanRepository{
		db:     db,
		logger: logger,
	}
}

func (v VerifikasiUsulanPerbaikanRepository) WithTrx(trxHandle *gorm.DB) VerifikasiUsulanPerbaikanRepository {
	if trxHandle == nil {
		v.logger.Zap.Error("transaction Database not found in gin context.")
		return v
	}

	v.db.DB = trxHandle

	return v
}

// Delete implements VerifikasiUsulanPerbaikanDefinition.
func (v VerifikasiUsulanPerbaikanRepository) Delete(id int64) (err error) {
	return v.db.DB.Where("id = ?", id).Delete(&models.VerifikasiUsulanPerbaikanResponse{}).Error
}

// GetData implements VerifikasiUsulanPerbaikanDefinition.
func (v VerifikasiUsulanPerbaikanRepository) GetData(id int64) (responses []models.VerifikasiUsulanPerbaikanResponse, err error) {
	db := v.db.DB.Table(`verifikasi_usulan_perbaikan`).
		Select(`id, verifikasi_id, usulan, deskripsi, aplikasi`).
		Where(`verifikasi_id = ?`, id)

	err = db.Scan(&responses).Error

	return responses, err
}

// Store implements VerifikasiUsulanPerbaikanDefinition.
func (v VerifikasiUsulanPerbaikanRepository) Store(request *models.VerifikasiUsulanPerbaikan, tx *gorm.DB) (responses *models.VerifikasiUsulanPerbaikan, err error) {
	return request, tx.Save(&request).Error
}
