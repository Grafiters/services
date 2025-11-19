package models

type PenyebabKejadianLv3Request struct {
	ID                      int64   `json:"id"`
	KodeSubKejadian         string  `json:"kode_sub_kejadian"`
	KodePenyebabKejadianLv3 string  `json:"kode_penyebab_kejadian_lv3"`
	PenyebabKejadianLv3     string  `json:"penyebab_kejadian_lv3"`
	Deskripsi               string  `json:"deskripsi"`
	Status                  bool    `json:"status"`
	CreatedAt               *string `json:"created_at"`
	UpdatedAt               *string `json:"updated_at"`
}

type PenyebabKejadianLv3Response struct {
	ID                      int64   `json:"id"`
	KodeSubKejadian         string  `json:"kode_sub_kejadian"`
	KodePenyebabKejadianLv3 string  `json:"kode_penyebab_kejadian_lv3"`
	PenyebabKejadianLv3     string  `json:"penyebab_kejadian_lv3"`
	Deskripsi               string  `json:"deskripsi"`
	Status                  bool    `json:"status"`
	CreatedAt               *string `json:"created_at"`
	UpdatedAt               *string `json:"updated_at"`
}

type KodeKejadian struct {
	KodeKejadian string `json:"kode_kejadian"`
}

type Paginate struct {
	Order  string `json:"order"`
	Sort   string `json:"sort"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Page   int    `json:"page"`
}

type PenyebabKejadianLv3Responses struct {
	ID                      int64   `json:"id"`
	KodeSubKejadian         string  `json:"kode_sub_kejadian"`
	PenyebabKejadianLv2     string  `json:"penyebab_kejadian_lv2"`
	KodePenyebabKejadianLv3 string  `json:"kode_penyebab_kejadian_lv3"`
	PenyebabKejadianLv3     string  `json:"penyebab_kejadian_lv3"`
	Deskripsi               string  `json:"deskripsi"`
	Status                  bool    `json:"status"`
	CreatedAt               *string `json:"created_at"`
	UpdatedAt               *string `json:"updated_at"`
}

type KodePenyebabKejadian struct {
	KodePenyebabKejadian string `json:"kode_penyebab_kejadian"`
}

type KodeSubKejadianRequest struct {
	KodeSubKejadian string `json:"kode_sub_kejadian"`
}

func (p PenyebabKejadianLv3Request) ParseRequest() PenyebabKejadianLv3 {
	return PenyebabKejadianLv3{
		ID:                      p.ID,
		KodeSubKejadian:         p.KodeSubKejadian,
		KodePenyebabKejadianLv3: p.KodePenyebabKejadianLv3,
		PenyebabKejadianLv3:     p.PenyebabKejadianLv3,
		Deskripsi:               p.Deskripsi,
		Status:                  p.Status,
	}
}

func (p PenyebabKejadianLv3Response) ParseResponse() PenyebabKejadianLv3 {
	return PenyebabKejadianLv3{
		ID:                      p.ID,
		KodeSubKejadian:         p.KodeSubKejadian,
		KodePenyebabKejadianLv3: p.KodePenyebabKejadianLv3,
		PenyebabKejadianLv3:     p.PenyebabKejadianLv3,
		Deskripsi:               p.Deskripsi,
		Status:                  p.Status,
		CreatedAt:               p.CreatedAt,
		UpdatedAt:               p.UpdatedAt,
	}
}

func (PK3 PenyebabKejadianLv3Request) TableName() string {
	return "penyebab_kejadian_lv3"
}

func (PK3 PenyebabKejadianLv3Response) TableName() string {
	return "penyebab_kejadian_lv3"
}
