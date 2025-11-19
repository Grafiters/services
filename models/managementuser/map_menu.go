package managementuser

type MapMenu struct {
	ID         int64
	IDJabatan  int64
	IDMenu     int64
	Keterangan *string
}

func (mu MapMenu) TableName() string {
	return "management_user_map_menu"
}
