package models

type PgsUser struct {
	ID            int64
	PN            string
	NamaPekerja   string
	UnitKerja     string
	REGION        string
	RGDESC        string
	RGNAME        string
	MAINBR        string
	MBDESC        string
	MBNAME        string
	BRANCH        string
	BRDESC        string
	BRNAME        string
	JabatanPgs    string
	PeriodeAwal   string
	PeriodeAkhir  string
	MakerID       string
	MakerDesc     string
	MakerDate     *string
	LastMakerID   string
	LastMakerDesc string
	LastMakerDate *string
	Status        string
	Action        string
	CreatedAt     *string
	UpdatedAt     *string
}

func (p PgsUser) TableName() string {
	return "pgs_user"
}
