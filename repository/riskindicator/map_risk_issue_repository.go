package riskindicator

import (
	"riskmanagement/lib"
	models "riskmanagement/models/riskindicator"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type MapRiskIssueDefinition interface {
	GetRiskIssue(id int64) (responses []models.MapRiskIssueResponse, err error)
	WithTrx(trxHandle *gorm.DB) MapRiskIssueRepository
}

type MapRiskIssueRepository struct {
	db     lib.Database
	logger logger.Logger
}

func NewMapRiskIssueRepository(
	db lib.Database,
	logger logger.Logger,
) MapRiskIssueDefinition {
	return MapRiskIssueRepository{
		db:     db,
		logger: logger,
	}
}

// GetRiskIssue implements MapRiskIssueDefinition
func (mr MapRiskIssueRepository) GetRiskIssue(id int64) (responses []models.MapRiskIssueResponse, err error) {
	db := mr.db.DB

	db = db.Table(`risk_issue_map_indicator rimi`).
		Select(`
			ri.risk_issue_code 'kode',
			ri.risk_issue  'risk_issue'`).
		Joins(`JOIN risk_issue ri ON rimi.id_risk_issue = ri.id`).
		Where(`rimi.id_indicator = ?`, id)

	err = db.Find(&responses).Error

	return responses, err
}

// WithTrx implements MapRiskIssueDefinition
func (mr MapRiskIssueRepository) WithTrx(trxHandle *gorm.DB) MapRiskIssueRepository {
	if trxHandle == nil {
		mr.logger.Zap.Error("transaction Database not found in gin context.")
		return mr
	}

	mr.db.DB = trxHandle
	return mr
}
