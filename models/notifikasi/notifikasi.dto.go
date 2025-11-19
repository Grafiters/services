package notifikasi

type NotifikasiResponse struct {
	ID         int    `json:"id"`
	TaskID     string `json:"task_id"`
	Tanggal    string `json:"tanggal"`
	Keterangan string `json:"keterangan"`
	Uker       string `json:"uker"`
	Status     int    `json:"status"`
	Jenis      string `json:"jenis"`
}

type NotifikasiSimpleResponse struct {
	Jenis string `json:"jenis"`
	Total string `json:"total"`
}

type NotifikasiUpdateStatus struct {
	ID     int `json:"id"`
	Status int `json:"status"`
}

type TasklistNotifikasi struct {
	ID         int    `json:"id"`
	Tanggal    string `json:"tanggal"`
	Keterangan string `json:"keterangan"`
	Status     int    `json:"status"`
	Jenis      string `json:"jenis"`
	CreatedBy  string `json:"created_by"`
	DateShow   string `json:"date_show"`
}

type NotifikasiRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Order  string `json:"order"`
	Sort   string `json:"sort"`
	Page   int    `json:"page"`
	Branch string `json:"branch"`
	PERNR  string `json:"pernr"`
}

type TasklistNotifikasiCreate struct {
	ID         int    `json:"id"`
	Tanggal    string `json:"tanggal"`
	Keterangan string `json:"keterangan"`
	Status     int    `json:"status"`
	Jenis      string `json:"jenis"`
}

type TasklistNotifikasiRequest struct {
	Keterangan string `json:"keterangan"`
	Jenis      string `json:"jenis"`
	Pernr      string `json:"pernr"`
	DateShow   string `json:"date_show"`
}

type NotifikasiTotalRequest struct {
	PERNR string `json:"pernr"`
	// Hilfm string `json:"hilfm"`
	Branch int64 `json:"branch"`
}
