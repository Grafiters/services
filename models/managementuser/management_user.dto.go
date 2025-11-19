package managementuser

import (
	"riskmanagement/lib"
)

type ManagementUserRequest struct {
	ID         int64   `json:"id"`
	RoleID     int64   `json:"role_id"`
	LevelUker  string  `json:"level_uker"`
	LevelID    int64   `json:"level_id"`
	AddonPernr string  `json:"addon_pernr"`
	CreatedAt  *string `json:"created_at"`
	UpdatedAt  *string `json:"updated_at"`
}

type Paginate struct {
	Order  string `json:"order"`
	Sort   string `json:"sort"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Page   int    `json:"page"`
	Role   int    `json:"role"`
}

type ManagementUserResponse struct {
	ID         int64   `json:"id"`
	RoleID     int64   `json:"role_id"`
	LevelUker  string  `json:"level_uker"`
	LevelID    int64   `json:"level_id"`
	AddonPernr string  `json:"addon_pernr"`
	CreatedAt  *string `json:"created_at"`
	UpdatedAt  *string `json:"updated_at"`
}

type ManagementUserFinResponse struct {
	ID int64 `json:"id"`
	// NamaJabatan string  `json:"nama_jabatan"`
	Role      string `json:"role"`
	LevelUker string `json:"level_uker"`
	StellTx   string `json:"stell_tx"`
	Jgpg      string `json:"jgpg"`
}

type ManagementUserResponses struct {
	ID        int64  `json:"id"`
	RoleID    int64  `json:"role_id"`
	LevelUker string `json:"level_uker"`
	LevelID   int64  `json:"level_id"`
	// ORGEH     string  `json:"orgeh"`
	MapMenu []MapMenuResponseFinal `json:"map_menu"`
}

type ManagementUserCheckAccessibilityRequest struct {
	LevelUker string `json:"level_uker"`
	LevelID   string `json:"level_id"`
	ORGEH     string `json:"orgeh"`
	PERNR     string `json:"pernr"`
}

type UkerKelolaanUserRequest struct {
	PERNR string `json:"PERNR"`
}

type UkerKelolaanUserResponse struct {
	Id        int64  `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	ExpiredAt string `json:"expired_at"`
	IsTemp    int64  `json:"is_temp"`
	PN        string `json:"pn"`
	REGION    string `json:"REGION"`
	RGDESC    string `json:"RGDESC"`
	MAINBR    string `json:"MAINBR"`
	MBDESC    string `json:"MBDESC"`
	BRANCH    string `json:"BRANCH"`
	BRDESC    string `json:"BRDESC"`
}

type UkerKelolaanUserResponseNull struct {
	Id        lib.NullInt64  `json:"id"`
	CreatedAt lib.NullString `json:"created_at"`
	UpdatedAt lib.NullString `json:"updated_at"`
	ExpiredAt lib.NullString `json:"expired_at"`
	IsTemp    lib.NullInt64  `json:"is_temp"`
	PN        lib.NullString `json:"pn"`
	REGION    lib.NullString `json:"REGION"`
	RGDESC    lib.NullString `json:"RGDESC"`
	MAINBR    lib.NullString `json:"MAINBR"`
	MBDESC    lib.NullString `json:"MBDESC"`
	BRANCH    lib.NullString `json:"BRANCH"`
	BRDESC    lib.NullString `json:"BRDESC"`
}

// levelUker
type LevelUkerResponse struct {
	Id        int    `json:"id"`
	LevelUker string `json:"level_uker"`
	Deskripsi string `json:"deskripsi"`
}

type MappingMenuRequest struct {
	ID      int64            `json:"id"`
	MapMenu []MapMenuRequest `json:"map_menu"`
}

// Enhance Management User By Panji 02/02/2024
type JabatanRolesResponse struct {
	Id          int    `json:"id"`
	HILFM       string `json:"hilfm"`
	Description string `json:"description"`
	Jgpg        string `json:"jgpg"`
}

// 06/02/2023
type MenuListRequest struct {
	TIPEUKER string `json:"tipe_uker"`
	HILFM    string `json:"hilfm"`
	KOSTL    string `json:"KOSTL"`
	PERNR    string `json:"pernr"`
	ORGEH    string `json:"orgeh"`
	StellTx  string `json:"stell_tx"`
	Jgpg     string `json:"jgpg"`
}

func (mu ManagementUserRequest) TableName() string {
	return "management_user_batch3"
}

func (mu ManagementUserResponse) TableName() string {
	return "management_user_batch3"
}

func (lv LevelUkerResponse) TableName() string {
	return "level_uker"
}

// Enhance Management User By Panji 02/02/2024
func (jr JabatanRolesResponse) TableName() string {
	return "mst_jabatan"
}
