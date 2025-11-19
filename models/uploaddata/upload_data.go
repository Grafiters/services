package model

type RiskEvent struct {
	ID             int64
	RiskTypeID     int64
	RiskIssueCode  string
	RiskIssue      string
	Deskripsi      string
	KategoriRisiko string
	Status         bool
	Likelihood     *string
	Impact         *string
	CreatedAt      *string
	UpdatedAt      *string
	DeleteFlag     bool
}

type RiskIssueRequest struct {
	ID             int64   `json:"id"`
	RiskTypeID     int64   `json:"risk_type_id"`
	RiskIssueCode  string  `json:"risk_issue_code"`
	RiskIssue      string  `json:"risk_issue"`
	Deskripsi      string  `json:"deskripsi"`
	KategoriRisiko string  `json:"kategori_risiko"`
	Status         bool    `json:"status"`
	Likelihood     *string `json:"likelihood"`
	Impact         *string `json:"impact"`
	DeleteFlag     bool    `json:"delete_flag"`
	CreatedAt      *string `json:"created_at"`
	UpdatedAt      *string `json:"updated_at"`
}

type MapProsesRequest struct {
	ID             int64  `json:"id"`
	IDRiskIssue    int64  `json:"id_risk_issue"`
	MegaProses     string `json:"mega_proses"`
	MajorProses    string `json:"major_proses"`
	SubMajorProses string `json:"sub_major_proses"`
}

type MapEventRequest struct {
	ID           int64  `json:"id"`
	IDRiskIssue  int64  `json:"id_risk_issue"`
	EventTypeLv1 string `json:"event_type_lv1"`
	EventTypeLv2 string `json:"event_type_lv2"`
	EventTypeLv3 string `json:"event_type_lv3"`
}

type MapKejadianRequest struct {
	ID                  int64  `json:"id"`
	IDRiskIssue         int64  `json:"id_risk_issue"`
	PenyebabKejadianLv1 string `json:"penyebab_kejadian_lv1"`
	PenyebabKejadianLv2 string `json:"penyebab_kejadian_lv2"`
	PenyebabKejadianLv3 string `json:"penyebab_kejadian_lv3"`
}

type MapProductRequest struct {
	ID          int64 `json:"id"`
	IDRiskIssue int64 `json:"id_risk_issue"`
	Product     int64 `json:"product"`
}

type MapLiniBisnisRequest struct {
	ID            int64  `json:"id"`
	IDRiskIssue   int64  `json:"id_risk_issue"`
	LiniBisnisLv1 string `json:"lini_bisnis_lv1"`
	LiniBisnisLv2 string `json:"lini_bisnis_lv2"`
	LiniBisnisLv3 string `json:"lini_bisnis_lv3"`
}

type MapAktifitasRequest struct {
	ID           int64 `json:"id"`
	IDRiskIssue  int64 `json:"id_risk_issue"`
	Aktifitas    int64 `json:"aktifitas"`
	SubAktifitas int64 `json:"sub_aktifitas"`
}

type RiskControlRequest struct {
	ID          int64   `json:"id"`
	Kode        string  `json:"kode"`
	RiskControl string  `json:"risk_control"`
	ControlType string  `json:"control_type"`
	Nature      string  `json:"nature"`
	KeyControl  string  `json:"key_control"`
	Deskripsi   string  `json:"deskripsi"`
	Status      bool    `json:"status"`
	CreatedAt   *string `json:"created_at"`
	UpdatedAt   *string `json:"updated_at"`
}

type RiskIndicatorRequest struct {
	ID                int64   `json:"id"`
	RiskIndicatorCode string  `json:"risk_indicator_code"`
	RiskIndicator     string  `json:"risk_indicator"`
	ActivityID        int64   `json:"activity_id"`
	ProductID         int64   `json:"product_id"`
	Deskripsi         string  `json:"deskripsi"`
	Satuan            string  `json:"satuan"`
	Sifat             string  `json:"sifat"`
	SLAVerifikasi     int64   `json:"sla_verifikasi"`
	SLATindakLanjut   int64   `json:"sla_tindak_lanjut"`
	SumberData        string  `json:"sumber_data"`
	SumberDataText    string  `json:"sumber_data_text"`
	PeriodePemantauan string  `json:"periode_pemantauan"`
	Owner             string  `json:"owner"`
	KPI               string  `json:"kpi"`
	StatusIndikator   string  `json:"status_indikator"`
	DataSourceAnomaly string  `json:"data_source_anomaly"`
	Status            bool    `json:"status"`
	CreatedAt         *string `json:"created_at"`
	UpdatedAt         *string `json:"updated_at"`
}

type UploadControlRequest struct {
	JenisData string             `json:"jenis_data"`
	ExcelData []RiskControlExcel `json:"excel_data"`
}

type RiskControlExcel struct {
	ID          int64   `json:"id"`
	Kode        string  `json:"kode"`
	RiskControl string  `json:"risk_control"`
	ControlType string  `json:"control_type"`
	Nature      string  `json:"nature"`
	KeyControl  string  `json:"key_control"`
	Deskripsi   string  `json:"deskripsi"`
	Status      string  `json:"status"`
	CreatedAt   *string `json:"created_at"`
	UpdatedAt   *string `json:"updated_at"`
}

type UploadIndicatorRequest struct {
	JenisData string               `json:"jenis_data"`
	ExcelData []RiskIndicatorExcel `json:"excel_data"`
}

type RiskIndicatorExcel struct {
	ID                int64   `json:"id"`
	RiskIndicatorCode string  `json:"risk_indicator_code"`
	RiskIndicator     string  `json:"risk_indicator"`
	ActivityID        int64   `json:"activity_id"`
	ProductID         int64   `json:"product_id"`
	Deskripsi         string  `json:"deskripsi"`
	Satuan            string  `json:"satuan"`
	Sifat             string  `json:"sifat"`
	SLAVerifikasi     int64   `json:"sla_verifikasi"`
	SLATindakLanjut   int64   `json:"sla_tindak_lanjut"`
	SumberData        string  `json:"sumber_data"`
	SumberDataText    string  `json:"sumber_data_text"`
	PeriodePemantauan string  `json:"periode_pemantauan"`
	Owner             string  `json:"owner"`
	KPI               string  `json:"kpi"`
	StatusIndikator   string  `json:"status_indikator"`
	DataSourceAnomaly string  `json:"data_source_anomaly"`
	Status            string  `json:"status"`
	CreatedAt         *string `json:"created_at"`
	UpdatedAt         *string `json:"updated_at"`
}

type UploadRiskIssueRequest struct {
	JenisData string           `json:"jenis_data"`
	ExcelData []RiskIssueExcel `json:"excel_data"`
}

type RiskIssueExcel struct {
	ID             int64                `json:"id"`
	RiskTypeID     int64                `json:"risk_type_id"`
	RiskIssueCode  string               `json:"risk_issue_code"`
	RiskIssue      string               `json:"risk_issue"`
	Deskripsi      string               `json:"deskripsi"`
	KategoriRisiko string               `json:"kategori_risiko"`
	Status         string               `json:"status"`
	Likelihood     *string              `json:"likelihood"`
	Impact         *string              `json:"impact"`
	DeleteFlag     bool                 `json:"delete_flag"`
	MapProses      []MapProsesExcel     `json:"map_proses"`
	MapEvent       []MapEventExcel      `json:"map_event"`
	MapKejadian    []MapKejadianExcel   `json:"map_kejadian"`
	MapProduct     []MapProductExcel    `json:"map_product"`
	MapLiniBisnis  []MapLiniBisnisExcel `json:"map_lini_bisnis"`
	MapAktifitas   []MapAktifitasExcel  `json:"map_aktifitas"`
	CreatedAt      *string              `json:"created_at"`
	UpdatedAt      *string              `json:"updated_at"`
}

type MapProsesExcel struct {
	ID               int64  `json:"id"`
	IDRiskIssue      int64  `json:"id_risk_issue"`
	IDMegaProses     string `json:"id_mega_proses"`
	MegaProses       string `json:"mega_proses"`
	IDMajorProses    string `json:"id_major_proses"`
	MajorProses      string `json:"major_proses"`
	IdSubMajorProses string `json:"id_sub_major_proses"`
	SubMajorProses   string `json:"sub_major_proses"`
}

type MapEventExcel struct {
	ID             int64  `json:"id"`
	IDRiskIssue    int64  `json:"id_risk_issue"`
	IDEventTypeLv1 string `json:"id_event_type_lv1"`
	IDEventTypeLv2 string `json:"id_event_type_lv2"`
	IDEventTypeLv3 string `json:"id_event_type_lv3"`
	EventTypeLv1   string `json:"event_type_lv1"`
	EventTypeLv2   string `json:"event_type_lv2"`
	EventTypeLv3   string `json:"event_type_lv3"`
}

type MapKejadianExcel struct {
	ID                    int64  `json:"id"`
	IDRiskIssue           int64  `json:"id_risk_issue"`
	IDPenyebabKejadianLv1 string `json:"id_penyebab_kejadian_lv1"`
	IDPenyebabKejadianLv2 string `json:"id_penyebab_kejadian_lv2"`
	IDPenyebabKejadianLv3 string `json:"id_penyebab_kejadian_lv3"`
	PenyebabKejadianLv1   string `json:"penyebab_kejadian_lv1"`
	PenyebabKejadianLv2   string `json:"penyebab_kejadian_lv2"`
	PenyebabKejadianLv3   string `json:"penyebab_kejadian_lv3"`
}

type MapProductExcel struct {
	ID          int64  `json:"id"`
	IDRiskIssue int64  `json:"id_risk_issue"`
	ProductId   int64  `json:"product_id"`
	Product     string `json:"product"`
}

type MapLiniBisnisExcel struct {
	ID              int64  `json:"id"`
	IDRiskIssue     int64  `json:"id_risk_issue"`
	IdLiniBisnisLv1 string `json:"id_lini_bisnis_lv1"`
	IdLiniBisnisLv2 string `json:"id_lini_bisnis_lv2"`
	IdLiniBisnisLv3 string `json:"id_lini_bisnis_lv3"`
	LiniBisnisLv1   string `json:"lini_bisnis_lv1"`
	LiniBisnisLv2   string `json:"lini_bisnis_lv2"`
	LiniBisnisLv3   string `json:"lini_bisnis_lv3"`
}

type MapAktifitasExcel struct {
	ID            int64  `json:"id"`
	IDRiskIssue   int64  `json:"id_risk_issue"`
	ActivityID    int64  `json:"activity_id"`
	SubActivityID int64  `json:"sub_activity_id"`
	Aktifitas     string `json:"aktifitas"`
	SubAktifitas  string `json:"sub_aktifitas"`
}

func (pr RiskEvent) TableName() string {
	return "risk_issue"
}

func (pr RiskIssueRequest) TableName() string {
	return "risk_issue"
}

func (rc RiskControlRequest) TableName() string {
	return "risk_control"
}

func (pr RiskIndicatorRequest) TableName() string {
	return "risk_indicator"
}

func (pr MapProsesRequest) TableName() string {
	return "risk_issue_map_proses"
}

func (pr MapEventRequest) TableName() string {
	return "risk_issue_map_event"
}

func (pr MapKejadianRequest) TableName() string {
	return "risk_issue_map_kejadian"
}

func (pr MapProductRequest) TableName() string {
	return "risk_issue_map_product"
}

func (pr MapLiniBisnisRequest) TableName() string {
	return "risk_issue_map_lini_bisnis"
}

func (pr MapAktifitasRequest) TableName() string {
	return "risk_issue_map_aktifitas"
}
