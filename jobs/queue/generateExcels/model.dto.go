package generateExcels

type Job struct {
	ID           int
	ReportId     int
	JSONData     []map[string]interface{}
	ColumnNames  []string
	Headers      map[string]string
	ExcelPath    string
	FileName     string
	GenerateInfo GenerateInfo
	RetryCode    string
}

type GenerateInfo struct {
	ID          int    `gorm:"column:id" json:"id"`
	NoFile      string `gorm:"column:no_file" json:"no_file"`
	ReportId    int    `gorm:"column:report_id" json:"report_id"`
	JSONPARAMS  string `gorm:"column:json_params" json:"json_params"`
	DOWNLOADURL string `gorm:"column:downloadUrl" json:"downloadUrl"`
	FILENAME    string `gorm:"column:filename" json:"filename"`
	FILEDESC    string `gorm:"column:file_desc" json:"file_desc"`
	MAKERID     string `gorm:"column:maker_id" json:"maker_id"`
	MAKERDESC   string `gorm:"column:maker_desc" json:"maker_desc"`
	RPTSTATUS   string `gorm:"column:rpt_status" json:"report_status"`
}

type DownloadUrl struct {
	DOWNLOADURL string `gorm:"column:downloadUrl" json:"downloadUrl"`
	FILEPATH    string `gorm:"column:filepath" json:"filepath"`
	FILENAME    string `gorm:"column:filename" json:"filename"`
}

func (g GenerateInfo) TableName() string {
	return "tbl_export_xls"
}

func (d DownloadUrl) TableName() string {
	return "tbl_download_link"
}
