package models

type VerifikasiAnomaliDataKRIDRequest struct {
	ID           int64  `json:"id"`
	VerifikasiID int64  `json:"verifikasi_id"`
	Periode      string `json:"periode"`
	Object       string `json:"object"`
	Status       bool   `json:"status"`
}

type VerifikasiAnomaliDataKRIDResponses struct {
	ID           int64  `json:"id"`
	VerifikasiID int64  `json:"verifikasi_id"`
	Periode      string `json:"periode"`
	Object       string `json:"object"`
	Status       bool   `json:"status"`
}

func (vad VerifikasiAnomaliDataKRIDRequest) TableName() string {
	return "verifikasi_data_anomali_krid"
}

func (vad VerifikasiAnomaliDataKRIDResponses) TableName() string {
	return "verifikasi_data_anomali_krid"
}
