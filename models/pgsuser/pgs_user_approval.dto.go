package models

type PgsUserApprovalRequest struct {
	ID             int64  `json:"id"`
	IDPgsUser      int64  `json:"id_pgs_user"`
	ApprovalID     string `json:"approval_id"`
	ApprovalDesc   string `json:"approval_desc"`
	ApprovalDate   string `json:"approval_date"`
	ApprovalStatus string `json:"approval_status"`
}

type PgsUserApprovalResponse struct {
	ID             int64  `json:"id"`
	IDPgsUser      int64  `json:"id_pgs_user"`
	ApprovalID     string `json:"approval_id"`
	ApprovalDesc   string `json:"approval_desc"`
	ApprovalDate   string `json:"approval_date"`
	ApprovalStatus string `json:"approval_status"`
}

type ApprovalUpdate struct {
	ID             int64  `json:"id"`
	ApprovalDesc   string `json:"approval_desc"`
	ApprovalDate   string `json:"approval_date"`
	ApprovalStatus string `json:"approval_status"`
}

func (pa PgsUserApprovalRequest) TableName() string {
	return "pgs_user_approval"
}

func (pa PgsUserApprovalResponse) TableName() string {
	return "pgs_user_approval"
}

func (pa ApprovalUpdate) TableName() string {
	return "pgs_user_approval"
}
