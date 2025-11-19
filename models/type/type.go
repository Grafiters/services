package models

import (
	"riskmanagement/lib"
)

type TypeCodeNullResponse struct {
	Code lib.NullString `json:"code"`
}

type TypeRequest struct {
	ID                  int64                  `json:"id"`
	IDMenu              int64                  `json:"id_menu"`
	Title               string                 `json:"title"`
	MenuType            string                 `json:"menu_type"`
	IDParent            int64                  `json:"id_parent"`
	LastWeight          int64                  `json:"last_weight"`
	MenuID              string                 `json:"menu_id"`
	ChildStatus         int64                  `json:"child_status"`
	Code                string                 `json:"code"`
	Name                string                 `json:"name"`
	IsQuiz              string                 `json:"is_quiz"`
	ApprovalResponse    string                 `json:"approval_response"`
	ApprovalTotal       string                 `json:"approval_total"`
	Role                RoleRequest            `json:"role"`
	Approver            []Approver             `json:"approver"`
	TypeThreshold       []TypeThresholdRequest `json:"threshold"`
	PartWeight          []PartDraft            `json:"part_weight"`
	ApprovalStatus      string                 `json:"approval_status"`
	OldStatus           string                 `json:"old_status"`
	Status              string                 `json:"status"`
	IsDelete            int64                  `json:"is_delete"`
	Order               string                 `json:"order"`
	Sort                string                 `json:"sort"`
	Offset              int                    `json:"offset"`
	Limit               int                    `json:"limit"`
	Page                int                    `json:"page"`
	PERNR               string                 `json:"pernr"`
	SNAME               string                 `json:"sname"`
	QuestionnaireStatus string                 `json:"questionnaire_status"`
	ActiveDate          *string                `json:"active_date"`
	NonactiveDate       *string                `json:"nonactive_date"`
	CreatedAt           *string                `json:"created_at"`
	UpdatedAt           *string                `json:"updated_at"`
}

type PartWeightRequest struct {
	ID     int64   `json:"id"`
	Weight *string `json:"weight"`
}

type TypeDraft struct {
	ID                  int64
	MenuType            string
	MenuID              string
	Code                string
	Name                string
	IsQuiz              string
	ApprovalResponse    string
	ApprovalStatus      string
	CurrentApprover     string
	ApprovalTotal       string
	Status              string
	QuestionnaireStatus string
	PERNR               string
	SNAME               string
	ActiveDate          *string
	NonactiveDate       *string
	CreatedAt           *string
	UpdatedAt           *string
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

type TypeFilter struct {
	ID                  int64   `json:"id"`
	MenuID              string  `json:"menu_id"`
	MenuType            string  `json:"menu_type"`
	Code                string  `json:"code"`
	Name                string  `json:"name"`
	IsQuiz              string  `json:"is_quiz"`
	Status              string  `json:"status"`
	ApprovalResponse    string  `json:"approval_response"`
	ApprovalTotal       string  `json:"approval_total"`
	ApprovalStatus      string  `json:"approval_status"`
	ApproverTurn        int64   `json:"approver_turn"`
	ApproveDate         *string `json:"approve_date"`
	QuestionnaireStatus string  `json:"questionnaire_status"`
	PERNR               string  `json:"pernr"`
	SNAME               string  `json:"sname"`
	TypeStatus          string  `json:"type_status"`
	ApprovedType        int64   `json:"approved_type"`
	ActiveDate          *string `json:"active_date"`
	NonactiveDate       *string `json:"nonactive_date"`
	CreatedAt           *string `json:"created_at"`
	UpdatedAt           *string `json:"updated_at"`
}

type GenerateQuestionnaireRequest struct {
	TypeID          int64 `json:"type_id"`
	PartID          int64 `json:"part_id"`
	QuestionnaireID int64 `json:"questionnaire_id"`
	AnswerID        int64 `json:"answer_id"`
}

type GenerateQuestionnaireResponse struct {
	ID       int64              `json:"id"`
	TypeID   int64              `json:"type_id"`
	Name     string             `json:"name"`
	Weight   *string            `json:"weight"`
	SubPart  string             `json:"sub_part"`
	Question []QuestionnaireRes `json:"question"`
}

type ORDMember struct {
	PERNR string `json:"pernr"`
	SNAME string `json:"sname"`
}

type Questionnaire struct {
	ID              int64  `json:"id"`
	Code            string `json:"code"`
	TypeID          int64  `json:"type_id"`
	PartID          int64  `json:"part_id"`
	Question        string `json:"question"`
	QuestionType    string `json:"question_type"`
	Mandatory       string `json:"mandatory"`
	FieldFormat     string `json:"field_format"`
	InputKeterangan string `json:"input_keterangan"`
	Nilai           string `json:"nilai"`
}

type QuestionnaireRes struct {
	ID              int64    `json:"id"`
	Code            string   `json:"code"`
	TypeID          int64    `json:"type_id"`
	PartID          int64    `json:"part_id"`
	Question        string   `json:"question"`
	QuestionType    string   `json:"question_type"`
	Answer          []Answer `json:"answer"`
	Mandatory       string   `json:"mandatory"`
	FieldFormat     string   `json:"field_format"`
	InputKeterangan string   `json:"input_keterangan"`
	Nilai           string   `json:"nilai"`
}

type Answer struct {
	AnswerOption string `json:"answer_option"`
	NextProcess  string `json:"next_process"`
	ResultDesc   string `json:"result_desc"`
	DefaultValue string `json:"default_value"`
	Nilai        string `json:"nilai"`
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
	IDParent int64  `json:"id_parent"`
}

type Type struct {
	ID                  int64
	MenuType            string
	MenuID              string
	Code                string
	Name                string
	IsQuiz              string
	ApprovalResponse    string
	ApprovalStatus      string
	ApprovalTotal       string
	Status              string
	QuestionnaireStatus string
	PERNR               string
	SNAME               string
	Versioning          int64
	ActiveDate          *string
	NonactiveDate       *string
	CreatedAt           *string
	UpdatedAt           *string
}

type TypeThresholdRequest struct {
	ID          int64  `json:"id"`
	TypeID      int64  `json:"type_id"`
	MinValue    string `json:"min_value"`
	MaxValue    string `json:"max_value"`
	Description string `json:"description"`
}

type TypeThreshold struct {
	ID          int64   `json:"id"`
	TypeID      int64   `json:"type_id"`
	MinValue    float64 `json:"min_value"`
	MaxValue    float64 `json:"max_value"`
	Description string  `json:"description"`
}

type TypeThresholdResponse struct {
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
	Role       []Role                  `json:"role"`
	Approver   []Approver              `json:"approver"`
	Threshold  []TypeThresholdResponse `json:"threshold"`
	Weight     []PartDraft             `json:"weight"`
	RejectNote []RejectedType          `json:"reject_note"`
}

type RejectedType struct {
	TypeID    int64  `json:"type_id"`
	PERNR     string `json:"pernr"`
	SNAME     string `json:"sname"`
	Notes     string `json:"notes"`
	CreatedAt *string
}

type TypeTX struct {
	ID                  int64
	DraftTypeID         int64
	MenuType            string
	MenuID              string
	Code                string
	Name                string
	IsQuiz              string
	ApprovalResponse    string
	ApprovalStatus      string
	ApprovalTotal       string
	Status              string
	QuestionnaireStatus string
	PERNR               string
	SNAME               string
	ActiveDate          *string
	NonactiveDate       *string
	Versioning          int64
	CreatedAt           *string
	UpdatedAt           *string
}

type RoleRequestUpdateType struct {
	MenuID string `json:"menu_id"`
	NewID  int64  `json:"new_id"`
}

type GenerateRequest struct {
	ID         int16  `json:"id"`
	TypeID     int16  `json:"type_id"`
	ResponseID int16  `json:"response_id"`
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
	Order      string `json:"order"`
	Offset     int    `json:"offset"`
	Sort       string `json:"sort"`
}

type QuestionnaireQuery struct {
	ID             int16   `json:"id" gorm:"primaryKey"`
	Code           string  `json:"code"`
	Name           string  `json:"name"`
	Quiz           string  `json:"quiz"`
	Approval       string  `json:"approval"`
	ApprovalTotal  string  `json:"approval_total"`
	PernrApproval  string  `json:"pernr_approval"`
	SnameApproval  string  `json:"sname_approval"`
	PosisiApprover string  `json:"posisi_approver"`
	NilaiAkhir     float32 `json:"nilai_akhir"`
	Kategori       string  `json:"kategori"`
	Status         string  `json:"status"`
	SubmitedDate   *string `json:"submited_date"`
	UpdatedDate    *string `json:"updated_date"`
}

type PartQuery struct {
	PartID    int16  `json:"part_id" gorm:"primaryKey"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	Weight    int    `json:"weight"`
	ViewValue int    `json:"view_value"`
}

type QuestionQuery struct {
	QuestID              int16  `json:"quest_id" gorm:"primaryKey"`
	TypeID               int16  `json:"type_id"`
	Code                 string `json:"code"`
	PartID               int16  `json:"part_id"`
	SubPartID            int16  `json:"sub_part_id"`
	ResponseUserDetailID string `json:"response_user_detail_id"`
	Question             string `json:"question"`
	QuestionType         string `json:"question_type"`
	FieldFormat          string `json:"field_format"`
	Nilai                int16  `json:"nilai"`
	InputKeterangan      string `json:"input_keterangan"`
	Answer               string `json:"answer"`
	KetAnswer            string `json:"ket_answer"`
	Mandatory            int16  `json:"mandatory"`
	Status               string `json:"status"`
	Readonly             int    `json:"readonly"`
}

type QuestionnaireType struct {
	ID             int16       `json:"id" gorm:"primaryKey"`
	Code           string      `json:"code"`
	Name           string      `json:"name"`
	Quiz           string      `json:"quiz"`
	Approval       string      `json:"approval"`
	ApprovalTotal  string      `json:"approval_total"`
	Part           []PartDraft `json:"part" gorm:"foreignKey:PartID"`
	PosisiApprover string      `json:"posisi_approver"`
	PernrApproval  string      `json:"pernr_approval"`
	SnameApproval  string      `json:"sname_approval"`
	NilaiAkhir     float32     `json:"nilai_akhir"`
	Kategori       string      `json:"kategori"`
	Status         string      `json:"status"`
	Readonly       bool        `json:"read_only"`
	SubmitedDate   *string     `json:"submited_date"`
	UpdatedDate    *string     `json:"updated_date"`
}

type SubPartQuery struct {
	SubPartID int16  `json:"sub_part_id" gorm:"primaryKey"`
	PartID    int16  `json:"part_id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
}

type OptionAnswerQuest struct {
	ID             int16  `json:"id"`
	QuestionID     int16  `json:"question_id"`
	AnswerOption   string `json:"answer_option"`
	NextProcess    string `json:"next_process"`
	NextProcessID  int16  `json:"next_process_id"`
	KetNextProcess string `json:"next_process_ket"`
	Nilai          int16  `json:"nilai"`
	DefaultValue   string `json:"default_value"`
	Status         string `json:"status"`
}

type Question struct {
	QuestID              int16               `json:"quest_id" gorm:"primaryKey"`
	TypeID               int16               `json:"type_id"`
	Code                 string              `json:"code"`
	PartID               int16               `json:"part_id"`
	SubPartID            int16               `json:"sub_part_id"`
	ResponseUserDetailID string              `json:"response_user_detail_id"`
	Question             string              `json:"question"`
	QuestionType         string              `json:"question_type"`
	FieldFormat          string              `json:"field_format"`
	OptionQuest          []OptionAnswerQuest `json:"option_answer"`
	CheckedAnswer        []string            `json:"checked_answer"`
	TextAnswer           string              `json:"text_answer"`
	SelectedAnswer       string              `json:"selected_answer"`
	Nilai                int16               `json:"nilai"`
	InputKeterangan      string              `json:"input_keterangan"`
	KetAnswer            string              `json:"ket_answer"`
	Mandatory            int16               `json:"mandatory"`
	Status               string              `json:"status"`
	Readonly             bool                `json:"readonly"`
	Urutan               int                 `json:"urutan"`
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

func (TypeThresholdResponse) TableName() string {
	return "q_type_threshold"
}

type WeighTotalResponse struct {
	WeightTotal int64 `json:"weight_total"`
}
