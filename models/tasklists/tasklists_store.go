package models

import "riskmanagement/models/files"

type TasklistsStoreRequest struct {
	Uker            []TasklistsUker      `json:"uker"`
	NoTasklist      string               `json:"no_tasklist"`
	NamaTasklist    string               `json:"nama_tasklist"`
	RiskIndicator   string               `json:"risk_indicator"`
	ActivityID      int64                `json:"activity_id"`
	ProductID       int64                `json:"product_id"`
	ProductName     string               `json:"product_name"`
	RiskIssueID     int64                `json:"risk_issue_id"`
	RiskIssue       string               `json:"risk_issue"`
	RiskIndicatorID int64                `json:"risk_indicator_id"`
	TaskType        int64                `json:"jenis_task"`
	TaskTypeName    string               `json:"task_type_name"`
	Kegiatan        string               `json:"kegiatan"`
	Period          string               `json:"period"`
	SumberData      string               `json:"sumber_data"`
	StartDate       *string              `json:"start_date"`
	EndDate         *string              `json:"end_date"`
	Files           []files.FilesRequest `json:"files"`
	RAP             string               `json:"rap"`
	Sample          int64                `json:"sample"`
	HeaderLampiran  string               `json:"header_lampiran"`
	IsiLampiran     []interface{}        `json:"isi_lampiran"`
	JumlahKolom     int64                `json:"jumlah_kolom"`
	Validation      string               `json:"validation"`
	ValidationName  string               `json:"validation_name"`
	Approval        string               `json:"approval"`
	ApprovalName    string               `json:"approval_name"`
	MakerID         string               `json:"maker_id"`
}

type LampiranIndikatorStore struct {
	RiskIssueID       int64  `json:"risk_issue_id"`
	RiskIndicatorID   int64  `json:"risk_indicator_id"`
	NamaTable         string `json:"nama_table"`
	JumlahKolom       int64  `json:"jumlah_kolom"`
	RiskIndicatorDesc string `json:"risk_indicator_desc"`
}

type LampiranIndikatorCheck struct {
	RiskIssueID     string `json:"risk_issue_id" example:"1"`
	RiskIndicatorID string `json:"risk_indicator_id" example:"1"`
}

type TasklistHeaderStore struct {
	TasklistID      int64  `json:"tasklist_id"`
	RiskIssueID     string `json:"risk_issue_id"`
	RiskIndicatorID string `json:"risk_indicator_id"`
	HeaderLampiran  string
}

type TasklistColumnStore struct {
	TasklistID      string `json:"tasklist_id" example:"123"`
	RiskIssueID     string `json:"risk_issue_id" example:"456"`
	RiskIndicatorID string `json:"risk_indicator_id" example:"789"`
	HeaderLampiran  string `json:"header_lampiran" example:"Header Example"`
	IsiLampiran     string `json:"isi_lampiran" example:"Isi Example"`
}

type TasklistsUpdateRequest struct {
	ID              int64                `json:"id" example:"123"`
	Uker            []TasklistsUker      `json:"uker"`
	NoTasklist      string               `json:"no_tasklist" example:"TL-001"`
	NamaTasklist    string               `json:"nama_tasklist" example:"Task List 1"`
	RiskIndicator   string               `json:"risk_indicator" example:"Low"`
	ActivityID      int64                `json:"activity_id" example:"456"`
	ProductID       int64                `json:"product_id" example:"789"`
	ProductName     string               `json:"product_name" example:"Product A"`
	RiskIssueID     int64                `json:"risk_issue_id" example:"112"`
	RiskIssue       string               `json:"risk_issue" example:"Issue"`
	RiskIndicatorID int64                `json:"risk_indicator_id" example:"223"`
	TaskType        int64                `json:"jenis_task" example:"1"`
	TaskTypeName    string               `json:"task_type_name" example:"Type A"`
	Kegiatan        string               `json:"kegiatan" example:"Activity"`
	Period          string               `json:"period" example:"2023-10"`
	SumberData      string               `json:"sumber_data" example:"External"`
	StartDate       *string              `json:"start_date" example:"2023-10-01"`
	EndDate         *string              `json:"end_date" example:"2023-12-31"`
	Sample          string               `json:"sample" example:"Sample Data"`
	Files           []files.FilesRequest `json:"files"`
	RAP             string               `json:"rap" example:"RAP Data"`
	HeaderLampiran  string               `json:"header_lampiran" example:"Header"`
	IsiLampiran     []interface{}        `json:"isi_lampiran"`
	JumlahKolom     int64                `json:"jumlah_kolom" example:"5"`
	Validation      string               `json:"validation" example:"Validated"`
	ValidationName  string               `json:"validation_name" example:"Validation Name"`
	Approval        string               `json:"approval" example:"Approved"`
	ApprovalName    string               `json:"approval_name" example:"Approval Name"`
	MakerID         string               `json:"maker_id" example:"Maker123"`
	CreatedAt       *string              `json:"created_at" example:"2023-10-01T00:00:00Z"`
}

type TasklistsUpdateEndDateRequest struct {
	ID        int64   `json:"id" example:"123"`
	EndDate   *string `json:"end_date" example:"2023-12-31"`
	UpdatedAt *string `json:"updated_at" example:"2023-10-01T00:00:00Z"`
}
