package models

import (
	"riskmanagement/lib"
	files "riskmanagement/models/files"
)

type VerifikasiRealisasiRequest struct {
	ID               int64                        `json:"id"`
	NoPelaporan      string                       `json:"no_pelaporan"`
	SumberData       string                       `json:"sumber_data"`
	REGION           string                       `json:"REGION"`
	RGDESC           string                       `json:"RGDESC"`
	MAINBR           string                       `json:"MAINBR"`
	MBDESC           string                       `json:"MBDESC"`
	BRANCH           string                       `json:"BRANCH"`
	BRDESC           string                       `json:"BRDESC"`
	ActivityID       int64                        `json:"activity_id"`
	ActivityName     string                       `json:"activity_name"`
	ProductID        int64                        `json:"product_id"`
	ProductName      string                       `json:"product_name"`
	SubActivityID    int64                        `json:"sub_activity_id"`
	SubActivityName  string                       `json:"sub_activity_name"`
	RestruckFlag     bool                         `json:"restruck_flag"`
	PeriodeData      string                       `json:"periode_data"`
	KunjunganNasabah bool                         `json:"kunjungan_nasabah"`
	TglKunjungan     string                       `json:"tgl_kunjungan"`
	ButuhPerbaikan   bool                         `json:"butuh_perbaikan"`
	IndikasiFraud    bool                         `json:"indikasi_fraud"`
	HasilVerifikasi  string                       `json:"hasil_verifikasi"`
	Status           string                       `json:"status"`
	Action           string                       `json:"action"`
	Deleted          bool                         `json:"deleted"`
	StatusValidasi   string                       `json:"status_validasi"`
	ActionValidasi   string                       `json:"action_validasi"`
	CreatedAt        string                       `json:"created_at"`
	CreatedID        string                       `json:"created_id"`
	CreatedDesc      string                       `json:"created_desc"`
	UpdatedAt        string                       `json:"updated_at"`
	UpdatedBy        string                       `json:"updated_by"`
	UpdatedDesc      string                       `json:"updated_desc"`
	SampleData       []SampleDataRealisasiRequest `json:"sample_data"`
	KriteriaData     string                       `json:"kriteria_data"`
	ListCriteria     string                       `json:"list_criteria"`
	Files            []files.FilesRequest         `json:"files"`
}

type VerifikasiRealisasiResponse struct {
	ID               int64                              `json:"id"`
	NoPelaporan      string                             `json:"no_pelaporan"`
	SumberData       string                             `json:"sumber_data"`
	REGION           string                             `json:"REGION"`
	RGDESC           string                             `json:"RGDESC"`
	MAINBR           string                             `json:"MAINBR"`
	MBDESC           string                             `json:"MBDESC"`
	BRANCH           string                             `json:"BRANCH"`
	BRDESC           string                             `json:"BRDESC"`
	ActivityID       int64                              `json:"activity_id"`
	ActivityName     string                             `json:"activity_name"`
	ProductID        int64                              `json:"product_id"`
	ProductName      string                             `json:"product_name"`
	SubActivityID    int64                              `json:"sub_activity_id"`
	SubActivityName  string                             `json:"sub_activity_name"`
	RestruckFlag     bool                               `json:"restruck_flag"`
	PeriodeData      string                             `json:"periode_data"`
	KunjunganNasabah bool                               `json:"kunjungan_nasabah"`
	TglKunjungan     string                             `json:"tgl_kunjungan"`
	ButuhPerbaikan   bool                               `json:"butuh_perbaikan"`
	IndikasiFraud    bool                               `json:"indikasi_fraud"`
	HasilVerifikasi  string                             `json:"hasil_verifikasi"`
	Status           string                             `json:"status"`
	Action           string                             `json:"action"`
	Deleted          bool                               `json:"deleted"`
	StatusValidasi   string                             `json:"status_validasi"`
	ActionValidasi   string                             `json:"action_validasi"`
	CreatedAt        string                             `json:"created_at"`
	CreatedID        string                             `json:"created_id"`
	CreatedDesc      string                             `json:"created_desc"`
	UpdatedAt        string                             `json:"updated_at"`
	UpdatedBy        string                             `json:"updated_by"`
	UpdatedDesc      string                             `json:"updated_desc"`
	SampleData       []SampleDataRealisasiResponse      `json:"sample_data"`
	KriteriaData     string                             `json:"kriteria_data"`
	ListCriteria     string                             `json:"list_criteria"`
	Files            []VerifikasiRealisasiFilesResponse `json:"files"`
}

type VerifikasiRealisasiDetailResponse struct {
	ID               int64  `json:"id"`
	NoPelaporan      string `json:"no_pelaporan"`
	SumberData       string `json:"sumber_data"`
	REGION           string `json:"REGION"`
	RGDESC           string `json:"RGDESC"`
	MAINBR           string `json:"MAINBR"`
	MBDESC           string `json:"MBDESC"`
	BRANCH           string `json:"BRANCH"`
	BRDESC           string `json:"BRDESC"`
	ActivityID       int64  `json:"activity_id"`
	ActivityName     string `json:"activity_name"`
	ProductID        int64  `json:"product_id"`
	ProductName      string `json:"product_name"`
	SubActivityID    int64  `json:"sub_activity_id"`
	SubActivityName  string `json:"sub_activity_name"`
	RestruckFlag     bool   `json:"restruck_flag"`
	PeriodeData      string `json:"periode_data"`
	KunjunganNasabah bool   `json:"kunjungan_nasabah"`
	TglKunjungan     string `json:"tgl_kunjungan"`
	ButuhPerbaikan   bool   `json:"butuh_perbaikan"`
	IndikasiFraud    bool   `json:"indikasi_fraud"`
	HasilVerifikasi  string `json:"hasil_verifikasi"`
	Status           string `json:"status"`
	Action           string `json:"action"`
	Deleted          bool   `json:"deleted"`
	StatusValidasi   string `json:"status_validasi"`
	ActionValidasi   string `json:"action_validasi"`
	CreatedAt        string `json:"created_at"`
	CreatedID        string `json:"created_id"`
	CreatedDesc      string `json:"created_desc"`
	UpdatedAt        string `json:"updated_at"`
	UpdatedBy        string `json:"updated_by"`
	UpdatedDesc      string `json:"updated_desc"`
	KriteriaData     string `json:"kriteria_data"`
	ListCriteria     string `json:"list_criteria"`
}

func (v VerifikasiRealisasiRequest) TableName() string {
	return "verifikasi_realisasi_kredit"
}

func (v VerifikasiRealisasiResponse) TableName() string {
	return "verifikasi_realisasi_kredit"
}

type VerifikasiRealisasiList struct {
	ID            int64  `json:"id"`
	No            int64  `json:"no"`
	NoPelaporan   string `json:"no_pelaporan"`
	UnitKerja     string `json:"unit_kerja"`
	Aktifitas     string `json:"aktifitas"`
	IndikasiFraud string `json:"indikasi_fraud"`
	StatusVerif   string `json:"status_verif"`
	StatusFraud   string `json:"status_fraud"`
}

type VerifikasiRealisasiPagination struct {
	Order    string `json:"order"`
	Sort     string `json:"sort"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
	Page     int    `json:"page"`
	Pernr    string `json:"pernr"`
	Branches string `json:"branches"`
	Kostl    string `json:"kostl"`
}

type VerifikasiRealisasiFilterRequest struct {
	Order           string `json:"order"`
	Sort            string `json:"sort"`
	Offset          int    `json:"offset"`
	Limit           int    `json:"limit"`
	Page            int    `json:"page"`
	Pernr           string `json:"pernr"`
	NoPelaporan     string `json:"no_pelaporan"`
	UnitKerja       string `json:"unit_kerja"`
	ActivityID      string `json:"activity_id"`
	IndikasiFraud   string `json:"indikasi_fraud"`
	Status          string `json:"status"`
	Branches        string `json:"branches"`
	Kostl           string `json:"kostl"`
	ProductID       string `json:"product_id"`
	Segment         string `json:"segment"`
	CriteriaID      string `json:"criteria_id"`
	SudahVerifikasi string `json:"sudah_verifikasi"`
	Efektif         string `json:"efektif"`
	REGION          string `json: REGION`
	MAINBR          string `json: MAINBR`
	BRANCH          string `json: BRANCH`
}

type VerifikasiRequestUpdateMaintain struct {
	LastMakerID   string `json:"last_maker_id"`
	LastMakerDesc string `json:"last_maker_desc"`
}

type VerifikasiRealisasiUpdateDelete struct {
	ID          int64   `json:"id"`
	UpdatedBy   string  `json:"updated_by"`
	UpdatedDesc string  `json:"updated_desc"`
	Deleted     bool    `json:"deleted"`
	UpdatedAt   *string `json:"updated_at"`
}

type NoPalaporanVerifikasiRealisasiRequest struct {
	ORGEH string `json:"ORGEH"`
}

type NoPelaporanVerifikasiRealisasiResponse struct {
	ORGEH       string `json:"ORGEH"`
	NoPelaporan string `json:"no_pelaporan"`
}

type NoPelaporanVerifikasiRealisasiNullResponse struct {
	NoPelaporan lib.NullString `json:"no_pelaporan"`
}

type VerifikasiRealisasiRequestID struct {
	ID int64 `json:"id"`
}
