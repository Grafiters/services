package models

type VerifikasiFilesRequest struct {
	ID           int64 `json:"id"`
	VerifikasiID int64 `json:"verifikasi_id"`
	FilesID      int64 `json:"files_id"`
	// UpdatedAt    *string `json:"updated_at"`
	// CreatedAt    *string `json:"created_at"`
}

type VerifikasiFilesResponse struct {
	ID           int64 `json:"id"`
	VerifikasiID int64 `json:"verifikasi_id"`
	FilesID      int64 `json:"files_id"`
	// UpdatedAt    *string `json:"updated_at"`
	// CreatedAt    *string `json:"created_at"`
}

type VerifikasiFilesResponses struct {
	IDLampiran   int64  `json:"id_lampiran"`
	VerifikasiID int64  `json:"verifikasi_id"`
	Filename     string `json:"filename"`
	Path         string `json:"path"`
	Ext          string `json:"ext"`
	Size         int64  `json:"size"`
}

func (p VerifikasiFilesRequest) ParseRequest() VerifikasiFiles {
	return VerifikasiFiles{
		ID:           p.ID,
		VerifikasiID: p.VerifikasiID,
		FilesID:      p.FilesID,
	}
}

func (p VerifikasiFilesResponse) ParseResponse() VerifikasiFiles {
	return VerifikasiFiles{
		ID:           p.ID,
		VerifikasiID: p.VerifikasiID,
		FilesID:      p.FilesID,
	}
}

func (vf VerifikasiFilesRequest) TableName() string {
	return "verifikasi_lampiran"
}

func (vf VerifikasiFilesResponse) TableName() string {
	return "verifikasi_lampiran"
}

func (vf VerifikasiFilesResponses) TableName() string {
	return "verifikasi_lampiran"
}
