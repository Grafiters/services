package dio

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

func DioRemainder(dbdefault *sql.DB, totalWorker, dbMaxConns, dbMaxIdleConns int) {

	fmt.Println("Kirim Remainder=>", time.Now())
	// summaries := getSummaries(dbdefault)
	// approval := getApprovalBisnis(dbdefault)

	// fmt.Println("summaries")
	// fmt.Println(summaries)
	// fmt.Println("approval")
	// fmt.Println(approval)

	// dioModel := DioModel{}
	// for _, summary := range summaries {
	// 	isi_surat := "<p>1.Peraturan OJK (POJK) nomor 23 /POJK.01/2019 tentang Perubahan atas POJK nomor 12/POJK.01/2017 tentang Penerapan Program APU PPT di Sektor Jasa Keuangan. </p>"
	// 	isi_surat += "<p>2.Peraturan OJK (POJK) Nomor 12/POJK.01/2017 tentang Penerapan Program APU PPT di Sektor Jasa Keuangan.</p>"
	// 	isi_surat += "<p>3.Surat Keputusan Direksi BRI PP No.1 -DIR/KPT/08/2018 tanggal 31 Agustus 2018 perihal Pedoman Pelaksanaan Penerapan Program APU PPT BRI</p>"
	// 	isi_surat += "<p>Menunjuk POJK dan SK Direksi BRI diatas, dengan ini kami sampaikan bahwa Bank wajib melakukan penilaian tingkat risiko TPPU/TPPT terhadap Bank Koresponden.<br> Adapun Compliance Division telah mengembangkan Sistem RBA Tools dalam rangka penerapan penilaian tingkat risiko Bank Koresponden dimaksud berdasarkan hasil <br> pelaksanaan due diligence dari pengisian AML Questionnaire melalui web yang telah dilengkapi dengan dokumen pendukung.</p>"
	// 	isi_surat += "Sehubungan dengan hal tersebut, kami mohon bantuan dari Divisi terkait untuk melakukan hal-hal sebagai berikut:"
	// 	isi_surat += `<ol type="1">`
	// 	isi_surat += "<li>Melakukan pemantauan terkait pemenuhan pengisian AML Questionnaire dan dokumen pendukung dari Bank Koresponden pada Aplikasi RBA Tools <br> Bank Koresponden dengan mengakses Bristars  (Menu BRISTARS)</li>"
	// 	isi_surat += "<li>Aplikasi RBA Tools Bank Koresponden telah memproses hasil pengisian AML Questionnaire Koresponden melalui web yang telah disediakan. <br> Pada bulan (Juli) terdapat hasil pengisian sebagai berikut: <br>"
	// 	isi_surat += `<ol type="a">`
	// 	isi_surat += "<li>AML Questionnaire masuk:" + summary.QuestionnaireMasuk + " </li>"
	// 	isi_surat += "<li>AML Questionnaire masuk yang belum direview: " + summary.QuestionnaireMasukBelumReview + " </li>"
	// 	isi_surat += "</ol></li>"
	// 	isi_surat += "<li>Atas pemantauan terkait hasil pengisian AMl Questionnaire Bank Koresponden, agar dilakukan hal-hal sbb: <br>"
	// 	isi_surat += `<ol type="a">`
	// 	isi_surat += "<li>Melakukan review kelengkapan pengisian field-field AML Questionnaire</li>"
	// 	isi_surat += "<li>Memeriksa kesesuaian hasil jawaban pada AML Questionnaire</li>"
	// 	isi_surat += "<li>Menginput hasil review pada Aplikasi RBA Tools Bank Koresponden</li>"
	// 	isi_surat += "</ol></li>"
	// 	isi_surat += "<li>Apabila pada pengisian field-field AML Questionnaire terdapat data tidak lengkap, agar dipilih opsi tidak sesuai pada Aplikasi RBA Tools <br> Bank Koresponden dan melakukan koordinasi dengan Bank Koresponden kembali untuk menginput data yang tidak lengkap tersebut.</li>"
	// 	isi_surat += "<li>Hasil penilaian tingkat risiko TPPU/TPPT Bank Koresponden diterapkan hanya untuk kepentingan internal BRI.</li></ol>"
	// 	isi_surat += "<p>Demikian, atas perhatian dan kerjasamanya yang baik disampaikan terima kasih.</p>"

	// 	dioModel.Id_template = "1"
	// 	dioModel.Branch_code_penerima = summary.DivisionOrg
	// 	dioModel.Orgeh_penerima = ""
	// 	dioModel.Branch_code_tindasan = ""
	// 	dioModel.Orgeh_tindasan = ""
	// 	dioModel.Pn_penerima = ""
	// 	dioModel.Pn_tindasan = ""
	// 	dioModel.Kode_surat = "R.-KEP/IMG/"
	// 	dioModel.Kerahasiaan = "B"
	// 	dioModel.Kesegeraan = "B"
	// 	dioModel.Kepada_yth = "Division Head of (Divisi Bisnis)"
	// 	dioModel.Perihal = "Review AML Questionnaire dan KYC Document"
	// 	dioModel.Isi_surat = isi_surat
	// 	dioModel.Id_maker = approval.MakerId
	// 	dioModel.Pn_approver = approval.ApprovalId
	// 	dioModel.Status_approver = "Signer"
	// 	dioModel.Folder_attachment = ""
	// 	dioModel.Nama_attachment_uploaded = ""
	// 	dioModel.Sla = 0
	// 	dioModel.Surat_keluar_approver = "y"
	// 	dioModel.App_id = os.Getenv("APP_ID_DIO")
	// 	dioModel.Password = os.Getenv("PASSWORD_DIO")
	// 	dioModel.Uker_edit = ""
	// 	dioModel.Jabatan_edit = ""
	// 	dioServer(dbdefault, dioModel)
	// }
}
func getSummaries(dbdefault *sql.DB) (summaries []Summaries) {

	query := `
	SELECT
		SUM( CASE WHEN status = '01a' THEN 1 ELSE 0 END ) AS questionnaireMasuk,
		SUM( CASE WHEN status = '01b' THEN 1 ELSE 0 END ) AS questionnaireMasukBelumDireview,
		divisionOrg,
		monthPeriod,
		yearPeriod,
		NOW() AS updated
	FROM
		tbl_form_summarize
	WHERE
		yearPeriod = '` + fmt.Sprint(time.Now().Year()) + `' and monthPeriod = '` + fmt.Sprint(int(time.Now().Month())) + `' group by divisionOrg`

	// err := dbdefault.QueryRow(query).Scan(&summaries.QuestionnaireMasuk, &summaries.QuestionnaireMasukBelumReview, &summaries.DivisionOrg, &summaries.MonthPeriod, &summaries.YearPeriod, &summaries.Updated)
	rows, err := dbdefault.Query(query)

	fmt.Println(query)

	if err == sql.ErrNoRows {
		return summaries
	}

	data := Summaries{}
	for rows.Next() {
		_ = rows.Scan(
			&data.QuestionnaireMasuk,
			&data.QuestionnaireMasukBelumReview,
			&data.DivisionOrg,
			&data.MonthPeriod,
			&data.YearPeriod,
			&data.Updated,
		)
		summaries = append(summaries, data)
	}

	return summaries
}

func getApprovalBisnis(dbdefault *sql.DB) (approval ApprovalBisnis) {

	query := `SELECT maker_id, approval_id from ref_maker_aproval_bisnis`

	err := dbdefault.QueryRow(query).Scan(&approval.MakerId, &approval.ApprovalId)

	fmt.Println(query)

	if err == sql.ErrNoRows {
		return approval
	}

	return approval
}

func dioServer(dbdefault *sql.DB, models DioModel) bool {
	var params = url.Values{}
	params.Set("id_template", models.Id_template)
	params.Set("branch_code_penerima", models.Branch_code_penerima)
	params.Set("orgeh_penerima", models.Orgeh_penerima)
	params.Set("branch_code_tindasan", models.Branch_code_tindasan)
	params.Set("orgeh_tindasan", models.Orgeh_tindasan)
	params.Set("pn_penerima", models.Pn_penerima)
	params.Set("pn_tindasan", models.Pn_tindasan)
	params.Set("kode_surat", models.Kode_surat)
	params.Set("kerahasiaan", models.Kerahasiaan)
	params.Set("kesegeraan", models.Kesegeraan)
	params.Set("kepada_yth", models.Kepada_yth)
	params.Set("perihal", models.Perihal)
	params.Set("isi_surat", models.Isi_surat)
	params.Set("id_maker", models.Id_maker)
	params.Set("pn_approver", models.Pn_approver)
	params.Set("status_approver", models.Status_approver)
	params.Set("folder_attachment", models.Folder_attachment)
	params.Set("nama_attachment_uploaded", models.Nama_attachment_uploaded)
	params.Set("sla", fmt.Sprint(models.Sla))
	params.Set("surat_keluar_approver", models.Surat_keluar_approver)
	params.Set("app_id", models.App_id)
	params.Set("password", models.Password)
	params.Set("uker_edit", models.Uker_edit)
	params.Set("jabatan_edit", models.Jabatan_edit)

	var payload = bytes.NewBufferString(params.Encode())
	client := http.DefaultClient
	req, err := http.NewRequest("POST", os.Getenv("WS_DIO"), payload)
	if err != nil {
		fmt.Println(err.Error())
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Cache-Control", "no-cache")

	res, err := client.Do(req)

	if err != nil {
		return false
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return false
	}

	var response responseDio
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return false
	}

	fmt.Println("Status From call DIO")
	fmt.Println(response.ResponseDesc)

	if response.ResponseCode == "05" {
		return false
	}

	return true
}
