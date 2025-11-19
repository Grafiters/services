package models

type AdminSettingRole struct {
	ID        int64
	KOSTL     string `json:"kostl"`
	Orgeh     string `json:"orgeh"`
	HILFM     string `json:"hilfm"`
	TipeUker  string `json:"tipe_uker"`
	Role      string `json:"role"`
	StellTX   string `json:"stell_tx"`
	Stell     string `json:"stell"`
	JG        string `json:"jg"`
	IDSetting int64
}

type AdminSettingRoleRequest struct {
	IDSetting int64
	KOSTL     string `json:"kostl"`
	Orgeh     string `json:"orgeh"`
	HILFM     string `json:"hilfm"`
	TipeUker  string `json:"tipe_uker"`
	Role      string `json:"role"`
	StellTX   string `json:"stell_tx"`
	Stell     string `json:"stell"`
	JG        string `json:"jg"`
}

type AdminSettingRoleResponse struct {
	KOSTL    string `json:"kostl"`
	Orgeh    string `json:"orgeh"`
	HILFM    string `json:"hilfm"`
	TipeUker string `json:"tipe_uker"`
	Role     string `json:"role"`
	StellTX  string `json:"stell_tx"`
	Stell    string `json:"stell"`
	JG       string `json:"jg"`
}

func (AdminSettingRole) TableName() string {
	return "admin_setting_role"
}

func (AdminSettingRoleRequest) TableName() string {
	return "admin_setting_role"
}
