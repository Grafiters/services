package models

type VerifikasiUsulanPerbaikan struct {
	ID           int64
	VerifikasiID int64
	Usulan       string
	Deskripsi    string
	Aplikasi     *string
}

func (v VerifikasiUsulanPerbaikan) TableName() string {
	return "verifikasi_usulan_perbaikan"
}
