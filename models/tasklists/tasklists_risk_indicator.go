package models

type TasklistsRiskIndicator struct {
	RiskIndicatorID  	string  	`json:"risk_indicator_id"`
	TasklistsID			int64
}

func (TasklistsRiskIndicator) TableName() string {
	return "tasklists_risk_indicator"
}