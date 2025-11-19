package models

type VerifikasiUsulanPerbaikanRequest struct {
	ID           int64   `json:"id"`
	VerifikasiID int64   `json:"verifikasi_id"`
	Usulan       string  `json:"usulan"`
	Deskripsi    string  `json:"deskripsi"`
	Aplikasi     *string `json:"aplikasi"`
}

type VerifikasiUsulanPerbaikanResponse struct {
	ID           int64   `json:"id"`
	VerifikasiID int64   `json:"verifikasi_id"`
	Usulan       string  `json:"usulan"`
	Deskripsi    string  `json:"deskripsi"`
	Aplikasi     *string `json:"aplikasi"`
}

func (v VerifikasiUsulanPerbaikanRequest) TableName() string {
	return "verifikasi_usulan_perbaikan"
}

func (v VerifikasiUsulanPerbaikanResponse) TableName() string {
	return "verifikasi_usulan_perbaikan"
}
