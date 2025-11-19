package tasklists

import (
	"fmt"
	"math"
	"riskmanagement/lib"
	models "riskmanagement/models/tasklists"
	"strconv"
	"strings"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type TasklistsDefinition interface {
	GetAll(request *models.Paginate) (response []models.TasklistsResponse, totalRows int, totalData int, err error)
	GetTaskByID(id int64) (response []models.TasklistsResponse, totalRows int, totalData int, err error)
	DoneSampleCount(request *models.CountDoneSampleRequest) (response models.CountDoneSampleResponse, err error)
	GetSubTask(id int64) (responses []models.TasklistSubTask, err error)
	GetByTaskType(request *models.Paginate) (response []models.TasklistsResponse, totalRows int, totalData int, err error)
	GetOne(id int64) (response models.TasklistDataResponse, err error)
	Filter(request models.TasklistsFilterRequest) (responses []models.TasklistsFilterApprovalResponses, totalRows int, totalData int, err error)
	GetUkerList(request models.UkerListReq) (responses []models.TasklistsUkerData, err error)
	FilterByID(request models.TasklistsFilterRequest) (responses []models.TasklistsFilterResponse, totalRows int, totalData int, err error)
	FilterOfficer(request models.TasklistsFilterOfficerRequest) (responses []models.TasklistsFilterResponses, totalRows int, totalData int, err error)
	CheckAvailability(request models.TasklistsCheckRequest) (response models.TasklistsCheckResponse, err error)
	CekRiskIssueAvail(request *models.TasklistRiskIssueAvailReq) (response models.TasklistRiskIssueAvailRes, err error)
	Store(request *models.Tasklists, tx *gorm.DB) (responses *models.Tasklists, err error)
	StoreNotif(request *models.TasklistNotif, tx *gorm.DB) (responses *models.TasklistNotif, err error)
	StoreTasklistDaily(request *models.TasklistsToday, tx *gorm.DB) (responses *models.TasklistsToday, err error)
	Update(request *models.TasklistsUpdate, tx *gorm.DB) (responses *models.TasklistsUpdate, err error)
	UpdateEndDate(request *models.TasklistsUpdateEndDate, tx *gorm.DB) (responses *models.TasklistsUpdateEndDate, err error)
	Delete(request *models.TasklistsUpdateDelete, tx *gorm.DB) (responses bool, err error)
	GetDataBRC(request *models.GetBRCRequest) (responses []models.GetBRCResponse, err error)
	Approval(request *models.TasklistsAprroval, tx *gorm.DB) (responses *models.TasklistsAprroval, err error)
	Validation(request *models.TasklistsValidation, tx *gorm.DB) (responses *models.TasklistsValidation, err error)
	CountTask(request models.TasklistCountRequest) (responses models.TasklistsCountResponse, err error)
	CountTaskDone(request models.TasklistCountRequest) (responses models.TasklistsCountResponse, err error)
	LeadAutocompleteVal(request models.LeadAutocompleteRequest) (responses []models.LeadAutocomplete, err error)
	LeadAutocompleteApr(request models.LeadAutocompleteRequest) (responses []models.LeadAutocomplete, err error)
	LeadAutocompleteKanwil(request models.LeadAutocompleteRequest) (responses []models.LeadAutocomplete, err error)
	UserRegion(request models.UserRegionRequest) (responses models.UserRegionResponse, err error)
	GetAllOfficer(request models.Paginate) (responses []models.TasklistsFilterResponses, totalRows int, totalData int, err error)
	GetDataVerifikasi(request models.DataVerifikasiRequest) (responses []models.DataVerifikasiResponse, err error)
	CreateTableLampiran(request *models.TasklistHeaderStore) (err error)
	InsertTableLampiran(request *models.TasklistColumnStore) (err error)
	InsertTableLampiranIndikator(request *models.LampiranIndikatorStore) (err error)
	GetLampiranIndicator(request *models.LampiranIndikatorCheck) (response models.LampiranIndikatorResponse, err error)
	DownloadLampiranIndikatorTemplate(request *models.LampiranIndikatorCheck) (response models.AnomaliHeaderResponse, err error)
	ShowUker(id int64) (responses []models.TasklistsUker, err error)
	InsertTasklistRejected(tasklistRejected *models.TasklistsRejected, tx *gorm.DB) (err error)
	GetNotes(id int64) (response models.TasklistRejectedNote, err error)

	LimitTask(id int64) (maxLimit int, err error)
	CheckMaxLimit(TaskID int64, BranchID string) (total int, err error)

	GetAnomaliHeader(request *models.AnomaliHeader) (response models.AnomaliHeaderResponse, err error)
	GetAnomaliValue(request *models.AnomaliValue) (response []models.AnomaliValueResponse, err error)

	GetFirstLampiran(request *models.GetFirstLampiranRequest) (response models.GetFirstLampiranResponse, err error)

	DeleteIsiLampiran(request *models.TasklistLampiranDelete, tx *gorm.DB) (err error)

	GetTipeUkerMaker(pnRequest string) (response models.PegawaiData, err error)
}

type TasklistsRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewTasklistsRepository(db lib.Database, dbRaw lib.Databases, logger logger.Logger) TasklistsDefinition {
	return TasklistsRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

func (r TasklistsRepository) GetAll(request *models.Paginate) (response []models.TasklistsResponse, totalRows int, totalData int, err error) {
	// rows, err := r.db.DB.Raw(`
	// 			SELECT tasklists.id AS "tasklist_id", tasklists.activity_id AS "activity_id",
	// 			activity.name AS "activity", tasklists.product_id AS "product_id", tasklists.*,
	// 			risk_issue.risk_issue_code "risk_issue_code", risk_issue.risk_issue, risk_indicator.risk_indicator, tasklists_uker.*,
	// 			tasklists_uker.branch AS "branch", product.product, admin_setting.kegiatan AS "kegiatan",
	// 			admin_setting.period AS "jenis_task"
	// 			FROM tasklists
	// 			JOIN admin_setting ON admin_setting.id = tasklists.task_type
	// 			JOIN risk_issue ON risk_issue.id = tasklists.risk_issue_id
	// 			LEFT JOIN risk_indicator ON risk_indicator.id = tasklists.risk_indicator_id
	// 			JOIN tasklists_uker ON tasklists_uker.tasklist_id = tasklists.id
	// 			JOIN activity ON activity.id = tasklists.activity_id
	// 			JOIN product ON product.id = tasklists.product_id
	// 			WHERE tasklists_uker.BRANCH = ?
	// 			AND tasklists.status = "Aktif" AND approval_status = "Disetujui"
	// 			ORDER BY tasklists.id ASC LIMIT ? OFFSET ?`, request.BRANCH, request.Limit, request.Offset).Rows()
	rows, err := r.db.DB.Raw(`SELECT tt.tasklist_id AS "tasklist_id", tt.activity_id AS "activity_id", 
				activity.name AS "activity", tt.product_id AS "product_id", tt.*, 
				tt.risk_issue, tt.risk_indicator, tt.risk_issue_id, tt.risk_indicator_id, tt.REGION,
				tt.RGDESC, tt.MAINBR, tt.MBDESC, tt.BRANCH AS "branch", tt.BRDESC,
				tt.progres AS "done_sample",
				tt.product, tt.kegiatan AS "kegiatan", 
				CONCAT(tt.task_type_name, " - ", tt.period) AS "jenis_task"
				FROM tasklists_today tt
				JOIN activity ON activity.id = tt.activity_id
				WHERE tt.PERNR = ?
				ORDER BY tt.id ASC LIMIT ? OFFSET ?`, request.PERNR, request.Limit, request.Offset).Rows()

	defer rows.Scan()

	var tasklists models.TasklistsResponse

	for rows.Next() {
		r.db.DB.ScanRows(rows, &tasklists)
		response = append(response, tasklists)
	}

	paginateQuery := fmt.Sprintf(`SELECT COUNT(*) FROM tasklists_today tt
								JOIN activity ON activity.id = tt.activity_id
								WHERE tt.PERNR = %s`, request.PERNR)

	err = r.dbRaw.DB.QueryRow(paginateQuery).Scan(&totalRows)

	result := float64(totalRows) / float64(request.Limit)
	resultFinal := int(math.Ceil(result))

	return response, resultFinal, totalRows, err
}

func (r TasklistsRepository) GetTaskByID(id int64) (response []models.TasklistsResponse, totalRows int, totalData int, err error) {
	// rows, err := r.db.DB.Raw(`
	// 	SELECT tt.tasklist_id AS "tasklist_id", tt.activity_id AS "activity_id",
	// 	activity.name AS "activity", tt.product_id AS "product_id", tt.*,
	// 	tt.risk_issue, tt.risk_indicator, tt.risk_issue_id, tt.risk_indicator_id, tt.REGION,
	// 	tt.RGDESC, tt.MAINBR, tt.MBDESC, tt.BRANCH AS "branch", tt.BRDESC,
	// 	tt.progres AS "done_sample",
	// 	tt.product, tt.kegiatan AS "kegiatan",
	// 	CONCAT(tt.task_type_name, " - ", tt.period) AS "jenis_task"
	// 	FROM tasklists_today tt
	// 	JOIN activity ON activity.id = tt.activity_id
	// 	WHERE tt.tasklist_id = ? ORDER BY tt.id ASC LIMIT ? OFFSET ?`, id, 1, 0).Rows()
	rows, err := r.db.DB.Raw(`
		SELECT tt.tasklist_id AS "tasklist_id", tt.activity_id AS "activity_id", 
		activity.name AS "activity", tt.product_id AS "product_id", tt.*, 
		tt.risk_issue, tt.risk_indicator, tt.risk_issue_id, tt.risk_indicator_id, tt.REGION,
		tt.RGDESC, tt.MAINBR, tt.MBDESC, tt.BRANCH AS "branch", tt.BRDESC,
		tt.progres AS "done_sample",
		tt.product, tt.kegiatan AS "kegiatan", 
		CONCAT(tt.task_type_name, " - ", tt.period) AS "jenis_task"
		FROM tasklists_today tt
		JOIN activity ON activity.id = tt.activity_id
		WHERE tt.id = ? ORDER BY tt.id ASC LIMIT ? OFFSET ?`, id, 1, 0).Rows()

	defer rows.Scan()

	var tasklists models.TasklistsResponse

	for rows.Next() {
		r.db.DB.ScanRows(rows, &tasklists)
		response = append(response, tasklists)
	}

	// paginateQuery := fmt.Sprintf(`SELECT count(*) FROM tasklists_today tt
	// 	JOIN activity ON activity.id = tt.activity_id
	// 	WHERE tt.tasklist_id = %d`, id)

	paginateQuery := fmt.Sprintf(`SELECT count(*) FROM tasklists_today tt
		JOIN activity ON activity.id = tt.activity_id
		WHERE tt.id = %d`, id)

	err = r.dbRaw.DB.QueryRow(paginateQuery).Scan(&totalRows)

	result := float64(totalRows) / 1
	resultFinal := int(math.Ceil(result))

	return response, resultFinal, totalRows, err
}

func (r TasklistsRepository) DoneSampleCount(request *models.CountDoneSampleRequest) (response models.CountDoneSampleResponse, err error) {
	if request.Kegiatan == "Verifikasi" {
		err = r.db.DB.Raw(`select count(*) as total from verifikasi v 
					join tasklists t on t.risk_issue_id  = v.risk_issue_id 
					join tasklists_uker tu on tu.tasklist_id = t.id 
					where v.activity_id = ? and v.product_id = ? and v.risk_issue_id = ?
					and t.activity_id = ? and t.product_id = ? and t.risk_issue_id = ?
					and tu.BRANCH = ? and t.status = "Aktif" and t.approval_status = 'Disetujui'`,
			request.ActivityID, request.ProductID, request.RiskIssueID,
			request.ActivityID, request.ProductID, request.RiskIssueID,
			request.Branch).Find(&response).Error
	} else if request.Kegiatan == "Briefing" {
		err = r.db.DB.Raw(`select count(*) as total from briefing_materis bm
						join briefing b on b.id = bm.briefing_id 
						where activity_id = ? and product_id = ? and risk_issue_code = ? and branch = ?`,
			request.ActivityID, request.ProductID, request.RiskIssueCode, request.Branch).Find(&response).Error
	} else if request.Kegiatan == "Coaching" {
		err = r.db.DB.Raw(`select count(distinct risk_issue_id) as total from coaching_activity ca 
						join coaching c on c.id = ca.coaching_id 
						WHERE activity_id = ? and product_id = ? and branch = ?`,
			request.ActivityID, request.ProductID, request.Branch).Find(&response).Error
	}

	if err != nil {
		r.logger.Zap.Error(err)
		return response, err
	}
	return response, err
}

func (r TasklistsRepository) GetSubTask(id int64) (responses []models.TasklistSubTask, err error) {
	err = r.db.DB.Raw(`
		SELECT risk_indicator.risk_indicator FROM tasklists_risk_indicator
		LEFT JOIN risk_indicator ON risk_indicator.id = tasklists_risk_indicator.risk_indicator_id 
		WHERE tasklists_id = ?`, id).Find(&responses).Error

	if err != nil {
		r.logger.Zap.Error(err)
		return responses, err
	}
	return responses, err
}

func (r TasklistsRepository) GetByTaskType(request *models.Paginate) (response []models.TasklistsResponse, totalRows int, totalData int, err error) {
	// if request.Period == "RAP" {
	// 	rows, err := r.db.DB.Raw(`
	// 			SELECT tasklists.id AS "tasklist_id", tasklists.*, risk_issue.risk_issue, risk_issue.risk_issue_code "risk_issue_code", risk_indicator.risk_indicator, tasklists_uker.*,
	// 			admin_setting.period AS "jenis_task"
	// 			FROM tasklists
	// 			JOIN risk_issue ON risk_issue.id = tasklists.risk_issue_id
	// 			LEFT JOIN risk_indicator ON risk_indicator.id = tasklists.risk_indicator_id
	// 			JOIN tasklists_uker ON tasklists_uker.tasklist_id = tasklists.id
	// 			JOIN admin_setting ON admin_setting.id = tasklists.task_type
	// 			WHERE tasklists.rap = 1 AND tasklists_uker.branch = ? AND tasklists.status = 'Aktif' AND approval_status = 'Disetujui'
	// 			ORDER BY tasklists.id ASC LIMIT ? OFFSET ?`,
	// 		request.BRANCH, request.Limit, request.Offset).Rows()

	// 	defer rows.Scan()

	// 	var tasklists models.TasklistsResponse

	// 	for rows.Next() {
	// 		r.db.DB.ScanRows(rows, &tasklists)
	// 		response = append(response, tasklists)
	// 	}

	// 	paginateQuery := fmt.Sprintf(`SELECT count(*)
	// 				FROM tasklists
	// 				JOIN risk_issue ON risk_issue.id = tasklists.risk_issue_id
	// 				JOIN tasklists_uker ON tasklists_uker.tasklist_id = tasklists.id
	// 				JOIN admin_setting ON admin_setting.id = tasklists.task_type
	// 				WHERE tasklists.rap = 1 AND tasklists_uker.BRANCH = %d
	// 				AND tasklists.status = "Aktif" AND approval_status = "Disetujui"`, request.BRANCH)
	// 	err = r.dbRaw.DB.QueryRow(paginateQuery).Scan(&totalRows)

	// 	result := float64(totalRows) / float64(request.Limit)
	// 	resultFinal := int(math.Ceil(result))

	// 	return response, resultFinal, totalRows, err
	// } else {
	rows, err := r.db.DB.Raw(`
				SELECT tt.tasklist_id AS "tasklist_id", tt.activity_id AS "activity_id", 
				activity.name AS "activity", tt.product_id AS "product_id", tt.*, 
				tt.risk_issue, tt.risk_indicator, tt.REGION,
				tt.RGDESC, tt.MAINBR, tt.MBDESC, tt.BRANCH AS "branch", tt.BRDESC,
				tt.progres AS "done_sample",
				tt.product, tt.kegiatan AS "kegiatan", 
				tt.period AS "jenis_task"
				FROM tasklists_today tt
				JOIN activity ON activity.id = tt.activity_id
				WHERE tt.PERNR = ? AND tt.task_type = ? 
				ORDER BY tt.id ASC LIMIT ? OFFSET ?`,
		request.PERNR, request.Period, request.Limit, request.Offset).Rows()

	defer rows.Close()
	defer rows.Scan()

	var tasklists models.TasklistsResponse

	for rows.Next() {
		r.db.DB.ScanRows(rows, &tasklists)
		response = append(response, tasklists)
	}

	paginateQuery := fmt.Sprintf(`SELECT count(*)
					FROM tasklists_today 
					WHERE task_type = "%s" AND PERNR = "%s"`, request.Period, request.PERNR)

	fmt.Println("isi query:", paginateQuery)
	err = r.dbRaw.DB.QueryRow(paginateQuery).Scan(&totalRows)

	result := float64(totalRows) / float64(request.Limit)
	resultFinal := int(math.Ceil(result))

	return response, resultFinal, totalRows, err
	// }
}

func (r TasklistsRepository) GetAllOfficer(request models.Paginate) (responses []models.TasklistsFilterResponses, totalRows int, totalData int, err error) {
	db := r.db.DB

	query := db.Table("tasklists")
	queryCount := db.Table("tasklists")

	query = query.Select(`
			DISTINCT
			tasklists.id 'id', 
			product.product,
			activity.name 'activity',
			tasklists.task_type,
			tasklists.task_type_name 'jenis_task', 
			IFNULL(tasklists.start_date, 'NULL') AS 'start_date',
			IFNULL(tasklists.end_date, 'NULL') AS 'end_date',
			tasklists.risk_issue 'risk_issue', 
			tasklists.risk_indicator 'risk_indicator', 
			tasklists.status, tasklists.approval, 
			tasklists.validation, 
			tasklists.approval_status 'status_approval',
			tasklists.maker_id 'maker_id', 
			tasklists.period 'period'
		`).
		Joins(`LEFT JOIN tasklists_uker ON tasklists_uker.tasklist_id = tasklists.id`).
		Joins(`left join activity on activity.id = tasklists.activity_id`).
		Joins(`left join product on product.id = tasklists.product_id`).
		Where(`tasklists.status = 'Aktif'`)

	queryCount = queryCount.Select(`
			COUNT(DISTINCT tasklists.id) AS 'total'
		`).Where(`tasklists.status = 'Aktif'`).
		Joins(`LEFT JOIN tasklists_uker ON tasklists_uker.tasklist_id = tasklists.id`)

	// if request.TipeUker != "KP" && request.TipeUker != "KW" {
	if request.HILFM == "033" || request.HILFM == "034" || request.HILFM == "228" {
		query = query.Where(`tasklists_uker.BRANCH in (select branch from uker_kelolaan_user where pn = ?)`, request.PERNR)
		queryCount = queryCount.Where(`tasklists_uker.BRANCH in (select branch from uker_kelolaan_user where pn = ?)`, request.PERNR)
	}

	if strings.Contains(request.StellTX, "RISK MANAGEMENT & COMPLIANCE") {
		query = query.Where(`tasklists_uker.REGION = ?`, request.REGION)
		queryCount = queryCount.Where(`tasklists_uker.REGION = ?`, request.REGION)
	}

	var count int64
	queryCount.Find(&count)

	totalData = int(count)
	query.Limit(request.Limit).Offset(request.Offset).Find(&responses)

	totalPages := int(math.Ceil(float64(totalData) / float64(request.Limit)))

	return responses, totalPages, totalData, err
}

func (r TasklistsRepository) Filter(request models.TasklistsFilterRequest) (responses []models.TasklistsFilterApprovalResponses, totalRows int, totalData int, err error) {
	db := r.db.DB
	query := db.Table("tasklists")
	queryCount := db.Table("tasklists")

	// query = query.Select(`
	// 		tasklists.id 'id',
	// 		lam.id 'id_lampiran',
	// 		lam.tasklists_id 'tasklists_id',
	// 		fl.filename 'filename',
	// 		fl.path 'path',
	// 		fl.extension 'ext',
	// 		fl.size 'size',
	// 		tasklists_uker.REGION,
	// 		tasklists_uker.RGDESC,
	// 		tasklists_uker.RGNAME,
	// 		tasklists_uker.MAINBR,
	// 		tasklists_uker.MBDESC,
	// 		tasklists_uker.MBNAME,
	// 		tasklists_uker.BRANCH 'branch',
	// 		tasklists_uker.BRDESC,
	// 		tasklists_uker.BRNAME,
	// 		tasklists_uker.unit_kerja,
	// 		activity.name 'activity',
	// 		product.product,
	// 		tasklists.task_type,
	// 		tasklists.task_type_name 'jenis_task',
	// 		tasklists.start_date,
	// 		tasklists.end_date,
	// 		tasklists.risk_issue 'risk_issue',
	// 		tasklists.risk_indicator 'risk_indicator',
	// 		tasklists.status, tasklists.approval,
	// 		tasklists.validation,
	// 		tasklists.approval_status 'status_approval',
	// 		tasklists.maker_id 'maker_id',
	// 		tasklists.period 'period'`).

	query = query.Select(`
			DISTINCT
			tasklists.id 'id', 
			lam.id 'id_lampiran', 
			lam.tasklists_id 'tasklists_id', 
			IFNULL(fl.filename, 'NULL') 'filename',
			fl.path 'path', 
			fl.extension 'ext',
			fl.size 'size', 
			activity.name 'activity', 
			product.product, 
			tasklists.task_type, 
			tasklists.task_type_name 'jenis_task', 
			IFNULL(tasklists.start_date, 'NULL') 'start_date',
			IFNULL(tasklists.end_date, 'NULL') 'end_date',
			tasklists.risk_issue 'risk_issue', 
			tasklists.risk_indicator 'risk_indicator', 
			tasklists.status, tasklists.approval, 
			tasklists.validation, 
			tasklists.approval_status 'status_approval', 
			tasklists.maker_id 'maker_id', 
			tasklists.period 'period'`).
		Joins(`LEFT JOIN tasklists_uker ON tasklists_uker.tasklist_id = tasklists.id`).
		Joins(`left join activity on activity.id = tasklists.activity_id left join product on product.id = tasklists.product_id`).
		Joins(`left join tasklists_lampiran lam on lam.tasklists_id = tasklists.id `).
		Joins(`left join files fl ON fl.id = lam.files_id`)

	queryCount = queryCount.Select(`
			COUNT(DISTINCT tasklists.id) AS 'total'
		`).
		Joins(`LEFT JOIN tasklists_uker ON tasklists_uker.tasklist_id = tasklists.id`)

	if request.PERNR != "" {
		query.Where("tasklists.validation = ? OR tasklists.approval = ?", request.PERNR, request.PERNR)
		queryCount.Where("tasklists.validation = ? OR tasklists.approval = ?", request.PERNR, request.PERNR)
	}

	if request.REGION != "all" {
		query.Where("tasklists_uker.REGION = ?", request.REGION)
		queryCount.Where("tasklists_uker.REGION = ?", request.REGION)
	}

	if request.MAINBR != "all" {
		query.Where("tasklists_uker.MAINBR = ?", request.MAINBR)
		queryCount.Where("tasklists_uker.MAINBR = ?", request.MAINBR)
	}

	if request.BRANCH != "all" {
		branches := strings.Split(request.BRANCH, ",")
		query.Where("tasklists_uker.BRANCH in (?)", branches)
		queryCount.Where("tasklists_uker.BRANCH in (?)", branches)
	}

	if request.ActivityID != 0 {
		query.Where("tasklists.activity_id = ?", request.ActivityID)
		queryCount.Where("tasklists.activity_id = ?", request.ActivityID)
	}

	if request.ProductID != 0 {
		query.Where("tasklists.product_id = ?", request.ProductID)
		queryCount.Where("tasklists.product_id = ?", request.ProductID)
	}

	if request.RiskIssueID != 0 {
		query.Where("tasklists.risk_issue_id = ?", request.RiskIssueID)
		queryCount.Where("tasklists.risk_issue_id = ?", request.RiskIssueID)
	}

	if request.RiskIndicatorID != 0 {
		query.Where("tasklists.risk_indicator_id  = ?", request.RiskIndicatorID)
		queryCount.Where("tasklists.risk_indicator_id  = ?", request.RiskIndicatorID)
	} else {
		if request.RiskIndicator != "" {
			query.Where("tasklists.risk_indicator  = ?", request.RiskIndicator)
			queryCount.Where("tasklists.risk_indicator  = ?", request.RiskIndicator)
		}
	}

	if request.JenisTask != "" {
		query.Where("tasklists.task_type = ?", request.JenisTask)
		queryCount.Where("tasklists.task_type = ?", request.JenisTask)
	}

	if request.Status != "" {
		query.Where("tasklists.status = ?", request.Status)
		queryCount.Where("tasklists.status = ?", request.Status)
	} else {
		query.Where("tasklists.status = 'Aktif'")
		queryCount.Where("tasklists.status = 'Aktif'")
	}

	if request.Approval != "Semua" && request.Approval != "" {
		if request.Approval == "Ditolak" {
			query.Where("tasklists.approval_status = 'Ditolak oleh Validator' OR tasklists.approval_status = 'Ditolak oleh Approver'")
			queryCount.Where("tasklists.approval_status = 'Ditolak oleh Validator' OR tasklists.approval_status = 'Ditolak oleh Approver'")
		} else {
			query.Where("tasklists.approval_status = ?", request.Approval)
			queryCount.Where("tasklists.approval_status = ?", request.Approval)
		}
	}

	var count int64
	queryCount.Find(&count)

	totalData = int(count)
	query.Order("tasklists.created_at desc")
	query.Limit(request.Limit).Offset(request.Offset).Find(&responses)

	totalPages := int(math.Ceil(float64(totalData) / float64(request.Limit)))

	return responses, totalPages, totalData, err
}

func (r TasklistsRepository) GetUkerList(request models.UkerListReq) (responses []models.TasklistsUkerData, err error) {
	// if request.TipeUker == "KC" || request.TipeUker == "KW" {
	if request.HILFM == "033" || request.HILFM == "034" || request.HILFM == "228" {
		rows, err := r.db.DB.Raw(`
			SELECT tu.REGION, tu.RGDESC, tu.MAINBR, tu.MBDESC, tu.BRANCH, tu.BRDESC, 'kelolaan' 
			FROM tasklists_uker tu WHERE tu.tasklist_id = ? AND tu.BRANCH IN (SELECT BRANCH FROM uker_kelolaan_user WHERE pn = ?)`, request.ID, request.PERNR).Rows()

		defer rows.Close()
		var uker models.TasklistsUkerData

		for rows.Next() {
			r.db.DB.ScanRows(rows, &uker)
			responses = append(responses, uker)
		}

		return responses, err
	} else if strings.Contains(request.StellTX, "RISK MANAGEMENT & COMPLIANCE") {
		rows, err := r.db.DB.Raw(`
			SELECT tu.REGION, tu.RGDESC, tu.MAINBR, tu.MBDESC, tu.BRANCH, tu.BRDESC, 
			CASE WHEN tu.BRANCH IN (SELECT BRANCH FROM uker_kelolaan_user WHERE pn = ?) THEN 'kelolaan' ELSE 'bukan' END AS 'kelolaan' 
			FROM tasklists_uker tu WHERE tu.tasklist_id = ? AND tu.REGION = ?`, request.PERNR, request.ID, request.REGION).Rows()

		defer rows.Close()
		var uker models.TasklistsUkerData

		for rows.Next() {
			r.db.DB.ScanRows(rows, &uker)
			responses = append(responses, uker)
		}

		return responses, err
	} else {
		rows, err := r.db.DB.Raw(`
			SELECT tu.REGION, tu.RGDESC, tu.MAINBR, tu.MBDESC, tu.BRANCH, tu.BRDESC
			FROM tasklists_uker tu WHERE tu.tasklist_id = ?`, request.ID).Rows()

		defer rows.Close()
		var uker models.TasklistsUkerData

		for rows.Next() {
			r.db.DB.ScanRows(rows, &uker)
			responses = append(responses, uker)
		}

		return responses, err
	}
}

func (r TasklistsRepository) FilterByID(request models.TasklistsFilterRequest) (responses []models.TasklistsFilterResponse, totalRows int, totalData int, err error) {
	rows, err := r.db.DB.Raw(`
		SELECT tasklists.id 'id', lam.id 'id_lampiran', lam.tasklists_id 'tasklists_id', fl.filename 'filename', fl.path 'path', fl.extension 'ext',
		fl.size 'size', tasklists_uker.REGION, tasklists_uker.RGDESC, tasklists_uker.RGNAME, tasklists_uker.MAINBR, tasklists_uker.MBDESC, tasklists_uker.MBNAME, tasklists_uker.BRANCH, tasklists_uker.BRDESC, tasklists_uker.BRNAME, tasklists_uker.unit_kerja, activity.name 'activity', product.product, tasklists.task_type, tasklists.task_type_name 'jenis_task', tasklists.period 'period', tasklists.start_date, tasklists.end_date, tasklists.risk_issue 'risk_issue', tasklists.risk_indicator 'risk_indicator', tasklists.status, tasklists.approval, tasklists.validation, tasklists.approval_status 'status_approval', tasklists.maker_id 'maker_id' FROM tasklists 
		LEFT JOIN tasklists_uker ON tasklists_uker.tasklist_id = tasklists.id 
		left join activity on activity.id = tasklists.activity_id left join product on product.id = tasklists.product_id 
		left join tasklists_lampiran lam on lam.tasklists_id = tasklists.id 
		left join files fl ON fl.id = lam.files_id
		WHERE tasklists.id = ? LIMIT ? OFFSET ?`,
		request.ID, 1, 0).Rows()

	defer rows.Close()
	defer rows.Scan()

	var tasklistsFilter models.TasklistsFilterResponse

	for rows.Next() {
		r.db.DB.ScanRows(rows, &tasklistsFilter)
		responses = append(responses, tasklistsFilter)
	}
	if err != nil {
		r.logger.Zap.Error(err)
		return responses, 0, 0, err
	}

	paginateQuery := fmt.Sprintf(`SELECT count(*) FROM tasklists 
	left join activity on activity.id = tasklists.activity_id left join product on product.id = tasklists.product_id 
	WHERE tasklists.id = %d`, request.ID)

	err = r.dbRaw.DB.QueryRow(paginateQuery).Scan(&totalRows)

	if err != nil {
		r.logger.Zap.Error(err)
		return responses, 0, 0, err
	}

	result := float64(totalRows) / float64(1)
	resultFinal := int(math.Ceil(result))

	return responses, resultFinal, totalRows, err
}

func (r TasklistsRepository) FilterOfficer(request models.TasklistsFilterOfficerRequest) (responses []models.TasklistsFilterResponses, totalRows int, totalData int, err error) {
	db := r.db.DB
	query := db.Table("tasklists")
	queryCount := db.Table("tasklists")

	// query = query.Select(`
	// 			tasklists.id 'id',
	// 			// tasklists_uker.REGION,
	// 			// tasklists_uker.RGDESC,
	// 			// tasklists_uker.RGNAME,
	// 			// tasklists_uker.MAINBR,
	// 			// tasklists_uker.MBDESC,
	// 			// tasklists_uker.MBNAME,
	// 			// tasklists_uker.BRANCH,
	// 			// tasklists_uker.BRDESC,
	// 			// tasklists_uker.BRNAME,
	// 			// tasklists_uker.unit_kerja,
	// 			// activity.name 'activity',
	// 			product.product,
	// 			tasklists.task_type,
	// 			tasklists.task_type_name 'jenis_task',
	// 			tasklists.start_date, tasklists.end_date,
	// 			tasklists.risk_issue 'risk_issue',
	// 			tasklists.risk_indicator 'risk_indicator',
	// 			tasklists.status, tasklists.approval,
	// 			tasklists.validation,
	// 			tasklists.approval_status 'status_approval',
	// 			tasklists.maker_id 'maker_id',
	// 			tasklists.period 'period'`).

	query = query.Select(`
				DISTINCT
				tasklists.id 'id', 
				product.product,
				activity.name 'activity',
				tasklists.task_type,
				tasklists.task_type_name 'jenis_task', 
				IFNULL(tasklists.start_date, 'NULL') AS 'start_date',
				IFNULL(tasklists.end_date, 'NULL') AS 'end_date',
				tasklists.risk_issue 'risk_issue', 
				tasklists.risk_indicator 'risk_indicator', 
				tasklists.status, tasklists.approval, 
				tasklists.validation, 
				tasklists.approval_status 'status_approval',
				tasklists.maker_id 'maker_id', 
				tasklists.period 'period'`).
		Joins(`LEFT JOIN tasklists_uker ON tasklists_uker.tasklist_id = tasklists.id`).
		Joins(`left join activity on activity.id = tasklists.activity_id`).
		Joins(`left join product on product.id = tasklists.product_id`).
		Where(`tasklists.status = 'Aktif'`)

	queryCount = queryCount.Select(`
			COUNT(DISTINCT tasklists.id) AS 'total'`).
		Joins(`LEFT JOIN tasklists_uker ON tasklists_uker.tasklist_id = tasklists.id`)

	if request.HILFM == "033" || request.HILFM == "034" || request.HILFM == "228" {
		query.Where("tasklists_uker.BRANCH in (select branch from uker_kelolaan_user where pn = ?)", request.MakerID)
		queryCount.Where("tasklists_uker.BRANCH in (select branch from uker_kelolaan_user where pn = ?)", request.MakerID)
	}

	if strings.Contains(request.StellTX, "RISK MANAGEMENT & COMPLIANCE") {
		query.Where("tasklists_uker.REGION = ?", request.REGION)
		queryCount.Where("tasklists_uker.REGION = ?", request.REGION)
	}

	if request.MAINBR != "all" {
		mainbrs := strings.Split(request.MAINBR, ",")

		query.Where("tasklists_uker.MAINBR in (?)", mainbrs)
		queryCount.Where("tasklists_uker.MAINBR in (?)", mainbrs)
	}

	if request.BRANCH != "all" {
		branches := strings.Split(request.BRANCH, ",")
		query.Where("tasklists_uker.BRANCH in (?)", branches)
		queryCount.Where("tasklists_uker.BRANCH in (?)", branches)
	}

	if request.ActivityID != 0 {
		query.Where("tasklists.activity_id = ?", request.ActivityID)
		queryCount.Where("tasklists.activity_id = ?", request.ActivityID)
	}

	if request.ProductID != 0 {
		query.Where("tasklists.product_id = ?", request.ProductID)
		queryCount.Where("tasklists.product_id = ?", request.ProductID)
	}

	if request.RiskIssueID != 0 {
		query.Where("tasklists.risk_issue_id = ?", request.RiskIssueID)
		queryCount.Where("tasklists.risk_issue_id = ?", request.RiskIssueID)
	}

	if request.RiskIndicatorID != 0 {
		query.Where("tasklists.risk_indicator_id  = ?", request.RiskIndicatorID)
		queryCount.Where("tasklists.risk_indicator_id  = ?", request.RiskIndicatorID)
	} else {
		if request.RiskIndicator != "" {
			query.Where("tasklists.risk_indicator  = ?", request.RiskIndicator)
			queryCount.Where("tasklists.risk_indicator  = ?", request.RiskIndicator)
		}
	}

	if request.JenisTask != "" {
		query.Where("tasklists.task_type = ?", request.JenisTask)
		queryCount.Where("tasklists.task_type = ?", request.JenisTask)
	}

	if request.Approval != "Semua" && request.Approval != "" {
		if request.Approval == "Ditolak" {
			query.Where("tasklists.approval_status = 'Ditolak oleh Validator' OR tasklists.approval_status = 'Ditolak oleh Approver'")
			queryCount.Where("tasklists.approval_status = 'Ditolak oleh Validator' OR tasklists.approval_status = 'Ditolak oleh Approver'")
		} else {
			query.Where("tasklists.approval_status = ?", request.Approval)
			queryCount.Where("tasklists.approval_status = ?", request.Approval)
		}
	}

	if request.Status != "" {
		query.Where("tasklists.status = ?", request.Status)
		queryCount.Where("tasklists.status = ?", request.Status)
	}

	var count int64
	queryCount.Find(&count)

	totalData = int(count)
	query.Limit(request.Limit).Offset(request.Offset).Find(&responses)

	totalPages := int(math.Ceil(float64(totalData) / float64(request.Limit)))

	return responses, totalPages, totalData, err
}

func (r TasklistsRepository) CheckAvailability(request models.TasklistsCheckRequest) (response models.TasklistsCheckResponse, err error) {
	err = r.db.DB.Raw(`
		SELECT COUNT(*) AS 'total' FROM tasklists JOIN tasklists_uker ON tasklists_uker.tasklist_id = tasklists.id WHERE task_type = ? AND tasklists_uker.region = ? AND tasklists_uker.rgdesc = ? AND tasklists_uker.rgname = ? AND tasklists_uker.mainbr = ? AND tasklists_uker.mbdesc = ? AND tasklists_uker.mbname = ? AND tasklists_uker.branch = ? AND tasklists_uker.brdesc = ? AND tasklists_uker.brname = ? AND tasklists_uker.unit_kerja = ? AND risk_issue_id = ?`,
		request.TaskType, request.REGION, request.RGDESC, request.RGNAME, request.MAINBR, request.MBDESC, request.MBNAME, request.BRANCH, request.BRDESC, request.BRNAME, request.UnitKerja, request.RiskIssue).Scan(&response).Error

	if err != nil {
		r.logger.Zap.Error(err)
		return response, err
	}
	return response, err
}

func (r TasklistsRepository) CekRiskIssueAvail(request *models.TasklistRiskIssueAvailReq) (response models.TasklistRiskIssueAvailRes, err error) {
	err = r.db.DB.Raw(`SELECT count(*) as total FROM tasklists t 
					JOIN tasklists_uker tu ON tu.tasklist_id = t.id 
					WHERE risk_issue_id = ? AND tu.BRANCH = ? AND activity_id = ? AND product_id = ? AND status = 'Aktif'`, request.RiskIssueID, request.Branch, request.ActivityID, request.ProductID).Scan(&response).Error

	if err != nil {
		r.logger.Zap.Error(err)
		return response, err
	}

	return response, err
}

func (r TasklistsRepository) Store(request *models.Tasklists, tx *gorm.DB) (responses *models.Tasklists, err error) {
	return request, tx.Save(&request).Error
}

func (r TasklistsRepository) LimitTask(id int64) (maxLimit int, err error) {
	err = r.db.DB.Raw(`select tasklist_max from admin_setting as2 where id = ?`,
		id).Scan(&maxLimit).Error

	if err != nil {
		r.logger.Zap.Error(err)
		return 0, err
	}

	return maxLimit, err
}

func (r TasklistsRepository) CheckMaxLimit(TaskID int64, BranchID string) (total int, err error) {
	query := `select count(*) from tasklists_uker tu 
	join tasklists t on t.id = tu.tasklist_id
	where t.task_type = ? and tu.BRANCH = ?`

	err = r.db.DB.Raw(query, TaskID, BranchID).Scan(&total).Error

	if err != nil {
		r.logger.Zap.Error(err)
		return 0, err
	}

	return total, err
}

func (r TasklistsRepository) StoreNotif(request *models.TasklistNotif, tx *gorm.DB) (responses *models.TasklistNotif, err error) {
	return request, tx.Create(&request).Error
}

func (r TasklistsRepository) StoreTasklistDaily(request *models.TasklistsToday, tx *gorm.DB) (responses *models.TasklistsToday, err error) {
	return request, tx.Create(&request).Error
}

func (r TasklistsRepository) Update(request *models.TasklistsUpdate, tx *gorm.DB) (responses *models.TasklistsUpdate, err error) {
	return request, tx.Save(&request).Error
}

func (r TasklistsRepository) UpdateEndDate(request *models.TasklistsUpdateEndDate, tx *gorm.DB) (responses *models.TasklistsUpdateEndDate, err error) {
	return request, tx.Save(&request).Error
}

func (r TasklistsRepository) GetOne(id int64) (response models.TasklistDataResponse, err error) {
	err = r.db.DB.Raw(`
		SELECT tasklists.*, admin_setting.period 'task_type_period', admin_setting.range 'range', admin_setting.upload 'upload', admin_setting.kegiatan 'kegiatan',
		activity.name 'activity'
		FROM tasklists 
		JOIN admin_setting ON admin_setting.id = tasklists.task_type 
		JOIN activity ON activity.id = tasklists.activity_id 
		WHERE tasklists.id = ?`, id).Find(&response).Error

	if err != nil {
		r.logger.Zap.Error(err)
		return response, err
	}

	return response, err
}

func (r TasklistsRepository) GetNotes(id int64) (response models.TasklistRejectedNote, err error) {
	err = r.db.DB.Raw(`
		SELECT notes
		FROM tasklists_rejected 
		WHERE tasklist_id = ? AND status = "Not Done" order by id desc LIMIT 1`, id).Find(&response).Error

	if err != nil {
		r.logger.Zap.Error(err)
		return response, err
	}

	return response, err
}

func (r TasklistsRepository) Delete(request *models.TasklistsUpdateDelete, tx *gorm.DB) (responses bool, err error) {
	return true, tx.Save(&request).Error
}

func (r TasklistsRepository) GetDataBRC(request *models.GetBRCRequest) (responses []models.GetBRCResponse, err error) {
	// branch, err := strconv.Atoi(request.Branch)

	// if err != nil {
	// 	r.logger.Zap.Error(err)
	// 	return responses, err
	// }

	// rows, err := r.db.DB.Raw(`
	// 	SELECT PERNR, BRANCH FROM pa0001_eof pe WHERE BRANCH = ? and HILFM in (033, 034)`, branch).Rows()
	// rows, err := r.db.DB.Raw(`
	// 	SELECT pn 'PERNR', sname, region, rgdesc, mainbr, mbdesc, branch, brdesc FROM uker_kelolaan_user uku WHERE branch IN (?)`, request.Uker).Rows()
	rows, err := r.db.DB.Raw(`
		SELECT pn 'PERNR', sname, region, rgdesc, mainbr, mbdesc, branch, brdesc FROM uker_kelolaan_user uku WHERE branch IN (` + request.Uker + `)`).Rows()

	defer rows.Close()
	if err != nil {
		r.logger.Zap.Error(err)
		return responses, err
	}

	defer rows.Scan()

	var brcList models.GetBRCResponse

	for rows.Next() {
		r.db.DB.ScanRows(rows, &brcList)
		responses = append(responses, brcList)
	}

	return responses, err
}

func (r TasklistsRepository) Approval(request *models.TasklistsAprroval, tx *gorm.DB) (responses *models.TasklistsAprroval, err error) {
	return request, tx.Save(&request).Error
}

func (r TasklistsRepository) Validation(request *models.TasklistsValidation, tx *gorm.DB) (responses *models.TasklistsValidation, err error) {
	return request, tx.Save(&request).Error
}

func (r TasklistsRepository) CountTask(request models.TasklistCountRequest) (responses models.TasklistsCountResponse, err error) {
	err = r.db.DB.Raw(`
		SELECT COUNT(*) AS 'total' FROM tasklists_today tt
		JOIN activity ON activity.id = tt.activity_id
		WHERE tt.PERNR = ?`, request.PERNR).Scan(&responses).Error

	if err != nil {
		r.logger.Zap.Error(err)
		return responses, err
	}
	return responses, err
}

func (r TasklistsRepository) CountTaskDone(request models.TasklistCountRequest) (responses models.TasklistsCountResponse, err error) {
	err = r.db.DB.Raw(`
		SELECT COUNT(*) AS 'total' FROM tasklists
		JOIN tasklists_uker ON tasklists_uker.tasklist_id = tasklists.id
		JOIN tasklists_done_history ON tasklists_done_history.tasklist_id = tasklists.id
		WHERE tasklists.approval_status = "Disetujui" AND tasklists.status = "Aktif" 
		AND tasklists_uker.branch = ? AND tasklists_done_history.pernr = ?`,
		request.BRANCH, request.PERNR).Scan(&responses).Error

	if err != nil {
		r.logger.Zap.Error(err)
		return responses, err
	}
	return responses, err
}

func (r TasklistsRepository) LeadAutocompleteVal(request models.LeadAutocompleteRequest) (responses []models.LeadAutocomplete, err error) {
	if request.Region == "V" {
		err = r.db.DB.Raw(`
		SELECT pe.PERNR, pe.SNAME, pe.ORGEH, pe.ORGEH_TX 
		FROM pa0001_eof pe 
		JOIN dwh_branch dwb ON dwb.branch = CAST(pe.BRANCH AS UNSIGNED) 
		WHERE pe.KOSTL = 'PS21014' AND pe.sname LIKE '%` + request.Keyword + `%'`).Find(&responses).Error

		// AND pe.STELL IN ('70512867', '70512324', '70522337', '70522570', '70522569')

		if err != nil {
			r.logger.Zap.Error(err)
			return responses, err
		}

		return responses, err
	} else if request.Region == "S" {
		err = r.db.DB.Raw(`
		SELECT pe.PERNR, pe.SNAME, pe.ORGEH, pe.ORGEH_TX 
		FROM pa0001_eof pe 
		JOIN dwh_branch dwb ON dwb.branch = CAST(pe.BRANCH AS UNSIGNED) 
		WHERE pe.STELL_TX = 'REGIONAL RISK MANAGEMENT HEAD' AND pe.TIPE_UKER = 'KCK' AND pe.sname LIKE '%` + request.Keyword + `%'`).Find(&responses).Error

		if err != nil {
			r.logger.Zap.Error(err)
			return responses, err
		}

		return responses, err
	} else {
		err = r.db.DB.Raw(`
		SELECT pe.PERNR, pe.SNAME, pe.ORGEH, pe.ORGEH_TX 
		FROM pa0001_eof pe 
		JOIN dwh_branch dwb ON dwb.branch = CAST(pe.BRANCH AS UNSIGNED) 
		WHERE dwb.region = ? AND (pe.STELL_TX LIKE '%RISK MANAGEMENT & COMPLIANCE%') AND pe.sname LIKE '%`+request.Keyword+`%'`, request.Region).Find(&responses).Error

		if err != nil {
			r.logger.Zap.Error(err)
			return responses, err
		}

		return responses, err
	}
}

func (r TasklistsRepository) LeadAutocompleteApr(request models.LeadAutocompleteRequest) (responses []models.LeadAutocomplete, err error) {
	if request.Region == "V" {
		err = r.db.DB.Raw(`
			SELECT pe.PERNR, pe.SNAME, pe.ORGEH, pe.ORGEH_TX 
			FROM pa0001_eof pe 
			JOIN dwh_branch dwb ON dwb.branch = CAST(pe.BRANCH AS UNSIGNED) 
			WHERE pe.KOSTL = 'PS21014' AND pe.sname LIKE '%` + request.Keyword + `%'`).Find(&responses).Error

		// AND pe.STELL IN ('70522337', '70522570', '70522569', '70358269', '70422782', '70491215', '70358274', '70522339')

		if err != nil {
			r.logger.Zap.Error(err)
			return responses, err
		}

		return responses, err
	} else if request.Region == "S" {
		err = r.db.DB.Raw(`
		SELECT pe.PERNR, pe.SNAME, pe.ORGEH, pe.ORGEH_TX 
		FROM pa0001_eof pe 
		JOIN dwh_branch dwb ON dwb.branch = CAST(pe.BRANCH AS UNSIGNED) 
		WHERE pe.STELL_TX = 'REGIONAL RISK MANAGEMENT HEAD' AND pe.TIPE_UKER = 'KCK' AND pe.sname LIKE '%` + request.Keyword + `%'`).Find(&responses).Error

		if err != nil {
			r.logger.Zap.Error(err)
			return responses, err
		}

		return responses, err
	} else {
		err = r.db.DB.Raw(`
			SELECT pe.PERNR, pe.SNAME, pe.ORGEH, pe.ORGEH_TX 
			FROM pa0001_eof pe 
			JOIN dwh_branch dwb ON dwb.branch = CAST(pe.BRANCH AS UNSIGNED) 
			WHERE dwb.region = ? AND (pe.STELL_TX LIKE '%RISK MANAGEMENT & COMPLIANCE%' OR pe.STELL_TX LIKE '%REGIONAL RISK MANAGEMENT HEAD%') AND pe.sname LIKE '%`+request.Keyword+`%'`, request.Region).Find(&responses).Error

		if err != nil {
			r.logger.Zap.Error(err)
			return responses, err
		}
		return responses, err
	}
}

func (r TasklistsRepository) LeadAutocompleteKanwil(request models.LeadAutocompleteRequest) (responses []models.LeadAutocomplete, err error) {
	// err = r.db.DB.Raw(`
	// 	SELECT pe.PERNR, pe.SNAME, pe.ORGEH, pe.ORGEH_TX
	// 	FROM pa0001_eof pe
	// 	WHERE pe.orgeh = ? AND pe.HILFM  in ('012','160', '161', '162','163','164', '014') AND pe.sname LIKE '%`+request.Keyword+`%'`, request.ORGEH).Find(&responses).Error
	err = r.db.DB.Raw(`
		SELECT pe.PERNR, pe.SNAME, pe.ORGEH, pe.ORGEH_TX 
		FROM pa0001_eof pe 
		JOIN dwh_branch dwb ON dwb.branch = CAST(pe.BRANCH AS UNSIGNED) 
		WHERE dwb.region = ? AND pe.STELL_TX LIKE '%RISK MANAGEMENT & COMPLIANCE%' AND pe.sname LIKE '%`+request.Keyword+`%'`, request.Region).Find(&responses).Error

	if err != nil {
		r.logger.Zap.Error(err)
		return responses, err
	}
	return responses, err
}

func (r TasklistsRepository) UserRegion(request models.UserRegionRequest) (responses models.UserRegionResponse, err error) {
	err = r.db.DB.Raw(`
		SELECT region FROM dwh_branch 
		WHERE dwh_branch.branch = ?`, request.Branch).Find(&responses).Error

	if err != nil {
		r.logger.Zap.Error(err)
		return responses, err
	}
	return responses, err
}

func (r TasklistsRepository) GetDataVerifikasi(request models.DataVerifikasiRequest) (responses []models.DataVerifikasiResponse, err error) {
	err = r.db.DB.Raw(`
		SELECT v.branch 'branch', v.brdesc 'brdesc', v.mbdesc 'mbdesc', v.rgdesc 'rgdesc', v.no_pelaporan 'no_pelaporan', v.maker_desc 'maker',
		v.risk_issue_id 'id_risk_event', v.risk_issue 'risk_event_name', v.hasil_verifikasi 'hasil_verifikasi', v.indikasi_fraud 'indikasi_fraud'
		FROM verifikasi v 
		JOIN tasklists t ON t.risk_issue_id = v.risk_issue_id AND t.product_id = v.product_id 
		JOIN activity a ON a.id = v.activity_id 
		JOIN sub_activity sa ON sa.id = v.sub_activity_id
		WHERE t.id = ?`, request.TasklistID).Find(&responses).Error

	if err != nil {
		r.logger.Zap.Error(err)
		return responses, err
	}
	return responses, err
}

func (r TasklistsRepository) CreateTableLampiran(request *models.TasklistHeaderStore) (err error) {
	// arr := strings.Split(request.HeaderLampiran[1:len(request.HeaderLampiran)-1], `","`)
	arr := strings.Split(request.HeaderLampiran, `,`)

	headerColumnMerge := ""

	// Menampilkan setiap elemen dalam array
	for i, value := range arr {
		if i != len(arr)-1 {
			headerColumnMerge += value + " TEXT, "
		} else {
			headerColumnMerge += value + " TEXT"
		}
	}

	namaTable := "lampiran_rap_" + request.RiskIssueID + "_" + request.RiskIndicatorID
	// taskIDStr := strconv.Itoa(int(request.TasklistID))
	// namaTable := "lampiran_rap_" + taskIDStr

	createTableSQL := `CREATE TABLE IF NOT EXISTS ` + namaTable + ` (id INT AUTO_INCREMENT PRIMARY KEY, tasklist_id INT, status TEXT, ` + headerColumnMerge + `) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;`
	err = r.db.DB.Exec(createTableSQL).Error

	return err
}

func (r TasklistsRepository) InsertTableLampiran(request *models.TasklistColumnStore) (err error) {
	namaTable := "lampiran_rap_" + request.RiskIssueID + "_" + request.RiskIndicatorID
	// namaTable := "lampiran_rap_" + request.TasklistID

	insertTableSQL := `INSERT INTO ` + namaTable + ` (tasklist_id, status, ` + request.HeaderLampiran + `) VALUES ` + request.IsiLampiran + `;`
	// insertTableSQL := `INSERT INTO ` + namaTable + ` (` + request.HeaderLampiran + `) VALUES ` + request.IsiLampiran + `;`

	err = r.db.DB.Exec(insertTableSQL).Error

	return err
}

func (r TasklistsRepository) InsertTableLampiranIndikator(request *models.LampiranIndikatorStore) (err error) {
	insertTableSQL := `INSERT INTO lampiran_indikator (risk_issue_id, risk_indicator_id, nama_table, jumlah_kolom, risk_indicator_desc) VALUES (?, ?, '` + request.NamaTable + `', ?, ?)`

	err = r.db.DB.Exec(insertTableSQL, request.RiskIssueID, request.RiskIndicatorID, request.JumlahKolom, request.RiskIndicatorDesc).Error

	return err
}

func (r TasklistsRepository) GetLampiranIndicator(request *models.LampiranIndikatorCheck) (response models.LampiranIndikatorResponse, err error) {
	err = r.db.DB.Raw(`
		SELECT * FROM lampiran_indikator WHERE risk_issue_id = ? AND risk_indicator_id = ?`, request.RiskIssueID, request.RiskIndicatorID).Find(&response).Error

	if err != nil {
		r.logger.Zap.Error(err)
		return response, err
	}

	return response, err
}

func (r TasklistsRepository) DownloadLampiranIndikatorTemplate(request *models.LampiranIndikatorCheck) (response models.AnomaliHeaderResponse, err error) {
	tableName := "lampiran_rap_" + request.RiskIssueID + "_" + request.RiskIndicatorID
	// tableName := "lampiran_rap_" + request.TasklistID
	rows, err := r.dbRaw.DB.Query(`SELECT DISTINCT COLUMN_NAME FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = "` + tableName + `" AND COLUMN_NAME NOT IN ('id', 'tasklist_id', 'status') ORDER BY ORDINAL_POSITION`)

	if err != nil {
		r.logger.Zap.Error(err)
		return response, err
	}
	defer rows.Close()

	var ColumnName string

	for rows.Next() {
		err := rows.Scan(&ColumnName)
		if err != nil {
			r.logger.Zap.Error(err)
			return response, err
		}
		response.Header = append(response.Header, ColumnName)
	}

	err = rows.Err()
	if err != nil {
		r.logger.Zap.Error(err)
		return response, err
	}

	return response, err
}

func (r TasklistsRepository) ShowUker(id int64) (responses []models.TasklistsUker, err error) {
	err = r.db.DB.Raw(`SELECT * FROM tasklists_uker WHERE tasklist_id = ?`, id).Find(&responses).Error

	if err != nil {
		r.logger.Zap.Error(err)
		return responses, err
	}
	return responses, err
}

func (r TasklistsRepository) InsertTasklistRejected(tasklistRejected *models.TasklistsRejected, tx *gorm.DB) (err error) {
	return tx.Create(&tasklistRejected).Error
}

func (r TasklistsRepository) GetAnomaliHeader(request *models.AnomaliHeader) (response models.AnomaliHeaderResponse, err error) {
	tableName := "lampiran_rap_" + request.RiskIssue + "_" + request.RiskIndicator
	// tableName := "lampiran_rap_" + request.TasklistID
	rows, err := r.dbRaw.DB.Query(`SELECT COLUMN_NAME FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = "` + tableName + `" AND COLUMN_NAME NOT IN ('id', 'tasklist_id', 'status') ORDER BY ORDINAL_POSITION`)

	if err != nil {
		r.logger.Zap.Error(err)
		return response, err
	}
	defer rows.Close()

	var ColumnName string

	for rows.Next() {
		err := rows.Scan(&ColumnName)
		if err != nil {
			r.logger.Zap.Error(err)
			return response, err
		}
		response.Header = append(response.Header, ColumnName)
	}

	err = rows.Err()
	if err != nil {
		r.logger.Zap.Error(err)
		return response, err
	}

	return response, err
}

func (r TasklistsRepository) GetAnomaliValue(request *models.AnomaliValue) (response []models.AnomaliValueResponse, err error) {
	tableName := "lampiran_rap_" + request.RiskIssue + "_" + request.RiskIndicator
	// tableName := "lampiran_rap_" + request.TasklistID

	columns, err := getTableColumns(r.db.DB, tableName)
	if err != nil {
		panic(err)
	}

	var results []models.AnomaliValueResponse

	query := "CONCAT("
	for i, column := range columns {
		if i > 0 {
			query += ", '\\\\', `"
		} else {
			query += "`"
		}
		query += column + "`"
	}
	query += ") as value"

	r.db.DB.Table(tableName).Select(query).Where("tasklist_id = ?", request.TasklistID).Scan(&results)
	// r.db.DB.Table(tableName).Select(query).Scan(&results)

	// for _, result := range results {
	// 	fmt.Printf("Value: %s\n", result.Value)
	// }

	return results, err
}

func getTableColumns(db *gorm.DB, tableName string) ([]string, error) {
	var columns []string

	result := db.Raw(`SELECT COLUMN_NAME FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = "` + tableName + `" AND COLUMN_NAME NOT IN ('id', 'tasklist_id', 'status') ORDER BY ORDINAL_POSITION`).Scan(&columns)
	if result.Error != nil {
		return nil, result.Error
	}

	return columns, nil
}

func (r TasklistsRepository) GetFirstLampiran(request *models.GetFirstLampiranRequest) (response models.GetFirstLampiranResponse, err error) {
	err = r.db.DB.Raw(`
			select f.filename, f.path from tasklists t 
			join tasklists_lampiran tl on tl.tasklists_id = t.id 
			join files f on f.id = tl.files_id
			where t.risk_issue_id = ? and t.risk_indicator_id = ?
			order by t.created_at asc limit 1`, request.RiskIssue, request.RiskIndicator).Find(&response).Error

	if err != nil {
		r.logger.Zap.Error(err)
		return response, err
	}

	return response, err
}

func (r TasklistsRepository) DeleteIsiLampiran(request *models.TasklistLampiranDelete, tx *gorm.DB) (err error) {
	riskIssueString := strconv.Itoa(int(request.RiskIssue))
	riskIndicatorString := strconv.Itoa(int(request.RiskIndicator))
	taskIDString := strconv.Itoa(int(request.TasklistID))

	tableName := "lampiran_rap_" + riskIssueString + "_" + riskIndicatorString

	deleteTableSQL := `DELETE FROM ` + tableName + ` WHERE tasklist_id = ` + taskIDString + `;`

	err = r.db.DB.Exec(deleteTableSQL).Error

	return err
}

func (r TasklistsRepository) GetTipeUkerMaker(pnRequest string) (response models.PegawaiData, err error) {
	err = r.db.DB.Raw(`SELECT TIPE_UKER 'tipe_uker',  CASE
						WHEN hilfm LIKE '%033%' OR hilfm LIKE '%034%' OR hilfm LIKE '%228%' THEN "['033', '034']"
						ELSE hilfm END AS 'hilfm', 
						SUBSTRING(STELL_TX, LOCATE('RISK MANAGEMENT & COMPLIANCE', stell_tx), 
						CHAR_LENGTH('RISK MANAGEMENT & COMPLIANCE')) 'stell_tx' FROM pa0001_eof WHERE PERNR = ?`, pnRequest).Find(&response).Error

	if err != nil {
		r.logger.Zap.Error(err)
		return response, err
	}
	return response, err
}
