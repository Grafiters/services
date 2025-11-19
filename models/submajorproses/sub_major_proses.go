package models

type SubMajorProses struct {
	ID                 int64
	IDMajorProses      string
	MajorProses		   string
	KodeSubMajorProses string
	SubMajorProses     string
	Deskripsi          string
	Status             bool
	CreatedAt          *string
	UpdatedAt          *string
}
