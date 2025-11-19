package models

type Krid struct {
	Periode     string
	UnitKerja   string
	IdIndikator string
	Aktivitas   string
	Produk      string
	RiskIssue   string
	Content     []Content
}

type Content map[string]interface{}

type KeyRiskIndicator struct {
	// ID               string `json:"id"`
	// KeyRiskIndicator string `json:"keyRiskIndicator"`
	// Aktivitas        string `json:"aktiitas"`
	// Produk           string `json:"produk"`
	// JenisIndikator   string `json:"jenisIndikator"`
	// IndikasiRisiko   string `json:"indikasiRisiko"`
	ID                   string `json:"id"`
	KodeKeyRiskIndicator string `json:"kode_key_risk_indicator"`
	KeyRiskIndicator     string `json:"key_risk_indicator"`
	Aktifitas            string `json:"aktifitas"`
	Produk               string `json:"produk"`
	JenisIndicator       string `json:"jenis_indicator"`
	IndikasiRisiko       string `json:"indikasi_risiko"`
}

type KeywordSearch struct {
	Keyword   string `json:"keyword"`
	Aktifitas string `json:"aktifitas"`
	Produk    string `json:"produk"`
}

type KeywordSearchEdit struct {
	Keyword   string `json:"keyword"`
	Aktifitas int64  `json:"aktifitas"`
	Produk    string `json:"produk"`
}
