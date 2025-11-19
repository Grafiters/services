package datatematik

type UpdaterData struct {
	RequestUpdate []RequestUpdate `json:"request_update"`
	Pernr         string          `json:"pernr"`
}
type RequestUpdate struct {
	NamaTable    string `json:"nama_table"`
	Id           int64  `json:"id"`
	Status       string `json:"status"`
	NoVerifikasi string `json:"no_verifikasi"`
}

type DataTematikRequest struct {
	RiskEvent     string `json:"risk_event"`
	RiskIndicator string `json:"risk_indicator"`
	NamaTable     string `json:"nama_table"`
	PeriodeData   string `json:"periode_data"`
	UnitKerja     string `json:"unit_kerja"`
	Limit         int64  `json:"limit"`
	Offset        int64  `json:"offset"`
}

type DataTematikResponse struct {
	Columns     string        `json:"columns"`
	ColumnsData []interface{} `json:"columns_data"`
}
