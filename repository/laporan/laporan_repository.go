package laporan

import (
	"math"
	"riskmanagement/lib"
	models "riskmanagement/models/laporan"
	"strconv"
	"strings"
	"time"

	"gitlab.com/golang-package-library/logger"
)

var (
	timeNow = lib.GetTimeNow("timestime")
)

type LaporanDefinition interface {
	GetTasklist(request models.HistoriTaskDataVerifikasiPagianted) (responses []models.HistoriTaskDataVerifikasiResult, totalRow int, totalData int, err error)
	GetTasklistDownload(request models.HistoriTaskDataVerifikasiDownload) (responses []models.HistoriTaskDataVerifikasiResult, err error)
	GetListVerifikasiByTaskID(request models.HistoriTaskDataVerifikasiDetailRequest) (response []models.HistoriTaskDataVerifikasiDetailResult, err error)
	GetPerhitunganPersentasePenyelesaianPerPekerja(request models.PerhitunganPersentasePenyelesaianPagianted) (responses []models.LaporanPerPekerjaQueryResult, totalRow int, totalData int, err error)
	GetPerhitunganPersentasePenyelesaianPerPekerjaDownload(request models.PerhitunganPersentasePenyelesaianDownload) (responses []models.LaporanPerPekerjaQueryResult, err error)
	GetPerhitunganPersentasePenyelesaianPerUker(request models.PerhitunganPersentasePenyelesaianPagianted) (responses []models.LaporanPerUkerResult, totalRow int, err error)
	GetPerhitunganPersentasePenyelesaianPerUkerDownload(request models.PerhitunganPersentasePenyelesaianDownload) (responses []models.LaporanPerUkerResult, err error)
	GetPerhitunganPersentasePenyelesaianPerPekerjaUker(request models.PerhitunganPersentasePenyelesaianPagianted) (responses []models.LaporanPerPekerjaUkerResult, totalRow int, err error)
	GetPerhitunganPersentasePenyelesaianPerPekerjaUkerDownload(request models.PerhitunganPersentasePenyelesaianDownload) (responses []models.LaporanPerPekerjaUkerResult, err error)
	GetRiskEventOnTaskList() (response []models.RiskEventOnTaskList, err error)
	GetMonitoringJob(request models.JobMonitoringRequest) (responses []models.JobMonitoringResponse, totalRow int, totalData int, err error)
	GetNamaJob(request *models.SearchNamaJobReq) (responses []models.SearchNamaJobRes, err error)
	GetActivityDaily(request models.ActivityDailyRequest) (responses []models.ActivityDaily, totalRow int, totalData int, err error)
	GetUkerList(request models.UkerListRequest) (responses []models.UkerListResponse, err error)
	GetTasklistPersentase(request models.PersentaseTotalRequest) (total float64, err error)
	GetActivityDailyDetail(request models.ActivityDailyDetailRequest) (responses []models.ActivityDailyDetail, totalRow int, totalData int, err error)
	GetPerhitunganBriefingDownload(request models.PerhitunganPersentasePenyelesaianDownload) (responses []models.LaporanBriefingResponse, err error)
	GetPerhitunganCoaching(request models.PerhitunganPersentasePenyelesaianPagianted) (responses []models.LaporanCoachingResponse, totalRow int, totalData int, err error)
	GetPerhitunganCoachingDownload(request models.PerhitunganPersentasePenyelesaianDownload) (responses []models.LaporanCoachingResponse, err error)
	GetPerhitunganVerifikasi(request models.PerhitunganPersentasePenyelesaianPagianted) (responses []models.LaporanVerifikasiResponse, totalRow int, totalData int, err error)
	GetPerhitunganVerifikasiDownload(request models.PerhitunganPersentasePenyelesaianDownload) (responses []models.LaporanVerifikasiResponse, err error)
	GetPerhitunganBriefing(request models.PerhitunganPersentasePenyelesaianPagianted) (responses []models.LaporanBriefingResponse, totalRow int, totalData int, err error)
}

type LaporanRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewLaporanRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) LaporanDefinition {
	return LaporanRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: 0,
	}
}

// GetPerhitunganBriefing implements LaporanDefinition.
func (l LaporanRepository) GetPerhitunganBriefing(request models.PerhitunganPersentasePenyelesaianPagianted) (responses []models.LaporanBriefingResponse, totalRow int, totalData int, err error) {
	db := l.db.DB

	//20/09/2024
	subQuery := db.Table(`tasklists_history`).Select(`*`).Group(`tasklist_id, branch`)

	query := db.Table("tasklists t").
		Select(`t.id , 
		th.pernr as "pn",
		th.name as "nama",
		tu.RGDESC "kanwil",
		tu.MBDESC  "kanca",
		tu.BRDESC "uker",
		t.task_type_name "jenis_task",
		t.kegiatan "kegiatan",
		COALESCE(th.activity, "-") "aktifitas",
		t.product_name "produk",
		t.risk_issue "risk_issue",
		t.risk_indicator "indikator",
		CASE 
			WHEN t.period = 'Custom' THEN DATE_FORMAT(t.start_date,"%d-%b-%Y")
			WHEN t.period = 'Monthly' THEN DATE_FORMAT(t.start_date,"%d")
			ELSE "-"
		END as "tanggal_mulai",
		CASE 
			WHEN t.period = 'Custom' THEN DATE_FORMAT(t.end_date,"%d-%b-%Y")
			WHEN t.period = 'Monthly' THEN DATE_FORMAT(t.end_date,"%d")
			ELSE "-"
		END as "tanggal_selesai",
		COALESCE(kegiatan.jumlah, 0) as "jumlah_kegiatan_dilakukan"`).
		Joins(`join tasklists_uker tu on t.id = tu.tasklist_id `).
		Joins(`join (?) th on t.id = th.tasklist_id and tu.branch = th.branch`, subQuery).
		Joins(`LEFT JOIN (
					SELECT b.REGION, b.MAINBR, b.BRANCH, bm.activity_id, bm.product_id, count(b.id) "jumlah" from briefing b 
					JOIN briefing_materis bm ON b.id = bm.briefing_id
					WHERE b.deleted = 0 
					AND (b.status = "01a" AND b.action = "Selesai")
					AND (b.status = "02b" AND (b.action = "Update" OR b.action ="Selesai"))
					GROUP BY b.REGION, b.MAINBR, b.BRANCH, bm.activity_id, bm.product_id
				)kegiatan ON t.activity_id = kegiatan.activity_id
					and t.product_id = kegiatan.product_id
					and tu.REGION = kegiatan.REGION
					and tu.MAINBR = kegiatan.MAINBR
					and tu.BRANCH = kegiatan.BRANCH`).
		Where(`t.kegiatan = "Briefing"`).
		Group("t.id").
		Order("t.id ASC")
	//tanpa group
	// .Group(`t.id`)

	if request.Nama != "" {
		// query = query.Where("t.maker_id = ? ", request.Nama)
		query = query.Where("th.pernr = ?", request.Nama)
	}

	if request.Kanwil != "all" && request.Kanwil != "" {
		query = query.Where("tu.REGION = ? ", request.Kanwil)
	}

	if request.Kanca != "" && request.Kanca != "all" {
		query = query.Where("tu.MAINBR = ? ", request.Kanca)
	}

	if request.Uker != "" && request.Uker != "all" {
		branches := strings.Split(request.Uker, ",")
		query = query.Where("tu.BRANCH in (?) ", branches)
	}

	if strings.ToLower(request.Aktifitas) != "0" && strings.ToLower(request.Aktifitas) != "" {
		query = query.Where("t.activity_id = ? ", request.Aktifitas)
	}

	if strings.ToLower(request.Produk) != "0" && strings.ToLower(request.Produk) != "" {
		query = query.Where("t.product_id = ? ", request.Produk)
	}

	if strings.ToLower(request.RiskIssue) != "0" && strings.ToLower(request.RiskIssue) != "" {
		query = query.Where("t.risk_issue_id = ? ", request.RiskIssue)
	}

	if strings.ToLower(request.Indikator) != "" {
		query = query.Where("t.risk_indicator_id = ? ", request.Indikator)
	}

	if strings.ToLower(request.Indikator) == "0" {
		query = query.Where("LOWER(t.risk_indicator) like ? ", "%"+request.IndikatorOther+"%")
	}

	if strings.ToLower(request.Status) != "" {
		if request.Status == "aktif" {
			query = query.Where("t.status = ? ", request.Status)
		} else {
			query = query.Where("t.status != ? ", "aktif")
		}

	}

	if strings.ToLower(request.JenisTaks) != "" {
		query = query.Where("t.task_type = ? ", request.JenisTaks)
	}

	var count int64
	query.Count(&count)

	totalRow = int(count)

	if request.Limit != 0 {
		query = query.Limit(request.Limit)
	}

	if request.Offset != 0 {
		query = query.Offset(request.Offset)
	}

	err = query.Scan(&responses).Error

	result := float64(totalRow) / float64(request.Limit)
	resultFinal := int(math.Ceil(result))

	if err != nil {
		l.logger.Zap.Error(err)
		return responses, 0, 0, err
	}

	l.logger.Zap.Info(responses)
	return responses, resultFinal, totalRow, err
}

// GetPerhitunganBriefingDownload implements LaporanDefinition.
func (l LaporanRepository) GetPerhitunganBriefingDownload(request models.PerhitunganPersentasePenyelesaianDownload) (responses []models.LaporanBriefingResponse, err error) {
	db := l.db.DB
	//20/09/2024
	subQuery := db.Table(`tasklists_history`).Select(`*`).Group(`tasklist_id, branch`)

	query := db.Table("tasklists t").
		Select(`t.id , 
			th.pernr as "pn",
			th.name as "nama",
			tu.RGDESC "kanwil",
			tu.MBDESC  "kanca",
			tu.BRDESC "uker",
			t.task_type_name "jenis_task",
			t.kegiatan "kegiatan",
			COALESCE(th.activity, "-") "aktifitas",
			t.product_name "produk",
			t.risk_issue "risk_issue",
			t.risk_indicator "indikator",
			CASE 
				WHEN t.period = 'Custom' THEN DATE_FORMAT(t.start_date,"%d-%b-%Y")
				WHEN t.period = 'Monthly' THEN DATE_FORMAT(t.start_date,"%d")
				ELSE "-"
			END as "tanggal_mulai",
			CASE 
				WHEN t.period = 'Custom' THEN DATE_FORMAT(t.end_date,"%d-%b-%Y")
				WHEN t.period = 'Monthly' THEN DATE_FORMAT(t.end_date,"%d")
				ELSE "-"
			END as "tanggal_selesai",
	COALESCE(kegiatan.jumlah, 0) as "jumlah_kegiatan_dilakukan"`).
		Joins(`join tasklists_uker tu on t.id = tu.tasklist_id `).
		Joins(`join (?) th on t.id = th.tasklist_id and tu.branch = th.branch`, subQuery).
		Joins(`LEFT JOIN (
				SELECT b.REGION, b.MAINBR, b.BRANCH, bm.activity_id, bm.product_id, count(b.id) "jumlah" from briefing b 
				JOIN briefing_materis bm ON b.id = bm.briefing_id
				WHERE b.deleted = 0 
				AND (b.status = "01a" AND b.action = "Selesai")
				AND (b.status = "02b" AND (b.action = "Update" OR b.action ="Selesai"))
				GROUP BY b.REGION, b.MAINBR, b.BRANCH, bm.activity_id, bm.product_id
			)kegiatan ON t.activity_id = kegiatan.activity_id
				and t.product_id = kegiatan.product_id
				and tu.REGION = kegiatan.REGION
				and tu.MAINBR = kegiatan.MAINBR
				and tu.BRANCH = kegiatan.BRANCH`).
		Where(`t.kegiatan = "Briefing"`).
		Group("t.id").
		Order("t.id ASC")

	if request.Nama != "" {
		// query = query.Where("t.maker_id = ? ", request.Nama)
		query = query.Where("th.pernr = ?", request.Nama)
	}

	if request.Kanwil != "all" && request.Kanwil != "" {
		query = query.Where("tu.REGION = ? ", request.Kanwil)
	}

	if request.Kanca != "" && request.Kanca != "all" {
		query = query.Where("tu.MAINBR = ? ", request.Kanca)
	}

	if request.Uker != "" && request.Uker != "all" {
		branches := strings.Split(request.Uker, ",")
		query = query.Where("tu.BRANCH in (?) ", branches)
	}

	if strings.ToLower(request.Aktifitas) != "0" && strings.ToLower(request.Aktifitas) != "" {
		query = query.Where("t.activity_id = ? ", request.Aktifitas)
	}

	if strings.ToLower(request.Produk) != "0" && strings.ToLower(request.Produk) != "" {
		query = query.Where("t.product_id = ? ", request.Produk)
	}

	if strings.ToLower(request.RiskIssue) != "0" && strings.ToLower(request.RiskIssue) != "" {
		query = query.Where("t.risk_issue_id = ? ", request.RiskIssue)
	}

	if strings.ToLower(request.Indikator) != "" {
		query = query.Where("t.risk_indicator_id = ? ", request.Indikator)
	}

	if strings.ToLower(request.Indikator) == "0" {
		query = query.Where("LOWER(t.risk_indicator) like ? ", "%"+request.IndikatorOther+"%")
	}

	if strings.ToLower(request.Status) != "" {
		if request.Status == "aktif" {
			query = query.Where("t.status = ? ", request.Status)
		} else {
			query = query.Where("t.status != ? ", "aktif")
		}

	}

	if strings.ToLower(request.JenisTaks) != "" {
		query = query.Where("t.task_type = ? ", request.JenisTaks)
	}

	err = query.Scan(&responses).Error

	if err != nil {
		l.logger.Zap.Error(err)
		return responses, err
	}

	// l.logger.Zap.Info(responses)
	return responses, err
}

// GetPerhitunganCoaching implements LaporanDefinition.
func (l LaporanRepository) GetPerhitunganCoaching(request models.PerhitunganPersentasePenyelesaianPagianted) (responses []models.LaporanCoachingResponse, totalRow int, totalData int, err error) {
	db := l.db.DB

	//20/09/2024
	subQuery := db.Table(`tasklists_history`).Select(`*`).Group(`tasklist_id, branch`)

	query := db.Table("tasklists t").
		Select(`t.id , 
		th.pernr as "pn",
		th.name as "nama",
		tu.RGDESC "kanwil",
		tu.MBDESC  "kanca",
		tu.BRDESC "uker",
		t.task_type_name "jenis_task",
		t.kegiatan "kegiatan",
		COALESCE(th.activity, "-") "aktifitas", 
		t.product_name "produk",
		t.risk_issue "risk_issue",
		t.risk_indicator "indikator",
		CASE 
			WHEN t.period = 'Custom' THEN DATE_FORMAT(t.start_date,"%d-%b-%Y")
			WHEN t.period = 'Monthly' THEN DATE_FORMAT(t.start_date,"%d")
			ELSE "-"
		END as "tanggal_mulai",
		CASE 
			WHEN t.period = 'Custom' THEN DATE_FORMAT(t.end_date,"%d-%b-%Y")
			WHEN t.period = 'Monthly' THEN DATE_FORMAT(t.end_date,"%d")
			ELSE "-"
		END as "tanggal_selesai",
		COALESCE(kegiatan.jumlah, 0) as "jumlah_kegiatan_dilakukan"`).
		Joins(`join tasklists_uker tu on t.id = tu.tasklist_id `).
		Joins(`join (?) th on t.id = th.tasklist_id and tu.branch = th.branch`, subQuery).
		Joins(`LEFT JOIN (
					SELECT c.REGION, c.MAINBR, c.BRANCH, c.activity_id, c.product_id, ca.risk_issue_id, count(c.id) "jumlah" FROM coaching c 
					JOIN coaching_activity ca ON c.id = ca.coaching_id 
					WHERE c.deleted = 0
					AND (c.status = "01a" AND c.action = "Selesai")
					AND (c.status = "02b" AND (c.action = "Update" OR c.action ="Selesai"))
					GROUP BY c.REGION, c.MAINBR, c.BRANCH, c.activity_id, c.product_id, ca.risk_issue_id
				) kegiatan ON t.activity_id = kegiatan.activity_id
					and t.product_id = kegiatan.product_id
					and t.risk_issue_id = kegiatan.risk_issue_id
					and tu.REGION = kegiatan.REGION
					and tu.MAINBR = kegiatan.MAINBR
					and tu.BRANCH = kegiatan.BRANCH `).
		Where(`t.kegiatan = "Coaching"`).Group("t.id").Order("t.id ASC")

	if request.Nama != "" {
		// query = query.Where("t.maker_id = ? ", request.Nama)
		query = query.Where("th.pernr = ?", request.Nama)
	}

	if request.Kanwil != "all" && request.Kanwil != "" {
		query = query.Where("tu.REGION = ? ", request.Kanwil)
	}

	if request.Kanca != "" && request.Kanca != "all" {
		query = query.Where("tu.MAINBR = ? ", request.Kanca)
	}

	if request.Uker != "" && request.Uker != "all" {
		branches := strings.Split(request.Uker, ",")
		query = query.Where("tu.BRANCH in (?) ", branches)
	}

	if strings.ToLower(request.Aktifitas) != "0" && strings.ToLower(request.Aktifitas) != "" {
		query = query.Where("t.activity_id = ? ", request.Aktifitas)
	}

	if strings.ToLower(request.Produk) != "0" && strings.ToLower(request.Produk) != "" {
		query = query.Where("t.product_id = ? ", request.Produk)
	}

	if strings.ToLower(request.RiskIssue) != "0" && strings.ToLower(request.RiskIssue) != "" {
		query = query.Where("t.risk_issue_id = ? ", request.RiskIssue)
	}

	if strings.ToLower(request.Indikator) != "" {
		query = query.Where("t.risk_indicator_id = ? ", request.Indikator)
	}

	if strings.ToLower(request.Indikator) == "0" {
		query = query.Where("LOWER(t.risk_indicator) like ? ", "%"+request.IndikatorOther+"%")
	}

	if strings.ToLower(request.Status) != "" {
		if request.Status == "aktif" {
			query = query.Where("t.status = ? ", request.Status)
		} else {
			query = query.Where("t.status != ? ", "aktif")
		}

	}

	if strings.ToLower(request.JenisTaks) != "" {
		query = query.Where("t.task_type = ? ", request.JenisTaks)
	}

	var count int64
	query.Count(&count)

	totalRow = int(count)

	if request.Limit != 0 {
		query = query.Limit(request.Limit)
	}

	if request.Offset != 0 {
		query = query.Offset(request.Offset)
	}

	err = query.Scan(&responses).Error

	result := float64(totalRow) / float64(request.Limit)
	resultFinal := int(math.Ceil(result))

	if err != nil {
		l.logger.Zap.Error(err)
		return responses, 0, 0, err
	}

	l.logger.Zap.Info(responses)
	return responses, resultFinal, totalRow, err
}

// GetPerhitunganCoachingDownload implements LaporanDefinition.
func (l LaporanRepository) GetPerhitunganCoachingDownload(request models.PerhitunganPersentasePenyelesaianDownload) (responses []models.LaporanCoachingResponse, err error) {
	db := l.db.DB

	//20/09/2024
	subQuery := db.Table(`tasklists_history`).Select(`*`).Group(`tasklist_id, branch`)

	query := db.Table("tasklists t").
		Select(`t.id , 
		th.pernr as "pn",
		th.name as "nama",
		tu.RGDESC "kanwil",
		tu.MBDESC  "kanca",
		tu.BRDESC "uker",
		t.task_type_name "jenis_task",
		t.kegiatan "kegiatan",
		COALESCE(th.activity, "-") "aktifitas", 
		t.product_name "produk",
		t.risk_issue "risk_issue",
		t.risk_indicator "indikator",
		CASE 
			WHEN t.period = 'Custom' THEN DATE_FORMAT(t.start_date,"%d-%b-%Y")
			WHEN t.period = 'Monthly' THEN DATE_FORMAT(t.start_date,"%d")
			ELSE "-"
		END as "tanggal_mulai",
		CASE 
			WHEN t.period = 'Custom' THEN DATE_FORMAT(t.end_date,"%d-%b-%Y")
			WHEN t.period = 'Monthly' THEN DATE_FORMAT(t.end_date,"%d")
			ELSE "-"
		END as "tanggal_selesai",
		COALESCE(kegiatan.jumlah, 0) as "jumlah_kegiatan_dilakukan"`).
		Joins(`join tasklists_uker tu on t.id = tu.tasklist_id `).
		Joins(`join (?) th on t.id = th.tasklist_id and tu.branch = th.branch`, subQuery).
		Joins(`LEFT JOIN (
					SELECT c.REGION, c.MAINBR, c.BRANCH, c.activity_id, c.product_id, ca.risk_issue_id, count(c.id) "jumlah" FROM coaching c 
					JOIN coaching_activity ca ON c.id = ca.coaching_id 
					WHERE c.deleted = 0
					AND (c.status = "01a" AND c.action = "Selesai")
					AND (c.status = "02b" AND (c.action = "Update" OR c.action ="Selesai"))
					GROUP BY c.REGION, c.MAINBR, c.BRANCH, c.activity_id, c.product_id, ca.risk_issue_id
				) kegiatan ON t.activity_id = kegiatan.activity_id
					and t.product_id = kegiatan.product_id
					and t.risk_issue_id = kegiatan.risk_issue_id
					and tu.REGION = kegiatan.REGION
					and tu.MAINBR = kegiatan.MAINBR
					and tu.BRANCH = kegiatan.BRANCH `).
		Where(`t.kegiatan = "Coaching"`).Group("t.id").Order("t.id ASC")

	if request.Nama != "" {
		// query = query.Where("t.maker_id = ? ", request.Nama)
		query = query.Where("th.pernr = ?", request.Nama)
	}

	if request.Kanwil != "all" && request.Kanwil != "" {
		query = query.Where("tu.REGION = ? ", request.Kanwil)
	}

	if request.Kanca != "" && request.Kanca != "all" {
		query = query.Where("tu.MAINBR = ? ", request.Kanca)
	}

	if request.Uker != "" && request.Uker != "all" {
		branches := strings.Split(request.Uker, ",")
		query = query.Where("tu.BRANCH in (?) ", branches)
	}

	if strings.ToLower(request.Aktifitas) != "0" && strings.ToLower(request.Aktifitas) != "" {
		query = query.Where("t.activity_id = ? ", request.Aktifitas)
	}

	if strings.ToLower(request.Produk) != "0" && strings.ToLower(request.Produk) != "" {
		query = query.Where("t.product_id = ? ", request.Produk)
	}

	if strings.ToLower(request.RiskIssue) != "0" && strings.ToLower(request.RiskIssue) != "" {
		query = query.Where("t.risk_issue_id = ? ", request.RiskIssue)
	}

	if strings.ToLower(request.Indikator) != "" {
		query = query.Where("t.risk_indicator_id = ? ", request.Indikator)
	}

	if strings.ToLower(request.Indikator) == "0" {
		query = query.Where("LOWER(t.risk_indicator) like ? ", "%"+request.IndikatorOther+"%")
	}

	if strings.ToLower(request.Status) != "" {
		if request.Status == "aktif" {
			query = query.Where("t.status = ? ", request.Status)
		} else {
			query = query.Where("t.status != ? ", "aktif")
		}

	}

	if strings.ToLower(request.JenisTaks) != "" {
		query = query.Where("t.task_type = ? ", request.JenisTaks)
	}

	err = query.Scan(&responses).Error

	if err != nil {
		l.logger.Zap.Error(err)
		return responses, err
	}

	// l.logger.Zap.Info(responses)
	return responses, err
}

// 28-08-2023
func (l LaporanRepository) GetPerhitunganVerifikasi(request models.PerhitunganPersentasePenyelesaianPagianted) (responses []models.LaporanVerifikasiResponse, totalRow int, totalData int, err error) {
	db := l.db.DB

	//09/09/2024
	// subQuery := db.Table(`tasklists_history`).Select(`*`).Group(`tasklist_id, branch`)

	tbl_task := `lampiran_rap_` + request.RiskIssue + `_` + request.Indikator

	queryVerifikasi := db.Table(tbl_task + " as task").
		Select(`
			task.id 'id',
			task.tasklist_id 'tasklist_id',
			task.no_verifikasi 'no_verifikasi',
			rlv.BRANCH 'uker',
			rlv.brd_urc 'pernr',
			rlv.butuh_perbaikan,
			rlv.status_perbaikan_konsolidasi,
			th.sample 'jumlah_data_anomali',
			rlv.jumlah_data_yg_diverifikasi 'jumlah_data_yg_diverifikasi',
			rlv.status_perbaikan_proses 'jumlah_data_yg_harus_diperbaiki',
			rlv.status_perbaikan_selesai 'jumlah_data_sudah_tindaklanjut',
			CASE
				WHEN rlv.butuh_perbaikan = "Ya"
					THEN rlv.jumlah_data_yg_diverifikasi
				ELSE "0"
			END AS 'jumlah_kegiatan_yg_dilakukan'
		`).
		Joins(`LEFT JOIN report_list_verifikasi rlv ON task.no_verifikasi = rlv.no_pelaporan`).
		Joins(`JOIN tasklists_history th ON task.tasklist_id = th.tasklist_id`).
		Where(`task.status = "Verify"`).
		Group(`task.no_verifikasi`).Order(`task.id`)

	querySummary := db.Table("(?) as mytask", queryVerifikasi).
		Select(`
			mytask.tasklist_id,
			mytask.pernr,
			mytask.jumlah_data_anomali,
			SUM(mytask.jumlah_data_yg_diverifikasi) 'jumlah_data_yg_diverifikasi',
			SUM(mytask.jumlah_data_sudah_tindaklanjut) 'jumlah_data_sudah_tindaklanjut',
			SUM(mytask.jumlah_data_yg_harus_diperbaiki) 'jumlah_data_yg_harus_diperbaiki',
			SUM(mytask.jumlah_kegiatan_yg_dilakukan) 'jumlah_kegiatan_yg_dilakukan'
		`).Group(`mytask.tasklist_id`).Group(`mytask.pernr`)

	query := db.Table("tasklists t").
		Select(`
			t.id,
			t.no_tasklist "no_tasklist",
			t.nama_tasklist "nama_tasklist",
			muk.pn,
			muk.sname 'nama',
			tu.RGDESC "kanwil",
			tu.MBDESC "kanca",
			tu.BRDESC "uker",
			task.tasklist_id 'tasklist_id',
			t.task_type_name "jenis_task",
			t.kegiatan "kegiatan",
			a.name  "aktifitas",
			t.product_name "produk",
			t.risk_issue "risk_issue",
			t.risk_indicator "indikator",
			-- task.butuh_perbaikan 'butuh_perbaikan',
			CASE
							WHEN t.period = 'Custom' THEN DATE_FORMAT(t.start_date, "%d-%b-%Y")
							WHEN t.period = 'Monthly' THEN DATE_FORMAT(t.start_date, "%d")
							ELSE "-"
			END as "tanggal_mulai",
			CASE
							WHEN t.period = 'Custom' THEN DATE_FORMAT(t.end_date,"%d-%b-%Y")
							WHEN t.period = 'Monthly' THEN DATE_FORMAT(t.end_date,"%d")
							ELSE "-"
			END as "tanggal_selesai",
			task.jumlah_data_anomali 'jumlah_data_anomali',
			SUM(task.jumlah_data_yg_diverifikasi) AS jumlah_data_verifikasi,
			0 as "persen_verifikasi",
			SUM(task.jumlah_data_sudah_tindaklanjut) 'jumlah_data_sudah_tindaklanjut',
			task.jumlah_data_yg_harus_diperbaiki 'jumlah_data_perlu_tindaklanjut',
			0 as "persen_sudah_tindaklanjut",
			SUM(task.jumlah_data_yg_diverifikasi) 'jumlah_kegiatan_dilakukan'	
		`).
		Joins(`LEFT JOIN (?) as task ON task.tasklist_id = t.id`, querySummary).
		Joins(`INNER JOIN tasklists_uker tu ON tu.tasklist_id = t.id`).
		Joins(`INNER JOIN mst_uker_kelolaan muk ON muk.pn = task.pernr`).
		Joins(`INNER JOIN activity a ON a.kode_activity = t.activity_id`).
		Where(`t.kegiatan = "Verifikasi"`).
		Group(`t.id`).
		// Group(`task.butuh_perbaikan`).
		Order(`t.id`)
		// Group(`t.activity_id`).Group(`t.product_id`).Group(`t.risk_issue_id`).Group(`tu.REGION`).Group(`tu.MAINBR`).Group(`tu.BRANCH`)

	if request.Nama != "" {
		// query = query.Where("t.maker_id = ? ", request.Nama)
		// query = query.Where("rlv.brd_urc = ?", request.Nama)
		query = query.Where("task.pernr = ?", request.Nama)
	}

	if request.JenisReport != "1" {
		if request.Nama == "" {
			query = query.Group(`task.pernr`)
		}
	}

	if request.Kanwil != "all" && request.Kanwil != "" {
		query = query.Where("tu.REGION = ? ", request.Kanwil)
	}

	if request.Kanca != "" && request.Kanca != "all" {
		query = query.Where("tu.MAINBR = ? ", request.Kanca)
	}

	if request.Uker != "" && request.Uker != "all" {
		branches := strings.Split(request.Uker, ",")
		query = query.Where("tu.BRANCH in (?) ", branches)
	}

	if strings.ToLower(request.Aktifitas) != "0" && strings.ToLower(request.Aktifitas) != "" {
		query = query.Where("t.activity_id = ? ", request.Aktifitas)
	}

	if strings.ToLower(request.Produk) != "0" && strings.ToLower(request.Produk) != "" {
		query = query.Where("t.product_id = ? ", request.Produk)
	}

	if strings.ToLower(request.RiskIssue) != "0" && strings.ToLower(request.RiskIssue) != "" {
		query = query.Where("t.risk_issue_id = ? ", request.RiskIssue)
	}

	if strings.ToLower(request.Indikator) != "" {
		query = query.Where("t.risk_indicator_id = ? ", request.Indikator)
	}

	if strings.ToLower(request.Indikator) == "0" {
		query = query.Where("LOWER(t.risk_indicator) like ? ", "%"+request.IndikatorOther+"%")
	}

	if strings.ToLower(request.Status) != "" {
		if request.Status == "aktif" {
			query = query.Where("t.status = ? ", request.Status)
		} else {
			query = query.Where("t.status != ? ", "aktif")
		}

	}

	if strings.ToLower(request.JenisTaks) != "" {
		query = query.Where("t.task_type = ? ", request.JenisTaks)
	}

	var count int64
	query.Count(&count)

	totalRow = int(count)

	if request.Limit != 0 {
		query = query.Limit(request.Limit)
	}

	if request.Offset != 0 {
		query = query.Offset(request.Offset)
	}

	err = query.Scan(&responses).Error

	result := float64(totalRow) / float64(request.Limit)
	resultFinal := int(math.Ceil(result))

	if err != nil {
		l.logger.Zap.Error(err)
		return responses, 0, 0, err
	}

	//
	return responses, resultFinal, totalRow, err
}

func (l LaporanRepository) GetPerhitunganVerifikasiDownload(request models.PerhitunganPersentasePenyelesaianDownload) (responses []models.LaporanVerifikasiResponse, err error) {
	db := l.db.DB

	//09/09/2024
	// subQuery := db.Table(`tasklists_history`).Select(`*`).Group(`tasklist_id, branch`)

	tbl_task := `lampiran_rap_` + request.RiskIssue + `_` + request.Indikator

	queryVerifikasi := db.Table(tbl_task + " as task").
		Select(`
			task.id 'id',
			task.tasklist_id 'tasklist_id',
			task.no_verifikasi 'no_verifikasi',
			rlv.BRANCH 'uker',
			rlv.brd_urc 'pernr',
			rlv.butuh_perbaikan,
			rlv.status_perbaikan_konsolidasi,
			th.sample 'jumlah_data_anomali',
			rlv.jumlah_data_yg_diverifikasi 'jumlah_data_yg_diverifikasi',
			rlv.status_perbaikan_proses 'jumlah_data_yg_harus_diperbaiki',
			rlv.status_perbaikan_selesai 'jumlah_data_sudah_tindaklanjut',
			CASE
				WHEN rlv.butuh_perbaikan = "Ya"
					THEN rlv.jumlah_data_yg_diverifikasi
				ELSE "0"
			END AS 'jumlah_kegiatan_yg_dilakukan'
		`).
		Joins(`LEFT JOIN report_list_verifikasi rlv ON task.no_verifikasi = rlv.no_pelaporan`).
		Joins(`JOIN tasklists_history th ON task.tasklist_id = th.tasklist_id`).
		Where(`task.status = "Verify"`).
		Group(`task.no_verifikasi`).Order(`task.id`)

	querySummary := db.Table("(?) as mytask", queryVerifikasi).
		Select(`
			mytask.tasklist_id,
			mytask.pernr,
			mytask.jumlah_data_anomali,
			SUM(mytask.jumlah_data_yg_diverifikasi) 'jumlah_data_yg_diverifikasi',
			SUM(mytask.jumlah_data_sudah_tindaklanjut) 'jumlah_data_sudah_tindaklanjut',
			SUM(mytask.jumlah_data_yg_harus_diperbaiki) 'jumlah_data_yg_harus_diperbaiki',
			SUM(mytask.jumlah_kegiatan_yg_dilakukan) 'jumlah_kegiatan_yg_dilakukan'
		`).Group(`mytask.tasklist_id`).Group(`mytask.pernr`)

	query := db.Table("tasklists t").
		Select(`
			t.id,
			t.no_tasklist "no_tasklist",
			t.nama_tasklist "nama_tasklist",
			muk.pn,
			muk.sname 'nama',
			tu.RGDESC "kanwil",
			tu.MBDESC "kanca",
			tu.BRDESC "uker",
			task.tasklist_id 'tasklist_id',
			t.task_type_name "jenis_task",
			t.kegiatan "kegiatan",
			a.name  "aktifitas",
			t.product_name "produk",
			t.risk_issue "risk_issue",
			t.risk_indicator "indikator",
			-- task.butuh_perbaikan 'butuh_perbaikan',
			CASE
							WHEN t.period = 'Custom' THEN DATE_FORMAT(t.start_date, "%d-%b-%Y")
							WHEN t.period = 'Monthly' THEN DATE_FORMAT(t.start_date, "%d")
							ELSE "-"
			END as "tanggal_mulai",
			CASE
							WHEN t.period = 'Custom' THEN DATE_FORMAT(t.end_date,"%d-%b-%Y")
							WHEN t.period = 'Monthly' THEN DATE_FORMAT(t.end_date,"%d")
							ELSE "-"
			END as "tanggal_selesai",
			task.jumlah_data_anomali 'jumlah_data_anomali',
			SUM(task.jumlah_data_yg_diverifikasi) AS jumlah_data_verifikasi,
			0 as "persen_verifikasi",
			SUM(task.jumlah_data_sudah_tindaklanjut) 'jumlah_data_sudah_tindaklanjut',
			task.jumlah_data_yg_harus_diperbaiki 'jumlah_data_perlu_tindaklanjut',
			0 as "persen_sudah_tindaklanjut",
			SUM(task.jumlah_data_yg_diverifikasi) 'jumlah_kegiatan_dilakukan'	
		`).
		Joins(`LEFT JOIN (?) as task ON task.tasklist_id = t.id`, querySummary).
		Joins(`INNER JOIN tasklists_uker tu ON tu.tasklist_id = t.id`).
		Joins(`INNER JOIN mst_uker_kelolaan muk ON muk.pn = task.pernr`).
		Joins(`INNER JOIN activity a ON a.kode_activity = t.activity_id`).
		Where(`t.kegiatan = "Verifikasi"`).
		Group(`t.id`).
		// Group(`task.butuh_perbaikan`).
		Order(`t.id`)
		// Group(`t.activity_id`).Group(`t.product_id`).Group(`t.risk_issue_id`).Group(`tu.REGION`).Group(`tu.MAINBR`).Group(`tu.BRANCH`)

	if request.Nama != "" {
		// query = query.Where("t.maker_id = ? ", request.Nama)
		// query = query.Where("rlv.brd_urc = ?", request.Nama)
		query = query.Where("th.pernr = ?", request.Nama)
	}

	if request.JenisReport != "1" {
		if request.Nama == "" {
			query = query.Group(`task.pernr`)
		}
	}

	if request.Kanwil != "all" && request.Kanwil != "" {
		query = query.Where("tu.REGION = ? ", request.Kanwil)
	}

	if request.Kanca != "" && request.Kanca != "all" {
		query = query.Where("tu.MAINBR = ? ", request.Kanca)
	}

	if request.Uker != "" && request.Uker != "all" {
		if request.Uker != "" && request.Uker != "all" {
			branches := strings.Split(request.Uker, ",")
			query = query.Where("tu.BRANCH in (?) ", branches)
		}
	}

	if strings.ToLower(request.Aktifitas) != "0" && strings.ToLower(request.Aktifitas) != "" {
		query = query.Where("t.activity_id = ? ", request.Aktifitas)
	}

	if strings.ToLower(request.Produk) != "0" && strings.ToLower(request.Produk) != "" {
		query = query.Where("t.product_id = ? ", request.Produk)
	}

	if strings.ToLower(request.RiskIssue) != "0" && strings.ToLower(request.RiskIssue) != "" {
		query = query.Where("t.risk_issue_id = ? ", request.RiskIssue)
	}

	if strings.ToLower(request.Indikator) != "" {
		query = query.Where("t.risk_indicator_id = ? ", request.Indikator)
	}

	if strings.ToLower(request.Indikator) == "0" {
		query = query.Where("LOWER(t.risk_indicator) like ? ", "%"+request.IndikatorOther+"%")
	}

	if strings.ToLower(request.Status) != "" {
		if request.Status == "aktif" {
			query = query.Where("t.status = ? ", request.Status)
		} else {
			query = query.Where("t.status != ? ", "aktif")
		}

	}

	if strings.ToLower(request.JenisTaks) != "" {
		query = query.Where("t.task_type = ? ", request.JenisTaks)
	}

	err = query.Scan(&responses).Error

	if err != nil {
		l.logger.Zap.Error(err)
		return responses, err
	}

	// l.logger.Zap.Info(responses)
	return responses, err
}

// lama
func (r LaporanRepository) GetRiskEventOnTaskList() (response []models.RiskEventOnTaskList, err error) {
	baseQuery := `select DISTINCT t.risk_issue_id, ri.risk_issue  from risk_issue ri 
					join tasklists t on t.risk_issue_id = ri.id 
					GROUP BY t.risk_issue_id`

	err = r.db.DB.Raw(baseQuery).Scan(&response).Error

	if err != nil {
		r.logger.Zap.Error(err)
		return response, err
	}
	return response, err

}

func (r LaporanRepository) GetPerhitunganPersentasePenyelesaianPerPekerja(request models.PerhitunganPersentasePenyelesaianPagianted) (responses []models.LaporanPerPekerjaQueryResult, totalRow int, totalData int, err error) {
	db := r.db.DB

	query := db.Table("tasklists t").
		Select(`t.id as id,
		t.maker_id as pn,
		pe.SNAME  as nama,
		tu.RGDESC as "kanwil",
		tu.MBDESC as "kanca",
		tu.unit_kerja as "uker", 
		a.name as "aktifitas", 
		p.product as "produk", 
		t.task_type_name as "jenis_task", 
		DATE_FORMAT(t.start_date,"%d-%b-%Y") as "tanggal_mulai", 
		DATE_FORMAT(t.end_date,"%d-%b-%Y") as "tanggal_selesai", 
		ri.risk_indicator as "indikator",
		CASE 
			WHEN v.sumber_data = "Non KRID" THEN
				COALESCE ( data_verif.jumlah_data_verifikasi , 0 )
			WHEN v.sumber_data = "KRID" THEN
				COALESCE ( data_verif_krid.jumlah_data_verifikasi , 0 )
			ELSE 0
		END as jumlah_data_anomali,
		CASE 
			WHEN v.sumber_data = "Non KRID" THEN 
				COALESCE ( perbaikan_selesai.jumlah_data_verifikasi  , 0 )
			WHEN v.sumber_data = "KRID" THEN 
				COALESCE ( perbaikan_selesai_krid.jumlah_data_verifikasi  , 0 )
			ELSE 0
		END as 'jumlah_data_verifikasi',
		CASE
		WHEN v.sumber_data = "Non KRID" THEN 
				COALESCE ( perbaikan_selesai.jumlah_data_verifikasi  , 0 )
		WHEN v.sumber_data = "KRID" THEN 
				COALESCE ( perbaikan_selesai_krid.jumlah_data_verifikasi  , 0 )
		ELSE 0 
		END as 'jumlah_data_perlu_tindaklanjut',
		0 as 'jumlah_data_sudah_tindaklanjut'`).
		Joins(`JOIN pa0001_eof pe on pe.PERNR = t.maker_id`).
		Joins(`join tasklists_uker tu on tu.tasklist_id`).
		Joins(`JOIN verifikasi v on v.activity_id = t.activity_id`).
		Joins(`JOIN activity a ON a.id = t.activity_id`).
		Joins(`JOIN product p on p.id = t.product_id`).
		Joins(`join risk_indicator ri on ri.id = t.risk_indicator_id`).
		Joins(`LEFT JOIN (
						SELECT
								vda.verifikasi_id 'id',
								COUNT(vda.verifikasi_id) 'jumlah_data_verifikasi'
						FROM verifikasi_data_anomali vda
						GROUP BY vda.verifikasi_id
				) data_verif ON data_verif.id = v.id`).
		Joins(`LEFT JOIN (
						SELECT
								vdak.verifikasi_id 'id',
								COUNT(vdak.verifikasi_id) 'jumlah_data_verifikasi'
						FROM verifikasi_data_anomali_krid vdak
						GROUP BY vdak.verifikasi_id
				) data_verif_krid ON data_verif_krid.id = v.id`).
		Joins(`LEFT JOIN (
						SELECT
								vda.verifikasi_id 'id',
								COUNT(vda.verifikasi_id) 'jumlah_data_verifikasi'
						FROM verifikasi_data_anomali vda
						JOIN verifikasi_pic_tindak_lanjut vptl ON vptl.verifikasi_id = vda.verifikasi_id 
						WHERE vptl.status in (2)
						GROUP BY vda.verifikasi_id
				) perbaikan_selesai ON perbaikan_selesai.id = v.id`).
		Joins(`LEFT JOIN (
						SELECT
								vdak.verifikasi_id 'id',
								COUNT(vdak.verifikasi_id) 'jumlah_data_verifikasi'
						FROM verifikasi_data_anomali_krid vdak
						JOIN verifikasi_pic_tindak_lanjut vptl ON vptl.verifikasi_id = vdak.verifikasi_id 
						WHERE vptl.status in (2)
						GROUP BY vdak.verifikasi_id
				) perbaikan_selesai_krid ON perbaikan_selesai_krid.id = v.id`).
		Joins(`LEFT JOIN (
						SELECT
								vda.verifikasi_id 'id',
								COUNT(vda.verifikasi_id) 'jumlah_data_verifikasi'
						FROM verifikasi_data_anomali vda
						JOIN verifikasi_pic_tindak_lanjut vptl ON vptl.verifikasi_id = vda.verifikasi_id 
						WHERE vptl.status in (0,1)
						GROUP BY vda.verifikasi_id
				) perbaikan_onprogres ON perbaikan_onprogres.id = v.id`).
		Joins(`LEFT JOIN (
						SELECT
								vdak.verifikasi_id 'id',
								COUNT(vdak.verifikasi_id) 'jumlah_data_verifikasi'
						FROM verifikasi_data_anomali_krid vdak
						JOIN verifikasi_pic_tindak_lanjut vptl ON vptl.verifikasi_id = vdak.verifikasi_id 
						WHERE vptl.status in (0,1)
						GROUP BY vdak.verifikasi_id
				) perbaikan_onprogres_krid ON perbaikan_onprogres_krid.id = v.id`)

	if request.Nama != "" {
		query = query.Where("t.maker_id = ? ", request.Nama)
	}

	if strings.ToLower(request.Aktifitas) != "0" && strings.ToLower(request.Aktifitas) != "" {
		query = query.Where("t.activity_id = ? ", request.Aktifitas)
	}

	if strings.ToLower(request.Produk) != "0" && strings.ToLower(request.Produk) != "" {
		query = query.Where("t.product_id = ? ", request.Produk)
	}

	if strings.ToLower(request.Indikator) != "0" && strings.ToLower(request.Indikator) != "" {
		query = query.Where("t.risk_indicator_id = ? ", request.Indikator)
	}

	if strings.ToLower(request.Status) != "" {
		if request.Status == "aktif" {
			query = query.Where("t.status = ? ", request.Status)
		} else {
			query = query.Where("t.status != ? ", "aktif")
		}

	}

	if strings.ToLower(request.JenisTaks) != "" {
		query = query.Where("t.task_type = ? ", request.JenisTaks)
	}

	var count int64
	query.Count(&count)

	totalRow = int(count)

	if request.Limit != 0 {
		query = query.Limit(request.Limit)
	}

	if request.Offset != 0 {
		query = query.Offset(request.Offset)
	}

	query.Scan(&responses)
	err = query.Error

	result := float64(totalRow) / float64(request.Limit)
	resultFinal := int(math.Ceil(result))

	if err != nil {
		r.logger.Zap.Error(err)
		return responses, 0, 0, err
	}
	return responses, totalRow, resultFinal, err
}

func (r LaporanRepository) GetPerhitunganPersentasePenyelesaianPerPekerjaDownload(request models.PerhitunganPersentasePenyelesaianDownload) (responses []models.LaporanPerPekerjaQueryResult, err error) {
	db := r.db.DB

	query := db.Table("tasklists t").
		Select(`t.id as id,
		t.maker_id as pn,
		pe.SNAME  as nama,
		tu.RGDESC as "kanwil",
		tu.MBDESC as "kanca",
		tu.unit_kerja as "uker", 
		a.name as "aktifitas", 
		p.product as "produk", 
		t.task_type_name as "jenis_task", 
		DATE_FORMAT(t.start_date,"%d-%b-%Y") as "tanggal_mulai", 
		DATE_FORMAT(t.end_date,"%d-%b-%Y") as "tanggal_selesai", 
		ri.risk_indicator as "indikator",
		CASE 
			WHEN v.sumber_data = "Non KRID" THEN
				COALESCE ( data_verif.jumlah_data_verifikasi , 0 )
			WHEN v.sumber_data = "KRID" THEN
				COALESCE ( data_verif_krid.jumlah_data_verifikasi , 0 )
			ELSE 0
		END as jumlah_data_anomali,
		CASE 
			WHEN v.sumber_data = "Non KRID" THEN 
				COALESCE ( perbaikan_selesai.jumlah_data_verifikasi  , 0 )
			WHEN v.sumber_data = "KRID" THEN 
				COALESCE ( perbaikan_selesai_krid.jumlah_data_verifikasi  , 0 )
			ELSE 0
		END as 'jumlah_data_verifikasi',
		CASE
		WHEN v.sumber_data = "Non KRID" THEN 
				COALESCE ( perbaikan_selesai.jumlah_data_verifikasi  , 0 )
		WHEN v.sumber_data = "KRID" THEN 
				COALESCE ( perbaikan_selesai_krid.jumlah_data_verifikasi  , 0 )
		ELSE 0 
		END as 'jumlah_data_perlu_tindaklanjut',
		0 as 'jumlah_data_sudah_tindaklanjut'`).
		Joins(`JOIN pa0001_eof pe on pe.PERNR = t.maker_id`).
		Joins(`join tasklists_uker tu on tu.tasklist_id`).
		Joins(`JOIN verifikasi v on v.activity_id = t.activity_id`).
		Joins(`JOIN activity a ON a.id = t.activity_id`).
		Joins(`JOIN product p on p.id = t.product_id`).
		Joins(`join risk_indicator ri on ri.id = t.risk_indicator_id`).
		Joins(`LEFT JOIN (
						SELECT
								vda.verifikasi_id 'id',
								COUNT(vda.verifikasi_id) 'jumlah_data_verifikasi'
						FROM verifikasi_data_anomali vda
						GROUP BY vda.verifikasi_id
				) data_verif ON data_verif.id = v.id`).
		Joins(`LEFT JOIN (
						SELECT
								vdak.verifikasi_id 'id',
								COUNT(vdak.verifikasi_id) 'jumlah_data_verifikasi'
						FROM verifikasi_data_anomali_krid vdak
						GROUP BY vdak.verifikasi_id
				) data_verif_krid ON data_verif_krid.id = v.id`).
		Joins(`LEFT JOIN (
						SELECT
								vda.verifikasi_id 'id',
								COUNT(vda.verifikasi_id) 'jumlah_data_verifikasi'
						FROM verifikasi_data_anomali vda
						JOIN verifikasi_pic_tindak_lanjut vptl ON vptl.verifikasi_id = vda.verifikasi_id 
						WHERE vptl.status in (2)
						GROUP BY vda.verifikasi_id
				) perbaikan_selesai ON perbaikan_selesai.id = v.id`).
		Joins(`LEFT JOIN (
						SELECT
								vdak.verifikasi_id 'id',
								COUNT(vdak.verifikasi_id) 'jumlah_data_verifikasi'
						FROM verifikasi_data_anomali_krid vdak
						JOIN verifikasi_pic_tindak_lanjut vptl ON vptl.verifikasi_id = vdak.verifikasi_id 
						WHERE vptl.status in (2)
						GROUP BY vdak.verifikasi_id
				) perbaikan_selesai_krid ON perbaikan_selesai_krid.id = v.id`).
		Joins(`LEFT JOIN (
						SELECT
								vda.verifikasi_id 'id',
								COUNT(vda.verifikasi_id) 'jumlah_data_verifikasi'
						FROM verifikasi_data_anomali vda
						JOIN verifikasi_pic_tindak_lanjut vptl ON vptl.verifikasi_id = vda.verifikasi_id 
						WHERE vptl.status in (0,1)
						GROUP BY vda.verifikasi_id
				) perbaikan_onprogres ON perbaikan_onprogres.id = v.id`).
		Joins(`LEFT JOIN (
						SELECT
								vdak.verifikasi_id 'id',
								COUNT(vdak.verifikasi_id) 'jumlah_data_verifikasi'
						FROM verifikasi_data_anomali_krid vdak
						JOIN verifikasi_pic_tindak_lanjut vptl ON vptl.verifikasi_id = vdak.verifikasi_id 
						WHERE vptl.status in (0,1)
						GROUP BY vdak.verifikasi_id
				) perbaikan_onprogres_krid ON perbaikan_onprogres_krid.id = v.id`)

	if request.Nama != "" {
		query = query.Where("t.maker_id = ? ", request.Nama)
	}

	if strings.ToLower(request.Aktifitas) != "0" && strings.ToLower(request.Aktifitas) != "" {
		query = query.Where("t.activity_id = ? ", request.Aktifitas)
	}

	if strings.ToLower(request.Produk) != "0" && strings.ToLower(request.Produk) != "" {
		query = query.Where("t.product_id = ? ", request.Produk)
	}

	if strings.ToLower(request.Indikator) != "0" && strings.ToLower(request.Indikator) != "" {
		query = query.Where("t.risk_indicator_id = ? ", request.Indikator)
	}

	if strings.ToLower(request.Status) != "" {
		if request.Status == "aktif" {
			query = query.Where("t.status = ? ", request.Status)
		} else {
			query = query.Where("t.status != ? ", "aktif")
		}

	}

	if strings.ToLower(request.JenisTaks) != "" {
		query = query.Where("t.task_type = ? ", request.JenisTaks)
	}
	query.Scan(&responses)
	err = query.Error

	if err != nil {
		r.logger.Zap.Error(err)
		return responses, err
	}
	return responses, err
}

func (r LaporanRepository) GetPerhitunganPersentasePenyelesaianPerUker(request models.PerhitunganPersentasePenyelesaianPagianted) (responses []models.LaporanPerUkerResult, totalRow int, err error) {
	sQueryKanwil := " lower(tu.RGDESC) like ? "
	sQueryKanca := "AND lower(tu.MBDESC) like ? "
	sQueryUker := "AND lower(tu.BRDESC) like ? "

	sQueryActivity := " AND t.activity_id = ? "
	sQueryProduct := " AND t.product_id = ? "
	sQueryIndikator := " AND t.risk_indicator_id = ? "

	sQueryStatus := "AND t.status = ? "
	sQueryJenisTask := "AND t.task_type = ? "

	if request.Kanwil != "all" {
		sQueryKanwil = " lower(tu.RGDESC) like ? "
	}

	if strings.ToLower(request.Kanca) != "" && request.Kanca != "all" {
		sQueryKanca = "AND lower(tu.MBDESC) like ? "
	}

	if strings.ToLower(request.Uker) != "" {
		sQueryUker = "AND lower(tu.BRDESC) like ? "
	}

	if strings.ToLower(request.Aktifitas) == "0" || strings.ToLower(request.Aktifitas) == "" {
		sQueryActivity = "AND t.activity_id != ?"
	}

	if strings.ToLower(request.Produk) == "0" || strings.ToLower(request.Produk) == "" {
		sQueryProduct = " AND t.product_id != ? "
	}

	if strings.ToLower(request.Indikator) == "0" || strings.ToLower(request.Indikator) == "" {
		sQueryIndikator = " AND t.risk_indicator_id != ? "
	}

	if strings.ToLower(request.Status) == "" {
		sQueryStatus = "AND t.status != ? "
	}

	if strings.ToLower(request.JenisTaks) == "" {
		sQueryJenisTask = "AND t.task_type != ? "
	}

	baseQuery := `SELECT 
		t.id as id,
		t.maker_id as pn,
		pe.SNAME  as nama,
		tu.RGDESC as "kanwil",
		tu.MBDESC as "kanca",
		tu.BRDESC as "uker", 
		a.name as "aktifitas", 
		p.product as "produk", 
		t.task_type as "jenis_task", 
		DATE_FORMAT(t.start_date,"%d %b %Y") as "tanggal_mulai", 
		DATE_FORMAT(t.end_date,"%d %b %Y") as  "tanggal_selesai",
		ri.risk_indicator "indikator",
		CASE 
			WHEN v.sumber_data = "Non KRID" THEN
				COALESCE ( data_verif.jumlah_data_verifikasi , 0 )
			WHEN v.sumber_data = "KRID" THEN
				COALESCE ( data_verif_krid.jumlah_data_verifikasi , 0 )
			ELSE 0
		END as jumlah_data_anomaly,
		CASE 
			WHEN v.sumber_data = "Non KRID" THEN 
				COALESCE ( perbaikan_selesai.jumlah_data_verifikasi  , 0 )
			WHEN v.sumber_data = "KRID" THEN 
				COALESCE ( perbaikan_selesai_krid.jumlah_data_verifikasi  , 0 )
			ELSE 0
		END as 'jumlah_data_sudah_verifikasi',
		CASE
		WHEN v.sumber_data = "Non KRID" THEN 
				COALESCE ( perbaikan_selesai.jumlah_data_verifikasi  , 0 )
		WHEN v.sumber_data = "KRID" THEN 
				COALESCE ( perbaikan_selesai_krid.jumlah_data_verifikasi  , 0 )
		ELSE 0 
		END as 'jumlah_data_perlu_tindaklanjut',
		0 as 'jumlah_data_sudah_tindaklanjut'
		
		from tasklists t 
		JOIN pa0001_eof pe on pe.PERNR = t.maker_id
		join tasklists_uker tu on tu.tasklist_id
		JOIN verifikasi v on v.activity_id = t.activity_id
		JOIN activity a ON a.id = t.activity_id
		JOIN product p on p.id = t.product_id
		join risk_indicator ri on ri.id = t.risk_indicator_id 
		LEFT JOIN (
					SELECT
						vda.verifikasi_id 'id',
						COUNT(vda.verifikasi_id) 'jumlah_data_verifikasi'
					FROM verifikasi_data_anomali vda
					GROUP BY vda.verifikasi_id
				) data_verif ON data_verif.id = v.id
		LEFT JOIN (
					SELECT
						vdak.verifikasi_id 'id',
						COUNT(vdak.verifikasi_id) 'jumlah_data_verifikasi'
					FROM verifikasi_data_anomali_krid vdak
					GROUP BY vdak.verifikasi_id
				) data_verif_krid ON data_verif_krid.id = v.id
		LEFT JOIN (
					SELECT
						vda.verifikasi_id 'id',
						COUNT(vda.verifikasi_id) 'jumlah_data_verifikasi'
					FROM verifikasi_data_anomali vda
					JOIN verifikasi_pic_tindak_lanjut vptl ON vptl.verifikasi_id = vda.verifikasi_id 
					WHERE vptl.status in (2)
					GROUP BY vda.verifikasi_id
				) perbaikan_selesai ON perbaikan_selesai.id = v.id
		LEFT JOIN (
					SELECT
						vdak.verifikasi_id 'id',
						COUNT(vdak.verifikasi_id) 'jumlah_data_verifikasi'
					FROM verifikasi_data_anomali_krid vdak
					JOIN verifikasi_pic_tindak_lanjut vptl ON vptl.verifikasi_id = vdak.verifikasi_id 
					WHERE vptl.status in (2)
					GROUP BY vdak.verifikasi_id
				) perbaikan_selesai_krid ON perbaikan_selesai_krid.id = v.id
		LEFT JOIN (
					SELECT
						vda.verifikasi_id 'id',
						COUNT(vda.verifikasi_id) 'jumlah_data_verifikasi'
					FROM verifikasi_data_anomali vda
					JOIN verifikasi_pic_tindak_lanjut vptl ON vptl.verifikasi_id = vda.verifikasi_id 
					WHERE vptl.status in (0,1)
					GROUP BY vda.verifikasi_id
				) perbaikan_onprogres ON perbaikan_onprogres.id = v.id
		LEFT JOIN (
					SELECT
						vdak.verifikasi_id 'id',
						COUNT(vdak.verifikasi_id) 'jumlah_data_verifikasi'
					FROM verifikasi_data_anomali_krid vdak
					JOIN verifikasi_pic_tindak_lanjut vptl ON vptl.verifikasi_id = vdak.verifikasi_id 
					WHERE vptl.status in (0,1)
					GROUP BY vdak.verifikasi_id
				) perbaikan_onprogres_krid ON perbaikan_onprogres_krid.id = v.id

                WHERE` + sQueryKanwil + sQueryKanca + sQueryUker + sQueryActivity + sQueryProduct + sQueryIndikator + sQueryStatus + sQueryJenisTask

	query := baseQuery + ` order by ` + request.Order + ` ` + request.Sort + ` LIMIT ? OFFSET ?`
	queryCount := `select count(*) from (` + baseQuery + `) as totalData`

	err = r.db.DB.Raw(query,
		strings.ToLower(request.Kanwil)+"%",
		strings.ToLower(request.Kanca)+"%",
		strings.ToLower(request.Uker)+"%",
		request.Aktifitas,
		request.Produk,
		request.Indikator,
		request.Status,
		request.JenisTaks,
		request.Limit,
		request.Offset).Scan(&responses).Error

	err = r.db.DB.Raw(queryCount,
		strings.ToLower(request.Kanwil)+"%",
		strings.ToLower(request.Kanca)+"%",
		strings.ToLower(request.Uker)+"%",
		request.Aktifitas,
		request.Produk,
		request.Indikator,
		request.Status,
		request.JenisTaks).Scan(&totalRow).Error

	if err != nil {
		r.logger.Zap.Error(err)
		return responses, 0, err
	}
	return responses, totalRow, err
}

func (r LaporanRepository) GetPerhitunganPersentasePenyelesaianPerUkerDownload(request models.PerhitunganPersentasePenyelesaianDownload) (responses []models.LaporanPerUkerResult, err error) {
	sQueryKanwil := " lower(tu.RGDESC) like ? "
	sQueryKanca := "AND lower(tu.MBDESC) like ? "
	sQueryUker := "AND lower(tu.BRDESC) like ? "

	sQueryActivity := " AND t.activity_id = ? "
	sQueryProduct := " AND t.product_id = ? "
	sQueryIndikator := " AND t.risk_indicator_id = ? "

	sQueryStatus := "AND t.status = ? "
	sQueryJenisTask := "AND t.task_type = ? "

	if request.Kanwil != "all" {
		sQueryKanwil = " lower(tu.RGDESC) like ? "
	}

	if strings.ToLower(request.Kanca) != "" {
		sQueryKanca = "AND lower(tu.MBDESC) like ? "
	}

	if strings.ToLower(request.Uker) != "" {
		sQueryUker = "AND lower(tu.BRDESC) like ? "
	}

	if strings.ToLower(request.Aktifitas) == "0" || strings.ToLower(request.Aktifitas) == "" {
		sQueryActivity = "AND t.activity_id != ?"
	}

	if strings.ToLower(request.Produk) == "0" || strings.ToLower(request.Produk) == "" {
		sQueryProduct = " AND t.product_id != ? "
	}

	if strings.ToLower(request.Indikator) == "0" || strings.ToLower(request.Indikator) == "" {
		sQueryIndikator = " AND t.risk_indicator_id != ? "
	}

	if strings.ToLower(request.Status) == "" {
		sQueryStatus = "AND t.status != ? "
	}

	if strings.ToLower(request.JenisTaks) == "" {
		sQueryJenisTask = "AND t.task_type != ? "
	}

	baseQuery := `SELECT 
		t.id as id,
		t.maker_id as pn,
		pe.SNAME  as nama,
		tu.RGDESC as "kanwil",
		tu.MBDESC as "kanca",
		tu.unit_kerja as "uker", 
		a.name as "aktifitas", 
		p.product as "produk", 
		t.task_type_name as "jenis_task", 
		DATE_FORMAT(t.start_date,"%d %b %Y") as "tanggal_mulai", 
		DATE_FORMAT(t.end_date,"%d %b %Y") as "tanggal_selesai", 
		ri.risk_indicator "indikator",
		CASE 
			WHEN v.sumber_data = "Non KRID" THEN
				COALESCE ( data_verif.jumlah_data_verifikasi , 0 )
			WHEN v.sumber_data = "KRID" THEN
				COALESCE ( data_verif_krid.jumlah_data_verifikasi , 0 )
			ELSE 0
		END as jumlah_data_anomaly,
		CASE 
								WHEN v.sumber_data = "Non KRID" THEN 
									COALESCE ( perbaikan_selesai.jumlah_data_verifikasi  , 0 )
								WHEN v.sumber_data = "KRID" THEN 
									COALESCE ( perbaikan_selesai_krid.jumlah_data_verifikasi  , 0 )
								ELSE 0
		END as 'jumlah_data_sudah_verifikasi',
		CASE
		WHEN v.sumber_data = "Non KRID" THEN 
				COALESCE ( perbaikan_selesai.jumlah_data_verifikasi  , 0 )
		WHEN v.sumber_data = "KRID" THEN 
				COALESCE ( perbaikan_selesai_krid.jumlah_data_verifikasi  , 0 )
		ELSE 0 
		END as 'jumlah_data_perlu_tindaklanjut',
		0 as 'jumlah_data_sudah_tindaklanjut'
		
		from tasklists t 
		JOIN pa0001_eof pe on pe.PERNR = t.maker_id
		join tasklists_uker tu on tu.tasklist_id
		JOIN verifikasi v on v.activity_id = t.activity_id
		JOIN activity a ON a.id = t.activity_id
		JOIN product p on p.id = t.product_id
		join risk_indicator ri on ri.id = t.risk_indicator_id 
		LEFT JOIN (
					SELECT
						vda.verifikasi_id 'id',
						COUNT(vda.verifikasi_id) 'jumlah_data_verifikasi'
					FROM verifikasi_data_anomali vda
					GROUP BY vda.verifikasi_id
				) data_verif ON data_verif.id = v.id
		LEFT JOIN (
					SELECT
						vdak.verifikasi_id 'id',
						COUNT(vdak.verifikasi_id) 'jumlah_data_verifikasi'
					FROM verifikasi_data_anomali_krid vdak
					GROUP BY vdak.verifikasi_id
				) data_verif_krid ON data_verif_krid.id = v.id
		LEFT JOIN (
					SELECT
						vda.verifikasi_id 'id',
						COUNT(vda.verifikasi_id) 'jumlah_data_verifikasi'
					FROM verifikasi_data_anomali vda
					JOIN verifikasi_pic_tindak_lanjut vptl ON vptl.verifikasi_id = vda.verifikasi_id 
					WHERE vptl.status in (2)
					GROUP BY vda.verifikasi_id
				) perbaikan_selesai ON perbaikan_selesai.id = v.id
		LEFT JOIN (
					SELECT
						vdak.verifikasi_id 'id',
						COUNT(vdak.verifikasi_id) 'jumlah_data_verifikasi'
					FROM verifikasi_data_anomali_krid vdak
					JOIN verifikasi_pic_tindak_lanjut vptl ON vptl.verifikasi_id = vdak.verifikasi_id 
					WHERE vptl.status in (2)
					GROUP BY vdak.verifikasi_id
				) perbaikan_selesai_krid ON perbaikan_selesai_krid.id = v.id
		LEFT JOIN (
					SELECT
						vda.verifikasi_id 'id',
						COUNT(vda.verifikasi_id) 'jumlah_data_verifikasi'
					FROM verifikasi_data_anomali vda
					JOIN verifikasi_pic_tindak_lanjut vptl ON vptl.verifikasi_id = vda.verifikasi_id 
					WHERE vptl.status in (0,1)
					GROUP BY vda.verifikasi_id
				) perbaikan_onprogres ON perbaikan_onprogres.id = v.id
		LEFT JOIN (
					SELECT
						vdak.verifikasi_id 'id',
						COUNT(vdak.verifikasi_id) 'jumlah_data_verifikasi'
					FROM verifikasi_data_anomali_krid vdak
					JOIN verifikasi_pic_tindak_lanjut vptl ON vptl.verifikasi_id = vdak.verifikasi_id 
					WHERE vptl.status in (0,1)
					GROUP BY vdak.verifikasi_id
				) perbaikan_onprogres_krid ON perbaikan_onprogres_krid.id = v.id

                WHERE` + sQueryKanwil + sQueryKanca + sQueryUker + sQueryActivity + sQueryProduct + sQueryIndikator + sQueryStatus + sQueryJenisTask

	query := baseQuery + ` order by ` + request.Order + ` ` + request.Sort

	err = r.db.DB.Raw(query,
		strings.ToLower(request.Kanwil)+"%",
		strings.ToLower(request.Kanca)+"%",
		strings.ToLower(request.Uker)+"%",
		request.Aktifitas,
		request.Produk,
		request.Indikator,
		request.Status,
		request.JenisTaks).Scan(&responses).Error

	if err != nil {
		r.logger.Zap.Error(err)
		return responses, err
	}
	return responses, err
}

func (r LaporanRepository) GetPerhitunganPersentasePenyelesaianPerPekerjaUker(request models.PerhitunganPersentasePenyelesaianPagianted) (responses []models.LaporanPerPekerjaUkerResult, totalRow int, err error) {
	sQueryMaker := "t.maker_id = ? "

	sQueryKanwil := "AND lower(tu.RGDESC) like ? "
	sQueryKanca := "AND lower(tu.MBDESC) like ? "
	sQueryUker := "AND lower(tu.BRDESC) like ? "

	sQueryActivity := " AND t.activity_id = ? "
	sQueryProduct := " AND t.product_id = ? "
	sQueryIndikator := " AND t.risk_indicator_id = ? "

	sQueryStatus := "AND t.status = ? "
	sQueryJenisTask := "AND t.task_type = ? "

	if strings.ToLower(request.Nama) == "" {
		sQueryMaker = "t.maker_id != ? "
	}

	if request.Kanwil != "all" {
		sQueryKanwil = "AND lower(tu.RGDESC) like ? "
	}

	if strings.ToLower(request.Kanca) != "" {
		sQueryKanca = "AND lower(tu.MBDESC) like ? "
	}

	if strings.ToLower(request.Uker) != "" {
		sQueryUker = "AND lower(tu.BRDESC) like ? "
	}

	if strings.ToLower(request.Aktifitas) == "0" || strings.ToLower(request.Aktifitas) == "" {
		sQueryActivity = "AND t.activity_id != ?"
	}

	if strings.ToLower(request.Produk) == "0" || strings.ToLower(request.Produk) == "" {
		sQueryProduct = " AND t.product_id != ? "
	}

	if strings.ToLower(request.Indikator) == "0" || strings.ToLower(request.Indikator) == "" {
		sQueryIndikator = " AND t.risk_indicator_id != ? "
	}

	if strings.ToLower(request.Status) == "" {
		sQueryStatus = "AND t.status != ? "
	}

	if strings.ToLower(request.JenisTaks) == "" {
		sQueryJenisTask = "AND t.task_type != ? "
	}

	baseQuery := `SELECT 
		t.id as id,
		t.maker_id as pn,
		pe.SNAME  as nama,
		tu.RGDESC as "kanwil",
		tu.MBDESC as "kanca",
		tu.BRDESC as "uker", 
		a.name as "aktifitas", 
		p.product as "produk", 
		t.task_type as "jenis_task", 
		DATE_FORMAT(t.start_date,"%d %b %Y") as "tanggal_mulai", 
		DATE_FORMAT(t.end_date,"%d %b %Y") as "tanggal_selesai", 
		ri.risk_indicator "indikator",
		CASE 
			WHEN v.sumber_data = "Non KRID" THEN
				COALESCE ( data_verif.jumlah_data_verifikasi , 0 )
			WHEN v.sumber_data = "KRID" THEN
				COALESCE ( data_verif_krid.jumlah_data_verifikasi , 0 )
			ELSE 0
		END as jumlah_data_anomaly,
		CASE 
								WHEN v.sumber_data = "Non KRID" THEN 
									COALESCE ( perbaikan_selesai.jumlah_data_verifikasi  , 0 )
								WHEN v.sumber_data = "KRID" THEN 
									COALESCE ( perbaikan_selesai_krid.jumlah_data_verifikasi  , 0 )
								ELSE 0
		END as 'jumlah_data_sudah_verifikasi',
		CASE
		WHEN v.sumber_data = "Non KRID" THEN 
				COALESCE ( perbaikan_selesai.jumlah_data_verifikasi  , 0 )
		WHEN v.sumber_data = "KRID" THEN 
				COALESCE ( perbaikan_selesai_krid.jumlah_data_verifikasi  , 0 )
		ELSE 0 
		END as 'jumlah_data_perlu_tindaklanjut',
		0 as 'jumlah_data_sudah_tindaklanjut'
		
		from tasklists t 
		JOIN pa0001_eof pe on pe.PERNR = t.maker_id
		join tasklists_uker tu on tu.tasklist_id
		JOIN verifikasi v on v.activity_id = t.activity_id
		JOIN activity a ON a.id = t.activity_id
		JOIN product p on p.id = t.product_id
		join risk_indicator ri on ri.id = t.risk_indicator_id 
		LEFT JOIN (
					SELECT
						vda.verifikasi_id 'id',
						COUNT(vda.verifikasi_id) 'jumlah_data_verifikasi'
					FROM verifikasi_data_anomali vda
					GROUP BY vda.verifikasi_id
				) data_verif ON data_verif.id = v.id
		LEFT JOIN (
					SELECT
						vdak.verifikasi_id 'id',
						COUNT(vdak.verifikasi_id) 'jumlah_data_verifikasi'
					FROM verifikasi_data_anomali_krid vdak
					GROUP BY vdak.verifikasi_id
				) data_verif_krid ON data_verif_krid.id = v.id
		LEFT JOIN (
					SELECT
						vda.verifikasi_id 'id',
						COUNT(vda.verifikasi_id) 'jumlah_data_verifikasi'
					FROM verifikasi_data_anomali vda
					JOIN verifikasi_pic_tindak_lanjut vptl ON vptl.verifikasi_id = vda.verifikasi_id 
					WHERE vptl.status in (2)
					GROUP BY vda.verifikasi_id
				) perbaikan_selesai ON perbaikan_selesai.id = v.id
		LEFT JOIN (
					SELECT
						vdak.verifikasi_id 'id',
						COUNT(vdak.verifikasi_id) 'jumlah_data_verifikasi'
					FROM verifikasi_data_anomali_krid vdak
					JOIN verifikasi_pic_tindak_lanjut vptl ON vptl.verifikasi_id = vdak.verifikasi_id 
					WHERE vptl.status in (2)
					GROUP BY vdak.verifikasi_id
				) perbaikan_selesai_krid ON perbaikan_selesai_krid.id = v.id
		LEFT JOIN (
					SELECT
						vda.verifikasi_id 'id',
						COUNT(vda.verifikasi_id) 'jumlah_data_verifikasi'
					FROM verifikasi_data_anomali vda
					JOIN verifikasi_pic_tindak_lanjut vptl ON vptl.verifikasi_id = vda.verifikasi_id 
					WHERE vptl.status in (0,1)
					GROUP BY vda.verifikasi_id
				) perbaikan_onprogres ON perbaikan_onprogres.id = v.id
		LEFT JOIN (
					SELECT
						vdak.verifikasi_id 'id',
						COUNT(vdak.verifikasi_id) 'jumlah_data_verifikasi'
					FROM verifikasi_data_anomali_krid vdak
					JOIN verifikasi_pic_tindak_lanjut vptl ON vptl.verifikasi_id = vdak.verifikasi_id 
					WHERE vptl.status in (0,1)
					GROUP BY vdak.verifikasi_id
				) perbaikan_onprogres_krid ON perbaikan_onprogres_krid.id = v.id

                WHERE
                ` + sQueryMaker + sQueryKanwil + sQueryKanca + sQueryUker +
		sQueryActivity + sQueryProduct + sQueryIndikator + sQueryStatus + sQueryJenisTask

	query := baseQuery + ` order by ` + request.Order + ` ` + request.Sort + ` LIMIT ? OFFSET ?`
	queryCount := `select count(*) from (` + baseQuery + `) as totalData`

	err = r.db.DB.Raw(query,
		request.Nama,
		strings.ToLower(request.Kanwil)+"%",
		strings.ToLower(request.Kanca)+"%",
		strings.ToLower(request.Uker)+"%",
		request.Aktifitas,
		request.Produk,
		request.Indikator,
		request.Status,
		request.JenisTaks,
		request.Limit,
		request.Offset).Scan(&responses).Error

	err = r.db.DB.Raw(queryCount,
		request.Nama,
		strings.ToLower(request.Kanwil)+"%",
		strings.ToLower(request.Kanca)+"%",
		strings.ToLower(request.Uker)+"%",
		request.Aktifitas,
		request.Produk,
		request.Indikator,
		request.Status,
		request.JenisTaks).Scan(&totalRow).Error

	if err != nil {
		r.logger.Zap.Error(err)
		return responses, 0, err
	}
	return responses, totalRow, err
}

func (r LaporanRepository) GetPerhitunganPersentasePenyelesaianPerPekerjaUkerDownload(request models.PerhitunganPersentasePenyelesaianDownload) (responses []models.LaporanPerPekerjaUkerResult, err error) {
	sQueryMaker := "t.maker_id = ? "

	sQueryKanwil := "AND lower(tu.RGDESC) like ? "
	sQueryKanca := "AND lower(tu.MBDESC) like ? "
	sQueryUker := "AND lower(tu.BRDESC) like ? "

	sQueryActivity := " AND t.activity_id = ? "
	sQueryProduct := " AND t.product_id = ? "
	sQueryIndikator := " AND t.risk_indicator_id = ? "

	sQueryStatus := "AND t.status = ? "
	sQueryJenisTask := "AND t.task_type = ? "

	if strings.ToLower(request.Nama) == "" {
		sQueryMaker = "t.maker_id != ? "
	}

	if request.Kanwil != "all" {
		sQueryKanwil = "AND lower(tu.RGDESC) like ? "
	}

	if strings.ToLower(request.Kanca) != "" {
		sQueryKanca = "AND lower(tu.MBDESC) like ? "
	}

	if strings.ToLower(request.Uker) != "" {
		sQueryUker = "AND lower(tu.BRDESC) like ? "
	}

	if strings.ToLower(request.Aktifitas) == "0" || strings.ToLower(request.Aktifitas) == "" {
		sQueryActivity = "AND t.activity_id != ?"
	}

	if strings.ToLower(request.Produk) == "0" || strings.ToLower(request.Produk) == "" {
		sQueryProduct = " AND t.product_id != ? "
	}

	if strings.ToLower(request.Indikator) == "0" || strings.ToLower(request.Indikator) == "" {
		sQueryIndikator = " AND t.risk_indicator_id != ? "
	}

	if strings.ToLower(request.Status) == "" {
		sQueryStatus = "AND t.status != ? "
	}

	if strings.ToLower(request.JenisTaks) == "" {
		sQueryJenisTask = "AND t.task_type != ? "
	}

	baseQuery := `SELECT 
		t.id as id,
		t.maker_id as pn,
		pe.SNAME  as nama,
		tu.RGDESC as "kanwil",
		tu.MBDESC as "kanca",
		tu.unit_kerja as "uker", 
		a.name as "aktifitas", 
		p.product as "produk", 
		t.task_type_name as "jenis_task", 
		DATE_FORMAT(t.start_date,"%d %b %Y") as "tanggal_mulai", 
		DATE_FORMAT(t.end_date,"%d %b %Y") as "tanggal_selesai", 
		ri.risk_indicator "indikator",
		CASE 
			WHEN v.sumber_data = "Non KRID" THEN
				COALESCE ( data_verif.jumlah_data_verifikasi , 0 )
			WHEN v.sumber_data = "KRID" THEN
				COALESCE ( data_verif_krid.jumlah_data_verifikasi , 0 )
			ELSE 0
		END as jumlah_data_anomaly,
		CASE 
								WHEN v.sumber_data = "Non KRID" THEN 
									COALESCE ( perbaikan_selesai.jumlah_data_verifikasi  , 0 )
								WHEN v.sumber_data = "KRID" THEN 
									COALESCE ( perbaikan_selesai_krid.jumlah_data_verifikasi  , 0 )
								ELSE 0
		END as 'jumlah_data_sudah_verifikasi',
		CASE
		WHEN v.sumber_data = "Non KRID" THEN 
				COALESCE ( perbaikan_selesai.jumlah_data_verifikasi  , 0 )
		WHEN v.sumber_data = "KRID" THEN 
				COALESCE ( perbaikan_selesai_krid.jumlah_data_verifikasi  , 0 )
		ELSE 0 
		END as 'jumlah_data_perlu_tindaklanjut',
		0 as 'jumlah_data_sudah_tindaklanjut'
		
		from tasklists t 
		JOIN pa0001_eof pe on pe.PERNR = t.maker_id
		join tasklists_uker tu on tu.tasklist_id
		JOIN verifikasi v on v.activity_id = t.activity_id
		JOIN activity a ON a.id = t.activity_id
		JOIN product p on p.id = t.product_id
		join risk_indicator ri on ri.id = t.risk_indicator_id 
		LEFT JOIN (
					SELECT
						vda.verifikasi_id 'id',
						COUNT(vda.verifikasi_id) 'jumlah_data_verifikasi'
					FROM verifikasi_data_anomali vda
					GROUP BY vda.verifikasi_id
				) data_verif ON data_verif.id = v.id
		LEFT JOIN (
					SELECT
						vdak.verifikasi_id 'id',
						COUNT(vdak.verifikasi_id) 'jumlah_data_verifikasi'
					FROM verifikasi_data_anomali_krid vdak
					GROUP BY vdak.verifikasi_id
				) data_verif_krid ON data_verif_krid.id = v.id
		LEFT JOIN (
					SELECT
						vda.verifikasi_id 'id',
						COUNT(vda.verifikasi_id) 'jumlah_data_verifikasi'
					FROM verifikasi_data_anomali vda
					JOIN verifikasi_pic_tindak_lanjut vptl ON vptl.verifikasi_id = vda.verifikasi_id 
					WHERE vptl.status in (2)
					GROUP BY vda.verifikasi_id
				) perbaikan_selesai ON perbaikan_selesai.id = v.id
		LEFT JOIN (
					SELECT
						vdak.verifikasi_id 'id',
						COUNT(vdak.verifikasi_id) 'jumlah_data_verifikasi'
					FROM verifikasi_data_anomali_krid vdak
					JOIN verifikasi_pic_tindak_lanjut vptl ON vptl.verifikasi_id = vdak.verifikasi_id 
					WHERE vptl.status in (2)
					GROUP BY vdak.verifikasi_id
				) perbaikan_selesai_krid ON perbaikan_selesai_krid.id = v.id
		LEFT JOIN (
					SELECT
						vda.verifikasi_id 'id',
						COUNT(vda.verifikasi_id) 'jumlah_data_verifikasi'
					FROM verifikasi_data_anomali vda
					JOIN verifikasi_pic_tindak_lanjut vptl ON vptl.verifikasi_id = vda.verifikasi_id 
					WHERE vptl.status in (0,1)
					GROUP BY vda.verifikasi_id
				) perbaikan_onprogres ON perbaikan_onprogres.id = v.id
		LEFT JOIN (
					SELECT
						vdak.verifikasi_id 'id',
						COUNT(vdak.verifikasi_id) 'jumlah_data_verifikasi'
					FROM verifikasi_data_anomali_krid vdak
					JOIN verifikasi_pic_tindak_lanjut vptl ON vptl.verifikasi_id = vdak.verifikasi_id 
					WHERE vptl.status in (0,1)
					GROUP BY vdak.verifikasi_id
				) perbaikan_onprogres_krid ON perbaikan_onprogres_krid.id = v.id

                WHERE` + sQueryMaker + sQueryKanwil + sQueryKanca + sQueryUker +
		sQueryActivity + sQueryProduct + sQueryIndikator + sQueryStatus + sQueryJenisTask

	query := baseQuery + ` order by ` + request.Order + ` ` + request.Sort

	err = r.db.DB.Raw(query,
		request.Nama,
		strings.ToLower(request.Kanwil)+"%",
		strings.ToLower(request.Kanca)+"%",
		strings.ToLower(request.Uker)+"%",
		request.Aktifitas,
		request.Produk,
		request.Indikator,
		request.Status,
		request.JenisTaks).Scan(&responses).Error

	if err != nil {
		r.logger.Zap.Error(err)
		return responses, err
	}
	return responses, err
}

func (r LaporanRepository) GetListVerifikasiByTaskID(request models.HistoriTaskDataVerifikasiDetailRequest) (response []models.HistoriTaskDataVerifikasiDetailResult, err error) {
	db := r.db.DB

	query := db.Table("verifikasi v").
		Select(`
			v.id,
			v.BRANCH "branch",
			v.BRDESC "unit_kerja",
			v.MBDESC "kanca",
			v.RGDESC "kanwil",
			v.no_pelaporan "no_pelaporan",
			a.name "aktifitas",
			sa.name "sub_aktifitas",
			vptl2.deskripsi_tindak_lanjut as "informasi_lainnya",
			CASE
					WHEN vptl2.status = 0 THEN "Backlog" 
					WHEN vptl2.status = 1 THEN "On Progress"
					WHEN vptl2.status = 2 THEN "Selesai"
			END as "status_perbaikan",
			v.last_maker_desc "maker",
			ri.risk_issue_code "risk_issue_id",
			ri.risk_issue "risk_issue",
			v.hasil_verifikasi "hasil_verifikasi",
			COALESCE ( data_anomali.jumlah_data_anomali, 0 ) as "jumlah_data_anomali",
			CASE
					WHEN v.perbaikan = 0 THEN "Tidak"
					ELSE "Ya"
			END as "butuh_perbaikan",
			CASE
					WHEN
							v.perbaikan = 0 THEN 0
					ELSE
						COALESCE (data_anomali.jumlah_data_anomali, 0)
			END as "yang_harus_diperbaiki",
			CASE
					WHEN v.perbaikan = 0 THEN
						COALESCE ( data_anomali.jumlah_data_anomali, 0 )
					ELSE
						COALESCE ( perbaikan_selesai.jumlah_data_verifikasi  , 0 )
			END as "status_perbaikan_selesai",
			CASE
					WHEN v.perbaikan = 0 THEN 0
					ELSE
					COALESCE ( perbaikan_onprogres.jumlah_data_verifikasi  , 0 )
			END as status_perbaikan_proses,
			DATE_FORMAT(v.tanggal_target_selesai ,"%d-%b-%Y") as "batas_waktu_perbaikan",
			CASE
					WHEN v.indikasi_fraud  = 0 THEN "Tidak"
					ELSE "Ya"
			END as "indikasi_fraud",
			v.rencana_tindak_lanjut "rtl_uker"
		`).
		Joins(`LEFT JOIN activity a on v.activity_id = a.id`).
		Joins(`LEFT JOIN sub_activity sa on v.sub_activity_id = sa.id`).
		Joins(`LEFT JOIN risk_issue ri ON ri.id = v.risk_issue_id`).
		Joins(`LEFT JOIN verifikasi_pic_tindak_lanjut vptl2 ON vptl2.verifikasi_id = v.id`).
		Joins(`
			LEFT JOIN (
				SELECT
				vdat.verifikasi_id 'id',
				COUNT(vdat.verifikasi_id) 'jumlah_data_anomali'
				FROM verifikasi_data_anomali_tematik vdat
				GROUP BY vdat.verifikasi_id
			) data_anomali ON data_anomali.id = v.id 
		`).
		Joins(`
			LEFT JOIN (
					SELECT
							vdat.verifikasi_id 'id',
							COUNT(vdat.verifikasi_id) 'jumlah_data_verifikasi'
					FROM verifikasi_data_anomali_tematik vdat
					JOIN verifikasi_pic_tindak_lanjut vptl ON vptl.verifikasi_id = vdat.verifikasi_id
					WHERE vptl.status in (0,1)
					GROUP BY vdat.verifikasi_id
			) perbaikan_onprogres ON perbaikan_onprogres.id = v.id 
		`).
		Joins(`
			LEFT JOIN (
					SELECT
							vdat.verifikasi_id 'id',
							COUNT(vdat.verifikasi_id) 'jumlah_data_verifikasi'
					FROM verifikasi_data_anomali_tematik vdat
					JOIN verifikasi_pic_tindak_lanjut vptl ON vptl.verifikasi_id = vdat.verifikasi_id
					WHERE vptl.status in (2)
					GROUP BY vdat.verifikasi_id
			) perbaikan_selesai ON perbaikan_selesai.id = v.id 
		`).Group(`v.id`)

	tblLampiran := "lampiran_rap_" + request.RiskIssue + "_" + request.RiskIndicator

	subQuery := db.Table(`tasklists t`).
		Select(`t.*, tu.REGION, tu.MAINBR, tu.BRANCH, lampiran.no_verifikasi`).
		Joins(`JOIN tasklists_uker tu on t.id = tu.tasklist_id`).
		Joins(`JOIN `+tblLampiran+` lampiran on lampiran.tasklist_id = t.id`).
		Where(`t.id = ? `, request.ID).
		Where(`tu.REGION = ? `, request.Region).
		Where(`tu.MAINBR = ? `, request.Mainbr).
		Where(`tu.BRANCH = ? `, request.Branch)

	query = query.Joins(`JOIN (?) 
			tasklist ON tasklist.activity_id = v.activity_id
			and tasklist.product_id = v.product_id
			and tasklist.risk_issue_id = v.risk_issue_id
			and tasklist.REGION = v.REGION
			and tasklist.MAINBR = v.MAINBR
			and tasklist.BRANCH = v.BRANCH
			and tasklist.no_verifikasi = v.no_pelaporan`, subQuery)

	query.Scan(&response)
	err = query.Error

	if err != nil {
		r.logger.Zap.Error(err)
		return response, err
	}
	return response, err
}

func (r LaporanRepository) GetTasklist(request models.HistoriTaskDataVerifikasiPagianted) (responses []models.HistoriTaskDataVerifikasiResult, totalRow int, totalData int, err error) {
	db := r.db.DB

	query := db.Table("tasklists").
		Select(`tasklists.id AS "id",
		tasklists.no_tasklist as "no_tasklist",
		tasklists.nama_tasklist as "nama_tasklist",
		tasklists.maker_id as "pn",
		pa0001_eof.SNAME as "nama",
		tasklists_uker.REGION as "region",
		tasklists_uker.RGDESC as "kanwil",
		tasklists_uker.MAINBR as "mainbr",
		tasklists_uker.MBDESC as "kanca",
		tasklists_uker.BRANCH as "branch",
		tasklists_uker.BRDESC as "uker",
		activity.name as "aktifitas",
		tasklists.product_name as "product",
		admin_setting.kegiatan as "kegiatan",
		admin_setting.task_type as "jenis_task",
		admin_setting.period as "period",
		CASE 
				WHEN admin_setting.period = 'Custom' THEN DATE_FORMAT(tasklists.start_date,"%d-%b-%Y")
				WHEN admin_setting.period = 'Monthly' THEN DATE_FORMAT(tasklists.start_date,"%d")
				ELSE "-"
			END as "tanggal_mulai",
		CASE 
			WHEN admin_setting.period = 'Custom' THEN DATE_FORMAT(tasklists.end_date,"%d-%b-%Y")
			WHEN admin_setting.period = 'Monthly' THEN DATE_FORMAT(tasklists.end_date,"%d")
			ELSE "-"
		END as "tanggal_akhir",
		tasklists.risk_indicator "indikator",
		tasklists.approval_status as "status_approval",
		tasklists.status as "status",
		tasklists.risk_issue as "risk_issue"`).
		Joins(`LEFT JOIN pa0001_eof on pa0001_eof.PERNR = tasklists.maker_id `).
		Joins(`left join activity on activity.id = tasklists.activity_id`).
		Joins(`left JOIN tasklists_uker ON tasklists_uker.tasklist_id = tasklists.id`).
		Joins(`join admin_setting on tasklists.task_type = admin_setting.id`).
		Group("tasklists.id, tasklists_uker.BRANCH")

		// Where(`admin_setting.kegiatan = 'Verifikasi'`)

	// if strings.ToLower(request.Pn) != "" {
	// 	query = query.Where("tasklists.maker_id = ? ", request.Pn)
	// }

	if request.KANWIL != "all" && request.KANWIL != "" {
		// query = query.Where("tasklists_uker.RGDESC = ? ", request.KANWIL)
		query = query.Where("tasklists_uker.REGION = ? ", request.KANWIL)
	}

	if strings.ToLower(request.KANCA) != "" && strings.ToLower(request.KANCA) != "all" {
		// query = query.Where("tasklists_uker.MBDESC = ? ", request.KANCA)
		query = query.Where("tasklists_uker.MAINBR = ? ", request.KANCA)
	}

	if strings.ToLower(request.UnitKerja) != "" && strings.ToLower(request.UnitKerja) != "all" {
		// query = query.Where("tasklists_uker.BRDESC = ? ", request.UnitKerja)
		branches := strings.Split(request.UnitKerja, ",")
		if len(branches) > 1 {
			//lebih dari 1
			query = query.Where("tasklists_uker.BRANCH in (?) ", branches)
		} else {
			//cuma 1
			query = query.Where("tasklists_uker.BRANCH = ? ", request.UnitKerja)
		}

	}

	if strings.ToLower(request.Aktifitas) != "0" && strings.ToLower(request.Aktifitas) != "" {
		query = query.Where("tasklists.activity_id = ? ", request.Aktifitas)
	}

	if strings.ToLower(request.Produk) != "0" && strings.ToLower(request.Produk) != "" {
		query = query.Where("tasklists.product_id = ? ", request.Produk)
	}

	if strings.ToLower(request.Indikator) != "" {
		query = query.Where("tasklists.risk_indicator_id = ? ", request.Indikator)
	}

	if strings.ToLower(request.Indikator) == "0" {
		query = query.Where("LOWER(tasklists.risk_indicator) like ?", "%"+strings.ToLower(request.IndikatorOther)+"%")
	}

	if strings.ToLower(request.RiskEvent) != "0" && strings.ToLower(request.RiskEvent) != "" {
		query = query.Where("tasklists.risk_issue_id = ? ", request.RiskEvent)
	}

	if strings.ToLower(request.JenisTask) != "" {
		query = query.Where("tasklists.task_type = ? ", request.JenisTask)
	}

	if strings.ToLower(request.Status) != "" {
		query = query.Where("tasklists.status = ? ", request.Status)
	}

	if strings.ToLower(request.StatusApproval) != "" && strings.ToLower(request.StatusApproval) != "all" {
		query = query.Where("tasklists.approval_status like ?", "%"+request.StatusApproval+"%")
	}

	// if request.SumberData != "" {
	// 	query = query.Where("tasklists.sumber_data = ?", request.SumberData)
	// }

	var count int64
	query.Count(&count)

	totalRow = int(count)

	if request.Limit != 0 {
		query = query.Limit(request.Limit)
	}

	if request.Offset != 0 {
		query = query.Offset(request.Offset)
	}

	query.Scan(&responses)

	result := float64(totalRow) / float64(request.Limit)
	resultFinal := int(math.Ceil(result))

	return responses, totalRow, resultFinal, err
}

func (r LaporanRepository) GetTasklistDownload(request models.HistoriTaskDataVerifikasiDownload) (responses []models.HistoriTaskDataVerifikasiResult, err error) {
	db := r.db.DB

	query := db.Table("tasklists").
		Select(`tasklists.id AS "id",
		tasklists.maker_id as "pn",
		tasklists.no_tasklist as "no_tasklist",
		tasklists.nama_tasklist as "nama_tasklist",
		pa0001_eof.SNAME as "nama",
		tasklists_uker.RGDESC as "kanwil",
		tasklists_uker.MBDESC as "kanca",
		tasklists_uker.BRDESC as "uker",
		activity.name as "aktifitas",
		tasklists.product_name as "product",
		admin_setting.kegiatan as "kegiatan",
		admin_setting.task_type as "jenis_task",
		CASE
			WHEN admin_setting.period = 'Custom' THEN DATE_FORMAT(tasklists.start_date,"%d-%b-%Y")
			WHEN admin_setting.period = 'Monthly' THEN DATE_FORMAT(tasklists.start_date,"%d")
			ELSE "-"
		END as "tanggal_mulai",
		CASE 
			WHEN admin_setting.period = 'Custom' THEN DATE_FORMAT(tasklists.end_date,"%d-%b-%Y")
			WHEN admin_setting.period = 'Monthly' THEN DATE_FORMAT(tasklists.end_date,"%d")
			ELSE "-"
		END as "tanggal_akhir",
		tasklists.risk_indicator "indikator",
		tasklists.approval_status as "status_approval",
		tasklists.status as "status",
		tasklists.risk_issue as "risk_issue"`).
		Joins(`LEFT JOIN pa0001_eof on pa0001_eof.PERNR = tasklists.maker_id `).
		Joins(`left join activity on activity.id = tasklists.activity_id`).
		Joins(`left JOIN tasklists_uker ON tasklists_uker.tasklist_id = tasklists.id`).
		Joins(` join admin_setting on tasklists.task_type = admin_setting.id`).
		Group("tasklists.id, tasklists_uker.BRANCH")

		// Where(`admin_setting.kegiatan = 'Verifikasi'`)

	// if strings.ToLower(request.Pn) != "" {
	// 	query = query.Where("tasklists.maker_id = ? ", request.Pn)
	// }

	if request.KANWIL != "all" && request.KANWIL != "" {
		// query = query.Where("tasklists_uker.RGDESC = ? ", request.KANWIL)
		query = query.Where("tasklists_uker.REGION = ? ", request.KANWIL)
	}

	if strings.ToLower(request.KANCA) != "" && strings.ToLower(request.KANCA) != "all" {
		// query = query.Where("tasklists_uker.MBDESC = ? ", request.KANCA)
		query = query.Where("tasklists_uker.MAINBR = ? ", request.KANCA)
	}

	if strings.ToLower(request.UnitKerja) != "" && strings.ToLower(request.UnitKerja) != "all" {
		// query = query.Where("tasklists_uker.BRDESC = ? ", request.UnitKerja)
		branches := strings.Split(request.UnitKerja, ",")
		if len(branches) > 1 {
			//lebih dari 1
			query = query.Where("tasklists_uker.BRANCH in (?) ", branches)
		} else {
			//cuma 1
			query = query.Where("tasklists_uker.BRANCH = ? ", request.UnitKerja)
		}

	}

	if strings.ToLower(request.Aktifitas) != "0" && strings.ToLower(request.Aktifitas) != "" {
		query = query.Where("tasklists.activity_id = ? ", request.Aktifitas)
	}

	if strings.ToLower(request.Produk) != "0" && strings.ToLower(request.Produk) != "" {
		query = query.Where("tasklists.product_id = ? ", request.Produk)
	}

	if strings.ToLower(request.Indikator) != "" {
		query = query.Where("tasklists.risk_indicator_id = ? ", request.Indikator)
	}

	if strings.ToLower(request.Indikator) == "0" {
		query = query.Where("LOWER(tasklists.risk_indicator) like ?", "%"+strings.ToLower(request.IndikatorOther)+"%")
	}

	if strings.ToLower(request.RiskEvent) != "0" && strings.ToLower(request.RiskEvent) != "" {
		query = query.Where("tasklists.risk_issue_id = ? ", request.RiskEvent)
	}

	if strings.ToLower(request.JenisTask) != "" {
		query = query.Where("tasklists.task_type = ? ", request.JenisTask)
	}

	if strings.ToLower(request.Status) != "" {
		query = query.Where("tasklists.status = ? ", request.Status)
	}

	if strings.ToLower(request.StatusApproval) != "" && strings.ToLower(request.StatusApproval) != "all" {
		query = query.Where("tasklists.approval_status like ?", "%"+request.StatusApproval+"%")
	}

	query.Scan(&responses)

	err = query.Error

	if err != nil {
		r.logger.Zap.Error(err)
		return responses, err
	}

	return responses, err
}

func (r LaporanRepository) GetMonitoringJob(request models.JobMonitoringRequest) (responses []models.JobMonitoringResponse, totalRow int, totalData int, err error) {
	db := r.db.DB

	query := db.Table("job_logs").
		Select(`job_logs.created_at as "tanggal",
		jobs.name as "nama_job",
		job_logs.process as "proses",
		job_logs.status as "status_proses",
		job_logs.status_description as "deskripsi_status"`).
		Joins(`JOIN jobs on jobs.id = job_logs.job_id`)

	if request.StartDate != "" && request.EndDate != "" {
		query.Where("DATE(job_logs.created_at) BETWEEN ? AND ?", request.StartDate, request.EndDate)
	}

	if request.NamaJob != "" {
		query.Where("jobs.name LIKE '%" + request.NamaJob + "%'")
		// query.Where("jobs.name = ?", request.NamaJob)
	}

	if request.StatusProses != "" {
		query.Where("job_logs.status = ?", request.StatusProses)
	}

	var count int64
	query.Count(&count)

	totalRow = int(count)

	if request.Limit != 0 {
		query = query.Limit(request.Limit)
	}

	if request.Offset != 0 {
		query = query.Offset(request.Offset)
	}

	query.Scan(&responses)

	result := float64(totalRow) / float64(request.Limit)
	resultFinal := int(math.Ceil(result))

	return responses, totalRow, resultFinal, err
}

func (l LaporanRepository) GetNamaJob(request *models.SearchNamaJobReq) (responses []models.SearchNamaJobRes, err error) {
	query := `SELECT name FROM jobs j 
			where j.name LIKE '%` + request.Keyword + `%' LIMIT ?`

	// fmt.Println(query)

	l.logger.Zap.Info(query)

	rows, err := l.dbRaw.DB.Query(query, request.Limit)

	if err != nil {
		return responses, err
	}

	response := models.SearchNamaJobRes{}

	for rows.Next() {
		_ = rows.Scan(
			&response.Name,
		)

		responses = append(responses, response)
	}

	return responses, err
}

func (r LaporanRepository) GetActivityDaily(request models.ActivityDailyRequest) (responses []models.ActivityDaily, totalRow int, totalData int, err error) {
	db := r.db.DB

	query := db.Table("tasklists_history th").
		Select(`th.PERNR as "PERNR",
		th.kegiatan,
		th.activity_id,
		th.product_id,
		th.risk_issue_id,
		th.name as "nama",
		th.kanwil as "kanwil",
		th.risk_event as "risk_event",
		sum(th.persentase) / count(*) as "persentase"`)

	if request.REGION != "" {
		// query.Where("th.kanwil = ?", request.REGION)
		// query.Where("th.kanwil LIKE ?", "%"+request.REGION+"%")
		query.Where("th.REGION = ?", request.REGION)
	}

	if request.PN != "" {
		query.Where("th.PERNR = ?", request.PN)
	}

	if request.Period != "" && request.Period != "-" {
		query.Where("DATE_FORMAT(th.created_at, '%Y-%m') = ?", request.Period)
	}

	if request.Persentase != "" {
		persentase := strings.Split(request.Persentase, "-")

		persen1Int, err := strconv.Atoi(persentase[0])

		if err != nil {
			r.logger.Zap.Error(err)
			return responses, 0, 0, err
		}

		persen2Int, err := strconv.Atoi(persentase[1])

		if err != nil {
			r.logger.Zap.Error(err)
			return responses, 0, 0, err
		}

		query.Having("sum(th.persentase) / count(*) BETWEEN ? AND ?", persen1Int, persen2Int)
	} else {
		query.Having("sum(th.persentase) / count(*) < 80")
	}

	query.Group("PERNR")
	// query.Group("kegiatan")
	query.Group("activity_id")
	query.Group("product_id")
	query.Group("risk_issue_id")
	// query.Group("kanca")

	// query.Group("DATE_FORMAT(th.created_at, '%Y-%m')")

	var count int64
	query.Count(&count)

	totalRow = int(count)

	if request.Limit != 0 {
		query = query.Limit(request.Limit)
	}

	if request.Offset != 0 {
		query = query.Offset(request.Offset)
	}

	query.Scan(&responses)

	result := float64(totalRow) / float64(request.Limit)
	resultFinal := int(math.Ceil(result))

	return responses, totalRow, resultFinal, err
}

func (r LaporanRepository) GetUkerList(request models.UkerListRequest) (responses []models.UkerListResponse, err error) {
	query := `SELECT DISTINCT(BRDESC) AS 'BRDESC' FROM uker_kelolaan_user uku 
			WHERE pn = '` + request.PERNR + `'`

	// fmt.Println(query)

	r.logger.Zap.Info(query)

	rows, err := r.dbRaw.DB.Query(query)

	if err != nil {
		return responses, err
	}

	response := models.UkerListResponse{}

	for rows.Next() {
		_ = rows.Scan(
			&response.BRDESC,
		)

		responses = append(responses, response)
	}

	return responses, err
}

func (r LaporanRepository) GetTasklistPersentase(request models.PersentaseTotalRequest) (total float64, err error) {
	db := r.db.DB

	query := db.Table("tasklists_today").
		Select(`sum(persentase) / count(*) as 'total'`).
		Where(`tasklists_today.PERNR = ?`, request.PERNR).Group("PERNR")

	query.Scan(&total)

	return total, err
}

func (r LaporanRepository) GetActivityDailyDetail(request models.ActivityDailyDetailRequest) (responses []models.ActivityDailyDetail, totalRow int, totalData int, err error) {
	db := r.db.DB

	query := db.Table("tasklists_history").
		Select(`tasklists_history.PERNR as "PERNR",
		tasklists_history.name as "nama",
		tasklists_history.kanwil as "kanwil",
		tasklists_history.kanca as "kanca",
		tasklists_history.unit_kerja as "unit_kerja",
		tasklists_history.kegiatan as "kegiatan",
		tasklists_history.risk_event as "risk_event",
		tasklists.task_type_name as "task_type",
		tasklists_history.period as "period",
		tasklists_history.sample as "sample",
		tasklists_history.progres as "progres",
		tasklists_history.assigned_created as "assigned_created",
		tasklists_history.start_date as "start_date",
		tasklists_history.end_date as "end_date",
		tasklists_history.persentase as "persentase"`).
		Joins(`LEFT JOIN tasklists ON tasklists.id = tasklists_history.tasklist_id`).
		Where("tasklists_history.PERNR = ?", request.PERNR).
		// Where("tasklists_history.kegiatan = ?", request.Kegiatan).
		Where("tasklists_history.activity_id = ?", request.ActivityID).
		Where("tasklists_history.product_id = ?", request.ProductID).
		Where("tasklists_history.risk_issue_id = ?", request.RiskIssueID)

	if request.Period != "" {
		query.Where("tasklists_history.period = ?", request.Period)
	}

	if request.TaskType != 0 {
		query.Where("tasklists.task_type = ?", request.TaskType)
	}

	var count int64
	query.Count(&count)

	totalRow = int(count)

	if request.Limit != 0 {
		query = query.Limit(request.Limit)
	}

	if request.Offset != 0 {
		query = query.Offset(request.Offset)
	}

	query.Scan(&responses)

	result := float64(totalRow) / float64(request.Limit)
	resultFinal := int(math.Ceil(result))

	return responses, totalRow, resultFinal, err
}
