package models

type UkerKelolaanRequest struct {
	ID        int64      `json:"id"`
	CreatedAt *string    `json:"created_at"`
	UpdatedAt *string    `json:"updated_at"`
	ExpiredAt string     `json:"expired_at"`
	IsTemp    bool       `json:"is_temp"`
	Pn        string     `json:"pn"`
	SNAME     string     `json:"SNAME"`
	ListUker  []ListUker `json:"list_uker"`
	Status    bool       `json:"status"`
}

type SaveMstRequest struct {
	Id    int64  `json:"id"`
	Pn    string `json:"pn"`
	Sname string `json:"sname"`
	Aktif bool   `json:"aktif"`
}

type ListUker struct {
	ID     int64  `json:"id"`
	REGION string `json:"REGION"`
	RGDESC string `json:"RGDESC"`
	MAINBR string `json:"MAINBR"`
	MBDESC string `json:"MBDESC"`
	BRANCH string `json:"BRANCH"`
	BRDESC string `json:"BRDESC"`
}

type UkerKelolaanResponseOne struct {
	ID        int64      `json:"id"`
	CreatedAt *string    `json:"created_at"`
	UpdatedAt *string    `json:"updated_at"`
	ExpiredAt string     `json:"expired_at"`
	IsTemp    bool       `json:"is_temp"`
	Pn        string     `json:"pn"`
	SNAME     string     `json:"SNAME"`
	ListUker  []ListUker `json:"list_uker"`
	Status    bool       `json:"status"`
}

type UkerKelolaanResponse struct {
	ID        int64   `json:"id"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
	ExpiredAt string  `json:"expired_at"`
	IsTemp    bool    `json:"is_temp"`
	Pn        string  `json:"pn"`
	SNAME     string  `json:"SNAME"`
	REGION    string  `json:"REGION"`
	RGDESC    string  `json:"RGDESC"`
	MAINBR    string  `json:"MAINBR"`
	MBDESC    string  `json:"MBDESC"`
	BRANCH    string  `json:"BRANCH"`
	BRDESC    string  `json:"BRDESC"`
	Status    bool    `json:"status"`
}

type KeywordRequest struct {
	Order  string `json:"order"`
	Sort   string `json:"sort"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Page   int    `json:"page"`
	Pn     string `json:"pn"`
	REGION string `json:"REGION"`
	MAINBR string `json:"MAINBR"`
	BRANCH string `json:"BRANCH"`
	Status bool   `json:"status"`
}

type Paginate struct {
	Order  string `json:"order"`
	Sort   string `json:"sort"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Page   int    `json:"page"`
}

type RequestOne struct {
	Id int64  `json:"id"`
	Pn string `json:"pn"`
}

type PencarianUker struct {
	Keyword  string `json:"keyword"`
	PERNR    string `json:"PERNR"`
	TIPEUKER string `json:"TIPE_UKER"`
	BRANCH   string `json:"BRANCH"`
	HILFM    string `json:"HILFM"`
	Limit    int64  `json:"limit"`
	Offset   int64  `json:"offset"`
}

type UkerList struct {
	REGION string `json:"REGION"`
	RGDESC string `json:"RGDESC"`
	MAINBR string `json:"MAINBR"`
	MBDESC string `json:"MBDESC"`
	BRANCH string `json:"BRANCH"`
	BRDESC string `json:"BRDESC"`
}

func (uk UkerKelolaanRequest) TableName() string {
	return "uker_kelolaan_user"
}

func (uk UkerKelolaanResponse) TableName() string {
	return "uker_kelolaan_user"
}

func (uk SaveMstRequest) TableName() string {
	return "mst_uker_kelolaan"
}
