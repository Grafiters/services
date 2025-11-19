package models

type SubMajorProsesRequest struct {
	ID                 int64   `json:"id"`
	IDMajorProses      string  `json:"id_major_proses"`
	KodeSubMajorProses string  `json:"kode_sub_major_proses"`
	SubMajorProses     string  `json:"sub_major_proses"`
	Deskripsi          string  `json:"deskripsi"`
	Status             bool    `json:"status"`
	CreatedAt          *string `json:"created_at"`
	UpdatedAt          *string `json:"updated_at"`
}

type Paginate struct {
	Order  string `json:"order"`
	Sort   string `json:"sort"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Page   int    `json:"page"`
}

type SubMajorProsesResponse struct {
	ID                 int64   `json:"id"`
	IDMajorProses      string  `json:"id_major_proses"`
	MajorProses        string  `json:"major_proses"`
	KodeSubMajorProses string  `json:"kode_sub_major_proses"`
	SubMajorProses     string  `json:"sub_major_proses"`
	Deskripsi          string  `json:"deskripsi"`
	Status             bool    `json:"status"`
	CreatedAt          *string `json:"created_at"`
	UpdatedAt          *string `json:"updated_at"`
}

type KodeMajor struct {
	KodeMajor string `json:"kode_major"`
}

func (p SubMajorProsesRequest) ParseRequest() SubMajorProses {
	return SubMajorProses{
		ID:                 p.ID,
		IDMajorProses:      p.IDMajorProses,
		KodeSubMajorProses: p.KodeSubMajorProses,
		SubMajorProses:     p.SubMajorProses,
		Deskripsi:          p.Deskripsi,
		Status:             p.Status,
	}
}

func (p SubMajorProsesResponse) ParseResponse() SubMajorProses {
	return SubMajorProses{
		ID:                 p.ID,
		IDMajorProses:      p.IDMajorProses,
		KodeSubMajorProses: p.KodeSubMajorProses,
		SubMajorProses:     p.SubMajorProses,
		Deskripsi:          p.Deskripsi,
		Status:             p.Status,
		CreatedAt:          p.CreatedAt,
		UpdatedAt:          p.UpdatedAt,
	}
}

func (ET1 SubMajorProsesRequest) TableName() string {
	return "sub_major_proses"
}

func (ET1 SubMajorProsesResponse) TableName() string {
	return "sub_major_proses"
}
