package models

type RiskTypeRequest struct {
	ID           int64   `json:"id"`
	RiskTypeCode string  `json:"risk_type_code"`
	RiskType     string  `json:"risk_type"`
	Deskripsi    string  `json:"deskripsi"`
	Status       bool    `json:"status"`
	CreatedAt    *string `json:"created_at"`
	UpdatedAt    *string `json:"updated_at"`
}

type Paginate struct {
	Order  string `json:"order"`
	Sort   string `json:"sort"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Page   int    `json:"page"`
}

type RiskTypeResponse struct {
	ID           int64   `json:"id"`
	RiskTypeCode string  `json:"risk_type_code"`
	RiskType     string  `json:"risk_type"`
	Deskripsi    string  `json:"deskripsi"`
	Status       bool    `json:"status"`
	CreatedAt    *string `json:"created_at"`
	UpdatedAt    *string `json:"updated_at"`
}

func (p RiskTypeRequest) ParseRequest() RiskType {
	return RiskType{
		ID:           p.ID,
		RiskTypeCode: p.RiskTypeCode,
		RiskType:     p.RiskType,
		Deskripsi:    p.Deskripsi,
		Status:       p.Status,
		CreatedAt:    new(string),
		UpdatedAt:    new(string),
	}
}

func (p RiskTypeResponse) ParseResponse() RiskType {
	return RiskType{
		ID:           p.ID,
		RiskTypeCode: p.RiskTypeCode,
		RiskType:     p.RiskType,
		Deskripsi:    p.Deskripsi,
		Status:       p.Status,
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
	}
}

func (pr RiskTypeRequest) TableName() string {
	return "risk_type"
}

func (pr RiskTypeResponse) TableName() string {
	return "risk_type"
}
