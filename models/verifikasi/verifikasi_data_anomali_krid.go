package models

type VerifikasiAnomaliDataKRID struct {
	ID           int64
	VerifikasiID int64
	Periode      string
	Object       string
	Status       bool
}

func (vad VerifikasiAnomaliDataKRID) TableName() string {
	return "verifikasi_data_anomali_krid"
}
