package models

type PerhitunganPersentasePenyelesaianPagianted struct {
	Order          string `json:"order"`
	Sort           string `json:"sort"`
	Offset         int    `json:"offset"`
	Limit          int    `json:"limit"`
	Page           int    `json:"page"`
	JenisReport    string `json:"jenis_report"`
	NoTasklist     string `json:"no_tasklist"`
	NamaTasklist   string `json:"nama_tasklist"`
	Nama           string `json:"namapn"`
	Kanwil         string `json:"kanwil"`
	Kanca          string `json:"kanca"`
	Uker           string `json:"uker"`
	Aktifitas      string `json:"activity_id"`
	Produk         string `json:"product_id"`
	Indikator      string `json:"risk_indicator_id"`
	IndikatorOther string `json:"risk_indicator_other"`
	Status         string `json:"status"`
	JenisTaks      string `json:"jenis_task"`
	RiskIssue      string `json:"risk_issue_id"`
	Kegiatan       string `json:"kegiatan"`
}

type PerhitunganPersentasePenyelesaianDownload struct {
	Order          string `json:"order"`
	Sort           string `json:"sort"`
	JenisReport    string `json:"jenis_report"`
	Nama           string `json:"namapn"`
	NoTasklist     string `json:"no_tasklist"`
	NamaTasklist   string `json:"nama_tasklist"`
	Kanwil         string `json:"kanwil"`
	Kanca          string `json:"kanca"`
	Uker           string `json:"uker"`
	Aktifitas      string `json:"activity_id"`
	Produk         string `json:"product_id"`
	Indikator      string `json:"risk_indicator_id"`
	IndikatorOther string `json:"risk_indicator_other"`
	Status         string `json:"status"`
	JenisTaks      string `json:"jenis_task"`
	RiskIssue      string `json:"risk_issue_id"`
	Kegiatan       string `json:"kegiatan"`
}

type LaporanPerPekerjaQueryResult struct {
	ID                          string `json:"id"`
	Pn                          string `json:"pn"`
	Nama                        string `json:"nama"`
	JenisTask                   string `json:"jenis_task"`
	Aktifitas                   string `json:"aktifitas"`
	Produk                      string `json:"produk"`
	Indikator                   string `json:"indikator"`
	TanggalMulai                string `json:"tanggal_mulai"`
	TanggalSelesai              string `json:"tanggal_selesai"`
	JumlahDataAnomali           int64  `json:"jumlah_data_anomali"`
	JumlahDataVerifikasi        int64  `json:"jumlah_data_sudah_verifikasi"`
	JumlahDataPerluTindaklanjut int64  `json:"jumlah_data_perlu_tindaklanjut"`
	JumlahDataSudahTindaklanjut int64  `json:"jumlah_data_sudah_tindaklanjut"`
}

type LaporanPerPekerjaResult struct {
	ID                          string  `json:"id"`
	Pn                          string  `json:"pn"`
	Nama                        string  `json:"nama"`
	JenisTask                   string  `json:"jenis_task"`
	Aktifitas                   string  `json:"aktifitas"`
	Produk                      string  `json:"produk"`
	Indikator                   string  `json:"indikator"`
	TanggalMulai                string  `json:"tanggal_mulai"`
	TanggalSelesai              string  `json:"tanggal_selesai"`
	JumlahDataAnomali           int64   `json:"jumlah_data_anomali"`
	JumlahDataVerifikasi        int64   `json:"jumlah_data_sudah_verifikasi"`
	JumlahDataPerluTindaklanjut int64   `json:"jumlah_data_perlu_tindaklanjut"`
	JumlahDataSudahTindaklanjut int64   `json:"jumlah_data_sudah_tindaklanjut"`
	PersenSudahVerifikasi       float64 `json:"persen_sudah_verifikasi"`
	PersenSudahTindaklanjut     float64 `json:"persen_sudah_tindaklanjut"`
}

type LaporanPerUkerResult struct {
	ID                          string  `json:"id"`
	Kanwil                      string  `json:"kanwil"`
	Kanca                       string  `json:"kanca"`
	Uker                        string  `json:"uker"`
	JenisTask                   string  `json:"jenis_task"`
	Aktifitas                   string  `json:"aktifitas"`
	Produk                      string  `json:"produk"`
	Indikator                   string  `json:"indikator"`
	TanggalMulai                string  `json:"tanggal_mulai"`
	TanggalSelesai              string  `json:"tanggal_selesai"`
	JumlahDataAnomali           int64   `json:"jumlah_data_anomali"`
	JumlahDataVerifikasi        int64   `json:"jumlah_data_verifikasi"`
	JumlahDataPerluTidaklanjut  int64   `json:"jumlah_data_perlu_tidaklanjut"`
	JumlahDataSudahTindaklanjut int64   `json:"jumlah_data_sudah_tindaklanjut"`
	PersenSudahVerifikasi       float64 `json:"persen_sudah_verifikasi"`
	PersenSudahTindaklanjut     float64 `json:"persen_sudah_tindaklanjut"`
}

type LaporanPerPekerjaUkerResult struct {
	ID                         string  `json:"id"`
	Pn                         string  `json:"pn"`
	Nama                       string  `json:"nama"`
	Kanwil                     string  `json:"kanwil"`
	Kanca                      string  `json:"kanca"`
	Uker                       string  `json:"uker"`
	JenisTask                  string  `json:"jenis_task"`
	Aktifitas                  string  `json:"aktifitas"`
	Produk                     string  `json:"produk"`
	Indikator                  string  `json:"indikator"`
	TanggalMulai               string  `json:"tanggal_mulai"`
	TanggalSelesai             string  `json:"tanggal_selesai"`
	JumlahDataAnomali          int64   `json:"jumlah_data_anomali"`
	JumlahDataVerifikasi       int64   `json:"jumlah_data_verifikasi"`
	JumlahDataPerluTidaklanjut int64   `json:"jumlah_data_perlu_tidaklanjut"`
	PersenSudahVerifikasi      float64 `json:"persen_sudah_verifikasi"`
	PersenSudahTindaklanjut    float64 `json:"persen_sudah_tindaklanjut"`
}

type RiskEventOnTaskList struct {
	RiskIssueId string `json:"risk_issue_id"`
	RiskIssue   string `json:"risk_issue"`
}

type LaporanVerifikasiResponse struct {
	ID                          string  `json:"id"`
	Pn                          string  `json:"pn"`
	Nama                        string  `json:"nama"`
	NoTasklist                  string  `json:"no_tasklist"`
	NamaTasklist                string  `json:"nama_tasklist"`
	Kanwil                      string  `json:"kanwil"`
	Kanca                       string  `json:"kanca"`
	Uker                        string  `json:"uker"`
	JenisTask                   string  `json:"jenis_task"`
	Kegiatan                    string  `json:"kegiatan"`
	Aktifitas                   string  `json:"aktifitas"`
	Produk                      string  `json:"produk"`
	RiskIssue                   string  `json:"risk_issue"`
	Indikator                   string  `json:"indikator"`
	TanggalMulai                string  `json:"tanggal_mulai"`
	TanggalSelesai              string  `json:"tanggal_selesai"`
	JumlahDataAnomali           int64   `json:"jumlah_data_anomali"`
	JumlahDataVerifikasi        int64   `json:"jumlah_data_verifikasi"`
	JumlahDataPerluTindaklanjut int64   `json:"jumlah_data_perlu_tindaklanjut"`
	JumlahDataSudahTindaklanjut int64   `json:"jumlah_data_sudah_tindaklanjut"`
	PersenSudahVerifikasi       float64 `json:"persen_verifikasi"`
	PersenSudahTindaklanjut     float64 `json:"persen_sudah_tindaklanjut"`
	JumlahKegiatanDilakukan     string  `json:"jumlah_kegiatan_dilakukan"`
	ButuhPerbaikan              string  `json:"butuh_perbaikan"`
}

type LaporanBriefingResponse struct {
	ID                              string `json:"id"`
	Pn                              string `json:"pn"`
	Nama                            string `json:"nama"`
	Kanwil                          string `json:"kanwil"`
	Kanca                           string `json:"kanca"`
	Uker                            string `json:"uker"`
	JenisTask                       string `json:"jenis_task"`
	Kegiatan                        string `json:"kegiatan"`
	Aktifitas                       string `json:"aktifitas"`
	Produk                          string `json:"produk"`
	RiskIssue                       string `json:"risk_issue"`
	Indikator                       string `json:"indikator"`
	TanggalMulai                    string `json:"tanggal_mulai"`
	TanggalSelesai                  string `json:"tanggal_selesai"`
	JumlahDataAnomali               string `json:"jumlah_data_anomali"`
	JumlahDataVerifikasi            string `json:"jumlah_data_verifikasi"`
	JumlahDataPerluTindaklanjut     string `json:"jumlah_data_perlu_tindaklanjut"`
	JumlahDataYangSudahTindaklanjut string `json:"jumlah_data_sudah_tindaklanjut"`
	PersenSudahVerifikasi           string `json:"persen_verifikasi"`
	PersenSudahTindaklanjut         string `json:"persen_sudah_tindaklanjut"`
	JumlahKegiatanDilakukan         string `json:"jumlah_kegiatan_dilakukan"`
}

type LaporanCoachingResponse struct {
	ID                              string `json:"id"`
	Pn                              string `json:"pn"`
	Nama                            string `json:"nama"`
	Kanwil                          string `json:"kanwil"`
	Kanca                           string `json:"kanca"`
	Uker                            string `json:"uker"`
	JenisTask                       string `json:"jenis_task"`
	Kegiatan                        string `json:"kegiatan"`
	Aktifitas                       string `json:"aktifitas"`
	Produk                          string `json:"produk"`
	RiskIssue                       string `json:"risk_issue"`
	Indikator                       string `json:"indikator"`
	TanggalMulai                    string `json:"tanggal_mulai"`
	TanggalSelesai                  string `json:"tanggal_selesai"`
	JumlahDataAnomali               string `json:"jumlah_data_anomali"`
	JumlahDataVerifikasi            string `json:"jumlah_data_verifikasi"`
	JumlahDataPerluTindaklanjut     string `json:"jumlah_data_perlu_tindaklanjut"`
	JumlahDataYangSudahTindaklanjut string `json:"jumlah_data_sudah_tindaklanjut"`
	PersenSudahVerifikasi           string `json:"persen_verifikasi"`
	PersenSudahTindaklanjut         string `json:"persen_sudah_tindaklanjut"`
	JumlahKegiatanDilakukan         string `json:"jumlah_kegiatan_dilakukan"`
}
