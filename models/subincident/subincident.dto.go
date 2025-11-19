package models

import "riskmanagement/lib"

type SubIncidentRequest struct {
	ID                       int64   `json:"id"`
	KodeKejadian             string  `json:"kode_kejadian"`
	KodeSubKejadian          string  `json:"kode_sub_kejadian"`
	KriteriaPenyebabKejadian string  `json:"kriteria_penyebab_kejadian"`
	Deskripsi                string  `json:"deskripsi"`
	Status                   bool    `json:"status"`
	CreatedAt                *string `json:"created_at"`
	UpdatedAt                *string `json:"updated_at"`
}

type SubIncidentResponse struct {
	ID                       int64   `json:"id"`
	KodeKejadian             string  `json:"kode_kejadian"`
	KodeSubKejadian          string  `json:"kode_sub_kejadian"`
	KriteriaPenyebabKejadian string  `json:"kriteria_penyebab_kejadian"`
	Deskripsi                string  `json:"deskripsi"`
	Status                   bool    `json:"status"`
	CreatedAt                *string `json:"created_at"`
	UpdatedAt                *string `json:"updated_at"`
}

type SubIncidentListFilter struct {
	ID                       lib.NullInt64  `json:"id"`
	KodeKejadian             lib.NullString `json:"kode_kejadian"`
	PenyebabKejadian         lib.NullString `json:"penyebab_kejadian"`
	KodeSubKejadian          lib.NullString `json:"kode_sub_kejadian"`
	KriteriaPenyebabKejadian lib.NullString `json:"kriteria_penyebab_kejadian"`
	CreatedAt                lib.NullString `json:"created_at"`
	UpdatedAt                lib.NullString `json:"updated_at"`
}

type SubIncidentFilterRequest struct {
	KodeKejadian string `json:"kode_kejadian"`
}

type KodePenyebabKejadian struct {
	KodePenyebabKejadian string `json:"kode_penyebab_kejadian"`
}

type SubIncidentResponses struct {
	ID                       int64   `json:"id"`
	KodeKejadian             string  `json:"kode_kejadian"`
	PenyebabKejadian         string  `json:"penyebab_kejadian"`
	KodeSubKejadian          string  `json:"kode_sub_kejadian"`
	KriteriaPenyebabKejadian string  `json:"kriteria_penyebab_kejadian"`
	Deskripsi                string  `json:"deskripsi"`
	Status                   bool    `json:"status"`
	CreatedAt                *string `json:"created_at"`
	UpdatedAt                *string `json:"updated_at"`
}

type Paginate struct {
	Order  string `json:"order"`
	Sort   string `json:"sort"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Page   int    `json:"page"`
}

func (p SubIncidentRequest) ParseRequest() SubIncident {
	return SubIncident{
		ID:                       p.ID,
		KodeKejadian:             p.KodeKejadian,
		KodeSubKejadian:          p.KodeSubKejadian,
		KriteriaPenyebabKejadian: p.KriteriaPenyebabKejadian,
	}
}

func (p SubIncidentResponse) ParseResponse() SubIncident {
	return SubIncident{
		ID:                       p.ID,
		KodeKejadian:             p.KodeKejadian,
		KodeSubKejadian:          p.KodeSubKejadian,
		KriteriaPenyebabKejadian: p.KriteriaPenyebabKejadian,
		CreatedAt:                p.CreatedAt,
		UpdatedAt:                p.UpdatedAt,
	}
}

func (pr SubIncidentRequest) TableName() string {
	return "penyebab_kejadian_lv2"
}

func (pr SubIncidentResponse) TableName() string {
	return "penyebab_kejadian_lv2"
}
