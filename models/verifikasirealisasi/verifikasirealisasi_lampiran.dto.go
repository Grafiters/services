package models

type VerifikasiRealisasiFilesRequest struct {
	ID           int64 `json:"id"`
	VerifikasiID int64 `json:"verifikasi_id"`
	FilesID      int64 `json:"files_id"`
}

type VerifikasiRealisasiFilesResponse struct {
	ID           int64  `json:"id"`
	VerifikasiID int64  `json:"verifikasi_id"`
	FilesID      int64  `json:"files_id"`
	Filename     string `json:"filename"`
	Path         string `json:"path"`
}

type VerifikasiRealisasiFilesResponses struct {
	IDLampiran   int64  `json:"id_lampiran"`
	VerifikasiID int64  `json:"verifikasi_id"`
	Filename     string `json:"filename"`
	Path         string `json:"path"`
	Ext          string `json:"ext"`
	Size         int64  `json:"size"`
}

func (p VerifikasiRealisasiFilesRequest) ParseRequest() VerifikasiRealisasiFiles {
	return VerifikasiRealisasiFiles{
		ID:           p.ID,
		VerifikasiID: p.VerifikasiID,
		FilesID:      p.FilesID,
	}
}

func (p VerifikasiRealisasiFilesResponse) ParseResponse() VerifikasiRealisasiFiles {
	return VerifikasiRealisasiFiles{
		ID:           p.ID,
		VerifikasiID: p.VerifikasiID,
		FilesID:      p.FilesID,
	}
}

func (vf VerifikasiRealisasiFilesRequest) TableName() string {
	return "verifikasi_realisasi_lampiran"
}

func (vf VerifikasiRealisasiFilesResponse) TableName() string {
	return "verifikasi_realisasi_lampiran"
}

func (vf VerifikasiRealisasiFilesResponses) TableName() string {
	return "verifikasi_realisasi_lampiran"
}
