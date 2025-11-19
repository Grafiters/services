package generateExcels

import (
	objJSON "encoding/json"
	"fmt"
	"riskmanagement/lib"
	AuditTrail "riskmanagement/models/audittrail"
	briefingModel "riskmanagement/models/briefing"
	coachingModel "riskmanagement/models/coaching"
	models "riskmanagement/models/mstkriteria"
	RiskIndicator "riskmanagement/models/riskindicator"
	verifModels "riskmanagement/models/verifikasi"
	verifRealpinModels "riskmanagement/models/verifikasireportrealisasi"
	"strings"
)

// Query Generator
func RptListVerifikasi(db *lib.Database, request string) (responses []verifModels.VerifikasiReportListResponse, err error) {
	var requestDownload verifModels.VerifikasiReportListRequest

	// Convert JSON string to struct
	if err := objJSON.Unmarshal([]byte(request), &requestDownload); err != nil {
		fmt.Println("Error - unmarshal from json string to json struct:", err)
	}

	sortQuery := ""
	if requestDownload.Sort == "desc" {
		// sortQuery = `verifikasi.created_at DESC`
		sortQuery = `rlv.id DESC`
	} else {
		// sortQuery = `verifikasi.created_at ASC`
		sortQuery = `rlv.id ASC`
	}

	query := db.DB.Table("report_list_verifikasi rlv").
		Select(`
			rlv.id,
			rlv.periode,
			rlv.BRANCH,
			rlv.BRDESC,
			rlv.MBDESC,
			rlv.RGDESC,
			rlv.no_pelaporan,
			rlv.aktifitas,
			rlv.sub_aktifitas,
			rlv.informasi_lain,
			rlv.status_perbaikan_konsolidasi,
			rlv.maker,
			rlv.risk_issue_code,
			rlv.risk_issue,
			rlv.risk_indicator,
			rlv.risk_control,
			rlv.hasil_verifikasi,
			rlv.jumlah_data_yg_diverifikasi,
			rlv.butuh_perbaikan,
			rlv.jumlah_data_yg_harus_diperbaiki,
			rlv.rtl_user,
			rlv.status_perbaikan_selesai,
			rlv.status_perbaikan_proses,
			rlv.batas_waktu_perbaikan,
			CASE
				WHEN rlv.indikasi_fraud = 1 THEN "Ya"
				ELSE "Tidak"
			END 'indikasi_fraud',
			rlv.filename,
			rlv.filepath`).
		Order(sortQuery)

	if requestDownload.NoPelaporan != "" {
		query = query.Where("rlv.no_pelaporan = ?", requestDownload.NoPelaporan)
	}

	if requestDownload.BrcUrc != "Semua" && requestDownload.BrcUrc != "" {
		query = query.Where("rlv.maker LIKE ?", fmt.Sprintf("%%%s%%", requestDownload.BrcUrc))
	}

	if requestDownload.REGION != "all" {
		query = query.Where("rlv.REGION = ?", requestDownload.REGION)
	}

	if requestDownload.MAINBR != "all" {
		mainbrs := strings.Split(requestDownload.MAINBR, ",")
		query = query.Where("rlv.MAINBR in (?)", mainbrs)
	}

	if requestDownload.BRANCH != "all" {
		branches := strings.Split(requestDownload.BRANCH, ",")
		query = query.Where("rlv.BRANCH in (?)", branches)
	}

	if requestDownload.RiskIssueID != "all" {
		query = query.Where("rlv.risk_issue_id = ?", requestDownload.RiskIssueID)
	}

	if requestDownload.RiskIndicator != "all" {
		query = query.Where("rlv.risk_indicator = ?", requestDownload.RiskIndicator)
	}

	if requestDownload.IndikasiFraud != "all" {
		query = query.Where("rlv.indikasi_fraud = ?", requestDownload.IndikasiFraud)
	}

	if requestDownload.Status != "" && requestDownload.Status != "all" {
		query = query.Where("rlv.status = ?", requestDownload.Status)
	}

	if requestDownload.StartDate != "" && lib.FixEndDate(requestDownload.EndDate) != "" {
		query = query.Where("rlv.periode BETWEEN ? AND ?", requestDownload.StartDate, lib.FixEndDate(requestDownload.EndDate))
	}

	err = query.Scan(&responses).Error

	// for i := range responses {
	// 	if responses[i].ButuhPerbaikan == "Tidak" {
	// 		responses[i].PresentasePerbaikan = (responses[i].StatusPerbaikanSelesai / responses[i].JumlahDataYgDiverifikasi) * 100
	// 	} else {
	// 		responses[i].PresentasePerbaikan = (responses[i].StatusPerbaikanSelesai / responses[i].JumlahDataYgHarusDiperbaiki) * 100
	// 	}
	// }

	return responses, err
}

func RptListBriefing(db *lib.Database, request string) (responses []briefingModel.BriefingReportListResponse, err error) {
	var requestDownload briefingModel.BriefingReportListRequest

	if err := objJSON.Unmarshal([]byte(request), &requestDownload); err != nil {
		fmt.Println("Error - unmarshal from json string to json struct:", err)
	}

	sorting := ""

	if requestDownload.Sort == "desc" {
		sorting = "rlb.id DESC"
	} else {
		sorting = "rlb.id ASC"
	}

	query := db.DB.Table("report_list_briefing rlb ").
		Select(`
			rlb.id,
			rlb.BRANCH,
			rlb.BRDESC,
			rlb.MBDESC,
			rlb.RGDESC,
			rlb.no_pelaporan,
			rlb.judul_materi,
			rlb.risk_event,
			rlb.rincian_materi,
			rlb.aktivitas,
			rlb.jumlah_peserta,
			rlb.jenis_peserta,
			rlb.jabatan_peserta,
			rlb.peserta,
			rlb.maker_id,
			rlb.status
		`).
		Order(sorting)
	// Where("briefing.maker_id = ?", requestDownload.Pernr)

	if requestDownload.NoPelaporan != "" {
		query = query.Where("rlb.no_pelaporan = ?", requestDownload.NoPelaporan)
	}

	if requestDownload.REGION != "all" {
		query = query.Where("rlb.REGION = ?", requestDownload.REGION)
	}

	if requestDownload.MAINBR != "all" {
		mainbrs := strings.Split(requestDownload.MAINBR, ",")
		query = query.Where("rlb.MAINBR in (?)", mainbrs)
	}

	if requestDownload.BRANCH != "all" {
		branches := strings.Split(requestDownload.BRANCH, ",")
		query = query.Where("rlb.BRANCH in (?)", branches)
	}

	if requestDownload.ActivityID != "all" {
		query = query.Where("rlb.aktivity_id LIKE ?", fmt.Sprintf("%%%s%%", requestDownload.ActivityID))
	}

	if requestDownload.JudulMateri != "" {
		query = query.Where("rlb.judul_materi LIKE ?", fmt.Sprintf("%%%s%%", requestDownload.JudulMateri))
	}

	if requestDownload.Status != "" && requestDownload.Status != "Semua" {
		query = query.Where("rlb.status = ?", requestDownload.Status)
	}

	if requestDownload.StartDate != "" && lib.FixEndDate(requestDownload.EndDate) != "" {
		query = query.Where(`rlb.periode BETWEEN ? AND ?`, requestDownload.StartDate, lib.FixEndDate(requestDownload.EndDate))
	}

	// execute the query
	err = query.Scan(&responses).Error

	return responses, err
}

func RptListCoaching(db *lib.Database, request string) (responses []coachingModel.CoachingReportListResponse, err error) {
	var requestDownload coachingModel.CoachingReportListRequest

	if err := objJSON.Unmarshal([]byte(request), &requestDownload); err != nil {
		fmt.Println("Error - unmarshal from json string to json struct:", err)
	}

	sorting := ""

	if requestDownload.Sort == "desc" {
		sorting = `rlc.id DESC`
	} else {
		sorting = `rlc.id ASC`
	}

	query := db.DB.Table("report_list_coaching rlc ").
		Select(`
				rlc.id,
				rlc.BRANCH,
				rlc.BRDESC,
				rlc.MBDESC,
				rlc.RGDESC,
				rlc.no_pelaporan,
				rlc.judul_materi,
				rlc.rincian_materi,
				rlc.jumlah_peserta,
				rlc.jenis_peserta,
				rlc.jabatan_peserta,
				rlc.peserta,
				rlc.maker_id,
				rlc.aktifitas,
				rlc.sub_aktifitas,
				rlc.isu_risiko,
				rlc.risk_indicator,
				rlc.status
			`).
		Order(sorting)

	if requestDownload.NoPelaporan != "" {
		query = query.Where("rlc.no_pelaporan = ?", requestDownload.NoPelaporan)
	}

	if requestDownload.REGION != "all" {
		query = query.Where("rlc.REGION = ?", requestDownload.REGION)
	}

	if requestDownload.MAINBR != "all" {
		mainbrs := strings.Split(requestDownload.MAINBR, ",")
		query = query.Where("rlc.MAINBR in (?)", mainbrs)
	}

	if requestDownload.BRANCH != "all" {
		branches := strings.Split(requestDownload.BRANCH, ",")
		if len(branches) > 1 {
			fmt.Println("Masuk 1")
			query = query.Where("rlc.BRANCH in (?)", branches)
		} else {
			fmt.Println("Masuk 2")
			query = query.Where("rlc.BRANCH = ?", requestDownload.BRANCH)
		}
	}

	if requestDownload.ActivityID != "all" {
		query = query.Where("rlc.activity_id = ?", requestDownload.ActivityID)
	}

	if requestDownload.RiskIssueID != "all" {
		query = query.Where("rlc.isu_risiko_id LIKE ?", fmt.Sprintf("%%%s%%", requestDownload.RiskIssueID))
	}

	if requestDownload.JudulMateri != "" {
		query = query.Where("rlc.judul_materi LIKE ?", fmt.Sprintf("%%%s%%", requestDownload.JudulMateri))
	}

	if requestDownload.Status != "" && requestDownload.Status != "Semua" {
		query = query.Where("rlc.status = ?", requestDownload.Status)
	}

	if requestDownload.StartDate != "" && lib.FixEndDate(requestDownload.EndDate) != "" {
		query = query.Where(`rlc.periode BETWEEN ? AND ?`, requestDownload.StartDate, lib.FixEndDate(requestDownload.EndDate))
	}

	// execute the query
	err = query.Scan(&responses).Error

	return responses, err
}

func RptAuditTrail(db *lib.Database, request string) (responses []AuditTrail.AuditTrailResponse, err error) {
	var requestDownload AuditTrail.FilterAudit

	if err := objJSON.Unmarshal([]byte(request), &requestDownload); err != nil {
		fmt.Println("Error - unmarshal from json string to json struct:", err)
	}

	query := db.DB.Table(`audit_trail`).
		Select(`
			id,
			DATE_FORMAT(tanggal , '%Y-%m-%d') 'tanggal',
			pn,
			nama_brc_urc,
			RGDESC 'Kanwil',
			CONCAT(MAINBR,' - ',MBDESC) 'Kanca',
			CONCAT(BRANCH,' - ',BRDESC) 'Uker',
			no_pelaporan,
			aktifitas,
			ip_address,
			lokasi
		`).Where("tanggal >= ? AND tanggal <= ?", requestDownload.StartDate, lib.FixEndDate(requestDownload.EndDate))

	if requestDownload.PERNR != "" && requestDownload.PERNR != "Semua" {
		query = query.Where(`pn = ?`, requestDownload.PERNR)
	}

	if requestDownload.Aktifitas != "all" {
		query = query.Where("aktifitas = ?", requestDownload.Aktifitas)
	}

	if requestDownload.REGION != "all" {
		query = query.Where("REGION = ?", requestDownload.REGION)
	}

	if requestDownload.MAINBR != "all" {
		query = query.Where("MAINBR = ?", requestDownload.MAINBR)
	}

	if requestDownload.BRANCH != "all" {
		query = query.Where("BRANCH = ?", requestDownload.BRANCH)
	}

	err = query.Scan(&responses).Error

	return responses, err
}

func RptRekapitulasiBCV(db *lib.Database, request string) (responses []verifModels.RptRekapitulasiBCVResponse, err error) {
	var requestDownload verifModels.RptRekapitulasiBCVRequest

	if err := objJSON.Unmarshal([]byte(request), &requestDownload); err != nil {
		fmt.Println("Error - unmarshal from json string to json struct:", err)
	}

	query := db.DB.Table(`rpt_rekapitulasi_bcv rrb`).
		Select(`
			rrb.pernr,
			rrb.brc,
			rrb.BRANCH,
			rrb.BRDESC,
			rrb.MBDESC,
			rrb.RGDESC,
			SUM(rrb.bdraft) 'b_draft',
			SUM(rrb.bfinish) 'b_finish',
			SUM(rrb.btotal) 'b_total',
			SUM(rrb.cdraft) 'c_draft',
			SUM(rrb.cfinish) 'c_finish',
			SUM(rrb.ctotal) 'c_total',
			SUM(rrb.vdraft) 'v_draft',
			SUM(rrb.vfinish) 'v_finish',
			SUM(rrb.vtotal) 'v_total'
		`).Group(`rrb.BRANCH`).Order(`rrb.pernr`)

	if requestDownload.REGION != "all" {
		query = query.Where("rrb.REGION = ?", requestDownload.REGION)
	}

	if requestDownload.MAINBR != "all" {
		mainbrs := strings.Split(requestDownload.MAINBR, ",")
		query = query.Where("rrb.MAINBR in (?)", mainbrs)
	}

	if requestDownload.BRANCH != "all" {
		branches := strings.Split(requestDownload.BRANCH, ",")
		query.Where("rrb.BRANCH in (?)", branches)
	}

	if requestDownload.BRC != "Semua" {
		query = query.Where("rrb.pn = ?", requestDownload.BRC)
	}

	if requestDownload.StartDate != "" && requestDownload.EndDate != "" {
		query = query.Where("rrb.Tanggal BETWEEN ? AND ?", requestDownload.StartDate, lib.FixEndDate(requestDownload.EndDate))
	}

	err = query.Scan(&responses).Error

	return responses, err
}

func RptRekomendasiRiskRequest(db *lib.Database, request string) (responses []verifModels.RptRekomendasiRiskResponse, err error) {
	var requestDownload verifModels.RptRekomendasiRiskRequest

	if err := objJSON.Unmarshal([]byte(request), &requestDownload); err != nil {
		fmt.Println("Error - unmarshal from json string to json struct:", err)
	}

	query := db.DB.Table(`rpt_rekomendasi_risk rrr`)

	switch requestDownload.JenisData {
	case "Risk Event":
		query.Select(`
				rrr.nama_risk 'risk_event',
				rrr.modul 'module',
				SUM(rrr.jumlah) 'count'
			`).
			Where(`rrr.Tanggal BETWEEN ? AND ?`, requestDownload.StartDate, lib.FixEndDate(requestDownload.EndDate)).
			Where("rrr.jenis_risk = ?", strings.ToLower(requestDownload.JenisData)).
			Group(`rrr.nama_risk, rrr.modul`).
			Order(`rrr.jumlah DESC`)
	case "Risk Indicator":
		query.Select(`
				rrr.nama_risk 'risk_indicator',
				rrr.modul 'module',
				SUM(rrr.jumlah) 'count'
			`).
			Where(`rrr.Tanggal BETWEEN ? AND ?`, requestDownload.StartDate, lib.FixEndDate(requestDownload.EndDate)).
			Where("rrr.jenis_risk = ?", strings.ToLower(requestDownload.JenisData)).
			Group(`rrr.nama_risk, rrr.modul`).
			Order(`rrr.jumlah DESC`)
	default:
		return nil, fmt.Errorf("Query tidak tersedia")
	}

	err = query.Scan(&responses).Error

	return responses, err
}

func GetDataThreshold(db *lib.Database, request string) (responses []RiskIndicator.ThresholdIndicator, err error) {
	var requestDownload RiskIndicator.RiskIndicatorGetOne

	if err := objJSON.Unmarshal([]byte(request), &requestDownload); err != nil {
		fmt.Println("Error - unmarshal from json string to json struct:", err)
	}

	query := db.DB.Table(`risk_indicator ri`).
		Select(`
			ri.id 'index',
			ri.risk_indicator_code 'id',
			ri.risk_indicator 'key_risk_indicator',
			a.name 'aktivitas',
			p.product 'produk',
			SUBSTRING(ri.sifat, 1, INSTR(ri.sifat, '/') - 1) 'jenis_indikator',
			"" AS 'indikasi_risiko',
			ri.deskripsi 'deskripsi',
			ri.sla_verifikasi 'sla_verifikasi',
			ri.sla_tindak_lanjut 'sla_tl',
			"" AS 'risk_awareness',
			"" As 'data_source',
			ri.satuan 'parameter',
			ri.status_indikator 'status_indikator',
			CASE 
				WHEN ri.status = 1 THEN 'aktif'
				ELSE 'non aktif '
			END 'is_aktif'`).
		Joins(`LEFT JOIN activity a ON a.id = ri.activity_id`).
		Joins(`LEFT JOIN product p ON p.id = ri.product_id`).
		Where(`ri.id = ?`, requestDownload.ID)

	err = query.Scan(&responses).Error

	return responses, err
}

func GetThreshold(db *lib.Database, id int64) (responses []RiskIndicator.MapThresholdResponse, err error) {
	return responses, db.DB.Where("id_indicator = ?", id).Find(&responses).Error
}

func RptRealisasiKreditList(db *lib.Database, request string) (response []verifRealpinModels.ReportRealisasiKreditListResponseRaw, err error) {
	var requestDownload verifRealpinModels.ReportRealisasiKreditListRequest

	if err := objJSON.Unmarshal([]byte(request), &requestDownload); err != nil {
		fmt.Println("Error - unmarshal from json string to json struct:", err)
	}

	query := db.DB.Table(`verifikasi_data_realisasi vdr`).
		Select(`vrk.no_pelaporan AS no_pelaporan,
					vrk.REGION AS REGION,
					vrk.RGDESC AS RGDESC,
					vrk.MAINBR AS MAINBR,
					vrk.MBDESC AS MBDESC,
					vrk.BRANCH AS BRANCH,
					vrk.BRDESC AS BRDESC,
					vrk.activity_id AS activity_id,
					vrk.activity_name AS activity_name,
					vrk.product_id AS product_id,
					vrk.product_name AS product_name,
					vrk.periode_data AS periode_data,
					vrk.restruck_flag AS restruck_flag,
					vrk.butuh_perbaikan AS butuh_perbaikan,
					vrk.kriteria_data AS kriteria_data,
					vrk.hasil_verifikasi AS hasil_verifikasi,
					vrk.kunjungan_nasabah AS kunjungan_nasabah,
					vrk.tgl_kunjungan AS tgl_kunjungan,
					JSON_UNQUOTE(JSON_EXTRACT(LOWER(vdr.data_realisasi), '$.segment')) AS 'segment',
					vrk.created_id AS created_id,
					vrk.created_desc AS created_desc,
					vdr.data_realisasi AS data_realisasi,
					vdr.status_verifikasi AS status_verifikasi`).
		Joins(`LEFT JOIN verifikasi_realisasi_kredit vrk ON vrk.id = vdr.verifikasi_id`).
		Where(`vrk.deleted = 0`)

	if requestDownload.Sort != "" {
		query = query.Order(`vdr.id` + requestDownload.Sort)
	} else {
		query = query.Order(`vdr.id DESC`)
	}

	if requestDownload.REGION != "" && requestDownload.REGION != "all" {
		query = query.Where(`vrk.REGION = ?`, requestDownload.REGION)
	}

	if requestDownload.MAINBR != "" && requestDownload.MAINBR != "all" {
		// query = query.Where(`vrk.MAINBR = ?`, requestDownload.MAINBR)
		mainbrs := strings.Split(requestDownload.MAINBR, ",")
		query = query.Where(`vrk.MAINBR in (?)`, mainbrs)
	}

	//Filter BRANCH
	if requestDownload.BRANCH != "" && requestDownload.BRANCH != "all" {
		// query = query.Where(`vrk.BRANCH = ?`, requestDownload.BRANCH)
		branches := strings.Split(requestDownload.BRANCH, ",")
		query = query.Where(`vrk.BRANCH in (?)`, branches)
	}

	//Filter NoPelaporan
	if requestDownload.NoPelaporan != "" {
		query = query.Where(`vrk.no_pelaporan = ?`, requestDownload.NoPelaporan)
	}

	//Filter Product
	if requestDownload.Product != "" && requestDownload.Product != "all" {
		query = query.Where("vrk.product_id = ?", requestDownload.Product)
	}

	//Filter Criteria
	if requestDownload.Criteria != "" && requestDownload.Criteria != "all" {
		query = query.Where(`JSON_CONTAINS(vrk.kriteria_data, ?)`, requestDownload.Criteria)
	}

	//Filter Segment
	if requestDownload.Segment != "" && requestDownload.Segment != "all" {
		query = query.Where(`JSON_UNQUOTE(JSON_EXTRACT(LOWER(vdr.data_realisasi), '$.segment')) = ?`, strings.ToLower(requestDownload.Segment))
	}

	//Filter Is Verified
	if requestDownload.StatusVerifikasi != "" && requestDownload.StatusVerifikasi != "all" {
		query = query.Where(`vdr.status_verifikasi = ?`, requestDownload.StatusVerifikasi)
	}

	if requestDownload.ButuhPerbaikan != "" && requestDownload.ButuhPerbaikan != "all" {
		query = query.Where(`vrk.butuh_perbaikan = ?`, requestDownload.ButuhPerbaikan)
	}

	//Filter Is Verified
	if requestDownload.IndikasiFraud != "" && requestDownload.IndikasiFraud != "all" {
		query = query.Where(`vrk.indikasi_fraud = ?`, requestDownload.IndikasiFraud)
	}

	if requestDownload.StartDate != "" && requestDownload.EndDate != "" {
		query = query.Where(`DATE(vrk.created_at) >= ?`, requestDownload.StartDate).Where(`DATE(vrk.created_at) <= ?`, requestDownload.EndDate)
	}

	err = query.Scan(&response).Error

	return response, err
}

func RptRealisasiKreditSummary(db *lib.Database, request string) (response []verifRealpinModels.ReportRealisasiKreditSummaryDownloadResponseRaw, err error) {
	var requestDownload verifRealpinModels.ReportRealisasiKreditSummaryRequest

	if err := objJSON.Unmarshal([]byte(request), &requestDownload); err != nil {
		fmt.Println("Error - unmarshal from json string to json struct:", err)
	}

	query := db.DB.Table(`verifikasi_data_realisasi vdr`).
		Select(`vrk.product_id AS product_id,
			 COUNT(vrk.id) AS total_verifikasi,
			vrk.product_name AS product_name, 
			vrk.created_id AS created_id, 
			vrk.created_desc AS created_desc, 
			vrk.REGION AS REGION, 
			CONCAT(vrk.REGION,' - ',vrk.RGDESC) AS RGDESC, 
			vrk.MAINBR AS MAINBR, 
			CONCAT(vrk.MAINBR,' - ',vrk.MBDESC) AS MBDESC, 
			vrk.BRANCH AS BRANCH, 
			CONCAT(vrk.BRANCH,' - ',vrk.BRDESC) AS BRDESC, 
			IF(vdr.status_verifikasi = 1, 'Ya', 'Tidak') AS status_verifikasi,
			(SELECT JSON_OBJECTAGG(id,criteria) FROM mst_kriteria WHERE (status = 1 AND active_date <= NOW() OR status = 0 AND inactive_date >= NOW())) AS mst_kriteria,
			vdr.data_realisasi AS data_realisasi,
			COUNT(CASE WHEN vrk.butuh_perbaikan = 0 THEN 1 END) AS efektif,
			COUNT(CASE WHEN vrk.butuh_perbaikan = 1 THEN 1 END) AS non_efektif,
			CONCAT('[',GROUP_CONCAT(REPLACE(REPLACE(vrk.kriteria_data, ']', ''), '[', '') SEPARATOR ','),']') AS kriteria_data`).
		Joins(`LEFT JOIN verifikasi_realisasi_kredit vrk ON vdr.verifikasi_id = vrk.id`).
		Where(`LOWER(vrk.action) = ?`, `selesai`)

	fmt.Println(query)

	if requestDownload.Sort != "" {
		query = query.Order(`vdr.id ` + requestDownload.Sort)
	} else {
		query = query.Order(`vdr.id DESC`)
	}

	//Filter REGION
	if requestDownload.REGION != "" && requestDownload.REGION != "all" {
		query = query.Where(`vrk.REGION = ?`, requestDownload.REGION)
	}

	//Filter MAINBR
	if requestDownload.MAINBR != "" && requestDownload.MAINBR != "all" {
		mainbrs := strings.Split(requestDownload.MAINBR, ",")
		query = query.Where(`vrk.MAINBR in (?)`, mainbrs)
	}

	//Filter BRANCH
	if requestDownload.BRANCH != "" && requestDownload.BRANCH != "all" {
		branches := strings.Split(requestDownload.BRANCH, ",")

		query = query.Where(`vrk.BRANCH in (?)`, branches)
	}

	if requestDownload.Product != nil {
		query = query.Where(`vrk.product_id IN (?)`, requestDownload.Product)
	}

	if requestDownload.StartDate != "" && requestDownload.EndDate != "" {
		query = query.Where(`DATE(vrk.created_at) >= ?`, requestDownload.StartDate).Where(`DATE(vrk.created_at) <= ?`, requestDownload.EndDate)
	}

	if requestDownload.GroupBy != nil {
		//option for group by
		optionGroupBy := map[string]string{
			"regional-office":   "REGION",
			"branch-office":     "MAINBR",
			"unit-kerja":        "BRANCH",
			"produk":            "product_name",
			"pn-brc-urc":        "created_id",
			"status-verifikasi": "status_verifikasi",
		}

		for _, value := range requestDownload.GroupBy {
			if optionGroupBy[value] != "" {
				query = query.Group(optionGroupBy[value])
			}
		}
	}

	err = query.Scan(&response).Error
	return response, err
}

// GetAll implements MstKriteriaDefinition
func GetMstKriteria(db *lib.Database, request string) (responses []models.MstKriteriaResponse, err error) {
	var requestDownload verifRealpinModels.ReportRealisasiKreditSummaryRequest

	if err := objJSON.Unmarshal([]byte(request), &requestDownload); err != nil {
		fmt.Println("Error - unmarshal from json string to json struct:", err)
	}

	query := db.DB.Table("mst_kriteria").Where("(status = 1 AND active_date <= NOW() OR status = 0 AND inactive_date >= NOW())")
	if requestDownload.Criteria != "" && requestDownload.Criteria != "all" {
		query.Where("id = ?", requestDownload.Criteria)
	}

	err = query.Find(&responses).Error
	return responses, err
}

func GetAllKriteria(db *lib.Database) (response []verifRealpinModels.ListKriteria, err error) {
	query := db.DB.Table("mst_kriteria").
		Select("id, criteria 'kriteria'").
		Where("(status = 1 AND active_date <= NOW() OR status = 0 AND inactive_date >= NOW())").
		Order("id ASC")

	err = query.Scan(&response).Error

	return response, err
}

func GetKriteriaByPeriodeList(db *lib.Database, request string) (responses []verifRealpinModels.ListKriteria, err error) {
	var requestDownload verifRealpinModels.ReportRealisasiKreditListRequest

	if err := objJSON.Unmarshal([]byte(request), &requestDownload); err != nil {
		fmt.Println("Error - unmarshal from json string to json struct:", err)
	}

	query := db.DB.Table(`mst_kriteria_history`).
		Select("id_criteria, criteria 'kriteria'").
		Where(`DATE(active_date) >= ? AND DATE(active_date) <= ?`, requestDownload.StartDate, requestDownload.EndDate).
		Group(`id_criteria`)

	err = query.Find(&responses).Error
	return responses, err
}

func GetKriteriaByPeriodeSummary(db *lib.Database, request string) (responses []verifRealpinModels.ListKriteria, err error) {
	var requestDownload verifRealpinModels.ReportRealisasiKreditSummaryRequest

	if err := objJSON.Unmarshal([]byte(request), &requestDownload); err != nil {
		fmt.Println("Error - unmarshal from json string to json struct:", err)
	}

	query := db.DB.Table(`mst_kriteria_history`).
		Select("id_criteria, criteria 'kriteria'").
		Where(`DATE(active_date) >= ? AND DATE(active_date) <= ?`, requestDownload.StartDate, requestDownload.EndDate).Group(`id_criteria`)

	err = query.Find(&responses).Error
	return responses, err
}
