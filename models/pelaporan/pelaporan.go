package models

type SuratDraftRequest struct {
	PnPenerima string `json:"pn_penerima"`
	PnMaker    string `json:"pn_maker"`
	PnSigner   string `json:"pn_signer"`
	Perihal    string `json:"perihal"`
	KodeKanca  string `json:"kode_kanca"`
	Kanca      string `json:"kanca"`
	KodeUker   string `json:"kode_uker"`
	Uker       string `json:"uker"`
	Semester   string `json:"semester"`
	Tahun      string `json:"tahun"`
}

type DraftSuratRequest struct {
	Kepentingan string      `json:"kepentingan"`
	Kerahasiaan string      `json:"kerahasiaan"`
	PnMaker     string      `json:"pn_maker"`
	NamaMaker   string      `json:"nama_maker"`
	Perihal     string      `json:"perihal"`
	Uker        []SuratUker `json:"uker"`
	Semester    string      `json:"semester"`
	Tahun       int64       `json:"tahun"`
	Signer      []Signer    `json:"signer"`
}

type ApprovalRequest struct {
	ID         int64  `json:"id"`
	PnApproval string `json:"pn_approval"`
}

type ApproveUpdate struct {
	ID             int64  `json:"id"`
	PosisiApprover string `gorm:"column:posisiApprover" json:"posisiApprover"`
	StatusMCS      string `gorm:"column:statusMCS" json:"statusMCS"`
	Status         string `gorm:"column:status" json:"status"`
	ResponseCode   string `gorm:"column:responseCode" json:"response_code"`
	ResponseStatus string `gorm:"column:responseStatus" json:"response_status"`
}
type PenolakanCatatan struct {
	IDPelaporan  int64   `json:"id_pelaporan"`
	Penolak      string  `json:"penolak"`
	Catatan      string  `json:"catatan"`
	TanggalTolak *string `json:"tanggal_tolak"`
}

type Signer struct {
	PnSigner   string `json:"pn_signer"`
	NamaSigner string `json:"nama_signer"`
	Tempat     string `json:"tempat"`
	Jabatan    string `json:"jabatan"`
}

type SuratUker struct {
	REGION string `json:"region"`
	MAINBR string `json:"mainbr"`
	MBDESC string `json:"mbdesc"`
	BRANCH string `json:"branch"`
	BRDESC string `json:"brdesc"`
}

type PelaporanDraft struct {
	IdTemplate          string  `gorm:"column:idTemplate" json:"idTemplate"`
	JenisSurat          string  `gorm:"column:jenisSurat" json:"jenisSurat"`
	KodeUnikSurat       string  `gorm:"column:kodeUnikSurat" json:"kodeUnikSurat"`
	BranchCodePenerima  string  `gorm:"column:branchCodePenerima" json:"branchCodePenerima"`
	OrgehPenerima       string  `gorm:"column:orgehPenerima" json:"orgehPenerima"`
	BranchCodeTindasan  string  `gorm:"column:branchCodeTindasan" json:"branchCodeTindasan"`
	OrgehTindasan       string  `gorm:"column:orgehTindasan" json:"orgehTindasan"`
	PnPenerima          string  `gorm:"column:pnPenerima" json:"pnPenerima"`
	PnTindasan          string  `gorm:"column:pnTindasan" json:"pnTindasan"`
	KodeSurat           string  `gorm:"column:kodeSurat" json:"kodeSurat"`
	Kerahasiaan         string  `gorm:"column:kerahasiaan" json:"kerahasiaan"`
	Kesegeraan          string  `gorm:"column:kesegeraan" json:"kesegeraan"`
	KepadaYth           string  `gorm:"column:kepadaYth" json:"kepadaYth"`
	Perihal             string  `gorm:"column:perihal" json:"perihal"`
	Semester            string  `gorm:"column:semester" json:"semester"`
	Tahun               string  `gorm:"tahun" json:"tahun"`
	IsiSurat            string  `gorm:"column:isiSurat" json:"isiSurat"`
	IdMaker             string  `gorm:"column:idMaker" json:"idMaker"`
	PnApprover          string  `gorm:"column:pnApprover" json:"pnApprover"`
	StatusApprover      string  `gorm:"column:statusApprover" json:"statusApprover"`
	PosisiApprover      string  `gorm:"column:posisiApprover" json:"posisiApprover"`
	SuratKeluarApprover string  `gorm:"column:suratKeluarApprover" json:"suratKeluarApprover"`
	StatusMCS           string  `gorm:"column:statusMCS" json:"statusMCS"`
	CreatedAt           *string `gorm:"column:createdAt" json:"createdAt"`
	Status              string  `gorm:"column:status" json:"status"`
}

// type PelaporanDraft struct {
// 	IdTemplate         string  `json:"idTemplate"`
// 	BranchCodePenerima string  `json:"branchCodePenerima"`
// 	OrgehPenerima      string  `json:"orgehPenerima"`
// 	BranchCodeTindasan string  `json:"branchCodeTindasan"`
// 	OrgehTindasan      string  `json:"orgehTindasan"`
// 	PnPenerima         string  `json:"pnPenerima"`
// 	PnTindasan         string  `json:"pnTindasan"`
// 	KodeSurat          string  `json:"kodeSurat"`
// 	Kerahasiaan        string  `json:"kerahasiaan"`
// 	Kesegeraan         string  `json:"kesegaraan"`
// 	KepadaYth          string  `json:"kepadaYth"`
// 	Perihal            string  `json:"perihal"`
// 	IsiSurat           string  `json:"isiSurat"`
// 	IdMaker            string  `json:"idMaker"`
// 	PnApprover         string  `json:"pnApprover"`
// 	StatusApprover     string  `json:"statusApprover"`
// 	CreatedAt          *string `json:"createdAt"`
// }

type SuratDraftResponse struct {
	Perihal  string `json:"perihal"`
	IsiSurat string `json:"isi_surat"`
}

type GenerateLaporanRequest struct {
	Branch    string `json:"branch"`
	Mainbr    string `json:"mainbr"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type GenerateLaporanResponse struct {
	ID        string `json:"id"`
	Aktifitas string `json:"aktifitas"`
	Product   string `json:"product"`
	RiskEvent string `json:"risk_event"`
	Prioritas string `json:"prioritas"`
}

type DraftListRequest struct {
	PnMaker        string `json:"pn_maker"`
	Order          string `json:"order"`
	Sort           string `json:"sort"`
	Offset         int    `json:"offset"`
	Limit          int    `json:"limit"`
	Page           int    `json:"page"`
	Filter         bool   `json:"filter"`
	JenisPencarian string `json:"jenis_pencarian"`
	Keyword        string `json:"keyword"`
	Status         string `json:"status"`
}

type DraftListResponse struct {
	ID             string  `json:"id"`
	KodeSurat      string  `json:"kode_surat"`
	NomorSurat     *string `json:"nomor_surat"`
	PnMaker        string  `json:"pn_maker"`
	NamaMaker      string  `json:"nama_maker"`
	Perihal        string  `json:"perihal"`
	Tujuan         string  `json:"tujuan"`
	StatusTerakhir string  `json:"status_terakhir"`
	ResponseStatus string  `json:"response_status"`
	PosisiApprover string  `json:"posisi_approver"`
	TanggalSurat   *string `json:"tanggal_surat"`
}

type PenerimaSuratResponse struct {
	Pernr string `json:"pernr"`
	Sname string `json:"sname"`
	Orgeh string `json:"orgeh"`
}

type KodeResponse struct {
	Kode string `json:"kode"`
}

type SuratRequestOne struct {
	ID int64 `json:"id"`
}

type SuratDetail struct {
	ID             string  `json:"id"`
	NomorSurat     string  `json:"nomor_surat"`
	KepadaYth      string  `json:"kepada_yth"`
	Penerima       string  `json:"penerima"`
	Pengirim       string  `json:"pengirim"`
	Perihal        string  `json:"perihal"`
	TanggalSurat   string  `json:"tanggal_surat"`
	IsiSurat       string  `json:"isi_surat"`
	PnApprover     string  `json:"pn_approver"`
	PosisiApprover string  `json:"posisi_approver"`
	StatusTerakhir string  `json:"status_terakhir"`
	Catatan        string  `json:"catatan"`
	Penolak        string  `json:"penolak"`
	TanggalTolak   *string `json:"tanggal_tolak"`
}

type SuratDetailResponse struct {
	ID             string   `json:"id"`
	NomorSurat     string   `json:"nomor_surat"`
	KepadaYth      string   `json:"kepada_yth"`
	Pengirim       string   `json:"pengirim"`
	Perihal        string   `json:"perihal"`
	TanggalSurat   string   `json:"tanggal_surat"`
	IsiSurat       string   `json:"isi_surat"`
	Signer         []Signer `json:"signer"`
	PosisiApprover string   `json:"posisi_approver"`
	StatusTerakhir string   `json:"status_terakhir"`
	Catatan        string   `json:"catatan"`
	Penolak        string   `json:"penolak"`
	TanggalTolak   *string  `json:"tanggal_tolak"`
}

func (ApproveUpdate) TableName() string {
	return "pelaporan_drafts"
}

func (PenolakanCatatan) TableName() string {
	return "pelaporan_rejected"
}
