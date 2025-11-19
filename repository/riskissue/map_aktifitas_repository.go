package riskissue

import (
	"riskmanagement/lib"
	models "riskmanagement/models/riskissue"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type MapAktifitasDefinition interface {
	GetAll() (responses []models.MapAktifitasResponse, err error)
	GetOne(id int64) (responses []models.MapAktifitasResponse, err error)
	Store(request *models.MapAktifitas, tx *gorm.DB) (responses *models.MapAktifitas, err error)
	Update(request *models.MapAktifitas, tx *gorm.DB) (responses bool, err error)
	Delete(id int64) (err error)
	GetOneDataByID(id int64) (responses []models.MapAktifitasResponseFinal, err error)
	DeleteDataByID(id int64, tx *gorm.DB) (err error)
	WithTrx(trxHandle *gorm.DB) MapAktifitasRepository
}

type MapAktifitasRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewMapAktifitasRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) MapAktifitasDefinition {
	return MapAktifitasRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// Delete implements MapAktifitasDefinition
func (mp MapAktifitasRepository) Delete(id int64) (err error) {
	return mp.db.DB.Where("id = ?", id).Delete(&models.MapAktifitasResponseFinal{}).Error
}

// DeleteDataByID implements MapAktifitasDefinition
func (mp MapAktifitasRepository) DeleteDataByID(id int64, tx *gorm.DB) (err error) {
	return tx.Where("id_risk_issue = ?", id).Delete(&models.MapAktifitasResponseFinal{}).Error
}

// GetAll implements MapAktifitasDefinition
func (mp MapAktifitasRepository) GetAll() (responses []models.MapAktifitasResponse, err error) {
	return responses, mp.db.DB.Find(&responses).Error
}

// GetOne implements MapAktifitasDefinition
func (mp MapAktifitasRepository) GetOne(id int64) (responses []models.MapAktifitasResponse, err error) {
	return responses, mp.db.DB.Where("id = ?", id).Find(&responses).Error
}

// GetOneDataByID implements MapAktifitasDefinition
func (mp MapAktifitasRepository) GetOneDataByID(id int64) (responses []models.MapAktifitasResponseFinal, err error) {
	rows, err := mp.db.DB.Raw(`
	SELECT
		ma.id 'id',
		ma.id_risk_issue 'id_risk_issue',
		ma.aktifitas 'aktifitas',
		act.name 'aktifitas_desc',
		ma.sub_aktifitas 'sub_aktifitas',
		subact.kode_sub_activity 'kode_sub_aktifitas',
		subact.name 'sub_aktifitas_desc'
	FROM risk_issue_map_aktifitas ma
	LEFT JOIN activity act ON act.id = ma.aktifitas
	LEFT JOIN sub_activity subact ON subact.id = ma.sub_aktifitas
	WHERE id_risk_issue = ?`, id).Rows()
	if err != nil {
		return responses, err
	}

	defer rows.Close()
	var DataMp models.MapAktifitasResponseFinal

	for rows.Next() {
		mp.db.DB.ScanRows(rows, &DataMp)
		responses = append(responses, DataMp)
	}

	return responses, err
}

// Store implements MapAktifitasDefinition
func (mp MapAktifitasRepository) Store(request *models.MapAktifitas, tx *gorm.DB) (responses *models.MapAktifitas, err error) {
	return request, tx.Save(&request).Error
}

// Update implements MapAktifitasDefinition
func (mp MapAktifitasRepository) Update(request *models.MapAktifitas, tx *gorm.DB) (responses bool, err error) {
	return true, tx.Save(&request).Error
}

// WithTrx implements MapAktifitasDefinition
func (mp MapAktifitasRepository) WithTrx(trxHandle *gorm.DB) MapAktifitasRepository {
	if trxHandle == nil {
		mp.logger.Zap.Error("transaction Database not found in gin context.")
		return mp
	}

	mp.db.DB = trxHandle
	return mp
}
