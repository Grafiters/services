package models

type LinkcageImage struct {
	LinkcageID int64 `json:"linkcage_id"`
	FileID     int64 `json:"file_id"`
	CreatedAt  *string
}

type LinkcageImageResponse struct {
	Filename string `json:"filename"`
	Path     string `json:"path"`
	Ext      string `json:"extension"`
	Size     int64  `json:"size"`
}

func (li LinkcageImage) TableName() string {
	return "linkcage_image"
}
