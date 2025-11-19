package pgsuser

import (
	"riskmanagement/lib"
	models "riskmanagement/models/pgsuser"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type PgsUserApprovalDefinition interface {
	WithTrx(trxHandle *gorm.DB) PgsUserApprovalRepository
	GetAll() (responses []models.PgsUserApprovalResponse, err error)
	GetOne(id int64) (responses models.PgsUserApprovalResponse, err error)
	Store(request *models.PgsUserApproval, tx *gorm.DB) (responses *models.PgsUserApproval, err error)
	Update(request *models.PgsUserApprovalRequest) (responses bool, err error)
	Delete(id int64) (err error)
	GeOneApprovalByID(id int64) (responses []models.PgsUserApprovalResponse, err error)
	ApprovalUpdate(request *models.ApprovalUpdate, tx *gorm.DB) (responses bool, err error)
}

type PgsUserApprovalRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewPgsUserApprovalRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) PgsUserApprovalDefinition {
	return PgsUserApprovalRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: 0,
	}
}

// WithTrx implements PgsUserApprovalDefinition
func (pgsUser PgsUserApprovalRepository) WithTrx(trxHandle *gorm.DB) PgsUserApprovalRepository {
	if trxHandle == nil {
		pgsUser.logger.Zap.Error("transaction Database not found in gin context.")
		return pgsUser
	}

	pgsUser.db.DB = trxHandle
	return pgsUser
}

// Delete implements PgsUserApprovalDefinition
func (pgsUser PgsUserApprovalRepository) Delete(id int64) (err error) {
	return pgsUser.db.DB.Where("id = ?", id).Delete(&models.PgsUserApprovalResponse{}).Error
}

// GetAll implements PgsUserApprovalDefinition
func (pgsUser PgsUserApprovalRepository) GetAll() (responses []models.PgsUserApprovalResponse, err error) {
	return responses, pgsUser.db.DB.Find(&responses).Error
}

// GetOne implements PgsUserApprovalDefinition
func (pgsUser PgsUserApprovalRepository) GetOne(id int64) (responses models.PgsUserApprovalResponse, err error) {
	return responses, pgsUser.db.DB.Where("id = ?", id).Find(&responses).Error
}

// GeOneApprovalByID implements PgsUserApprovalDefinition
func (pgsUser PgsUserApprovalRepository) GeOneApprovalByID(id int64) (responses []models.PgsUserApprovalResponse, err error) {
	rows, err := pgsUser.db.DB.Raw(`SELECT * FROM pgs_user_approval WHERE id_pgs_user = ? AND approval_status = 0`, id).Rows()

	defer rows.Close()
	var approval models.PgsUserApprovalResponse

	for rows.Next() {
		pgsUser.db.DB.ScanRows(rows, &approval)
		responses = append(responses, approval)
	}

	return responses, err
}

// Store implements PgsUserApprovalDefinition
func (pgsUser PgsUserApprovalRepository) Store(request *models.PgsUserApproval, tx *gorm.DB) (responses *models.PgsUserApproval, err error) {
	return request, tx.Save(&request).Error
}

// Update implements PgsUserApprovalDefinition
func (pgsUser PgsUserApprovalRepository) Update(request *models.PgsUserApprovalRequest) (responses bool, err error) {
	return true, pgsUser.db.DB.Save(&request).Error
}

// ApprovalUpdate implements PgsUserApprovalDefinition
func (pgsUser PgsUserApprovalRepository) ApprovalUpdate(request *models.ApprovalUpdate, tx *gorm.DB) (responses bool, err error) {
	return true, pgsUser.db.DB.Save(&request).Error
}
