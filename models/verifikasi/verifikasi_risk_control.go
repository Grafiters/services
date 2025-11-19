package models

type VerifikasiRiskControl struct {
	ID            int64
	VerifikasiId  int64
	RiskControlID int64
	RiskControl   string
}

func (vc VerifikasiRiskControl) TableName() string {
	return "verifikasi_risk_control"
}
