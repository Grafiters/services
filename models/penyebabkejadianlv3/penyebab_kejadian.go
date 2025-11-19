package models

type PenyebabKejadianLv3 struct {
	ID                      int64
	KodeSubKejadian         string
	KodePenyebabKejadianLv3 string
	PenyebabKejadianLv3     string
	Deskripsi               string
	Status                  bool
	CreatedAt               *string
	UpdatedAt               *string
}
