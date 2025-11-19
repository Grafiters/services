package models

type JobMonitoringRequest struct {
	Order        string `json:"order"`
	Sort         string `json:"sort"`
	Offset       int    `json:"offset"`
	Limit        int    `json:"limit"`
	Page         int    `json:"page"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
	NamaJob      string `json:"nama_job"`
	StatusProses string `json:"status_proses"`
}

type JobMonitoringResponse struct {
	Tanggal         *string `json:"tanggal"`
	NamaJob         string  `json:"nama_job"`
	Proses          string  `json:"proses"`
	StatusProses    string  `json:"status_proses"`
	DeskripsiStatus string  `json:"deskripsi_status"`
}

type SearchNamaJobReq struct {
	Keyword string `json:"keyword"`
	Limit   int64  `json:"limit"`
}

type SearchNamaJobRes struct {
	Name string `json:"name"`
}
