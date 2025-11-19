package models

type RiskControl struct {
	ID          int64
	Kode        string
	RiskControl string
	Deskripsi   string
	Status      bool
	CreatedAt   *string
	UpdatedAt   *string
}
