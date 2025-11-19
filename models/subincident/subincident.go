package models

type SubIncident struct {
	ID                       int64
	KodeKejadian             string
	KodeSubKejadian          string
	KriteriaPenyebabKejadian string
	Deskripsi                string
	Status                   bool
	CreatedAt                *string
	UpdatedAt                *string
}
