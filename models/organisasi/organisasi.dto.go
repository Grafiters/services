package organisasi

type CostCenterRequest struct {
	Type string `json:"type"`
}

type CostCenterResponse struct {
	Werks   string `json:"werks"`
	WerksTx string `json:"werks_tx"`
}

type DepartmentRequest struct {
	Type    string `json:"type"`
	Werks   string `json:"werks"`
	Keyword string `json:"keyword"`
	Limit   int    `json:"limit"`
	Offset  int    `json:"offset"`
}

type DepartmentResponse struct {
	Orgeh   string `json:"orgeh"`
	OrgehTx string `json:"orgeh_tx"`
}

type JabatanRequest struct {
	Keyword string `json:"keyword"`
	Limit   int    `json:"limit"`
	Offset  int    `json:"offset"`
}

type JabatanResponse struct {
	Hilfm string `json:"hilfm"`
	Htext string `json:"htext"`
}
