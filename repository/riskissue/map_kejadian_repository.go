package riskissue

import (
	"riskmanagement/lib"
	models "riskmanagement/models/riskissue"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type MapKejadianDefinition interface {
	GetAll() (responses []models.MapKejadianResponse, err error)
	GetOne(id int64) (responses []models.MapKejadianResponse, err error)
	Store(request *models.MapKejadian, tx *gorm.DB) (responses *models.MapKejadian, err error)
	Update(request *models.MapKejadian, tx *gorm.DB) (responses bool, err error)
	Delete(id int64) (err error)
	GetOneDataByID(id int64) (responses []models.MapKejadianResponseFinal, err error)
	DeleteDataByID(id int64, tx *gorm.DB) (err error)
	WithTrx(trxHandle *gorm.DB) MapKejadianRepository
}

type MapKejadianRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewMapKejadianRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) MapKejadianDefinition {
	return MapKejadianRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// Delete implements MapKejadianDefinition
func (mp MapKejadianRepository) Delete(id int64) (err error) {
	return mp.db.DB.Where("id = ?", id).Delete(&models.MapKejadianResponseFinal{}).Error
}

// DeleteDataByID implements MapKejadianDefinition
func (mp MapKejadianRepository) DeleteDataByID(id int64, tx *gorm.DB) (err error) {
	return tx.Where("id_risk_issue = ?", id).Delete(&models.MapKejadianResponseFinal{}).Error
}

// GetAll implements MapKejadianDefinition
func (mp MapKejadianRepository) GetAll() (responses []models.MapKejadianResponse, err error) {
	return responses, mp.db.DB.Find(&responses).Error
}

// GetOne implements MapKejadianDefinition
func (mp MapKejadianRepository) GetOne(id int64) (responses []models.MapKejadianResponse, err error) {
	return responses, mp.db.DB.Where("id = ?", id).Find(&responses).Error
}

// GetOneDataByID implements MapKejadianDefinition
func (mp MapKejadianRepository) GetOneDataByID(id int64) (responses []models.MapKejadianResponseFinal, err error) {
	rows, err := mp.db.DB.Raw(`
	SELECT
		mk.id 'id',
		mk.id_risk_issue 'id_risk_issue',
		mk.penyebab_kejadian_lv1 'penyebab_kejadian_lv1',
		pk1.penyebab_kejadian 'penyebab_kejadian_lv1_desc',
		mk.penyebab_kejadian_lv2 'penyebab_kejadian_lv2',
		pk2.kriteria_penyebab_kejadian 'penyebab_kejadian_lv2_desc',
		mk.penyebab_kejadian_lv3 'penyebab_kejadian_lv3',
		pk3.penyebab_kejadian_lv3 'penyebab_kejadian_lv3_desc'
	FROM risk_issue_map_kejadian mk
	LEFT JOIN penyebab_kejadian_lv1 pk1 ON pk1.kode_kejadian = mk.penyebab_kejadian_lv1
	LEFT JOIN penyebab_kejadian_lv2 pk2 ON pk2.kode_sub_kejadian = mk.penyebab_kejadian_lv2
	LEFT JOIN penyebab_kejadian_lv3 pk3 ON pk3.kode_penyebab_kejadian_lv3 = mk.penyebab_kejadian_lv3
	WHERE id_risk_issue = ?`, id).Rows()
	if err != nil {
		return responses, err
	}

	defer rows.Close()
	var DataMp models.MapKejadianResponseFinal

	for rows.Next() {
		mp.db.DB.ScanRows(rows, &DataMp)
		responses = append(responses, DataMp)
	}

	return responses, err
}

// Store implements MapKejadianDefinition
func (mp MapKejadianRepository) Store(request *models.MapKejadian, tx *gorm.DB) (responses *models.MapKejadian, err error) {
	return request, tx.Save(&request).Error
}

// Update implements MapKejadianDefinition
func (mp MapKejadianRepository) Update(request *models.MapKejadian, tx *gorm.DB) (responses bool, err error) {
	return true, tx.Save(&request).Error
}

// WithTrx implements MapKejadianDefinition
func (mp MapKejadianRepository) WithTrx(trxHandle *gorm.DB) MapKejadianRepository {
	if trxHandle == nil {
		mp.logger.Zap.Error("transaction Database not found in gin context.")
		return mp
	}

	mp.db.DB = trxHandle
	return mp
}
