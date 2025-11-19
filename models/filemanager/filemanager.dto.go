package filemanager

import (
	"mime/multipart"
	"net/url"
)

type FileManagerRequest struct {
	File       *multipart.FileHeader
	BucketName string `json:"bucket_name"`
	ObjectName string `json:"object_name"`
	Subdir     string `json:"subdir"`
	Filename   string `json:"filename"`
}

type FileManagerResponse struct {
	Path      string `json:"path"`
	Size      string `json:"size"`
	Filename  string `json:"filename"`
	Extension string `json:"extension"`
}

type FileManagerResponseUrl struct {
	MinioPath  string   `json:"file_path"`
	PreSignUrl *url.URL `json:"pre_sign_url"`
}
