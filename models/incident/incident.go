package models

type Incident struct {
	ID               int64
	KodeKejadian     string
	PenyebabKejadian string
	Deskripsi        string
	Status           bool
	CreatedAt        *string
	UpdatedAt        *string
}
