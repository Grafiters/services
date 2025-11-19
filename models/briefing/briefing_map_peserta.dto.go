package models

type BriefingMapPesertaRequest struct {
	ID          int64  `json:"id"`
	IDBriefing  int64  `json:"id_briefing"`
	PERNR       string `json:"pernr"`
	NamaPeserta string `json:"nama_peserta"`
	SteelTx     string `json:"steel_tx"`
}

type BriefingMapPesertaResponse struct {
	ID          int64  `json:"id"`
	IDBriefing  int64  `json:"id_briefing"`
	PERNR       string `json:"pernr"`
	NamaPeserta string `json:"nama_peserta"`
	SteelTx     string `json:"steel_tx"`
}

func (brief BriefingMapPesertaRequest) TableName() string {
	return "briefing_map_peserta"
}

func (brief BriefingMapPesertaResponse) TableName() string {
	return "briefing_map_peserta"
}
