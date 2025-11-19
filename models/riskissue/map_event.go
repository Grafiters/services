package models

type MapEvent struct {
	ID           int64
	IDRiskIssue  int64
	EventTypeLv1 string
	EventTypeLv2 string
	EventTypeLv3 string
}

func (me MapEvent) TableName() string {
	return "risk_issue_map_event"
}
