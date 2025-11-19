package models

type EventTypeLv2 struct {
	ID               int64
	IDEventTypeLv1   string
	KodeEventTypeLv2 string
	EventTypeLv2     string
	Deskripsi        string
	Status           bool
	CreatedAt        *string
	UpdatedAt        *string
}
