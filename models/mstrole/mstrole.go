package mstrole

type MstRole struct {
	ID             int64
	RoleName       string
	AddonPernr     bool
	Menu           string
	AdditionalMenu string
	DeleteFlag     bool
	CreatedAt      *string
	UpdatedAt      *string
}

func (mu MstRole) TableName() string {
	return "mst_roles_batch3"
}
