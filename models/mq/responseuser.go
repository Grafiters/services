package models

type RequestResponseUserList struct {
	TypeID       int16  `json:"type_id"`
	Status       string `json:"status"`
	TanggalAwal  string `json:"tanggal_awal"`
	TanggalAkhir string `json:"tanggal_akhir"`
	Pernr        string `json:"pernr"`
	Page         int    `json:"page"`
	Limit        int    `json:"limit"`
	Order        string `json:"order"`
	Offset       int    `json:"offset"`
	Sort         string `json:"sort"`
}

type RequestResponseApprovalList struct {
	TypeID         int16  `json:"type_id"`
	Status         string `json:"status"`
	Code           string `json:"code"`
	Pernr          string `json:"pernr"`
	PosisiApprover string `json:"posisi_approver"`
	Page           int    `json:"page"`
	Limit          int    `json:"limit"`
	Order          string `json:"order"`
	Offset         int    `json:"offset"`
	Sort           string `json:"sort"`
}

// store
type RequestUserHistory struct {
	ID           int16        `json:"id"`
	TypeID       int16        `json:"type_id"`
	Pernr        string       `json:"pernr"`
	Sname        string       `json:"sname"`
	RGDESC       string       `json:"rgdesc"`
	MBDESC       string       `json:"mbdesc"`
	BRDESC       string       `json:"brdesc"`
	Jabatan      string       `json:"jabatan"`
	Pertanyaan   []Pertanyaan `json:"pertanyaan"`
	Approval     []Approval   `json:"approval"`
	Bagian       []Bagian     `json:"nilai_bagian"`
	Status       string       `json:"status"`
	Quiz         string       `json:"quiz"`
	Versioning   int16        `json:"versioning"`
	SubmitedDate *string      `json:"submited_date"`
	UpdatedDate  *string      `json:"updated_date"`
}

type Pertanyaan struct {
	ID          int16  `json:"response_user_detail_id"`
	QuestID     int16  `json:"quest_id"`
	Answer      string `json:"answer"`
	AnswerValue string `json:"answer_value"`
	Keterangan  string `json:"ket_answer"`
	NilaiAkhir  int16  `json:"nilai_akhir"`
	Readonly    bool   `json:"readonly"`
}

type Approval struct {
	PernrApproval string `json:"pn_approval"`
	NamaApproval  string `json:"nama_approval"`
}

type Bagian struct {
	ID          int16 `json:"response"`
	PartID      int16 `json:"part_id"`
	NilaiBagian int16 `json:"nilai_bagian"`
	Weight      int16 `json:"weight"`
}

// =====================
type GenerateRequest struct {
	ID         int16  `json:"id"`
	TypeID     int16  `json:"type_id"`
	ResponseID int16  `json:"response_id"`
	Versioning int16  `json:"versioning"`
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
	Order      string `json:"order"`
	Offset     int    `json:"offset"`
	Sort       string `json:"sort"`
	Pernr      string `json:"pernr"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
}

type ApprovalUpdate struct {
	ID             int16  `json:"id"`
	PosisiApprover string `json:"posisi_approver"`
	Status         string `json:"status"`
	Pernr          string `json:"pernr"`
}

type RejectedUpdate struct {
	ResponseID   int16   `json:"response_id"`
	Penolak      string  `json:"penolak"`
	Catatan      string  `json:"catatan"`
	TanggalTolak *string `json:"tanggal_tolak"`
	Pernr        string  `json:"pernr"`
}

type UpdateResponseUserHistory struct {
	ID             int16   `json:"id"`
	PernrApproval  string  `json:"pernr_approval"`
	SnameApproval  string  `json:"sname_approval"`
	PosisiApprover string  `json:"posisi_approver"`
	UpdatedDate    *string `json:"updated_date"`
	Status         string  `json:"status"`
	Pernr          string  `json:"pernr"`
}

type RequestPartid struct {
	ResponseID int    `json:"response_id"`
	PartID     string `json:"part_id"`
	Readonly   bool   `json:"read_only"`
	EnablePart string `json:"enable_part"`
	Pernr      string `json:"pernr"`
}
