package models

type MapProses struct {
	ID             int64
	IDRiskIssue    int64
	MegaProses     string
	MajorProses    string
	SubMajorProses string
}

func (mp MapProses) TableName() string {
	return "risk_issue_map_proses"
}
