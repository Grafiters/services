package models

type MstKriteriaRequest struct {
	ID           int64   `json:"id"`
	KodeCriteria string  `json:"kode_criteria"`
	Criteria     string  `json:"criteria"`
	Restruck     int8    `json:"restruck"`
	Status       int8    `json:"status"`
	ActiveDate   *string `json:"active_date"`
	InactiveDate *string `json:"inactive_date"`
	CreatedAt    *string `json:"created_at"`
	CreatedBy    string  `json:"created_by"`
	CreatedDesc  string  `json:"created_desc"`
	EnabledDate  *string `json:"enabled_date"`
	EnabledBy    string  `json:"enabled_by"`
	EnabledDesc  string  `json:"enabled_desc"`
	DisabledDate *string `json:"disabled_date"`
	DisabledBy   string  `json:"disabled_by"`
	DisabledDesc string  `json:"disabled_desc"`
}

type FilterRequest struct {
	Order    string `json:"order"`
	Sort     string `json:"sort"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
	Page     int    `json:"page"`
	Keyword  string `json:"keyword"`
	Restruck string `json:"restruck"`
}

type MstKriteriaResponse struct {
	ID           int64   `json:"id"`
	KodeCriteria string  `json:"kode_criteria"`
	Criteria     string  `json:"criteria"`
	Restruck     int8    `json:"restruck"`
	Status       int8    `json:"status"`
	ActiveDate   *string `json:"active_date"`
	InactiveDate *string `json:"inactive_date"`
	CreatedAt    *string `json:"created_at"`
	CreatedBy    string  `json:"created_by"`
	CreatedDesc  string  `json:"created_desc"`
	EnabledDate  *string `json:"enabled_date"`
	EnabledBy    string  `json:"enabled_by"`
	EnabledDesc  string  `json:"enabled_desc"`
	DisabledDate *string `json:"disabled_date"`
	DisabledBy   string  `json:"disabled_by"`
	DisabledDesc string  `json:"disabled_desc"`
}

type KodeMstKriteria struct {
	KodeCriteria string `json:"kode_criteria"`
}

func (p MstKriteriaRequest) ParseRequest() MstKriteria {
	return MstKriteria{
		ID:           p.ID,
		KodeCriteria: p.KodeCriteria,
		Criteria:     p.Criteria,
		Restruck:     p.Restruck,
		Status:       p.Status,
		ActiveDate:   p.ActiveDate,
		InactiveDate: p.InactiveDate,
		CreatedAt:    p.CreatedAt,
		CreatedBy:    p.CreatedBy,
		CreatedDesc:  p.CreatedDesc,
		EnabledDate:  p.EnabledDate,
		EnabledBy:    p.EnabledBy,
		EnabledDesc:  p.EnabledDesc,
		DisabledDate: p.DisabledDate,
		DisabledBy:   p.DisabledBy,
		DisabledDesc: p.DisabledDesc,
	}
}

func (p MstKriteriaResponse) ParseResponse() MstKriteria {
	return MstKriteria{
		ID:           p.ID,
		KodeCriteria: p.KodeCriteria,
		Criteria:     p.Criteria,
		Restruck:     p.Restruck,
		Status:       p.Status,
		ActiveDate:   p.ActiveDate,
		InactiveDate: p.InactiveDate,
		CreatedAt:    p.CreatedAt,
		CreatedBy:    p.CreatedBy,
		CreatedDesc:  p.CreatedDesc,
		EnabledDate:  p.EnabledDate,
		EnabledBy:    p.EnabledBy,
		EnabledDesc:  p.EnabledDesc,
		DisabledDate: p.DisabledDate,
		DisabledBy:   p.DisabledBy,
		DisabledDesc: p.DisabledDesc,
	}
}

func (mk MstKriteriaRequest) TableName() string {
	return "mst_kriteria"
}

func (mk MstKriteriaResponse) TableName() string {
	return "mst_kriteria"
}

// add by panji 18/12/2024
type CriteriaRequestById struct {
	Id       string `json:"id"`
	Restruck string `json:"restruck"`
}

type PeriodeRequest struct {
	TglAwal  string `json:"tgl_awal"`
	TglAkhir string `json:"tgl_akhir"`
}

type MstKriteriaHistoryRequest struct {
	ID           int64   `json:"id"`
	IdCriteria   int64   `json:"id_criteria"`
	KodeCriteria string  `json:"kode_criteria"`
	Criteria     string  `json:"criteria"`
	Restruck     int8    `json:"restruck"`
	Status       int8    `json:"status"`
	CreatedAt    *string `json:"created_at"`
	CreatedBy    *string `json:"created_by"`
	CreatedDesc  *string `json:"created_desc"`
}

type MstKriteriaHistoryResponses struct {
	ID           int64   `json:"id"`
	IdCriteria   int64   `json:"id_criteria"`
	KodeCriteria string  `json:"kode_criteria"`
	Criteria     string  `json:"criteria"`
	Restruck     int8    `json:"restruck"`
	Status       int8    `json:"status"`
	CreatedAt    *string `json:"created_at"`
	CreatedBy    *string `json:"created_by"`
	CreatedDesc  *string `json:"created_desc"`
}

func (mk MstKriteriaHistory) TableName() string {
	return "mst_kriteria_history"
}

func (mk MstKriteriaHistoryRequest) TableName() string {
	return "mst_kriteria_history"
}
