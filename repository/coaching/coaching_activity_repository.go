package coaching

import (
	"riskmanagement/lib"
	models "riskmanagement/models/coaching"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type CoachingActivityDefinition interface {
	WithTrx(trxHandle *gorm.DB) CoachingActivityRepository
	GetAll() (responses []models.CoachingActivityResponse, err error)
	GetOne(id int64) (responses models.CoachingActivityResponse, err error)
	GetOneActivity(id int64) (responses []models.CoachingActivityResponses, err error)
	Store(request *models.CoachingActivity, tx *gorm.DB) (responses *models.CoachingActivity, err error)
	UpdatedIT(request *models.CoachingActivity, tx *gorm.DB) (responses bool, err error)
	Update(request *models.CoachingActivityRequest) (responses bool, err error)
	Delete(id int64) (err error)
	DeleteCoachingID(id int64, tx *gorm.DB) (err error)
	GetAktifitasReport(id int64) (responses []models.ActivityReportResponse, err error)
}

type CoachingActivityRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewCoachingActivityRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) CoachingActivityDefinition {
	return CoachingActivityRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// Delete implements CoachingActivityDefinition
func (CoachingActivity CoachingActivityRepository) Delete(id int64) (err error) {
	return CoachingActivity.db.DB.Find("id = ?", id).Delete(&models.CoachingActivityResponse{}).Error
}

// DeleteCoachingID implements CoachingActivityDefinition
func (CoachingActivity CoachingActivityRepository) DeleteCoachingID(id int64, tx *gorm.DB) (err error) {
	return tx.Where("coaching_id = ?", id).Delete(&models.CoachingActivityResponse{}).Error
}

// GetAll implements CoachingActivityDefinition
func (CoachingActivity CoachingActivityRepository) GetAll() (responses []models.CoachingActivityResponse, err error) {
	return responses, CoachingActivity.db.DB.Find(&responses).Error
}

// GetOne implements CoachingActivityDefinition
func (CoachingActivity CoachingActivityRepository) GetOne(id int64) (responses models.CoachingActivityResponse, err error) {
	return responses, CoachingActivity.db.DB.Where("id = ?", id).Find(&responses).Error
}

// GetOneActivity implements CoachingActivityDefinition
func (CoachingActivity CoachingActivityRepository) GetOneActivity(id int64) (responses []models.CoachingActivityResponses, err error) {
	rows, err := CoachingActivity.db.DB.Raw(`
		SELECT ca.* FROM coaching_activity ca WHERE ca.coaching_id = ?`, id).Rows()

	defer rows.Close()
	var activity models.CoachingActivityResponses

	for rows.Next() {
		CoachingActivity.db.DB.ScanRows(rows, &activity)
		responses = append(responses, activity)
	}

	return responses, err
}

// Store implements CoachingActivityDefinition
func (CoachingActivity CoachingActivityRepository) Store(request *models.CoachingActivity, tx *gorm.DB) (responses *models.CoachingActivity, err error) {
	return request, tx.Save(&request).Error
}

// Update implements CoachingActivityDefinition
func (CoachingActivity CoachingActivityRepository) Update(request *models.CoachingActivityRequest) (responses bool, err error) {
	return true, CoachingActivity.db.DB.Find(&request).Error
}

// UpdatedIT implements CoachingActivityDefinition
func (CoachingActivity CoachingActivityRepository) UpdatedIT(request *models.CoachingActivity, tx *gorm.DB) (responses bool, err error) {
	return true, tx.Save(&request).Error
}

// WithTrx implements CoachingActivityDefinition
func (CoachingActivity CoachingActivityRepository) WithTrx(trxHandle *gorm.DB) CoachingActivityRepository {
	if trxHandle == nil {
		CoachingActivity.logger.Zap.Error("transaction Database not found in gin context.")
		return CoachingActivity
	}
	CoachingActivity.db.DB = trxHandle
	return CoachingActivity
}

// GetAktifitasReport implements CoachingActivityDefinition
func (repo CoachingActivityRepository) GetAktifitasReport(id int64) (responses []models.ActivityReportResponse, err error) {
	rows, err := repo.db.DB.Raw(`
				SELECT 
					ca.judul_materi,
					ca.materi_tambahan  'rincian_materi',
					ri.risk_issue_code  'isu_risiko'
				FROM coaching_activity ca
				JOIN risk_issue ri ON ri.id = ca.risk_issue_id 
				WHERE ca.coaching_id = ?`, id).Rows()

	defer rows.Close()
	var activity models.ActivityReportResponse

	for rows.Next() {
		repo.db.DB.ScanRows(rows, &activity)
		responses = append(responses, activity)
	}

	return responses, err
}
