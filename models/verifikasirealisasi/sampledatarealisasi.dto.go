package models

type SampleDataRealisasi struct {
	ID               int64
	VerifikasiID     int64
	DataRealisasi    string
	StatusVerifikasi bool
}

type SampleDataRealisasiRequest struct {
	ID               int64  `json:"id"`
	VerifikasiID     int64  `json:"verifikasi_id"`
	DataRealisasi    string `json:"data_realisasi"`
	StatusVerifikasi bool   `json:"status_verifikasi"`
}

type SampleDataRealisasiResponse struct {
	ID               int64  `json:"id"`
	VerifikasiID     int64  `json:"verifikasi_id"`
	DataRealisasi    string `json:"data_realisasi"`
	StatusVerifikasi bool   `json:"status_verifikasi"`
}

type RealisasiKreditKriteria struct {
	ID           int64
	VerifikasiID int64
	KriteriaID   int64
	Value        bool
}

type RealisasiKreditKriteriaRequest struct {
	ID           int64 `json:"id"`
	VerifikasiID int64 `json:"verifikasi_id"`
	KriteriaID   int64 `json:"kriteria_id"`
	Value        bool  `json:"value"`
}

type RealisasiKreditKriteriaResponse struct {
	ID           int64 `json:"id"`
	VerifikasiID int64 `json:"verifikasi_id"`
	KriteriaID   int64 `json:"kriteria_id"`
	Value        bool  `json:"value"`
}

func (v SampleDataRealisasi) TableName() string {
	return "verifikasi_data_realisasi"
}

func (v SampleDataRealisasiRequest) TableName() string {
	return "verifikasi_data_realisasi"
}
func (v SampleDataRealisasiResponse) TableName() string {
	return "verifikasi_data_realisasi"
}

func (v RealisasiKreditKriteria) TableName() string {
	return "verifikasi_realisasi_kredit_kriteria"
}

func (v RealisasiKreditKriteriaRequest) TableName() string {
	return "verifikasi_realisasi_kredit_kriteria"
}
func (v RealisasiKreditKriteriaResponse) TableName() string {
	return "verifikasi_realisasi_kredit_kriteria"
}
