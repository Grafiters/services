package models

import "riskmanagement/lib"

type RiskIssueRequest struct {
	ID             int64                  `json:"id"`
	RiskTypeID     int64                  `json:"risk_type_id"`
	RiskIssueCode  string                 `json:"risk_issue_code"`
	RiskIssue      string                 `json:"risk_issue"`
	Deskripsi      string                 `json:"deskripsi"`
	KategoriRisiko string                 `json:"kategori_risiko"`
	Status         bool                   `json:"status"`
	Likelihood     *string                `json:"likelihood"`
	Impact         *string                `json:"impact"`
	DeleteFlag     bool                   `json:"delete_flag"`
	MapProses      []MapProsesRequest     `json:"map_proses"`
	MapEvent       []MapEventRequest      `json:"map_event"`
	MapKejadian    []MapKejadianRequest   `json:"map_kejadian"`
	MapProduct     []MapProductRequest    `json:"map_product"`
	MapLiniBisnis  []MapLiniBisnisRequest `json:"map_lini_bisnis"`
	MapAktifitas   []MapAktifitasRequest  `json:"map_aktifitas"`
	CreatedAt      *string                `json:"created_at"`
	UpdatedAt      *string                `json:"updated_at"`
}

type Paginate struct {
	Order  string `json:"order"`
	Sort   string `json:"sort"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Page   int    `json:"page"`
}

type RiskIssueResponse struct {
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

type RiskIssueResponseGetOne struct {
	ID             int64                        `json:"id"`
	RiskTypeID     int64                        `json:"risk_type_id"`
	RiskIssueCode  string                       `json:"risk_issue_code"`
	RiskIssue      string                       `json:"risk_issue"`
	Deskripsi      string                       `json:"deskripsi"`
	KategoriRisiko string                       `json:"kategori_risiko"`
	Status         bool                         `json:"status"`
	Likelihood     *string                      `json:"likelihood"`
	Impact         *string                      `json:"impact"`
	DeleteFlag     bool                         `json:"delete_flag"`
	MapProses      []MapProsesResponseFinal     `json:"map_proses"`
	MapEvent       []MapEventResponseFinal      `json:"map_event"`
	MapKejadian    []MapKejadianResponseFinal   `json:"map_kejadian"`
	MapProduct     []MapProductResponseFinal    `json:"map_product"`
	MapLiniBisnis  []MapLiniBisnisResponseFinal `json:"map_lini_bisnis"`
	MapAktifitas   []MapAktifitasResponseFinal  `json:"map_aktifitas"`
	MapControl     []MapControlResponseFinal    `json:"map_control"`
	MapIndicator   []MapIndicatorResponseFinal  `json:"map_indicator"`
	CreatedAt      *string                      `json:"created_at"`
	UpdatedAt      *string                      `json:"updated_at"`
}

type MappingControlRequest struct {
	ID         int64               `json:"id"`
	MapControl []MapControlRequest `json:"map_control"`
}

type MappingIndicatorRequest struct {
	ID           int64                 `json:"id"`
	MapIndicator []MapIndicatorRequest `json:"map_indicator"`
}

type Kode struct {
	KodeSubMajor string `json:"kode_sub_major"`
}

type KodeResponsNull struct {
	Kode lib.NullString `json:"kode"`
}

type KodeRespon struct {
	KodeSubMajor string `json:"kode_sub_major"`
	Kode         string `json:"kode"`
}

type KeywordRequest struct {
	Order         string `json:"order"`
	Sort          string `json:"sort"`
	Offset        int    `json:"offset"`
	Limit         int    `json:"limit"`
	Page          int    `json:"page"`
	Keyword       string `json:"keyword"`
	Aktivitas     string `json:"aktivitas"`
	SubActivityID string `json:"sub_activity_id"`
}

type RiskIssueWithoutSub struct {
	Order      string `json:"order"`
	Sort       string `json:"sort"`
	Offset     int    `json:"offset"`
	Limit      int    `json:"limit"`
	Page       int    `json:"page"`
	Keyword    string `json:"keyword"`
	ActivityID int    `json:"activity_id"`
	ProductID  int    `json:"product_id"`
}

type RiskIssueResponses struct {
	ID            int64  `json:"id"`
	RiskTypeID    int64  `json:"risk_type_id"`
	RiskIssueCode string `json:"risk_issue_code"`
	RiskIssue     string `json:"risk_issue"`
}

type RiskIssueResponsesNull struct {
	ID            lib.NullInt64  `json:"id"`
	RiskTypeID    lib.NullInt64  `json:"risk_type_id"`
	RiskIssueCode lib.NullString `json:"risk_issue_code"`
	RiskIssue     lib.NullString `json:"risk_issue"`
}

type FilterRiskIssueRequest struct {
	Order          string `json:"order"`
	Sort           string `json:"sort"`
	Offset         int    `json:"offset"`
	Limit          int    `json:"limit"`
	Page           int    `json:"page"`
	Kode           string `json:"kode"`
	RiskIssue      string `json:"risk_issue"`
	RiskTypeID     string `json:"risk_type_id"`
	KategoriRisiko string `json:"kategori_risiko"`
	Status         bool   `json:"status"`
	Product        string `json:"product"`
	LiniBisnis     string `json:"lini_bisnis"`
	Aktifitas      string `json:"aktifitas"`
	Proses         string `json:"proses"`
	EventType      string `json:"event_type"`
	Kejadian       string `json:"kejadian"`
}

type RiskIssueFilterResponses struct {
	ID             int64  `json:"id"`
	RiskTypeID     int64  `json:"risk_type_id"`
	RiskIssueCode  string `json:"risk_issue_code"`
	RiskIssue      string `json:"risk_issue"`
	KategoriRisiko string `json:"kategori_risiko"`
	Status         bool   `json:"status"`
}

type RiskIssueFilterReponsesNull struct {
	ID             lib.NullInt64  `json:"id"`
	RiskTypeID     lib.NullInt64  `json:"risk_type_id"`
	RiskIssueCode  lib.NullString `json:"risk_issue_code"`
	RiskIssue      lib.NullString `json:"risk_issue"`
	KategoriRisiko lib.NullString `json:"kategori_risiko"`
	Status         lib.NullBool   `json:"status"`
}

type RiskIssueResponseByActivity struct {
	ID            int64  `json:"id"`
	RiskIssueCode string `json:"risk_issue_code"`
	RiskIssue     string `json:"risk_issue"`
}

type RiskIssueResponseByActivityNull struct {
	ID            lib.NullInt64  `json:"id"`
	RiskIssueCode lib.NullString `json:"risk_issue_code"`
	RiskIssue     lib.NullString `json:"risk_issue"`
}

type RekomendasiMateri struct {
	ID           int64  `json:"id"`
	IDIndicator  int64  `json:"id_indicator"`
	NamaLampiran string `json:"nama_lampiran"`
	Path         string `json:"path"`
	Filename     string `json:"filename"`
}

type RekomendasiMateriNull struct {
	ID           lib.NullInt64  `json:"id"`
	IDIndicator  lib.NullInt64  `json:"id_indicator"`
	NamaLampiran lib.NullString `json:"nama_lampiran"`
	Path         lib.NullString `json:"path"`
	Filename     lib.NullString `json:"filename"`
}

type RiskIssueCode struct {
	RiskIssueCode string `json:"risk_issue_code"`
}

type ListMateri struct {
	ID            int64  `json:"id"`
	IDIndicator   int64  `json:"id_indicator"`
	NamaLampiran  string `json:"nama_lampiran"`
	NomorLampiran string `json:"nomor_lampiran"`
	JenisFile     string `json:"jenis_file"`
	Path          string `json:"path"`
	Filename      string `json:"filename"`
}

type ListMateriNull struct {
	ID            lib.NullInt64  `json:"id"`
	IDIndicator   lib.NullInt64  `json:"id_indicator"`
	NamaLampiran  lib.NullString `json:"nama_lampiran"`
	NomorLampiran lib.NullString `json:"nomor_lampiran"`
	JenisFile     lib.NullString `json:"jenis_file"`
	Path          lib.NullString `json:"path"`
	Filename      lib.NullString `json:"filename"`
}

type RiskIssueName struct {
	RiskIssue string `json:"risk_issue"`
}

/*
func (p RiskIssueRequest) ParseRequest() RiskIssue {
	return RiskIssue{
		ID:             p.ID,
		RiskTypeID:     p.RiskTypeID,
		RiskIssueCode:  p.RiskIssueCode,
		RiskIssue:      p.RiskIssue,
		Deskripsi:      p.Deskripsi,
		KategoriRisiko: p.KategoriRisiko,
		Status:         p.Status,
	}
}

func (p RiskIssueResponse) ParseRequest() RiskIssue {
	return RiskIssue{
		ID:             p.ID,
		RiskTypeID:     p.RiskTypeID,
		RiskIssueCode:  p.RiskIssueCode,
		RiskIssue:      p.RiskIssue,
		Deskripsi:      p.DeskripNullString
		KategoriRisiko: p.KategoriRisiko,
		Status:         p.Status,
		CreatedAt:      p.CreatedAt,
		UpdatedAt:      p.UpdatedAt,
	}
}
*/

func (pr RiskIssueRequest) TableName() string {
	return "risk_issue"
}

func (pr RiskIssueResponse) TableName() string {
	return "risk_issue"
}

func (pr RiskIssueFilterResponses) TableName() string {
	return "risk_issue"
}
