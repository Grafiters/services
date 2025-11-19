package models

type TasklistsFilterRequest struct {
	ID              int    `json:"id" example:"1"`
	Order           string `json:"order" example:"DESC"`
	Sort            string `json:"sort" example:"id"`
	Offset          int    `json:"offset" example:"0"`
	Limit           int    `json:"limit" example:"10"`
	Page            int    `json:"page" example:"1"`
	PERNR           string `json:"pernr" example:"100001"`
	REGION          string `json:"region" example:"JAKARTA"`
	RGDESC          string `json:"rgdesc" example:"JAKARTA"`
	MAINBR          string `json:"mainbr" example:"JAKARTA"`
	MBDESC          string `json:"mbdesc" example:"JAKARTA"`
	BRANCH          string `json:"branch" example:"JAKARTA"`
	BRDESC          string `json:"brdesc" example:"JAKARTA"`
	ActivityID      int64  `json:"activity_id" example:"1"`
	ProductID       int64  `json:"product_id" example:"1"`
	RiskIssueID     int64  `json:"risk_issue_id" example:"1"`
	RiskIndicatorID int64  `json:"risk_indikator_id" example:"1"`
	RiskIndicator   string `json:"risk_indicator" example:"TEST"`
	Approval        string `json:"status_approval" example:"APPROVED"`
	JenisTask       string `json:"jenis_task" example:"TEST"`
	StartDate       string `json:"start_date" example:"2021-01-01"`
	EndData         string `json:"end_date" example:"2021-01-31"`
	TipeUker        string `json:"tipe_uker" example:"TEST"`
	Status          string `json:"status" example:"PENDING"`
}

type TasklistsFilterByIDRequest struct {
	ID int64 `json:"id" example:"1"`
}

type TasklistsFilterResponse struct {
	ID int64 `json:"id"`
	// RGDESC         string `json:"kanwil"`
	// MBDESC         string `json:"kanca"`
	// BRDESC         string `json:"uker"`
	// BRANCH         string `json:"branch"`
	StartDate      string              `json:"start_date"`
	EndDate        string              `json:"end_date"`
	RiskIssue      string              `json:"risk_issue"`
	Activity       string              `json:"aktivitas"`
	Product        string              `json:"product"`
	RiskIndicator  string              `json:"risk_indicator"`
	StatusApproval string              `json:"status_approval"`
	MakerID        string              `json:"maker_id"`
	JenisTask      string              `json:"jenis_task"`
	Period         string              `json:"period"`
	Validation     string              `json:"validation"`
	Approval       string              `json:"approval"`
	IDLampiran     int64               `json:"id_lampiran"`
	TasklistID     int64               `json:"tasklist_id"`
	Filename       string              `json:"filename"`
	Path           string              `json:"path"`
	Ext            string              `json:"ext"`
	Size           int64               `json:"size"`
	Status         string              `json:"status"`
	Uker           []TasklistsUkerData `json:"uker"`
	DataPegawai    PegawaiData         `json:"data_pegawai"`
}

type TasklistsFilterOfficerRequest struct {
	Order           string `json:"order" example:"DESC"`
	Sort            string `json:"sort" example:"id"`
	Offset          int    `json:"offset" example:"0"`
	Limit           int    `json:"limit" example:"10"`
	Page            int    `json:"page" example:"1"`
	PERNR           string `json:"pernr" example:"100001"`
	REGION          string `json:"region" example:"JAKARTA"`
	RGDESC          string `json:"rgdesc" example:"JAKARTA"`
	MAINBR          string `json:"mainbr" example:"JAKARTA"`
	MBDESC          string `json:"mbdesc" example:"JAKARTA"`
	BRANCH          string `json:"branch" example:"JAKARTA"`
	BRDESC          string `json:"brdesc" example:"JAKARTA"`
	ActivityID      int64  `json:"activity_id" example:"1"`
	ProductID       int64  `json:"product_id" example:"1"`
	RiskIssueID     int64  `json:"risk_issue_id" example:"1"`
	RiskIndicatorID int64  `json:"risk_indicator_id" example:"1"`
	RiskIndicator   string `json:"risk_indicator" example:"TEST"`
	Approval        string `json:"status_approval" example:"APPROVED"`
	JenisTask       string `json:"jenis_task" example:"TEST"`
	StartDate       string `json:"start_date" example:"2021-01-01"`
	EndData         string `json:"end_date" example:"2021-01-31"`
	Status          string `json:"status" example:"PENDING"`
	MakerID         string `json:"maker_id" example:"100001"`
	TipeUker        string `json:"tipe_uker" example:"TEST"`
	HILFM           string `json:"hilfm" example:"TEST"`
	StellTX         string `json:"stell_tx" example:"TEST"`
}

type TasklistsFilterOfficerResponse struct {
	ID int64 `json:"id"`
	// RGDESC         string `json:"kanwil"`
	// MBDESC         string `json:"kanca"`
	// BRDESC         string `json:"uker"`
	StartDate      string              `json:"start_date"`
	EndDate        string              `json:"end_date"`
	RiskIssue      string              `json:"risk_issue"`
	Activity       string              `json:"aktivitas"`
	Product        string              `json:"product"`
	RiskIndicator  string              `json:"risk_indicator"`
	StatusApproval string              `json:"status_approval"`
	MakerID        string              `json:"maker_id"`
	JenisTask      string              `json:"jenis_task"`
	Period         string              `json:"period"`
	Validation     string              `json:"validation"`
	Approval       string              `json:"approval"`
	Status         string              `json:"status"`
	Uker           []TasklistsUkerData `json:"uker"`
	DataPegawai    PegawaiData         `json:"data_pegawai"`
}

type TasklistsFilterResponses struct {
	ID int64 `json:"id"`
	// RGDESC         string `json:"kanwil"`
	// MBDESC         string `json:"kanca"`
	// BRDESC         string `json:"uker"`
	StartDate      string `json:"start_date"`
	EndDate        string `json:"end_date"`
	RiskIssue      string `json:"risk_issue"`
	Activity       string `json:"aktivitas"`
	Product        string `json:"product"`
	RiskIndicator  string `json:"risk_indicator"`
	StatusApproval string `json:"status_approval"`
	MakerID        string `json:"maker_id"`
	JenisTask      string `json:"jenis_task"`
	Period         string `json:"period"`
	Validation     string `json:"validation"`
	Approval       string `json:"approval"`
	Status         string `json:"status"`
}

type TasklistsFilterApprovalResponses struct {
	ID int64 `json:"id"`
	// RGDESC         string `json:"kanwil"`
	// MBDESC         string `json:"kanca"`
	// BRDESC         string `json:"uker"`
	// BRANCH         string `json:"branch"`
	StartDate      string `json:"start_date"`
	EndDate        string `json:"end_date"`
	RiskIssue      string `json:"risk_issue"`
	Activity       string `json:"aktivitas"`
	Product        string `json:"product"`
	RiskIndicator  string `json:"risk_indicator"`
	StatusApproval string `json:"status_approval"`
	MakerID        string `json:"maker_id"`
	JenisTask      string `json:"jenis_task"`
	Period         string `json:"period"`
	Validation     string `json:"validation"`
	Approval       string `json:"approval"`
	IDLampiran     int64  `json:"id_lampiran"`
	TasklistID     int64  `json:"tasklist_id"`
	Filename       string `json:"filename"`
	Path           string `json:"path"`
	Ext            string `json:"ext"`
	Size           int64  `json:"size"`
	Status         string `json:"status"`
}
