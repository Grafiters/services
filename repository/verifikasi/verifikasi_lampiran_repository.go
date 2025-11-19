package verifikasi

import (
	"riskmanagement/lib"
	models "riskmanagement/models/verifikasi"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type VerifikasiFilesDefinition interface {
	GetAll() (responses []models.VerifikasiFilesResponse, err error)
	GetOne(id int64) (responses models.VerifikasiFilesResponse, err error)
	GetOneFileByID(id int64) (responses []models.VerifikasiFilesResponses, err error)
	Store(request *models.VerifikasiFiles, tx *gorm.DB) (responses *models.VerifikasiFiles, err error)
	Update(request *models.VerifikasiFiles, tx *gorm.DB) (responses bool, err error)
	Delete(id int64) (err error)
	DeleteFilesByID(id int64, tx *gorm.DB) (err error)
	WithTrx(trxHandle *gorm.DB) VerifikasiFilesRepository
}

type VerifikasiFilesRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewVerfikasiFilesRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) VerifikasiFilesDefinition {
	return VerifikasiFilesRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// Delete implements VerifikasiFilesDefinition
func (lampiranFiles VerifikasiFilesRepository) Delete(id int64) (err error) {
	return lampiranFiles.db.DB.Where("id = ?", id).Delete(&models.VerifikasiFilesResponse{}).Error
}

// DeleteFilesByID implements VerifikasiFilesDefinition
func (lampiranFiles VerifikasiFilesRepository) DeleteFilesByID(id int64, tx *gorm.DB) (err error) {
	return tx.Where("verifikasi_id = ?", id).Delete(&models.VerifikasiFilesResponse{}).Error
}

// GetAll implements VerifikasiFilesDefinition
func (lampiranFiles VerifikasiFilesRepository) GetAll() (responses []models.VerifikasiFilesResponse, err error) {
	return responses, lampiranFiles.db.DB.Find(&responses).Error
}

// GetOne implements VerifikasiFilesDefinition
func (lampiranFiles VerifikasiFilesRepository) GetOne(id int64) (responses models.VerifikasiFilesResponse, err error) {
	return responses, lampiranFiles.db.DB.Where("id = ?", id).Find(&responses).Error
}

// GetOneFileByID implements VerifikasiFilesDefinition
func (lampiranFiles VerifikasiFilesRepository) GetOneFileByID(id int64) (responses []models.VerifikasiFilesResponses, err error) {
	rows, err := lampiranFiles.db.DB.Raw(`
		SELECT 
			lam.id 'id_lampiran',
			lam.verifikasi_id 'verifikasi_id',
			fl.filename 'filename',
			fl.path 'path',
			fl.extension 'ext',
			fl.size 'size'
		FROM verifikasi_lampiran lam 
		JOIN files fl ON fl.id = lam.files_id
		WHERE lam.verifikasi_id = ?`, id).Rows()

	defer rows.Close()
	var lampiran models.VerifikasiFilesResponses

	for rows.Next() {
		lampiranFiles.db.DB.ScanRows(rows, &lampiran)
		responses = append(responses, lampiran)
	}

	return responses, err
}

// Store implements VerifikasiFilesDefinition
func (lampiranFiles VerifikasiFilesRepository) Store(request *models.VerifikasiFiles, tx *gorm.DB) (responses *models.VerifikasiFiles, err error) {
	return request, tx.Save(&request).Error
}

// Update implements VerifikasiFilesDefinition
func (lampiranFiles VerifikasiFilesRepository) Update(request *models.VerifikasiFiles, tx *gorm.DB) (responses bool, err error) {
	return true, tx.Save(*request).Error
}

// WithTrx implements VerifikasiFilesDefinition
func (lampiranFiles VerifikasiFilesRepository) WithTrx(trxHandle *gorm.DB) VerifikasiFilesRepository {
	if trxHandle == nil {
		lampiranFiles.logger.Zap.Error("transaction Database not found in gin context.")
		return lampiranFiles
	}

	lampiranFiles.db.DB = trxHandle
	return lampiranFiles
}
