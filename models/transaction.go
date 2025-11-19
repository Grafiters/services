package models

// User model
type Transaction struct {
	Id            string
	Appname       string      `json:"appname"`
	Prefix        string      `json:"prefix"`
	Data          interface{} `json:"data"`
	ExpiredDate   string      `json:"expiredDate"`
	ReferenceCode string      `json:"referenceCode"`
	Status        string      `json:"status"`
	Created       string      `json:"created"`
	LastUpdate    string      `json:"lastUpdate"`
}

type TransactionExpired struct {
	Id            string
	Appname       string      `json:"appname"`
	Prefix        string      `json:"prefix"`
	Data          interface{} `json:"data"`
	ExpiredDate   string      `json:"expiredDate"`
	ReferenceCode string      `json:"referenceCode"`
	Status        string      `json:"status"`
	Created       string      `json:"created"`
	LastUpdate    string      `json:"lastUpdate"`
}

type TransactionExecuted struct {
	Id            string
	Appname       string      `json:"appname"`
	Prefix        string      `json:"prefix"`
	Data          interface{} `json:"data"`
	ExpiredDate   string      `json:"expiredDate"`
	ReferenceCode string      `json:"referenceCode"`
	Status        string      `json:"status"`
	Created       string      `json:"created"`
	LastUpdate    string      `json:"lastUpdate"`
}

type ReferenceCodeCounter struct {
	Id       int64  `json:"id"`
	Prefix   string `json:"prefix"`
	Sequence int64  `json:"sequence"`
}

type LogAccess struct {
	Id        string
	Uri       string `json:"uri"`
	Agent     string `json:"agent"`
	Ipaddress string `json:"ipaddress"`
	Request   string `json:"request"`
	Response  string `json:"response"`
	Date      string `json:"date"`
}

type LogError struct {
	Id            string
	Filename      string `json:"filename"`
	Method        string `json:"method"`
	Line          string `json:"line"`
	MessageCustom string `json:"messageCustom"`
	MessageSystem string `json:"messageSystem"`
	Date          string `json:"date"`
}

type Data map[string]interface{}

// Table Index gives table name of model
func (u Transaction) IndexTransactionOpen() string {
	return "transaction_open"
}

func (u Transaction) IndexTransactionExecuted() string {
	return "transaction_executed"
}

func (u TransactionExpired) IndexTransactionExpired() string {
	return "transaction_expired"
}

func (u Transaction) IndexReferenceSequence() string {
	return "reference_code_counter"
}

func (u LogAccess) IndexLogAccess(period string) string {
	return "log_access_" + period
}

func (u LogError) IndexLogError(period string) string {
	return "log_error_" + period
}
