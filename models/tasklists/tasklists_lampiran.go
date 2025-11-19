package models

type TasklistsFiles struct {
	TasklistsID int64
	FilesID     int64
	CreatedAt   *string
}

type Lampiran struct {
	isi_kolom string
}

type TasklistFilesResponses struct {
	IDLampiran int64  `json:"id_lampiran"`
	TasklistID int64  `json:"tasklist_id"`
	Filename   string `json:"filename"`
	Path       string `json:"path"`
	Ext        string `json:"ext"`
	Size       int64  `json:"size"`
}

type TasklistFileResponses struct {
	IDLampiran int64  `json:"id_lampiran"`
	TasklistID int64  `json:"tasklist_id"`
	Filename   string `json:"filename"`
	Path       string `json:"path"`
	Ext        string `json:"ext"`
	Size       int64  `json:"size"`
}

func (TasklistsFiles) TableName() string {
	return "tasklists_lampiran"
}
