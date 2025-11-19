package realisasimodels

type ParameterGetHeaderRequest struct {
	Pernr   string                     `json:"pernr"`
	Request ParameterBodyFilterRequest `json:"request"`
}

type ParameterBodyFilterRequest struct {
	PeriodeAwal  string `json:"periode_awal"`
	PeriodeAkhir string `json:"periode_akhir"`
	Limit        int64  `json:"limit"`
	Offset       int64  `json:"offset"`
	Pernr        string `json:"pernr"`
}

type ResponseData struct {
	Status     string        `json:"status"`
	Message    string        `json:"message"`
	Data       []interface{} `json:"data"`
	Pagination int           `json:"pagination"`
}

type StoreResponse struct {
	Status     string      `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Pagination interface{} `json:"pagination"`
}

type ParameterStoreHeaderRequest struct {
	Pernr   string               `json:"pernr"`
	Request ParameterBodyRequest `json:"request"`
}

type ParameterBodyRequest struct {
	NilaiNpl          float64 `json:"nilai_npl"`
	NilaiDpk          float64 `json:"nilai_dpk"`
	Rugi              string  `json:"rugi"`
	PeriodeKeragaan   string  `json:"periode_keragaan"`
	LastUpdatePernr   string  `json:"last_update_pernr"`
	LastUpdateSname   string  `json:"last_update_sname"`
	LastUpdateStelltx string  `json:"last_update_stelltx"`
	Pernr             string  `json:"pernr"`
}

type RevisiUkerGetHeaderRequest struct {
	Pernr   string               `json:"pernr"`
	Request RevisUkerBodyRequest `json:"request"`
}

type RevisiUkerStoreHeaderRequest struct {
	Pernr   string                 `json:"pernr"`
	Request RevisiUkerStoreRequest `json:"request"`
}

type RevisiUkerDeleteHeaderRequest struct {
	Pernr   string                  `json:"pernr"`
	Request RevisiUkerDeleteRequest `json:"request"`
}

type RevisUkerBodyRequest struct {
	REGION      string `json:"REGION"`
	MAINBR      string `json:"MAINBR"`
	BRANCH      string `json:"BRANCH"`
	JenisRevisi string `json:"jenis_revisi"`
	Limit       int64  `json:"limit"`
	Offset      int64  `json:"offset"`
	Pernr       string `json:"pernr"`
}

type RevisiUkerStoreRequest struct {
	REGION      string `json:"REGION"`
	RGDESC      string `json:"RGDESC"`
	MAINBR      string `json:"MAINBR"`
	MBDESC      string `json:"MBDESC`
	BRANCH      string `json:"BRANCH"`
	BRDESC      string `json:"BRDESC"`
	JenisRevisi string `json:"jenis_revisi"`
	UpdateId    string `json:"update_id"`
	UpdateName  string `json:"update_name"`
	UpdateStell string `json:"update_stell"`
	Pernr       string `json:"pernr"`
}

type RevisiUkerDeleteRequest struct {
	Id    int64  `json:"id"`
	Pernr string `json:"pernr"`
}

type RealisasiGetHeaderRequest struct {
	Pernr   string               `json:"pernr"`
	Request RealisasiBodyRequest `json:"request"`
}

type RealisasiBodyRequest struct {
	Periode  string `json:"periode"`
	Branch   string `json:"branch"`
	Restruck string `json:"restruck"`
	Limit    int    `json:"limit"`
	Offset   int    `json:"offset"`
	Pernr    string `json:"pernr"`
}

type ResponseSampleData struct {
	Status     string                `json:"status"`
	Message    string                `json:"message"`
	Data       []DataRealisasiKredit `json:"data"`
	Pagination int                   `json:"pagination"`
}

type DataRealisasiKredit struct {
	ID             int     `json:"id"`
	Periode        string  `json:"periode"`
	REGION         string  `json:"REGION"`
	RGDESC         string  `json:"RGDESC"`
	MAINBR         string  `json:"MAINBR"`
	MBDESC         string  `json:"MBDESC"`
	BRANCH         string  `json:"BRANCH"`
	BRDESC         string  `json:"BRDESC"`
	SEGMENT        string  `json:"SEGMENT"`
	PRODUK         string  `json:"PRODUK"`
	TIPE           string  `json:"TIPE"`
	Officr         string  `json:"officr"`
	NamaPemrakarsa string  `json:"nama_pemrakarsa"`
	PnPemrakarsa   string  `json:"pn_pemrakarsa"`
	NOMOR_REKENING string  `json:"NOMOR_REKENING"`
	NAMA_KREDITUR  string  `json:"NAMA_KREDITUR"`
	KOLEK          string  `json:"KOLEK"`
	LANCAR         string  `json:"LANCAR"`
	DPK            float64 `json:"DPK"`
	KURANG_LANCAR  float64 `json:"KURANG_LANCAR"`
	DIRAGUKAN      float64 `json:"DIRAGUKAN"`
	MACET          float64 `json:"MACET"`
	LOAN_TYPE      string  `json:"LOAN_TYPE"`
	PLAFOND        float64 `json:"PLAFOND"`
	NPDT           string  `json:"NPDT"`
	NIPDT7         string  `json:"NIPDT7"`
	RATE           float64 `json:"RATE"`
	FRELDT         string  `json:"FRELDT"`
	MATDT          string  `json:"MATDT"`
	FLGRES         string  `json:"FLGRES"`
	Restruck       string  `json:"restruck"`
	TglRestruck    string  `json:"tgl_restruck"`
	CIFNO          string  `json:"CIFNO"`
	Jgkwkt         string  `json:"jgkwkt"`
}

type RealisasiUpdateFlagRequest struct {
	Pernr   string            `json:"pernr"`
	Request RequestUpdateFlag `json:"request"`
}

type RequestUpdateFlag struct {
	Id             int64  `json:"id"`
	VerifikasiFlag bool   `json:"verifikasi_flag"`
	Pernr          string `json:"pernr"`
}

type UpdateFlagResponse struct {
	Status     string      `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Pagination interface{} `json:"pagination"`
}
