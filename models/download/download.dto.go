package models

type DownloadRequest struct {
	ReportId      int    `json:"report_id"`
	JSONPARAMS    string `json:json_params`
	PERNR         string `json:pernr`
	RequestStatus string `json:"request_status"`
}

type DownloadResponse struct {
	ID          int    `gorm:"column:id" json:"id"`
	NoFile      string `gorm:"column:no_file" json:"no_file"`
	ReportId    int    `gorm:"column:report_id" json:"report_id"`
	JSONPARAMS  string `gorm:"column:json_params" json:"json_params"`
	DOWNLOADURL string `gorm:"column:downloadUrl" json:"downloadUrl"`
	FILENAME    string `gorm:"column:filename" json:"filename"`
	MAKERID     string `gorm:"column:maker_id" json:"maker_id"`
	MAKERDESC   string `gorm:"column:maker_desc" json:"maker_desc"`
}

type DownloadUrl struct {
	DOWNLOADURL string `gorm:"column:downloadUrl" json:"downloadUrl"`
	FILEPATH    string `gorm:"column:filepath" json:"filepath"`
	FILENAME    string `gorm:"column:filename" json:"filename"`
}

type ListDownloadRequest struct {
	NamaLaporan string `json:"nama_laporan"`
	Kanwil      string `json:"kanwil"`
	Kanca       string `json:"kanca"`
	UnitKerja   string `json:"unit_kerja"`
	PeriodeData string `json:"periode_data"`
	MakerId     string `json:"maker_id"`
	Limit       int    `json:"limit"`
	Offset      int    `json:"offset"`
}

type ListDownloadResponse struct {
	No          int64   `json:"no"`
	Id          int64   `json:"id"`
	ReportId    int64   `json:"report_id"`
	NamaLaporan string  `json:"nama_laporan"`
	Kanwil      string  `json:"kanwil"`
	Kanca       string  `json:"kanca"`
	UnitKerja   string  `json:"unit_kerja"`
	PeriodeData string  `json:"periode_data"`
	Status      string  `json:"status"`
	FileDesc    string  `json:"file_desc"`
	Filename    *string `json:"filename"`
	Filepath    *string `json:"filepath"`
	JsonParams  *string `json:"json_params"`
}

type RetryRequest struct {
	InsertId   int    `json:"insert_id"`
	ReportId   int    `json:"report_id"`
	Filename   string `json:"filename"`
	JsonParams string `json:"json_params"`
	RetryCode  string `json:"retry_code"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
	Pernr      string `json:"pernr"`
	Timetstamp string `json:"timestamp"`
}

type ReportTypeResponse struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type GeneratorTemplate struct {
	ReportId   int64  `json:"report_id"`
	Pernr      string `json:"pernr"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"EndDate"`
	Timestamp  string `json:"Timestamp"`
	JsonParams string `json:"jsonparams"`
}

func (d DownloadResponse) TableName() string {
	return "tbl_export_xls"
}

func (d DownloadUrl) TableName() string {
	return "tbl_download_link"
}
