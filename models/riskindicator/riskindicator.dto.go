package models

import (
	"riskmanagement/lib"
)

type RiskIndicatorRequest struct {
	ID                    int64                      `json:"id"`
	RiskIndicatorCode     string                     `json:"risk_indicator_code"`
	RiskIndicator         string                     `json:"risk_indicator"`
	ActivityID            int64                      `json:"activity_id"`
	ProductID             int64                      `json:"product_id"`
	Deskripsi             string                     `json:"deskripsi"`
	Satuan                string                     `json:"satuan"`
	Sifat                 string                     `json:"sifat"`
	BusinessCycleActivity string                     `json:"business_cycle_activity"`
	Batasan               string                     `json:"batasan"`
	Kondisi               string                     `json:"kondisi"`
	Type                  string                     `json:"type"`
	SLAVerifikasi         int64                      `json:"sla_verifikasi"`
	SLATindakLanjut       int64                      `json:"sla_tindak_lanjut"`
	SumberData            string                     `json:"sumber_data"`
	SumberDataText        string                     `json:"sumber_data_text"`
	PeriodePemantauan     string                     `json:"periode_pemantauan"`
	Owner                 string                     `json:"owner"`
	KPI                   string                     `json:"kpi"`
	StatusIndikator       string                     `json:"status_indikator"`
	DataSourceAnomaly     string                     `json:"data_source_anomaly"`
	Status                bool                       `json:"status"`
	LampiranIndicator     []LampiranIndicatorRequest `json:"lampiran_indicator"`
	MapThreshold          []MapThresholdRequest      `json:"map_threshold"`
	CreatedAt             *string                    `json:"created_at"`
	UpdatedAt             *string                    `json:"updated_at"`
}

type Paginate struct {
	Search    string `json:"search"`
	CreatedAt string `json:"created_at"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	Batasan   string `json:"batasan"`
	Active    bool   `json:"active"`
	Inactive  bool   `json:"inactive"`
	Order     string `json:"order"`
	Sort      string `json:"sort"`
	Offset    int    `json:"offset"`
	Limit     int    `json:"limit"`
	Page      int    `json:"page"`
}

type FilterRequest struct {
	Order     string `json:"order"`
	Sort      string `json:"sort"`
	Offset    int    `json:"offset"`
	Limit     int    `json:"limit"`
	Page      int    `json:"page"`
	Kode      string `json:"kode"`
	Indikator string `json:"indikator"`
	Status    bool   `json:"status"`
}

type UpdateDelete struct {
	ID         int64   `json:"id"`
	DeleteFlag bool    `json:"delete_flag"`
	UpdatedAt  *string `json:"updated_at"`
}

type RiskIndicatorResponse struct {
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

type RiskIndicatorGetOne struct {
	ID                    int64                       `json:"id"`
	RiskIndicatorCode     string                      `json:"risk_indicator_code"`
	RiskIndicator         string                      `json:"risk_indicator"`
	ActivityID            int64                       `json:"activity_id"`
	ProductID             int64                       `json:"product_id"`
	Deskripsi             string                      `json:"deskripsi"`
	Satuan                string                      `json:"satuan"`
	Sifat                 string                      `json:"sifat"`
	BusinessCycleActivity string                      `json:"business_cycle_activity"`
	Batasan               string                      `json:"batasan"`
	Kondisi               string                      `json:"kondisi"`
	Type                  string                      `json:"type"`
	SLAVerifikasi         int64                       `json:"sla_verifikasi"`
	SLATindakLanjut       int64                       `json:"sla_tindak_lanjut"`
	SumberData            string                      `json:"sumber_data"`
	SumberDataText        string                      `json:"sumber_data_text"`
	PeriodePemantauan     string                      `json:"periode_pemantauan"`
	Owner                 string                      `json:"owner"`
	KPI                   string                      `json:"kpi"`
	StatusIndikator       string                      `json:"status_indikator"`
	DataSourceAnomaly     string                      `json:"data_source_anomaly"`
	Status                bool                        `json:"status"`
	LampiranIndicator     []LampiranIndicatorResponse `json:"lampiran_indicator"`
	MinioLink             []MinioLink                 `json:"minio_link"`
	MapRiskIssue          []MapRiskIssueResponse      `json:"map_risk_issue"`
	CreatedAt             *string                     `json:"created_at"`
	UpdatedAt             *string                     `json:"updated_at"`
}

type ThresholdIndicator struct {
	Index            int64  `json:"index"`
	Id               string `json:"id"`
	KeyRiskIndicator string `json:"key_risk_indicator"`
	Aktivitas        string `json:"aktivitas"`
	Produk           string `json:"produk"`
	JenisIndikator   string `json:"jenis_indikator"`
	IndikasiRisiko   string `json:"indikasi_risiko"`
	Deskripsi        string `json:"deskripsi"`
	SlaVerifikasi    string `json:"sla_verifikasi"`
	SlaTl            string `json:"sla_tl"`
	RiskAwarness     string `json:"risk_awarness"`
	DataSource       string `json:"data_source"`
	Parameter        string `json:"parameter"`
	StatusIndikator  string `json:"status_indikator"`
	IsAktif          string `json:"is_aktif"`
}

type ThresholdIndicatorResponse struct {
	Index            int64                  `json:"index"`
	Id               string                 `json:"id"`
	KeyRiskIndicator string                 `json:"key_risk_indicator"`
	Aktivitas        string                 `json:"aktivitas"`
	Produk           string                 `json:"produk"`
	JenisIndikator   string                 `json:"jenis_indikator"`
	IndikasiRisiko   string                 `json:"indikasi_risiko"`
	Deskripsi        string                 `json:"deskripsi"`
	SlaVerifikasi    string                 `json:"sla_verifikasi"`
	SlaTl            string                 `json:"sla_tl"`
	RiskAwarness     string                 `json:"risk_awarness"`
	DataSource       string                 `json:"data_source"`
	Parameter        string                 `json:"parameter"`
	StatusIndikator  string                 `json:"status_indikator"`
	IsAktif          string                 `json:"is_aktif"`
	MapThreshold     []MapThresholdResponse `json:"map_threshold"`
}

type RiskIndicatorResponses struct {
	ID                int64   `json:"id"`
	RiskIndicatorCode string  `json:"risk_indicator_code"`
	RiskIndicator     string  `json:"risk_indicator"`
	ActivityID        int64   `json:"activity_id"`
	Activity          string  `json:"activity"`
	ProductID         int64   `json:"product_id"`
	Product           string  `json:"product"`
	CreatedAt         *string `json:"created_at"`
	UpdatedAt         *string `json:"updated_at"`
}

//	type SearchRequest struct {
//		Order      string `json:"order"`
//		Sort       string `json:"sort"`
//		Offset     int    `json:"offset"`
//		Limit      int    `json:"limit"`
//		Page       int    `json:"page"`
//		Keyword    string `json:"keyword"`
//		ActivityID string `json:"activity_id"`
//		ProductID  string `json:"product_id"`
//		// RiskIssueId string `json:"risk_issue_id"`
//	}
type MinioLink struct {
	Index    int64  `json:"Index"`
	Filepath string `json:"filepath"`
}

type SearchRequest struct {
	Order       string `json:"order"`
	Sort        string `json:"sort"`
	Offset      int    `json:"offset"`
	Limit       int    `json:"limit"`
	Page        int    `json:"page"`
	Keyword     string `json:"keyword"`
	ActivityID  string `json:"activity_id"`
	ProductID   string `json:"product_id"`
	RiskIssueId string `json:"risk_issue_id"`
}

type KodeResponseNull struct {
	Kode lib.NullString `json:"kode"`
}

type KodeResponse struct {
	Kode string `json:"kode"`
}

type RiskIndicatorResponsesFinal struct {
	ID                int64  `json:"id"`
	RiskIndicatorCode string `json:"risk_indicator_code"`
	RiskIndicator     string `json:"risk_indicator"`
}

type RiskIndicatorResponseNull struct {
	ID                lib.NullInt64  `json:"id"`
	RiskIndicatorCode lib.NullString `json:"risk_indicator_code"`
	RiskIndicator     lib.NullString `json:"risk_indicator"`
}

type RekomendasiMateri struct {
	ID           int64  `json:"id"`
	IDIndicator  int64  `json:"id_indicator"`
	NamaLampiran string `json:"nama_lampiran"`
	Filename     string `json:"filename"`
	Path         string `json:"path"`
}

type RekomendasiMateriNull struct {
	ID           lib.NullInt64  `json:"id"`
	IDIndicator  lib.NullInt64  `json:"id_indicator"`
	NamaLampiran lib.NullString `json:"nama_lampiran"`
	Filename     lib.NullString `json:"filename"`
	Path         lib.NullString `json:"path"`
}

type KeyRiskRequest struct {
	Order   string `json:"order"`
	Sort    string `json:"sort"`
	Offset  int    `json:"offset"`
	Limit   int    `json:"limit"`
	Page    int    `json:"page"`
	Keyword string `json:"keyword"`
}

type RiskIndicatorKRIDRequest struct {
	ID                   int64  `json:"id"`
	KodeKeyRiskIndicator int64  `json:"kode_key_risk_indicator"`
	KeyRiskIndicator     string `json:"key_risk_indicator"`
	Aktifitas            string `json:"aktifitas"`
	Produk               string `json:"produk"`
	JenisIndicator       string `json:"jenis_indicator"`
	IndikasiRisiko       string `json:"indikasi_risiko"`
}

type RiskIndicatorKRIDResponses struct {
	ID                   int64  `json:"id"`
	KodeKeyRiskIndicator int64  `json:"kode_key_risk_indicator"`
	KeyRiskIndicator     string `json:"key_risk_indicator"`
	Aktifitas            string `json:"aktifitas"`
	Produk               string `json:"produk"`
	JenisIndicator       string `json:"jenis_indicator"`
	IndikasiRisiko       string `json:"indikasi_risiko"`
}

type RiskIndicatorKRIDResponseNull struct {
	ID                   lib.NullInt64  `json:"id"`
	KodeKeyRiskIndicator lib.NullInt64  `json:"kode_key_risk_indicator"`
	KeyRiskIndicator     lib.NullString `json:"key_risk_indicator"`
	Aktifitas            lib.NullString `json:"aktifitas"`
	Produk               lib.NullString `json:"produk"`
	JenisIndicator       lib.NullString `json:"jenis_indicator"`
	IndikasiRisiko       lib.NullString `json:"indikasi_risiko"`
}

type IndicatorRequest struct {
	Aktivitas int64 `json:"aktivitas"`
	Produk    int64 `json:"produk"`
}

type IndikatorResponse struct {
	ID                int64  `json:"id"`
	RiskIndicatorCode string `json:"risk_indicator_code"`
	RiskIndicator     string `json:"risk_indicator"`
}

type IndicatorTematikResponse struct {
	Id            int64  `json:"id"`
	RiskIndicator string `json:"risk_indicator"`
	NamaTable     string `json:"nama_table"`
}

type TematikDataRequest struct {
	RiskEvent     string `json:"risk_event"`
	RiskIndicator string `json:"risk_indicator"`
	NamaTable     string `json:"nama_table"`
	PeriodeData   string `json:"periode_data"`
	UnitKerja     string `json:"unit_kerja"`
}

type TematikDataResponse struct {
	Header string   `json:"header"`
	Data   []string `json:"data"`
}

func (p RiskIndicatorRequest) ParseRequest() RiskIndicator {
	return RiskIndicator{
		ID:                p.ID,
		RiskIndicatorCode: p.RiskIndicatorCode,
		RiskIndicator:     p.RiskIndicator,
		ActivityID:        p.ActivityID,
		ProductID:         p.ProductID,
		Deskripsi:         p.Deskripsi,
		Satuan:            p.Satuan,
		Sifat:             p.Sifat,
		SLAVerifikasi:     p.SLAVerifikasi,
		SLATindakLanjut:   p.SLATindakLanjut,
		SumberData:        p.SumberData,
		PeriodePemantauan: p.PeriodePemantauan,
		Owner:             p.Owner,
		KPI:               p.KPI,
		Status:            p.Status,
	}
}

func (p RiskIndicatorResponse) ParseResponse() RiskIndicator {
	return RiskIndicator{
		ID:                p.ID,
		RiskIndicatorCode: p.RiskIndicatorCode,
		RiskIndicator:     p.RiskIndicator,
		ActivityID:        p.ActivityID,
		ProductID:         p.ProductID,
		Deskripsi:         p.Deskripsi,
		Satuan:            p.Satuan,
		Sifat:             p.Sifat,
		SLAVerifikasi:     p.SLAVerifikasi,
		SLATindakLanjut:   p.SLATindakLanjut,
		SumberData:        p.SumberData,
		PeriodePemantauan: p.PeriodePemantauan,
		Owner:             p.Owner,
		KPI:               p.KPI,
		Status:            p.Status,
		CreatedAt:         p.CreatedAt,
		UpdatedAt:         p.UpdatedAt,
	}
}

func (pr RiskIndicatorRequest) TableName() string {
	return "risk_indicator"
}

func (pr RiskIndicatorResponse) TableName() string {
	return "risk_indicator"
}

func (pr UpdateDelete) TableName() string {
	return "risk_indicator"
}

type ActivityResponse struct {
	ID           int64   `json:"id"`
	KodeActivity string  `json:"kode_activity"`
	Name         string  `json:"name"`
	CreateAt     *string `json:"create_at"`
	UpdateAt     *string `json:"update_at"`
}

func (ar ActivityResponse) TableName() string {
	return "activity"
}

type RequestMateriIfFinish struct {
	RequestId string `json:"request_id"`
}
