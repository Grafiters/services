package responses

type TransactionResponse struct {
	Appname       string      `json:"appname"`
	Data          interface{} `json:"data"`
	Prefix        string      `json:"prefix"`
	ExpiredDate   string      `json:"expiredDate"`
	ReferenceCode string      `json:"referenceCode"`
	Status        string      `json:"status"`
}

type TransactionInquiryResponse struct {
	Data interface{} `json:"data"`
}

type TransactionCreateResponse struct {
	ReferenceCode string `json:"referenceCode"`
}

type Data interface{}

type ReferenceSequenceResponse struct {
	Id            int64  `json:"id"`
	Prefix        string `json:"prefix"`
	Sequence      int64  `json:"sequence"`
	ReferenceCode string `json:"referenceCode"`
}
