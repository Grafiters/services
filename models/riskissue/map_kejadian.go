package models

type MapKejadian struct {
	ID                  int64
	IDRiskIssue         int64
	PenyebabKejadianLv1 string
	PenyebabKejadianLv2 string
	PenyebabKejadianLv3 string
}

func (me MapKejadian) TableName() string {
	return "risk_issue_map_kejadian"
}
