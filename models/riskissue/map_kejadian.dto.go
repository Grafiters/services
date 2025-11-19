package models

type MapKejadianRequest struct {
	ID                  int64  `json:"id"`
	IDRiskIssue         int64  `json:"id_risk_issue"`
	PenyebabKejadianLv1 string `json:"penyebab_kejadian_lv1"`
	PenyebabKejadianLv2 string `json:"penyebab_kejadian_lv2"`
	PenyebabKejadianLv3 string `json:"penyebab_kejadian_lv3"`
}

type MapKejadianResponse struct {
	ID                  int64  `json:"id"`
	IDRiskIssue         int64  `json:"id_risk_issue"`
	PenyebabKejadianLv1 string `json:"penyebab_kejadian_lv1"`
	PenyebabKejadianLv2 string `json:"penyebab_kejadian_lv2"`
	PenyebabKejadianLv3 string `json:"penyebab_kejadian_lv3"`
}

type MapKejadianResponseFinal struct {
	ID                      int64  `json:"id"`
	IDRiskIssue             int64  `json:"id_risk_issue"`
	PenyebabKejadianLv1     string `json:"penyebab_kejadian_lv1"`
	PenyebabKejadianLv1Desc string `json:"penyebab_kejadian_lv1_desc"`
	PenyebabKejadianLv2     string `json:"penyebab_kejadian_lv2"`
	PenyebabKejadianLv2Desc string `json:"penyebab_kejadian_lv2_desc"`
	PenyebabKejadianLv3     string `json:"penyebab_kejadian_lv3"`
	PenyebabKejadianLv3Desc string `json:"penyebab_kejadian_lv3_desc"`
}

func (p MapKejadianRequest) ParseRequest() MapKejadian {
	return MapKejadian{
		ID:                  p.ID,
		IDRiskIssue:         p.IDRiskIssue,
		PenyebabKejadianLv1: p.PenyebabKejadianLv1,
		PenyebabKejadianLv2: p.PenyebabKejadianLv2,
		PenyebabKejadianLv3: p.PenyebabKejadianLv3,
	}
}

func (p MapKejadianRequest) ParseResponse() MapKejadian {
	return MapKejadian{
		ID:                  p.ID,
		IDRiskIssue:         p.IDRiskIssue,
		PenyebabKejadianLv1: p.PenyebabKejadianLv1,
		PenyebabKejadianLv2: p.PenyebabKejadianLv2,
		PenyebabKejadianLv3: p.PenyebabKejadianLv3,
	}
}

func (me MapKejadianRequest) TableName() string {
	return "risk_issue_map_penyebab_kejadian"
}

func (me MapKejadianResponse) TableName() string {
	return "risk_issue_map_penyebab_kejadian"
}
