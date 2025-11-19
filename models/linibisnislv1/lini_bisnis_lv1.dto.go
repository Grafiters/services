package models

type LiniBisnisLv1Request struct {
	ID             int64   `json:"id"`
	KodeLiniBisnis string  `json:"kode_lini_bisnis"`
	LiniBisnis1    string  `json:"lini_bisnis1"`
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

type LiniBisnisLv1Response struct {
	ID             int64   `json:"id"`
	KodeLiniBisnis string  `json:"kode_lini_bisnis"`
	LiniBisnis1    string  `json:"lini_bisnis1"`
	Deskripsi      string  `json:"deskripsi"`
	Status         bool    `json:"status"`
	CreatedAt      *string `json:"created_at"`
	UpdatedAt      *string `json:"updated_at"`
}

type KodeLiniBisnis struct {
	KodeLiniBisnis string `json:"kode_lini_bisnis"`
}

func (p LiniBisnisLv1Request) ParseRequest() LiniBisnisLv1 {
	return LiniBisnisLv1{
		ID:             p.ID,
		KodeLiniBisnis: p.KodeLiniBisnis,
		LiniBisnisLv1:  p.LiniBisnis1,
		Deskripsi:      p.Deskripsi,
		Status:         p.Status,
	}
}

func (p LiniBisnisLv1Response) ParseResponse() LiniBisnisLv1 {
	return LiniBisnisLv1{
		ID:             p.ID,
		KodeLiniBisnis: p.KodeLiniBisnis,
		LiniBisnisLv1:  p.LiniBisnis1,
		Deskripsi:      p.Deskripsi,
		Status:         p.Status,
		CreatedAt:      p.CreatedAt,
		UpdatedAt:      p.UpdatedAt,
	}
}

func (LB1 LiniBisnisLv1Request) TableName() string {
	return "lini_bisnis_lv1"
}

func (LB1 LiniBisnisLv1Response) TableName() string {
	return "lini_bisnis_lv1"
}
