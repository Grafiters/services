package models

import "time"

type Task struct {
	ID              int64        `json:"id" form:"id"`
	NoTasklist      string       `json:"no_tasklist" form:"no_tasklist" example:"Test Create something"`
	NamaTasklist    string       `json:"nama_tasklist" form:"nama_tasklist" example:"Test Create something"`
	RiskIndicator   string       `json:"risk_indicator" form:"risk_indicator" example:"Test Create something"`
	ActivityID      int64        `json:"activity_id" form:"activity_id" example:"1"`
	ProductID       int64        `json:"product_id" form:"product_id" example:"1"`
	ProductName     string       `json:"product_name" form:"product_name" example:"Test Create something"`
	RiskIssueID     int64        `json:"risk_issue_id" form:"risk_issue_id" example:"1"`
	RiskIssue       string       `json:"risk_issue" form:"risk_issue" example:"Test Create something"`
	RiskIndicatorID int64        `json:"risk_indicator_id" form:"risk_indicator_id" example:"1"`
	TaskType        int64        `json:"task_type" form:"task_type" example:"1"`
	TaskTypeName    string       `json:"task_type_name" form:"task_type_name" example:"Test Create something"`
	Kegiatan        string       `json:"kegiatan" form:"kegiatan" example:"Test Create something"`
	Period          string       `json:"period" form:"period" example:"Test Create something"`
	EnableRange     string       `json:"enable_range" form:"enable_range" example:"Quarter"`
	SumberData      string       `json:"sumber_data" form:"sumber_data" example:"Test Create something"`
	RangeDate       string       `json:"range_date" form:"range_date" example:"1"`
	StartDate       *string      `json:"start_date" form:"start_date" example:"Test Create something"`
	EndDate         *string      `json:"end_date" form:"end_date" example:"Test Create something"`
	RAP             bool         `json:"rap" form:"rap" example:"true"`
	Sample          int64        `json:"sample" form:"sample" example:"1"`
	Validation      string       `json:"validation" form:"validation" example:"Test Create something"`
	ValidationName  string       `json:"validation_name" form:"validation_name" example:"Test Create something"`
	Approval        string       `json:"approval" form:"approval" example:"Test Create something"`
	ApprovalName    string       `json:"approval_name" form:"approval_name" example:"Test Create something"`
	ApprovalStatus  ApprovalEnum `json:"approval_status" form:"approval_status" example:"Test Create something"`
	Status          string       `json:"status" form:"status" example:"Test Create something"`
	MakerID         string       `json:"maker_id" form:"maker_id" example:"Test Create something"`
	StatusFile      string       `json:"status_file" form:"status_file" example:"Sedang diproses"`
	CreatedAt       string       `json:"created_at" form:"created_at"`
	UpdatedAt       string       `json:"updated_at" form:"updated_at"`
}

type ApprovalEnum string

// ApprovalEnum constants
const (
	M1 ApprovalEnum = "Minta Persetujuan Validasi"
	C1 ApprovalEnum = "Minta Persetujuan Approval"
	C0 ApprovalEnum = "Ditolak Checker"
	S1 ApprovalEnum = "Disetujui"
	S0 ApprovalEnum = "Ditolak Signer"
)

func (Task) TableName() string {
	return "tasklists"
}

type TasklistUker struct {
	ID              int64  `json:"id" form:"id"`
	TasklistId      int64  `json:"tasklist_id"`
	REGION          string `json:"REGION"`
	RGDESC          string `json:"RGDESC"`
	MAINBR          string `json:"MAINBR"`
	MBDESC          string `json:"MBDESC"`
	BRANCH          string `json:"BRANCH"`
	BRDESC          string `json:"BRDESC"`
	JumlahNominatif int    `json:"jumlah_nominatif"`
}

func (TasklistUker) TableName() string {
	return "tasklists_uker"
}

type DWHBranch struct {
	REGION string `json:"REGION"`
	RGDESC string `json:"RGDESC"`
	MAINBR int    `json:"MAINBR"`
	MBDESC string `json:"MBDESC"`
	BRANCH string `json:"BRANCH"`
	BRDESC string `json:"BRDESC"`
}

func (DWHBranch) TableName() string {
	return "dwh_branch"
}

type TasklistsAprroval struct {
	ID             int64   `json:"id" example:"123"`
	ApprovalStatus string  `json:"approval" example:"approved"`
	UpdatedAt      *string `json:"updated_at" example:"2023-10-01T00:00:00Z"`
}

type TasklistsRejected struct {
	ID         int64
	TasklistId int64
	Notes      string
	Status     string
}

type AdminSetting struct {
	ID          int64     `json:"id"`           // Primary key
	TaskType    string    `json:"task_type"`    // Varchar(100)
	Period      string    `json:"period"`       // Varchar(100)
	Kegiatan    string    `json:"kegiatan"`     // Varchar(100)
	Range       string    `json:"range"`        // Varchar(100), nullable
	Upload      string    `json:"upload"`       // Varchar(100), nullable
	TasklistMax uint      `json:"tasklist_max"` // Unsigned int
	Status      string    `json:"status"`       // Varchar(100)
	CreatedAt   time.Time `json:"created_at"`   // Datetime
	UpdatedAt   time.Time `json:"updated_at"`   // Datetime
}

func (AdminSetting) TableName() string {
	return "admin_setting"
}

type TasklistsToday struct {
	ID              uint `gorm:"primaryKey"`
	TasklistID      int64
	ActivityID      int64
	ProductID       int64
	Product         string
	RiskIssueID     int64
	RiskIssue       string
	RiskIndicatorID int64
	RiskIndicator   string
	StartDate       *string
	EndDate         *string
	Status          string
	TaskType        int64
	TaskTypeName    string
	Kegiatan        string
	Period          string
	Sample          int64
	Progres         int64
	Persentase      int64
	MakerID         string
	PERNR           string
	Assigned        string
	REGION          string
	RGDESC          string
	MAINBR          string
	MBDESC          string
	BRANCH          string
	BRDESC          string
	CreatedAt       *string
	UpdatedAt       *string
}

func (TasklistsToday) TableName() string {
	return "tasklists_today"
}

type TasklistNotif struct {
	TaskID     int64   `json:"task_id"`
	Tanggal    *string `json:"tanggal"`
	Keterangan string  `json:"keterangan"`
	Status     int64   `json:"status"`
	Jenis      string  `json:"jenis"`
	Receiver   string  `json:"receiver"`
	Uker       string  `json:"uker"`
	// Branch     string  `json:"branch"`
}

func (TasklistNotif) TableName() string {
	return "tasklist_notifikasis"
}

type TasklistRejected struct {
	ID         int64  `json:"id" db:"id"`
	TasklistID int64  `json:"tasklist_id" db:"tasklist_id"`
	Notes      string `json:"notes" db:"notes"`
	Status     string `json:"status" db:"status"`
}

func (TasklistRejected) TableName() string {
	return "tasklists_rejected"
}

type LampiranIndicatorRequest struct {
	RiskIssueId       int64  `json:"risk_issue_id"`
	RiskIndicatorId   int64  `json:"risk_indicator_id"`
	NamaTable         string `json:"nama_table"`
	JumlahKolom       int64  `json:"jumlah_kolom"`
	RiskIndicatorDesc string `json:"risk_indicator_desc"`
}

func (LampiranIndicatorRequest) TableName() string {
	return "lampiran_indikator"
}

type RequestMyTasklist struct {
	Branch string `json:"branch"`
	Pernr  string `json:"pernr"`
}

type MyTasklist struct {
	Total int64 `json:"total"`
}
