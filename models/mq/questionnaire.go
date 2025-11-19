package models

type QuestionListRequest struct {
	ID        int    `json:"id"`
	MenuID    int    `json:"menu_id"`
	Code      string `json:"code"`
	TypeID    int    `json:"type_id"`
	PartID    int    `json:"part_id"`
	SubPartID int    `json:"sub_part_id"`
	Page      int    `json:"page"`
	Limit     int    `json:"limit"`
	Order     string `json:"order"`
	Offset    int    `json:"offset"`
	Sort      string `json:"sort"`
	PERNR     string `json:"pernr"`
	Counter   int    `json:"counter"`
}

type RequestQuestion struct {
	ID        int64                `json:"id"`
	TypeID    int16                `json:"type_id"`
	PartID    int16                `json:"part_id"`
	SubPartID int16                `json:"sub_part_id"`
	Isi       []RequestQuestionIsi `json:"isi"`
	Pernr     string               `json:"pernr"`
	Sname     string               `json:"sname"`
}

type RequestQuestionIsi struct {
	Code       string `json:"code"`
	Pertanyaan string `json:"pertanyaan"`
	// BobotPertanyaan string                       `json:"bobot_pertanyaan"`
	JenisJawaban string                       `json:"jenis_jawaban"`
	Format       string                       `json:"format"`
	OpsiJawaban  []RequestQuestionOpsiJawaban `json:"opsi_jawaban"`
	Default      bool                         `json:"default"`
	DefaultValue string                       `json:"default_value"`
	KetJawaban   bool                         `json:"ket_jawaban"`
	Status       string                       `json:"status"`
	Mandatory    bool                         `json:"mandatory"`
	Nilai        string                       `json:"nilai"`
}

type RequestQuestionOpsiJawaban struct {
	ID             int    `json:"id"`
	Opsi           string `json:"opsi"`
	AnswerOption   string `json:"answer_option"`
	NextProcess    string `json:"next_process"`
	NextProcessID  int16  `json:"next_process_id"`
	KetNextProcess string `json:"next_process_ket"`
	Result         string `json:"result"`
	KetResult      string `json:"ket_result"`
	Nilai          string `json:"nilai"`
}

type RequestQuestionUpdate struct {
	ID           int64                        `json:"id"`
	TypeID       int16                        `json:"type_id"`
	PartID       int16                        `json:"part_id"`
	SubPartID    int16                        `json:"sub_part_id"`
	Pertanyaan   string                       `json:"pertanyaan"`
	JenisJawaban string                       `json:"jenis_jawaban"`
	Format       string                       `json:"format"`
	OpsiJawaban  []RequestQuestionOpsiJawaban `json:"opsi_jawaban"`
	Default      bool                         `json:"default"`
	DefaultValue string                       `json:"default_value"`
	KetJawaban   bool                         `json:"input_keterangan"`
	Status       string                       `json:"status"`
	Mandatory    bool                         `json:"mandatory"`
	Nilai        string                       `json:"nilai"`
	Pernr        string                       `json:"pernr"`
	Sname        string                       `json:"sname"`
}
