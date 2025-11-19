package models

type MajorProsesRequest struct {
	ID              int64   `json:"id"`
	IDMegaProses    string  `json:"id_mega_proses"`
	KodeMajorProses string  `json:"kode_major_proses"`
	MajorProses     string  `json:"major_proses"`
	Deskripsi       string  `json:"deskripsi"`
	Status          bool    `json:"status"`
	CreatedAt       *string `json:"created_at"`
	UpdatedAt       *string `json:"updated_at"`
}

type MajorProsesResponse struct {
	ID              int64   `json:"id"`
	IDMegaProses    string  `json:"id_mega_proses"`
	MegaProsesName  string  `json:"mega_proses_name"`
	KodeMajorProses string  `json:"kode_major_proses"`
	MajorProses     string  `json:"major_proses"`
	Deskripsi       string  `json:"deskripsi"`
	Status          bool    `json:"status"`
	CreatedAt       *string `json:"created_at"`
	UpdatedAt       *string `json:"updated_at"`
}

type Paginate struct {
	Order  string `json:"order"`
	Sort   string `json:"sort"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Page   int    `json:"page"`
}

type KodeMegaProses struct {
	KodeMegaProses string `json:"kode_mega_proses"`
}

type KodeMajorProses struct {
	KodeMegaProses  string `json:"kode_mega_proses"`
	KodeMajorProses string `json:"kode_major_proses"`
}

func (p MajorProsesRequest) ParseRequest() MajorProses {
	return MajorProses{
		ID:              p.ID,
		IDMegaProses:    p.IDMegaProses,
		KodeMajorProses: p.KodeMajorProses,
		MajorProses:     p.MajorProses,
		Deskripsi:       p.Deskripsi,
		Status:          p.Status,
	}
}

func (p MajorProsesResponse) ParseResponse() MajorProses {
	return MajorProses{
		ID:              p.ID,
		IDMegaProses:    p.IDMegaProses,
		MegaProsesName:  p.MegaProsesName,
		KodeMajorProses: p.KodeMajorProses,
		MajorProses:     p.MajorProses,
		Deskripsi:       p.Deskripsi,
		Status:          p.Status,
		CreatedAt:       p.CreatedAt,
		UpdatedAt:       p.UpdatedAt,
	}
}

func (ET1 MajorProsesRequest) TableName() string {
	return "major_proses"
}

func (ET1 MajorProsesResponse) TableName() string {
	return "major_proses"
}
