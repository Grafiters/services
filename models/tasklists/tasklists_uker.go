package models

type TasklistsUker struct {
	REGION          string `json:"region"`
	RGDESC          string `json:"rgdesc"`
	MAINBR          string `json:"mainbr"`
	MBDESC          string `json:"mbdesc"`
	BRANCH          string `json:"branch"`
	BRDESC          string `json:"brdesc"`
	JumlahNominatif int64  `json:"jumlah_nominatif"`
	TasklistID      int64  `json:"tasklist_id"`
}

type TasklistsUkerData struct {
	REGION          string `json:"region"`
	RGDESC          string `json:"rgdesc"`
	MAINBR          string `json:"mainbr"`
	MBDESC          string `json:"mbdesc"`
	BRANCH          string `json:"branch"`
	BRDESC          string `json:"brdesc"`
	JumlahNominatif int64  `json:"jumlah_nominatif"`
	TasklistID      int64  `json:"tasklist_id"`
}

func (TasklistsUker) TableName() string {
	return "tasklists_uker"
}
