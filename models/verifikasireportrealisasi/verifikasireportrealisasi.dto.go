package models

type ReportRealisasiKreditListRequest struct {
	Limit            int    `json:"limit"`
	Offset           int    `json:"offset"`
	Page             int    `json:"page"`
	ReportType       string `json:"report_type"`
	REGION           string `json:"REGION"`
	MAINBR           string `json:"MAINBR"`
	BRANCH           string `json:"BRANCH"`
	NoPelaporan      string `json:"no_pelaporan"`
	Criteria         string `json:"criteria"`
	Segment          string `json:"segment"`
	Product          string `json:"product"`
	StatusVerifikasi string `json:"status_verifikasi"`
	ButuhPerbaikan   string `json:"butuh_perbaikan"`
	IndikasiFraud    string `json:"indikasi_fraud"`
	StartDate        string `json:"start_date"`
	EndDate          string `json:"end_date"`
	Sort             string `json:"sort"`
	Pernr            string `json:"pernr"`
	Timestime        string `json:"timestime"`
}

type ReportRealisasiKreditSummaryRequest struct {
	Limit      int      `json:"limit"`
	Offset     int      `json:"offset"`
	Page       int      `json:"page"`
	ReportType string   `json:"report_type"`
	REGION     string   `json:"REGION"`
	MAINBR     string   `json:"MAINBR"`
	BRANCH     string   `json:"BRANCH"`
	Criteria   string   `json:"criteria"`
	GroupBy    []string `json:"group_by"`
	Product    []string `json:"product"`
	Periode    string   `json:"periode"`
	StartDate  string   `json:"start_date"`
	EndDate    string   `json:"end_date"`
	Sort       string   `json:"sort"`
	Pernr      string   `json:"pernr"`
	Timestime  string   `json:"timestime"`
}

type SegmentRealisasiKreditRequest struct {
	Segment  string `json:"segment"`
	IsActive *int   `json:"is_active"`
	Limit    int    `json:"limit"`
	Offset   int    `json:"offset"`
}

type ReportRealisasiKreditListResponseRaw struct {
	VerifikasiID     int64  `json:"verifikasi_id"`
	NoPelaporan      string `json:"no_pelaporan"`
	REGION           string `json:"REGION"`
	RGDESC           string `json:"RGDESC"`
	MAINBR           string `json:"MAINBR"`
	MBDESC           string `json:"MBDESC"`
	BRANCH           string `json:"BRANCH"`
	BRDESC           string `json:"BRDESC"`
	ActivityId       int64  `json:"activity_id"`
	ActivityName     string `json:"activity_name"`
	ProductId        int64  `json:"product_id"`
	ProductName      string `json:"product_name"`
	PeriodeData      string `json:"periode_data"`
	RestruckFlag     int64  `json:"restruck_flag"`
	ButuhPerbaikan   int64  `json:"butuh_perbaikan"`
	KriteriaData     string `json:"kriteria_data"`
	Segment          string `json:"segment"`
	CreatedId        string `json:"created_id"`
	CreatedDesc      string `json:"created_desc"`
	DataRealisasi    string `json:"data_realisasi"`
	StatusVerifikasi int64  `json:"status_verifikasi"`
	HasilVerifikasi  string `json:"hasil_verifikasi"`
	KunjunganNasabah int64  `json:"kunjungan_nasabah"`
	TglKunjungan     string `json:"tgl_kunjungan"`
	LampiranName     string `json:"lampiran_name"`
	LampiranPath     string `json:"lampiran_path"`
}

type ReportRealisasiKreditSummaryResponseRaw struct {
	TotalVerifikasi  int64  `json:"total_verifikasi"`
	ProductId        int64  `json:"product_id"`
	ProductName      string `json:"product_name"`
	CreatedId        string `json:"created_id"`
	CreatedDesc      string `json:"created_desc"`
	REGION           string `json:"REGION"`
	RGDESC           string `json:"RGDESC"`
	MAINBR           string `json:"MAINBR"`
	MBDESC           string `json:"MBDESC"`
	BRANCH           string `json:"BRANCH"`
	BRDESC           string `json:"BRDESC"`
	StatusVerifikasi int64  `json:"status_verifikasi"`
	DataRealisasi    string `json:"data_realisasi"`
	Efektif          int64  `json:"efektif"`
	NonEfektif       int64  `json:"non_efektif"`
	KriteriaData     string `json:"kriteria_data"`
}

type ReportRealisasiKreditSummaryDownloadResponseRaw struct {
	TotalVerifikasi  int64  `json:"total_verifikasi"`
	ProductId        int64  `json:"product_id"`
	ProductName      string `json:"product_name"`
	CreatedId        string `json:"created_id"`
	CreatedDesc      string `json:"created_desc"`
	REGION           string `json:"REGION"`
	RGDESC           string `json:"RGDESC"`
	MAINBR           string `json:"MAINBR"`
	MBDESC           string `json:"MBDESC"`
	BRANCH           string `json:"BRANCH"`
	BRDESC           string `json:"BRDESC"`
	StatusVerifikasi string `json:"status_verifikasi"`
	DataRealisasi    string `json:"data_realisasi"`
	Efektif          int64  `json:"efektif"`
	NonEfektif       int64  `json:"non_efektif"`
	KriteriaData     string `json:"kriteria_data"`
}

type ReportRealisasiKreditListResponse struct {
	VerifikasiID     int64       `json:"verifikasi_id"`
	NoPelaporan      string      `json:"no_pelaporan"`
	REGION           string      `json:"REGION"`
	RGDESC           string      `json:"RGDESC"`
	MAINBR           string      `json:"MAINBR"`
	MBDESC           string      `json:"MBDESC"`
	BRANCH           string      `json:"BRANCH"`
	BRDESC           string      `json:"BRDESC"`
	ActivityId       int64       `json:"activity_id"`
	ActivityName     string      `json:"activity_name"`
	ProductId        int64       `json:"product_id"`
	ProductDesc      int64       `json:"product_desc"`
	ProductName      string      `json:"product_name"`
	PeriodeData      string      `json:"periode_data"`
	RestruckFlag     int64       `json:"restruck_flag"`
	ButuhPerbaikan   int64       `json:"butuh_perbaikan"`
	KriteriaData     []string    `json:"kriteria_data"`
	Segment          string      `json:"segment"`
	CreatedId        string      `json:"created_id"`
	CreatedDesc      string      `json:"created_desc"`
	DataRealisasi    interface{} `json:"data_realisasi"`
	StatusVerifikasi int64       `json:"status_verifikasi"`
	HasilVerifikasi  string      `json:"hasil_verifikasi"`
	KunjunganNasabah int64       `json:"kunjungan_nasabah"`
	TglKunjungan     string      `json:"tgl_kunjungan"`
	LampiranName     string      `json:"lampiran_name"`
	LampiranPath     string      `json:"lampiran_path"`
}

type ReportRealisasiKreditSummaryResponse struct {
	TotalVerifikasi  int64       `json:"total_verifikasi"`
	ProductId        int64       `json:"product_id"`
	ProductName      string      `json:"product_name"`
	CreatedId        string      `json:"created_id"`
	CreatedDesc      string      `json:"created_desc"`
	REGION           string      `json:"REGION"`
	RGDESC           string      `json:"RGDESC"`
	MAINBR           string      `json:"MAINBR"`
	MBDESC           string      `json:"MBDESC"`
	BRANCH           string      `json:"BRANCH"`
	BRDESC           string      `json:"BRDESC"`
	StatusVerifikasi int64       `json:"status_verifikasi"`
	DataRealisasi    interface{} `json:"data_realisasi"`
	Efektif          int64       `json:"efektif"`
	NonEfektif       int64       `json:"non_efektif"`
	KriteriaData     []string    `json:"kriteria_data"`
}

type SegmentRealisasiKreditResponse struct {
	Id       int64  `json:"id"`
	Segment  string `json:"segment"`
	IsActive int64  `json:"product_id"`
}

type ReportRealisasiKreditSummaryDownloadResponse struct {
	TotalVerifikasi  int64       `json:"total_verifikasi"`
	ProductId        int64       `json:"product_id"`
	ProductName      string      `json:"product_name"`
	CreatedId        string      `json:"created_id"`
	CreatedDesc      string      `json:"created_desc"`
	REGION           string      `json:"REGION"`
	RGDESC           string      `json:"RGDESC"`
	MAINBR           string      `json:"MAINBR"`
	MBDESC           string      `json:"MBDESC"`
	BRANCH           string      `json:"BRANCH"`
	BRDESC           string      `json:"BRDESC"`
	StatusVerifikasi string      `json:"status_verifikasi"`
	DataRealisasi    interface{} `json:"data_realisasi"`
	Efektif          int64       `json:"efektif"`
	NonEfektif       int64       `json:"non_efektif"`
	KriteriaData     []string    `json:"kriteria_data"`
	MstKriteria      interface{} `json:"mst_kriteria"`
}

type GeneratorRealPinRptList struct {
	NoPelaporan      string         `json:"no_pelaporan"`
	REGION           string         `json:"REGION"`
	RGDESC           string         `json:"RGDESC"`
	MAINBR           string         `json:"MAINBR"`
	MBDESC           string         `json:"MBDESC"`
	BRANCH           string         `json:"BRANCH"`
	BRDESC           string         `json:"BRDESC"`
	ActivityId       int64          `json:"activity_id"`
	ActivityName     string         `json:"activity_name"`
	ProductId        int64          `json:"product_id"`
	ProductDesc      int64          `json:"product_desc"`
	ProductName      string         `json:"product_name"`
	PeriodeData      string         `json:"periode_data"`
	RestruckFlag     int64          `json:"restruck_flag"`
	ButuhPerbaikan   int64          `json:"butuh_perbaikan"`
	KriteriaData     []string       `json:"kriteria_data"`
	Segment          string         `json:"segment"`
	CreatedId        string         `json:"created_id"`
	CreatedDesc      string         `json:"created_desc"`
	DataRealisasi    interface{}    `json:"data_realisasi"`
	StatusVerifikasi int64          `json:"status_verifikasi"`
	HasilVerifikasi  string         `json:"hasil_verifikasi"`
	KunjunganNasabah int64          `json:"kunjungan_nasabah"`
	TglKunjungan     string         `json:"tgl_kunjungan"`
	ListKriteria     []ListKriteria `json:"list_kriteria"`
}

type ListKriteria struct {
	IdCriteria int64  `json:"id_criteria"`
	Kriteria   string `json:"kriteria"`
}
