package verifikasi

import (
	"riskmanagement/lib"
	models "riskmanagement/models/verifikasi"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type VerifikasiRiskControlDefinition interface {
	GetAll() (responses []models.VerifikasiRiskControlResponse, err error)
	GetOne(id int64) (responses models.VerifikasiRiskControlResponse, err error)
	GetOneDataByID(id int64) (responses []models.VerifikasiRiskControlResponse, err error)
	Store(request *models.VerifikasiRiskControl, tx *gorm.DB) (responses *models.VerifikasiRiskControl, err error)
	Update(request *models.VerifikasiRiskControl, tx *gorm.DB) (responses bool, err error)
	Delete(id int64) (err error)
	DeleteDataByID(id int64, tx *gorm.DB) (err error)
	WithTrx(trxHandle *gorm.DB) VerifikasiRiskControlRepository
}

type VerifikasiRiskControlRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewVerifikasiRiskControlRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) VerifikasiRiskControlDefinition {
	return VerifikasiRiskControlRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: 0,
	}
}

// Delete implements VerifikasiRiskControlDefinition
func (vRiskControl VerifikasiRiskControlRepository) Delete(id int64) (err error) {
	return vRiskControl.db.DB.Where("id = ?", id).Delete(&models.VerifikasiRiskControlResponse{}).Error
}

// DeleteDataByID implements VerifikasiRiskControlDefinition
func (vRiskControl VerifikasiRiskControlRepository) DeleteDataByID(id int64, tx *gorm.DB) (err error) {
	return tx.Where("verifikasi_id = ?", id).Delete(&models.VerifikasiRiskControlResponse{}).Error
}

// GetAll implements VerifikasiRiskControlDefinition
func (vRiskControl VerifikasiRiskControlRepository) GetAll() (responses []models.VerifikasiRiskControlResponse, err error) {
	return responses, vRiskControl.db.DB.Find(&responses).Error
}

// GetOne implements VerifikasiRiskControlDefinition
func (vRiskControl VerifikasiRiskControlRepository) GetOne(id int64) (responses models.VerifikasiRiskControlResponse, err error) {
	return responses, vRiskControl.db.DB.Where("id = ?", id).Find(&responses).Error
}

// GetOneDataByID implements VerifikasiRiskControlDefinition
func (vRiskControl VerifikasiRiskControlRepository) GetOneDataByID(id int64) (responses []models.VerifikasiRiskControlResponse, err error) {
	// rows, err := vRiskControl.db.DB.Raw(`
	// 	SELECT
	// 		vrCon.id 'id',
	// 		vrCon.verifikasi_id 'verifikasi_id',
	// 		rCon.risk_control 'risk_control'
	// 	FROM verifikasi_risk_control vrCon
	// 	JOIN risk_control rCon on vrCon.risk_control_id = rCon.id
	// 	WHERE vrCon.verifikasi_id = ?`, id).Rows()

	rows, err := vRiskControl.db.DB.Raw(`SELECT * FROM verifikasi_risk_control vrc WHERE vrc.verifikasi_id = ?`, id).Rows()

	defer rows.Close()

	var vrCon models.VerifikasiRiskControlResponse

	for rows.Next() {
		vRiskControl.db.DB.ScanRows(rows, &vrCon)
		responses = append(responses, vrCon)
	}

	return responses, err
}

// Store implements VerifikasiRiskControlDefinition
func (vRiskControl VerifikasiRiskControlRepository) Store(request *models.VerifikasiRiskControl, tx *gorm.DB) (responses *models.VerifikasiRiskControl, err error) {
	return request, tx.Save(&request).Error
}

// Update implements VerifikasiRiskControlDefinition
func (vRiskControl VerifikasiRiskControlRepository) Update(request *models.VerifikasiRiskControl, tx *gorm.DB) (responses bool, err error) {
	return true, tx.Save(&request).Error
}

// WithTrx implements VerifikasiRiskControlDefinition
func (vRiskControl VerifikasiRiskControlRepository) WithTrx(trxHandle *gorm.DB) VerifikasiRiskControlRepository {
	if trxHandle == nil {
		vRiskControl.logger.Zap.Error("transaction Database not found in gin context.")
		return vRiskControl
	}

	vRiskControl.db.DB = trxHandle
	return vRiskControl
}
