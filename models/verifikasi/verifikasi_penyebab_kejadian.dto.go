package models

type VerifikasiPenyababKejadianRequest struct {
	ID                    int64 `json:"id"`
	VerifikasiID          int64 `json:"verifikasi_id"`
	IDPenyebabKejadian    int64 `json:"id_penyebab_kejadian"`
	IDSubPenyebabKejadian int64 `json:"id_sub_penyebab_kejadian"`
}

type VerifikasiPenyababKejadianResponse struct {
	ID                    int64 `json:"id"`
	VerifikasiID          int64 `json:"verifikasi_id"`
	IDPenyebabKejadian    int64 `json:"id_penyebab_kejadian"`
	IDSubPenyebabKejadian int64 `json:"id_sub_penyebab_kejadian"`
}

type VerifikasiPenyababKejadianDetailResponse struct {
	ID                  int64  `json:"id"`
	VerifikasiID        int64  `json:"verifikasi_id"`
	PenyebabKejadian    string `json:"penyebab_kejadian"`
	SubPenyebabKejadian string `json:"sub_penyebab_kejadian"`
}

func (vpk VerifikasiPenyababKejadianRequest) TableName() string {
	return "verifikasi_penyebab_kejadian"
}

func (vpk VerifikasiPenyababKejadianResponse) TableName() string {
	return "verifikasi_penyebab_kejadian"
}
