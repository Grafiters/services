package models

type RiskIssue struct {
	ID             int64
	RiskTypeID     int64
	RiskIssueCode  string
	RiskIssue      string
	Deskripsi      string
	KategoriRisiko string
	Status         bool
	Likelihood     *string
	Impact         *string
	CreatedAt      *string
	UpdatedAt      *string
	DeleteFlag     bool
}

type RiskIssueDeleteRequest struct {
	ID         int64
	DeleteFlag bool
	UpdatedAt  *string
}

type RiskIssueUpdate struct {
	ID             int64
	RiskTypeID     int64
	RiskIssueCode  string
	RiskIssue      string
	Deskripsi      string
	KategoriRisiko string
	Status         bool
	Likelihood     *string
	Impact         *string
	UpdatedAt      *string
	DeleteFlag     bool
}

func (r RiskIssueUpdate) TableName() string {
	return "risk_issue"
}

func (r RiskIssueDeleteRequest) TableName() string {
	return "risk_issue"
}

func (r RiskIssue) TableName() string {
	return "risk_issue"
}
