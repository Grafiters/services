package models

type LiniBisnisLv2 struct {
	ID                int64
	IDLiniBisnisLv1   string
	KodeLiniBisnisLv2 string
	LiniBisnisLv2     string
	Deskripsi         string
	Status            bool
	CreatedAt         *string
	UpdatedAt         *string
}
