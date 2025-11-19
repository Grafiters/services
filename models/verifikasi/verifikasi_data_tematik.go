package models

type VerifikasiDataTematik struct {
	ID           int64
	VerifikasiId int64
	Periode      string
	Columns      string
	ColumnsData  string
	Status       bool
}

type VerifikasiDataTematikRequest struct {
	ID           int64  `json:"id"`
	VerifikasiId int64  `json:"verifikasi_id"`
	Periode      string `json:"periode"`
	Columns      string `json:"columns"`
	ColumnsData  string `json:"columns_data"`
	Status       bool   `json:"status"`
}

type VerifikasiDataTematikResponse struct {
	ID           int64  `json:"id"`
	VerifikasiId int64  `json:"verifikasi_id"`
	Periode      string `json:"periode"`
	Columns      string `json:"columns"`
	ColumnsData  string `json:"columns_data"`
	Status       bool   `json:"status"`
}

func (v VerifikasiDataTematik) TableName() string {
	return "verifikasi_data_anomali_tematik"
}

func (v VerifikasiDataTematikRequest) TableName() string {
	return "verifikasi_data_anomali_tematik"
}

func (v VerifikasiDataTematikResponse) TableName() string {
	return "verifikasi_data_anomali_tematik"
}
