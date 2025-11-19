package models

type VerifikasiPenyababKejadian struct {
	ID                    int64
	VerifikasiID          int64
	IDPenyebabKejadian    int64
	IDSubPenyebabKejadian int64
}

func (vpk VerifikasiPenyababKejadian) TableName() string {
	return "verifikasi_penyebab_kejadian"
}
