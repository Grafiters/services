package riskcontrol

import (
	"errors"
	"fmt"
	"math"
	"riskmanagement/lib"
	models "riskmanagement/models/riskcontrol"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type RiskControlDefinition interface {
	GetAll() (responses []models.RiskControlResponse, err error)
	GetAllWithPaginate(request *models.Paginate) (responses []models.RiskControlResponse, totalData int64, totalRows int64, err error)
	GetOne(id int64) (responses models.RiskControlResponse, err error)
	Store(request *models.RiskControlRequest) (responses bool, err error)
	Update(request *models.RiskControlRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) RiskControlRepository
	GetKodeRiskControl() (responses []models.KodeRiskControl, err error)
	GenLastCode() (string, error)
	UpdateStatus(id int64, status bool) error
	BulkCreateRiskControl(items []models.RiskControlRequest, tx *gorm.DB) error
	SearchRiskControlByIssue(request *models.KeywordRequest) (responses []models.RiskControlResponses, totalRows int, totalData int, err error)
}

type RiskControlRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewRiskControlRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) RiskControlDefinition {
	return RiskControlRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// Delete implements RiskControlDefinition
func (riskControl RiskControlRepository) Delete(id int64) (err error) {
	return riskControl.db.DB.Where("id = ?", id).Delete(&models.RiskControlResponse{}).Error
}

// GetAll implements RiskControlDefinition
func (riskControl RiskControlRepository) GetAll() (responses []models.RiskControlResponse, err error) {
	return responses, riskControl.db.DB.Find(&responses).Error
}

// GetAllWithPaginate implements RiskControlDefinition
func (rc RiskControlRepository) GetAllWithPaginate(request *models.Paginate) (responses []models.RiskControlResponse, totalData int64, totalRows int64, err error) {
	db := rc.db.DB

	// create a new query builder with the base query
	queryBuilder := db.Model(&responses).
		Select(
			`id 'id',
			kode 'kode',
			risk_control 'risk_control',
			deskripsi 'deskripsi',
			status 'status',
			control_type 'control_type',
			nature 'nature',
			key_control 'key_control',
			owner_lvl 'owner_lvl',
			owner_group 'owner_group',
			owner 'owner',
			document 'document',
			created_at 'created_at',
			updated_at 'updated_at'
	`).Order("created_at ASC")

	// add dynamic where clauses
	if request.Kode != "" {
		queryBuilder = queryBuilder.Where("kode = ?", request.Kode)
	}
	if request.RiskControl != "" {
		queryBuilder = queryBuilder.Where("risk_control LIKE ?", fmt.Sprintf("%%%s%%", request.RiskControl))
	}

	if request.Status != "" {
		if request.Status == "Aktif" {
			queryBuilder = queryBuilder.Where("status = 1")
		} else {
			queryBuilder = queryBuilder.Where("status = 0")
		}
	}

	if request.Search != "" {
		search := "%" + request.Search + "%"
		queryBuilder = queryBuilder.Where("kode LIKE ? OR risk_control LIKE ?", search, search)
	}
	// if request.Status != nil {
	// 	queryBuilder = queryBuilder.Where("status = ?", request.Status)
	// }

	// count the total rows
	err = queryBuilder.Count(&totalData).Error
	if err != nil {
		return responses, totalRows, totalData, err
	}

	// execute the query
	queryBuilder.Limit(request.Limit).Offset(request.Offset).Find(&responses)

	if totalData > 0 {
		totalRows = int64(math.Ceil(float64(totalData) / float64(request.Limit)))
	}

	return responses, totalData, totalRows, err
}

// GetOne implements RiskControlDefinition
func (riskControl RiskControlRepository) GetOne(id int64) (responses models.RiskControlResponse, err error) {
	err = riskControl.db.DB.Select(`id, kode,risk_control, control_type, nature, key_control,deskripsi, owner_lvl, owner_group, owner, document, status, created_at`).
		Where("id = ?", id).Find(&responses).Error

	return responses, err

}

// Store implements RiskControlDefinition
func (riskControl RiskControlRepository) Store(request *models.RiskControlRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	err = riskControl.db.DB.Save(&models.RiskControlRequest{
		Kode:        request.Kode,
		RiskControl: request.RiskControl,
		ControlType: request.ControlType,
		Nature:      request.Nature,
		KeyControl:  request.KeyControl,
		Deskripsi:   request.Deskripsi,
		Status:      request.Status,
		CreatedAt:   &timeNow,
		OwnerLvl:    request.OwnerLvl,
		OwnerGroup:  request.OwnerGroup,
		Owner:       request.Owner,
		Document:    request.Document,
	}).Error
	if err != nil {
		return false, err
	}

	return true, err
}

// Update implements RiskControlDefinition
func (riskControl RiskControlRepository) Update(request *models.RiskControlRequest) (responses bool, err error) {
	err = riskControl.db.DB.Save(request).Error
	if err != nil {
		return false, err
	}

	return true, err
}

// WithTrx implements RiskControlDefinition
func (riskControl RiskControlRepository) WithTrx(trxHandle *gorm.DB) RiskControlRepository {
	if trxHandle == nil {
		riskControl.logger.Zap.Error("transaction Database not found in gin context")
		return riskControl
	}

	riskControl.db.DB = trxHandle
	return riskControl
}

// GetKodeRiskControl implements RiskControlDefinition
func (riskControl RiskControlRepository) GetKodeRiskControl() (responses []models.KodeRiskControl, err error) {

	//Local
	// query := `SELECT * from (
	// 			SELECT
	// 				CAST(
	// 					CASE
	// 						WHEN LENGTH(kode) = 5 THEN RIGHT(kode,4)
	// 						WHEN LENGTH(kode) = 4  THEN RIGHT(kode,3)
	// 						WHEN LENGTH(kode) = 3  THEN RIGHT(kode,2)
	// 						WHEN LENGTH(kode) = 2  THEN RIGHT(kode,1)
	// 					END AS int
	// 				) + 1 'kode_risk_control'
	// 			FROM risk_control
	// 		) as t
	// 		ORDER BY t.kode_risk_control DESC LIMIT 1`

	//server dev
	query := `SELECT * FROM (SELECT 
				CASE
						WHEN LENGTH(kode) = 5 THEN CAST(RIGHT(kode,4) AS DECIMAL) 
						WHEN LENGTH(kode) = 4  THEN CAST(RIGHT(kode,3) AS DECIMAL)
						WHEN LENGTH(kode) = 3  THEN CAST(RIGHT(kode,2) AS DECIMAL)
						WHEN LENGTH(kode) = 2  THEN CAST(RIGHT(kode,1) AS DECIMAL)
				END + 1 'kode_risk_control' 
				FROM risk_control rc) as t
				ORDER BY t.kode_risk_control DESC LIMIT 1`

	rows, err := riskControl.dbRaw.DB.Query(query)
	if err != nil {
		return responses, err
	}
	defer rows.Close()

	riskControl.logger.Zap.Info("rows =>", rows)

	response := models.KodeRiskControl{}
	for rows.Next() {
		_ = rows.Scan(
			&response.KodeRiskControl,
		)
		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, err
	}

	return responses, nil
}

func (riskControl RiskControlRepository) GenLastCode() (string, error) {
	var last models.RiskControlResponse
	const prefix = "RC-"

	err := riskControl.db.DB.
		Select("kode").
		Order("created_at DESC").
		First(&last).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Sprintf("%s%05d", prefix, 1), nil
	}
	if err != nil {
		return "", err
	}

	lastNumber := 0
	if _, err := fmt.Sscanf(last.Kode, prefix+"%d", &lastNumber); err != nil {
		lastNumber = 0
	}

	newNumber := lastNumber + 1
	newCode := fmt.Sprintf("%s%05d", prefix, newNumber)

	return newCode, nil
}

// SearchRiskControlByIssue implements RiskControlDefinition
func (rc RiskControlRepository) SearchRiskControlByIssue(request *models.KeywordRequest) (responses []models.RiskControlResponses, totalRows int, totalData int, err error) {
	db := rc.db.DB.Table("risk_control rc")

	query := db.Select(`
			rc.id 'id',
			rc.kode 'kode',
			rc.risk_control 'risk_control'`).
		Joins(`INNER JOIN risk_issue_map_control rimc on rimc.id_control = rc.id`).
		Where(`rimc.id_risk_issue = ?`, request.RiskIssueId)

	if request.Keyword != "" {
		query = query.Where("rc.risk_control LIKE ?", fmt.Sprintf("%%%s%%", request.Keyword))
	}

	if request.Limit != 0 && request.Offset != 0 {
		query = query.Limit(request.Limit).Offset(request.Offset)
	}

	var Count int64
	query = query.Count(&Count)

	totalData = int(Count)

	totalRows = int(math.Ceil(float64(totalData) / float64(request.Limit)))

	if err = query.Scan(&responses).Error; err != nil {
		rc.logger.Zap.Error(err)
		return responses, totalRows, totalData, err
	}

	return responses, totalRows, totalData, nil
}

// Update data risk control for status only
func (rc RiskControlRepository) UpdateStatus(id int64, status bool) error {
	err := rc.db.DB.
		Table("risk_control"). // pastikan ini sesuai nama table di DB
		Where("id = ?", id).
		Update("status", status).Error // GORM syntax: Update(column, value)

	return err
}

// Store multiple data with batch
func (rc RiskControlRepository) BulkCreateRiskControl(items []models.RiskControlRequest, tx *gorm.DB) error {
	if len(items) == 0 {
		return nil
	}

	batchSize := 100
	if len(items) > 1000 {
		batchSize = 500
	} else if len(items) > 5000 {
		batchSize = 1000
	}

	return tx.CreateInBatches(items, batchSize).Error
}
