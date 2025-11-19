package models

type HistoriTaskDataVerifikasiPagianted struct {
	Order          string `json:"order"`
	Sort           string `json:"sort"`
	Offset         int    `json:"offset"`
	Limit          int    `json:"limit"`
	Page           int    `json:"page"`
	Pn             string `json:"pernr"`
	NoTasklist     string `json:"no_tasklist"`
	NamaTasklist   string `json:"nama_tasklist"`
	Nama           string `json:"nama"`
	KANWIL         string `json:"kanwil"`     //kanwil
	KANCA          string `json:"kanca"`      //kanca
	UnitKerja      string `json:"unit_kerja"` //uker
	RiskEvent      string `json:"risk_issue_id"`
	Aktifitas      string `json:"activity_id"`
	Produk         string `json:"product_id"`
	Indikator      string `json:"risk_indikator_id"`
	IndikatorOther string `json:"risk_indikator_other"`
	StatusApproval string `json:"status_approval"`
	Status         string `json:"status"`
	JenisTask      string `json:"jenis_task"`
	SumberData     string `json:"sumber_data"`
}

type HistoriTaskDataVerifikasiDownload struct {
	Order          string `json:"order"`
	Sort           string `json:"sort"`
	Pn             string `json:"pernr"`
	Nama           string `json:"nama"`
	NoTasklist     string `json:"no_tasklist"`
	NamaTasklist   string `json:"nama_tasklist"`
	KANWIL         string `json:"kanwil"`     //kanwil
	KANCA          string `json:"kanca"`      //kanca
	UnitKerja      string `json:"unit_kerja"` //uker
	RiskEvent      string `json:"risk_issue_id"`
	Aktifitas      string `json:"activity_id"`
	Produk         string `json:"product_id"`
	Indikator      string `json:"risk_indikator_id"`
	IndikatorOther string `json:"risk_indikator_other"`
	StatusApproval string `json:"status_approval"`
	Status         string `json:"status"`
	JenisTask      string `json:"jenis_task"`
	SumberData     string `json:"sumber_data"`
}

type HistoriTaskDataVerifikasiResult struct {
	ID             string `json:"id"`
	PN             string `json:"pn"`
	Nama           string `json:"nama"`
	NoTasklist     string `json:"no_tasklist"`
	NamaTasklist   string `json:"nama_tasklist"`
	Region         string `json:"region"`
	Kanwil         string `json:"kanwil"`
	Mainbr         string `json:"mainbr"`
	Kanca          string `json:"kanca"`
	Branch         string `json:"branch"`
	Uker           string `json:"uker"`
	Aktifitas      string `json:"aktifitas"`
	Product        string `json:"product"`
	Indikator      string `json:"indikator"`
	JenisTask      string `json:"jenis_task"`
	Period         string `json:"period"`
	TanggalMulai   string `json:"tanggal_mulai"`
	TanggalAkhir   string `json:"tanggal_akhir"`
	StatusApproval string `json:"status_approval"`
	Status         string `json:"status"`
	Kegiatan       string `json:"kegiatan"`
	RiskIssue      string `json:"risk_issue"`
}

type HistoriTaskDataVerifikasiDetailRequest struct {
	ID            string `json:"id"`
	Region        string `json:"region"`
	Mainbr        string `json:"mainbr"`
	Branch        string `json:"branch"`
	RiskIssue     string `json:"risk_issue"`
	RiskIndicator string `json:"risk_indicator"`
}
type HistoriTaskDataVerifikasiDetailResult struct {
	Branch                 string `json:"branch"`
	UnitKerja              string `json:"unit_kerja"`
	Kanwil                 string `json:"kanwil"`
	Kanca                  string `json:"kanca"`
	NoPelaporan            string `json:"no_pelaporan"`
	Aktifitas              string `json:"aktifitas"`
	SubAktifitas           string `json:"sub_aktifitas"`
	InformasiLainnya       string `json:"informasi_lainnya"`
	StatusPerbaikan        string `json:"status_perbaikan"`
	Maker                  string `json:"maker"`
	IDRiskEvent            string `json:"id_riskevent"`
	RiskEventName          string `json:"risk_event"`
	RiskIssueId            string `json:"risk_issue_id"`
	RiskIssue              string `json:"risk_issue"`
	HasilVerifikasi        string `json:"hasil_verifikasi"`
	JumlahDataDiverifikasi int64  `json:"jumlah_data_verifikasi"`
	JumlahDataAnomali      int64  `json:"jumlah_data_anomali"`
	ButuhPerbaikan         string `json:"butuh_perbaikan"`
	YangHarusDiperbaiki    int64  `json:"yang_harus_diperbaiki"`
	RtlUker                string `json:"rtl_uker"`
	StatusPerbaikanSelesai string `json:"status_perbaikan_selesai"`
	StatusPerbaikanProses  string `json:"status_perbaikan_proses"`
	PersentasePerbaikan    int64  `json:"persentase_perbaikan"`
	BatasWaktuPerbaikan    string `json:"batas_waktu_perbaikan"`
	IndikasiFraud          string `json:"indikasi_fraud"`
}

type DownloadHistoriTaskDataVerifikasiResult struct {
	ID             string `json:"id"`
	PN             string `json:"pn"`
	Nama           string `json:"nama"`
	Kanwil         string `json:"kanwil"`
	Kanca          string `json:"kanca"`
	Uker           string `json:"uker"`
	Branch         string `json:"branch"`
	Aktifitas      string `json:"aktifitas"`
	Produk         string `json:"produk"`
	Indikator      string `json:"indikator"`
	JenisTask      string `json:"jenis_task"`
	TanggalMulai   string `json:"tanggal_mulai"`
	TanggalAkhir   string `json:"tanggal_akhir"`
	StatusApproval string `json:"status_approval"`
}
