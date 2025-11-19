package models

type TasklistDaily struct {
	// TasklistID      int64
	ActivityID      int64
	ProductID       int64
	ProductName     string
	RiskIssueID     int64
	RiskIssue       string
	RiskIndicatorID int64
	RiskIndicator   string
	StartDate       *string
	EndDate         *string
	TaskType        int64
	TaskTypeName    string
	SumberData      string
	RAP             int64
	Validation      string
	ValidationName  string
	Approval        string
	ApprovalName    string
	Sample          string
	REGION          string
	RGDESC          string
	MAINBR          string
	MBDESC          string
	BRANCH          string
	BRDESC          string
	PERNR           string
	ApprovalStatus  string
	Status          string
	MakerID         string
}

type TasklistCheckRequest struct {
	RiskIssueID int64  `json:"risk_issue_id"`
	ActivityID  int64  `json:"activity_id"`
	ProductID   int64  `json:"product_id"`
	BRANCH      string `json:"branch"`
}

type TasklistCheckResponse struct {
	ID              int64   `json:"id"`
	ActivityID      int64   `json:"activity_id"`
	ProductID       int64   `json:"product_id"`
	ProductName     string  `json:"product_name"`
	RiskIssueID     int64   `json:"risk_issue_id"`
	RiskIssue       string  `json:"risk_issue"`
	RiskIndicatorID int64   `json:"risk_indicator_id"`
	RiskIndicator   string  `json:"risk_indicator"`
	StartDate       *string `json:"start_date"`
	EndDate         *string `json:"end_date"`
	TaskType        int64   `json:"task_type"`
	TaskTypeName    string  `json:"task_type_name"`
	SumberData      string  `json:"sumber_data"`
	RAP             int64   `json:"rap"`
	Validation      string  `json:"validation"`
	ValidationName  string  `json:"validation_name"`
	Approval        string  `json:"approval"`
	ApprovalName    string  `json:"approval_name"`
	Sample          string  `json:"sample"`
	REGION          string  `json:"region"`
	RGDESC          string  `json:"rgdesc"`
	MAINBR          string  `json:"mainbr"`
	MBDESC          string  `json:"mbdesc"`
	BRANCH          string  `json:"branch"`
	BRDESC          string  `json:"brdesc"`
	PERNR           string  `json:"pernr"`
	ApprovalStatus  string  `json:"approval_status"`
	Status          string  `json:"status"`
	MakerID         string  `json:"maker_id"`
}

type TasklistDailyStore struct {
	// TasklistID      int64   `json:"tasklist_id"`
	ActivityID      int64   `json:"activity_id" example:"1"`
	ProductID       int64   `json:"product_id" example:"1"`
	ProductName     string  `json:"product_name" example:"Kredit"`
	RiskIssueID     int64   `json:"risk_issue_id" example:"1"`
	RiskIssue       string  `json:"risk_issue" example:"Kredit Macet"`
	RiskIndicatorID int64   `json:"risk_indicator_id" example:"1"`
	RiskIndicator   string  `json:"risk_indicator" example:"Kredit Macet > 3 Bulan"`
	StartDate       *string `json:"start_date" example:"2022-01-01"`
	EndDate         *string `json:"end_date" example:"2022-01-31"`
	TaskType        int64   `json:"task_type" example:"1"`
	TaskTypeName    string  `json:"task_type_name" example:"Daily"`
	SumberData      string  `json:"sumber_data" example:"Kanal"`
	RAP             int64   `json:"rap" example:"1"`
	Validation      string  `json:"validation" example:"Manual"`
	ValidationName  string  `json:"validation_name" example:"Manual"`
	Approval        string  `json:"approval" example:"Manual"`
	ApprovalName    string  `json:"approval_name" example:"Manual"`
	Progres         int64   `json:"progres" example:"1"`
	Sample          int64   `json:"sample" example:"1"`
	REGION          string  `json:"region" example:"Jawa"`
	RGDESC          string  `json:"rgdesc" example:"Jawa"`
	MAINBR          string  `json:"mainbr" example:"KANWIL JAWA"`
	MBDESC          string  `json:"mbdesc" example:"KANWIL JAWA"`
	BRANCH          string  `json:"branch" example:"KC JAKARTA"`
	BRDESC          string  `json:"brdesc" example:"KC JAKARTA"`
	PERNR           string  `json:"pernr" example:"123456"`
	ApprovalStatus  string  `json:"approval_status" example:"Waiting"`
	Status          string  `json:"status" example:"Open"`
	MakerID         string  `json:"maker_id" example:"123456"`
	CreatedAt       *string
	UpdatedAt       *string
}

type ProgresUpdateRequest struct {
	ID      int64 `json:"id" example:"1"`
	Progres int64 `json:"progres" example:"50"`
}

func (TasklistDaily) TableName() string {
	return "tasklists_daily"
}
