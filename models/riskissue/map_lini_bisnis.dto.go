package models

type MapLiniBisnisRequest struct {
	ID            int64  `json:"id"`
	IDRiskIssue   int64  `json:"id_risk_issue"`
	LiniBisnisLv1 string `json:"lini_bisnis_lv1"`
	LiniBisnisLv2 string `json:"lini_bisnis_lv2"`
	LiniBisnisLv3 string `json:"lini_bisnis_lv3"`
}

type MapLiniBisnisResponse struct {
	ID            int64  `json:"id"`
	IDRiskIssue   int64  `json:"id_risk_issue"`
	LiniBisnisLv1 string `json:"lini_bisnis_lv1"`
	LiniBisnisLv2 string `json:"lini_bisnis_lv2"`
	LiniBisnisLv3 string `json:"lini_bisnis_lv3"`
}

type MapLiniBisnisResponseFinal struct {
	ID                int64  `json:"id"`
	IDRiskIssue       int64  `json:"id_risk_issue"`
	LiniBisnisLv1     string `json:"lini_bisnis_lv1"`
	LiniBisnisLv1Desc string `json:"lini_bisnis_lv1_desc"`
	LiniBisnisLv2     string `json:"lini_bisnis_lv2"`
	LiniBisnisLv2Desc string `json:"lini_bisnis_lv2_desc"`
	LiniBisnisLv3     string `json:"lini_bisnis_lv3"`
	LiniBisnisLv3Desc string `json:"lini_bisnis_lv3_desc"`
}

func (p MapLiniBisnisRequest) ParseRequest() MapLiniBisnis {
	return MapLiniBisnis{
		ID:            p.ID,
		IDRiskIssue:   p.IDRiskIssue,
		LiniBisnisLv1: p.LiniBisnisLv1,
		LiniBisnisLv2: p.LiniBisnisLv2,
		LiniBisnisLv3: p.LiniBisnisLv3,
	}
}

func (p MapLiniBisnisRequest) ParseResponse() MapLiniBisnis {
	return MapLiniBisnis{
		ID:            p.ID,
		IDRiskIssue:   p.IDRiskIssue,
		LiniBisnisLv1: p.LiniBisnisLv1,
		LiniBisnisLv2: p.LiniBisnisLv2,
		LiniBisnisLv3: p.LiniBisnisLv3,
	}
}

func (me MapLiniBisnisRequest) TableName() string {
	return "risk_issue_map_lini_bisnis"
}

func (me MapLiniBisnisResponse) TableName() string {
	return "risk_issue_map_lini_bisnis"
}
