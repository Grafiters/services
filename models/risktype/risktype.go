package models

type RiskType struct {
	ID           int64
	RiskTypeCode string
	RiskType     string
	Deskripsi    string
	Status       bool
	CreatedAt    *string
	UpdatedAt    *string
}
