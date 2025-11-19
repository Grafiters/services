package dio

type DioModel struct {
	Id_template              string
	Branch_code_penerima     string
	Orgeh_penerima           string
	Branch_code_tindasan     string
	Orgeh_tindasan           string
	Pn_penerima              string
	Pn_tindasan              string
	Kode_surat               string
	Kerahasiaan              string
	Kesegeraan               string
	Kepada_yth               string
	Perihal                  string
	Isi_surat                string
	Id_maker                 string
	Pn_approver              string
	Status_approver          string
	Folder_attachment        string
	Nama_attachment_uploaded string
	Sla                      int
	Surat_keluar_approver    string
	App_id                   string
	Password                 string
	Uker_edit                string
	Jabatan_edit             string
}

type Summaries struct {
	QuestionnaireMasuk            string
	QuestionnaireMasukBelumReview string
	DivisionOrg                   string
	MonthPeriod                   string
	YearPeriod                    string
	Updated                       string
}
type responseDio struct {
	ResponseCode string `json:"responseCode"`
	ResponseDesc string `json:"responseDesc"`
}

type ApprovalBisnis struct {
	MakerId    string
	ApprovalId string
}
