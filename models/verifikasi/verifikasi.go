package models

type Verifikasi struct {
	ID                        int64
	NoPelaporan               string
	REGION                    string
	RGDESC                    string
	MAINBR                    string
	MBDESC                    string
	BRANCH                    string
	BRDESC                    string
	UnitKerja                 string
	ActivityID                int64
	SubActivityID             int64
	ProductID                 int64
	RiskIssueID               int64
	RiskIssue                 string
	RiskIndicatorID           int64
	RiskIndicator             string
	SumberData                string
	ApplicationID             string
	HasilVerifikasi           string
	KunjunganNasabah          bool
	Perbaikan                 bool
	IndikasiFraud             bool
	TerdapatKerugianFinansial bool
	JenisKerugianFinansial    string
	JumlahPerkiraanKerugian   int64
	JenisKerugianNonFinansial string
	JenisRekomendasi          string
	RekomendasiTindakLanjut   string
	RencanaTindakLanjut       string
	RiskTypeID                int64
	AdaUsulanPerbaikan        bool
	TanggalDitemukan          *string
	TanggalMulaiRTL           *string
	TanggalTargetSelesai      *string
	MakerID                   string
	MakerDesc                 string
	MakerDate                 *string
	LastMakerID               string
	LastMakerDesc             string
	LastMakerDate             *string
	Status                    string
	Action                    string // create, updateApproval, updateMaintain, delete, publish, unpublish
	StatusIndikasiFraud       string
	ActionIndikasiFraud       string
	Deleted                   bool
	UpdatedAt                 *string
	CreatedAt                 *string
}

type VerifikasiUpdateDelete struct {
	ID            int64
	LastMakerID   string
	LastMakerDesc string
	LastMakerDate *string
	Status        string
	Action        string // create, updateApproval, updateMaintain, delete, publish, unpublish
	Deleted       bool
	UpdatedAt     *string
}

type VerifikasiUpdateMaintain struct {
	ID            int64
	LastMakerID   string
	LastMakerDesc string
	LastMakerDate *string
	Status        string
	Action        string // create, updateApproval, updateMaintain, delete, publish, unpublish
	UpdatedAt     *string
}

type VerifikasiUpdateAll struct {
	ID                        int64
	NoPelaporan               string
	REGION                    string
	RGDESC                    string
	MAINBR                    string
	MBDESC                    string
	BRANCH                    string
	BRDESC                    string
	UnitKerja                 string
	ActivityID                int64
	SubActivityID             int64
	ProductID                 int64
	RiskIssueID               int64
	RiskIssue                 string
	RiskIndicatorID           int64
	RiskIndicator             string
	SumberData                string
	ApplicationID             string
	HasilVerifikasi           string
	KunjunganNasabah          bool
	Perbaikan                 bool
	IndikasiFraud             bool
	TerdapatKerugianFinansial bool
	JenisKerugianFinansial    string
	JumlahPerkiraanKerugian   int64
	JenisKerugianNonFinansial string
	JenisRekomendasi          string
	RekomendasiTindakLanjut   string
	RencanaTindakLanjut       string
	RiskTypeID                int64
	AdaUsulanPerbaikan        bool
	TanggalDitemukan          *string
	TanggalMulaiRTL           *string
	TanggalTargetSelesai      *string
	LastMakerID               string
	LastMakerDesc             string
	LastMakerDate             *string
	Status                    string
	Action                    string // create, updateApproval, updateMaintain, delete, publish, unpublish
	StatusIndikasiFraud       string
	ActionIndikasiFraud       string
	Deleted                   bool
	UpdatedAt                 *string
}

type VerifyDataRequest struct {
	ID           int64 `json:"id"`
	VerifikasiID int64 `json:"verifikasi_id"`
}

func (v Verifikasi) TableName() string {
	return "verifikasi"
}

func (v VerifikasiUpdateDelete) TableName() string {
	return "verifikasi"
}

func (v VerifikasiUpdateMaintain) TableName() string {
	return "verifikasi"
}

func (v VerifikasiUpdateAll) TableName() string {
	return "verifikasi"
}
