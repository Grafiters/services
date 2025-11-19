package models

type RequestReportList struct {
	TypeID int    `json:"type_id"`
	Code   string `json:"code"`
	Pernr  string `json:"pernr"`
	Page   int    `json:"page"`
	Limit  int    `json:"limit"`
	Order  string `json:"order"`
	Offset int    `json:"offset"`
	Sort   string `json:"sort"`
}

type ReportListQuery struct {
	ID          int     `json:"id"`
	Code        string  `json:"code"`
	Name        string  `json:"name"`
	CreatedDate *string `json:"created_date"`
	Pernr       string  `json:"pernr"`
}

type ReportUrl struct {
	Url      string `json:"url"`
	Filename string `json:"filename"`
}
