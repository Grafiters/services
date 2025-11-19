package models

type CoachingMapPesertaRequest struct {
	ID          int64  `json:"id"`
	IDCoaching  int64  `json:"id_coaching"`
	PERNR       string `json:"pernr"`
	NamaPeserta string `json:"nama_peserta"`
	SteelTx     string `json:"steel_tx"`
}

type CoachingMapPesertaResponse struct {
	ID          int64  `json:"id"`
	IDCoaching  int64  `json:"id_coaching"`
	PERNR       string `json:"pernr"`
	NamaPeserta string `json:"nama_peserta"`
	SteelTx     string `json:"steel_tx"`
}

func (co CoachingMapPesertaRequest) TableName() string {
	return "coaching_map_peserta"
}

func (co CoachingMapPesertaResponse) TableName() string {
	return "coaching_map_peserta"
}
