package models

type MapAktifitasRequest struct {
	ID           int64 `json:"id"`
	IDRiskIssue  int64 `json:"id_risk_issue"`
	Aktifitas    int64 `json:"aktifitas"`
	SubAktifitas int64 `json:"sub_aktifitas"`
}

type MapAktifitasResponse struct {
	ID           int64 `json:"id"`
	IDRiskIssue  int64 `json:"id_risk_issue"`
	Aktifitas    int64 `json:"aktifitas"`
	SubAktifitas int64 `json:"sub_aktifitas"`
}

type MapAktifitasResponseFinal struct {
	ID               int64  `json:"id"`
	IDRiskIssue      int64  `json:"id_risk_issue"`
	Aktifitas        int64  `json:"aktifitas"`
	AktifitasDesc    string `json:"aktifitas_desc"`
	SubAktifitas     int64  `json:"sub_aktifitas"`
	KodeSubAktifitas string `json:"kode_sub_aktifitas"`
	SubAktifitasDesc string `json:"sub_aktifitas_desc"`
}

func (p MapAktifitasRequest) ParseRequest() MapAktifitas {
	return MapAktifitas{
		ID:           p.ID,
		IDRiskIssue:  p.IDRiskIssue,
		Aktifitas:    p.Aktifitas,
		SubAktifitas: p.SubAktifitas,
	}
}

func (p MapAktifitasResponse) ParseResponse() MapAktifitas {
	return MapAktifitas{
		ID:           p.ID,
		IDRiskIssue:  p.IDRiskIssue,
		Aktifitas:    p.Aktifitas,
		SubAktifitas: p.SubAktifitas,
	}
}

func (ma MapAktifitasRequest) TableName() string {
	return "risk_issue_map_aktifitas"
}

func (ma MapAktifitasResponse) TableName() string {
	return "risk_issue_map_aktifitas"
}
