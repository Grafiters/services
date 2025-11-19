package models

type CommonRequest struct {
	ID        int64
	MenuID    int    `json:"menu_id"`
	SubMenu   int64  `json:"sub_menu_id"`
	TypeID    int64  `json:"type_id"`
	PartID    int64  `json:"part_id"`
	SubPartID int64  `json:"sub_part_id"`
	Keyword   string `json:"keyword"`
	Status    string `json:"status"`
	PERNR     string `json:"pernr"`
}
