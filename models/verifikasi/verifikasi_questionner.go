package models

type VerifikasiQuestionner struct {
	ID                int64
	VerifikasiID      int64
	Questionner       string
	DataSumber        string
	Checker           string
	Signer            string
	ApprovalOrd       string
	JenisFraud        string
	StatusValidasiRmc string
}
