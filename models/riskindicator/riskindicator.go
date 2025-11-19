package models

type RiskIndicator struct {
	ID                int64
	RiskIndicatorCode string
	RiskIndicator     string
	ActivityID        int64
	ProductID         int64
	Deskripsi         string
	Satuan            string
	Sifat             string
	SLAVerifikasi     int64
	SLATindakLanjut   int64
	SumberData        string
	SumberDataText    string
	PeriodePemantauan string
	Owner             string
	KPI               string
	StatusIndikator   string
	DataSourceAnomaly string
	Status            bool
	CreatedAt         *string
	UpdatedAt         *string
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
