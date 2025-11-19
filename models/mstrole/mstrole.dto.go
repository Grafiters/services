package mstrole

type MstRoleRequest struct {
	ID             int64  `json:"id"`
	Menu           string `json:"menu"`
	AdditionalMenu string `json:"additional_menu"`
	RoleName       string `json:"role_name"`
	AddonPernr     bool   `json:"addon_pernr"`
	DeleteFlag     bool   `json:"delete_flag"`
	// Menu       []MstRoleMapMenuRequest `json:"menu"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

type Paginate struct {
	Order  string `json:"order"`
	Sort   string `json:"sort"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Page   int    `json:"page"`
}

type MstRoleResponse struct {
	ID             int64   `json:"id"`
	Menu           string  `json:"menu"`
	AdditionalMenu string  `json:"additional_menu"`
	RoleName       string  `json:"role_name"`
	AddonPernr     bool    `json:"addon_pernr"`
	DeleteFlag     bool    `json:"delete_flag"`
	CreatedAt      *string `json:"created_at"`
	UpdatedAt      *string `json:"updated_at"`
}

type MstRoleResponseOne struct {
	ID             int64   `json:"id"`
	RoleName       string  `json:"role_name"`
	AdditionalMenu string  `json:"additional_menu"`
	MfeMenu        string  `json:"mfe_menu"`
	AddonPernr     bool    `json:"addon_pernr"`
	DeleteFlag     bool    `json:"delete_flag"`
	Menu           string  `json:"menu"`
	CreatedAt      *string `json:"created_at"`
	UpdatedAt      *string `json:"updated_at"`
}

type MenuListRequest struct {
	TIPEUKER string `json:"tipe_uker"`
	HILFM    string `json:"hilfm"`
	ORGEH    string `json:"orgeh"`
	KOSTL    string `json:"kostl"`
	PERNR    string `json:"pernr"`
	StellTx  string `json:"stell_tx"`
	Jgpg     string `json:"jgpg"`
}

type MstRoleRequestUpdate struct {
	ID             int64   `json:"id"`
	RoleName       string  `json:"role_name"`
	AddonPernr     bool    `json:"addon_pernr"`
	Menu           string  `json:"menu"`
	AdditionalMenu string  `json:"additional_menu"`
	UpdatedAt      *string `json:"updated_at"`
}

type MstRoleRequestDelete struct {
	ID         int64   `json:"id"`
	DeleteFlag bool    `json:"delete_flag"`
	UpdatedAt  *string `json:"updated_at"`
}

// type MstRoleQuestionnaireResponseOne struct {
// 	RoleID     int64   `json:"role_id"`
// 	DeleteFlag bool    `json:"delete_flag"`
// 	MenuID     string  `json:"menu_id"`
// 	CreatedAt  *string `json:"created_at"`
// 	UpdatedAt  *string `json:"updated_at"`
// }

type MstRoleQuestionnaireResponseOne struct {
	// ID         int64  `json:"id"`
	ID         string `json:"id"`
	RoleName   string `json:"role_name"`
	DeleteFlag bool   `json:"delete_flag"`
	Menu       string `json:"menu"`
	// Menu       []MstRoleMapMenuResponseOne `json:"menu"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

func (mu MstRoleRequestUpdate) TableName() string {
	return "mst_roles_batch3"
}

func (mu MstRoleRequestDelete) TableName() string {
	return "mst_roles_batch3"
}

func (mu MstRoleRequest) TableName() string {
	return "mst_roles_batch3"
}

func (mu MstRoleResponse) TableName() string {
	return "mst_roles_batch3"
}

func (mu MstRoleResponseOne) TableName() string {
	return "mst_roles_batch3"
}

func (mu MstRoleQuestionnaireResponseOne) TableName() string {
	return "mst_roles_questionnaire"
}
