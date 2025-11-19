package models

type MapAktifitas struct {
	ID           int64
	IDRiskIssue  int64
	Aktifitas    int64
	SubAktifitas int64
}

func (me MapAktifitas) TableName() string {
	return "risk_issue_map_aktifitas"
}
