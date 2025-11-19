package models

type McsRequest struct {
	Clientid     string `json:"clientid"`
	Clientsecret string `json:"clientsecret"`
	Keyword      string `json:"keyword"`
	Limit        int    `json:"limit"`
	Offset       int    `json:"offset"`
}

type PICResponse struct {
	PERNR string `json:"PERNR"`
	HTEXT string `json:"HTEXT"`
	NAMA  string `json:"NAMA"`
}

type PICResponseString struct {
	PERNR string `json:"PERNR"`
	HTEXT string `json:"HTEXT"`
	NAMA  string `json:"NAMA"`
}

type UkerResponse struct {
	BRNAME string `json:"BRNAME"`
	BRANCH string `json:"BRANCH"`
}

type UkerResponseString struct {
	BRNAME string `json:"BRNAME"`
	BRANCH string `json:"BRANCH"`
}

type KeywordRequest struct {
	Keyword string `json:"keyword"`
	Limit   int    `json:"limit"`
	Offset  int    `json:"offset"`
}

type JabatanResponse struct {
	HILFM string `json:"HILFM"`
	HTEXT string `json:"HTEXT"`
}

type OrgehResponse struct {
	// WERKS   string `json:"WERKS"`
	// WERKSTX string `json:"WERKSTX"`
	// BTRTL   string `json:"BTRTL"`
	// BTRTLTX string `json:"BTRTLTX"`
	// KOSTL   string `json:"KOSTL"`
	// KOSTLTX string `json:"KOSTLTX"`
	ORGEH   string `json:"ORGEH"`
	ORGEHTX string `json:"ORGEHTX"`
}
