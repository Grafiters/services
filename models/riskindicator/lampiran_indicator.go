package models

type LampiranIndicator struct {
	ID            int64
	IDIndicator   int64
	NamaLampiran  string
	NomorLampiran string
	JenisFile     string
	Path          string
	Filename      string
}

func (li LampiranIndicator) TableName() string {
	return "risk_indicator_map_files"
}
