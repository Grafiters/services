package verifikasi

import (
	"riskmanagement/lib"
	models "riskmanagement/models/verifikasi"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type VerifikasiAnomaliDefinition interface {
	GetAll() (responses []models.VerifikasiAnomaliDataResponse, err error)
	GetOne(id int64) (responses models.VerifikasiAnomaliDataResponse, err error)
	GetOneByVerifikasi(id int64) (responses []models.VerifikasiAnomaliDataResponses, err error)
	Store(request *models.VerifikasiAnomaliData, tx *gorm.DB) (responses *models.VerifikasiAnomaliData, err error)
	Update(request *models.VerifikasiAnomaliData, tx *gorm.DB) (responses bool, err error)
	Delete(id int64) (err error)
	DeleteAnomaliByID(id int64, tx *gorm.DB) (err error)
	WithTrx(trxHandle *gorm.DB) VerifikasiAnomaliRepository
}

type VerifikasiAnomaliRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewVerifikasiAnomaliRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) VerifikasiAnomaliDefinition {
	return VerifikasiAnomaliRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// Delete implements VerifikasiAnomaliDefinition
func (anomaliData VerifikasiAnomaliRepository) Delete(id int64) (err error) {
	return anomaliData.db.DB.Where("id = ?", id).Delete(&models.VerifikasiAnomaliDataResponse{}).Error
}

// DeleteAnomaliByID implements VerifikasiAnomaliDefinition
func (anomaliData VerifikasiAnomaliRepository) DeleteAnomaliByID(id int64, tx *gorm.DB) (err error) {
	return tx.Where("verifikasi_id = ?", id).Delete(&models.VerifikasiAnomaliDataResponse{}).Error
}

// GetAll implements VerifikasiAnomaliDefinition
func (anomaliData VerifikasiAnomaliRepository) GetAll() (responses []models.VerifikasiAnomaliDataResponse, err error) {
	return responses, anomaliData.db.DB.Find(&responses).Error
}

// GetOne implements VerifikasiAnomaliDefinition
func (anomaliData VerifikasiAnomaliRepository) GetOne(id int64) (responses models.VerifikasiAnomaliDataResponse, err error) {
	return responses, anomaliData.db.DB.Where("id = ?", id).Find(&responses).Error
}

// GetOneByVerifikasi implements VerifikasiAnomaliDefinition
func (anomaliData VerifikasiAnomaliRepository) GetOneByVerifikasi(id int64) (responses []models.VerifikasiAnomaliDataResponses, err error) {
	rows, err := anomaliData.db.DB.Raw(`
		SELECT vda.*
		FROM verifikasi_data_anomali vda WHERE vda.verifikasi_id = ?`, id).Rows()

	defer rows.Close()
	var anomali models.VerifikasiAnomaliDataResponses

	for rows.Next() {
		anomaliData.db.DB.ScanRows(rows, &anomali)
		responses = append(responses, anomali)
	}

	return responses, err
}

// Store implements VerifikasiAnomaliDefinition
func (anomaliData VerifikasiAnomaliRepository) Store(request *models.VerifikasiAnomaliData, tx *gorm.DB) (responses *models.VerifikasiAnomaliData, err error) {
	return request, tx.Save(&request).Error
}

// Update implements VerifikasiAnomaliDefinition
func (anomaliData VerifikasiAnomaliRepository) Update(request *models.VerifikasiAnomaliData, tx *gorm.DB) (responses bool, err error) {
	return true, tx.Save(&request).Error
}

// WithTrx implements VerifikasiAnomaliDefinition
func (anomaliData VerifikasiAnomaliRepository) WithTrx(trxHandle *gorm.DB) VerifikasiAnomaliRepository {
	if trxHandle == nil {
		anomaliData.logger.Zap.Error("transaction Database not found in gin context.")
		return anomaliData
	}
	anomaliData.db.DB = trxHandle
	return anomaliData
}
