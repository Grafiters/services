package verifikasi

import (
	"riskmanagement/lib"
	models "riskmanagement/models/verifikasi"
	"time"

	"gorm.io/gorm"

	"gitlab.com/golang-package-library/logger"
)

type VerifikasiAnomaliDataKRIDDefinition interface {
	GetOneByVerifikasi(id int64) (responses []models.VerifikasiAnomaliDataKRIDResponses, err error)
	Store(request *models.VerifikasiAnomaliDataKRID, tx *gorm.DB) (responses *models.VerifikasiAnomaliDataKRID, err error)
	Delete(id int64, tx *gorm.DB) (err error)
	WithTrx(trxHandle *gorm.DB) VerifikasiAnomaliDataKRIDRepository
}

type VerifikasiAnomaliDataKRIDRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewVerifikasiAnomaliDataKRIDRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) VerifikasiAnomaliDataKRIDDefinition {
	return VerifikasiAnomaliDataKRIDRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// Delete implements VerifikasiAnomaliDataKRIDDefinition
func (VerifikasiAnomaliDataKRIDRepository) Delete(id int64, tx *gorm.DB) (err error) {
	panic("unimplemented")
}

// GetOneByVerifikasi implements VerifikasiAnomaliDataKRIDDefinition
func (repo VerifikasiAnomaliDataKRIDRepository) GetOneByVerifikasi(id int64) (responses []models.VerifikasiAnomaliDataKRIDResponses, err error) {
	rows, err := repo.db.DB.Raw(`SELECT * FROM verifikasi_data_anomali_krid where verifikasi_id = ?`, id).Rows()

	defer rows.Close()
	var dataAnomali models.VerifikasiAnomaliDataKRIDResponses

	for rows.Next() {
		repo.db.DB.ScanRows(rows, &dataAnomali)
		responses = append(responses, dataAnomali)
	}

	return responses, err
}

// Store implements VerifikasiAnomaliDataKRIDDefinition
func (repo VerifikasiAnomaliDataKRIDRepository) Store(request *models.VerifikasiAnomaliDataKRID, tx *gorm.DB) (responses *models.VerifikasiAnomaliDataKRID, err error) {
	return request, tx.Save(&request).Error
}

// WithTrx implements VerifikasiAnomaliDataKRIDDefinition
func (repo VerifikasiAnomaliDataKRIDRepository) WithTrx(trxHandle *gorm.DB) VerifikasiAnomaliDataKRIDRepository {
	if trxHandle == nil {
		repo.logger.Zap.Error("transaction Database not found in gin context.")
		return repo
	}

	repo.db.DB = trxHandle
	return repo
}
