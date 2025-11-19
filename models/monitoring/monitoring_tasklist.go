package models

type MonitoringTasklistRequest struct {
	PN          string `json:"pn"`
	Order       string `json:"order"`
	Sort        string `json:"sort"`
	Offset      int    `json:"offset"`
	Limit       int    `json:"limit"`
	Page        int    `json:"page"`
	JenisTask   string `json:"jenis_task"`
	JenisReport string `json:"jenis_report"`
	Periode     string `json:"periode"`
	Kanwil      string `json:"kanwil"`
}

type MonitoringTasklistUnitKerjaResponse struct {
	ID         string `json:"id"`
	Kanwil     string `json:"kanwil"`
	Kanca      string `json:"kanca"`
	KodeBranch string `json:"branch"`
	UnitKerja  string `json:"unit_kerja"`
	Pengelola  string `json:"pengelola"`
	PN         string `json:"pn"`
}

type MonitoringTasklistPekerjaResponse struct {
	ID        string `json:"id"`
	PN        string `json:"pn"`
	Nama      string `json:"nama"`
	Kanwil    string `json:"kanwil"`
	UnitKerja string `json:"unit_kerja"`
	// UnitKelolaan []string `json:"unit_kelolaan"`
}
