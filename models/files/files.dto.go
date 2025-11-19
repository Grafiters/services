package files

type FilesRequest struct {
	ID        int64   `json:"id"`
	Filename  string  `json:"filename"`
	Path      string  `json:"path"`
	Extension string  `json:"extension"`
	Size      int64   `json:"size"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

type FilesResponse struct {
	ID        int64   `json:"id"`
	Filename  string  `json:"filename"`
	Path      string  `json:"path"`
	Extension string  `json:"extension"`
	Size      int64   `json:"size"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

type FilesResponses struct {
	ID        int64  `json:"id"`
	Filename  string `json:"filename"`
	Path      string `json:"path"`
	Extension string `json:"extension"`
	Size      int64  `json:"size"`
}

func (p FilesRequest) ParseRequest() Files {
	return Files{
		ID:        p.ID,
		Filename:  p.Filename,
		Path:      p.Path,
		Extension: p.Extension,
		Size:      p.Size,
	}
}

func (p FilesResponse) ParseResponse() Files {
	return Files{
		ID:        p.ID,
		Filename:  p.Filename,
		Path:      p.Path,
		Extension: p.Extension,
		Size:      p.Size,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func (i FilesRequest) TableName() string {
	return "files"
}

func (i FilesResponse) TableName() string {
	return "files"
}
