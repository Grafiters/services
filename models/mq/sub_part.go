package models

type SubPartListRequest struct {
	ID     int64  `json:"id"`
	MenuID int64  `json:"menu_id"`
	Code   string `json:"code"`
	TypeID int64  `json:"type_id"`
	PartID int64  `json:"part_id"`
	Nama   string `json:"nama"`
	Page   int    `json:"page"`
	Limit  int    `json:"limit"`
	Order  string `json:"order"`
	Offset int    `json:"offset"`
	Sort   string `json:"sort"`
	PERNR  string `json:"pernr"`
}

type SubPart struct {
	ID            int16   `json:"id"`
	Code          string  `json:"code"`
	TypeID        int64   `json:"type_id"`
	PartID        int64   `json:"part_id"`
	Name          string  `json:"name"`
	SubPart       int64   `json:"sub_part"`
	Status        string  `json:"status"`
	Pernr         string  `json:"pernr"`
	SName         string  `gorm:"column:sname" json:"sname"`
	ActiveDate    *string `json:"active_date"`
	NonactiveDate *string `json:"nonactive_date"`
	CreatedAt     *string
	UpdatedAt     *string
}

type SubPartEditRequest struct {
	ID            int16   `json:"id"`
	Code          string  `json:"code"`
	TypeID        int64   `json:"type_id"`
	PartID        int64   `json:"part_id"`
	Name          string  `json:"name"`
	SubPart       int64   `json:"sub_part"`
	Status        string  `json:"status"`
	StatusOld     string  `json:"status_old"`
	Pernr         string  `json:"pernr"`
	SName         string  `gorm:"column:sname" json:"sname"`
	ActiveDate    *string `json:"active_date"`
	NonactiveDate *string `json:"nonactive_date"`
	CreatedAt     *string `json:"created_at"`
	UpdatedAt     *string
}
