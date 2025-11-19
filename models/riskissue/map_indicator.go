package models

type MapIndicator struct {
	ID          int64
	IDRiskIssue int64
	IDIndicator int64
	IsChecked   bool
}

func (mi MapIndicator) TableName() string {
	return "risk_issue_map_indicator"
}
