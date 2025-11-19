package models

type MegaProsesRequest struct {
	ID             int64   `json:"id"`
	KodeMegaProses string  `json:"kode_mega_proses"`
	MegaProses     string  `json:"mega_proses"`
	Deskripsi      string  `json:"deskripsi"`
	Status         bool    `json:"status"`
	CreatedAt      *string `json:"created_at"`
	UpdatedAt      *string `json:"updated_at"`
}

type Paginate struct {
	Order  string `json:"order"`
	Sort   string `json:"sort"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Page   int    `json:"page"`
}

type MegaProsesResponse struct {
	ID             int64   `json:"id"`
	KodeMegaProses string  `json:"kode_mega_proses"`
	MegaProses     string  `json:"mega_proses"`
	Deskripsi      string  `json:"deskripsi"`
	Status         bool    `json:"status"`
	CreatedAt      *string `json:"created_at"`
	UpdatedAt      *string `json:"updated_at"`
}

type KodeMegaProses struct {
	KodeMegaProses string `json:"kode_mega_proses"`
}

func (p MegaProsesRequest) ParseRequest() MegaProses {
	return MegaProses{
		ID:             p.ID,
		KodeMegaProses: p.KodeMegaProses,
		MegaProses:     p.MegaProses,
		Deskripsi:      p.Deskripsi,
		Status:         p.Status,
	}
}

func (p MegaProsesResponse) ParseResponse() MegaProses {
	return MegaProses{
		ID:             p.ID,
		KodeMegaProses: p.KodeMegaProses,
		MegaProses:     p.MegaProses,
		Deskripsi:      p.Deskripsi,
		Status:         p.Status,
		CreatedAt:      p.CreatedAt,
		UpdatedAt:      p.UpdatedAt,
	}
}

func (ET1 MegaProsesRequest) TableName() string {
	return "mega_proses"
}

func (ET1 MegaProsesResponse) TableName() string {
	return "mega_proses"
}
