package models

import "riskmanagement/lib"

type PartCodeNullResponse struct {
	Code lib.NullString `json:"code"`
}

type PartRequest struct {
	ID            int64   `json:"id"`
	Code          string  `json:"code"`
	TypeID        int64   `json:"type_id"`
	Name          string  `json:"name"`
	SubPart       int64   `json:"sub_part"`
	PartID        int64   `json:"part_id"`
	ViewValue     string  `json:"view_value"`
	Status        string  `json:"status"`
	OldStatus     string  `json:"old_status"`
	Order         string  `json:"order"`
	Sort          string  `json:"sort"`
	Offset        int     `json:"offset"`
	Limit         int     `json:"limit"`
	Page          int     `json:"page"`
	PERNR         string  `json:"pernr"`
	SNAME         string  `json:"sname"`
	ActiveDate    *string `json:"active_date"`
	NonactiveDate *string `json:"nonactive_date"`
	CreatedAt     *string `json:"created_at"`
	UpdatedAt     *string `json:"updated_at"`
}

type Part struct {
	ID            int64
	Code          string
	TypeID        int64
	Name          string
	SubPart       int64
	PartID        int64
	ViewValue     string
	Status        string
	PERNR         string
	SNAME         string
	ActiveDate    *string
	NonactiveDate *string
	CreatedAt     *string
	UpdatedAt     *string
}

type PartResponse struct {
	ID            int64       `json:"id"`
	MenuID        int64       `json:"menu_id"`
	Code          string      `json:"code"`
	TypeID        int64       `json:"type_id"`
	Type          []TypeDraft `json:"type"`
	Name          string      `json:"name"`
	SubPart       int64       `json:"sub_part"`
	PartID        int64       `json:"part_id"`
	ViewValue     string      `json:"view_value"`
	Status        string      `json:"status"`
	IsQuiz        string      `json:"is_quiz"`
	PERNR         string      `json:"pernr"`
	SNAME         string      `json:"sname"`
	ApprovedPart  int64       `json:"approved_part"`
	ActiveDate    *string     `json:"active_date"`
	NonactiveDate *string     `json:"nonactive_date"`
	CreatedAt     *string     `json:"created_at"`
	UpdatedAt     *string     `json:"updated_at"`
}

type PartTX struct {
	ID            int64
	Code          string
	TypeID        int64
	Name          string
	SubPart       int64
	PartID        int64
	ViewValue     string
	Status        string
	PERNR         string
	SNAME         string
	ActiveDate    *string
	NonactiveDate *string
	Versioning    int64
	CreatedAt     *string
	UpdatedAt     *string
}

func (Part) TableName() string {
	return "q_part"
}

func (PartTX) TableName() string {
	return "q_part_tx"
}
