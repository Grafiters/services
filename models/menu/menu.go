package menu

import models "riskmanagement/models/type"

type MenuRequest struct {
	TipeUker string `json:"tipe_uker"`
	Hilfm    string `json:"hilfm"`
	Kostl    string `json:"kostl"`
	Pernr    string `json:"pernr"`
	Orgeh    string `json:"orgeh"`
	StellTx  string `json:"stell_tx"`
	Jgpg     string `json:"jgpg"`
}

type Menu struct {
	Id        int64   `gorm:"primaryKey" json:"id_menu"`
	Title     string  `gorm:"column:title" json:"title"`
	Icon      string  `gorm:"column:icon" json:"icon,omitempty"`
	Path      *string `gorm:"column:path" json:"path,omitempty"`
	IsSection bool    `gorm:"column:is_section" json:"is_section"`
	ParentID  int64   `gorm:"column:parent_id" json:"parent_id,omitempty"`
}

type MenuResponse struct {
	Id        int64     `gorm:"primaryKey" json:"id_menu"`
	Title     string    `gorm:"column:title" json:"title"`
	Icon      string    `gorm:"column:icon" json:"icon,omitempty"`
	Path      *string   `gorm:"column:path" json:"path,omitempty"`
	IsSection bool      `gorm:"column:is_section" json:"is_section"`
	ParentID  int64     `gorm:"column:parent_id" json:"-"`
	Submenu   []SubMenu `gorm:"foreignKey:ParentID" json:"submenu,omitempty"`
}

type SubMenu struct {
	Id           int64   `gorm:"primaryKey" json:"id_menu"`
	Title        string  `gorm:"column:title" json:"title"`
	Icon         string  `gorm:"column:icon" json:"icon,omitempty"`
	Path         *string `gorm:"column:path" json:"path,omitempty"`
	IsSection    bool    `gorm:"column:is_section" json:"is_section"`
	ParentID     int64   `gorm:"column:parent_id" json:"-"`
	SubChildMenu []Menu  `gorm:"foreignKey:ParentID" json:"subChildMenu,omitempty"`
}

// MQ Enhancement
type RequestKuisioner struct {
	Keyword  string `json:"keyword`
	TipeUker string `json:"tipe_uker"`
	Hilfm    string `json:"hilfm"`
	Kostl    string `json:"kostl"`
	Pernr    string `json:"pernr"`
	Orgeh    string `json:"orgeh"`
	StellTx  string `json:"stell_tx"`
	Jgpg     string `json:"jgpg"`
	Limit    int64  `json:"limit"`
	Offset   int64  `json:"offset"`
}

type RoleDesc struct {
	Id       string `json:"id"`
	RoleDesc string `json:"role_desc"`
}

type MenuKuisioner struct {
	Id        string  `gorm:"primaryKey" json:"id_menu"`
	Title     string  `gorm:"column:title" json:"title"`
	Icon      string  `gorm:"column:icon" json:"icon,omitempty"`
	Path      *string `gorm:"column:path" json:"path,omitempty"`
	IsSection bool    `gorm:"column:is_section" json:"is_section"`
	ParentID  int64   `gorm:"column:parent_id" json:"parent_id,omitempty"`
}

// MQ
type MenuQna struct {
	ID      int64
	Name    string
	Submenu string
}

type MenuQnaRequest struct {
	ID     int64  `json:"id"`
	Status string `json:"status"`
}

type MenuQnaResponse struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Submenu string `json:"submenu"`
}

type MstMenuResponse struct {
	Title    string `json:"title"`
	IDParent string `json:"id_parent"`
}

func (m Menu) TableName() string {
	return "menu"
}

type MstMenu struct {
	IDMenu      string `json:"id_menu"`
	Title       string `json:"title"`
	Url         string `json:"url"`
	FontIcon    string `json:"font_icon"`
	IDParent    int64  `json:"id_parent"`
	ChildStatus int64  `json:"child_status"`
	RoleAccess  string `json:"role_access"`
	Urutan      int64  `json:"urutan"`
	Status      string `json:"status"`
}

type MstMenuRequest struct {
	ID          int64
	IDMenu      string `json:"id_menu"`
	Title       string `json:"title"`
	Url         string `json:"url"`
	FontIcon    string `json:"font_icon"`
	IDParent    int64  `json:"id_parent"`
	ChildStatus int64  `json:"child_status"`
	Urutan      int64  `json:"urutan"`
}

type Role struct {
	RoleID int64  `json:"role_id"`
	MenuID string `json:"menu_id"`
}

func (mm MstMenu) TableName() string {
	return "mst_menu_questionnaire"
}

func (mm MstMenuRequest) TableName() string {
	return "mst_menu_questionnaire"
}

type MenuRoleRequest struct {
	MstMenu MstMenu       `json:"mst_menu"`
	Role    []models.Role `json:"role"`
}

type ResponseData struct {
	Status     string      `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Pagination interface{} `json:"pagination"`
}
