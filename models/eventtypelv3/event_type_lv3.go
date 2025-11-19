package models

type EventTypeLv3 struct {
	ID               int64
	IDEventTypeLv2   string
	KodeEventTypeLv3 string
	EventTypeLv3     string
	Deskripsi        string
	Status           bool
	CreatedAt        *string
	UpdatedAt        *string
}
