package models

type LiniBisnisLv2Request struct {
	ID                int64   `json:"id"`
	IDLiniBisnisLv1   string  `json:"id_lini_bisnis_lv1"`
	KodeLiniBisnisLv2 string  `json:"kode_lini_bisnis_lv2"`
	LiniBisnisLv2     string  `json:"lini_bisnis_lv2"`
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

type LiniBisnisLv2Response struct {
	ID                int64   `json:"id"`
	IDLiniBisnisLv1   string  `json:"id_lini_bisnis_lv1"`
	KodeLiniBisnisLv2 string  `json:"kode_lini_bisnis_lv2"`
	LiniBisnisLv2     string  `json:"lini_bisnis_lv2"`
	Deskripsi         string  `json:"deskripsi"`
	Status            bool    `json:"status"`
	CreatedAt         *string `json:"created_at"`
	UpdatedAt         *string `json:"updated_at"`
}

type KodeLiniBisnis struct {
	KodeLiniBisnis string `json:"kode_lini_bisnis"`
}

type KodeLB1 struct {
	KodeLB string `json:"kode_lb"`
}

func (p LiniBisnisLv2Request) ParseRequest() LiniBisnisLv2 {
	return LiniBisnisLv2{
		ID:                p.ID,
		IDLiniBisnisLv1:   p.IDLiniBisnisLv1,
		KodeLiniBisnisLv2: p.KodeLiniBisnisLv2,
		LiniBisnisLv2:     p.LiniBisnisLv2,
		Deskripsi:         p.Deskripsi,
		Status:            p.Status,
	}
}

func (p LiniBisnisLv2Response) ParseResponse() LiniBisnisLv2 {
	return LiniBisnisLv2{
		ID:                p.ID,
		IDLiniBisnisLv1:   p.IDLiniBisnisLv1,
		KodeLiniBisnisLv2: p.KodeLiniBisnisLv2,
		LiniBisnisLv2:     p.LiniBisnisLv2,
		Deskripsi:         p.Deskripsi,
		Status:            p.Status,
		CreatedAt:         p.CreatedAt,
		UpdatedAt:         p.UpdatedAt,
	}
}

func (LB2 LiniBisnisLv2Request) TableName() string {
	return "lini_bisnis_lv2"
}

func (LB2 LiniBisnisLv2Response) TableName() string {
	return "lini_bisnis_lv2"
}
