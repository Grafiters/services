package models

type EventTypeLv1 struct {
	ID            int64
	KodeEventType string
	EventType     string
	Deskripsi     string
	Status        bool
	CreatedAt     *string
	UpdatedAt     *string
}
