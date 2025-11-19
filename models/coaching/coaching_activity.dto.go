package models

type CoachingActivityRequest struct {
	ID                int64   `json:"id"`
	CoachingID        int64   `json:"coaching_id"`
	RiskIssueID       int64   `json:"risk_issue_id"`
	RiskIssue         string  `json:"risk_issue"`
	RiskIssueCode     string  `json:"risk_issue_code"`
	JudulMateri       string  `json:"judul_materi"`
	RiskIndicatorID   int64   `json:"risk_indicator_id"`
	RekomendasiMateri string  `json:"rekomendasi_materi"`
	TitleMateries     string  `json:"title_materies"`
	MateriTambahan    string  `json:"materi_tambahan"`
	UpdatedAt         *string `json:"updated_at"`
	CreatedAt         *string `json:"created_at"`
}

type CoachingActivityResponse struct {
	ID                int64   `json:"id"`
	CoachingID        int64   `json:"coaching_id"`
	RiskIssueID       int64   `json:"risk_issue_id"`
	RiskIssue         string  `json:"risk_issue"`
	RiskIssueCode     string  `json:"risk_issue_code"`
	JudulMateri       string  `json:"judul_materi"`
	RiskIndicatorID   int64   `json:"risk_indicator_id"`
	RekomendasiMateri string  `json:"rekomendasi_materi"`
	TitleMateries     string  `json:"title_materies"`
	MateriTambahan    string  `json:"materi_tambahan"`
	UpdatedAt         *string `json:"updated_at"`
	CreatedAt         *string `json:"created_at"`
}

type CoachingActivityResponses struct {
	ID                int64  `json:"id"`
	CoachingID        int64  `json:"coaching_id"`
	RiskIssueID       int64  `json:"risk_issue_id"`
	RiskIssue         string `json:"risk_issue"`
	RiskIssueCode     string `json:"risk_issue_code"`
	JudulMateri       string `json:"judul_materi"`
	RiskIndicatorID   int64  `json:"risk_indicator_id"`
	RekomendasiMateri string `json:"rekomendasi_materi"`
	TitleMateries     string `json:"title_materies"`
	MateriTambahan    string `json:"materi_tambahan"`
}

func (ca CoachingActivityRequest) TableName() string {
	return "coaching_activity"
}

func (ca CoachingActivityResponse) TableName() string {
	return "coaching_activity"
}
