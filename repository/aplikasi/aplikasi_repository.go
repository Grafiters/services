package aplikasi

import (
	"riskmanagement/lib"
	models "riskmanagement/models/aplikasi"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type AplikasiDefinition interface {
	GetAll() (responses []models.AplikasiResponse, err error)
	GetOne(id int64) (responses models.AplikasiResponse, err error)
	Store(request *models.AplikasiRequest) (response bool, err error)
	Update(requests *models.AplikasiRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) AplikasiRepository
}

type AplikasiRepository struct {
	db      lib.Database
	dbRaw   lib.Database
	logger  logger.Logger
	timeout time.Duration
}

func NewAplikasiRepository(
	db lib.Database,
	dbRaw lib.Database,
	logger logger.Logger,
) AplikasiDefinition {
	return AplikasiRepository{
		db:      db,
		dbRaw:   db,
		logger:  logger,
		timeout: 0,
	}
}

// Delete implements AplikasiDefinition
func (aplikasi AplikasiRepository) Delete(id int64) (err error) {
	return aplikasi.db.DB.Where("id = ?", id).Delete(&models.AplikasiResponse{}).Error
}

// GetAll implements AplikasiDefinition
func (aplikasi AplikasiRepository) GetAll() (responses []models.AplikasiResponse, err error) {
	return responses, aplikasi.db.DB.Find(&responses).Error
}

// GetOne implements AplikasiDefinition
func (aplikasi AplikasiRepository) GetOne(id int64) (responses models.AplikasiResponse, err error) {
	return responses, aplikasi.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements AplikasiDefinition
func (aplikasi AplikasiRepository) Store(request *models.AplikasiRequest) (response bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	err = aplikasi.db.DB.Save(&models.AplikasiRequest{
		Kode:            request.Kode,
		ApplicationName: request.ApplicationName,
		CreatedAt:       &timeNow,
	}).Error
	if err != nil {
		return false, err
	}

	return true, err
}

// Update implements AplikasiDefinition
func (aplikasi AplikasiRepository) Update(requests *models.AplikasiRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	err = aplikasi.db.DB.Save(&models.AplikasiRequest{
		ID:              requests.ID,
		Kode:            requests.Kode,
		ApplicationName: requests.ApplicationName,
		CreatedAt:       requests.CreatedAt,
		UpdatedAt:       &timeNow,
	}).Error
	if err != nil {
		return false, err
	}

	return true, err
}

// WithTrx implements AplikasiDefinition
func (aplikasi AplikasiRepository) WithTrx(trxHandle *gorm.DB) AplikasiRepository {
	if trxHandle == nil {
		aplikasi.logger.Zap.Error("transaction Database not found in gin context")
		return aplikasi
	}

	aplikasi.db.DB = trxHandle
	return aplikasi
}
