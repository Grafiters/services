package models

type Tasklists struct {
	ID              int64
	NoTasklist      string
	NamaTasklist    string
	ActivityID      int64
	ProductID       int64
	ProductName     string
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
	SumberData      string
	RAP             string
	Sample          int64
	Validation      string
	ValidationName  string
	Approval        string
	ApprovalName    string
	IsiLampiran     string
	ApprovalStatus  string
	MakerID         string
	CreatedAt       *string
	UpdatedAt       *string
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

type TasklistsUpdate struct {
	ID                 int64
	NoTasklist         string
	NamaTasklist       string
	ActivityID         int64
	ProductID          int64
	ProductName        string
	RiskIssueID        int64
	RiskIssue          string
	RiskIndicatorID    int64
	RiskIndicator      string
	StartDate          *string
	EndDate            *string
	Status             string
	TaskType           int64
	TaskTypeName       string
	Kegiatan           string
	Period             string
	SumberData         string
	RAP                string
	Sample             string
	Validation         string
	ValidationName     string
	Approval           string
	ApprovalName       string
	IsiLampiran        string
	ApprovalStatus     string
	VerificationStatus string
	MakerID            string
	CreatedAt          *string
	UpdatedAt          *string
}

type TasklistsUpdateEndDate struct {
	ID      int64
	EndDate *string
}

type TasklistsUpdateDelete struct {
	ID        int64   `json:"id" example:"123"`
	Status    string  `json:"status" example:"active"`
	UpdatedAt *string `json:"updated_at" example:"2023-10-01T00:00:00Z"`
}

type TasklistsAprroval struct {
	ID             int64   `json:"id" example:"123"`
	ApprovalStatus string  `json:"approval" example:"approved"`
	UpdatedAt      *string `json:"updated_at" example:"2023-10-01T00:00:00Z"`
}

type TasklistsAprrovalRequest struct {
	ID             int64   `json:"id" example:"123"`
	Notes          string  `json:"notes" example:"notes"`
	ApprovalStatus string  `json:"approval" example:"approve"`
	Receiver       string  `json:"receiver" example:"receiver"`
	Branch         string  `json:"branch" example:"branch"`
	Uker           string  `json:"uker" example:"uker"`
	UpdatedAt      *string `json:"updated_at" example:"2023-10-01T00:00:00Z"`
}

type TasklistRejectedNote struct {
	Notes string `json:"notes" example:"notes"`
}

type TasklistsValidation struct {
	Notes            string  `json:"notes" example:"notes"`
	ID               int64   `json:"id"`
	Receiver         string  `json:"receiver" example:"receiver"`
	ValidationStatus string  `json:"validation" example:"rejected"`
	Uker             string  `json:"uker" example:"uker"`
	Branch           string  `json:"branch" example:"branch"`
	UpdatedAt        *string `json:"updated_at" example:"2023-10-01T00:00:00Z"`
}

type TasklistsDoneHistoryRequest struct {
	TasklistID int64  `json:"tasklist_id" example:"123"`
	PERNR      string `json:"pernr" example:"456"`
}

type TasklistsDoneHistoryCheckRequest struct {
	TasklistID int64  `json:"tasklist_id" example:"123"`
	PERNR      string `json:"pernr" example:"456"`
	Date       string `json:"date" example:"2023-10-01"`
}

type TasklistsDoneHistory struct {
	ID         int64
	PERNR      string
	TasklistID int64
	CreatedAt  *string
}

type TasklistsDoneHistoryResponse struct {
	PERNR      string  `json:"pernr"`
	TasklistID int64   `json:"tasklist_id"`
	CreatedAt  *string `json:"created_at"`
}

func (TasklistsUpdate) TableName() string {
	return "tasklists"
}

func (TasklistsUpdateDelete) TableName() string {
	return "tasklists"
}

func (TasklistsUpdateEndDate) TableName() string {
	return "tasklists"
}

func (TasklistsAprroval) TableName() string {
	return "tasklists"
}

func (TasklistsValidation) TableName() string {
	return "tasklists"
}

func (TasklistsDoneHistory) TableName() string {
	return "tasklists_done_history"
}

func (TasklistNotif) TableName() string {
	return "tasklist_notifikasis"
}

func (TasklistsToday) TableName() string {
	return "tasklists_today"
}

type AnomaliHeader struct {
	TasklistID    string `json:"tasklist_id"`
	RiskIssue     string `json:"risk_issue"`
	RiskIndicator string `json:"risk_indicator"`
}

type AnomaliValue struct {
	RiskIssue     string `json:"risk_issue" example:"Risk Issue Example"`
	RiskIndicator string `json:"risk_indicator" example:"Risk Indicator Example"`
	TasklistID    string `json:"tasklist_id" example:"1"`
}

type AnomaliHeaderResponse struct {
	Header []string `json:"header"`
}

type AnomaliValueResponse struct {
	Value string `json:"value"`
}

type GetFirstLampiranRequest struct {
	// TasklistID    string `json:"tasklist_id"`
	RiskIssue     string `json:"risk_issue" example:"Risk Issue Example"`
	RiskIndicator string `json:"risk_indicator" example:"Risk Indicator Example"`
}

type GetFirstLampiranResponse struct {
	Path     string `json:"path"`
	Filename string `json:"filename"`
}

type TasklistLampiranDelete struct {
	TasklistID    int64 `json:"tasklist_id"`
	RiskIssue     int64 `json:"risk_issue"`
	RiskIndicator int64 `json:"risk_indicator"`
}

type PegawaiData struct {
	Stell_TX string `json:"stell_tx"`
	HILFM    string `json:"hilfm"`
	TipeUker string `json:"tipe_uker"`
}
