package models

import "riskmanagement/lib"

type BriefingRequest struct {
	ID             int64                       `json:"id"`
	NoPelaporan    string                      `json:"no_pelaporan"`
	REGION         string                      `json:"REGION"`
	RGDESC         string                      `json:"RGDESC"`
	MAINBR         string                      `json:"MAINBR"`
	MBDESC         string                      `json:"MBDESC"`
	BRANCH         string                      `json:"BRANCH"`
	BRDESC         string                      `json:"BRDESC"`
	JenisPeserta   string                      `json:"jenis_peserta"`
	JabatanPeserta string                      `json:"jabatan_peserta"`
	Peserta        []BriefingMapPesertaRequest `json:"peserta"`
	JumlahPeserta  int64                       `json:"jumlah_peserta"`
	ListPeserta    string                      `json:"list_peserta"`
	MakerID        string                      `json:"maker_id"`
	MakerDesc      string                      `json:"maker_desc"`
	MakerDate      *string                     `json:"maker_date"`
	LastMakerID    string                      `json:"last_maker_id"`
	LastMakerDesc  string                      `json:"last_maker_desc"`
	LastMakerDate  *string                     `json:"last_maker_date"`
	Status         string                      `json:"status"`
	Action         string                      `json:"action"` // create, updateApproval, updateMaintain, delete, publish, unpublish
	Deleted        bool                        `json:"deleted"`
	Materi         []BriefingMateriRequest     `json:"materi"`
}

type BriefingResponse struct {
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
	MakerID        string  `json:"maker_id"`
	MakerDesc      string  `json:"maker_desc"`
	MakerDate      *string `json:"maker_date"`
	LastMakerID    string  `json:"last_maker_id"`
	LastMakerDesc  string  `json:"last_maker_desc"`
	LastMakerDate  *string `json:"last_maker_date"`
	Status         string  `json:"status"`
	Action         string  `json:"action"` // create, updateApproval, updateMaintain, delete, publish, unpublish
	Deleted        bool    `json:"deleted"`
	CreatedAt      *string `json:"created_at"`
	UpdatedAt      *string `json:"updated_at"`
}

type BriefingResponses struct {
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
	MakerID        string  `json:"maker_id"`
	MakerDesc      string  `json:"maker_desc"`
	MakerDate      *string `json:"maker_date"`
	LastMakerID    string  `json:"last_maker_id"`
	LastMakerDesc  string  `json:"last_maker_desc"`
	LastMakerDate  *string `json:"last_maker_date"`
	Status         string  `json:"status"`
	Action         string  `json:"action"` // create, updateApproval, updateMaintain, delete, publish, unpublish
	Deleted        bool    `json:"deleted"`
}

type BriefingResponseGetOneString struct {
	ID             int64                        `json:"id"`
	NoPelaporan    string                       `json:"no_pelaporan"`
	REGION         string                       `json:"REGION"`
	RGDESC         string                       `json:"RGDESC"`
	MAINBR         string                       `json:"MAINBR"`
	MBDESC         string                       `json:"MBDESC"`
	BRANCH         string                       `json:"BRANCH"`
	BRDESC         string                       `json:"BRDESC"`
	UnitKerja      string                       `json:"unit_kerja"`
	JenisPeserta   string                       `json:"jenis_peserta"`
	JabatanPeserta string                       `json:"jabatan_peserta"`
	Peserta        []BriefingMapPesertaResponse `json:"peserta"`
	JumlahPeserta  int64                        `json:"jumlah_peserta"`
	ListPeserta    string                       `json:"list_peserta"`
	MakerID        string                       `json:"maker_id"`
	MakerDesc      string                       `json:"maker_desc"`
	MakerDate      *string                      `json:"maker_date"`
	LastMakerID    string                       `json:"last_maker_id"`
	LastMakerDesc  string                       `json:"last_maker_desc"`
	LastMakerDate  *string                      `json:"last_maker_date"`
	Status         string                       `json:"status"`
	Action         string                       `json:"action"` // create, updateApproval, updateMaintain, delete, publish, unpublish
	Deleted        bool                         `json:"deleted"`
	Materi         []BriefingMateriResponses    `json:"materi"`
	CreatedAt      *string                      `json:"created_at"`
	UpdatedAt      *string                      `json:"updated_at"`
}

type BriefingRequestUpdate struct {
	ID            int64   `json:"id"`
	LastMakerID   string  `json:"last_maker_id"`
	LastMakerDesc string  `json:"last_maker_desc"`
	LastMakerDate *string `json:"last_maker_date"`
}

type BriefingResponseMaintain struct {
	ID             int64                        `json:"id"`
	NoPelaporan    string                       `json:"no_pelaporan"`
	REGION         string                       `json:"REGION"`
	RGDESC         string                       `json:"RGDESC"`
	MAINBR         string                       `json:"MAINBR"`
	MBDESC         string                       `json:"MBDESC"`
	BRANCH         string                       `json:"BRANCH"`
	BRDESC         string                       `json:"BRDESC"`
	UnitKerja      string                       `json:"unit_kerja"`
	JenisPeserta   string                       `json:"jenis_peserta"`
	JabatanPeserta string                       `json:"jabatan_peserta"`
	Peserta        []BriefingMapPesertaResponse `json:"peserta"`
	JumlahPeserta  int64                        `json:"jumlah_peserta"`
	ListPeserta    string                       `json:"list_peserta"`
	LastMakerID    string                       `json:"last_maker_id"`
	LastMakerDesc  string                       `json:"last_maker_desc"`
	LastMakerDate  *string                      `json:"last_maker_date"`
	Status         string                       `json:"status"`
	Action         string                       `json:"action"` // create, updateApproval, updateMaintain, delete, publish, unpublish
	Deleted        bool                         `json:"deleted"`
	Materi         []BriefingMateriResponse     `json:"materi"`
	UpdatedAt      *string                      `json:"updated_at"`
}

type BriefingResponseData struct {
	ID          int64                 `json:"id"`
	NoPelaporan string                `json:"no_pelaporan"`
	UnitKerja   string                `json:"unit_kerja"`
	JudulMateri []JudulMateriBriefing `json:"judul_materi"`
	StatusBrf   string                `json:"status_brf"`
}

type BriefingResponseDataNull struct {
	ID          lib.NullInt64  `json:"id"`
	NoPelaporan lib.NullString `json:"no_pelaporan"`
	UnitKerja   lib.NullString `json:"unit_kerja"`
	// Aktifitas   lib.NullString `json:"aktifitas"`
	JudulMateri lib.NullString `json:"judul_materi"`
	StatusBrf   lib.NullString `json:"status_brf"`
}

type BriefingFilterRequest struct {
	Order       string `json:"order"`
	Sort        string `json:"sort"`
	Offset      int    `json:"offset"`
	Limit       int    `json:"limit"`
	Page        int    `json:"page"`
	Pernr       string `json:"pernr"`
	NoPelaporan string `json:"no_pelaporan"`
	UnitKerja   string `json:"unit_kerja"`
	ActivityID  string `json:"activity_id"`
	JudulMateri string `json:"judul_materi"`
	Status      string `json:"status"`
	TglAwal     string `json:"tgl_awal"`
	TglAkhir    string `json:"tgl_akhir"`
	Branches    string `json:"branches"`
	Kostl       string `json:kostl`
}

type BriefingPagination struct {
	Order    string `json:"order"`
	Sort     string `json:"sort"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
	Page     int    `json:"page"`
	Pernr    string `json:"pernr"`
	Branches string `json:"branches"`
	Kostl    string `json:"kostl"`
}

type BriefingFilterReport struct {
	Limit      int    `json:"limit"`
	Offset     int    `json:"offset"`
	Page       int    `json:"page"`
	ReportType string `json:"report_type"`
	// batch 2
	REGION string `json: REGION`
	MAINBR string `json: MAINBR`
	BRANCH string `json: BRANCH`
	// end of batch 2
	Activity  string `json:"activity"`
	Product   string `json:"product"`
	Title     string `json:"title"`
	Periode   string `json:"periode"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Sort      string `json:"sort"`
}

type BriefingFilterReportByUker struct {
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

type BriefingGetOneRequest struct {
	ID int `json:"id"`
}

type BriefingFilterReportResponse struct {
	Id    int64  `json:"id"`
	Code  string `json:"code"`
	Name  string `json:"name"`
	Total int64  `json:"total"`
}

type BriefingReportFilteredByUkerResponse struct {
	REGION          string `json:"REGION"`
	RGDESC          string `json:"RGDESC"`
	MAINBR          string `json:"MAINBR"`
	MBDESC          string `json:"MBDESC"`
	BRANCH          string `json:"BRANCH"`
	BRDESC          string `json:"BRDESC"`
	TOTALBRIEFING   string `json:"TOTALBRIEFING"`
	TOTALBRC        string `json:"TOTALBRC"`
	PERCENTBRIEFING string `json:"PERCENTBRIEFING"`
}

type BriefingFilterReportFinalResponse struct {
	Id        int64  `json:"id"`
	Date      string `json:"date"`
	BRANCH    string `json:"BRANCH"`
	BRDESC    string `json:"BRDESC"`
	Activity  string `json:"activity"`
	Product   string `json:"product"`
	RiskIssue string `json:"risk_issue"`
}

type BriefingFilterReportResponseNull struct {
	Id    lib.NullInt64  `json:"id"`
	Code  lib.NullString `json:"code"`
	Name  lib.NullString `json:"name"`
	Total lib.NullInt64  `json:"total"`
}

//reportType uker
type BriefingFilterReportByUkerResponse struct {
	REGION          string `json:"REGION"`
	RGDESC          string `json:"RGDESC"`
	MAINBR          string `json:"MAINBR"`
	MBDESC          string `json:"MBDESC"`
	BRANCH          string `json:"BRANCH"`
	BRDESC          string `json:"BRDESC"`
	TOTALBRIEFING   string `json:"TOTALBRIEFING"`
	TOTALBRC        string `json:"TOTALBRC"`
	PERCENTBRIEFING string `json:"PERCENTBRIEFING"`
}

type BriefingFilterReportFinalResponseNull struct {
	Id        lib.NullInt64  `json:"id"`
	Date      lib.NullString `json:"date"`
	BRANCH    lib.NullString `json:"BRANCH"`
	BRDESC    lib.NullString `json:"BRDESC"`
	Activity  lib.NullString `json:"activity"`
	Product   lib.NullString `json:"product"`
	RiskIssue lib.NullString `json:"risk_issue"`
}

type BriefingReportResponse struct {
	Data      []BriefingFilterReportResponse `json:"data"`
	TotalData int64                          `json:"totalData"`
}

type BriefingReportFinalResponse struct {
	Data      []BriefingFilterReportFinalResponse `json:"data"`
	TotalData int64                               `json:"totalData"`
}

type BriefingReportDetailRequest struct {
	ID string `json:"id"`
}

type BriefingReportDetail struct {
	ID            string `json:"id"`
	NoPelaporan   string `json:"no_pelaporan"`
	UnitKerja     string `json:"unit_kerja"`
	JenisPeserta  string `json:"jenis_peserta"`
	JumlahPeserta string `json:"jumlah_peserta"`
}

type BriefingReportDetailMateri struct {
	ID                int    `json:"id"`
	Activity          string `json:"activity"`
	SubActivity       string `json:"sub_activity"`
	Product           string `json:"product`
	JudulMateri       string `json:"judul_materi`
	RekomendasiMateri string `json:"rekomendasi_materi"`
	MateriTambahan    string `json:"materi_tambahan"`
}

type BriefingReportMateriRequest struct {
	ID string `json:"id`
}

type BriefingDetailMateriResponse struct {
	ID           int64  `json:"id"`
	NamaLampiran string `json:"nama_lampiran"`
	Filename     string `json:"filename"`
	Path         string `json:"path"`
}

type BriefingDetailMateriResponseNull struct {
	ID           lib.NullInt64  `json:"id"`
	NamaLampiran lib.NullString `json:"nama_lampiran"`
	Filename     lib.NullString `json:"filename"`
	Path         lib.NullString `json:"path"`
}

type BriefingReportDetailResponse struct {
	BriefingDetail  BriefingReportDetail         `json:"briefing_detail"`
	BriefingMateris []BriefingReportDetailMateri `json:"briefing_materis"`
}

//versioning 23/10/2023 add by panji
type JudulMateriBriefing struct {
	ID          int64  `json:"id"`
	JudulMateri string `json:"judul_materi"`
	Aktifitas   string `json:"aktifitas"`
}

type NoPelaporanRequest struct {
	ORGEH string `json:"ORGEH"`
}

type NoPelaporanNullResponse struct {
	NoPelaporan lib.NullString `json:"no_pelaporan"`
}

type NoPelaporanResponse struct {
	ORGEH       string `json:"ORGEH"`
	NoPelaporan string `json:"no_pelaporan"`
}

//Batch 2 Report List Briefing
type BriefingReportListRequest struct {
	Order       string `json:"order"`
	Sort        string `json:"sort"`
	Offset      int    `json:"offset"`
	Limit       int    `json:"limit"`
	Page        int    `json:"page"`
	Pernr       string `json:"pernr"`
	NoPelaporan string `json:"no_pelaporan"`
	// batch 2
	REGION string `json: REGION`
	MAINBR string `json: MAINBR`
	BRANCH string `json: BRANCH`
	// end of batch 2
	ActivityID  string `json:"activity_id"`
	JudulMateri string `json:"judul_materi"`
	Status      string `json:"status"`
	StartDate   string `json:"StartDate"`
	EndDate     string `json:"EndDate"`
	Timestime   string `json:"timestime"`
}

type MateriReportResponse struct {
	JudulMateri   string `json:"judul_materi"`
	RincianMateri string `json:"rincian_materi"`
	Aktifitas     string `json:"aktifitas"`
}

type PesertaReportResponse struct {
	Peserta string `json:"peserta"` //nama
}

type BriefingReportListFinalResponse struct {
	ID             int64  `json:"id"`
	NoPelaporan    string `json:"no_pelaporan"`
	RGDESC         string `json:"RGDESC"`
	MBDESC         string `json:"MBDESC"`
	BRANCH         string `json:"BRANCH"`
	BRDESC         string `json:"BRDESC"`
	JudulMateri    string `json:"judul_materi"`
	RiskEvent      string `json:"risk_event"`
	RincianMateri  string `json:"rincian_materi"`
	Aktivitas      string `json:"aktivitas"`
	JumlahPeserta  int64  `json:"jumlah_peserta"`
	JabatanPeserta string `json:"jabatan_peserta"`
	JenisPeserta   string `json:"jenis_peserta"`
	Peserta        string `json:"peserta"`
	MakerID        string `json:"maker_id"`
	Status         string `json:"status"`
	// Materi        []MateriReportResponse  `json:"materi"`
	// Peserta       []PesertaReportResponse `json:"peserta"`
}

type BriefingReportListResponse struct {
	ID             int64  `json:"id"`
	NoPelaporan    string `json:"no_pelaporan"`
	RGDESC         string `json:"RGDESC"`
	MBDESC         string `json:"MBDESC"`
	BRANCH         string `json:"BRANCH"`
	BRDESC         string `json:"BRDESC"`
	JudulMateri    string `json:"judul_materi"`
	RiskEvent      string `json:"risk_event"`
	RincianMateri  string `json:"rincian_materi"`
	Aktivitas      string `json:"aktivitas"`
	JumlahPeserta  int64  `json:"jumlah_peserta"`
	JabatanPeserta string `json:"jabatan_peserta"`
	JenisPeserta   string `json:"jenis_peserta"`
	Peserta        string `json:"peserta"`
	MakerID        string `json:"maker_id"`
	Status         string `json:"status"`
}

type GroupDataRequest struct {
	Column string `json:"column"`
}

type FrekuensiBriefingRequest struct {
	Order              string             `json:"order"`
	Sort               string             `json:"sort"`
	Offset             int                `json:"offset"`
	Limit              int                `json:"limit"`
	Page               int                `json:"page"`
	JenisPengelompokan string             `json:"jenis_pengelompokan"`
	REGION             string             `json:"REGION"`
	MAINBR             string             `json:"MAINBR"`
	BRANCH             string             `json:"BRANCH"`
	Kegiatan           string             `json:"kegiatan"`
	GroupData          []GroupDataRequest `json:"group_data"`
	StartDate          string             `json:"start_date"`
	EndDate            string             `json:"end_date"`
}

type FrekuensiBriefingResponse struct {
	Aktivitas *string `json:"aktivitas,omitempty"`
	Produk    *string `json:"produk,omitempty"`
	RiskEvent *string `json:"risk_event,omitempty"`
	Jumlah    int64   `json:"jumlah"`
}

func (b BriefingRequest) TableName() string {
	return "briefing"
}

func (b BriefingResponse) TableName() string {
	return "briefing"
}

func (b BriefingResponseData) TableName() string {
	return "briefing"
}

func (b BriefingFilterReportFinalResponse) TableName() string {
	return "briefing_materis"
}

//Batch 2 Report List Briefing
func (b BriefingReportListResponse) TableName() string {
	return "briefing"
}
