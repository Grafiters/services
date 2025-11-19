package coaching

import (
	"fmt"
	"math"
	"riskmanagement/lib"
	models "riskmanagement/models/coaching"
	"strconv"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"

	// "encoding/json"
	"database/sql"

	"strings"
)

type CoachingDefinition interface {
	WithTrx(trxHandle *gorm.DB) CoachingRepository
	GetAll() (responses []models.CoachingResponse, err error)
	GetOne(id int64) (responses models.CoachingResponse, err error)
	Store(request *models.Coaching, tx *gorm.DB) (responses *models.Coaching, err error)
	Delete(request *models.CoachingUpdateDelete, include []string, tx *gorm.DB) (responses bool, err error)
	DeleteCoachingActivity(id int64, tx *gorm.DB) (err error)
	UpdateAllCoaching(request *models.CoachingUpdateActivity, include []string, tx *gorm.DB) (responses bool, err error)
	FilterCoaching(request *models.CoachingFilterRequest) (responses []models.CoachingResponse, totalRows int, totalData int, err error)
	GetNoPelaporan(request *models.NoPalaporanRequest) (responses []models.NoPelaporanNullResponse, err error)
	GetData() (responses []models.CoachingResponse, err error)
	GetDataWithPagination(request *models.CoachingPagination) (responses []models.CoachingResponse, totalRows int, totalData int, err error)
	CoachingReportFilter(request *models.CoachingFilterReportRequest) (responses []models.CoachingFilterReportResponse, totalAktivitas int64, totalRows int64, err error)
	CoachingFinalReportFilter(request *models.CoachingFilterReportRequest) (responses []models.CoachingFilterReportFinalResponse, totalRows int64, err error)
	CoachingReportDetail(request *models.CoachingReportDetailRequest) (responses models.CoachingReportDetailResponse, err error)
	CoachingReportMateriList(request *models.CoachingReportMateriRequest) (responses []models.CoachingDetailMateriResponseNull, err error)
	CoachingReportFilterByUkerAllActivity(request *models.CoachingFilterReportRequest) (responses []models.CoachingFilterByUkerAllActivityReportResponse, totalRows int64, totalData int64, err error)
	CoachingReportByUkerFilter(request *models.CoachingFilterReportByUker) (responses []models.CoachingFilterReportByUkerResponse, totalRows int64, err error)
	CoachingReportFilterByUkerComplete(request *models.CoachingFilterReportByUker) (responses []models.CoachingFilterReportFinalResponseNull, totalRows int, err error)
	CoachingReportList(request *models.CoachingReportListRequest) (responses []models.CoachingReportListResponse, totalRows int, err error)
	CoachingFrekuensiRpt(request *models.FrekuensiCoachingRequest) (responses []models.FrekuensiCoachingResponse, totalRows int, err error)

	// Versioning 23/10/2023
	GetJudulMateri(id int64) (responses []models.JudulMateriCoaching, err error)
}

type CoachingRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewCoachingRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) CoachingDefinition {
	return CoachingRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// Delete implements CoachingDefinition
func (coaching CoachingRepository) Delete(request *models.CoachingUpdateDelete, include []string, tx *gorm.DB) (responses bool, err error) {
	return true, tx.Save(&request).Error
}

// DeleteCoachingActivity implements CoachingDefinition
func (coaching CoachingRepository) DeleteCoachingActivity(id int64, tx *gorm.DB) (err error) {
	return tx.Where("id = ?", id).Delete(&models.CoachingActivityRequest{}).Error
}

// GetAll implements CoachingDefinition
func (coaching CoachingRepository) GetAll() (responses []models.CoachingResponse, err error) {
	return responses, coaching.db.DB.Find(&responses).Error
}

// GetOne implements CoachingDefinition
func (coaching CoachingRepository) GetOne(id int64) (responses models.CoachingResponse, err error) {
	err = coaching.db.DB.Raw(`
	SELECT
		coa.id,
		coa.no_pelaporan,
		coa.REGION,
		coa.RGDESC,
		coa.MAINBR,
		coa.MBDESC,
		coa.BRANCH,
		coa.BRDESC,
		coa.unit_kerja,
		coa.jenis_peserta,
		coa.jabatan_peserta,
		coa.jumlah_peserta,
		coa.list_peserta,
		coa.activity_id,
		coa.sub_activity_id,
		coa.product_id,
		coa.maker_id,
		coa.maker_desc,
		coa.maker_date,
		coa.last_maker_id,
		coa.last_maker_desc,
		coa.last_maker_date,
		CASE
				WHEN coa.status = "01a" AND coa.action = "Selesai" THEN "Selesai"
				WHEN coa.status = "01a" AND coa.action = "Draft" THEN "Draft"
				WHEN coa.status = "02b" AND (coa.action = "Update" OR coa.action ="Selesai")   THEN "Selesai"
				ELSE "Delete"
		END 'status',
		coa.deleted,
		coa.created_at,
		coa.updated_at
	FROM coaching coa
	WHERE coa.id = ?`, id).Find(&responses).Error

	if err != nil {
		coaching.logger.Zap.Error(err)
		return responses, err
	}
	return responses, err
}

// Store implements CoachingDefinition
func (coaching CoachingRepository) Store(request *models.Coaching, tx *gorm.DB) (responses *models.Coaching, err error) {
	return request, tx.Save(&request).Error
}

// UpdateAllCoaching implements CoachingDefinition
func (coaching CoachingRepository) UpdateAllCoaching(request *models.CoachingUpdateActivity, include []string, tx *gorm.DB) (responses bool, err error) {
	return true, tx.Save(&request).Error
}

// WithTrx implements CoachingDefinition
func (coaching CoachingRepository) WithTrx(trxHandle *gorm.DB) CoachingRepository {
	if trxHandle == nil {
		coaching.logger.Zap.Error("transaction Database not found in gin context")
		return coaching
	}
	coaching.db.DB = trxHandle
	return coaching
}

// FilterCoaching implements CoachingDefinition
func (coaching CoachingRepository) FilterCoaching(request *models.CoachingFilterRequest) (responses []models.CoachingResponse, totalRows int, totalData int, err error) {
	branches := strings.Split(request.Branches, ",")
	db := coaching.db.DB

	queryBuilder := db.Model(&responses).
		Select(`
			coaching.id 'id',
			coaching.no_pelaporan 'no_pelaporan',
			act.name 'aktifitas',
			coaching.BRDESC 'unit_kerja',
			CASE
				WHEN coaching.status = "01a" AND coaching.action = "Selesai" THEN "Selesai"
				WHEN coaching.status = "01a" AND coaching.action = "Draft" THEN "Draft"
				WHEN coaching.status = "02b" AND (coaching.action = "Update" OR coaching.action ="Selesai")   THEN "Selesai"
				ELSE "Delete"
			END 'status'
		`).
		Joins("JOIN coaching_activity coAct ON  coAct.coaching_id = coaching.id").
		Joins("JOIN activity act ON act.kode_activity = coaching.activity_id").
		Where("coaching.deleted != 1").
		// Where(`coach.BRANCH in (?)`, branches).
		Order(("coaching.id DESC"))

	// if request.Branches != "" {
	// 	branches := strings.Split(request.Branches, ",")
	// 	db = db.Where(`coach.BRANCH in (?)`, branches)
	// }

	if request.Kostl != "" {
		// db = db.Where(`coaching.maker_id = ?`, request.Pernr)
		queryBuilder = queryBuilder.Where(`coaching.maker_id = ?`, request.Pernr)
	} else {
		queryBuilder = queryBuilder.Where(`coaching.BRANCH in (?)`, branches)
	}

	if request.NoPelaporan != "" {
		queryBuilder = queryBuilder.Where("coaching.no_pelaporan = ?", request.NoPelaporan)
	}

	if request.UnitKerja != "" {
		// queryBuilder = queryBuilder.Where("coaching.BRDESC like ?", fmt.Sprintf("%%%s%%", request.UnitKerja))
		queryBuilder = queryBuilder.Where("coaching.BRANCH = ?", request.UnitKerja)
	}

	if request.ActivityID != "" {
		queryBuilder = queryBuilder.Where("coaching.activity_id = ?", request.ActivityID)
	}

	if request.RiskIssueID != "" {
		queryBuilder = queryBuilder.Where("coAct.risk_issue_id = ?", request.RiskIssueID)
	}

	if request.JudulMateri != "" {
		queryBuilder = queryBuilder.Where("coAct.judul_materi like ?", fmt.Sprintf("%%%s%%", request.JudulMateri))
	}

	if request.Status != "" && request.Status != "Semua" && request.Status != "Selesai" {
		queryBuilder = queryBuilder.Where("coaching.action = ?", request.Status)
	}

	if request.Status == "Selesai" {
		queryBuilder = queryBuilder.Where("coaching.action IN(?, ?)", "Update", "Selesai")
	}

	if request.TglAwal != "" && request.TglAkhir != "" {
		queryBuilder = queryBuilder.Where("coaching.created_at >= ? AND coaching.created_at <= ?", request.TglAwal, lib.FixEndDate(request.TglAkhir))
	}

	var count int64
	queryBuilder.
		Group("coAct.coaching_id").
		Count(&count)

	totalRows = int(count)

	queryBuilder.
		Limit(request.Limit).
		Offset(request.Offset).
		Find(&responses)

	// calculate the total pages
	totalPages := int(math.Ceil(float64(totalRows) / float64(request.Limit)))
	return responses, totalPages, totalRows, err

}

// GetNoPelaporan implements CoachingDefinition
func (coaching CoachingRepository) GetNoPelaporan(request *models.NoPalaporanRequest) (responses []models.NoPelaporanNullResponse, err error) {
	kode := "CO-"
	today := lib.GetTimeNow("date2")

	if request.ORGEH != "" {
		kode += request.ORGEH + "-" + today
	}

	query := `SELECT RIGHT(CONCAT("0000",(count(*) + 1)), 4) 'no_pelaporan' FROM coaching WHERE no_pelaporan like ?`

	coaching.logger.Zap.Info(query)
	rows, err := coaching.dbRaw.DB.Query(query, fmt.Sprintf("%%%s%%", kode))
	if err != nil {
		coaching.logger.Zap.Error(err)
		return responses, err
	}
	defer rows.Close()

	coaching.logger.Zap.Info("rows ", rows)

	for rows.Next() {
		response := models.NoPelaporanNullResponse{}
		if err := rows.Scan(&response.NoPelaporan); err != nil {
			coaching.logger.Zap.Error(err)
			return responses, err
		}
		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		coaching.logger.Zap.Error(err)
		return responses, err
	}

	return responses, nil
}

// GetData implements CoachingDefinition
func (coaching CoachingRepository) GetData() (responses []models.CoachingResponse, err error) {
	rows, err := coaching.db.DB.Raw(`
		SELECT
			DISTINCT 
			coach.id 'id',
			coach.no_pelaporan 'no_pelaporan',
			act.name 'aktifitas',
			coach.BRDESC 'unit_kerja',
			coAct.risk_issue 'risk_issue',
			CASE
				WHEN coach.status = "01a" AND coach.action = "Draft" THEN "Draft"
				WHEN coach.status = "02b" AND (coach.action = "Update" OR coach.action ="Selesai")   THEN "Selesai"
				ELSE "Delete"
			END 'status'
		FROM coaching coach
		LEFT JOIN coaching_activity coAct ON coach.id = coAct.coaching_id
		LEFT JOIN activity act ON coach.activity_id = act.id
		WHERE coach.deleted != 1`).Rows()

	defer rows.Close()

	var Coach models.CoachingResponse
	for rows.Next() {
		coaching.db.DB.ScanRows(rows, &Coach)
		responses = append(responses, Coach)
	}

	return responses, err
}

// GetDataWithPagination implements CoachingDefinition
func (coa CoachingRepository) GetDataWithPagination(request *models.CoachingPagination) (responses []models.CoachingResponse, totalRows int, totalData int, err error) {
	branches := strings.Split(request.Branches, ",")
	db := coa.db.DB.Table("coaching coach")

	db.Select(`
		coach.id 'id',
		coach.no_pelaporan 'no_pelaporan',
		act.name 'aktifitas',
		coach.BRDESC 'unit_kerja',
		CASE
			WHEN coach.status = "01a" AND coach.action = "Selesai" THEN "Selesai"
			WHEN coach.status = "01a" AND coach.action = "Draft" THEN "Draft"
			WHEN coach.status = "02b" AND (coach.action = "Update" OR coach.action ="Selesai")   THEN "Selesai"
			ELSE "Delete"
		END 'status'`).
		Joins(`LEFT JOIN activity act ON coach.activity_id = act.id`).
		Where(`coach.deleted != 1`).
		// Where(`coach.BRANCH in (?)`, branches).
		Order(`coach.id DESC`)

	// if request.Branches != "" {

	// 	db = db.Where(`coach.BRANCH in (?)`, branches)
	// }

	if request.Kostl != "" {
		db = db.Where(`coach.maker_id = ?`, request.Pernr)
	} else {
		db = db.Where(`coach.BRANCH in (?)`, branches)
	}

	var count int64
	db.
		Group(`coach.id`).
		Count(&count)

	totalData = int(count)

	err = db.Limit(request.Limit).Offset(request.Offset).Find(&responses).Error
	if err != nil {
		return responses, totalRows, totalData, err
	}

	if totalData > 0 {
		totalRows = int(math.Ceil(float64(totalData) / float64(request.Limit)))
	}

	return responses, totalRows, totalData, err
}

func (repo CoachingRepository) CoachingReportFilter(request *models.CoachingFilterReportRequest) (responses []models.CoachingFilterReportResponse, totalAktivitas int64, totalRows int64, err error) {
	db := repo.db.DB
	query := db
	queryPagination := db
	sortQuery := ""
	var errPagination error

	//bank wide
	filter1 := request.ReportType == "aktivitas" &&
		request.Activity == "all" &&
		request.Product == "all" &&
		(request.Title == "all" || request.Title == "Semua")

	filter2 := request.ReportType == "aktivitas" &&
		request.Activity != "all" &&
		request.Product == "all" &&
		(request.Title == "all" || request.Title == "Semua")

	filter3 := request.ReportType == "aktivitas" &&
		request.Activity != "all" &&
		request.Product != "all" &&
		(request.Title == "all" || request.Title == "Semua")

	if request.Sort == "DESC" {
		sortQuery = `total DESC, id`
	} else {
		sortQuery = `total ASC, id`
	}

	if filter1 {
		fmt.Println("====== query 1")

		subquery := db.Table("coaching").
			Select("COUNT(*) as total")

		totalAktvityQuery := db.Table("coaching").Select("COUNT(*) totalAktivitas")

		if request.REGION != "all" {
			subquery = subquery.Where("coaching.REGION = ?", request.REGION)
			totalAktvityQuery = totalAktvityQuery.Where("coaching.REGION = ?", request.REGION)
		}

		if request.REGION != "all" && request.MAINBR != "all" {
			mainbrs := strings.Split(request.MAINBR, ",")

			subquery = subquery.Where("coaching.MAINBR in (?)", mainbrs)
			totalAktvityQuery = totalAktvityQuery.Where("coaching.MAINBR in (?)", mainbrs)
		}

		if request.REGION != "all" && request.MAINBR != "all" && request.BRANCH != "all" {
			branches := strings.Split(request.BRANCH, ",")
			subquery = subquery.Where("coaching.BRANCH in (?)", branches)
			totalAktvityQuery = totalAktvityQuery.Where("coaching.BRANCH in (?)", branches)
		}

		subquery = subquery.
			Where("coaching.activity_id = activity.id").
			Where("coaching.created_at >= ? AND coaching.created_at <= ?", request.StartDate, lib.FixEndDate(request.EndDate))

		totalAktvityQuery = totalAktvityQuery.
			Where("coaching.created_at >= ? AND coaching.created_at <= ?", request.StartDate, lib.FixEndDate(request.EndDate)).
			Find(&totalAktivitas)

		// Define query
		query = db.Table("activity").
			Select("activity.id, activity.kode_activity as code, activity.name, (?) total", subquery)

		// query pagination
		// errPagination = db.Table("activity").Count(&totalRows).Error
		errPagination = query.Count(&totalRows).Error

		query = query.Limit(request.Limit).
			Offset(request.Offset).
			Order(sortQuery).
			Find(&responses)

	} else if filter2 {
		fmt.Println("====== query 2")

		// Define subquery
		subquery := db.Table("coaching").
			Select("COUNT(*) as total")

		totalAktvityQuery := db.Table("coaching").Select("COUNT(*) totalAktivitas")

		if request.REGION != "all" {
			subquery = subquery.Where("coaching.REGION = ?", request.REGION)
			totalAktvityQuery = totalAktvityQuery.Where("coaching.REGION = ?", request.REGION)
		}

		if request.REGION != "all" && request.MAINBR != "all" {
			mainbrs := strings.Split(request.MAINBR, ",")

			subquery = subquery.Where("coaching.MAINBR in (?)", mainbrs)
			totalAktvityQuery = totalAktvityQuery.Where("coaching.MAINBR in (?)", mainbrs)

		}

		if request.REGION != "all" && request.MAINBR != "all" && request.BRANCH != "all" {
			branches := strings.Split(request.BRANCH, ",")
			subquery = subquery.Where("coaching.BRANCH in (?)", branches)
			totalAktvityQuery = totalAktvityQuery.Where("coaching.BRANCH in (?)", branches)
		}

		subquery = subquery.
			Where("coaching.activity_id = ?", request.Activity).
			Where("coaching.product_id = product.id").
			Where("coaching.created_at >= ? AND coaching.created_at <= ?", request.StartDate, lib.FixEndDate(request.EndDate))

		totalAktvityQuery = totalAktvityQuery.
			Where("coaching.activity_id = ?", request.Activity).
			Where("coaching.created_at >= ? AND coaching.created_at <= ?", request.StartDate, lib.FixEndDate(request.EndDate)).
			Find(&totalAktivitas)

		// Define query
		query = db.Table("product").
			Select("product.id, product.kode_product as code, product.product as name, (?) total", subquery).
			Where("product.activity_id = ?", request.Activity)

		// query pagination
		// errPagination = db.Table("product").Where("activity_id = ?", request.Activity).Count(&totalRows).Error
		errPagination = query.Count(&totalRows).Error

		query = query.Limit(request.Limit).
			Offset(request.Offset).
			Order(sortQuery).
			Find(&responses)

	} else if filter3 {
		fmt.Println("====== query 3")

		// Define query
		query = db.Table("coaching").
			Select("coaching.activity_id  as id, coaching.product_id as code, ca.judul_materi as name, COUNT(*) as total").
			Joins("JOIN coaching_activity ca ON ca.coaching_id = coaching.id").
			Where("coaching.activity_id = ?", request.Activity).
			Where("coaching.product_id = ?", request.Product).
			Where("coaching.created_at >= ? AND coaching.created_at <= ?", request.StartDate, lib.FixEndDate(request.EndDate)).
			Group("ca.judul_materi").
			Count(&totalRows) //for pagination

		// query pagination
		// queryPagination = db.Table("coaching").
		// 	Select("COUNT(coaching.activity_id) as pagination").
		// 	Joins("JOIN coaching_activity ca ON ca.coaching_id = coaching.id").
		// 	Where("coaching.activity_id = ?", request.Activity).
		// 	Where("coaching.product_id = ?", request.Product).
		// 	Where("coaching.created_at >= ? AND coaching.created_at <= ?", request.StartDate, lib.FixEndDate(request.EndDate))
		// 	// Group("ca.judul_materi")

		// totalAktvityQuery := db.Table("coaching").
		// 	Select("COUNT(*) totalAktivitas").
		// 	Joins("JOIN coaching_activity ca ON ca.coaching_id = coaching.id").
		// 	Where("coaching.activity_id = ?", request.Activity).
		// 	Where("coaching.product_id = ?", request.Product).
		// 	Where("coaching.created_at >= ? AND coaching.created_at <= ?", request.StartDate, lib.FixEndDate(request.EndDate)).
		// 	Group("ca.judul_materi").
		// 	Scan(&totalAktivitas)

		totalAktvityQuery := db.Table("coaching").
			Select("COUNT(*) totalAktivitas").
			Where("coaching.activity_id = ?", request.Activity).
			Where("coaching.product_id = ?", request.Product).
			Where("coaching.created_at >= ? AND coaching.created_at <= ?", request.StartDate, lib.FixEndDate(request.EndDate)).
			Find(&totalAktivitas)

		if request.REGION != "all" {
			query = query.Where("coaching.REGION = ?", request.REGION)
			// queryPagination = queryPagination.Where("coaching.REGION = ?", request.REGION)
			totalAktvityQuery = totalAktvityQuery.Where("coaching.REGION = ?", request.REGION)
		}

		if request.REGION != "all" && request.MAINBR != "all" {
			mainbrs := strings.Split(request.MAINBR, ",")

			query = query.Where("coaching.MAINBR in (?)", mainbrs)
			// queryPagination = queryPagination.Where("coaching.MAINBR = ?", request.MAINBR)
			totalAktvityQuery = totalAktvityQuery.Where("coaching.MAINBR in (?)", mainbrs)
		}

		if request.REGION != "all" && request.MAINBR != "all" && request.BRANCH != "all" {
			branches := strings.Split(request.BRANCH, ",")
			query = query.Where("coaching.BRANCH  in (?)", branches)
			// queryPagination = queryPagination.Where("coaching.BRANCH in (?)", branches)
			totalAktvityQuery = totalAktvityQuery.Where("coaching.BRANCH in (?)", branches)
		}

		query = query.Limit(request.Limit).
			Offset(request.Offset).
			Order(sortQuery).
			Find(&responses)

		err = query.Error

		// errPagination = queryPagination.Scan(&totalRows).Error
	}

	//QUERY PAGINATION
	repo.logger.Zap.Info("coaching-queryPagination-activity-unknown", queryPagination)

	if errPagination != nil {
		return responses, totalAktivitas, totalRows, err
	}

	repo.logger.Zap.Info("coaching-query-activity-unknown", query)
	if err != nil {
		return responses, totalAktivitas, totalRows, err
	}

	return responses, totalAktivitas, totalRows, err
}

// reportType unit kerja
func (repo CoachingRepository) CoachingReportByUkerFilter(request *models.CoachingFilterReportByUker) (responses []models.CoachingFilterReportByUkerResponse, totalRows int64, err error) {
	db := repo.db.DB
	sortQuery := ""
	var errPagination error

	fmt.Println("Module CoachingReportByUkerFilter")

	query := db

	//bank wide
	filter1 := request.ReportType == "unitkerja" &&
		request.REGION == "all" &&
		request.MAINBR == "all" &&
		request.BRANCH == "all"

	filter2 := request.ReportType == "unitkerja" &&
		request.REGION != "all" &&
		request.MAINBR == "all" &&
		request.BRANCH == "all"

	filter3 := request.ReportType == "unitkerja" &&
		request.REGION != "all" &&
		request.MAINBR != "all" &&
		request.BRANCH == "all"

	filter4 := request.ReportType == "unitkerja" &&
		request.REGION != "all" &&
		request.MAINBR != "all" &&
		request.BRANCH != "all"

	mainbrs := strings.Split(request.MAINBR, ",")
	branches := strings.Split(request.BRANCH, ",")

	if filter1 {
		if request.Sort == "DESC" {
			sortQuery = `PERCENTCOACHING DESC, uker.REGION`
		} else {
			sortQuery = `PERCENTCOACHING ASC, uker.REGION`
		}

		query = db.Table("dwh_branch uker").
			Select(`
						uker.REGION, 
						uker.RGDESC, 
						uker.MAINBR, 
						uker.MBDESC, 
						uker.BRANCH, 
						uker.BRDESC,
						COALESCE(total_coaching.TOTAL, 0) AS TOTALCOACHING, 
						COALESCE(brc.TOTAL, 0) AS TOTALBRC,
						ROUND(COALESCE((total_coaching.TOTAL/brc.TOTAL)*100, 0), 0) AS PERCENTCOACHING`).
			Joins(
				`LEFT JOIN (SELECT REGION, COUNT(id) AS TOTAL FROM coaching coaching_a
						WHERE (created_at >= ? AND created_at <= ?) GROUP BY REGION) total_coaching
						ON total_coaching.REGION = uker.REGION
						LEFT JOIN (SELECT kelolaan.REGION, COUNT(kelolaan.REGION) AS TOTAL FROM
						(SELECT REGION, COUNT(pn) AS TOTAL FROM uker_kelolaan_user GROUP BY pn) kelolaan
						GROUP BY kelolaan.REGION) brc ON brc.REGION = uker.REGION`, request.StartDate, lib.FixEndDate(request.EndDate)).
			Where(
				"uker.BRUNIT = 'B' AND uker.MAINBR = uker.BRANCH AND (uker.BRDESC LIKE ? OR uker.BRDESC = ?)",
				"kanwil%", "Jkt KCK").
			Order(sortQuery).
			Limit(request.Limit).
			Offset(request.Offset).
			Find(&responses)

		err = query.Error

		errPagination = db.Table("dwh_branch").
			Where("BRUNIT = ? AND (BRDESC LIKE ? OR BRDESC = ?)", "B", "kanwil%", "Jkt KCK").
			Count(&totalRows).Error

	} else if filter2 {
		if request.Sort == "DESC" {
			sortQuery = `PERCENTCOACHING DESC, uker.MAINBR`
		} else {
			sortQuery = `PERCENTCOACHING ASC, uker.MAINBR`
		}

		query = db.Table("dwh_branch uker").
			Select(`uker.REGION, 
					uker.RGDESC, 
					uker.MAINBR, 
					uker.MBDESC, 
					uker.BRANCH, 
					uker.BRDESC, 
					COALESCE(total_coaching.TOTAL, 0) AS TOTALCOACHING, 
					COALESCE(brc.TOTAL, 0) AS TOTALBRC, 
					ROUND(COALESCE((total_coaching.TOTAL / brc.TOTAL) * 100, 0), 0) AS PERCENTCOACHING`).
			Joins(`LEFT JOIN (
					SELECT 
						REGION, 
						MAINBR, 
						COUNT(id) AS TOTAL 
					FROM coaching 
					WHERE 
						(created_at >= ? AND created_at <= ?)
					GROUP BY 
						MAINBR
					) total_coaching ON total_coaching.REGION = uker.REGION AND total_coaching.MAINBR = uker.MAINBR`,
				request.StartDate, lib.FixEndDate(request.EndDate)).
			Joins(`LEFT JOIN (
					SELECT 
						kelolaan.REGION, 
						kelolaan.MAINBR, 
						kelolaan.BRANCH, 
						COUNT(kelolaan.REGION) AS TOTAL 
					FROM 
						(
							SELECT 
								REGION, 
								MAINBR, 
								BRANCH, 
								COUNT(pn) AS TOTAL 
							FROM uker_kelolaan_user 
							GROUP BY BRANCH, pn
						) kelolaan 
					GROUP BY 
						kelolaan.BRANCH
					) brc ON brc.REGION = uker.REGION AND brc.MAINBR = uker.MAINBR`).
			Where("uker.REGION = ?", request.REGION).
			Group("uker.MBDESC").
			Order(sortQuery).
			Limit(request.Limit).
			Offset(request.Offset).
			Find(&responses)

		err = query.Error

		errPagination = db.Table("dwh_branch").
			Select("COUNT(DISTINCT MAINBR) AS pagination").
			Where("REGION = ?", request.REGION).
			Scan(&totalRows).Error

	} else if filter3 {
		if request.Sort == "DESC" {
			sortQuery = `PERCENTCOACHING DESC, uker.BRANCH`
		} else {
			sortQuery = `PERCENTCOACHING ASC, uker.BRANCH`
		}

		query = db.Table("dwh_branch uker").
			Select(`
				uker.REGION, 
				uker.RGDESC, 
				uker.MAINBR, 
				uker.MBDESC, 
				uker.BRANCH, 
				uker.BRDESC, 
				COALESCE(total_coaching.TOTAL,0) AS TOTALCOACHING, 
				COALESCE(brc.TOTAL,0) AS TOTALBRC, 
				ROUND(COALESCE((total_coaching.TOTAL / brc.TOTAL)*100,0),0) AS PERCENTCOACHING`).
			Joins(`
				LEFT JOIN (
					SELECT 
						REGION, 
						MAINBR, 
						BRANCH, 
						COUNT(id) AS TOTAL 
					FROM coaching 
					WHERE created_at >= ? AND created_at <= ? 
					GROUP BY BRANCH
				) total_coaching ON total_coaching.REGION = uker.REGION AND total_coaching.MAINBR = uker.MAINBR AND total_coaching.BRANCH = uker.BRANCH`,
				request.StartDate, lib.FixEndDate(request.EndDate)).
			Joins(`LEFT JOIN (
					SELECT 
						kelolaan.REGION, 
						kelolaan.MAINBR, 
						kelolaan.BRANCH, 
						COUNT(kelolaan.REGION) AS TOTAL 
					FROM (
						SELECT 
							REGION, 
							MAINBR, 
							BRANCH, 
							COUNT(pn) AS TOTAL 
						FROM uker_kelolaan_user 
						GROUP BY BRANCH, pn
						) kelolaan 
						GROUP BY kelolaan.BRANCH) brc ON brc.REGION = uker.REGION AND brc.MAINBR = uker.MAINBR AND brc.BRANCH = uker.BRANCH`).
			Where("uker.REGION = ? AND uker.MAINBR in (?)", request.REGION, mainbrs).
			Order(sortQuery).
			Limit(request.Limit).
			Offset(request.Offset).
			Scan(&responses)

		err = query.Error

		errPagination = db.Table("dwh_branch").
			Select("COUNT(DISTINCT BRANCH) AS pagination").
			Where("REGION = ? AND MAINBR in (?)", request.REGION, mainbrs).
			Scan(&totalRows).Error
	} else if filter4 {
		if request.Sort == "DESC" {
			sortQuery = `PERCENTCOACHING DESC, uker.BRANCH`
		} else {
			sortQuery = `PERCENTCOACHING ASC, uker.BRANCH`
		}

		query = db.Table("dwh_branch uker").
			Select(`
				uker.REGION, 
				uker.RGDESC, 
				uker.MAINBR, 
				uker.MBDESC, 
				uker.BRANCH, 
				uker.BRDESC, 
				COALESCE(total_coaching.TOTAL,0) AS TOTALCOACHING, 
				COALESCE(brc.TOTAL,0) AS TOTALBRC, 
				ROUND(COALESCE((total_coaching.TOTAL / brc.TOTAL)*100,0),0) AS PERCENTCOACHING`).
			Joins(`
				LEFT JOIN (
					SELECT 
						REGION, 
						MAINBR, 
						BRANCH, 
						COUNT(id) AS TOTAL 
					FROM coaching 
					WHERE created_at >= ? AND created_at <= ?
					GROUP BY BRANCH
				) total_coaching ON total_coaching.REGION = uker.REGION AND total_coaching.MAINBR = uker.MAINBR AND total_coaching.BRANCH = uker.BRANCH`,
				request.StartDate, lib.FixEndDate(request.EndDate)).
			Joins(`LEFT JOIN (
					SELECT 
						kelolaan.REGION, 
						kelolaan.MAINBR, 
						kelolaan.BRANCH, 
						COUNT(kelolaan.REGION) AS TOTAL 
					FROM (
						SELECT 
							REGION, 
							MAINBR, 
							BRANCH, 
							COUNT(pn) AS TOTAL 
						FROM uker_kelolaan_user 
						GROUP BY BRANCH, pn
						) kelolaan 
						GROUP BY kelolaan.BRANCH) brc ON brc.REGION = uker.REGION AND brc.MAINBR = uker.MAINBR AND brc.BRANCH = uker.BRANCH`).
			Where("uker.REGION = ? AND uker.MAINBR in (?) AND uker.BRANCH in (?)", request.REGION, mainbrs, branches).
			Order(sortQuery).
			Limit(request.Limit).
			Offset(request.Offset).
			Scan(&responses)

		err = query.Error

		errPagination = db.Table("dwh_branch").
			Select("COUNT(DISTINCT BRANCH) AS pagination").
			Where("REGION = ? AND MAINBR in (?) AND BRANCH in (?)", request.REGION, mainbrs, branches).
			Scan(&totalRows).Error
	}

	repo.logger.Zap.Info("coaching-query-activity-unknown", query)

	if errPagination != nil {
		return responses, totalRows, err
	}

	if err != nil {
		return responses, totalRows, err
	}

	return responses, totalRows, err
}

// 2nd fase
func (repo CoachingRepository) CoachingReportFilterByUkerAllActivity(request *models.CoachingFilterReportRequest) (responses []models.CoachingFilterByUkerAllActivityReportResponse, totalRows int64, totalData int64, err error) {
	query := ""

	filter1 := request.ReportType == "aktivitas_all_uker" &&
		request.Activity == "all" &&
		request.Product == "all" &&
		(request.Title == "all" || request.Title == "Semua") &&
		request.Uker == "bank_wide"

	if filter1 {
		fmt.Println("======= query 1")
		query = `SELECT DISTINCT db.MAINBR, db.RGDESC, db.MBDESC`

		// get activity
		queryGetAllActivity := "SELECT id, kode_activity as code, name FROM activity a"
		rowsGetAllActivity, err := repo.dbRaw.DB.Query(queryGetAllActivity)
		if err != nil {
			return responses, totalRows, totalData, err
		}

		fmt.Println("======= response query 1")
		activity := models.ActivityList{}
		for rowsGetAllActivity.Next() {
			err := rowsGetAllActivity.Scan(
				&activity.Id,
				&activity.Code,
				&activity.Name,
			)
			if err != nil {
				return responses, totalRows, totalData, err
			}

			fmt.Println(activity)
			query +=
				`,(
				SELECT COUNT(*) 
				FROM coaching c 
				WHERE c.REGION = db.REGION
				AND c.activity_id = ` + activity.Id + `
			) as activity` + activity.Id + `
			`
		}

		query += `FROM dwh_branch db WHERE MBDESC LIKE '%kanwil%'`
	}

	repo.logger.Zap.Info("coaching-query-activity-unknown", query)
	rows, err := repo.dbRaw.DB.Query(query)
	defer rows.Close()

	fmt.Println(rows)

	if err != nil {
		return responses, totalRows, totalData, err
	}

	// var results []map[string]interface{}
	// for _, row := range rows {
	// 	key, ok := row["key"].(string)
	// 	if !ok {
	// 		// handle error: key is not a string
	// 	}
	// 	results[key] = row["value"]
	// }

	// fmt.Println("responseee results")
	// fmt.Println(results)

	// response1 := []models.CoachingFilterByUkerAllActivityReportResponse{}\
	// for rows.Next() {
	// 	_ = rows.Scan(&result)
	// }
	// fmt.Println("responseee bytes")
	// fmt.Println(result)
	// bytes, _ := json.Marshal(result)
	// fmt.Println("responseee bytes")
	// fmt.Println(string(bytes))

	return responses, totalRows, totalData, err

}

// end of 2nd fase

func (repo CoachingRepository) CoachingFinalReportFilter(request *models.CoachingFilterReportRequest) (responses []models.CoachingFilterReportFinalResponse, totalRows int64, err error) {
	db := repo.db.DB
	var errPagination error
	sortQuery := ""

	if request.Sort == "DESC" {
		sortQuery = `c.created_at DESC, id`
	} else {
		sortQuery = `c.created_at ASC, id`
	}

	query := db.Table("coaching c ").
		Select(`
			c.id, 
			c.created_at as date, 
			c.BRANCH, 
			c.BRDESC, 
			a.name as activity, 
			p.product, 
			ca.risk_issue, 
			ca.judul_materi as materi
		`).
		Joins(`
			JOIN coaching_activity ca ON ca.coaching_id = c.id 
			JOIN product p ON p.id = c.product_id 
			JOIN activity a ON a.kode_activity = c.activity_id
		`).
		Where(`ca.judul_materi = ?`, request.Title).
		Where(`c.activity_id = ?`, request.Activity).
		Where(`c.product_id = ?`, request.Product).
		Where(`c.created_at >= ? AND c.created_at <= ?`, request.StartDate, lib.FixEndDate(request.EndDate))

	queryPagination := db.Table("coaching c").
		Select("COUNT(c.id) as pagination").
		Joins(`
			JOIN coaching_activity ca ON ca.coaching_id = c.id 
			JOIN product p ON p.id = c.product_id 
			JOIN activity a ON a.kode_activity = c.activity_id
		`).
		Where("ca.judul_materi = ?", request.Title).
		Where("c.activity_id = ?", request.Activity).
		Where("c.product_id = ?", request.Product).
		Where("c.created_at >= ? AND c.created_at <= ?", request.StartDate, lib.FixEndDate(request.EndDate))

	if request.REGION != "all" {
		query = query.Where("c.REGION = ?", request.REGION)
		queryPagination = queryPagination.Where("c.REGION = ?", request.REGION)
	}

	if request.REGION != "all" && request.MAINBR != "all" {
		mainbrs := strings.Split(request.MAINBR, ",")

		query = query.Where("c.MAINBR in (?)", mainbrs)
		queryPagination = queryPagination.Where("c.MAINBR in (?)", mainbrs)
	}

	if request.REGION != "all" && request.MAINBR != "all" && request.BRANCH != "all" {
		branches := strings.Split(request.BRANCH, ",")

		query = query.Where("c.BRANCH in (?)", branches)
		queryPagination = queryPagination.Where("c.MAINBR = (?)", branches)
	}

	// if request.BRANCH != "all" {
	// 	branches := strings.Split(request.BRANCH, ",")

	// 	query = query.Where("c.BRANCH in (?)", branches)
	// 	queryPagination = queryPagination.Where("c.MAINBR = (?)", branches)
	// }

	query = query.Order(sortQuery).
		Limit(request.Limit).
		Offset(request.Offset).
		Find(&responses)

	queryPagination = queryPagination.Scan(&totalRows)

	err = query.Error

	errPagination = queryPagination.Error

	repo.logger.Zap.Info("coaching-queryPagination-activity-unknown", queryPagination)

	if errPagination != nil {
		return responses, totalRows, err
	}

	repo.logger.Zap.Info("coaching-query-activity-unknown", query)

	if err != nil {
		return responses, totalRows, err
	}

	return responses, totalRows, err
}

func (repo CoachingRepository) CoachingReportFilterByUkerComplete(request *models.CoachingFilterReportByUker) (responses []models.CoachingFilterReportFinalResponseNull, totalRows int, err error) {
	var rows *sql.Rows
	var errPagination error

	query := `SELECT c.id, c.created_at, c.BRANCH, c.BRDESC, a.name as activity, p.product, ca.risk_issue, ca.judul_materi as materi
			FROM coaching c 
			JOIN coaching_activity ca ON ca.coaching_id = c.id 
			JOIN product p ON p.id = c.product_id 
			JOIN activity a ON a.kode_activity = c.activity_id
			WHERE c.BRANCH = ?
			AND c.created_at >= ? AND c.created_at <= ?
			GROUP BY c.id
			ORDER BY c.id
			LIMIT ? OFFSET ?`

	queryPagination := `
			SELECT 
				COUNT(DISTINCT c.id)
			FROM coaching c 
			JOIN coaching_activity ca ON ca.coaching_id = c.id 
			JOIN product p ON p.id = c.product_id 
			JOIN activity a ON a.kode_activity = c.activity_id
			WHERE c.BRANCH = ?
			AND c.created_at >= ? AND c.created_at <= ?
		`
	rows, err = repo.db.DB.Raw(query, request.BRANCH, request.StartDate, lib.FixEndDate(request.EndDate), strconv.Itoa(request.Limit), strconv.Itoa(request.Offset)).Rows()
	defer rows.Close()

	//QUERY PAGINATION
	errPagination = repo.db.DB.Raw(queryPagination, request.BRANCH, request.StartDate, lib.FixEndDate(request.EndDate)).Scan(&totalRows).Error

	repo.logger.Zap.Info("coaching-queryPagination-activity-unknown", queryPagination)

	fmt.Println("errPagination", errPagination)

	if errPagination != nil {
		return responses, totalRows, err
	}

	repo.logger.Zap.Info("coaching-query-activity-unknown", query)
	fmt.Println(rows)

	repo.logger.Zap.Info("coaching-rows-activity-unknown", rows)
	if err != nil {
		return responses, totalRows, err
	}

	response := models.CoachingFilterReportFinalResponseNull{}
	for rows.Next() {
		_ = rows.Scan(
			&response.Id,
			&response.Date,
			&response.BRANCH,
			&response.BRDESC,
			&response.Activity,
			&response.Product,
			&response.RiskIssue,
			&response.Materi,
		)
		responses = append(responses, response)

		fmt.Println(responses)
	}

	if err = rows.Err(); err != nil {
		return responses, totalRows, err
	}

	return responses, totalRows, err
}

func (repo CoachingRepository) CoachingReportDetail(request *models.CoachingReportDetailRequest) (responses models.CoachingReportDetailResponse, err error) {
	var rows1 *sql.Rows
	var rows2 *sql.Rows
	var err1 error
	var err2 error

	//get coaching detail
	query1 := `
				SELECT c.id, c.no_pelaporan, c.BRDESC as unit_kerja, c.jenis_peserta, c.jumlah_peserta 
				FROM coaching c WHERE c.id = ?
			  `

	//get coaching materis
	query2 := `
				SELECT ca.id, a.name as activity, sa.name as sub_activity, p.product, ca.risk_issue, ca.judul_materi, ca.rekomendasi_materi, ca.materi_tambahan
				FROM coaching_activity ca 
				JOIN coaching c ON c.id = ca.coaching_id
				JOIN activity a ON a.kode_activity = c.activity_id 
				JOIN sub_activity sa ON sa.id = c.sub_activity_id 
				JOIN product p ON p.id = c.product_id 
				WHERE ca.coaching_id = ?
			  `

	repo.logger.Zap.Info("coaching-query-detail", query1)
	repo.logger.Zap.Info("coaching-query-detail-materis", query2)

	rows1, err1 = repo.db.DB.Raw(query1, request.ID).Rows()
	defer rows1.Close()

	repo.logger.Zap.Info("coaching-rows-detail", rows1)

	if err1 != nil {
		err = err1
		return responses, err
	}

	rows2, err2 = repo.db.DB.Raw(query2, request.ID).Rows()
	defer rows2.Close()

	repo.logger.Zap.Info("coaching-rows-detail-materis", rows2)

	if err2 != nil {
		err = err2
		return responses, err
	}

	response1 := models.CoachingReportDetail{}
	for rows1.Next() {
		_ = rows1.Scan(
			&response1.ID,
			&response1.NoPelaporan,
			&response1.UnitKerja,
			&response1.JenisPeserta,
			&response1.JumlahPeserta,
		)

		responses.CoachingDetail = response1
	}

	fmt.Println("")
	fmt.Println("========================== response1", response1)
	fmt.Println("")

	var responseMateris []models.CoachingReportDetailMateri
	response2 := models.CoachingReportDetailMateri{}
	for rows2.Next() {
		_ = rows2.Scan(
			&response2.ID,
			&response2.Activity,
			&response2.SubActivity,
			&response2.Product,
			&response2.RiskIssue,
			&response2.JudulMateri,
			&response2.RekomendasiMateri,
			&response2.MateriTambahan,
		)
		responseMateris = append(responseMateris, response2)
	}

	fmt.Println("")
	fmt.Println("========================== responseMateris", &responseMateris)
	fmt.Println("")

	responses.CoachingMateris = responseMateris

	if err1 = rows1.Err(); err1 != nil {
		return responses, err
	}

	if err2 = rows2.Err(); err1 != nil {
		return responses, err
	}

	return responses, err
}

func (repo CoachingRepository) CoachingReportMateriList(request *models.CoachingReportMateriRequest) (responses []models.CoachingDetailMateriResponseNull, err error) {
	var rows *sql.Rows
	fileId := strings.Split(request.ID, ",")

	query := `SELECT id, nama_lampiran, filename, path FROM risk_indicator_map_files rimf WHERE id IN ?`

	// repo.logger.Zap.Info("coaching-query-activity-unknown", query)
	rows, err = repo.db.DB.Raw(query, fileId).Rows()
	defer rows.Close()

	// repo.logger.Zap.Info("coaching-rows-activity-unknown", rows)
	if err != nil {
		return responses, err
	}

	response := models.CoachingDetailMateriResponseNull{}
	for rows.Next() {
		_ = rows.Scan(
			&response.ID,
			&response.NamaLampiran,
			&response.Filename,
			&response.Path,
		)
		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, err
	}

	return responses, err
}

// CoachingReportList implements CoachingDefinition
func (repo CoachingRepository) CoachingReportList(request *models.CoachingReportListRequest) (responses []models.CoachingReportListResponse, totalRows int, err error) {
	db := repo.db.DB.Table("report_list_coaching rlc ")

	sorting := ""

	if request.Sort == "desc" {
		sorting = `rlc.id DESC`
	} else {
		sorting = `rlc.id ASC`
	}

	query := db.Select(`
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
	`).Order(sorting)

	if request.NoPelaporan != "" {
		query = query.Where("rlc.no_pelaporan = ?", request.NoPelaporan)
	}

	if request.REGION != "all" {
		query = query.Where("rlc.REGION = ?", request.REGION)
	}

	if request.MAINBR != "all" {
		mainbrs := strings.Split(request.MAINBR, ",")
		query = query.Where("rlc.MAINBR in (?)", mainbrs)
	}

	if request.BRANCH != "all" {
		branches := strings.Split(request.BRANCH, ",")
		query = query.Where("rlc.BRANCH in (?)", branches)
	}

	if request.ActivityID != "all" {
		query = query.Where("rlc.activity_id = ?", request.ActivityID)
	}

	if request.RiskIssueID != "all" {
		query = query.Where("rlc.isu_risiko_id LIKE ?", fmt.Sprintf("%%%s%%", request.RiskIssueID))
	}

	if request.JudulMateri != "" {
		query = query.Where("rlc.judul_materi LIKE ?", fmt.Sprintf("%%%s%%", request.JudulMateri))
	}

	if request.Status != "" && request.Status != "Semua" {
		query = query.Where("rlc.status = ?", request.Status)
	}

	if request.StartDate != "" && lib.FixEndDate(request.EndDate) != "" {
		query = query.Where(`rlc.periode BETWEEN ? AND ?`, request.StartDate, lib.FixEndDate(request.EndDate))
	}

	var count int64

	query.Count(&count)
	totalRows = int(count)

	if request.Limit != 0 {
		query = query.Limit(request.Limit)
	}

	if request.Offset != 0 {
		query = query.Offset(request.Offset)
	}

	err = query.Find(&responses).Error

	return responses, totalRows, err
}

// Versioning 23/10/2023
// GetJudulMateri implements CoachingDefinition.
func (repo CoachingRepository) GetJudulMateri(id int64) (responses []models.JudulMateriCoaching, err error) {
	db := repo.db.DB.Table(`coaching_activity`).
		Select(`id, risk_issue,judul_materi`).Where(`coaching_id = ?`, id)

	err = db.Find(&responses).Error

	return responses, err
}

func (coaching CoachingRepository) CoachingFrekuensiRpt(request *models.FrekuensiCoachingRequest) (responses []models.FrekuensiCoachingResponse, totalRows int, err error) {
	tableName := ""
	if request.Kegiatan == "coaching" {
		tableName = "coaching c"

		columnMappings := map[string]string{
			"activity":       "a.name AS 'aktivitas'",
			"product":        "p.product AS 'produk'",
			"risk_event":     "ca.risk_issue AS 'risk_event'",
			"risk_indicator": "ca.judul_materi AS 'risk_indicator'",
		}

		joinMappings := map[string]string{
			"activity": "JOIN activity a ON a.id = c.activity_id",
			"product":  "JOIN product p ON p.id = c.product_id",
		}

		groupMappings := map[string]string{
			"activity":       "c.activity_id",
			"product":        "c.product_id",
			"risk_event":     "ca.risk_issue",
			"risk_indicator": "ca.judul_materi",
		}

		selectParts := []string{"COUNT(*) AS 'jumlah'"}
		joinParts := make(map[string]bool)
		groupByParts := make(map[string]bool)

		for _, gd := range request.GroupData {
			if column, exists := columnMappings[gd.Column]; exists {
				selectParts = append(selectParts, column)
				if join, ok := joinMappings[gd.Column]; ok {
					joinParts[join] = true
				}

				if groupBy, ok := groupMappings[gd.Column]; ok {
					groupByParts[groupBy] = true
				}
			}
		}

		var joins []string
		for j := range joinParts {
			joins = append(joins, j)
		}

		var groupBy []string
		for g := range groupByParts {
			groupBy = append(groupBy, g)
		}

		query := coaching.db.DB.Table(tableName).
			Select(strings.Join(selectParts, ", ")).
			Joins("JOIN coaching_activity ca ON ca.coaching_id = c.id").
			Joins(strings.Join(joins, " ")).
			Where("c.action IN(?, ?)", "Update", "Selesai").
			Group(strings.Join(groupBy, ", "))

		queryPaging := coaching.db.DB.Table(tableName).
			Select(`Count(*)`).
			Joins("JOIN coaching_activity ca ON ca.coaching_id = c.id").
			Joins(strings.Join(joins, " ")).
			Where("c.action IN(?, ?)", "Update", "Selesai").
			Group(strings.Join(groupBy, ", "))

		if request.REGION != "all" && request.REGION != "" {
			query = query.Where("c.REGION = ?", request.REGION)
			queryPaging = queryPaging.Where("c.REGION = ?", request.REGION)
		}

		if request.MAINBR != "all" && request.MAINBR != "" {
			mainbrs := strings.Split(request.MAINBR, ",")

			query = query.Where("c.MAINBR in (?)", mainbrs)
			queryPaging = queryPaging.Where("c.MAINBR in (?)", mainbrs)
		}

		if request.BRANCH != "all" && request.BRANCH != "" {
			branches := strings.Split(request.BRANCH, ",")

			query = query.Where("c.BRANCH in (?)", branches)
			queryPaging = queryPaging.Where("c.BRANCH in (?)", branches)
		}

		if request.StartDate != "" && request.EndDate != "" {
			query = query.Where("c.created_at BETWEEN ? AND ?", request.StartDate, lib.FixEndDate(request.EndDate))
			queryPaging = queryPaging.Where("c.created_at BETWEEN ? AND ?", request.StartDate, lib.FixEndDate((request.EndDate)))
		}

		var Count int64
		queryPaging = queryPaging.Count(&Count)

		totalRows = int(Count)

		query = query.Order("jumlah DESC").Limit(request.Limit).Offset(request.Offset)

		responses = []models.FrekuensiCoachingResponse{}

		err = query.Scan(&responses).Error

		return responses, totalRows, err

	} else {
		return nil, 0, err
	}
}
