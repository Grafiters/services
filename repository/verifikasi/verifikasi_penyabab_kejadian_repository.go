package verifikasi

import (
	"riskmanagement/lib"
	models "riskmanagement/models/verifikasi"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type VerifikasiPenyababKejadianDefinition interface {
	GetData(id int64) (responses []models.VerifikasiPenyababKejadianResponse, err error)
	Store(request *models.VerifikasiPenyababKejadian, tx *gorm.DB) (responses *models.VerifikasiPenyababKejadian, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) VerifikasiPenyababKejadianRepository
	GetDataDetail(id int64) (responses []models.VerifikasiPenyababKejadianDetailResponse, err error)
}

type VerifikasiPenyababKejadianRepository struct {
	db     lib.Database
	logger logger.Logger
}

func NewVerifikasiPenyababKejadianRepository(
	db lib.Database,
	logger logger.Logger,
) VerifikasiPenyababKejadianDefinition {
	return VerifikasiPenyababKejadianRepository{
		db:     db,
		logger: logger,
	}
}

func (v VerifikasiPenyababKejadianRepository) WithTrx(trxHandle *gorm.DB) VerifikasiPenyababKejadianRepository {
	if trxHandle == nil {
		v.logger.Zap.Error("transaction Database not found in gin context.")
		return v
	}

	v.db.DB = trxHandle

	return v
}

// Delete implements VerifikasiPenyababKejadianDefinition.
func (v VerifikasiPenyababKejadianRepository) Delete(id int64) (err error) {
	return v.db.DB.Where("id = ?", id).Delete(&models.VerifikasiPenyababKejadianResponse{}).Error
}

// GetData implements VerifikasiPenyababKejadianDefinition.
func (v VerifikasiPenyababKejadianRepository) GetData(id int64) (responses []models.VerifikasiPenyababKejadianResponse, err error) {
	db := v.db.DB.Table(`verifikasi_penyebab_kejadian`).Where(`verifikasi_id = ?`, id)

	err = db.Scan(&responses).Error

	return responses, err
}

// Store implements VerifikasiPenyababKejadianDefinition.
func (v VerifikasiPenyababKejadianRepository) Store(request *models.VerifikasiPenyababKejadian, tx *gorm.DB) (responses *models.VerifikasiPenyababKejadian, err error) {
	return request, tx.Save(&request).Error
}

func (v VerifikasiPenyababKejadianRepository) GetDataDetail(id int64) (responses []models.VerifikasiPenyababKejadianDetailResponse, err error) {
	db := v.db.DB.Table(`verifikasi_penyebab_kejadian vpk`).
		Select(`
			vpk.id,
			vpk.verifikasi_id,
			pkl.penyebab_kejadian 'penyebab_kejadian',
			pkl2.penyebab_kejadian_lv3 'sub_penyebab_kejadian'`).
		Joins(`LEFT JOIN penyebab_kejadian_lv1 pkl ON pkl.id = vpk.id_penyebab_kejadian`).
		Joins(`LEFT JOIN penyebab_kejadian_lv3 pkl2 ON pkl2.id = vpk.id_sub_penyebab_kejadian`).
		Where(`verifikasi_id = ?`, id)

	err = db.Scan(&responses).Error

	return responses, err
}
