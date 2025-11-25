package models

type RiskControl struct {
	ID          int64
	Kode        string
	RiskControl string
	ControlType string
	Nature      string
	KeyControl  string
	Deskripsi   string
	OwnerLvl    string
	OwnerGroup  string
	Owner       string
	Document    string
	Status      bool
	CreatedAt   *string
	UpdatedAt   *string
}
