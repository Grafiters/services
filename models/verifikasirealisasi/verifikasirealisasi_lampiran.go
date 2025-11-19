package models

type VerifikasiRealisasiFiles struct {
	ID           int64
	VerifikasiID int64
	FilesID      int64
}

func (vf VerifikasiRealisasiFiles) TableName() string {
	return "verifikasi_realisasi_lampiran"
}
