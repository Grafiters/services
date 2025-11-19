package models

type MapLiniBisnis struct {
	ID            int64
	IDRiskIssue   int64
	LiniBisnisLv1 string
	LiniBisnisLv2 string
	LiniBisnisLv3 string
}

func (me MapLiniBisnis) TableName() string {
	return "risk_issue_map_lini_bisnis"
}
