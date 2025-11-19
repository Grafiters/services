package models

type MapIndicatorRequest struct {
	ID          int64 `json:"id"`
	IDRiskIssue int64 `json:"id_risk_issue"`
	IDIndicator int64 `json:"id_indicator"`
	IsChecked   bool  `json:"is_checked"`
}

type MapIndicatorResponse struct {
	ID          int64 `json:"id"`
	IDRiskIssue int64 `json:"id_risk_issue"`
	IDIndicator int64 `json:"id_indicator"`
	IsChecked   bool  `json:"is_checked"`
}

type MapIndicatorResponseFinal struct {
	ID            int64  `json:"id"`
	IDRiskIssue   int64  `json:"id_risk_issue"`
	IDIndicator   int64  `json:"id_indicator"`
	Kode          string `json:"kode"`
	RiskIndicator string `json:"risk_indicator"`
	IsChecked     bool   `json:"is_checked"`
}

func (p MapIndicatorRequest) ParseRequest() MapIndicator {
	return MapIndicator{
		ID:          p.ID,
		IDRiskIssue: p.IDRiskIssue,
		IDIndicator: p.IDIndicator,
		IsChecked:   p.IsChecked,
	}
}

func (p MapIndicatorResponse) ParseResponse() MapIndicator {
	return MapIndicator{
		ID:          p.ID,
		IDRiskIssue: p.IDRiskIssue,
		IDIndicator: p.IDIndicator,
		IsChecked:   p.IsChecked,
	}
}

func (ma MapIndicatorRequest) TableName() string {
	return "risk_issue_map_indicator"
}

func (ma MapIndicatorResponse) TableName() string {
	return "risk_issue_map_indicator"
}
