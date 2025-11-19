package mstrole

type MstRoleMapMenuRequest struct {
	ID     int64 `json:"id"`
	IDRole int64 `json:"id_role"`
	IDMenu int64 `json:"id_menu"`
}

type MstRoleMapMenuResponse struct {
	ID     int64 `json:"id"`
	IDRole int64 `json:"id_role"`
	IDMenu int64 `json:"id_menu"`
}

type MstRoleMapMenuResponseOne struct {
	ID     int64 `json:"id"`
	IDRole int64 `json:"id_role"`
	IDMenu int64 `json:"id_menu"`
	// Title  string `json:"Title"`
}

func (p MstRoleMapMenuRequest) ParseRequest() MstRoleMapMenu {
	return MstRoleMapMenu{
		ID:     p.ID,
		IDRole: p.IDRole,
		IDMenu: p.IDMenu,
	}
}

func (p MstRoleMapMenuResponse) ParseResponse() MstRoleMapMenu {
	return MstRoleMapMenu{
		ID:     p.ID,
		IDRole: p.IDRole,
		IDMenu: p.IDMenu,
	}
}

func (mu MstRoleMapMenuRequest) TableName() string {
	return "mst_role_map_menu"
}

func (mu MstRoleMapMenuResponse) TableName() string {
	return "mst_role_map_menu"
}
