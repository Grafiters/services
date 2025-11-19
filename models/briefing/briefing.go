package models

type Briefing struct {
	ID             int64
	NoPelaporan    string
	UnitKerja      string
	REGION         string
	RGDESC         string
	RGNAME         string
	MAINBR         string
	MBDESC         string
	MBNAME         string
	BRANCH         string
	BRDESC         string
	BRNAME         string
	JenisPeserta   string
	JabatanPeserta string
	JumlahPeserta  int64
	ListPeserta    string
	MakerID        string
	MakerDesc      string
	MakerDate      *string
	LastMakerID    string
	LastMakerDesc  string
	LastMakerDate  *string
	Status         string
	Action         string // create, updateApproval, updateMaintain, delete, publish, unpublish
	Deleted        bool
	UpdatedAt      *string
	CreatedAt      *string
}

type BriefingUpdateDelete struct {
	ID            int64
	LastMakerID   string
	LastMakerDesc string
	LastMakerDate *string
	Deleted       bool
	Action        string
	Status        string
	UpdatedAt     *string
}

type BriefingUpdateMateri struct {
	ID             int64
	REGION         string
	RGDESC         string
	MAINBR         string
	MBDESC         string
	BRANCH         string
	BRDESC         string
	UnitKerja      string
	JenisPeserta   string
	JabatanPeserta string
	JumlahPeserta  int64
	ListPeserta    string
	LastMakerID    string
	LastMakerDesc  string
	LastMakerDate  *string
	Deleted        bool
	Action         string
	Status         string
	UpdatedAt      *string
}

type BriefMateriRequest struct {
	ID         int64 `json:"id"`
	BriefingID int64 `json:"briefing_id"`
}

func (b Briefing) TableName() string {
	return "briefing"
}

func (b BriefingUpdateDelete) TableName() string {
	return "briefing"
}

func (b BriefingUpdateMateri) TableName() string {
	return "briefing"
}
