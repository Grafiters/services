package models

type MapControlRequest struct {
	ID          int64 `json:"id"`
	IDRiskIssue int64 `json:"id_risk_issue"`
	IDControl   int64 `json:"id_control"`
	IsChecked   bool  `json:"is_checked"`
}

type MapControlResponse struct {
	ID          int64 `json:"id"`
	IDRiskIssue int64 `json:"id_risk_issue"`
	IDControl   int64 `json:"id_control"`
	IsChecked   bool  `json:"is_checked"`
}

type MapControlResponseFinal struct {
	ID          int64  `json:"id"`
	IDRiskIssue int64  `json:"id_risk_issue"`
	IDControl   int64  `json:"id_control"`
	Kode        string `json:"kode"`
	RiskControl string `json:"risk_control"`
	IsChecked   bool   `json:"is_checked"`
}

func (p MapControlRequest) ParseRequest() MapControl {
	return MapControl{
		ID:          p.ID,
		IDRiskIssue: p.IDRiskIssue,
		IDControl:   p.IDControl,
		IsChecked:   p.IsChecked,
	}
}

func (p MapControlResponse) ParseResponse() MapControl {
	return MapControl{
		ID:          p.ID,
		IDRiskIssue: p.IDRiskIssue,
		IDControl:   p.IDControl,
		IsChecked:   p.IsChecked,
	}
}

func (ma MapControlRequest) TableName() string {
	return "risk_issue_map_control"
}

func (ma MapControlResponse) TableName() string {
	return "risk_issue_map_control"
}
