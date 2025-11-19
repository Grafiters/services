package common

type PnNamaRequest struct {
	HILFM string `json:"hilfm"`
	ORGEH string `json:"orgeh"`
	PERN  string `json:"pern"`
}

type KeywordRequest struct {
	Order   string `json:"order"`
	Sort    string `json:"sort"`
	Offset  int    `json:"offset"`
	Limit   int    `json:"limit"`
	Page    int    `json:"page"`
	Keyword string `json:"keyword"`
	HILFM   string `json:"hilfm"`
	ORGEH   string `json:"orgeh"`
	PERN    string `json:"pernr"`
}

type PnNamaResult struct {
	Pernr string `json:"pernr"`
	Sname string `json:"sname"`
}

type ORDMember struct {
	PERNR string `json:"pernr"`
	SNAME string `json:"sname"`
}

type KanwilResult struct {
	Nama string `json:"nama"`
}

type KancaResult struct {
	Nama string `json:"nama"`
}

type UkerResult struct {
	Kode string `json:"kode"`
	Nama string `json:"nama"`
}

type RiskEventRequest struct {
	ActivityID string `json:"activity_id"`
	ProductID  string `json:"product_id"`
}

type RiskEventResult struct {
	ID            int    `json:"id"`
	RiskIssueCode string `json:"risk_issue_code"`
	RiskIssue     string `json:"risk_issue"`
}

type RiskIndikatorRequest struct {
	RiskEventID string `json:"riskevent_id"`
}

type RiskIndikatorResult struct {
	ID                int    `json:"id"`
	RiskIndicatorCode string `json:"risk_indicator_code"`
	RiskIndicator     string `json:"risk_indicator"`
}

type RRMHeadRequest struct {
	BTRTL   string `json:"btrtl"`
	Keyword string `json:"keyword"`
}

type PimpinanUkerRequest struct {
	BRANCH string `json:"branch"`
}

type CommonRequest struct {
	ID        int64
	MenuID    int    `json:"menu_id"`
	SubMenu   int64  `json:"sub_menu_id"`
	TypeID    int64  `json:"type_id"`
	PartID    int64  `json:"part_id"`
	SubPartID int64  `json:"sub_part_id"`
	Keyword   string `json:"keyword"`
	Status    string `json:"status"`
}

// Mq Enhance
type MasterDataOption struct {
	ID    string `json:"id"`
	Value string `json:"value"`
	Label string `json:"label"`
	Api   string `json:"api"`
}

type BrcKeywordRequest struct {
	Keyword  string `json:"keyword"`
	TipeUker string `json:"tipe_uker"`
	Region   string `json:"region"`
	Order    string `json:"order"`
	Sort     string `json:"sort"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
	Page     int    `json:"page"`
}
