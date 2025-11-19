package models

type LiniBisnisLv3Request struct {
	ID                int64   `json:"id"`
	IDLiniBisnisLv2   string  `json:"id_lini_bisnis_lv2"`
	KodeLiniBisnisLv3 string  `json:"kode_lini_bisnis_lv3"`
	LiniBisnisLv3     string  `json:"lini_bisnis_lv3"`
	Deskripsi         string  `json:"deskripsi"`
	Status            bool    `json:"status"`
	CreatedAt         *string `json:"created_at"`
	UpdatedAt         *string `json:"updated_at"`
}

type Paginate struct {
	Order  string `json:"order"`
	Sort   string `json:"sort"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Page   int    `json:"page"`
}

type LiniBisnisLv3Response struct {
	ID                int64   `json:"id"`
	IDLiniBisnisLv2   string  `json:"id_lini_bisnis_lv2"`
	KodeLiniBisnisLv3 string  `json:"kode_lini_bisnis_lv3"`
	LiniBisnisLv3     string  `json:"lini_bisnis_lv3"`
	Deskripsi         string  `json:"deskripsi"`
	Status            bool    `json:"status"`
	CreatedAt         *string `json:"created_at"`
	UpdatedAt         *string `json:"updated_at"`
}

type KodeLiniBisnis struct {
	KodeLiniBisnis string `json:"kode_lini_bisnis"`
}

type KodeLB2 struct {
	KodeLB string `json:"kode_lb"`
}

func (p LiniBisnisLv3Request) ParseRequest() LiniBisnisLv3 {
	return LiniBisnisLv3{
		ID:                p.ID,
		IDLiniBisnisLv2:   p.IDLiniBisnisLv2,
		KodeLiniBisnisLv3: p.KodeLiniBisnisLv3,
		LiniBisnisLv3:     p.LiniBisnisLv3,
		Deskripsi:         p.Deskripsi,
		Status:            p.Status,
	}
}

func (p LiniBisnisLv3Response) ParseResponse() LiniBisnisLv3 {
	return LiniBisnisLv3{
		ID:                p.ID,
		IDLiniBisnisLv2:   p.IDLiniBisnisLv2,
		KodeLiniBisnisLv3: p.KodeLiniBisnisLv3,
		LiniBisnisLv3:     p.LiniBisnisLv3,
		Deskripsi:         p.Deskripsi,
		Status:            p.Status,
		CreatedAt:         p.CreatedAt,
		UpdatedAt:         p.UpdatedAt,
	}
}

func (LB3 LiniBisnisLv3Request) TableName() string {
	return "lini_bisnis_lv3"
}

func (LB3 LiniBisnisLv3Response) TableName() string {
	return "lini_bisnis_lv3"
}
