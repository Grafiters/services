package managementuser

import (
	"riskmanagement/lib"
	models "riskmanagement/models/managementuser"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type MapMenuDefinition interface {
	GetAll() (responses []models.MapMenuResponse, err error)
	GetOne(id int64) (responses []models.MapMenuResponse, err error)
	Store(request *models.MapMenu, tx *gorm.DB) (responses *models.MapMenu, err error)
	Update(request *models.MapMenu, tx *gorm.DB) (responses bool, err error)
	Delete(id int64) (err error)
	GetOneDataByID(id int64) (responses []models.MapMenuResponseFinal, err error)
	DeleteDataByID(id int64, tx *gorm.DB) (err error)
	WithTrx(trxHandle *gorm.DB) MapMenuRepository
}

type MapMenuRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewMapMenuRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) MapMenuDefinition {
	return MapMenuRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// Delete implements MapMenuDefinition
func (mp MapMenuRepository) Delete(id int64) (err error) {
	return mp.db.DB.Where("id = ?", id).Delete(&models.MapMenuResponseFinal{}).Error
}

// DeleteDataByID implements MapMenuDefinition
func (mp MapMenuRepository) DeleteDataByID(id int64, tx *gorm.DB) (err error) {
	return tx.Where("id_jabatan = ?", id).Delete(&models.MapMenuResponseFinal{}).Error
}

// GetAll implements MapMenuDefinition
func (mp MapMenuRepository) GetAll() (responses []models.MapMenuResponse, err error) {
	return responses, mp.db.DB.Find(&responses).Error
}

// GetOne implements MapMenuDefinition
func (mp MapMenuRepository) GetOne(id int64) (responses []models.MapMenuResponse, err error) {
	return responses, mp.db.DB.Where("id = ?", id).Find(&responses).Error
}

// GetOneDataByID implements MapMenuDefinition
func (mp MapMenuRepository) GetOneDataByID(id int64) (responses []models.MapMenuResponseFinal, err error) {
	rows, err := mp.db.DB.Raw(`
	SELECT 
		mapMenu.id 'id',
		mapMenu.id_jabatan 'id_jabatan',
		mapMenu.id_menu 'id_menu',
		menu.Title 'Title',
		mapMenu.Keterangan 'Keterangan'
	FROM management_user_map_menu mapMenu
	JOIN mst_menu menu ON menu.id_menu = mapMenu.id_menu
	WHERE mapMenu.id_jabatan = ?`, id).Rows()

	defer rows.Close()
	var DataMp models.MapMenuResponseFinal

	for rows.Next() {
		mp.db.DB.ScanRows(rows, &DataMp)
		responses = append(responses, DataMp)
	}

	return responses, err
}

// Store implements MapMenuDefinition
func (mp MapMenuRepository) Store(request *models.MapMenu, tx *gorm.DB) (responses *models.MapMenu, err error) {
	return request, tx.Save(&request).Error
}

// Update implements MapMenuDefinition
func (mp MapMenuRepository) Update(request *models.MapMenu, tx *gorm.DB) (responses bool, err error) {
	return true, tx.Save(&request).Error
}

// WithTrx implements MapMenuDefinition
func (mp MapMenuRepository) WithTrx(trxHandle *gorm.DB) MapMenuRepository {
	if trxHandle == nil {
		mp.logger.Zap.Error("transaction Database not found in gin context.")
		return mp
	}

	mp.db.DB = trxHandle
	return mp
}
