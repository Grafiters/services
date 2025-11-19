package models

type LampiranIndicatorRequest struct {
	ID            int64  `json:"id"`
	IDIndicator   int64  `json:"id_indicator"`
	NamaLampiran  string `json:"nama_lampiran"`
	NomorLampiran string `json:"nomor_lampiran"`
	JenisFile     string `json:"jenis_file"`
	Path          string `json:"path"`
	Filename      string `json:"filename"`
}

type LampiranIndicatorResponse struct {
	ID            int64  `json:"id"`
	IDIndicator   int64  `json:"id_indicator"`
	NamaLampiran  string `json:"nama_lampiran"`
	NomorLampiran string `json:"nomor_lampiran"`
	JenisFile     string `json:"jenis_file"`
	Path          string `json:"path"`
	Filename      string `json:"filename"`
}

func (p LampiranIndicatorRequest) ParseRequest() LampiranIndicator {
	return LampiranIndicator{
		ID:            p.ID,
		IDIndicator:   p.IDIndicator,
		NamaLampiran:  p.NamaLampiran,
		NomorLampiran: p.NomorLampiran,
		JenisFile:     p.JenisFile,
		Path:          p.Path,
		Filename:      p.Filename,
	}
}

func (p LampiranIndicatorResponse) ParseResponse() LampiranIndicator {
	return LampiranIndicator{
		ID:            p.ID,
		IDIndicator:   p.IDIndicator,
		NamaLampiran:  p.NamaLampiran,
		NomorLampiran: p.NomorLampiran,
		JenisFile:     p.JenisFile,
		Path:          p.Path,
		Filename:      p.Filename,
	}
}

func (li LampiranIndicatorRequest) TableName() string {
	return "risk_indicator_map_files"
}

func (li LampiranIndicatorResponse) TableName() string {
	return "risk_indicator_map_files"
}
