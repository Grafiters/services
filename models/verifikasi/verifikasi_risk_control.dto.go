package models

type VerifikasiRiskControlRequest struct {
	ID            int64  `json:"id"`
	VerifikasiId  int64  `json:"verifikasi_id"`
	RiskControlID int64  `json:"risk_control_id"`
	RiskControl   string `json:"risk_control"`
}

type VerifikasiRiskControlResponse struct {
	ID            int64  `json:"id"`
	VerifikasiId  int64  `json:"verifikasi_id"`
	RiskControlID int64  `json:"risk_control_id"`
	RiskControl   string `json:"risk_control"`
}

type VerifikasiRiskControlResponses struct {
	ID           int64  `json:"id"`
	VerifikasiId int64  `json:"verifikasi_id"`
	RiskControl  string `json:"risk_control"`
}

func (p VerifikasiRiskControlRequest) ParseRequest() VerifikasiRiskControl {
	return VerifikasiRiskControl{
		ID:            p.ID,
		VerifikasiId:  p.VerifikasiId,
		RiskControlID: p.RiskControlID,
	}
}

func (p VerifikasiRiskControlResponse) ParseResponse() VerifikasiRiskControl {
	return VerifikasiRiskControl{
		ID:            p.ID,
		VerifikasiId:  p.VerifikasiId,
		RiskControlID: p.RiskControlID,
	}
}

func (vc VerifikasiRiskControlRequest) TableName() string {
	return "verifikasi_risk_control"
}

func (vc VerifikasiRiskControlResponse) TableName() string {
	return "verifikasi_risk_control"
}

func (vc VerifikasiRiskControlResponses) TableName() string {
	return "verifikasi_risk_control"
}
