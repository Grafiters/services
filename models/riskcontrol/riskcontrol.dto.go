package models

import "riskmanagement/lib"

type RiskControlRequest struct {
	ID          int64   `json:"id"`
	Kode        string  `json:"kode"`
	RiskControl string  `json:"risk_control"`
	ControlType string  `json:"control_type"`
	Nature      string  `json:"nature"`
	KeyControl  string  `json:"key_control"`
	Deskripsi   string  `json:"deskripsi"`
	Status      bool    `json:"status"`
	CreatedAt   *string `json:"created_at"`
	UpdatedAt   *string `json:"updated_at"`
}

type Paginate struct {
	Order       string `json:"order"`
	Sort        string `json:"sort"`
	Offset      int    `json:"offset"`
	Limit       int    `json:"limit"`
	Page        int    `json:"page"`
	Kode        string `json:"kode"`
	RiskControl string `json:"risk_control"`
	Status      string `json:"status"`
}

type RiskControlResponse struct {
	ID          int64   `json:"id"`
	Kode        string  `json:"kode"`
	RiskControl string  `json:"risk_control"`
	ControlType string  `json:"control_type"`
	Nature      string  `json:"nature"`
	KeyControl  string  `json:"key_control"`
	Deskripsi   string  `json:"deskripsi"`
	Status      bool    `json:"status"`
	CreatedAt   *string `json:"created_at"`
	UpdatedAt   *string `json:"updated_at"`
}

type KeywordRequest struct {
	Order       string `json:"order"`
	Sort        string `json:"sort"`
	Offset      int    `json:"offset"`
	Limit       int    `json:"limit"`
	Page        int    `json:"page"`
	Keyword     string `json:"keyword"`
	RiskIssueId string `json:"risk_issue_id"`
}

type RiskControlResponses struct {
	ID          int64  `json:"id"`
	Kode        string `json:"kode"`
	RiskControl string `json:"risk_control"`
}

type RiskControlResponsesNull struct {
	ID          lib.NullInt64  `json:"id"`
	Kode        lib.NullString `json:"kode"`
	RiskControl lib.NullString `json:"risk_control"`
}

type KodeRiskControl struct {
	KodeRiskControl string `json:"kode_risk_control"`
}

func (p RiskControlRequest) ParseRequest() RiskControl {
	return RiskControl{
		ID:          p.ID,
		Kode:        p.Kode,
		RiskControl: p.RiskControl,
		Deskripsi:   p.Deskripsi,
		Status:      p.Status,
	}
}

func (p RiskControlResponse) ParseResonse() RiskControl {
	return RiskControl{
		ID:          p.ID,
		Kode:        p.Kode,
		RiskControl: p.RiskControl,
		Deskripsi:   p.Deskripsi,
		Status:      p.Status,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func (rc RiskControlRequest) TableName() string {
	return "risk_control"
}

func (rc RiskControlResponse) TableName() string {
	return "risk_control"
}
