package pekerja

type DataPekerjaResponse struct {
	Pernr   string `json:"pernr"`
	Sname   string `json:"sname"`
	Stell   string `json:"stell"`
	StellTx string `json:"stell_tx"`
	Branch  string `json:"branch"`
}

type RequestApproval struct {
	Keyword string `json:"keyword"`
	Limit   int64  `json:"limit"`
	Offset  int64  `json:"offset"`
}

type PekerjaUkerRequest struct {
	Branch string `json:"branch"`
	Kostl  string `json:"kostl"`
	Werks  string `json:"werks"`
}
