package models

type MapProsesRequest struct {
	ID             int64  `json:"id"`
	IDRiskIssue    int64  `json:"id_risk_issue"`
	MegaProses     string `json:"mega_proses"`
	MajorProses    string `json:"major_proses"`
	SubMajorProses string `json:"sub_major_proses"`
}

type MapProsesResponse struct {
	ID             int64  `json:"id"`
	IDRiskIssue    int64  `json:"id_risk_issue"`
	MegaProses     string `json:"mega_proses"`
	MajorProses    string `json:"major_proses"`
	SubMajorProses string `json:"sub_major_proses"`
}

type MapProsesResponseFinal struct {
	ID                 int64  `json:"id"`
	IDRiskIssue        int64  `json:"id_risk_issue"`
	MegaProses         string `json:"mega_proses"`
	MegaProsesDesc     string `json:"mega_proses_desc"`
	MajorProses        string `json:"major_proses"`
	MajorProsesDesc    string `json:"major_proses_desc"`
	SubMajorProses     string `json:"sub_major_proses"`
	SubMajorProsesDesc string `json:"sub_major_proses_desc"`
}

func (p MapProsesRequest) ParseRequest() MapProses {
	return MapProses{
		ID:             p.ID,
		IDRiskIssue:    p.IDRiskIssue,
		MegaProses:     p.MegaProses,
		MajorProses:    p.MajorProses,
		SubMajorProses: p.SubMajorProses,
	}
}

func (p MapProsesResponse) ParseResponse() MapProses {
	return MapProses{
		ID:             p.ID,
		IDRiskIssue:    p.IDRiskIssue,
		MegaProses:     p.MegaProses,
		MajorProses:    p.MajorProses,
		SubMajorProses: p.SubMajorProses,
	}
}

func (mp MapProsesRequest) TableName() string {
	return "risk_issue_map_proses"
}

func (mp MapProsesResponse) TableName() string {
	return "risk_issue_map_proses"
}
