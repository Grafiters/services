package mstrole

type MstRoleMapMenu struct {
	ID     int64
	IDRole int64
	IDMenu int64
}

func (mst MstRoleMapMenu) TableName() string {
	return "mst_role_map_menu"
}
