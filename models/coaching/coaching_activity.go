package models

type CoachingActivity struct {
	ID                int64
	CoachingID        int64
	RiskIssueID       int64
	RiskIssue         string
	RiskIssueCode     string
	JudulMateri       string
	RiskIndicatorID   int64
	TitleMateries     string
	RekomendasiMateri string
	MateriTambahan    string
	UpdatedAt         *string
	CreatedAt         *string
}

func (ca CoachingActivity) TableName() string {
	return "coaching_activity"
}
