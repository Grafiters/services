package models

type AuditTrail struct {
	ID          int64  `json:"id"`
	Tanggal     string `json:"tanggal"`
	PN          string `json:"pn"`
	NamaBrcUrc  string `json:"nama_brc_urc"`
	REGION      string `json:"REGION"`
	RGDESC      string `json:"RGDESC"`
	MAINBR      string `json:"MAINBR"`
	MBDESC      string `json:"MBDESC"`
	BRANCH      string `json:"BRANCH"`
	BRDESC      string `json:"BRDESC"`
	NoPelaporan string `json:"no_pelaporan"`
	Aktifitas   string `json:"aktifitas"`
	IpAddress   string `json:"ip_address"`
	Lokasi      string `json:"lokasi"`
}

type AuditTrailResponse struct {
	ID          int64  `json:"id"`
	Tanggal     string `json:"tanggal"`
	PN          string `json:"pn"`
	NamaBrcUrc  string `json:"nama_brc_urc"`
	Kanwil      string `json:"Kanwil"`
	Kanca       string `json:"Kanca"`
	Uker        string `json:"Uker"`
	NoPelaporan string `json:"no_pelaporan"`
	Aktifitas   string `json:"aktifitas"`
	IpAddress   string `json:"ip_address"`
	Lokasi      string `json:"lokasi"`
}

type FilterAudit struct {
	Order     string `json:"order"`
	Sort      string `json:"sort"`
	Offset    int    `json:"offset"`
	Limit     int    `json:"limit"`
	Page      int    `json:"page"`
	PERNR     string `json:"pn"`
	Aktifitas string `json:"aktifitas"`
	REGION    string `json:"REGION"`
	MAINBR    string `json:"MAINBR"`
	BRANCH    string `json:"BRANCH"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Timestime string `json:"timestime"`
}

func (at AuditTrail) TableName() string {
	return "audit_trail"
}
