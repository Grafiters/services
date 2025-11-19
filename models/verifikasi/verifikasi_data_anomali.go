package models

type VerifikasiAnomaliData struct {
	ID              int64
	VerifikasiID    int64
	TanggalKejadian *string
	NomorRekening   string
	Nominal         int64
	Keterangan      string
	// CreatedAt       *string
	// UpdatedAt       *string
}

func (vad VerifikasiAnomaliData) TableName() string {
	return "verifikasi_data_anomali"
}
