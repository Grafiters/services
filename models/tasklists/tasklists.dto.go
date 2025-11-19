package models

type TasklistsRequest struct {
	Period string `json:"period"`
}

type TasklistsCheckRequest struct {
	TaskType  string `json:"task_type" example:"1"`
	REGION    string `json:"region" example:"Banten"`
	RGDESC    string `json:"rgdesc" example:"Banten"`
	RGNAME    string `json:"rgname" example:"Banten"`
	MAINBR    string `json:"mainbr" example:"001"`
	MBDESC    string `json:"mbdesc" example:"KCP BSD"`
	MBNAME    string `json:"mbname" example:"KCP BSD"`
	BRANCH    string `json:"branch" example:"001"`
	BRDESC    string `json:"brdesc" example:"KCP BSD"`
	BRNAME    string `json:"brname" example:"KCP BSD"`
	UnitKerja string `json:"unit_kerja" example:"KCP BSD"`
	RiskIssue int64  `json:"risk_issue_id" example:"1"`
}

type TasklistRiskIssueAvailReq struct {
	RiskIssueID int64  `json:"risk_issue_id"`
	Branch      string `json:"branch"`
	ActivityID  int64  `json:"activity_id"`
	ProductID   int64  `json:"product_id"`
}

type TasklistRiskIssueAvailRes struct {
	Total int64 `json:"total"`
}

type TasklistsResponse struct {
	TasklistID      int64  `json:"id"`
	Branch          string `json:"branch"`
	BRDESC          string `json:"brdesc"`
	ActivityID      int64  `json:"activity_id"`
	Activity        string `json:"activity"`
	ProductID       int64  `json:"product_id"`
	Product         string `json:"product"`
	RiskIssueID     int64  `json:"risk_issue_id"`
	RiskIssueCode   string `json:"risk_issue_code"`
	RiskIssue       string `json:"task"`
	RiskIndicatorID string `json:"risk_indicator_id"`
	RiskIndicator   string `json:"sub_task"`
	Kegiatan        string `json:"kegiatan"`
	// SubTask    []TasklistSubTask `json:"sub_task"`
	StartDate  *string `json:"start_date"`
	EndDate    *string `json:"end_date"`
	DoneSample int64   `json:"done_sample"`
	Sample     int64   `json:"sample"`
	Status     string  `json:"status"`
	JenisTask  string  `json:"jenis_task"`
}

type TasklistDataResponse struct {
	ID              int64   `json:"id"`
	NoTasklist      string  `json:"no_tasklist"`
	NamaTasklist    string  `json:"nama_tasklist"`
	SumberData      string  `json:"sumber_data"`
	RAP             string  `json:"rap"`
	ActivityID      int64   `json:"activity_id"`
	Activity        string  `json:"activity"`
	ProductID       int64   `json:"product_id"`
	ProductName     string  `json:"product_name"`
	RiskIssueID     int64   `json:"task"`
	RiskIssue       string  `json:"risk_issue"`
	RiskIndicatorID int64   `json:"sub_task"`
	RiskIndicator   string  `json:"risk_indicator"`
	StartDate       *string `json:"start_date"`
	EndDate         *string `json:"end_date"`
	Status          string  `json:"status"`
	IsiLampiran     string  `json:"isi_lampiran"`
	TaskType        int64   `json:"task_type"`
	TaskTypeName    string  `json:"task_type_name"`
	TaskTypePeriod  string  `json:"task_type_period"`
	Validation      string  `json:"validation"`
	ValidationName  string  `json:"validation_name"`
	Approval        string  `json:"approval"`
	ApprovalName    string  `json:"approval_name"`
	ApprovalStatus  string  `json:"approval_status"`
	Range           string  `json:"range"`
	Upload          string  `json:"upload"`
	Kegiatan        string  `json:"kegiatan"`
	Period          string  `json:"period"`
	Sample          int64   `json:"sample"`
	MakerID         string  `json:"maker_id"`
	CreatedAt       *string `json:"created_at"`
}

type TasklistResponse struct {
	ID              int64                    `json:"id"`
	NoTasklist      string                   `json:"no_tasklist"`
	NamaTasklist    string                   `json:"nama_tasklist"`
	Uker            []TasklistsUker          `json:"uker"`
	SumberData      string                   `json:"sumber_data"`
	RAP             string                   `json:"rap"`
	ActivityID      int64                    `json:"activity_id"`
	Activity        string                   `json:"activity"`
	ProductID       int64                    `json:"product_id"`
	ProductName     string                   `json:"product_name"`
	RiskIssueID     int64                    `json:"task"`
	RiskIssue       string                   `json:"risk_issue"`
	RiskIndicatorID int64                    `json:"sub_task"`
	RiskIndicator   string                   `json:"risk_indicator"`
	StartDate       *string                  `json:"start_date"`
	EndDate         *string                  `json:"end_date"`
	Status          string                   `json:"status"`
	IsiLampiran     string                   `json:"isi_lampiran"`
	TaskType        int64                    `json:"task_type"`
	TaskTypeName    string                   `json:"task_type_name"`
	TaskTypePeriod  string                   `json:"task_type_period"`
	Validation      string                   `json:"validation"`
	ValidationName  string                   `json:"validation_name"`
	Approval        string                   `json:"approval"`
	ApprovalName    string                   `json:"approval_name"`
	ApprovalStatus  string                   `json:"approval_status"`
	Sample          int64                    `json:"sample"`
	Notes           TasklistRejectedNote     `json:"notes"`
	Range           string                   `json:"range"`
	Upload          string                   `json:"upload"`
	Kegiatan        string                   `json:"kegiatan"`
	Period          string                   `json:"period"`
	Files           []TasklistFilesResponses `json:"files"`
	Progres         int64                    `json:"progres"`
	Persentase      int64                    `json:"persentase"`
	MakerID         string                   `json:"maker_id"`
	CreatedAt       *string                  `json:"created_at"`
}

type LampiranIndikatorResponse struct {
	ID              int64  `json:"id"`
	RiskIssueID     int64  `json:"risk_issue_id"`
	RiskIndicatorID int64  `json:"risk_indicator_id"`
	NamaTable       string `json:"nama_table"`
	JumlahKolom     string `json:"jumlah_kolom"`
}

type TasklistSubTask struct {
	RiskIndicator string `json:"risk_indicator"`
}

type TasklistsCountResponse struct {
	Total int64 `json:"total"`
}

type TasklistsCheckResponse struct {
	Total int64 `json:"total"`
}

type TasklistCountRequest struct {
	BRANCH int64  `json:"branch" example:"123"`
	PERNR  string `json:"pernr" example:"456"`
}

type UkerListReq struct {
	ID       int64  `json:"id"`
	PERNR    string `json:"pernr"`
	TipeUker string `json:"tipe_uker"`
	HILFM    string `json:"hilfm"`
	StellTX  string `json:"stell_tx"`
	REGION   string `json:"region"`
}

type Paginate struct {
	ID        int    `json:"id" example:"1"`
	Order     string `json:"order" example:"asc"`
	Sort      string `json:"sort" example:"id"`
	Offset    int    `json:"offset" example:"0"`
	Limit     int    `json:"limit" example:"10"`
	Page      int    `json:"page" example:"1"`
	Period    string `json:"period" example:"2022"`
	REGION    string `json:"region" example:"DKI JAKARTA"`
	MAINBR    string `json:"mainbr" example:"KCP PONDOK INDAH"`
	BRANCH    string `json:"branch" example:"KCP PONDOK INDAH"`
	UnitKerja string `json:"unit_kerja" example:"KCP PONDOK INDAH"`
	PERNR     string `json:"pernr" example:"123456"`
	TipeUker  string `json:"tipe_uker" example:"KCP"`
	HILFM     string `json:"hilfm" example:"KCP001"`
	StellTX   string `json:"stell_tx" example:"KCP PONDOK INDAH"`
}

type GetTaskByID struct {
	ID int64 `json:"id" example:"1"`
}

type CountDoneSampleRequest struct {
	Kegiatan      string `json:"kegiatan"`
	ActivityID    int64  `json:"activity_id"`
	ProductID     int64  `json:"product_id"`
	RiskIssueID   int64  `json:"risk_issue_id"`
	RiskIssueCode string `json:"risk_issue_code"`
	Branch        string `json:"branch"`
}

type CountDoneSampleResponse struct {
	Total int64 `json:"total"`
}

type PaginateTaskToday struct {
	ID        int    `json:"id"`
	Order     string `json:"order"`
	Sort      string `json:"sort"`
	Offset    int    `json:"offset"`
	Limit     int    `json:"limit"`
	Page      int    `json:"page"`
	Period    string `json:"period"`
	REGION    string `json:"region"`
	MAINBR    int64  `json:"mainbr"`
	BRANCH    string `json:"branch"`
	UnitKerja string `json:"unit_kerja"`
	PERNR     string `json:"pernr"`
	TipeUker  string `json:"tipe_uker"`
}

type GetBRCRequest struct {
	// Branch int64 `json:"branch"`
	Uker string `json:"uker"`
}

type GetBRCResponse struct {
	PERNR  string `json:"pernr"`
	SNAME  string `json:"sname"`
	REGION string `json:"region"`
	RGDESC string `json:"rgdesc"`
	MAINBR string `json:"mainbr"`
	MBDESC string `json:"mbdesc"`
	BRANCH string `json:"branch"`
	BRDESC string `json:"brdesc"`
}

type LeadAutocomplete struct {
	PERNR    string `json:"pernr"`
	SNAME    string `json:"sname"`
	ORGEH    int64  `json:"orgeh"`
	ORGEH_TX string `json:"orgeh_tx"`
}

//	type LeadAutocompleteRequest struct {
//		ORGEH   string `json:"orgeh"`
//		Keyword string `json:"keyword"`
//	}

type LeadAutocompleteRequest struct {
	Region  string `json:"region" example:"DKI JAKARTA"`
	Keyword string `json:"keyword" example:"KCP PONDOK INDAH"`
}

type TasklistRequestOne struct {
	ID int64 `json:"id"  example:"1"`
}

type DataVerifikasiRequest struct {
	TasklistID int64 `json:"tasklist_id" example:"1"`
}

type DataVerifikasiResponse struct {
	Branch          string `json:"kode_branch"`
	BRDESC          string `json:"unit_kerja"`
	MBDESC          string `json:"mbdesc"`
	RGDESC          string `json:"rgdesc"`
	NoPelaporan     string `json:"no_pelaporan"`
	Activity        string `json:"activity"`
	SubAktivitas    string `json:"sub_activity"`
	Maker           string `json:"maker"`
	IDRiskEvent     string `json:"id_risk_event"`
	RiskEventName   string `json:"risk_event_name"`
	HasilVerifikasi string `json:"hasil_verifikasi"`
	IndikasiFraud   string `json:"indikasi_fraud"`
}

type TasklistsRejected struct {
	TasklistID int64  `json:"tasklist_id" example:"1"`
	Notes      string `json:"notes" example:"notes"`
	Status     string `json:"status" example:"rejected"`
}

type UserRegionRequest struct {
	Branch int64 `json:"branch" example:"123"`
}

type UserRegionResponse struct {
	Region string `json:"region"`
}

func (pr UserRegionResponse) TableName() string {
	return "dwh_branch"
}

func (pr TasklistsRejected) TableName() string {
	return "tasklists_rejected"
}

func (pr TasklistsRequest) TableName() string {
	return "tasklists"
}

func (pr TasklistsResponse) TableName() string {
	return "tasklists"
}

func (pr TasklistsCountResponse) TableName() string {
	return "tasklists"
}

func (pr TasklistsCheckResponse) TableName() string {
	return "tasklists"
}
