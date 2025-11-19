package models

type TasklistsAnomaliDataKRIDRequest struct {
	TasklistID int64  `json:"tasklist_id"`
	Object     string `json:"object"`
}

type TasklistDataAnomaliRequest struct {
	TasklistID string `json:"tasklist_id" example:"1"`
}

type TasklistDataAnomaliResponse struct {
	Object string `json:"object"`
}

type TasklistsAnomaliDataKRIDDelete struct {
	TasklistID int64 `json:"tasklist_id"`
}

func (TasklistsAnomaliDataKRIDRequest) TableName() string {
	return "tasklists_data_anomali_krid"
}

func (TasklistsAnomaliDataKRIDDelete) TableName() string {
	return "tasklists_data_anomali_krid"
}
