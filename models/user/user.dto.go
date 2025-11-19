package user

type LoginRequest struct {
	Pernr    string `json:"pernr"`
	Password string `json:"password"`
}

type Login struct {
	ClientID     string `json:"clientid"`
	ClientSecret string `json:"clientsecret"`
}

type MenuRequest struct {
	Search    string `json:"search"`
	Order     string `json:"order"`
	Sort      string `json:"sort"`
	Limit     int    `json:"limit"`
	Page      int    `json:"page"`
	LevelUker string `json:"levelUker"`
	LevelID   string `json:"levelId"`
	Orgeh     string `json:"orgeh"`
	Kostl     string `json:"kostl"`
}

type UserSession struct {
	PERNR        string       `json:"PERNR"`
	NIP          string       `json:"NIP"`
	SNAME        string       `json:"SNAME"`
	CORP_TITLE   string       `json:"CORP_TITLE"`
	JGPG         string       `json:"JGPG"`
	AGAMA        string       `json:"AGAMA"`
	WERKS        string       `json:"WERKS"`
	BTRTL        string       `json:"BTRTL"`
	KOSTL        string       `json:"KOSTL"`
	ORGEH        string       `json:"ORGEH"`
	ORGEH_PGS    string       `json:"ORGEH_PGS"`
	TIPE_UKER    string       `json:"TIPE_UKER"`
	STELL        string       `json:"STELL"`
	WERKS_TX     string       `json:"WERKS_TX"`
	BTRTL_TX     string       `json:"BTRTL_TX"`
	KOSTL_TX     string       `json:"KOSTL_TX"`
	ORGEH_TX     string       `json:"ORGEH_TX"`
	ORGEH_PGS_TX string       `json:"ORGEH_PGS_TX"`
	PLANS_PGS    string       `json:"PLANS_PGS"`
	PLANS_PGS_TX string       `json:"PLANS_PGS_TX"`
	STELL_TX     string       `json:"STELL_TX"`
	PLANS_TX     string       `json:"PLANS_TX"`
	BRANCH       string       `json:"BRANCH"`
	HILFM        string       `json:"HILFM"`
	HILFM_PGS    string       `json:"HILFM_PGS"`
	HTEXT        string       `json:"HTEXT"`
	HTEXT_PGS    string       `json:"HTEXT_PGS"`
	ADD_AREA     string       `json:"ADD_AREA"`
	SISA_CT      int64        `json:"SISA_CT"`
	SISA_CB      int64        `json:"SISA_CB"`
	KAWIN        string       `json:"KAWIN"`
	STATUS       string       `json:"STATUS"`
	LAST_SYNC    string       `json:"LAST_SYNC"`
	MENU         MenuResponse `json:"MENU"`
}

type UserSessionIncognito struct {
	PERNR      string       `json:"PERNR"`
	WERKS      string       `json:"WERKS"`
	BTRTL      string       `json:"BTRTL"`
	KOSTL      string       `json:"KOSTL"`
	ORGEH      string       `json:"ORGEH"`
	ORGEHPGS   string       `json:"ORGEH_PGS"`
	STELL      string       `json:"STELL"`
	SNAME      string       `json:"SNAME"`
	WERKSTX    string       `json:"WERKS_TX"`
	BTRTLTX    string       `json:"BTRTL_TX"`
	KOSTLTX    string       `json:"KOSTL_TX"`
	ORGEHTX    string       `json:"ORGEH_TX"`
	ORGEHPGSTX string       `json:"ORGEH_PGS_TX"`
	STELLTX    string       `json:"STELL_TX"`
	BRANCH     string       `json:"BRANCH"`
	TIPEUKER   string       `json:"TIPE_UKER"`
	HILFM      string       `json:"HILFM"`
	HILFMPGS   string       `json:"HILFM_PGS"`
	HTEXT      string       `json:"HTEXT"`
	HTEXTPGS   string       `json:"HTEXT_PGS"`
	CORPTITLE  string       `json:"CORP_TITLE"`
	MENU       MenuResponse `json:"MENU"`
}

type Menu struct {
	IDMenu    int64  `json:"IDMenu"`
	Title     string `json:"title"`
	Url       string `json:"url"`
	Deskripsi string `json:"deskripsi"`
	Icon      string `json:"icon"`
	SvgIcon   string `json:"svgIcon"`
	FontIcon  string `json:"fontIcon"`
}

type Menus []Menu

type MenuResponse struct {
	IDMenu    int64               `json:"idMenu"`
	Title     string              `json:"title"`
	Url       string              `json:"url"`
	Deskripsi string              `json:"deskripsi"`
	Icon      string              `json:"icon"`
	SvgIcon   string              `json:"svgIcon"`
	FontIcon  string              `json:"fontIcon"`
	Child     []ChildMenuResponse `json:"child"`
}

type ChildMenuResponse struct {
	IDMenu   int64                  `json:"idMenu"`
	Title    string                 `json:"title"`
	Url      string                 `json:"url"`
	Icon     string                 `json:"icon"`
	SvgIcon  string                 `json:"svgIcon"`
	FontIcon string                 `json:"fontIcon"`
	SubChild []SubChildMenuResponse `json:"subChild"`
}

type SubChildMenuResponse struct {
	Title    string `json:"title"`
	Url      string `json:"url"`
	Icon     string `json:"icon"`
	SvgIcon  string `json:"svgIcon"`
	FontIcon string `json:"fontIcon"`
}
