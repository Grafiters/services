package models

type VerifikasiQuestionnerRequest struct {
	ID                   int64  `json:"id"`
	VerifikasiID         int64  `json:"verifikasi_id"`
	Questionner          string `json:"questionner"`
	DataSumber           string `json:"data_sumber"`
	Checker              string `json:"checker"`
	Signer               string `json:"signer"`
	ApprovalOrd          string `json:"approval_ord"`
	JenisFraud           string `json:"jenis_fraud"`
	StatusValidasiRmc    string `json:"status_validasi_rmc"`
	StatusValidasiSigner string `json:"status_validasi_signer"`
	StatusValidasiOrd    string `json:"status_validasi_ord"`
}

type VerifikasiQuestionnerResponse struct {
	ID                   int64  `json:"id"`
	VerifikasiID         int64  `json:"verifikasi_id"`
	Questionner          string `json:"questionner"`
	DataSumber           string `json:"data_sumber"`
	Checker              string `json:"checker"`
	Signer               string `json:"signer"`
	ApprovalOrd          string `json:"approval_ord"`
	JenisFraud           string `json:"jenis_fraud"`
	StatusValidasiRmc    string `json:"status_validasi_rmc"`
	StatusValidasiSigner string `json:"status_validasi_signer"`
	StatusValidasiOrd    string `json:"status_validasi_ord"`
}
