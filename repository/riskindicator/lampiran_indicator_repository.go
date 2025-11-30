package riskindicator

import (
	"riskmanagement/lib"
	models "riskmanagement/models/riskindicator"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type LampiranIndicatorDefinition interface {
	GetAll() (responses []models.LampiranIndicatorResponse, err error)
	GetOne(id int64) (responses models.LampiranIndicatorResponse, err error)
	GetOneFileByID(id int64) (responses []models.LampiranIndicatorResponse, err error)
	Store(request *models.LampiranIndicator, tx *gorm.DB) (responses *models.LampiranIndicator, err error)
	Update(request *models.LampiranIndicator, tx *gorm.DB) (responses bool, err error)
	Delete(id int64) (err error)
	DeleteFilesByIndicator(id int64, tx *gorm.DB) (err error)
	WithTrx(trxHandle *gorm.DB) LampiranIndicatorRepository
}

type LampiranIndicatorRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewLampiranIndicatorRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) LampiranIndicatorDefinition {
	return LampiranIndicatorRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: 0,
	}
}

// Delete implements LampiranIndicatorDefinition
func (LI LampiranIndicatorRepository) Delete(id int64) (err error) {
	return LI.db.DB.Where("id = ?", id).Delete(&models.LampiranIndicatorResponse{}).Error
}

// DeleteFilesByID implements LampiranIndicatorDefinition
func (LI LampiranIndicatorRepository) DeleteFilesByIndicator(id int64, tx *gorm.DB) (err error) {
	return tx.Where("id_indicator  = ?", id).Delete(&models.LampiranIndicatorResponse{}).Error
}

// GetAll implements LampiranIndicatorDefinition
func (LI LampiranIndicatorRepository) GetAll() (responses []models.LampiranIndicatorResponse, err error) {
	return responses, LI.db.DB.Find(&responses).Error
}

// GetOne implements LampiranIndicatorDefinition
func (LI LampiranIndicatorRepository) GetOne(id int64) (responses models.LampiranIndicatorResponse, err error) {
	return responses, LI.db.DB.Where("id = ?", id).Find(&responses).Error
}
func (LI LampiranIndicatorRepository) GetOneFileByID(id int64) (responses []models.LampiranIndicatorResponse, err error) {
	// Jalankan query
	rows, err := LI.db.DB.Raw(`
        SELECT
            files.id AS id,
            files.id_indicator AS id_indicator,
            files.nama_lampiran AS nama_lampiran,
            files.nomor_lampiran AS nomor_lampiran,
            files.jenis_file AS jenis_file,
            files.path AS path,
            files.filename AS filename
        FROM risk_indicator_map_files files
        WHERE files.id_indicator = ?
    `, id).Rows()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Loop rows
	for rows.Next() {
		var lampiran models.LampiranIndicatorResponse // declare di dalam loop agar tidak reuse nilai lama

		if err := LI.db.DB.ScanRows(rows, &lampiran); err != nil {
			return nil, err
		}

		responses = append(responses, lampiran)
	}

	// Cek jika ada error di rows iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return responses, nil
}

// Store implements LampiranIndicatorDefinition
func (LI LampiranIndicatorRepository) Store(request *models.LampiranIndicator, tx *gorm.DB) (responses *models.LampiranIndicator, err error) {
	return request, tx.Save(&request).Error
}

// Update implements LampiranIndicatorDefinition
func (LI LampiranIndicatorRepository) Update(request *models.LampiranIndicator, tx *gorm.DB) (responses bool, err error) {
	return true, tx.Save(*request).Error
}

// WithTrx implements LampiranIndicatorDefinition
func (LI LampiranIndicatorRepository) WithTrx(trxHandle *gorm.DB) LampiranIndicatorRepository {
	if trxHandle == nil {
		LI.logger.Zap.Error("transaction Database not found in gin context.")
		return LI
	}

	LI.db.DB = trxHandle
	return LI
}
