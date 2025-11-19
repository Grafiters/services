package models

type HeaderRequest struct {
	Username string      `json:"username"`
	Password string      `json:"password"`
	Request  RequestBody `json:"request"`
}

// type Request struct {
// 	Request []RequestBody `json:"request"`
// }

type RequestBody struct {
	Periode       string `json:"periode"`
	UnitKerja     string `json:"unitKerja"`
	KodeIndikator string `json:"kodeIndikator"`
	Produk        string `json:"produk"`
}
