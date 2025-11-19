package models

type TypeRequest struct {
	ID               int64           `json:"id"`
	IDMenu           string          `json:"id_menu"`
	Title            string          `json:"title"`
	MenuType         string          `json:"menu_type"`
	IDParent         string          `json:"id_parent"`
	LastWeight       int64           `json:"last_weight"`
	MenuID           string          `json:"menu_id"`
	ChildStatus      int64           `json:"child_status"`
	Code             string          `json:"code"`
	Name             string          `json:"name"`
	IsQuiz           string          `json:"is_quiz"`
	ApprovalStatus   string          `json:"approval_status"`
	ApprovalResponse string          `json:"approval_response"`
	ApprovalTotal    string          `json:"approval_total"`
	Role             RoleRequest     `json:"role"`
	Approver         []Approver      `json:"approver"`
	TypeThreshold    []TypeThreshold `json:"threshold"`
	PartWeight       []PartDraft     `json:"part_weight"`
	OldStatus        string          `json:"old_status"`
	Status           string          `json:"status"`
	Order            string          `json:"order"`
	Sort             string          `json:"sort"`
	Offset           int             `json:"offset"`
	Limit            int             `json:"limit"`
	Page             int             `json:"page"`
	PERNR            string          `json:"pernr"`
	SNAME            string          `json:"sname"`
	ActiveDate       *string         `json:"active_date"`
	NonactiveDate    *string         `json:"nonactive_date"`
	CreatedAt        *string         `json:"created_at"`
	UpdatedAt        *string         `json:"updated_at"`
}

type TypeDraft struct {
	ID               int64
	MenuType         string
	MenuID           string
	Code             string
	Name             string
	IsQuiz           string
	ApprovalResponse string
	ApprovalStatus   string
	ApprovalTotal    string
	Status           string
	PERNR            string
	SNAME            string
	ActiveDate       *string
	NonactiveDate    *string
	CreatedAt        *string
	UpdatedAt        *string
}

type TypeResponse struct {
	ID                  int64           `json:"id"`
	MenuID              string          `json:"menu_id"`
	MenuType            string          `json:"menu_type"`
	Code                string          `json:"code"`
	Name                string          `json:"name"`
	IsQuiz              string          `json:"is_quiz"`
	MstMenu             MstMenuResponse `json:"mst_menu"`
	MstMenuRRM          MstMenuResponse `json:"mst_menu_rrm"`
	Status              string          `json:"status"`
	QuestionnaireStatus string          `json:"questionnaire_status"`
	ApprovalResponse    string          `json:"approval_response"`
	ApprovalTotal       string          `json:"approval_total"`
	ApprovalStatus      string          `json:"approval_status"`
	ApproverTurn        int64           `json:"approver_turn"`
	ApproveDate         *string         `json:"approve_date"`
	PERNR               string          `json:"pernr"`
	SNAME               string          `json:"sname"`
	TypeStatus          string          `json:"type_status"`
	ApprovedType        int64           `json:"approved_type"`
	ActiveDate          *string         `json:"active_date"`
	NonactiveDate       *string         `json:"nonactive_date"`
	CreatedAt           *string         `json:"created_at"`
	UpdatedAt           *string         `json:"updated_at"`
}

type ResponseFilter struct {
	Data       []TypeResponse `json:"data"`
	Message    string         `json:"message"`
	Status     string         `json:"status"`
	Pagination Pagination     `json:"pagination"`
}

type ResponseStore struct {
	Data       bool       `json:"data"`
	Message    string     `json:"message"`
	Status     string     `json:"status"`
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	CurrentPage int `json:"current_page"`
	LastPage    int `json:"last_page"`
	PerPage     int `json:"per_page"`
	Total       int `json:"total"`
	TotalData   int `json:"total_data"`
}

type GenerateQuestionnaireRequest struct {
	TypeID          int64  `json:"type_id"`
	PartID          int64  `json:"part_id"`
	QuestionnaireID int64  `json:"questionnaire_id"`
	AnswerID        int64  `json:"answer_id"`
	PERNR           string `json:"pernr"`
}

type GenerateQuestionnaireResponse struct {
	ID       int64              `json:"id"`
	TypeID   int64              `json:"type_id"`
	Name     string             `json:"name"`
	Weight   string             `json:"weight"`
	SubPart  string             `json:"sub_part"`
	Question []QuestionnaireRes `json:"question"`
}

type ORDMember struct {
	PERNR string `json:"pernr"`
	SNAME string `json:"sname"`
}

type Questionnaire struct {
	ID           int64  `json:"id"`
	Code         string `json:"code"`
	TypeID       int64  `json:"type_id"`
	PartID       int64  `json:"part_id"`
	Question     string `json:"question"`
	QuestionType string `json:"question_type"`
}

type QuestionnaireRes struct {
	ID           int64    `json:"id"`
	Code         string   `json:"code"`
	TypeID       int64    `json:"type_id"`
	PartID       int64    `json:"part_id"`
	Question     string   `json:"question"`
	QuestionType string   `json:"question_type"`
	Answer       []Answer `json:"answer"`
}

type Answer struct {
	AnswerOption string `json:"answer_option"`
	NextProcess  string `json:"next_process"`
	ResultDesc   string `json:"result_desc"`
	DefaultValue string `json:"default_value"`
}

type RoleRequest struct {
	RoleID []string `json:"role_id"`
	MenuID string   `json:"menu_id"`
}

type Role struct {
	RoleID string `json:"role_id"`
	MenuID string `json:"menu_id"`
}

type ApproverRequest struct {
	PERNR          string  `json:"pernr"`
	SNAME          string  `json:"sname"`
	TypeID         int64   `json:"type_id"`
	ApproverOrder  int64   `json:"approver_order"`
	Notes          string  `json:"notes"`
	ApproveDate    *string `json:"approver_date"`
	ApprovalStatus string  `json:"approval_status"`
}

type Approver struct {
	ID            int64
	PERNR         string  `json:"pernr"`
	SNAME         string  `json:"sname"`
	TypeID        int64   `json:"type_id"`
	ApproverOrder int64   `json:"approver_order"`
	ApproveDate   *string `json:"approver_date"`
}

type TypeTableRequest struct {
	TypeID     int64 `json:"type_id"`
	HeaderType string
}

type MstMenuResponse struct {
	Title    string `json:"title"`
	IDParent string `json:"id_parent"`
}

type Type struct {
	ID               int64
	MenuType         string
	MenuID           string
	Code             string
	Name             string
	IsQuiz           string
	ApprovalResponse string
	ApprovalStatus   string
	ApprovalTotal    string
	Status           string
	PERNR            string
	SNAME            string
	ActiveDate       *string
	NonactiveDate    *string
	CreatedAt        *string
	UpdatedAt        *string
}

type TypeThreshold struct {
	ID          int64  `json:"id"`
	TypeID      int64  `json:"type_id"`
	MinValue    string `json:"min_value"`
	MaxValue    string `json:"max_value"`
	Description string `json:"description"`
}

type PartDraft struct {
	ID         int64
	Code       string
	TypeID     int64
	Name       string
	ResultType string
	SubPart    int64
	PartID     int64
	ViewValue  string
	Weight     *string
	Status     string
	CreatedAt  *string
	UpdatedAt  *string
}

type DataQuestForm struct {
	Role      []Role          `json:"role"`
	Approver  []Approver      `json:"approver"`
	Threshold []TypeThreshold `json:"threshold"`
	Weight    []PartDraft     `json:"weight"`
}

type RejectedType struct {
	TypeID    int64  `json:"type_id"`
	PERNR     string `json:"pernr"`
	SNAME     string `json:"sname"`
	Notes     string `json:"notes"`
	CreatedAt *string
}

type TypeTX struct {
	ID               int64
	DraftTypeID      int64
	MenuType         string
	MenuID           string
	Code             string
	Name             string
	IsQuiz           string
	ApprovalResponse string
	ApprovalStatus   string
	ApprovalTotal    string
	Status           string
	PERNR            string
	SNAME            string
	ActiveDate       *string
	NonactiveDate    *string
	Versioning       int64
	CreatedAt        *string
	UpdatedAt        *string
}

type MenuRequest struct {
	PERNR string `json:"pernr"`
}

type PartWeightRequest struct {
	ID     int64   `json:"id"`
	Weight *string `json:"weight"`
	PERNR  string  `json:"pernr"`
}

func (TypeDraft) TableName() string {
	return "q_type_draft"
}

func (Type) TableName() string {
	return "q_type"
}

func (TypeTX) TableName() string {
	return "q_type_tx"
}

func (PartDraft) TableName() string {
	return "q_part_draft"
}

func (TypeThreshold) TableName() string {
	return "q_type_threshold"
}

func (Role) TableName() string {
	return "mst_roles_questionnaire"
}

func (Approver) TableName() string {
	return "q_approver"
}

func (Questionnaire) TableName() string {
	return "q_questionnaire_draft"
}

func (Answer) TableName() string {
	return "q_answer_draft"
}

func (MstMenuResponse) TableName() string {
	return "mst_menu"
}

func (ORDMember) TableName() string {
	return "pa0001_eof"
}

func (RejectedType) TableName() string {
	return "note_reject"
}
