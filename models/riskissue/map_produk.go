package models

type MapProduct struct {
	ID          int64
	IDRiskIssue int64
	Product     int64
}

func (me MapProduct) TableName() string {
	return "risk_issue_map_product"
}
