package models

import "riskmanagement/models/files"

type LinkcageRequest struct {
	ID        int64   `json:"id"`
	Name      string  `json:"name"`
	URL       string  `json:"url"`
	Status    string  `json:"status"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`

	Files []files.FilesRequest `json:"files"`

	Order  string `json:"order"`
	Sort   string `json:"sort"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Page   int    `json:"page"`

	PERNR string `json:"pernr"`
}

type Linkcage struct {
	ID        int64
	Name      string
	URL       string
	Status    string
	CreatedAt *string
	UpdatedAt *string
}

type LinkcageResponse struct {
	ID        int64                   `json:"id"`
	Name      string                  `json:"name"`
	URL       string                  `json:"url"`
	Banner    string                  `json:"banner"`
	Files     []LinkcageImageResponse `json:"files"`
	Status    string                  `json:"status"`
	CreatedAt *string
	UpdatedAt *string
}

func (l Linkcage) TableName() string {
	return "linkcage"
}

func (l LinkcageRequest) TableName() string {
	return "linkcage"
}

type ActiveLinkcage struct {
	ID       int64
	Name     string
	URL      string
	Filename string `json:"filename"`
	Path     string `json:"path"`
	Ext      string `json:"ext"`
	Size     int64  `json:"size"`
	Status   string
}

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
