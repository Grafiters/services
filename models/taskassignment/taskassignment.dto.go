package models

type CreateTaskDTO struct {
	NoTasklist      string  `json:"no_tasklist" form:"no_tasklist" example:"Test Create something"`
	NamaTasklist    string  `json:"nama_tasklist" form:"nama_tasklist" example:"Test Create something"`
	RiskIndicator   string  `json:"risk_indicator" form:"risk_indicator" example:"Test Create something"`
	ActivityID      int64   `json:"activity_id" form:"activity_id" example:"1"`
	ProductID       int64   `json:"product_id" form:"product_id" example:"1"`
	ProductName     string  `json:"product_name" form:"product_name" example:"Test Create something"`
	RiskIssueID     int64   `json:"risk_issue_id" form:"risk_issue_id" example:"1"`
	RiskIssue       string  `json:"risk_issue" form:"risk_issue" example:"Test Create something"`
	RiskIndicatorID int64   `json:"risk_indicator_id" form:"risk_indicator_id" example:"1"`
	TaskType        int64   `json:"task_type" form:"task_type" example:"1"`
	TaskTypeName    string  `json:"task_type_name" form:"task_type_name" example:"Test Create something"`
	Kegiatan        string  `json:"kegiatan" form:"kegiatan" example:"Test Create something"`
	Period          string  `json:"period" form:"period" example:"Test Create something"`
	EnableRange     string  `json:"enable_range" form:"enable_range" example:"Quarter, Semester, Angka, Montly"`
	SumberData      string  `json:"sumber_data" form:"sumber_data" example:"Test Create something"`
	RangeDate       string  `json:"range_date" form:"range_date" example:"1"`
	StartDate       *string `json:"start_date" form:"start_date" example:"Test Create something"`
	EndDate         *string `json:"end_date" form:"end_date" example:"Test Create something"`
	RAP             bool    `json:"rap" form:"rap" example:"true"`
	Sample          int64   `json:"sample" form:"sample" example:"1"`
	Validation      string  `json:"validation" form:"validation" example:"Test Create something"`
	ValidationName  string  `json:"validation_name" form:"validation_name" example:"Test Create something"`
	Approval        string  `json:"approval" form:"approval" example:"Test Create something"`
	ApprovalName    string  `json:"approval_name" form:"approval_name" example:"Test Create something"`
	MakerID         string  `json:"maker_id" form:"maker_id" example:"Test Create something"`
}

type TaskFile struct {
	TasklistsID int64  `json:"tasklists_id"`
	FilesID     int64  `json:"files_id"`
	CreatedAt   string `json:"created_at"`
}

type CheckTableRequest struct {
	RiskIssue     int64 `json:"risk_issue_id"`
	RiskIndicator int64 `json:"risk_indicator_id"`
}

type TemplateResponse struct {
	Status  string `json:"status"`
	Columns string `json:"columns"`
}

type NoTaskRequest struct {
	Orgeh string `json:"orgeh"`
	Pernr string `json:"pernr"`
}

type TaskDetailRequest struct {
	ID    int64  `json:"id"`
	Pernr string `json:"pernr"`
}

type TaskFilterRequest struct {
	Kanwil         string `json:"kanwil" form:"kanwil" example:"Test"`
	Kanca          string `json:"kanca" form:"kanca" example:"Test"`
	Uker           string `json:"uker" form:"uker" example:"Test"`
	Aktifitas      int64  `json:"aktifitas" form:"aktifitas" example:"1"`
	Produck        int64  `json:"product" form:"product" example:"1"`
	RiskEvent      int64  `json:"risk_event" form:"risk_event" example:"1"`
	RiskIndicator  int64  `json:"risk_indicator" form:"risk_indicator" example:"1"`
	TaskType       int64  `json:"task_type" form:"task_type" example:"1"`
	StatusApproval string `json:"status_approval" form:"status_approval" example:"Test"`
	StatusDocument string `json:"status_document" form:"status_document" example:"Test"`
	StatusTask     string `json:"status_task" form:"status_task" example:"Test"`
	Pernr          string `json:"pernr" form:"pernr" example:"Test"`
	Kostl          string `json:"kostl" form:"kostl" example:"Test"`
	Branches       string `json:"branches" form:"branches" example:"Test"`
	TglAwal        string `json:"tgl_awal" form:"tgl_awal" example:"2025-02-01"`
	TglAkhir       string `json:"tgl_akhir" form:"tgl_akhir" example:"2025-02-10"`
	Limit          int64  `json:"limit" form:"limit" example:"10"`
	Offset         int64  `json:"offset" form:"offset" example:"0"`
	Page           int64  `json:"page" form:"page" example:"1"`
}

type TaskApprovalRequest struct {
	Kanwil         string `json:"kanwil" form:"kanwil" example:"Test Kanwil"`
	Kanca          string `json:"kanca" form:"kanca" example:"Test Kanca"`
	Uker           string `json:"uker" form:"uker" example:"Test Uker"`
	Aktifitas      int64  `json:"aktifitas" form:"aktifitas" example:"1"`
	Produck        int64  `json:"product" form:"product" example:"1"`
	RiskEvent      int64  `json:"risk_event" form:"risk_event" example:"1"`
	RiskIndicator  int64  `json:"risk_indicator" form:"risk_indicator" example:"1"`
	TaskType       int64  `json:"task_type" form:"task_type" example:"1"`
	StatusApproval string `json:"status_approval" form:"status_approval" example:"Approved"`
	StatusDocument string `json:"status_document" form:"status_document" example:"Completed"`
	StatusTask     string `json:"status_task" form:"status_task" example:"In Progress"`
	Pernr          string `json:"pernr" form:"pernr" example:"P0000001"`
	Kostl          string `json:"kostl" form:"kostl" example:"K0001"`
	Branches       string `json:"branches" form:"branches" example:"Branch1"`
	Validator      string `json:"validator" form:"validator" example:"Validator1"`
	TglAwal        string `json:"tgl_awal" form:"tgl_awal" example:"2025-02-01"`
	TglAkhir       string `json:"tgl_akhir" form:"tgl_akhir" example:"2025-02-10"`
	Limit          int64  `json:"limit" form:"limit" example:"10"`
	Offset         int64  `json:"offset" form:"offset" example:"0"`
	Page           int64  `json:"page" form:"page" example:"1"`
}

type TaskResponses struct {
	No              int64    `json:"no"`
	ID              int64    `json:"id"`
	NoTasklist      *string  `json:"no_tasklist"`
	NamaTasklist    *string  `json:"nama_tasklist"`
	RiskIndicator   string   `json:"risk_indicator"`
	ActivityID      int64    `json:"activity_id"`
	ActivityName    string   `json:"activity_name"`
	ProductID       int64    `json:"product_id"`
	ProductName     string   `json:"product_name"`
	RiskIssueID     int64    `json:"risk_issue_id"`
	RiskIssue       string   `json:"risk_issue"`
	RiskIndicatorID int64    `json:"risk_indicator_id"`
	TaskType        int64    `json:"jenis_task"`
	TaskTypeName    string   `json:"task_type_name"`
	Kegiatan        string   `json:"kegiatan"`
	Period          string   `json:"period"`
	EnableRange     string   `json:"enable_range"`
	SumberData      string   `json:"sumber_data"`
	RangeDate       string   `json:"range_date"`
	StartDate       *string  `json:"start_date"`
	EndDate         *string  `json:"end_date"`
	RAP             bool     `json:"rap"`
	Validation      string   `json:"validation"`
	ValidationName  string   `json:"validation_name"`
	Approval        string   `json:"approval"`
	ApprovalName    string   `json:"approval_name"`
	MakerID         string   `json:"maker_id"`
	Status          string   `json:"status"`
	ApprovalStatus  string   `json:"approval_status"`
	StatusFile      string   `json:"status_file"`
	Branch          []string `json:"branch"`
}

type MyTasklistResponse struct {
	No              int64   `json:"no"`
	ID              int64   `json:"id"`
	BRANCH          int     `json:"BRANCH" `
	BRDESC          string  `json:"BRDESC" `
	NoTasklist      string  `json:"no_tasklist"`
	NamaTasklist    string  `json:"nama_tasklist"`
	Periode         string  `json:"periode"`
	RiskIndicator   string  `json:"risk_indicator"`
	ActivityID      int64   `json:"activity_id"`
	ActivityName    string  `json:"activity_name"`
	ProductID       int64   `json:"product_id"`
	ProductName     string  `json:"product_name"`
	RiskIssueID     int64   `json:"risk_issue_id"`
	RiskIssue       string  `json:"risk_issue"`
	RiskIndicatorID int64   `json:"risk_indicator_id"`
	TaskType        int64   `json:"jenis_task"`
	TaskTypeName    string  `json:"task_type_name"`
	Kegiatan        string  `json:"kegiatan"`
	Period          string  `json:"period"`
	RangeDate       string  `json:"range_date"`
	StartDate       *string `json:"start_date"`
	EndDate         *string `json:"end_date"`
	Progress        int64   `json:"progress"`
	JumlahNominatif int64   `json:"jumlah_nominatif"`
}

type TaskRejectionNotes struct {
	Notes string `json:"notes"`
}

type DataTematikRequest struct {
	Id            int64  `json:"id"`
	RiskEvent     int64  `json:"risk_event"`
	RiskIndicator int64  `json:"risk_indicator"`
	Region        string `json:"region"`
	Branch        string `json:"branch"`
	Limit         int64  `json:"limit"`
	Offset        int64  `json:"offset"`
}

type DataTematikResponse struct {
	Columns     string        `json:"columns"`
	ColumnsData []interface{} `json:"columns_data"`
}

func (t CreateTaskDTO) TableName() string {
	return "tasklists"
}

func (t TaskFile) TableName() string {
	return "tasklists_lampiran"
}

type ValidateLaporanRAPDTO struct {
	REGION string `json:"REGION" `
	RGDESC string `json:"RGDESC" `
	MAINBR int    `json:"MAINBR" `
	MBDESC string `json:"MBDESC" `
	BRANCH int    `json:"BRANCH" `
	BRDESC string `json:"BRDESC" `
	Count  int    `json:"count" `
}

func (t ValidateLaporanRAPDTO) TableName() string {
	return "dwh_branch"
}

type RequestApprovalEnum string

const (
	Approved RequestApprovalEnum = "approved"
	Rejected RequestApprovalEnum = "rejected"
)

type ApprovalRequest struct {
	ID       int64    `json:"id" example:"1"`
	Approval string   `json:"approval" example:"approve" enum:"approve,reject"` // This is of type ApprovalEnum
	Notes    string   `json:"notes" example:"Test Create something"`
	Branch   []string `json:"branch" example:"0"`
	Perner   string   `json:"pernr" example:"P0000001"`
}

type GetBRCResponse struct {
	PN              string `json:"pn"`
	SNAME           string `json:"sname"`
	REGION          string `json:"region"`
	RGDESC          string `json:"rgdesc"`
	MAINBR          string `json:"mainbr"`
	MBDESC          string `json:"mbdesc"`
	BRANCH          string `json:"branch"`
	BRDESC          string `json:"brdesc"`
	JumlahNominatif int64  `json:"jumlah_nominatif"`
}

type UpdateTasklistDTO struct {
	NoTasklist      string  `json:"no_tasklist" form:"no_tasklist" example:"Test Create something"`
	NamaTasklist    string  `json:"nama_tasklist" form:"nama_tasklist" example:"Test Create something"`
	RiskIndicator   string  `json:"risk_indicator" form:"risk_indicator" example:"Test Create something"`
	ActivityID      int64   `json:"activity_id" form:"activity_id" example:"1"`
	ProductID       int64   `json:"product_id" form:"product_id" example:"1"`
	ProductName     string  `json:"product_name" form:"product_name" example:"Test Create something"`
	RiskIssueID     int64   `json:"risk_issue_id" form:"risk_issue_id" example:"1"`
	RiskIssue       string  `json:"risk_issue" form:"risk_issue" example:"Test Create something"`
	RiskIndicatorID int64   `json:"risk_indicator_id" form:"risk_indicator_id" example:"1"`
	TaskType        int64   `json:"task_type" form:"task_type" example:"1"`
	TaskTypeName    string  `json:"task_type_name" form:"task_type_name" example:"Test Create something"`
	Kegiatan        string  `json:"kegiatan" form:"kegiatan" example:"Test Create something"`
	Period          string  `json:"period" form:"period" example:"Test Create something"`
	EnableRange     string  `json:"enable_range" form:"enable_range" example:"Quarter, Semester, Angka, Montly"`
	SumberData      string  `json:"sumber_data" form:"sumber_data" example:"Test Create something"`
	RangeDate       string  `json:"range_date" form:"range_date" example:"1"`
	StartDate       *string `json:"start_date" form:"start_date" example:"Test Create something"`
	EndDate         *string `json:"end_date" form:"end_date" example:"Test Create something"`
	RAP             bool    `json:"rap" form:"rap" example:"true"`
	Sample          int64   `json:"sample" form:"sample" example:"1"`
	Validation      string  `json:"validation" form:"validation" example:"Test Create something"`
	ValidationName  string  `json:"validation_name" form:"validation_name" example:"Test Create something"`
	Approval        string  `json:"approval" form:"approval" example:"Test Create something"`
	ApprovalName    string  `json:"approval_name" form:"approval_name" example:"Test Create something"`
	MakerID         string  `json:"maker_id" form:"maker_id" example:"Test Create something"`
}

type DeleteTasklistDTO struct {
	ID     int64  `json:"id" example:"1"`
	Perner string `json:"pernr" example:"P0000001"`
}
