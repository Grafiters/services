package models

type MapProductRequest struct {
	ID          int64 `json:"id"`
	IDRiskIssue int64 `json:"id_risk_issue"`
	Product     int64 `json:"product"`
}

type MapProductResponse struct {
	ID          int64 `json:"id"`
	IDRiskIssue int64 `json:"id_risk_issue"`
	Product     int64 `json:"product"`
}

type MapProductResponseFinal struct {
	ID          int64  `json:"id"`
	IDRiskIssue int64  `json:"id_risk_issue"`
	Product     int64  `json:"product"`
	ProductDesc string `json:"product_desc"`
}

func (p MapProductRequest) ParseRequest() MapProduct {
	return MapProduct{
		ID:          p.ID,
		IDRiskIssue: p.IDRiskIssue,
		Product:     p.Product,
	}
}

func (p MapProductResponse) ParseResponse() MapProduct {
	return MapProduct{
		ID:          p.ID,
		IDRiskIssue: p.IDRiskIssue,
		Product:     p.Product,
	}
}

func (mp MapProductRequest) TableName() string {
	return "risk_issue_map_product"
}

func (mp MapProductResponse) TableName() string {
	return "risk_issue_map_product"
}
