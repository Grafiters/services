package models

type VerifikasiFiles struct {
	ID           int64
	VerifikasiID int64
	FilesID      int64
	// UpdatedAt    *string
	// CreatedAt    *string
}

func (vf VerifikasiFiles) TableName() string {
	return "verifikasi_lampiran"
}
