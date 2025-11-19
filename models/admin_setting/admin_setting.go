package models

import "riskmanagement/lib"

type AdminSetting struct {
	ID          int64
	TaskType    string  `json:"task_type"`
	Kegiatan    string  `json:"kegiatan"`
	Period      string  `json:"period"`
	Range       string  `json:"range"`
	Upload      string  `json:"upload"`
	TasklistMax int64   `json:"tasklist_max"`
	Status      string  `json:"status"`
	CreatedAt   *string `json:"created_at"`
	UpdatedAt   *string `json:"updated_at"`
}

type AdminSettingRequest struct {
	ID          int64
	TaskType    string                    `json:"task_type"`
	Role        []AdminSettingRoleRequest `json:"role"`
	Kegiatan    string                    `json:"kegiatan"`
	Period      string                    `json:"period"`
	Range       string                    `json:"range"`
	Upload      string                    `json:"upload"`
	TasklistMax int64                     `json:"tasklist_max"`
	Status      string                    `json:"status"`
	CreatedAt   *string                   `json:"created_at"`
	UpdatedAt   *string                   `json:"updated_at"`
}

type AdminSettingUpdate struct {
	ID          int64  `json:"id"`
	TaskType    string `json:"task_type"`
	Kegiatan    string `json:"kegiatan"`
	Period      string `json:"period"`
	Range       string `json:"range"`
	Upload      string `json:"upload"`
	TasklistMax int64  `json:"tasklist_max"`
	UpdatedAt   *string
}

type AdminSettingUpdateRequest struct {
	ID               int64              `json:"id"`
	TaskType         string             `json:"task_type"`
	AdminSettingRole []AdminSettingRole `json:"role"`
	Kegiatan         string             `json:"kegiatan"`
	Period           string             `json:"period"`
	Range            string             `json:"range"`
	Upload           string             `json:"upload"`
	TasklistMax      int64              `json:"tasklist_max"`
	UpdatedAt        *string
}
type AdminSettingDelete struct {
	ID        int64 `json:"id"`
	Status    string
	UpdatedAt *string
}

type TaskTypeRequestOne struct {
	ID int64 `json:"id"`
}

type AdminSettingResponse struct {
	ID          int64                  `json:"id"`
	TaskType    string                 `json:"task_type"`
	Role        []TaskTypeRoleResponse `json:"role"`
	Kegiatan    string                 `json:"kegiatan"`
	Period      string                 `json:"period"`
	Range       string                 `json:"range"`
	Upload      string                 `json:"upload"`
	TasklistMax int64                  `json:"tasklist_max"`
}

type KeywordRequest struct {
	Order    string `json:"order"`
	Sort     string `json:"sort"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
	Page     int    `json:"page"`
	Keyword  string `json:"keyword"`
	TipeUker string `json:"tipe_uker"`
	HILFM    string `json:"hilfm"`
	KOSTL    string `json:"kostl"`
	Orgeh    string `json:"orgeh"`
	StellTX  string `json:"stell_tx"`
	Stell    string `json:"stell"`
	JG       string `json:"jg"`
	Kegiatan string `json:"kegiatan"`
}

type TaskType struct {
	ID          int64  `json:"id"`
	TaskType    string `json:"task_type"`
	Kegiatan    string `json:"kegiatan"`
	Period      string `json:"period"`
	Range       string `json:"range"`
	Upload      string `json:"upload"`
	TasklistMax int64  `json:"tasklist_max"`
}

type TaskTypeResponse struct {
	ID          int64                  `json:"id"`
	TaskType    string                 `json:"task_type"`
	Kegiatan    string                 `json:"kegiatan"`
	Period      string                 `json:"period"`
	Range       string                 `json:"range"`
	Upload      string                 `json:"upload"`
	TasklistMax int64                  `json:"tasklist_max"`
	Role        []TaskTypeRoleResponse `json:"role"`
}

type TaskTypeRoleResponse struct {
	ID        int64  `json:"id"`
	IDSetting int64  `json:"id_setting"`
	KOSTL     string `json:"kostl"`
	Orgeh     string `json:"orgeh"`
	HILFM     string `json:"hilfm"`
	TipeUker  string `json:"tipe_uker"`
	Role      string `json:"role"`
	StellTX   string `json:"stell_tx"`
	Stell     string `json:"stell"`
	JG        string `json:"jg"`
}

type Paginate struct {
	Order    string `json:"order"`
	Sort     string `json:"sort"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
	Page     int    `json:"page"`
	TipeUker string `json:"tipe_uker"`
	HILFM    string `json:"hilfm"`
	KOSTL    string `json:"kostl"`
	Orgeh    string `json:"orgeh"`
	StellTX  string `json:"stell_tx"`
	Stell    string `json:"stell"`
	JG       string `json:"jg"`
}

type TaskTypeResponsesNull struct {
	ID          lib.NullInt64  `json:"id"`
	TaskType    lib.NullString `json:"task_type"`
	Kegiatan    lib.NullString `json:"kegiatan"`
	Period      lib.NullString `json:"period"`
	Range       lib.NullString `json:"range"`
	Upload      lib.NullString `json:"upload"`
	TasklistMax lib.NullString `json:"tasklist_max"`
}

type TaskTypeCheckRequest struct {
	TaskType string `json:"task_type"`
}

type TaskTypeCheckResponse struct {
	Total int64 `json:"total"`
}

func (AdminSetting) TableName() string {
	return "admin_setting"
}

func (AdminSettingUpdate) TableName() string {
	return "admin_setting"
}

func (AdminSettingDelete) TableName() string {
	return "admin_setting"
}
