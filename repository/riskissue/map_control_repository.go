package riskissue

import (
	"database/sql"
	"fmt"
	"riskmanagement/lib"
	riskControlModels "riskmanagement/models/riskcontrol"
	models "riskmanagement/models/riskissue"
	"strings"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type MapControlDefinition interface {
	GetAll() (responses []models.MapControlResponse, err error)
	GetOne(id int64) (responses []models.MapControlResponse, err error)
	Store(request *models.MapControl, tx *gorm.DB) (responses *models.MapControl, err error)
	Update(request *models.MapControl, tx *gorm.DB) (responses bool, err error)
	Delete(id int64) (err error)
	GetOneDataByID(id int64) (responses []models.MapControlResponseFinal, err error)
	DeleteDataByID(id int64, tx *gorm.DB) (err error)
	WithTrx(trxHandle *gorm.DB) MapControlRepository
	GetWithPagination(id int, filter riskControlModels.Paginate) (responses []models.MapControlResponseFinal, total int, err error)
}

type MapControlRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewMapControlRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) MapControlDefinition {
	return MapControlRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

func checkCount(rows *sql.Rows) (count int) {
	for rows.Next() {
		err := rows.Scan(&count)
		checkErr(err)
	}
	return count
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (mp MapControlRepository) GetWithPagination(id int, filter riskControlModels.Paginate) (responses []models.MapControlResponseFinal, total int, err error) {
	where := ` WHERE mc.id_risk_issue = ? `
	args := []interface{}{id}

	// ================= FILTER =================
	if filter.Search != "" {
		where += `
		AND (
			rc.kode LIKE ?
			OR rc.risk_control LIKE ?
		)
		`
		search := "%" + filter.Search + "%"
		args = append(args, search, search)
	}

	if filter.Kode != "" {
		where += ` AND rc.kode LIKE ? `
		args = append(args, "%"+filter.Kode+"%")
	}

	if filter.RiskControl != "" {
		where += ` AND rc.risk_control LIKE ? `
		args = append(args, "%"+filter.RiskControl+"%")
	}

	if filter.Status != "" {
		switch strings.ToLower(filter.Status) {
		case "active":
			where += ` AND mc.is_checked = true `
		case "inactive":
			where += ` AND mc.is_checked = false `
		}
	}

	// ================= COUNT QUERY =================
	countQuery := `
	SELECT COUNT(1)
	FROM risk_issue_map_control mc
	JOIN risk_control rc ON rc.id = mc.id_control
	` + where

	if err = mp.db.DB.Raw(countQuery, args...).Scan(&total).Error; err != nil {
		return responses, 0, err
	}

	// ================= DATA QUERY =================
	dataQuery := `
		SELECT
			mc.id AS id,
			mc.id_risk_issue AS id_risk_issue,
			mc.id_control AS id_control,
			rc.kode AS kode,
			rc.risk_control AS risk_control,
			mc.is_checked AS is_checked
		FROM risk_issue_map_control mc
		JOIN risk_control rc ON rc.id = mc.id_control
		` + where +
		fmt.Sprintf(" ORDER BY %s %s LIMIT ? OFFSET ? ", filter.Order, filter.Sort)

	dataArgs := append(args, filter.Limit, filter.Offset)

	rows, err := mp.db.DB.Raw(dataQuery, dataArgs...).Rows()
	if err != nil {
		return responses, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var data models.MapControlResponseFinal
		if err := mp.db.DB.ScanRows(rows, &data); err != nil {
			return responses, 0, err
		}
		responses = append(responses, data)
	}

	return responses, total, nil
}

// Delete implements MapControlDefinition
func (mp MapControlRepository) Delete(id int64) (err error) {
	return mp.db.DB.Where("id = ?", id).Delete(&models.MapControlResponseFinal{}).Error
}

// DeleteDataByID implements MapControlDefinition
func (mp MapControlRepository) DeleteDataByID(id int64, tx *gorm.DB) (err error) {
	return tx.Where("id_risk_issue = ?", id).Delete(&models.MapControlResponseFinal{}).Error
}

// GetAll implements MapControlDefinition
func (mp MapControlRepository) GetAll() (responses []models.MapControlResponse, err error) {
	return responses, mp.db.DB.Find(&responses).Error
}

// GetOne implements MapControlDefinition
func (mp MapControlRepository) GetOne(id int64) (responses []models.MapControlResponse, err error) {
	return responses, mp.db.DB.Where("id = ?", id).Find(&responses).Error
}

// GetOneDataByID implements MapControlDefinition
func (mp MapControlRepository) GetOneDataByID(id int64) (responses []models.MapControlResponseFinal, err error) {
	rows, err := mp.db.DB.Raw(`
	SELECT
		mc.id 'id',
		mc.id_risk_issue 'id_risk_issue',
		mc.id_control 'id_control',
		rc.kode 'kode',
		rc.risk_control 'risk_control',
		mc.is_checked 'is_checked'
	FROM risk_issue_map_control mc
	JOIN risk_control rc ON rc.id = mc.id_control
	WHERE id_risk_issue = ?`, id).Rows()
	if err != nil {
		return responses, err
	}

	defer rows.Close()
	var DataMp models.MapControlResponseFinal

	for rows.Next() {
		mp.db.DB.ScanRows(rows, &DataMp)
		responses = append(responses, DataMp)
	}

	return responses, err
}

// Store implements MapControlDefinition
func (mp MapControlRepository) Store(request *models.MapControl, tx *gorm.DB) (responses *models.MapControl, err error) {
	return request, tx.Save(&request).Error
}

// Update implements MapControlDefinition
func (mp MapControlRepository) Update(request *models.MapControl, tx *gorm.DB) (responses bool, err error) {
	return true, tx.Save(&request).Error
}

// WithTrx implements MapControlDefinition
func (mp MapControlRepository) WithTrx(trxHandle *gorm.DB) MapControlRepository {
	if trxHandle == nil {
		mp.logger.Zap.Error("transaction Database not found in gin context.")
		return mp
	}

	mp.db.DB = trxHandle
	return mp
}
