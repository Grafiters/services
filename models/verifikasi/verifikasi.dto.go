package models

import (
	"riskmanagement/lib"
	files "riskmanagement/models/files"
)

type VerifikasiRequest struct {
	ID                        int64                               `json:"id"`
	NoPelaporan               string                              `json:"no_pelaporan"`
	REGION                    string                              `json:"REGION"`
	RGDESC                    string                              `json:"RGDESC"`
	MAINBR                    string                              `json:"MAINBR"`
	MBDESC                    string                              `json:"MBDESC"`
	BRANCH                    string                              `json:"BRANCH"`
	BRDESC                    string                              `json:"BRDESC"`
	ActivityID                int64                               `json:"activity_id"`
	SubActivityID             int64                               `json:"sub_activity_id"`
	ProductID                 int64                               `json:"product_id"`
	RiskIssueID               int64                               `json:"risk_issue_id"`
	RiskIssue                 string                              `json:"risk_issue"`
	RiskIndicatorID           int64                               `json:"risk_indicator_id"`
	RiskIndicator             string                              `json:"risk_indicator"`
	SumberData                string                              `json:"sumber_data"`
	ApplicationID             string                              `json:"application_id"`
	HasilVerifikasi           string                              `json:"hasil_verifikasi"`
	KunjunganNasabah          *bool                               `json:"kunjungan_nasabah"`
	Perbaikan                 bool                                `json:"perbaikan"`
	IndikasiFraud             bool                                `json:"indikasi_fraud"`
	TerdapatKerugianFinansial bool                                `json:"terdapat_kerugian_finansial"`
	JenisKerugianFinansial    string                              `json:"jenis_kerugian_finansial"`
	JumlahPerkiraanKerugian   int64                               `json:"jumlah_perkiraan_kerugian"`
	JenisKerugianNonFinansial string                              `json:"jenis_kerugian_non_finansial"`
	JenisRekomendasi          string                              `json:"jenis_rekomendasi"`
	RekomendasiTindakLanjut   string                              `json:"rekomendasi_tindak_lanjut"`
	RencanaTindakLanjut       string                              `json:"rencana_tindak_lanjut"`
	RiskTypeID                int64                               `json:"risk_type_id"`
	TanggalDitemukan          *string                             `json:"tanggal_ditemukan"`
	TanggalMulaiRTL           *string                             `json:"tanggal_mulai_rtl"`
	TanggalTargetSelesai      *string                             `json:"tanggal_target_selesai"`
	MakerID                   string                              `json:"maker_id"`
	MakerDesc                 string                              `json:"maker_desc"`
	MakerDate                 *string                             `json:"maker_date"`
	LastMakerID               string                              `json:"last_maker_id"`
	LastMakerDesc             string                              `json:"last_maker_desc"`
	LastMakerDate             *string                             `json:"last_maker_date"`
	Status                    string                              `json:"status"`
	Action                    string                              `json:"action"` // create, updateApproval, updateMaintain, delete, publish, unpublish
	StatusIndikasiFraud       string                              `json:"status_indikasi_fraud"`
	ActionIndikasiFraud       string                              `json:"action_indikasi_fraud"`
	Deleted                   bool                                `json:"deleted"`
	DataAnomali               []VerifikasiAnomaliDataRequest      `json:"data_anomali"`
	DataAnomaliKRID           []VerifikasiAnomaliDataKRIDRequest  `json:"data_anomali_krid"`
	PICTindakLanjut           []VerifikasiPICTindakLanjutRequest  `json:"pic_tindak_lanjut"`
	Questionner               []VerifikasiQuestionnerRequest      `json:"questionner"`
	Files                     []files.FilesRequest                `json:"files"`
	RiskControl               []VerifikasiRiskControlRequest      `json:"risk_control"`
	PenyababKejadian          []VerifikasiPenyababKejadianRequest `json:"penyebab_kejadian"`
	AdaUsulanPerbaikan        bool                                `json:"ada_usulan_perbaikan"`
	UsulanPerbaikan           []VerifikasiUsulanPerbaikanRequest  `json:"usulan_perbaikan"`
	SampleDataTeamatik        []VerifikasiDataTematikRequest      `json:"sample_data_tematik"`
	UpdatedAt                 *string                             `json:"updated_at"`
	CreatedAt                 *string                             `json:"created_at"`
}

type VerifikasiRequestUpdateMaintain struct {
	ID                        int64   `json:"id"`
	NoPelaporan               string  `json:"no_pelaporan"`
	REGION                    string  `json:"REGION"`
	RGDESC                    string  `json:"RGDESC"`
	MAINBR                    string  `json:"MAINBR"`
	MBDESC                    string  `json:"MBDESC"`
	BRANCH                    string  `json:"BRANCH"`
	BRDESC                    string  `json:"BRDESC"`
	ActivityID                int64   `json:"activity_id"`
	SubActivityID             int64   `json:"sub_activity_id"`
	ProductID                 int64   `json:"product_id"`
	RiskIssueID               int64   `json:"risk_issue_id"`
	RiskIssue                 string  `json:"risk_issue"`
	RiskIssueOther            string  `json:"risk_issue_other"`
	RiskIndicatorID           int64   `json:"risk_indicator_id"`
	RiskIndicator             string  `json:"risk_indicator"`
	RiskIndicatorOther        string  `json:"risk_indicator_other"`
	SumberData                string  `json:"sumber_data"`
	IncidentCauseID           int64   `json:"incident_cause_id"`
	SubIncidentCauseID        int64   `json:"sub_incident_cause_id"`
	ApplicationID             string  `json:"application_id"`
	HasilVerifikasi           string  `json:"hasil_verifikasi"`
	KunjunganNasabah          bool    `json:"kunjungan_nasabah"`
	Perbaikan                 bool    `json:"perbaikan"`
	IndikasiFraud             bool    `json:"indikasi_fraud"`
	JenisKerugianFinansial    string  `json:"jenis_kerugian_finansial"`
	JumlahPerkiraanKerugian   int64   `json:"jumlah_perkiraan_kerugian"`
	JenisKerugianNonFinansial string  `json:"jenis_kerugian_non_finansial"`
	RekomendasiTindakLanjut   string  `json:"rekomendasi_tindak_lanjut"`
	RencanaTindakLanjut       string  `json:"rencana_tindak_lanjut"`
	RiskTypeID                int64   `json:"risk_type_id"`
	TanggalDitemukan          *string `json:"tanggal_ditemukan"`
	TanggalMulaiRTL           *string `json:"tanggal_mulai_rtl"`
	TanggalTargetSelesai      *string `json:"tanggal_target_selesai"`
	MakerID                   string  `json:"maker_id"`
	MakerDesc                 string  `json:"maker_desc"`
	MakerDate                 *string `json:"maker_date"`
	LastMakerID               string  `json:"last_maker_id"`
	LastMakerDesc             string  `json:"last_maker_desc"`
	LastMakerDate             *string `json:"last_maker_date"`
	Status                    string  `json:"status"`
	Action                    string  `json:"action"` // create, updateApproval, updateMaintain, delete, publish, unpublish
	Deleted                   bool    `json:"deleted"`
	UpdatedAt                 *string `json:"updated_at"`
}

type VerifikasiResponse struct {
	ID                        int64   `json:"id"`
	NoPelaporan               string  `json:"no_pelaporan"`
	REGION                    string  `json:"REGION"`
	RGDESC                    string  `json:"RGDESC"`
	MAINBR                    string  `json:"MAINBR"`
	MBDESC                    string  `json:"MBDESC"`
	BRANCH                    string  `json:"BRANCH"`
	BRDESC                    string  `json:"BRDESC"`
	ActivityID                int64   `json:"activity_id"`
	SubActivityID             int64   `json:"sub_activity_id"`
	ProductID                 int64   `json:"product_id"`
	RiskIssueID               int64   `json:"risk_issue_id"`
	RiskIssue                 string  `json:"risk_issue"`
	RiskIndicatorID           int64   `json:"risk_indicator_id"`
	RiskIndicator             string  `json:"risk_indicator"`
	SumberData                string  `json:"sumber_data"`
	ApplicationID             string  `json:"application_id"`
	HasilVerifikasi           string  `json:"hasil_verifikasi"`
	KunjunganNasabah          *bool   `json:"kunjungan_nasabah"`
	Perbaikan                 bool    `json:"perbaikan"`
	IndikasiFraud             bool    `json:"indikasi_fraud"`
	TerdapatKerugianFinansial bool    `json:"terdapat_kerugian_finansial"`
	JenisKerugianFinansial    string  `json:"jenis_kerugian_finansial"`
	JumlahPerkiraanKerugian   int64   `json:"jumlah_perkiraan_kerugian"`
	JenisKerugianNonFinansial string  `json:"jenis_kerugian_non_finansial"`
	JenisRekomendasi          string  `json:"jenis_rekomendasi"`
	RekomendasiTindakLanjut   string  `json:"rekomendasi_tindak_lanjut"`
	RencanaTindakLanjut       string  `json:"rencana_tindak_lanjut"`
	RiskTypeID                int64   `json:"risk_type_id"`
	TanggalDitemukan          *string `json:"tanggal_ditemukan"`
	TanggalMulaiRTL           *string `json:"tanggal_mulai_rtl"`
	TanggalTargetSelesai      *string `json:"tanggal_target_selesai"`
	MakerID                   string  `json:"maker_id"`
	MakerDesc                 string  `json:"maker_desc"`
	MakerDate                 *string `json:"maker_date"`
	LastMakerID               string  `json:"last_maker_id"`
	LastMakerDesc             string  `json:"last_maker_desc"`
	LastMakerDate             *string `json:"last_maker_date"`
	Status                    string  `json:"status"`
	Action                    string  `json:"action"` // create, updateApproval, updateMaintain, delete, publish, unpublish
	Deleted                   bool    `json:"deleted"`
	AdaUsulanPerbaikan        bool    `json:"ada_usulan_perbaikan"`
	UpdatedAt                 *string `json:"updated_at"`
	CreatedAt                 *string `json:"created_at"`
}

type VerifikasiList struct {
	ID            int64  `json:"id"`
	NoPelaporan   string `json:"no_pelaporan"`
	UnitKerja     string `json:"unit_kerja"`
	Aktifitas     string `json:"aktifitas"`
	IndikasiFraud string `json:"indikasi_fraud"`
	StatusVerif   string `json:"status_verif"`
	StatusRtl     string `json:"status_rtl"`
	StatusFraud   string `json:"status_fraud"`
}

type VerifikasiListFilter struct {
	ID            lib.NullInt64  `json:"id"`
	NoPelaporan   lib.NullString `json:"no_pelaporan"`
	UnitKerja     lib.NullString `json:"unit_kerja"`
	Aktifitas     lib.NullString `json:"aktifitas"`
	IndikasiFraud lib.NullString `json:"indikasi_fraud"`
	StatusVerif   lib.NullString `json:"status_verif"`
	StatusRtl     lib.NullString `json:"status_rtl"`
	StatusFraud   lib.NullString `json:"status_fraud"`
}

type VerifikasiResponseGetOne struct {
	ID                        int64                                `json:"id"`
	NoPelaporan               string                               `json:"no_pelaporan"`
	REGION                    string                               `json:"REGION"`
	RGDESC                    string                               `json:"RGDESC"`
	MAINBR                    string                               `json:"MAINBR"`
	MBDESC                    string                               `json:"MBDESC"`
	BRANCH                    string                               `json:"BRANCH"`
	BRDESC                    string                               `json:"BRDESC"`
	ActivityID                int64                                `json:"activity_id"`
	SubActivityID             int64                                `json:"sub_activity_id"`
	ProductID                 int64                                `json:"product_id"`
	RiskIssueID               int64                                `json:"risk_issue_id"`
	RiskIssue                 string                               `json:"risk_issue"`
	RiskIndicatorID           int64                                `json:"risk_indicator_id"`
	RiskIndicator             string                               `json:"risk_indicator"`
	SumberData                string                               `json:"sumber_data"`
	ApplicationID             string                               `json:"application_id"`
	HasilVerifikasi           string                               `json:"hasil_verifikasi"`
	KunjunganNasabah          *bool                                `json:"kunjungan_nasabah"`
	Perbaikan                 bool                                 `json:"perbaikan"`
	IndikasiFraud             bool                                 `json:"indikasi_fraud"`
	TerdapatKerugianFinansial bool                                 `json:"terdapat_kerugian_finansial"`
	JenisKerugianFinansial    string                               `json:"jenis_kerugian_finansial"`
	JumlahPerkiraanKerugian   int64                                `json:"jumlah_perkiraan_kerugian"`
	JenisKerugianNonFinansial string                               `json:"jenis_kerugian_non_finansial"`
	JenisRekomendasi          string                               `json:"jenis_rekomendasi"`
	RekomendasiTindakLanjut   string                               `json:"rekomendasi_tindak_lanjut"`
	RencanaTindakLanjut       string                               `json:"rencana_tindak_lanjut"`
	RiskTypeID                int64                                `json:"risk_type_id"`
	TanggalDitemukan          *string                              `json:"tanggal_ditemukan"`
	TanggalMulaiRTL           *string                              `json:"tanggal_mulai_rtl"`
	TanggalTargetSelesai      *string                              `json:"tanggal_target_selesai"`
	MakerID                   string                               `json:"maker_id"`
	MakerDesc                 string                               `json:"maker_desc"`
	MakerDate                 *string                              `json:"maker_date"`
	LastMakerID               string                               `json:"last_maker_id"`
	LastMakerDesc             string                               `json:"last_maker_desc"`
	LastMakerDate             *string                              `json:"last_maker_date"`
	Status                    string                               `json:"status"`
	Action                    string                               `json:"action"` // create, updateApproval, updateMaintain, delete, publish, unpublish
	Deleted                   bool                                 `json:"deleted"`
	DataAnomali               []VerifikasiAnomaliDataResponses     `json:"data_anomali"`
	DataAnomaliKRID           []VerifikasiAnomaliDataKRIDResponses `json:"data_anomali_krid"`
	PICTindakLanjut           []VerifikasiPICTindakLanjutResponses `json:"pic_tindak_lanjut"`
	Files                     []VerifikasiFilesResponses           `json:"files"` // Files                     []files.FilesResponses               `json:"files"`
	RiskControl               []VerifikasiRiskControlResponse      `json:"risk_control"`
	Questionner               []VerifikasiQuestionnerResponse      `json:"questionner"`
	PenyababKejadian          []VerifikasiPenyababKejadianResponse `json:"penyebab_kejadian"`
	AdaUsulanPerbaikan        bool                                 `json:"ada_usulan_perbaikan"`
	UsulanPerbaikan           []VerifikasiUsulanPerbaikanResponse  `json:"usulan_perbaikan"`
	SampleDataTeamatik        []VerifikasiDataTematikResponse      `json:"sample_data_tematik"`
	UpdatedAt                 *string                              `json:"updated_at"`
	CreatedAt                 *string                              `json:"created_at"`
}

type VerifikasiRequestMaintain struct {
	ID                        int64                               `json:"id"`
	NoPelaporan               string                              `json:"no_pelaporan"`
	REGION                    string                              `json:"REGION"`
	RGDESC                    string                              `json:"RGDESC"`
	MAINBR                    string                              `json:"MAINBR"`
	MBDESC                    string                              `json:"MBDESC"`
	BRANCH                    string                              `json:"BRANCH"`
	BRDESC                    string                              `json:"BRDESC"`
	ActivityID                int64                               `json:"activity_id"`
	SubActivityID             int64                               `json:"sub_activity_id"`
	ProductID                 int64                               `json:"product_id"`
	RiskIssueID               int64                               `json:"risk_issue_id"`
	RiskIssue                 string                              `json:"risk_issue"`
	RiskIndicatorID           int64                               `json:"risk_indicator_id"`
	RiskIndicator             string                              `json:"risk_indicator"`
	SumberData                string                              `json:"sumber_data"`
	ApplicationID             string                              `json:"application_id"`
	HasilVerifikasi           string                              `json:"hasil_verifikasi"`
	KunjunganNasabah          bool                                `json:"kunjungan_nasabah"`
	Perbaikan                 bool                                `json:"perbaikan"`
	IndikasiFraud             bool                                `json:"indikasi_fraud"`
	TerdapatKerugianFinansial bool                                `json:"terdapat_kerugian_finansial"`
	JenisKerugianFinansial    string                              `json:"jenis_kerugian_finansial"`
	JumlahPerkiraanKerugian   int64                               `json:"jumlah_perkiraan_kerugian"`
	JenisKerugianNonFinansial string                              `json:"jenis_kerugian_non_finansial"`
	JenisRekomendasi          string                              `json:"jenis_rekomendasi"`
	RekomendasiTindakLanjut   string                              `json:"rekomendasi_tindak_lanjut"`
	RencanaTindakLanjut       string                              `json:"rencana_tindak_lanjut"`
	RiskTypeID                int64                               `json:"risk_type_id"`
	TanggalDitemukan          *string                             `json:"tanggal_ditemukan"`
	TanggalMulaiRTL           *string                             `json:"tanggal_mulai_rtl"`
	TanggalTargetSelesai      *string                             `json:"tanggal_target_selesai"`
	MakerID                   string                              `json:"maker_id"`
	MakerDesc                 string                              `json:"maker_desc"`
	MakerDate                 *string                             `json:"maker_date"`
	LastMakerID               string                              `json:"last_maker_id"`
	LastMakerDesc             string                              `json:"last_maker_desc"`
	LastMakerDate             *string                             `json:"last_maker_date"`
	Status                    string                              `json:"status"`
	Action                    string                              `json:"action"` // create, updateApproval, updateMaintain, delete, publish, unpublish
	StatusIndikasiFraud       string                              `json:"status_indikasi_fraud"`
	ActionIndikasiFraud       string                              `json:"action_indikasi_fraud"`
	Deleted                   bool                                `json:"deleted"`
	DataAnomali               []VerifikasiAnomaliDataRequest      `json:"data_anomali"`
	DataAnomaliKRID           []VerifikasiAnomaliDataKRIDRequest  `json:"data_anomali_krid"`
	PICTindakLanjut           []VerifikasiPICTindakLanjutRequest  `json:"pic_tindak_lanjut"`
	Questionner               []VerifikasiQuestionnerRequest      `json:"questionner"`
	Files                     []files.FilesRequest                `json:"files"`
	RiskControl               []VerifikasiRiskControlRequest      `json:"risk_control"`
	PenyababKejadian          []VerifikasiPenyababKejadianRequest `json:"penyebab_kejadian"`
	AdaUsulanPerbaikan        bool                                `json:"ada_usulan_perbaikan"`
	UsulanPerbaikan           []VerifikasiUsulanPerbaikanRequest  `json:"usulan_perbaikan"`
	SampleDataTeamatik        []VerifikasiDataTematikRequest      `json:"sample_data_tematik"`
	UpdatedAt                 *string                             `json:"updated_at"`
	CreatedAt                 *string                             `json:"created_at"`
}

type VerifikasiFileRequest struct {
	FilesID              int64  `json:"files_id"`
	VerifikasiLampiranID int64  `json:"verifikasi_lampiran_id"`
	Path                 string `json:"path"`
}

type VerifikasiFilterRequest struct {
	Order         string `json:"order"`
	Sort          string `json:"sort"`
	Offset        int    `json:"offset"`
	Limit         int    `json:"limit"`
	Page          int    `json:"page"`
	Pernr         string `json:"pernr"`
	NoPelaporan   string `json:"no_pelaporan"`
	UnitKerja     string `json:"unit_kerja"`
	ActivityID    string `json:"activity_id"`
	RiskIssueID   string `json:"risk_issue_id"`
	StatusRtl     string `json:"status_rtl"`
	IndikasiFraud string `json:"indikasi_fraud"`
	Status        string `json:"status"`
	TglAwal       string `json:"tgl_awal"`
	TglAkhir      string `json:"tgl_akhir"`
	Branches      string `json:"branches"`
	Kostl         string `json:"kostl"`
}

type VerifikasiPagination struct {
	Order    string `json:"order"`
	Sort     string `json:"sort"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
	Page     int    `json:"page"`
	Pernr    string `json:"pernr"`
	Branches string `json:"branches"`
	Kostl    string `json:"kostl"`
}

type VerifikasiFilterReport struct {
	Order         string `json:"order"`
	Sort          string `json:"sort"`
	Offset        int    `json:"offset"`
	Limit         int    `json:"limit"`
	Page          int    `json:"page"`
	JenisLaporan  string `json:"jenis_laporan"`
	UnitKerja     string `json:"unit_kerja"`
	ActivityID    string `json:"activity_id"`
	ProductID     string `json:"product_id"`
	RiskIssueID   string `json:"risk_issue_id"`
	JudulMateri   string `json:"judul_materi"`
	Perbaikan     string `json:"perbaikan"`
	IndikasiFraud string `json:"indikasi_fraud"`
	Status        string `json:"status"`
	Periode       string `json:"periode"`
	TglAwal       string `json:"tgl_awal"`
	TglAkhir      string `json:"tgl_akhir"`
}

type VerifikasiReportResponseNull struct {
	ID          lib.NullInt64  `json:"id"`
	Tanggal     lib.NullString `json:"tanggal"`
	KodeBranch  lib.NullString `json:"kode_branch"`
	Aktifitas   lib.NullString `json:"aktifitas"`
	Produk      lib.NullString `json:"produk"`
	RiskIssue   lib.NullString `json:"risk_issue"`
	JudulMateri lib.NullString `json:"judul_materi"` //mapping riskindicator dengan riskissue
}

type VerifikasiReportResponse struct {
	ID          int64  `json:"id"`
	Tanggal     string `json:"tanggal"`
	KodeBranch  string `json:"kode_branch"`
	Aktifitas   string `json:"aktifitas"`
	Produk      string `json:"produk"`
	RiskIssue   string `json:"risk_issue"`
	JudulMateri string `json:"judul_materi"` //mapping riskindicator dengan riskissue
}

type ActivityList struct {
	ID           int64  `json:"id"`
	KodeActivity string `json:kode_activity`
	Name         string `json:name`
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

type VerifikasiLastID struct {
	ID lib.NullInt64 `json:"id"`
}

type VerifikasiLastIDResponse struct {
	ID int64 `json:"id"`
}

type VerfikasiRequestID struct {
	ID int64 `json:"id"`
}

// part 3 load questionnare
type QuestionnareResponse struct {
	Id              int64   `json:"id"`
	NamaQuistionner string  `json:"nama_questionner"`
	Pertanyaan1     *string `json:"pertanyaan1"`
	Pertanyaan2     *string `json:"pertanyaan2"`
	Pertanyaan3     *string `json:"pertanyaan3"`
	Pertanyaan4     *string `json:"pertanyaan4"`
	Pertanyaan5     *string `json:"pertanyaan5"`
	Pertanyaan6     *string `json:"pertanyaan6"`
	Pertanyaan7     *string `json:"pertanyaan7"`
	Pertanyaan8     *string `json:"pertanyaan8"`
	Pertanyaan9     *string `json:"pertanyaan9"`
}

func (v VerifikasiRequest) TableName() string {
	return "verifikasi"
}

func (v VerifikasiResponse) TableName() string {
	return "verifikasi"
}

func (v VerifikasiRequestUpdateMaintain) TableName() string {
	return "verifikasi"
}

// report
// request =============
type VerifikasiFilterReportRequest struct {
	Limit           int    `json:"limit"`
	Offset          int    `json:"offset"`
	Page            int    `json:"page"`
	ReportType      string `json:"report_type"`
	Uker            string `json:"uker"`
	REGION          string `json:"REGION"`
	MAINBR          string `json:"MAINBR"`
	BRANCH          string `json:"BRANCH"`
	Activity        string `json:"activity"`
	RiskIssue       string `json:"risk_issue"`
	RiskIndicator   string `json:"risk_indicator"`
	Product         string `json:"product"`
	Title           string `json:"title"`
	Periode         string `json:"periode"`
	FraudIndication string `json:"fraud_indication"`
	Weakness        string `json:"weakness"`
	Status          string `json:"status"`
	StartDate       string `json:"startDate"`
	EndDate         string `json:"endDate"`
	Sort            string `json:"sort"`
}

type VerifikasiReportDetailRequest struct {
	Id string `json:"id"`
}

type VerifikasiMateriRequest struct {
	Id string `json:"id"`
}

type VerifikasiDetailMateriResponse struct {
	ID           int64  `json:"id"`
	Filename     string `json:"filename"`
	NamaLampiran string `json:"nama_lampiran"`
	Path         string `json:"path"`
}

type VerifikasiDetailMateriResponseNull struct {
	ID           lib.NullInt64  `json:"id"`
	Filename     lib.NullString `json:"filename"`
	NamaLampiran lib.NullString `json:"nama_lampiran"`
	Path         lib.NullString `json:"path"`
}

type VerificationFilterReportByUkerRequest struct {
	Limit           int    `json:"limit"`
	ID              string `json:"id"`
	Offset          int    `json:"offset"`
	Page            int    `json:"page"`
	ReportType      string `json:"report_type"`
	TotalType       string `json:"total_type"`
	Uker            string `json:"uker"`
	REGION          string `json:"REGION"`
	MAINBR          string `json:"MAINBR"`
	BRANCH          string `json:"BRANCH"`
	Status          string `json:"status"`
	FraudIndication string `json:"fraud_indication"`
	RiskIssue       string `json:"risk_issue"`
	Activity        string `json:"activity"`
	Product         string `json:"product"`
	Title           string `json:"title"`
	Periode         string `json:"periode"`
	StartDate       string `json:"startDate"`
	EndDate         string `json:"endDate"`
	Sort            string `json:"sort"`
}

// data risk control
type DataRiskControlRequest struct {
	Activity      string `json:"activity"`
	RiskIssue     string `json:"risk_issue"`
	RiskIndicator string `json:"risk_indicator"`
	Product       string `json:"product"`
	Weakness      string `json:"weakness"`
	REGION        string `json:"REGION"`
	MAINBR        string `json:"MAINBR"`
	BRANCH        string `json:"BRANCH"`
	Limit         int    `json:"limit"`
	Offset        int    `json:"offset"`
	StartDate     string `json:"startDate"`
	EndDate       string `json:"endDate"`
}

type VerificationReportFilterByUkerAndTotalRequest struct {
	Limit     int    `json:"limit"`
	Offset    int    `json:"offset"`
	Page      int    `json:"page"`
	BRANCH    string `json:"BRANCH"`
	TotalType string `json:"totalType"`
	Periode   string `json:"periode"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Sort      string `json:"sort"`
}

//end of request =============

// response ================
// all ====
type VerifikasiFilterReportResponse struct {
	Id                 int64   `json:"id"`
	Code               string  `json:"code"`
	Name               string  `json:"name"`
	TotalWeakness      int64   `json:"total_weakness"`
	PercentWeakness    float64 `json:"percent_weakness"`
	TotalNonWeakness   int64   `json:"total_non_weakness"`
	PercentNonWeakness float64 `json:"percent_non_weakness"`
	GrandTotal         int64   `json:"grand_total"`
	PercentGrandTotal  float64 `json:"percent_grand_total"`
}

type VerifikasiFilterReportResponseNull struct {
	Id                 lib.NullInt64   `json:"id"`
	Code               lib.NullString  `json:"code"`
	Name               lib.NullString  `json:"name"`
	TotalWeakness      lib.NullInt64   `json:"total_weakness"`
	PercentWeakness    lib.NullFloat64 `json:"percent_weakness"`
	TotalNonWeakness   lib.NullInt64   `json:"total_non_weakness"`
	PercentNonWeakness lib.NullFloat64 `json:"percent_non_weakness"`
	GrandTotal         lib.NullInt64   `json:"grand_total"`
	PercentGrandTotal  lib.NullFloat64 `json:"percent_grand_total"`
}

type VerifikasiFilterReportResponsWithoutPercent struct {
	Id               int64  `json:"id"`
	Code             string `json:"code"`
	Name             string `json:"name"`
	TotalWeakness    int64  `json:"total_weakness"`
	TotalNonWeakness int64  `json:"total_non_weakness"`
	GrandTotal       int64  `json:"grand_total"`
}

type VerifikasiFilterReportResponsWithoutPercentNull struct {
	Id               lib.NullInt64  `json:"id"`
	Code             lib.NullString `json:"code"`
	Name             lib.NullString `json:"name"`
	TotalWeakness    lib.NullInt64  `json:"total_weakness"`
	TotalNonWeakness lib.NullInt64  `json:"total_non_weakness"`
	GrandTotal       lib.NullInt64  `json:"grand_total"`
}

// all ====

// weakness only
type VerifikasiFilterReportWeaknessOnlyResponse struct {
	Id              int64   `json:"id"`
	Code            string  `json:"code"`
	Name            string  `json:"name"`
	TotalWeakness   int64   `json:"total_weakness"`
	PercentWeakness float64 `json:"percent_weakness"`
}

type VerifikasiFilterReportWeaknessOnlyResponseNull struct {
	Id              lib.NullInt64   `json:"id"`
	Code            lib.NullString  `json:"code"`
	Name            lib.NullString  `json:"name"`
	TotalWeakness   lib.NullInt64   `json:"total_weakness"`
	PercentWeakness lib.NullFloat64 `json:"percent_weakness"`
}

type VerifikasiFilterReportWeaknessOnlyResponsWithoutPercent struct {
	Id            int64  `json:"id"`
	Code          string `json:"code"`
	Name          string `json:"name"`
	TotalWeakness int64  `json:"total_weakness"`
}

type VerifikasiFilterReportWeaknessOnlyResponsWithoutPercentNull struct {
	Id            lib.NullInt64  `json:"id"`
	Code          lib.NullString `json:"code"`
	Name          lib.NullString `json:"name"`
	TotalWeakness lib.NullInt64  `json:"total_weakness"`
}

// weakness only

// non weakness only
type VerifikasiFilterReportNonWeaknessOnlyResponse struct {
	Id                 int64   `json:"id"`
	Code               string  `json:"code"`
	Name               string  `json:"name"`
	TotalNonWeakness   int64   `json:"total_non_weakness"`
	PercentNonWeakness float64 `json:"percent_non_weakness"`
}

type VerifikasiFilterReportNonWeaknessOnlyResponseNull struct {
	Id                 lib.NullInt64   `json:"id"`
	Code               lib.NullString  `json:"code"`
	Name               lib.NullString  `json:"name"`
	TotalNonWeakness   lib.NullInt64   `json:"total_non_weakness"`
	PercentNonWeakness lib.NullFloat64 `json:"percent_non_weakness"`
}

type VerifikasiFilterReportNonWeaknessOnlyResponsWithoutPercent struct {
	Id               int64  `json:"id"`
	Code             string `json:"code"`
	Name             string `json:"name"`
	TotalNonWeakness int64  `json:"total_non_weakness"`
}

type VerifikasiFilterReportNonWeaknessOnlyResponsWithoutPercentNull struct {
	Id               lib.NullInt64  `json:"id"`
	Code             lib.NullString `json:"code"`
	Name             lib.NullString `json:"name"`
	TotalNonWeakness lib.NullInt64  `json:"total_non_weakness"`
}

// non weakness only
type VerifikasiFilterReportCompleteResponse struct {
	Id           int64  `json:"id"`
	Date         string `json:"date"`
	BRANCH       string `json:"BRANCH"`
	BRDESC       string `json:"BRDESC"`
	ActivityName string `json:"activity_name"`
	ProductName  string `json:"product_name"`
	RiskIssue    string `json:"risk_issue"`
	JudulMateri  string `json:"judul_materi"`
}

type VerifikasiFilterReportCompleteResponseNull struct {
	Id           lib.NullInt64  `json:"id"`
	Date         lib.NullString `json:"date"`
	BRANCH       lib.NullString `json:"BRANCH"`
	BRDESC       lib.NullString `json:"BRDESC"`
	ActivityName lib.NullString `json:"activity_name"`
	ProductName  lib.NullString `json:"product_name"`
	RiskIssue    lib.NullString `json:"risk_issue"`
	JudulMateri  lib.NullString `json:"judul_materi"`
}

type VerifikasiReportDetailResponseWithoutDataAnomaliNull struct {
	Id                   lib.NullInt64  `json:"id"`
	NoPelaporan          lib.NullString `json:"no_pelaporan"`
	BRANCH               lib.NullString `json:"BRANCH"`
	BRDESC               lib.NullString `json:"BRDESC"`
	MAINBR               lib.NullString `json:"MAINBR"`
	MBDESC               lib.NullString `json:"MBDESC"`
	REGION               lib.NullString `json:"REGION"`
	RGDESC               lib.NullString `json:"RGDESC"`
	ActivityName         lib.NullString `json:"activity_name"`
	SubActivityName      lib.NullString `json:"sub_activity_name"`
	ProductName          lib.NullString `json:"product_name"`
	RiskIssue            lib.NullString `json:"risk_issue"`
	RiskIndicator        lib.NullString `json:"risk_indicator"`
	IncidentCauseCode    lib.NullString `json:"incident_cause_code"`
	IncidentCauseName    lib.NullString `json:"incident_cause_name"`
	SubIncidentCauseCode lib.NullString `json:"sub_incident_cause_code"`
	SubIncidentCauseName lib.NullString `json:"sub_incident_cause_name"`
	VerificationResult   lib.NullString `json:"verification_result"`
	DataSource           lib.NullString `json:"data_source"`
	Perbaikan            lib.NullBool   `json:"perbaikan"`
	IndikasiFraud        lib.NullBool   `json:"indikasi_fraud"`
}

type VerifikasiReportDetailResponseWithoutDataAnomali struct {
	Id                   int64          `json:"id"`
	NoPelaporan          string         `json:"no_pelaporan"`
	BRANCH               string         `json:"BRANCH"`
	BRDESC               string         `json:"BRDESC"`
	ActivityName         string         `json:"activity_name"`
	SubActivityName      string         `json:"sub_activity_name"`
	ProductName          string         `json:"product_name"`
	RiskIssue            string         `json:"risk_issue"`
	RiskIndicator        string         `json:"risk_indicator"`
	IncidentCauseCode    lib.NullString `json:"incident_cause_code"`
	IncidentCauseName    lib.NullString `json:"incident_cause_name"`
	SubIncidentCauseCode lib.NullString `json:"sub_incident_cause_code"`
	SubIncidentCause     lib.NullString `json:"sub_incident_cause"`
	VerificationResult   string         `json:"verification_result"`
	DataSource           string         `json:""data_source"`
	Perbaikan            bool           `json:"perbaikan"`
	IndikasiFraud        bool           `json:"indikasi_fraud"`
}

type VerifikasiReportDetailResponseNull struct {
	Id              lib.NullInt64  `json:"id"`
	NoPelaporan     lib.NullString `json:"no_pelaporan"`
	BRANCH          lib.NullString `json:"BRANCH"`
	BRDESC          lib.NullString `json:"BRDESC"`
	ActivityName    lib.NullString `json:"activity_name"`
	SubActivityName lib.NullString `json:"sub_activity_name"`
	ProductName     lib.NullString `json:"product_name"`
	RiskIssue       lib.NullString `json:"risk_issue"`
	RiskIndicator   lib.NullString `json:"risk_indicator"`
	DataAnomali     []DataAnomali  `json:"data_anomali"`

	IncidentCauseCode lib.NullString `json:"incident_cause_code"`
	IncidentCauseName lib.NullString `json:"incident_cause_name"`

	SubIncidentCauseCode lib.NullString `json:"sub_incident_cause_code"`
	SubIncidentCauseName lib.NullString `json:"sub_incident_cause_name"`

	VerificationResult lib.NullString `json:"verification_result"`
	DataSource         lib.NullString `json:"data_source"`
	Perbaikan          lib.NullBool   `json:"perbaikan"`
	IndikasiFraud      lib.NullBool   `json:"indikasi_fraud"`
}

type VerifikasiReportDetailResponse struct {
	Id              int64  `json:"id"`
	NoPelaporan     string `json:"no_pelaporan"`
	BRANCH          string `json:"BRANCH"`
	BRDESC          string `json:"BRDESC"`
	MAINBR          string `json:"MAINBR"`
	MBDESC          string `json:"MBDESC"`
	REGION          string `json:"REGION"`
	RGDESC          string `json:"RGDESC"`
	ActivityName    string `json:"activity_name"`
	SubActivityName string `json:"sub_activity_name"`
	ProductName     string `json:"product_name"`
	RiskIssue       string `json:"risk_issue"`
	RiskIndicator   string `json:"risk_indicator"`
	// DataAnomali     []DataAnomali `json:"data_anomali"`
	DataAnomali       []VerifikasiAnomaliDataResponses     `json:"data_anomali"`
	DataAnomaliKRID   []VerifikasiAnomaliDataKRIDResponses `json:"data_anomali_krid"`
	Files             []VerifikasiFilesResponses           `json:"files"`
	IncidentCauseCode string                               `json:"incident_cause_code"`
	IncidentCauseName string                               `json:"incident_cause_name"`

	SubIncidentCauseCode string `json:"sub_incident_cause_code"`
	SubIncidentCauseName string `json:"sub_incident_cause_name"`

	VerificationResult string `json:"verification_result"`
	DataSource         string `json:"data_source"`
	Perbaikan          bool   `json:"perbaikan"`
	IndikasiFraud      bool   `json:"indikasi_fraud"`
}

type VerifikasiReportAllUker struct {
	ActivityList []ActivityList           `json:"activity_list"`
	Data         []map[string]interface{} `json:"data"`
	Colors       []Colors                 `json:"colors"`
}

// data anomali
type DataAnomali struct {
	Date       string `json"date"`
	NoRek      string `json"no_rek"`
	Nominal    string `json"nominal"`
	Keterangan string `json"keterangan"`
}

type DataAnomaliNull struct {
	Date       lib.NullString `json"date"`
	NoRek      lib.NullString `json"no_rek"`
	Nominal    lib.NullString `json"nominal"`
	Keterangan lib.NullString `json"keterangan"`
}

//

// data risk
type DataRiskIndicatorResponse struct {
	RiskControl string  `json:"risk_control"`
	Total       int64   `json:"total"`
	Percent     float64 `json:"percent"`
}

type DataRiskIndicatorResponseNull struct {
	RiskControl lib.NullString  `json:"risk_control"`
	Total       lib.NullInt64   `json:"total"`
	Percent     lib.NullFloat64 `json:"percent"`
}

type DataRiskIndicatorResponseWithNoPercent struct {
	RiskControl string `json:"risk_control"`
	Total       int64  `json:"total"`
}

type DataRiskIndicatorResponseWithNoPercentNull struct {
	RiskControl lib.NullString `json:"risk_control"`
	Total       lib.NullInt64  `json:"total"`
}

// risk indicator as materi
type GetRiskIndicatorAsMateriResponseNull struct {
	ID   lib.NullString `json:"id"`
	Code lib.NullString `json:"code"`
	Name lib.NullString `json:"name"`
}

type GetRiskIndicatorAsMateriResponse struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

// filter report by uker
type VerificationFilterReportByUkerResponseNull struct {
	REGION                   lib.NullString `json:"REGION"`
	RGDESC                   lib.NullString `json:"RGDESC"`
	MAINBR                   lib.NullString `json:"MAINBR"`
	MBDESC                   lib.NullString `json:"MBDESC"`
	BRANCH                   lib.NullString `json:"BRANCH"`
	BRDESC                   lib.NullString `json:"BRDESC"`
	TOTALVERIFICATION        lib.NullInt64  `json:"TOTALVERIFICATION"`
	TOTALBRC                 lib.NullInt64  `json:"TOTALBRC"`
	TOTALNONWEAKNESS         lib.NullInt64  `json:"TOTALNONWEAKNESS"`
	TOTALNWEAKNESS           lib.NullInt64  `json:"TOTALNWEAKNESS"`
	TOTALPERBAIKANONPROGRESS lib.NullInt64  `json:"TOTALPERBAIKANONPROGRESS"`
	TOTALPERBAIKANDONE       lib.NullInt64  `json:"TOTALPERBAIKANDONE"`
}

type VerificationFilterReportByUkerResponse struct {
	REGION                     string  `json:"REGION"`
	RGDESC                     string  `json:"RGDESC"`
	MAINBR                     string  `json:"MAINBR"`
	MBDESC                     string  `json:"MBDESC"`
	BRANCH                     string  `json:"BRANCH"`
	BRDESC                     string  `json:"BRDESC"`
	TOTALVERIFICATION          int64   `json:"TOTALVERIFICATION"`
	TOTALBRC                   int64   `json:"TOTALBRC"`
	PERCENTVERIFICATION        float64 `json:"PERCENTVERIFICATION"`
	TOTALNONWEAKNESS           int64   `json:"TOTALNONWEAKNESS"`
	TOTALNWEAKNESS             int64   `json:"TOTALNWEAKNESS"`
	PERCENTWEAKNESS            float64 `json:"PERCENTWEAKNESS"`
	TOTALPERBAIKANONPROGRESS   int64   `json:"TOTALPERBAIKANONPROGRESS"`
	TOTALPERBAIKANDONE         int64   `json:"TOTALPERBAIKANDONE"`
	PERCENTPERBAIKANONPROGRESS float64 `json:"PERCENTPERBAIKANONPROGRESS"`
}

type VerificationFilterByUkerReportCompleteResponse struct {
	Id               int64  `json:"id"`
	Date             string `json:"date"`
	BRANCH           string `json:"BRANCH"`
	BRDESC           string `json:"BRDESC"`
	Activity         string `json:"activity"`
	Product          string `json:"product"`
	RiskIssue        string `json:"risk_issue"`
	Materi           string `json:"materi"`
	IsRequiredFixing string `json:"is_required_fixing"`
	FixingStatus     string `json:"fixing_status"`
}

type VerificationFilterByUkerReportCompleteResponseNull struct {
	Id               lib.NullInt64  `json:"id"`
	Date             lib.NullString `json:"date"`
	BRANCH           lib.NullString `json:"BRANCH"`
	BRDESC           lib.NullString `json:"BRDESC"`
	Activity         lib.NullString `json:"activity"`
	Product          lib.NullString `json:"product"`
	RiskIssue        lib.NullString `json:"risk_issue"`
	Materi           lib.NullString `json:"materi"`
	IsRequiredFixing lib.NullString `json:"is_required_fixing"`
	FixingStatus     lib.NullString `json:"fixing_status"`
}

// filter by fraud indicator
type VerificationFilterReportByFraudIndicatorResponseNull struct {
	REGION     lib.NullString `json:"REGION"`
	RGDESC     lib.NullString `json:"RGDESC"`
	MAINBR     lib.NullString `json:"MAINBR"`
	MBDESC     lib.NullString `json:"MBDESC"`
	BRANCH     lib.NullString `json:"BRANCH"`
	BRDESC     lib.NullString `json:"BRDESC"`
	TOTALFRAUD lib.NullInt64  `json:"TOTALFRAUD"`
}

type VerificationFilterReportByFraudIndicatorResponse struct {
	REGION     string `json:"REGION"`
	RGDESC     string `json:"RGDESC"`
	MAINBR     string `json:"MAINBR"`
	MBDESC     string `json:"MBDESC"`
	BRANCH     string `json:"BRANCH"`
	BRDESC     string `json:"BRDESC"`
	TOTALFRAUD int64  `json:"TOTALFRAUD"`
}

type TempVerificationAllActivity struct {
	REGION   string `json:"REGION"`
	RGDESC   string `json:"RGDESC"`
	MAINBR   string `json:"MAINBR"`
	MBDESC   string `json:"MBDESC"`
	BRANCH   string `json:"BRANCH"`
	BRDESC   string `json:"BRDESC"`
	TOTAL    int64  `json:"TOTAL"`
	WEAKNESS int64  `json:"WEAKNESS"`
}

type ResponsesAllActivityComplete struct {
	RiskIssue string `json:"risk_issue"`
	WEAKNESS  int    `json:"WEAKNESS"`
	TOTAL     int    `json:"TOTAL"`
}

type Colors struct {
	Name        string `json:"name"`
	Condition   string `json:"condition"`
	BGCOLOR     string `json:"bgcolor"`
	TXCOLOR     string `json:"txcolor"`
	Description string `json:"description"`
}

type DownloadRequest struct {
	ReportId   int    `json:"report_id"`
	JSONPARAMS string `json:jsonparams`
	PERNR      string `json:pernr`
}

type GenerateInfo struct {
	NoFile     string `json:"no_file"`
	ReportId   int    `json:"report_id"`
	JSONPARAMS string `json:"json_params"`
	FILEPATH   string `json:"file_path"`
	FILEDESC   string `json:"file_desc"`
	MAKERID    string `json:"maker_id"`
	MAKERDESC  string `json:"maker_desc"`
}

// end of response ===========

// filter RptListVerifikasi 28 Apr 2023
type VerifikasiReportListRequest struct {
	Order           string `json:"order"`
	Sort            string `json:"sort"`
	Offset          int    `json:"offset"`
	Limit           int    `json:"limit"`
	Page            int    `json:"page"`
	Pernr           string `json:"pernr"`
	NoPelaporan     string `json:"no_pelaporan"`
	BrcUrc          string `json:"brc_urc"`
	BrcUrcName      string `json:"brc_urc_name"`
	REGION          string `json: REGION`
	MAINBR          string `json: MAINBR`
	BRANCH          string `json: BRANCH`
	ProductID       int64  `json:"product_id"`
	RiskIssueID     string `json:"risk_issue_id"`
	RiskIndicator   string `json:"risk_indicator"`
	RiskIndicatorID int64  `json:"risk_indicator_id"`
	IndikasiFraud   string `json:"indikasi_fraud"`
	Status          string `json:"status"`
	StartDate       string `json:"StartDate"`
	EndDate         string `json:"EndDate"`
	Timestime       string `json:"timestime"`
}

type VerifikasiReportListResponse struct {
	ID                          int64   `json:"id"`
	Periode                     string  `json:"periode"`
	RGDESC                      string  `json:"RGDESC"`
	MBDESC                      string  `json:"MBDESC"`
	BRANCH                      string  `json:"BRANCH"`
	BRDESC                      string  `json:"BRDESC"`
	NoPelaporan                 string  `json:"no_pelaporan"`
	Aktifitas                   string  `json:"aktifitas"`
	SubAktifitas                string  `json:"sub_aktifitas"`
	InformasiLain               string  `json:"informasi_lain"`
	StatusPerbaikanKonsolidasi  string  `json:"status_perbaikan_konsolidasi"`
	Maker                       string  `json:"maker"`
	RiskIssueCode               string  `json:"risk_issue_code"`
	RiskIssue                   string  `json:"risk_issue"`
	RiskIndicator               string  `json:"risk_indicator"`
	RiskControl                 string  `json:"risk_control"`
	HasilVerifikasi             string  `json:"hasil_verifikasi"`
	JumlahDataYgDiverifikasi    int     `json:"jumlah_data_yg_diverifikasi"`
	ButuhPerbaikan              string  `json:"butuh_perbaikan"`
	JumlahDataYgHarusDiperbaiki int     `json:"jumlah_data_yg_harus_diperbaiki"`
	RTLUser                     string  `json:"rtl_user"`
	StatusPerbaikanSelesai      int     `json:"status_perbaikan_selesai"`
	StatusPerbaikanProses       int     `json:"status_perbaikan_proses"`
	PresentasePerbaikan         int     `json:"presentase_perbaikan"`
	BatasWaktuPerbaikan         string  `json:"batas_waktu_perbaikan"`
	IndikasiFraud               string  `json:"indikasi_fraud"`
	Filename                    *string `json:"filename"`
	Filepath                    *string `json:"filepath"`
}

type RptRekapitulasiBCVRequest struct {
	Order     string `json:"order"`
	Sort      string `json:"sort"`
	Offset    int    `json:"offset"`
	Limit     int    `json:"limit"`
	Page      int    `json:"page"`
	REGION    string `json:"REGION"`
	MAINBR    string `json:"MAINBR"`
	BRANCH    string `json:"BRANCH"`
	BRC       string `json:"brc"`
	StartDate string `json:"StartDate"`
	EndDate   string `json:"EndDate"`
}

type RptRekapitulasiBCVResponse struct {
	Pernr   string `json:"pernr"`
	BRC     string `json:"brc"`
	BRANCH  string `json:"BRANCH"`
	BRDESC  string `json:"BRDESC"`
	MBDESC  string `json:"MBDESC"`
	RGDESC  string `json:"RGDESC"`
	BDraft  string `json:"b_draft"`
	BFinish string `json:"b_finish"`
	BTotal  string `json:"b_total"`
	CDraft  string `json:"c_draft"`
	CFinish string `json:"c_finish"`
	CTotal  string `json:"c_total"`
	VDraft  string `json:"v_draft"`
	VFinish string `json:"v_finish"`
	VTotal  string `json:"v_total"`
}

// end filter RptListVerifikasi

// Validasi Verifikasi
type ValidasiVerifikasiResponse struct {
	VerifikasiId int64   `json:"verifikasi_id"`
	NoPelaporan  string  `json:"no_pelaporan"`
	UnitKerja    string  `json:"unit_kerja"`
	Aktivitas    string  `json:"aktivitas"`
	RiskIssue    string  `json:"risk_issue"`
	MakerDesc    string  `json:"maker_desc"`
	ValidasiId   int64   `json:"validasi_id"`
	ValidatorRmc *string `json:"validator_rmc"`
	StatusRmc    *string `json:"status_rmc"`
	ValidatorRrm *string `json:"validator_rrm"`
	StatusSigner *string `json:"status_signer"`
	ValidatorOrd *string `json:"validator_ord"`
	StatusOrd    *string `json:"status_ord"`
}

type ValidasiVerifikasiRequest struct {
	Order     string `json:"order"`
	Sort      string `json:"sort"`
	Offset    int    `json:"offset"`
	Limit     int    `json:"limit"`
	Page      int    `json:"page"`
	Pn        string `json:"pn"`
	REGION    string `json: REGION`
	MAINBR    string `json: MAINBR`
	BRANCH    string `json: BRANCH`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Validator string `json:"validator"`
	TipeUker  string `json:"tipe_uker"`
}

type AcceptValidasiRequest struct {
	ID                              int64  `json:"id"`
	StatusValidasiRmc               string `json:"status_validasi_rmc"`
	TindakLanjutIndikasiFraudRmc    bool   `json:"tindak_lanjut_indikasi_fraud_rmc"`
	TindakLanjutRmc                 string `json:"tindak_lanjut_rmc"`
	CatatanRmc                      string `json:"catatan_rmc"`
	StatusValidasiSigner            string `json:"status_validasi_signer"`
	TindakLanjutIndikasiFraudSigner bool   `json:"tindak_lanjut_indikasi_fraud_signer"`
	TindakLanjutSigner              string `json:"tindak_lanjut_signer"`
	CatatanSigner                   string `json:"catatan_signer"`
	StatusValidasiOrd               string `json:"status_validasi_ord"`
	ValidasiIndikasiFraudOrd        bool   `json:"validasi_indikasi_fraud_ord"`
	TindakLanjutOrd                 string `json:"tindak_lanjut_ord"`
	CatatanOrd                      string `json:"catatan_ord"`
	ValidationBy                    string `json:"validationBy"`
}

type RejectValidasiRequest struct {
	ID                   int64  `json:"id"`
	StatusValidasiRmc    string `json:"status_validasi_rmc"`
	CatatanRmc           string `json:"catatan_rmc"`
	StatusValidasiSigner string `json:"status_validasi_signer"`
	CatatanSigner        string `json:"catatan_signer"`
	StatusValidasiOrd    string `json:"status_validasi_ord"`
	CatatanOrd           string `json:"catatan_ord"`
	RejectBy             string `json:"rejectBy"`
}

type UpdateStatusVerifikasi struct {
	ID                  int64  `json:"id"`
	IsTask              string `json:"isTask"`
	ValidationBy        string `json:"validationBy"`
	IndikasiFraud       bool   `json:"indikasi_fraud"`
	StatusIndikasiFraud string `json:"status_indikasi_fraud"`
	ActionIndikasiFraud string `json:"action_indikasi_fraud"`
}

//End Of Validasi Verifikasi

// RptRekomendasiRisk
type RptRekomendasiRiskRequest struct {
	Order     string `json:"order"`
	Sort      string `json:"sort"`
	Offset    int    `json:"offset"`
	Limit     int    `json:"limit"`
	Page      int    `json:"page"`
	JenisData string `json:"jenis_data"`
	StartDate string `json:"StartDate"`
	EndDate   string `json:"EndDate"`
}

type RptRekomendasiRiskResponse struct {
	RiskEvent     *string `json:"risk_event"`
	RiskIndicator *string `json:"risk_indicator"`
	Module        string  `json:"module"`
	Count         string  `json:"count"`
}

type ReqRtlIndikasiFraud struct {
	StartDate string `json:"StartDate"`
	EndDate   string `json:"EndDate"`
	Limit     int    `json:"limit"`
	Offset    int    `json:"offset"`
}

type RtlIndikasiFraudResponse struct {
	RtlIndikasiFraud []RtlIndikasiFraud `json:RtlIndikasiFraud`
	TotalVerifikasi  TotalVerifikasi    `json:TotalVerifikasi`
}

type RtlIndikasiFraud struct {
	REGION         string `json:"REGION"`
	RGDESC         string `json:"RGDESC"`
	MAINBR         string `json:"MAINBR"`
	MBDESC         string `json:"MBDESC"`
	BRANCH         string `json:"BRANCH"`
	BRDESC         string `json:"BRDESC"`
	INPUTBRCURC    string `json:"INPUTBRCURC"`
	INDIKASIFRAUD  string `json:"INDIKASIFRAUD"`
	VALIDASIORD    string `json:"VALIDASIORD"`
	SAAIW          string `json:"SAAIW"`
	KOORDINASIRCEO string `json:"KOORDINASIRCEO"`
}

type TotalVerifikasi struct {
	TOTALINPUTBRCURC    string `json:"TOTALINPUTBRCURC"`
	TOTALINDIKASIFRAUD  string `json:"TOTALINDIKASIFRAUD"`
	TOTALVALIDASIORD    string `json:"TOTALVALIDASIORD"`
	TOTALSAAIW          string `json:"TOTALSAAIW"`
	TOTALKOORDINASIRCEO string `json:"TOTALKOORDINASIRCEO"`
}

type ValidasiVerifikasiDetailResponse struct {
	Id                        int64   `json:"id"`
	NoPelaporan               string  `json:"no_pelaporan"`
	BRANCH                    string  `json:"BRANCH"`
	BRDESC                    string  `json:"BRDESC"`
	MAINBR                    string  `json:"MAINBR"`
	MBDESC                    string  `json:"MBDESC"`
	REGION                    string  `json:"REGION"`
	RGDESC                    string  `json:"RGDESC"`
	ActivityName              string  `json:"activity_name"`
	SubActivityName           string  `json:"sub_activity_name"`
	ProductName               string  `json:"product_name"`
	RiskIssue                 string  `json:"risk_issue"`
	RiskIndicator             string  `json:"risk_indicator"`
	VerificationResult        string  `json:"verification_result"`
	DataSource                string  `json:"data_source"`
	Perbaikan                 bool    `json:"perbaikan"`
	IndikasiFraud             bool    `json:"indikasi_fraud"`
	TerdapatKerugianFinansial bool    `json:"terdapat_kerugian_finansial"`
	JenisKerugianFinansial    string  `json:"jenis_kerugian_finansial"`
	JumlahPerkiraanKerugian   int64   `json:"jumlah_perkiraan_kerugian"`
	JenisRekomendasi          string  `json:"jenis_rekomedasi"`
	RekomendasiTindakLanjut   string  `json:"rekomendasi_tindak_lanjut"`
	RencanaTindakLanjut       string  `json:"rencana_tindak_lanjut"`
	RiskType                  string  `json:"risk_type"`
	TanggalDitemukan          *string `json:"tanggal_ditemukan"`
	TanggalMulaiRTL           *string `json:"tanggal_mulai_rtl"`
	TanggalTargetSelesai      *string `json:"tanggal_target_selesai"`
	AdaUsulanPerbaikan        bool    `json:"ada_usulan_perbaikan"`
}

type ValidasiVerifikasiDetailedResponse struct {
	Id                        int64                                      `json:"id"`
	NoPelaporan               string                                     `json:"no_pelaporan"`
	BRANCH                    string                                     `json:"BRANCH"`
	BRDESC                    string                                     `json:"BRDESC"`
	MAINBR                    string                                     `json:"MAINBR"`
	MBDESC                    string                                     `json:"MBDESC"`
	REGION                    string                                     `json:"REGION"`
	RGDESC                    string                                     `json:"RGDESC"`
	ActivityName              string                                     `json:"activity_name"`
	SubActivityName           string                                     `json:"sub_activity_name"`
	ProductName               string                                     `json:"product_name"`
	RiskIssue                 string                                     `json:"risk_issue"`
	RiskIndicator             string                                     `json:"risk_indicator"`
	VerificationResult        string                                     `json:"verification_result"`
	DataSource                string                                     `json:"data_source"`
	Perbaikan                 bool                                       `json:"perbaikan"`
	IndikasiFraud             bool                                       `json:"indikasi_fraud"`
	TerdapatKerugianFinansial bool                                       `json:"terdapat_kerugian_finansial"`
	JenisKerugianFinansial    string                                     `json:"jenis_kerugian_finansial"`
	JumlahPerkiraanKerugian   int64                                      `json:"jumlah_perkiraan_kerugian"`
	JenisRekomendasi          string                                     `json:"jenis_rekomedasi"`
	RekomendasiTindakLanjut   string                                     `json:"rekomendasi_tindak_lanjut"`
	RencanaTindakLanjut       string                                     `json:"rencana_tindak_lanjut"`
	RiskType                  string                                     `json:"risk_type"`
	TanggalDitemukan          *string                                    `json:"tanggal_ditemukan"`
	TanggalMulaiRTL           *string                                    `json:"tanggal_mulai_rtl"`
	TanggalTargetSelesai      *string                                    `json:"tanggal_target_selesai"`
	AdaUsulanPerbaikan        bool                                       `json:"ada_usulan_perbaikan"`
	VerifikasiMateriDetail    []VerifikasiDetailMateriResponse           `json:"verifikasi_materi_detail"`
	DataAnomali               []VerifikasiAnomaliDataResponses           `json:"data_anomali"`
	DataAnomaliKRID           []VerifikasiAnomaliDataKRIDResponses       `json:"data_anomali_krid"`
	PICTindakLanjut           []VerifikasiPICTindakLanjutResponses       `json:"pic_tindak_lanjut"`
	Files                     []VerifikasiFilesResponses                 `json:"files"` // Files                     []files.FilesResponses               `json:"files"`
	RiskControl               []VerifikasiRiskControlResponse            `json:"risk_control"`
	PenyababKejadian          []VerifikasiPenyababKejadianDetailResponse `json:"penyebab_kejadian"`
	UsulanPerbaikan           []VerifikasiUsulanPerbaikanResponse        `json:"usulan_perbaikan"`
}

//End of RptRekomendasiRisk

// Batch 3 Add By Panji 02/04/2024
type RTLRequest struct {
	Produk        string `json:"produk"`
	RiskEvent     string `json:"risk_event"`
	RiskIndicator string `json:"risk_indicator"`
	RiskControl   string `json:"risk_control"`
}

type RTLResponses struct {
	No                  int    `json:"no"`
	Kanwil              string `json:"kanwil"`
	Kanca               string `json:"kanca"`
	Uker                string `json:"uker"`
	Produk              string `json:"produk"`
	RiskEvent           string `json:"risk_event"`
	RiskIndicator       string `json:"risk_indicator"`
	KelemahanKontrol    string `json:"kelemahan_kontrol"`
	PenyebabKejadianLv3 string `json:"penyebab_kejadian_lv3`
	RencanaTindakLanjut string `json:"rencana_tindak_lanjut"`
}

type SummaryVerifikasiRequest struct {
	Order              string             `json:"order"`
	Sort               string             `json:"sort"`
	Offset             int                `json:"offset"`
	Limit              int                `json:"limit"`
	Page               int                `json:"page"`
	JenisPengelompokan string             `json:"jenis_pengelompokan"`
	JenisKerugian      string             `json:"jenis_kerugian"`
	REGION             string             `json:"REGION"`
	MAINBR             string             `json:"MAINBR"`
	BRANCH             string             `json:"BRANCH"`
	Kegiatan           string             `json:"kegiatan"`
	GroupData          []GroupDataRequest `json:"group_data"`
	StartDate          string             `json:"start_date"`
	EndDate            string             `json:"end_date"`
}

type FrekuensiVerifikasiRequest struct {
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

type GroupDataRequest struct {
	Column string `json:"column"`
}

type SummaryVerifikasiResponse struct {
	Aktivitas     *string `json:"aktivitas,omitempty"`
	Produk        *string `json:"produk,omitempty"`
	RiskEvent     *string `json:"risk_event,omitempty"`
	RiskIndicator *string `json:"risk_indicator,omitempty"`
	RiskControl   *string `json:"risk_control,omitempty"`
	Jumlah        int64   `json:"jumlah"`
}

type FrekuensiVerifikasiResponse struct {
	Aktivitas     *string `json:"aktivitas,omitempty"`
	Produk        *string `json:"produk,omitempty"`
	RiskEvent     *string `json:"risk_event,omitempty"`
	RiskIndicator *string `json:"risk_indicator,omitempty"`
	RiskControl   *string `json:"risk_control,omitempty"`
	Jumlah        int64   `json:"jumlah"`
}

// verfikiasFilter List

func (v VerifikasiList) TableName() string {
	return "verifikasi"
}

func (v VerifikasiFilterReportCompleteResponse) TableName() string {
	return "verifikasi"
}

func (v DataRiskIndicatorResponseWithNoPercent) TableName() string {
	return "verifikasi"
}

func (v VerificationFilterReportByFraudIndicatorResponseNull) TableName() string {
	return "dwh_branch"
}

// func (v VerificationFilterByUkerReportCompleteResponseNull) TableName() string {
// 	return "verifikasi"
// }

func (v VerificationFilterByUkerReportCompleteResponse) TableName() string {
	return "verifikasi"
}

func (v VerificationFilterReportByFraudIndicatorResponse) TableName() string {
	return "dwh_branch"
}
