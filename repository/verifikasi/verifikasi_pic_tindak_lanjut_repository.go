package verifikasi

import (
	"riskmanagement/lib"
	models "riskmanagement/models/verifikasi"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type VerifikasiPICDefinition interface {
	GetAll() (responses []models.VerifikasiPICTindakLanjutRequest, err error)
	GetOne(id int64) (responses models.VerifikasiPICTindakLanjutResponse, err error)
	GetOneByPIC(id int64) (responses []models.VerifikasiPICTindakLanjutResponses, err error)
	Store(request *models.VerifikasiPICTindakLanjut, tx *gorm.DB) (responses *models.VerifikasiPICTindakLanjut, err error)
	Update(request *models.VerifikasiPICTindakLanjut, tx *gorm.DB) (responses bool, err error)
	Delete(id int64) (err error)
	DeletePICByID(id int64, tx *gorm.DB) (err error)
	WithTrx(trxHandle *gorm.DB) VerifikasiPICRepository
}

type VerifikasiPICRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewVerifikasiPICRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) VerifikasiPICDefinition {
	return VerifikasiPICRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// Delete implements VerifikasiPICDefinition
func (picTindakLanjut VerifikasiPICRepository) Delete(id int64) (err error) {
	return picTindakLanjut.db.DB.Where("id = ?", id).Delete(&models.VerifikasiPICTindakLanjutResponse{}).Error
}

// DeletePICByID implements VerifikasiPICDefinition
func (picTindakLanjut VerifikasiPICRepository) DeletePICByID(id int64, tx *gorm.DB) (err error) {
	return tx.Where("verifikasi_id = ?", id).Delete(&models.VerifikasiPICTindakLanjutResponse{}).Error
}

// GetAll implements VerifikasiPICDefinition
func (picTindakLanjut VerifikasiPICRepository) GetAll() (responses []models.VerifikasiPICTindakLanjutRequest, err error) {
	return responses, picTindakLanjut.db.DB.Find(&responses).Error
}

// GetOne implements VerifikasiPICDefinition
func (picTindakLanjut VerifikasiPICRepository) GetOne(id int64) (responses models.VerifikasiPICTindakLanjutResponse, err error) {
	return responses, picTindakLanjut.db.DB.Where("id = ?", id).Find(&responses).Error
}

// GetOneByPIC implements VerifikasiPICDefinition
func (picTindakLanjut VerifikasiPICRepository) GetOneByPIC(id int64) (responses []models.VerifikasiPICTindakLanjutResponses, err error) {
	rows, err := picTindakLanjut.db.DB.Raw(`
		SELECT vpic.*
		FROM verifikasi_pic_tindak_lanjut vpic WHERE vpic.verifikasi_id = ?`, id).Rows()

	defer rows.Close()
	var verifPIC models.VerifikasiPICTindakLanjutResponses

	for rows.Next() {
		picTindakLanjut.db.DB.ScanRows(rows, &verifPIC)
		responses = append(responses, verifPIC)
	}

	return responses, err
}

// Store implements VerifikasiPICDefinition
func (picTindakLanjut VerifikasiPICRepository) Store(request *models.VerifikasiPICTindakLanjut, tx *gorm.DB) (responses *models.VerifikasiPICTindakLanjut, err error) {
	return request, tx.Save(&request).Error
}

// Update implements VerifikasiPICDefinition
func (picTindakLanjut VerifikasiPICRepository) Update(request *models.VerifikasiPICTindakLanjut, tx *gorm.DB) (responses bool, err error) {
	return true, tx.Save(&request).Error
}

// WithTrx implements VerifikasiPICDefinition
func (picTindakLanjut VerifikasiPICRepository) WithTrx(trxHandle *gorm.DB) VerifikasiPICRepository {
	if trxHandle == nil {
		picTindakLanjut.logger.Zap.Error("transaction Database not found in gin context.")
		return picTindakLanjut
	}

	picTindakLanjut.db.DB = trxHandle
	return picTindakLanjut
}
