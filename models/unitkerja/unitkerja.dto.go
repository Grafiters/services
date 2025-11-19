package models

type UnitKerjaRequest struct {
	ID         int64   `json:"id"`
	KodeUker   int64   `json:"kode_uker"`
	NamaUker   string  `json:"nama_uker"`
	KodeCabang int64   `json:"kode_cabang"`
	NamaCabang string  `json:"nama_cabang"`
	KanwilID   int64   `json:"kanwil_id"`
	KodeKanwil string  `json:"kode_kanwil"`
	Kanwil     string  `json:"kanwil"`
	Status     int64   `json:"status"`
	CreatedAt  *string `json:"created_at"`
	UpdatedAt  *string `json:"updated_at"`
}

type UnitKerjaResponse struct {
	ID         int64   `json:"id"`
	KodeUker   int64   `json:"kode_uker"`
	NamaUker   string  `json:"nama_uker"`
	KodeCabang int64   `json:"kode_cabang"`
	NamaCabang string  `json:"nama_cabang"`
	KanwilID   int64   `json:"kanwil_id"`
	KodeKanwil string  `json:"kode_kanwil"`
	Kanwil     string  `json:"kanwil"`
	Status     int64   `json:"status"`
	CreatedAt  *string `json:"created_at"`
	UpdatedAt  *string `json:"updated_at"`
}

// batch 2
type RegionList struct {
	REGION string `json:REGION`
	BRDESC string `json:BRDESC`
}
type MainbrList struct {
	MAINBR string `json:MAINBR`
	BRDESC string `json:BRDESC`
}
type BranchList struct {
	BRANCH string `json:BRANCH`
	BRDESC string `json:BRDESC`
}
type MainbrRequest struct {
	PERNR    string `json:"pernr"`
	TipeUker string `json:"tipe_uker"`
	HILFM    string `json:"hilfm"`
	BRANCH   string `json:"BRANCH"`
	REGION   string `json:REGION`
	MAINBR   string `json:"MAINBR"`
	WERKS    string `json:"werks"`
	FlagBrc  bool   `json:"flag_Brc"`
}
type BranchRequest struct {
	PERNR    string `json:"pernr"`
	TipeUker string `json:"tipe_uker"`
	HILFM    string `json:"hilfm"`
	BRANCH   string `json:"BRANCH"`
	REGION   string `json:REGION`
	MAINBR   string `json:MAINBR`
	WERKS    string `json:"werks"`
	FlagBrc  bool   `json:"flag_Brc"`
}
type UkerName struct {
	BRDESC string `json:BRDESC`
}

// Add By Panji
type RegionRequest struct {
	PERNR    string `json:"pernr"`
	WERKS    string `json:"werks"`
	TipeUker string `json:"tipe_uker"`
	BRANCH   string `json:"BRANCH"`
}

// end of batch 2

func (p UnitKerjaRequest) ParseRequest() UnitKerja {
	return UnitKerja{
		ID:         p.ID,
		KodeUker:   p.KodeUker,
		NamaUker:   p.NamaUker,
		KodeCabang: p.KanwilID,
		NamaCabang: p.NamaCabang,
		KanwilID:   p.KanwilID,
		KodeKanwil: p.KodeKanwil,
		Kanwil:     p.Kanwil,
		Status:     p.Status,
	}
}

func (p UnitKerjaResponse) ParseResponse() UnitKerja {
	return UnitKerja{
		ID:         p.ID,
		KodeUker:   p.KodeUker,
		NamaUker:   p.NamaUker,
		KodeCabang: p.KanwilID,
		NamaCabang: p.NamaCabang,
		KanwilID:   p.KanwilID,
		KodeKanwil: p.KodeKanwil,
		Kanwil:     p.Kanwil,
		Status:     p.Status,
		CreatedAt:  p.CreatedAt,
		UpdatedAt:  p.UpdatedAt,
	}
}

type MainbrKWRequest struct {
	// Branch int64 `json:branch`
	REGION string `json:"REGION"`
}

type EmployeeRegionRequest struct {
	Branch int64 `json:"branch"`
}

type EmployeeRegionResponse struct {
	REGION string `json:"region"`
	RGDESC string `json:"rgdesc"`
	MAINBR string `json:"mainbr"`
	MBDESC string `json:"mbdesc"`
	BRANCH string `json:"branch"`
	BRDESC string `json:"brdesc"`
}

type DetailUkerResponse struct {
	REGION string `json:"REGION"`
	RGDESC string `json:"RGDESC"`
	MAINBR string `json:"MAINBR"`
	MBDESC string `json:"MBDESC"`
	BRANCH string `json:"BRANCH"`
	BRDESC string `json:"BRDESC"`
}

func (pr UnitKerjaRequest) TableName() string {
	return "unit_kerja"
}

func (pr UnitKerjaResponse) TableName() string {
	return "unit_kerja"
}

func (er EmployeeRegionResponse) TableName() string {
	return "dwh_branch"
}

type MapLocationRequest struct {
	Keyword    string `json:"keyword"`
	LevelUker  string `json:"level_uker"`
	KodeRegion string `json:"kode_region"`
	KodeBranch string `json:"kode_branch"`
	KodeUnit   string `json:"kode_unit"`
	IdPekerja  string `json:"id_pekerja"`
	FlagAll    bool   `json:"flag_all"`
}

type MapRegionOffice struct {
	Region     string `json:"region"`
	RegionCode string `json:"region_code"`
	RegionName string `json:"region_name"`
	Longitude  string `json:"longitude"`
	Latitude   string `json:"latitude"`
	Address    string `json:"address"`
}

type MapBranchOffice struct {
	Branch     string `json:"branch"`
	BranchCode string `json:"branch_code"`
	BranchName string `json:"branch_name"`
	Longitude  string `json:"longitude"`
	Latitude   string `json:"latitude"`
	Address    string `json:"address"`
}

type MapUnitOffice struct {
	Unit      string `json:"unit"`
	UnitCode  string `json:"unit_code"`
	UnitName  string `json:"unit_name"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
	Address   string `json:"address"`
}
