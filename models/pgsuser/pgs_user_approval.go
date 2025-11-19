package models

type PgsUserApproval struct {
	ID             int64
	IDPgsUser      int64
	ApprovalID     string
	ApprovalDesc   string
	ApprovalDate   string
	ApprovalStatus string
}

func (p PgsUserApproval) TableName() string {
	return "pgs_user_approval"
}
