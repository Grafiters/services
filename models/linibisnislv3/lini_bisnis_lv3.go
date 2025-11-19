package models

type LiniBisnisLv3 struct {
	ID                int64
	IDLiniBisnisLv2   string
	KodeLiniBisnisLv3 string
	LiniBisnisLv3     string
	Deskripsi         string
	Status            bool
	CreatedAt         *string
	UpdatedAt         *string
}
