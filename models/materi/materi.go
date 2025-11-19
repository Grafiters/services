package models

type Materi struct {
	ID        int64
	Name      string
	UpdatedAt *string
	CreatedAt *string
}

type MateriFiles struct {
	ID       int64
	MateriID int64
	FilesID  int64
}
