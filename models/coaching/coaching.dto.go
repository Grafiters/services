package models

import (
	"riskmanagement/lib"
)

type CoachingRequest struct {
	ID             int64                       `json:"id"`
	NoPelaporan    string                      `json:"no_pelaporan"`
	REGION         string                      `json:"REGION"`
	RGDESC         string                      `json:"RGDESC"`
	MAINBR         string                      `json:"MAINBR"`
	MBDESC         string                      `json:"MBDESC"`
	BRANCH         string                      `json:"BRANCH"`
	BRDESC         string                      `json:"BRDESC"`
	JenisPeserta   string                      `json:"jenis_peserta"`
	Peserta        []CoachingMapPesertaRequest `json:"peserta"`
	JabatanPeserta string                      `json:"jabatan_peserta"`
	JumlahPeserta  int64                       `json:"jumlah_peserta"`
	ListPeserta    string                      `json:"list_peserta"`
	ActivityID     int64                       `json:"activity_id"`
	SubActivityID  int64                       `json:"sub_activity_id"`
	ProductID      int64                       `json:"product_id"`
	MakerID        string                      `json:"maker_id"`
	MakerDesc      string                      `json:"maker_desc"`
	MakerDate      *string                     `json:"maker_date"`
	LastMakerID    string                      `json:"last_maker_id"`
	LastMakerDesc  string                      `json:"last_maker_desc"`
	LastMakerDate  *string                     `json:"last_maker_Date"`
	Status         string                      `json:"status"`
	Action         string                      `json:"action"` // create, updateApproval, updateMaintain, delete, publish, unpublish
	Deleted        bool                        `json:"deleted"`
	Activity       []CoachingActivityRequest   `json:"activity"`
}

type CoachingResponse struct {
	ID             int64   `json:"id"`
	NoPelaporan    string  `json:"no_pelaporan"`
	REGION         string  `json:"REGION"`
	RGDESC         string  `json:"RGDESC"`
	MAINBR         string  `json:"MAINBR"`
	MBDESC         string  `json:"MBDESC"`
	BRANCH         string  `json:"BRANCH"`
	BRDESC         string  `json:"BRDESC"`
	UnitKerja      string  `json:"unit_kerja"`
	JenisPeserta   string  `json:"jenis_peserta"`
	JabatanPeserta string  `json:"jabatan_peserta"`
	JumlahPeserta  int64   `json:"jumlah_peserta"`
	ListPeserta    string  `json:"list_peserta"`
	ActivityID     int64   `json:"activity_id"`
	Aktifitas      string  `json:"aktifitas"`
	SubActivityID  int64   `json:"sub_activity_id"`
	ProductID      int64   `json:"product_id"`
	MakerID        string  `json:"maker_id"`
	MakerDesc      string  `json:"maker_desc"`
	MakerDate      *string `json:"maker_date"`
	LastMakerID    string  `json:"last_maker_id"`
	LastMakerDesc  string  `json:"last_maker_desc"`
	LastMakerDate  *string `json:"last_maker_Date"`
	Status         string  `json:"status"`
	Action         string  `json:"action"` // create, updateApproval, updateMaintain, delete, publish, unpublish
	Deleted        bool    `json:"deleted"`
	UpdatedAt      *string `json:"updated_at"`
	CreatedAt      *string `json:"created_at"`
}

type CoachingResponses struct {
	ID             int64   `json:"id"`
	NoPelaporan    string  `json:"no_pelaporan"`
	REGION         string  `json:"REGION"`
	RGDESC         string  `json:"RGDESC"`
	MAINBR         string  `json:"MAINBR"`
	MBDESC         string  `json:"MBDESC"`
	BRANCH         string  `json:"BRANCH"`
	BRDESC         string  `json:"BRDESC"`
	UnitKerja      string  `json:"unit_kerja"`
	JenisPeserta   string  `json:"jenis_peserta"`
	JabatanPeserta string  `json:"jabatan_peserta"`
	ListPeserta    string  `json:"list_peserta"`
	JumlahPeserta  int64   `json:"jumlah_peserta"`
	ActivityID     int64   `json:"activity_id"`
	SubActivityID  int64   `json:"sub_activity_id"`
	ProductID      int64   `json:"product_id"`
	MakerID        string  `json:"maker_id"`
	MakerDesc      string  `json:"maker_desc"`
	MakerDate      *string `json:"maker_date"`
	LastMakerID    string  `json:"last_maker_id"`
	LastMakerDesc  string  `json:"last_maker_desc"`
	LastMakerDate  *string `json:"last_maker_Date"`
	Status         string  `json:"status"`
	Action         string  `json:"action"` // create, updateApproval, updateMaintain, delete, publish, unpublish
	Deleted        bool    `json:"deleted"`
}

type CoachingResponsesGetOneString struct {
	ID             int64                        `json:"id"`
	NoPelaporan    string                       `json:"no_pelaporan"`
	REGION         string                       `json:"REGION"`
	RGDESC         string                       `json:"RGDESC"`
	MAINBR         string                       `json:"MAINBR"`
	MBDESC         string                       `json:"MBDESC"`
	BRANCH         string                       `json:"BRANCH"`
	BRDESC         string                       `json:"BRDESC"`
	JenisPeserta   string                       `json:"jenis_peserta"`
	JabatanPeserta string                       `json:"jabatan_peserta"`
	ListPeserta    string                       `json:"list_peserta"`
	Peserta        []CoachingMapPesertaResponse `json:"peserta"`
	JumlahPeserta  int64                        `json:"jumlah_peserta"`
	ActivityID     int64                        `json:"activity_id"`
	SubActivityID  int64                        `json:"sub_activity_id"`
	ProductID      int64                        `json:"product_id"`
	MakerID        string                       `json:"maker_id"`
	MakerDesc      string                       `json:"maker_desc"`
	MakerDate      *string                      `json:"maker_date"`
	LastMakerID    string                       `json:"last_maker_id"`
	LastMakerDesc  string                       `json:"last_maker_desc"`
	LastMakerDate  *string                      `json:"last_maker_Date"`
	Status         string                       `json:"status"`
	Action         string                       `json:"action"` // create, updateApproval, updateMaintain, delete, publish, unpublish
	Deleted        bool                         `json:"deleted"`
	Activity       []CoachingActivityResponses  `json:"activity"`
	UpdatedAt      *string                      `json:"updated_at"`
	CreatedAt      *string                      `json:"created_at"`
}

type CoachingRequestUpdate struct {
	ID            int64   `json:"id"`
	LastMakerID   string  `json:"last_maker_id"`
	LastMakerDesc string  `json:"last_maker_desc"`
	LastMakerDate *string `json:"last_maker_Date"`
}

type CoachingResponseMaintain struct {
	ID             int64                        `json:"id"`
	NoPelaporan    string                       `json:"no_pelaporan"`
	REGION         string                       `json:"REGION"`
	RGDESC         string                       `json:"RGDESC"`
	MAINBR         string                       `json:"MAINBR"`
	MBDESC         string                       `json:"MBDESC"`
	BRANCH         string                       `json:"BRANCH"`
	BRDESC         string                       `json:"BRDESC"`
	UnitKerja      string                       `json:"unit_kerja"`
	JabatanPeserta string                       `json:"jabatan_peserta"`
	JenisPeserta   string                       `json:"jenis_peserta"`
	ListPeserta    string                       `json:"list_peserta"`
	Peserta        []CoachingMapPesertaResponse `json:"peserta"`
	JumlahPeserta  int64                        `json:"jumlah_peserta"`
	ActivityID     int64                        `json:"activity_id"`
	SubActivityID  int64                        `json:"sub_activity_id"`
	ProductID      int64                        `json:"product_id"`
	LastMakerID    string                       `json:"last_maker_id"`
	LastMakerDesc  string                       `json:"last_maker_desc"`
	LastMakerDate  *string                      `json:"last_maker_Date"`
	Status         string                       `json:"status"`
	Action         string                       `json:"action"` // create, updateApproval, updateMaintain, delete, publish, unpublish
	Deleted        bool                         `json:"deleted"`
	Activity       []CoachingActivityResponse   `json:"activity"`
	UpdatedAt      *string                      `json:"updated_at"`
}

type CoachingFilterRequest struct {
	Order       string `json:"order"`
	Sort        string `json:"sort"`
	Offset      int    `json:"offset"`
	Limit       int    `json:"limit"`
	Page        int    `json:"page"`
	Pernr       string `json:"pernr"`
	NoPelaporan string `json:"no_pelaporan"`
	UnitKerja   string `json:"unit_kerja"`
	ActivityID  string `json:"activity_id"`
	ProductID   int64  `json:"product_id"`
	RiskIssueID string `json:"risk_issue_id"`
	JudulMateri string `json:"judul_materi"`
	Status      string `json:"status"`
	TglAwal     string `json:"tgl_awal"`
	TglAkhir    string `json:"tgl_akhir"`
	Branches    string `json:"branches"`
	Kostl       string `json:kostl`
}

type CoachingPagination struct {
	Order    string `json:"order"`
	Sort     string `json:"sort"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
	Page     int    `json:"page"`
	Pernr    string `json:"pernr"`
	Branches string `json:"branches"`
	Kostl    string `json:kostl`
}

type CoachingResponseDataNull struct {
	ID          lib.NullInt64  `json:"id"`
	NoPelaporan lib.NullString `json:"no_pelaporan"`
	UnitKerja   lib.NullString `json:"unit_kerja"`
	Aktifitas   lib.NullString `json:"aktifitas"`
	RiskIssue   lib.NullString `json:"risk_issue"`
	JudulMateri lib.NullString `json:"judul_materi"`
	StatusCoach lib.NullString `json:"status_coach"`
}

type CoachingResponseData struct {
	ID          int64                 `json:"id"`
	NoPelaporan string                `json:"no_pelaporan"`
	UnitKerja   string                `json:"unit_kerja"`
	Aktifitas   string                `json:"aktifitas"`
	Materi      []JudulMateriCoaching `json:"materi"`
	StatusCoach string                `json:"status_coach"`
}

type CoachingGetOneRequest struct {
	ID int `json:"id"`
}

type NoPalaporanRequest struct {
	ORGEH string `json:"ORGEH"`
}

type NoPelaporanNullResponse struct {
	NoPelaporan lib.NullString `json:"no_pelaporan"`
}

type NoPelaporanResponse struct {
	ORGEH       string `json:"ORGEH"`
	NoPelaporan string `json:"no_pelaporan"`
}

func (c CoachingRequest) TableName() string {
	return "coaching"
}

func (c CoachingResponse) TableName() string {
	return "coaching"
}

// report
type CoachingFilterReportRequest struct {
	Limit      int    `json:"limit"`
	Offset     int    `json:"offset"`
	Page       int    `json:"page"`
	ReportType string `json:"report_type"`
	Uker       string `json:"uker"`
	REGION     string `json: REGION`
	MAINBR     string `json: MAINBR`
	BRANCH     string `json: region`
	Activity   string `json:"activity"`
	RiskIssue  string `json:"riskIssue"`
	Product    string `json:"product"`
	Title      string `json:"title"`
	Periode    string `json:"periode"`
	StartDate  string `json:"startDate"`
	EndDate    string `json:"endDate"`
	Sort       string `json:"sort"`
}

type CoachingFilterReportByUker struct {
	Limit      int    `json:"limit"`
	Offset     int    `json:"offset"`
	Page       int    `json:"page"`
	ReportType string `json:"report_type"`
	Uker       string `json:"uker"`
	REGION     string `json: REGION`
	MAINBR     string `json: MAINBR`
	BRANCH     string `json: region`
	Activity   string `json:"activity"`
	Product    string `json:"product"`
	Title      string `json:"title"`
	Periode    string `json:"periode"`
	StartDate  string `json:"startDate"`
	EndDate    string `json:"endDate"`
	Sort       string `json:"sort"`
}

type CoachingFilterReportResponse struct {
	Id    int64  `json:"id"`
	Code  string `json:"code"`
	Name  string `json:"name"`
	Total int64  `json:"total"`
}

type CoachingReportFilteredByUkerResponse struct {
	REGION          string `json:"REGION"`
	RGDESC          string `json:"RGDESC"`
	MAINBR          string `json:"MAINBR"`
	MBDESC          string `json:"MBDESC"`
	BRANCH          string `json:"BRANCH"`
	BRDESC          string `json:"BRDESC"`
	TOTALCOACHING   string `json:"TOTALCOACHING"`
	TOTALBRC        string `json:"TOTALBRC"`
	PERCENTCOACHING string `json:"PERCENTCOACHING"`
}

type CoachingFilterReportResponseNull struct {
	Id    lib.NullInt64  `json:"id"`
	Code  lib.NullString `json:"code"`
	Name  lib.NullString `json:"name"`
	Total lib.NullInt64  `json:"total"`
}

type CoachingFilterReportByUkerResponseNull struct {
	REGION          lib.NullString `json:"REGION"`
	RGDESC          lib.NullString `json:"RGDESC"`
	MAINBR          lib.NullString `json:"MAINBR"`
	MBDESC          lib.NullString `json:"MBDESC"`
	BRANCH          lib.NullString `json:"BRANCH"`
	BRDESC          lib.NullString `json:"BRDESC"`
	TOTALCOACHING   lib.NullString `json:"TOTALCOACHING"`
	TOTALBRC        lib.NullString `json:"TOTALBRC"`
	PERCENTCOACHING lib.NullString `json:"PERCENTCOACHING"`
}

type CoachingFilterReportByUkerResponse struct {
	REGION          string `json:"REGION"`
	RGDESC          string `json:"RGDESC"`
	MAINBR          string `json:"MAINBR"`
	MBDESC          string `json:"MBDESC"`
	BRANCH          string `json:"BRANCH"`
	BRDESC          string `json:"BRDESC"`
	TOTALCOACHING   string `json:"TOTALCOACHING"`
	TOTALBRC        string `json:"TOTALBRC"`
	PERCENTCOACHING string `json:"PERCENTCOACHING"`
}

type CoachingReportResponse struct {
	Data      []CoachingFilterReportResponse `json:"data"`
	TotalData int64                          `json:"totalData"`
}

// report activity by all uker
type ActivityList struct {
	Id   string `json:id`
	Code string `json:code`
	Name string `json:name`
}

type ActivityListTotal struct {
	Name  string `json:name`
	Total string `json:total`
}

type CoachingDataUker struct {
	MAINBR string `json:"MAINBR"`
	RGDESC string `json:"RGDESC"`
	MBDESC string `json:"MBDESC"`
}

type CoachingFilterByUkerAllActivityReportResponse struct {
	MAINBR   int64               `json:"MAINBR"`
	RGDESC   string              `json:"RGDESC"`
	MBDESC   string              `json:"MBDESC"`
	Activity []ActivityListTotal `json:"activity"`
}

type CoachingFilterByUkerAllActivityReportResponseNull struct {
	MAINBR   lib.NullInt64       `json:"MAINBR"`
	RGDESC   lib.NullString      `json:"RGDESC"`
	MBDESC   lib.NullString      `json:"MBDESC"`
	Activity []ActivityListTotal `json:"activity"`
}

//final

type CoachingFilterReportFinalResponse struct {
	Id        int64  `json:"id"`
	Date      string `json:"date"`
	BRANCH    string `json:"BRANCH"`
	BRDESC    string `json:"BRDESC"`
	Activity  string `json:"activity"`
	Product   string `json:"product"`
	RiskIssue string `json:"risk_issue"`
	Materi    string `json:"materi"`
}

type CoachingFilterReportFinalResponseNull struct {
	Id        lib.NullInt64  `json:"id"`
	Date      lib.NullString `json:"date"`
	BRANCH    lib.NullString `json:"BRANCH"`
	BRDESC    lib.NullString `json:"BRDESC"`
	Activity  lib.NullString `json:"activity"`
	Product   lib.NullString `json:"product"`
	RiskIssue lib.NullString `json:"risk_issue"`
	Materi    lib.NullString `json:"materi"`
}

//end of final

// detail
type CoachingReportDetailRequest struct {
	ID string `json:"id"`
}
type CoachingReportDetail struct {
	ID            string `json:"id"`
	NoPelaporan   string `json:"no_pelaporan"`
	UnitKerja     string `json:"unit_kerja"`
	JenisPeserta  string `json:"jenis_peserta"`
	JumlahPeserta string `json:"jumlah_peserta"`
}

type CoachingReportDetailMateri struct {
	ID                int    `json:"id"`
	Activity          string `json:"activity"`
	SubActivity       string `json:"sub_activity"`
	Product           string `json:"product`
	RiskIssue         string `json:"risk_issue`
	JudulMateri       string `json:"judul_materi`
	RekomendasiMateri string `json:"rekomendasi_materi"`
	MateriTambahan    string `json:"materi_tambahan"`
}

// // materi detail
type CoachingDetailMateriResponse struct {
	ID           int64          `json:"id"`
	NamaLampiran string         `json:"nama_lampiran"`
	Filename     lib.NullString `json:"filename"`
	Path         string         `json:"path"`
}

type CoachingDetailMateriResponseNull struct {
	ID           lib.NullInt64  `json:"id"`
	NamaLampiran lib.NullString `json:"nama_lampiran"`
	Filename     lib.NullString `json:"filename"`
	Path         lib.NullString `json:"path"`
}

// // end of materi detail

type CoachingReportMateriRequest struct {
	ID string `json:"id`
}

type CoachingReportDetailResponse struct {
	CoachingDetail  CoachingReportDetail         `json:"coaching_detail"`
	CoachingMateris []CoachingReportDetailMateri `json:"coaching_materis"`
}

// end of detail
// end of report

// add coachingResponseData tablename
func (c CoachingResponseData) TableName() string {
	return "coaching"
}

// Batch 2 Report List Coaching
type CoachingReportListRequest struct {
	Order       string `json:"order"`
	Sort        string `json:"sort"`
	Offset      int    `json:"offset"`
	Limit       int    `json:"limit"`
	Page        int    `json:"page"`
	Pernr       string `json:"pernr"`
	NoPelaporan string `json:"no_pelaporan"`
	REGION      string `json: REGION`
	MAINBR      string `json: MAINBR`
	BRANCH      string `json: BRANCH`
	ActivityID  string `json:"activity_id"`
	ProductID   int64  `json:"product_id"`
	RiskIssueID string `json:"risk_issue_id"`
	JudulMateri string `json:"judul_materi"`
	Status      string `json:"status"`
	StartDate   string `json:"StartDate"`
	EndDate     string `json:"EndDate"`
	Timestime   string `json:"timestime"`
}

type ActivityReportResponse struct {
	JudulMateri   string `json:"judul_materi"`
	RincianMateri string `json:"rincian_materi"`
	IsuRisiko     string `json:"isu_risiko"`
}

type PesertaReportResponse struct {
	Peserta string `json:"peserta"`
}

type CoachingReportListResponse struct {
	ID             int64  `json:"id"`
	NoPelaporan    string `json:"no_pelaporan"`
	RGDESC         string `json:"RGDESC"`
	MBDESC         string `json:"MBDESC"`
	BRANCH         string `json:"BRANCH"`
	BRDESC         string `json:"BRDESC"`
	JudulMateri    string `json:"judul_materi"`
	RincianMateri  string `json:"rincian_materi"`
	JumlahPeserta  int64  `json:"jumlah_peserta"`
	JabatanPeserta string `json:"jabatan_peserta"`
	JenisPeserta   string `json:"jenis_peserta"`
	Peserta        string `json:"peserta"`
	MakerID        string `json:"maker_id"`
	Aktifitas      string `json:"aktifitas"`
	SubAktifitas   string `json:"sub_aktifitas"`
	IsuRisiko      string `json:"isu_risiko"`
	RiskIndicator  string `json:"risk_indicator"`
	Status         string `json:"status"`
}

type CoachingReportListFinalResponse struct {
	NoPelaporan    string `json:"no_pelaporan"`
	RGDESC         string `json:"RGDESC"`
	MBDESC         string `json:"MBDESC"`
	BRANCH         string `json:"BRANCH"`
	BRDESC         string `json:"BRDESC"`
	JudulMateri    string `json:"judul_materi"`
	RincianMateri  string `json:"rincian_materi"`
	JumlahPeserta  int64  `json:"jumlah_peserta"`
	JabatanPeserta string `json:"jabatan_peserta"`
	JenisPeserta   string `json:"jenis_peserta"`
	Peserta        string `json:"peserta"`
	Aktifitas      string `json:"aktifitas"`
	SubAktifitas   string `json:"sub_aktifitas"`
	IsuRisiko      string `json:"isu_risiko"`
	RiskIndicator  string `json:"risk_indicator"`
	MakerID        string `json:"maker_id"`
	Status         string `json:"status"`
}

func (c CoachingReportListResponse) TableName() string {
	return "coaching"
}

// versioning 23/10/2023 by panji
type JudulMateriCoaching struct {
	ID          int64  `json:"id"`
	RiskIssue   string `json:"risk_issue"`
	JudulMateri string `json:"judul_materi"`
}

type FrekuensiCoachingRequest struct {
	Order              string `json:"order"`
	Sort               string `json:"sort"`
	Offset             int    `json:"offset"`
	Limit              int    `json:"limit"`
	Page               int    `json:"page"`
	JenisPengelompokan string `json:"jenis_pengelompokan"`

	REGION string `json:"REGION"`
	MAINBR string `json:"MAINBR"`
	BRANCH string `json:"BRANCH"`

	Kegiatan  string             `json:"kegiatan"`
	GroupData []GroupDataRequest `json:"group_data"`
	StartDate string             `json:"start_date"`
	EndDate   string             `json:"end_date"`
}

type FrekuensiCoachingResponse struct {
	Aktivitas     *string `json:"aktivitas,omitempty"`
	Produk        *string `json:"produk,omitempty"`
	RiskEvent     *string `json:"risk_event,omitempty"`
	RiskIndicator *string `json:"risk_indicator,omitempty"`
	Jumlah        int64   `json:"jumlah"`
}

type GroupDataRequest struct {
	Column string `json:"column"`
}
