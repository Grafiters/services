package models

type MapRiskIssue struct {
	Kode      string
	RiskIssue string
}

type MapRiskIssueResponse struct {
	Kode      string `json:"kode"`
	RiskIssue string `json:"risk_issue"`
}
