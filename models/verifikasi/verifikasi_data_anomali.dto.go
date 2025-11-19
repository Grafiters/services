package models

type VerifikasiAnomaliDataRequest struct {
	ID              int64   `json:"id"`
	VerifikasiID    int64   `json:"verifikasi_id"`
	TanggalKejadian *string `json:"tanggal_kejadian"`
	NomorRekening   string  `json:"nomor_rekening"`
	Nominal         int64   `json:"nominal"`
	Keterangan      string  `json:"keterangan"`
	// CreatedAt       *string `json:"created_at"`
	// UpdatedAt       *string `json:"updated_at"`
}

type VerifikasiAnomaliDataResponse struct {
	ID              int64   `json:"id"`
	VerifikasiID    int64   `json:"verifikasi_id"`
	TanggalKejadian *string `json:"tanggal_kejadian"`
	NomorRekening   string  `json:"nomor_rekening"`
	Nominal         int64   `json:"nominal"`
	Keterangan      string  `json:"keterangan"`
	// CreatedAt       *string `json:"created_at"`
	// UpdatedAt       *string `json:"updated_at"`
}

type VerifikasiAnomaliDataResponses struct {
	ID              int64   `json:"id"`
	VerifikasiID    int64   `json:"verifikasi_id"`
	TanggalKejadian *string `json:"tanggal_kejadian"`
	NomorRekening   string  `json:"nomor_rekening"`
	Nominal         float64 `json:"nominal"`
	Keterangan      string  `json:"keterangan"`
}

func (vad VerifikasiAnomaliDataRequest) TableName() string {
	return "verifikasi_data_anomali"
}

func (vad VerifikasiAnomaliDataResponse) TableName() string {
	return "verifikasi_data_anomali"
}
