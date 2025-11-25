package dto

type Response[T any] struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       T      `json:"data"`
	Errors     *any   `json:"errors"`
}

type Pagination struct {
	Total int64 `json:"total"`
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
	Next  int   `json:"next"`
	Prev  int   `json:"prev"`
}

type ControlAttributeFiler struct {
	CodeIDs []string `form:"code_ids"`
	Code    string   `form:"code"`
	Status  string   `form:"status"`
}

type DtoRiskControlAttributeResponse struct {
	List       []DtoAttribute `json:"list"`
	Pagination Pagination     `json:"pagination"`
}

type DtoAttribute struct {
	ID        string `json:"id"`
	CodeID    string `json:"code_id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}
