package requests

type TransactionRequest struct {
	Appname       string      `json:"appname" binding:"required" `
	Prefix        string      `json:"prefix" binding:"required" "max=5,min=5` //harus 5 char
	ExpiredDate   string      `json:"expiredDate" binding:"required"`         // format yyyy-mm-dd
	Data          interface{} `json:"data" binding:"required"`
	ReferenceCode string      `json:"referenceCode"`
	Status        string      `json:"status"`
	Created       string      `json:"created"`
	LastUpdate    string      `json:"lastUpdate"`
}

type Data map[string]interface{}

type UpdateRequest struct {
	Appname       string `json:"appname" binding:"required" `
	ReferenceCode string `json:"referenceCode" binding:"required"`
}

type InquiryRequest struct {
	Appname       string `json:"appname" binding:"required" `
	ReferenceCode string `json:"referenceCode" binding:"required"`
}

type ReferenceSequenceRequest struct {
	Id       int64  `json:"id"`
	Prefix   string `json:"prefix"`
	Sequence int64  `json:"sequence"`
}
