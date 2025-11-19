package models

type VerifikasiPICTindakLanjutRequest struct {
	ID                    int64  `json:"id"`
	VerifikasiID          int64  `json:"verifikasi_id"`
	PICID                 string `json:"pic_id"`
	PICDetail             string `json:"pic_detail"`
	TanggalTindakLanjut   string `json:"tanggal_tindak_lanjut"`
	DeskripsiTindakLanjut string `json:"deskripsi_tindak_lanjut"`
	Status                string `json:"status"`
	// CreatedAt             *string `json:"created_at"`
	// UpdatedAt             *string `json:"updated_at"`
}

type VerifikasiPICTindakLanjutResponse struct {
	ID                    int64  `json:"id"`
	VerifikasiID          int64  `json:"verifikasi_id"`
	PICID                 string `json:"pic_id"`
	PICDetail             string `json:"pic_detail"`
	TanggalTindakLanjut   string `json:"tanggal_tindak_lanjut"`
	DeskripsiTindakLanjut string `json:"deskripsi_tindak_lanjut"`
	Status                string `json:"status"`
	// CreatedAt             *string `json:"created_at"`
	// UpdatedAt             *string `json:"updated_at"`
}

type VerifikasiPICTindakLanjutResponses struct {
	ID                    int64  `json:"id"`
	VerifikasiID          int64  `json:"verifikasi_id"`
	PICID                 string `json:"pic_id"`
	PICDetail             string `json:"pic_detail"`
	TanggalTindakLanjut   string `json:"tanggal_tindak_lanjut"`
	DeskripsiTindakLanjut string `json:"deskripsi_tindak_lanjut"`
	Status                string `json:"status"`
}

func (vp VerifikasiPICTindakLanjutRequest) TableName() string {
	return "verifikasi_pic_tindak_lanjut"
}

func (vp VerifikasiPICTindakLanjutResponse) TableName() string {
	return "verifikasi_pic_tindak_lanjut"
}
