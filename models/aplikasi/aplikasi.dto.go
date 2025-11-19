package models

type AplikasiRequest struct {
	ID              int64   `json:"id"`
	Kode            string  `json:"kode"`
	ApplicationName string  `json:"application_name"`
	CreatedAt       *string `json:"CreatedAt"`
	UpdatedAt       *string `json:"UpdatedAt"`
}

type AplikasiResponse struct {
	ID              int64   `json:"id"`
	Kode            string  `json:"kode"`
	ApplicationName string  `json:"application_name"`
	CreatedAt       *string `json:"CreatedAt"`
	UpdatedAt       *string `json:"UpdatedAt"`
}

func (p AplikasiRequest) ParseRequest() Aplikasi {
	return Aplikasi{
		ID:              p.ID,
		Kode:            p.Kode,
		ApplicationName: p.ApplicationName,
	}
}

func (p AplikasiResponse) ParseResponse() Aplikasi {
	return Aplikasi{
		ID:              p.ID,
		Kode:            p.Kode,
		ApplicationName: p.ApplicationName,
		CreatedAt:       p.CreatedAt,
		UpdatedAt:       p.UpdatedAt,
	}
}

func (apr AplikasiRequest) TableName() string {
	return "aplikasi"
}

func (apr AplikasiResponse) TableName() string {
	return "aplikasi"
}
