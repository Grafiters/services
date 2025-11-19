package models

type KriteriaPinjaman struct {
	ID       int
	Kriteria string
	Status   bool
}

type KriteriaPinjamanRequest struct {
	ID       int    `json:"id"`
	Kriteria string `json:"kriteria"`
	Status   bool   `json:"status"`
}

type KriteriaPinjamanResponse struct {
	ID       int    `json:"id"`
	Kriteria string `json:"kriteria"`
	Status   bool   `json:"status"`
}

type Paginate struct {
	Order  string `json:"order"`
	Sort   string `json:"sort"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Page   int    `json:"page"`
}

func (KP KriteriaPinjaman) TableName() string {
	return "tbl_mst_kriteria_pinjaman"
}

func (KP KriteriaPinjamanRequest) TableName() string {
	return "tbl_mst_kriteria_pinjaman"
}

func (KP KriteriaPinjamanResponse) TableName() string {
	return "tbl_mst_kriteria_pinjaman"
}
