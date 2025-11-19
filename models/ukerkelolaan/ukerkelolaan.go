package models

type UkerKelolaan struct {
	ID        int64
	CreatedAt *string
	UpdatedAt *string
	ExpiredAt string
	IsTemp    bool
	Pn        string
	SNAME     string
	REGION    string
	RGDESC    string
	MAINBR    string
	MBDESC    string
	BRANCH    string
	BRDESC    string
	Status    bool
}

type UkerKelolaanReqDelete struct {
	ID        int64
	UpdatedAt *string
	Status    bool
}

func (uk UkerKelolaan) TableName() string {
	return "uker_kelolaan_user"
}

func (uk UkerKelolaanReqDelete) TableName() string {
	return "uker_kelolaan_user"
}
