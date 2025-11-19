package models

type ActivityDailyRequest struct {
	Order      string `json:"order"`
	Sort       string `json:"sort"`
	Offset     int    `json:"offset"`
	Limit      int    `json:"limit"`
	Page       int    `json:"page"`
	REGION     string `json:"region"`
	PERNR      string `json:"pernr"`
	PN         string `json:"pn"`
	Period     string `json:"period"`
	Persentase string `json:"persentase"`
}

type ActivityDailyResponse struct {
	PERNR        string             `json:"pernr"`
	Nama         string             `json:"nama"`
	Kanwil       string             `json:"kanwil"`
	UkerKelolaan []UkerListResponse `json:"uker_kelolaan"`
	Kegiatan     string             `json:"kegiatan"`
	ActivityID   string             `json:"activity_id"`
	ProductID    string             `json:"product_id"`
	RiskIssueID  int                `json:"risk_issue_id"`
	Sample       string             `json:"sample"`
	Progres      string             `json:"progres"`
	Persentase   int64              `json:"persentase"`
	RiskEvent    string             `json:"risk_event"`
}

type UkerListRequest struct {
	PERNR string `json:"pernr"`
}

type PersentaseTotalRequest struct {
	PERNR      string `json:"pernr"`
	Persentase string `json:"persentase"`
}

type UkerListResponse struct {
	BRDESC string `json:"brdesc"`
}

type ActivityDailyDetailRequest struct {
	Order       string `json:"order"`
	Sort        string `json:"sort"`
	Offset      int    `json:"offset"`
	Limit       int    `json:"limit"`
	Page        int    `json:"page"`
	PERNR       string `json:"pernr"`
	Period      string `json:"period"`
	Kegiatan    string `json:"kegiatan"`
	ActivityID  int    `json:"activity_id"`
	ProductID   int    `json:"product_id"`
	RiskIssueID int    `json:"risk_issue_id"`
	TaskType    int    `json:"task_type"`
}

type ActivityDailyDetailResponse struct {
	PERNR           string             `json:"pernr"`
	Nama            string             `json:"nama"`
	Kanwil          string             `json:"kanwil"`
	Kanca           string             `json:"kanca"`
	UnitKerja       string             `json:"unit_kerja"`
	UkerKelolaan    []UkerListResponse `json:"uker_kelolaan"`
	Kegiatan        string             `json:"kegiatan"`
	RiskEvent       string             `json:"risk_event"`
	TaskType        string             `json:"task_type"`
	Period          string             `json:"period"`
	Sample          string             `json:"sample"`
	Progres         string             `json:"progres"`
	Persentase      string             `json:"persentase"`
	AssignedCreated string             `json:"assigned_created"`
	StartDate       string             `json:"start_date"`
	EndDate         string             `json:"end_date"`
}

type ActivityDaily struct {
	PERNR       string `json:"pernr"`
	Nama        string `json:"nama"`
	Kanwil      string `json:"kanwil"`
	Kegiatan    string `json:"kegiatan"`
	ActivityID  string `json:"activity_id"`
	ProductID   string `json:"product_id"`
	RiskIssueID int    `json:"risk_issue_id"`
	Sample      string `json:"sample"`
	Progres     string `json:"progres"`
	Persentase  string `json:"persentase"`
	RiskEvent   string `json:"risk_event"`
}

type ActivityDailyDetail struct {
	PERNR           string `json:"pernr"`
	Nama            string `json:"nama"`
	Kanwil          string `json:"kanwil"`
	Kanca           string `json:"kanca"`
	UnitKerja       string `json:"unit_kerja"`
	Kegiatan        string `json:"kegiatan"`
	RiskEvent       string `json:"risk_event"`
	TaskType        string `json:"task_type"`
	Period          string `json:"period"`
	Sample          string `json:"sample"`
	Progres         string `json:"progres"`
	Persentase      string `json:"persentase"`
	AssignedCreated string `json:"assigned_created"`
	StartDate       string `json:"start_date"`
	EndDate         string `json:"end_date"`
}
