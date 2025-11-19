package managementuser

type MapMenuRequest struct {
	ID         int64   `json:"id"`
	IDJabatan  int64   `json:"id_jabatan"`
	IDMenu     int64   `json:"id_menu"`
	Keterangan *string `json:"keterangan"`
}

type MapMenuResponse struct {
	ID         int64   `json:"id"`
	IDJabatan  int64   `json:"id_jabatan"`
	IDMenu     int64   `json:"id_menu"`
	Keterangan *string `json:"keterangan"`
}

type MapMenuResponseFinal struct {
	ID         int64   `json:"id"`
	IDJabatan  int64   `json:"id_jabatan"`
	IDMenu     int64   `json:"id_menu"`
	Title      string  `json:"title"`
	Keterangan *string `json:"keterangan"`
}

type Menu struct {
	IDMenu      string `json:"IDMenu"`
	Title       string `json:"title"`
	Url         string `json:"url"`
	Deskripsi   string `json:"deskripsi"`
	Icon        string `json:"icon"`
	SvgIcon     string `json:"svg_icon"`
	FontIcon    string `json:"font_icon"`
	IDParent    string `json:"id_parent"`
	ChildStatus int64  `json:"child_status"`
}

type MenuRequest struct {
	Search   string `json:"search"`
	Order    string `json:"order"`
	Sort     string `json:"sort"`
	Limit    int    `json:"limit"`
	Page     int    `json:"page"`
	TIPEUKER string `json:"tipe_uker"`
	HILFM    string `json:"hilfm"`
	ORGEH    string `json:"orgeh"`
	KOSTL    string `json:"kostl"`
	PERNR    string `json:"pernr"`
	MenuRole string `json:"menu_role"`
}

type Menus []Menu

type MenuResponse struct {
	IDMenu    string              `json:"idMenu"`
	Title     string              `json:"title"`
	Url       string              `json:"url"`
	Deskripsi string              `json:"deskripsi"`
	Icon      string              `json:"icon"`
	SvgIcon   string              `json:"svgIcon"`
	FontIcon  string              `json:"fontIcon"`
	Child     []ChildMenuResponse `json:"child"`
}

type ChildMenuResponse struct {
	IDMenu   string                 `json:"idMenu"`
	Title    string                 `json:"title"`
	Url      string                 `json:"url"`
	Icon     string                 `json:"icon"`
	SvgIcon  string                 `json:"svgIcon"`
	FontIcon string                 `json:"fontIcon"`
	SubChild []SubChildMenuResponse `json:"subChild"`
}

type SubChildMenuResponse struct {
	IDMenu   int64  `json:"idMenu"`
	Title    string `json:"title"`
	Url      string `json:"url"`
	Icon     string `json:"icon"`
	SvgIcon  string `json:"svgIcon"`
	FontIcon string `json:"fontIcon"`
}

type AdditionalMenuResponse struct {
	Id   int64  `json:"id"`
	Nama string `json:"nama"`
	Url  string `json:"url"`
	Icon string `json:"icon"`
}

func (menu MapMenuRequest) TableName() string {
	return "management_user_map_menu"
}

func (menu MapMenuResponse) TableName() string {
	return "management_user_map_menu"
}

func (menu Menu) TableName() string {
	return "mst_menu"
}
