package risktype

import (
	"math"
	"riskmanagement/lib"
	models "riskmanagement/models/risktype"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type RiskTypeDefinition interface {
	GetAll() (responses []models.RiskTypeResponse, err error)
	GetAllWithPaginate(request *models.Paginate) (responses []models.RiskTypeResponse, totalRows int, totalData int, err error)
	GetOne(id int64) (responses models.RiskTypeResponse, err error)
	Store(request *models.RiskTypeRequest) (responses bool, err error)
	Update(request *models.RiskTypeRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) RiskTypeRepository
}

type RiskTypeRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

// GetAllWithPaginate implements RiskTypeDefinition
func (rt RiskTypeRepository) GetAllWithPaginate(request *models.Paginate) (responses []models.RiskTypeResponse, totalRows int, totalData int, err error) {
	rows, err := rt.db.DB.Raw(`
	SELECT
		rt.id 'id',
		rt.risk_type_code 'risk_type_code',
		rt.risk_type 'risk_type',
		rt.deskripsi 'deskripsi',
		rt.status 'status',
		rt.created_at 'created_at',
		rt.updated_at 'updated_at'
	FROM risk_type rt  ORDER BY rt.id LIMIT ? OFFSET ?
	`, request.Limit, request.Offset).Rows()
	if err != nil {
		return responses, totalRows, totalData, err
	}
	defer rows.Close()

	var linibisnislv1 models.RiskTypeResponse
	for rows.Next() {
		rt.db.DB.ScanRows(rows, &linibisnislv1)
		responses = append(responses, linibisnislv1)
	}

	paginateQuery := `SELECT COUNT(*) FROM risk_type`
	err = rt.dbRaw.DB.QueryRow(paginateQuery).Scan(&totalData)
	if err != nil {
		return responses, totalRows, totalData, err
	}

	if totalData > 0 {
		totalRows = int(math.Ceil((float64(totalData)) / float64(request.Limit)))
	}

	return responses, totalRows, totalData, err
}

// Delete implements RiskTypeDefinition
func (riskType RiskTypeRepository) Delete(id int64) (err error) {
	return riskType.db.DB.Where("id = ?", id).Delete(&models.RiskTypeResponse{}).Error
}

// GetAll implements RiskTypeDefinition
func (riskType RiskTypeRepository) GetAll() (responses []models.RiskTypeResponse, err error) {
	return responses, riskType.db.DB.Find(&responses).Error
}

// GetOne implements RiskTypeDefinition
func (riskType RiskTypeRepository) GetOne(id int64) (responses models.RiskTypeResponse, err error) {
	return responses, riskType.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements RiskTypeDefinition
func (riskType RiskTypeRepository) Store(request *models.RiskTypeRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	err = riskType.db.DB.Save(&models.RiskTypeRequest{
		RiskTypeCode: request.RiskTypeCode,
		RiskType:     request.RiskType,
		Deskripsi:    request.Deskripsi,
		Status:       request.Status,
		CreatedAt:    &timeNow,
	}).Error
	if err != nil {
		return false, err
	}

	return true, err
}

// Update implements RiskTypeDefinition
func (riskType RiskTypeRepository) Update(request *models.RiskTypeRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	err = riskType.db.DB.Save(&models.RiskTypeRequest{
		ID:           request.ID,
		RiskTypeCode: request.RiskTypeCode,
		RiskType:     request.RiskType,
		Deskripsi:    request.Deskripsi,
		Status:       request.Status,
		CreatedAt:    request.CreatedAt,
		UpdatedAt:    &timeNow,
	}).Error
	if err != nil {
		return false, err
	}

	return true, err
}

// WithTrx implements RiskTypeDefinition
func (riskType RiskTypeRepository) WithTrx(trxHandle *gorm.DB) RiskTypeRepository {
	if trxHandle == nil {
		riskType.logger.Zap.Error("transaction Database not found in gin context")
		return riskType
	}

	riskType.db.DB = trxHandle
	return riskType
}

func NewRiskTypeRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) RiskTypeDefinition {
	return RiskTypeRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}
