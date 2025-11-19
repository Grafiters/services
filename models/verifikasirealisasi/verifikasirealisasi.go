package models

// import (
// 	"riskmanagement/lib"
// 	files "riskmanagement/models/files"
// )

type VerifikasiRealisasi struct {
	ID               int64
	NoPelaporan      string
	SumberData       string
	REGION           string
	RGDESC           string
	MAINBR           string
	MBDESC           string
	BRANCH           string
	BRDESC           string
	ActivityID       int64
	ActivityName     string
	ProductID        int64
	ProductName      string
	SubActivityID    int64
	SubActivityName  string
	RestruckFlag     bool
	PeriodeData      string
	KunjunganNasabah bool
	TglKunjungan     string
	ButuhPerbaikan   bool
	IndikasiFraud    bool
	HasilVerifikasi  string
	Status           string
	Action           string
	Deleted          bool
	StatusValidasi   string
	ActionValidasi   string
	CreatedAt        string
	CreatedID        string
	CreatedDesc      string
	UpdatedAt        string
	UpdatedBy        string
	UpdatedDesc      string
	KriteriaData     string
	ListCriteria     string
}

type VerifikasiRealisasiUpdate struct {
	ID               int64
	NoPelaporan      string
	SumberData       string
	REGION           string
	RGDESC           string
	MAINBR           string
	MBDESC           string
	BRANCH           string
	BRDESC           string
	ActivityID       int64
	ActivityName     string
	ProductID        int64
	ProductName      string
	SubActivityID    int64
	SubActivityName  string
	RestruckFlag     bool
	PeriodeData      string
	KunjunganNasabah bool
	TglKunjungan     string
	ButuhPerbaikan   bool
	IndikasiFraud    bool
	HasilVerifikasi  string
	Status           string
	Action           string
	Deleted          bool
	StatusValidasi   string
	ActionValidasi   string
	UpdatedAt        string
	UpdatedBy        string
	UpdatedDesc      string
	KriteriaData     string
	ListCriteria     string
}

func (v VerifikasiRealisasi) TableName() string {
	return "verifikasi_realisasi_kredit"
}

func (v VerifikasiRealisasiUpdate) Table() string {
	return "verifikasi_realisasi_kredit"
}
