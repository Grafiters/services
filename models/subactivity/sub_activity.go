package models

type SubActivity struct {
	ID              int64
	ActivityID      int64
	KodeSubActivity string
	Name            string
	CreatedAt       *string
	UpdatedAt       *string
}
