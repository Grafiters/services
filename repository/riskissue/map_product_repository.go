package riskissue

import (
	"riskmanagement/lib"
	models "riskmanagement/models/riskissue"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type MapProductDefinition interface {
	GetAll() (responses []models.MapProductResponse, err error)
	GetOne(id int64) (responses []models.MapProductResponse, err error)
	Store(request *models.MapProduct, tx *gorm.DB) (responses *models.MapProduct, err error)
	Update(request *models.MapProduct, tx *gorm.DB) (responses bool, err error)
	Delete(id int64) (err error)
	GetOneDataByID(id int64) (responses []models.MapProductResponseFinal, err error)
	DeleteDataByID(id int64, tx *gorm.DB) (err error)
	BulkCreate(
		req []*models.MapProduct,
		tx *gorm.DB,
	) error
	WithTrx(trxHandle *gorm.DB) MapProductRepository
}

type MapProductRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewMapProductRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) MapProductDefinition {
	return MapProductRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// Delete implements MapProductDefinition
func (mp MapProductRepository) Delete(id int64) (err error) {
	return mp.db.DB.Where("id = ?", id).Delete(&models.MapProductResponseFinal{}).Error
}

// DeleteDataByID implements MapProductDefinition
func (mp MapProductRepository) DeleteDataByID(id int64, tx *gorm.DB) (err error) {
	return tx.Where("id_risk_issue = ?", id).Delete(&models.MapProductResponseFinal{}).Error
}

// GetAll implements MapProductDefinition
func (mp MapProductRepository) GetAll() (responses []models.MapProductResponse, err error) {
	return responses, mp.db.DB.Find(&responses).Error
}

// GetOne implements MapProductDefinition
func (mp MapProductRepository) GetOne(id int64) (responses []models.MapProductResponse, err error) {
	return responses, mp.db.DB.Where("id = ?", id).Find(&responses).Error
}

// GetOneDataByID implements MapProductDefinition
func (mp MapProductRepository) GetOneDataByID(id int64) (responses []models.MapProductResponseFinal, err error) {
	rows, err := mp.db.DB.Raw(`
	SELECT
		mp.id 'id',
		mp.id_risk_issue 'id_risk_issue',
		mp.product 'product',
		prod.product 'product_desc'
	FROM risk_issue_map_product mp
	LEFT JOIN product prod ON prod.id = mp.product
	WHERE id_risk_issue = ?`, id).Rows()
	if err != nil {
		return responses, err
	}

	defer rows.Close()
	var DataMp models.MapProductResponseFinal

	for rows.Next() {
		mp.db.DB.ScanRows(rows, &DataMp)
		responses = append(responses, DataMp)
	}

	return responses, err
}

// Store implements MapProductDefinition
func (mp MapProductRepository) Store(request *models.MapProduct, tx *gorm.DB) (responses *models.MapProduct, err error) {
	return request, tx.Save(&request).Error
}

// Update implements MapProductDefinition
func (mp MapProductRepository) Update(request *models.MapProduct, tx *gorm.DB) (responses bool, err error) {
	return true, tx.Save(&request).Error
}

func (mp MapProductRepository) BulkCreate(
	req []*models.MapProduct,
	tx *gorm.DB,
) error {
	if len(req) == 0 {
		return nil
	}

	return tx.Create(req).Error
}

// WithTrx implements MapProductDefinition
func (mp MapProductRepository) WithTrx(trxHandle *gorm.DB) MapProductRepository {
	if trxHandle == nil {
		mp.logger.Zap.Error("transaction Database not found in gin context.")
		return mp
	}

	mp.db.DB = trxHandle
	return mp
}
