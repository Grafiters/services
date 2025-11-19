package models

type MapControl struct {
	ID          int64
	IDRiskIssue int64
	IDControl   int64
	IsChecked   bool
}

func (me MapControl) TableName() string {
	return "risk_issue_map_control"
}
