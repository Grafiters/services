package models

type RiskIndicator struct {
	ID                    int64   `json:"id"`
	RiskIndicatorCode     string  `json:"risk_indicator_code"`
	RiskIndicator         string  `json:"risk_indicator"`
	ActivityID            int64   `json:"activity_id"`
	ProductID             int64   `json:"product_id"`
	Deskripsi             string  `json:"deskripsi"`
	Satuan                string  `json:"satuan"`
	Sifat                 string  `json:"sifat"`
	BusinessCycleActivity string  `json:"business_cycle_activity"`
	Batasan               string  `json:"batasan"`
	Kondisi               string  `json:"kondisi"`
	Type                  string  `json:"type"`
	SLAVerifikasi         int64   `json:"sla_verifikasi"`
	SLATindakLanjut       int64   `json:"sla_tindak_lanjut"`
	SumberData            string  `json:"sumber_data"`
	SumberDataText        string  `json:"sumber_data_text"`
	PeriodePemantauan     string  `json:"periode_pemantauan"`
	Owner                 string  `json:"owner"`
	KPI                   string  `json:"kpi"`
	StatusIndikator       string  `json:"status_indikator"`
	DataSourceAnomaly     string  `json:"data_source_anomaly"`
	Status                bool    `json:"status"`
	CreatedAt             *string `json:"created_at"`
	UpdatedAt             *string `json:"updated_at"`
}

type RiskIndicatorKRID struct {
	ID                   int64
	KodeKeyRiskIndicator string
	KeyRiskIndicator     string
	Aktifitas            string
	Produk               string
	JenisIndicator       string
	IndikasiRisiko       string
}

func (pr RiskIndicator) TableName() string {
	return "risk_indicator"
}

func (riKRI RiskIndicatorKRID) TableName() string {
	return "risk_indicator_krid"
}
