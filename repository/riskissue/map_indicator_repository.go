package riskissue

import (
	"fmt"
	"riskmanagement/lib"
	modelsIndicator "riskmanagement/models/riskindicator"
	models "riskmanagement/models/riskissue"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type MapIndicatorDefinition interface {
	GetAll() (responses []models.MapIndicatorResponse, err error)
	GetOne(id int64) (responses []models.MapIndicatorResponse, err error)
	Store(request *models.MapIndicator, tx *gorm.DB) (responses *models.MapIndicator, err error)
	Update(request *models.MapIndicator, tx *gorm.DB) (responses bool, err error)
	Delete(id int64) (err error)
	GetOneDataByID(id int64) (responses []models.MapIndicatorResponseFinal, err error)
	DeleteDataByID(id int64, tx *gorm.DB) (err error)
	WithTrx(trxHandle *gorm.DB) MapIndicatorRepository
	GetWithPaginate(id int, filter modelsIndicator.Paginate) (responses []models.MapIndicatorResponseFinal, total int, err error)
}

type MapIndicatorRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewMapIndicatorRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) MapIndicatorDefinition {
	return MapIndicatorRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

func (mp MapIndicatorRepository) GetWithPaginate(id int, filter modelsIndicator.Paginate) (responses []models.MapIndicatorResponseFinal, total int, err error) {
	where := ` WHERE mi.id_risk_issue = ? `
	args := []interface{}{id}

	// ================= FILTER =================
	if filter.Search != "" {
		where += `
		AND (
			ri.risk_indicator_code LIKE ?
			OR ri.risk_indicator LIKE ?
		)
		`
		search := "%" + filter.Search + "%"
		args = append(args, search, search)
	}

	if filter.Code != "" {
		where += ` AND ri.risk_indicator_code LIKE ? `
		args = append(args, "%"+filter.Code+"%")
	}

	if filter.Name != "" {
		where += ` AND ri.risk_indicator LIKE ? `
		args = append(args, "%"+filter.Name+"%")
	}

	if filter.CreatedAt != "" {
		where += ` AND DATE(mi.created_at) = ? `
		args = append(args, filter.CreatedAt)
	}

	if filter.Batasan != "" {
		where += ` AND ri.batasan LIKE ? `
		args = append(args, "%"+filter.Batasan+"%")
	}

	// ================= STATUS FILTER =================
	if filter.Active && !filter.Inactive {
		where += ` AND mi.is_checked = true `
	}
	if filter.Inactive && !filter.Active {
		where += ` AND mi.is_checked = false `
	}

	// ================= COUNT QUERY =================
	countQuery := `
	SELECT COUNT(1)
	FROM risk_issue_map_indicator mi
	JOIN risk_indicator ri ON ri.id = mi.id_indicator
	` + where

	if err = mp.db.DB.Raw(countQuery, args...).Scan(&total).Error; err != nil {
		return responses, 0, err
	}

	// ================= DATA QUERY =================
	dataQuery := `
	SELECT
		mi.id AS id,
		mi.id_risk_issue AS id_risk_issue,
		mi.id_indicator AS id_indicator,
		ri.risk_indicator_code AS kode,
		ri.risk_indicator AS risk_indicator,
		mi.is_checked AS is_checked
	FROM risk_issue_map_indicator mi
	JOIN risk_indicator ri ON ri.id = mi.id_indicator
	` + where +
		fmt.Sprintf(" ORDER BY %s %s LIMIT ? OFFSET ? ", filter.Order, filter.Sort)

	dataArgs := append(args, filter.Limit, filter.Offset)

	rows, err := mp.db.DB.Raw(dataQuery, dataArgs...).Rows()
	if err != nil {
		return responses, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var data models.MapIndicatorResponseFinal
		if err := mp.db.DB.ScanRows(rows, &data); err != nil {
			return responses, 0, err
		}
		responses = append(responses, data)
	}

	return responses, total, nil
}

// Delete implements MapIndicatorDefinition
func (mp MapIndicatorRepository) Delete(id int64) (err error) {
	return mp.db.DB.Where("id = ?", id).Delete(&models.MapIndicatorResponseFinal{}).Error
}

// DeleteDataByID implements MapIndicatorDefinition
func (mp MapIndicatorRepository) DeleteDataByID(id int64, tx *gorm.DB) (err error) {
	return tx.Where("id_risk_issue = ?", id).Delete(&models.MapIndicatorResponseFinal{}).Error
}

// GetAll implements MapIndicatorDefinition
func (mp MapIndicatorRepository) GetAll() (responses []models.MapIndicatorResponse, err error) {
	return responses, mp.db.DB.Find(&responses).Error
}

// GetOne implements MapIndicatorDefinition
func (mp MapIndicatorRepository) GetOne(id int64) (responses []models.MapIndicatorResponse, err error) {
	return responses, mp.db.DB.Where("id = ?", id).Find(&responses).Error
}

// GetOneDataByID implements MapIndicatorDefinition
func (mp MapIndicatorRepository) GetOneDataByID(id int64) (responses []models.MapIndicatorResponseFinal, err error) {
	rows, err := mp.db.DB.Raw(`
	SELECT
		mi.id 'id',
		mi.id_risk_issue 'id_risk_issue',
		mi.id_indicator 'id_indicator',
		ri.risk_indicator_code 'kode',
		ri.risk_indicator 'risk_indicator',
		mi.is_checked 'is_checked'
	FROM risk_issue_map_indicator mi
	JOIN risk_indicator ri ON ri.id = mi.id_indicator
	WHERE id_risk_issue = ?`, id).Rows()
	if err != nil {
		return responses, err
	}

	defer rows.Close()
	var DataMp models.MapIndicatorResponseFinal

	for rows.Next() {
		mp.db.DB.ScanRows(rows, &DataMp)
		responses = append(responses, DataMp)
	}

	return responses, err
}

// Store implements MapIndicatorDefinition
func (mp MapIndicatorRepository) Store(request *models.MapIndicator, tx *gorm.DB) (responses *models.MapIndicator, err error) {
	return request, tx.Save(&request).Error
}

// Update implements MapIndicatorDefinition
func (mp MapIndicatorRepository) Update(request *models.MapIndicator, tx *gorm.DB) (responses bool, err error) {
	return true, tx.Save(&request).Error
}

// WithTrx implements MapIndicatorDefinition
func (mp MapIndicatorRepository) WithTrx(trxHandle *gorm.DB) MapIndicatorRepository {
	if trxHandle == nil {
		mp.logger.Zap.Error("transaction Database not found in gin context.")
		return mp
	}

	mp.db.DB = trxHandle
	return mp
}
