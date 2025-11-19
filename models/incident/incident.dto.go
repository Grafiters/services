package models

type IncidentRequest struct {
	ID               int64   `json:"id"`
	KodeKejadian     string  `json:"kode_kejadian"`
	PenyebabKejadian string  `json:"penyebab_kejadian"`
	Deskripsi        string  `json:"deskripsi"`
	Status           bool    `json:"status"`
	CreatedAt        *string `json:"created_at"`
	UpdatedAt        *string `json:"updated_at"`
}

type Paginate struct {
	Order  string `json:"order"`
	Sort   string `json:"sort"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Page   int    `json:"page"`
}

type IncidentResponse struct {
	ID               int64   `json:"id"`
	KodeKejadian     string  `json:"kode_kejadian"`
	PenyebabKejadian string  `json:"penyebab_kejadian"`
	Deskripsi        string  `json:"deskripsi"`
	Status           bool    `json:"status"`
	CreatedAt        *string `json:"created_at"`
	UpdatedAt        *string `json:"updated_at"`
}

type IncidentResponses struct {
	ID               int64   `json:"id"`
	KodeKejadian     string  `json:"kode_kejadian"`
	PenyebabKejadian string  `json:"penyebab_kejadian"`
	Deskripsi        string  `json:"deskripsi"`
	Status           bool    `json:"status"`
	CreatedAt        *string `json:"created_at"`
	UpdatedAt        *string `json:"updated_at"`
}

type KodePenyebabKejadian struct {
	KodePenyebabKejadian string `json:"kode_penyebab_kejadian"`
}

func (p IncidentRequest) ParseRequest() Incident {
	return Incident{
		ID:               p.ID,
		KodeKejadian:     p.KodeKejadian,
		PenyebabKejadian: p.PenyebabKejadian,
		Deskripsi:        p.Deskripsi,
		Status:           p.Status,
	}
}

func (p IncidentResponse) ParseResponse() Incident {
	return Incident{
		ID:               p.ID,
		KodeKejadian:     p.KodeKejadian,
		PenyebabKejadian: p.PenyebabKejadian,
		Deskripsi:        p.Deskripsi,
		Status:           p.Status,
		CreatedAt:        p.CreatedAt,
		UpdatedAt:        p.UpdatedAt,
	}
}

func (pr IncidentRequest) TableName() string {
	return "penyebab_kejadian_lv1"
}

func (pr IncidentResponse) TableName() string {
	return "penyebab_kejadian_lv1"
}
