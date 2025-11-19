package briefing

import (
	"fmt"
	"math"
	"riskmanagement/lib"
	models "riskmanagement/models/briefing"
	"strconv"
	"time"

	"database/sql"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"

	"strings"
)

type BriefingDefinition interface {
	WithTrx(trxHandle *gorm.DB) BriefingRepository
	GetAll() (responses []models.BriefingResponse, err error)
	GetData() (responses []models.BriefingResponse, err error)
	GetDataWithPagination(request *models.BriefingPagination) (responses []models.BriefingResponse, totalRows int, totalData int, err error)
	GetOne(id int64) (responses models.BriefingResponse, err error)
	Store(request *models.Briefing, tx *gorm.DB) (responses *models.Briefing, err error)
	Delete(request *models.BriefingUpdateDelete, include []string, tx *gorm.DB) (responses bool, err error)
	DeleteBriefingMateri(id int64, tx *gorm.DB) (err error)
	UpdateAllBrief(request *models.BriefingUpdateMateri, include []string, tx *gorm.DB) (responses bool, err error)
	FilterBriefing(request *models.BriefingFilterRequest) (responses []models.BriefingResponse, totalRows int, totalData int, err error)
	GetNoPelaporan(request *models.NoPelaporanRequest) (responses []models.NoPelaporanNullResponse, err error)
	BriefingReportFilter(request *models.BriefingFilterReport) (responses []models.BriefingFilterReportResponse, totalAktivitas int64, totalRows int64, err error)
	BriefingReportFilterComplete(request *models.BriefingFilterReport) (responses []models.BriefingFilterReportFinalResponse, totalRows int, err error)
	BriefingReportDetail(request *models.BriefingReportDetailRequest) (responses models.BriefingReportDetailResponse, err error)
	BriefingReportMateriList(request *models.BriefingReportMateriRequest) (responses []models.BriefingDetailMateriResponseNull, err error)
	BriefingReportByUkerFilter(request *models.BriefingFilterReportByUker) (responses []models.BriefingFilterReportByUkerResponse, totalRows int64, err error)
	BriefingReportFilterByUkerComplete(request *models.BriefingFilterReportByUker) (responses []models.BriefingFilterReportFinalResponseNull, totalRows int, err error)
	BriefingReportList(request *models.BriefingReportListRequest) (responses []models.BriefingReportListResponse, totalRows int, err error)
	BriefingFrekuensiRpt(request *models.FrekuensiBriefingRequest) (responses []models.FrekuensiBriefingResponse, totalRows int, err error)
	// versioning Add by panji 23/10/2023
	GetJudulMateri(id int64) (responses []models.JudulMateriBriefing, err error)
}

type BriefingRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewBriefingRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) BriefingDefinition {
	return BriefingRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// Delete implements BriefingDefinition
func (briefing BriefingRepository) Delete(request *models.BriefingUpdateDelete, include []string, tx *gorm.DB) (responses bool, err error) {
	return true, tx.Save(&request).Error
}

// DeleteBriefingMateri implements BriefingDefinition
func (briefing BriefingRepository) DeleteBriefingMateri(id int64, tx *gorm.DB) (err error) {
	return tx.Where("id = ?", id).Delete(&models.BriefingMateriRequest{}).Error
}

// GetAll implements BriefingDefinition
func (briefing BriefingRepository) GetAll() (responses []models.BriefingResponse, err error) {
	return responses, briefing.db.DB.Find(&responses).Error
}

// GetOne implements BriefingDefinition
func (briefing BriefingRepository) GetOne(id int64) (responses models.BriefingResponse, err error) {
	err = briefing.db.DB.Raw(`
	SELECT 
		brf.id,
		brf.no_pelaporan,
		brf.REGION,
		brf.RGDESC,
		brf.MAINBR,
		brf.MBDESC,
		brf.BRANCH,
		brf.BRDESC,
		brf.unit_kerja,
		brf.jenis_peserta,
		brf.jabatan_peserta,
		brf.jumlah_peserta,
		brf.list_peserta,
		brf.maker_id,
		brf.maker_desc,
		brf.maker_date,
		brf.last_maker_id,
		brf.last_maker_desc,
		brf.last_maker_date,
		CASE
			WHEN brf.status = "01a" AND brf.action = "Selesai" THEN "Selesai"
			WHEN brf.status = "01a" AND brf.action = "Draft" THEN "Draft"
			WHEN brf.status = "02b" AND (brf.action = "Update" OR brf.action ="Selesai")   THEN "Selesai"
			ELSE "Delete"
		END 'status',
		brf.deleted,
		brf.created_at,
		brf.updated_at
	FROM briefing brf
	WHERE brf.id = ?`, id).Find(&responses).Error

	if err != nil {
		briefing.logger.Zap.Error(err)
		return responses, err
	}
	return responses, err
}

// Store implements BriefingDefinition
func (briefing BriefingRepository) Store(request *models.Briefing, tx *gorm.DB) (responses *models.Briefing, err error) {
	return request, tx.Save(&request).Error
}

// WithTrx implements BriefingDefinition
func (briefing BriefingRepository) WithTrx(trxHandle *gorm.DB) BriefingRepository {
	if trxHandle == nil {
		briefing.logger.Zap.Error("transaction Database not found in gin context")
		return briefing
	}
	briefing.db.DB = trxHandle
	return briefing
}

// GetData implements BriefingDefinition
func (briefing BriefingRepository) GetData() (responses []models.BriefingResponse, err error) {
	rows, err := briefing.db.DB.Raw(`
		SELECT 
			brf.id 'id',
			brf.no_pelaporan 'no_pelaporan',
			brf.BRDESC 'unit_kerja',
			CASE
				WHEN brf.status = "01a" AND brf.action = "Selesai" THEN "Selesai"
				WHEN brf.status = "01a" AND brf.action = "Draft" THEN "Draft"
				WHEN brf.status = "02b" AND (brf.action = "Update" OR brf.action ="Selesai")   THEN "Selesai"
				ELSE "Delete"
			END 'status_brf'
		FROM briefing brf 
		WHERE brf.deleted != 1
	`).Rows()

	// SELECT
	// 		brf.id 'id',
	// 		brf.no_pelaporan 'no_pelaporan',
	// 		brf.unit_kerja 'unit_kerja',
	// 		act.name 'aktifitas',
	// 		bm.judul_materi 'judul_materi',
	// 		CASE
	// 			WHEN brf.status = "01a" && brf.action = "Draft" THEN "Draft"
	// 			WHEN brf.status = "02b" && (brf.action = "Update" || brf.action ="Selesai")   THEN "Selesai"
	// 			ELSE "Delete"
	// 		END 'status_brf'
	// 	FROM briefing brf
	// 	JOIN briefing_materis bm ON bm.briefing_id = brf.id
	// 	JOIN activity act ON bm.activity_id = act.id
	// 	WHERE brf.deleted != 1
	// 	GROUP BY brf.id

	defer rows.Close()

	var brief models.BriefingResponse
	for rows.Next() {
		briefing.db.DB.ScanRows(rows, &brief)
		responses = append(responses, brief)
	}
	return responses, err
}

// GetDataWithPagination implements BriefingDefinition
func (brf BriefingRepository) GetDataWithPagination(request *models.BriefingPagination) (responses []models.BriefingResponse, totalRows int, totalData int, err error) {

	branches := strings.Split(request.Branches, ",")

	db := brf.db.DB.Table("briefing brf").
		Select(`brf.id 'id',
				brf.no_pelaporan 'no_pelaporan',
				brf.BRDESC 'unit_kerja',
				CASE
					WHEN brf.status = "01a" AND brf.action = "Selesai" THEN "Selesai"
					WHEN brf.status = "01a" AND brf.action = "Draft" THEN "Draft"
					WHEN brf.status = "02b" AND (brf.action = "Update" OR brf.action ="Selesai")   THEN "Selesai"
					ELSE "Delete"
				END 'status'`).
		Where(`brf.deleted != 1`).
		// Where(`brf.BRANCH in (?)`, branches).
		Order(`brf.id DESC`)

	// if request.Branches != "" {
	// 	branches := strings.Split(request.Branches, ",")
	// 	db = db.Where(`brf.BRANCH in (?)`, branches)
	// }

	if request.Kostl != "" {
		db = db.Where("brf.maker_id = ?", request.Pernr)
	} else {
		db = db.Where(`brf.BRANCH in (?)`, branches)
	}

	var count int64
	db.
		Group(`brf.id`).
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

// UpdateAllBrief implements BriefingDefinition
func (briefing BriefingRepository) UpdateAllBrief(request *models.BriefingUpdateMateri, include []string, tx *gorm.DB) (responses bool, err error) {
	return true, tx.Save(&request).Error
}

// FilterBriefing implements BriefingDefinition
func (briefing BriefingRepository) FilterBriefing(request *models.BriefingFilterRequest) (responses []models.BriefingResponse, totalRows int, totalData int, err error) {
	branches := strings.Split(request.Branches, ",")

	db := briefing.db.DB

	// create a new query builder with the base query
	queryBuilder := db.Model(&responses).
		Select(`
            briefing.id AS id,
            briefing.no_pelaporan AS no_pelaporan,
            briefing.BRDESC AS unit_kerja,
            CASE
                WHEN briefing.status = "01a" AND briefing.action = "Selesai" THEN "Selesai"
                WHEN briefing.status = "01a" AND briefing.action = "Draft" THEN "Draft"
                WHEN briefing.status = "02b" AND (briefing.action = "Update" OR briefing.action ="Selesai") THEN "Selesai"
                ELSE "Delete"
            END AS status
        `).
		Joins("LEFT JOIN briefing_materis bm ON bm.briefing_id = briefing.id").
		Joins("LEFT JOIN activity a ON a.kode_activity = bm.activity_id").
		Where("briefing.deleted != 1").
		// Where("briefing.BRANCH in (?)", branches).
		Order("briefing.id DESC")

	// if request.Branches != "" {
	// 	queryBuilder = queryBuilder.Where("briefing.BRANCH in (?)", branches)
	// }

	if request.Kostl != "" {
		queryBuilder = queryBuilder.Where("briefing.maker_id = ?", request.Pernr)
	} else {
		queryBuilder = queryBuilder.Where("briefing.BRANCH in (?)", branches)
	}

	// add dynamic where clauses
	if request.NoPelaporan != "" {
		queryBuilder = queryBuilder.Where("briefing.no_pelaporan = ?", request.NoPelaporan)
	}

	if request.UnitKerja != "" {
		// queryBuilder = queryBuilder.Where("briefing.BRDESC LIKE ?", fmt.Sprintf("%%%s%%", request.UnitKerja))
		queryBuilder = queryBuilder.Where("briefing.BRANCH = ?", request.UnitKerja)
	}

	if request.ActivityID != "" {
		queryBuilder = queryBuilder.Where("bm.activity_id = ?", request.ActivityID)
	}

	if request.JudulMateri != "" {
		queryBuilder = queryBuilder.Where("bm.judul_materi LIKE ?", fmt.Sprintf("%%%s%%", request.JudulMateri))
	}

	if request.Status != "" && request.Status != "Semua" && request.Status != "Selesai" {
		queryBuilder = queryBuilder.Where("briefing.action = ?", request.Status)
	}

	if request.Status == "Selesai" {
		queryBuilder = queryBuilder.Where("briefing.action IN (?, ?)", "Update", "Selesai")
	}

	if request.TglAwal != "" && request.TglAkhir != "" {
		// queryBuilder = queryBuilder.Where("CAST(briefing.created_at AS DATE) BETWEEN ? AND ?", request.TglAwal, request.TglAkhir)
		queryBuilder = queryBuilder.Where("(briefing.created_at >= ? AND briefing.created_at <= ?)", request.TglAwal, lib.FixEndDate(request.TglAkhir))
	}

	// count the total rows
	var count int64
	queryBuilder.
		Group("bm.briefing_id").
		Count(&count)

	totalRows = int(count)

	fmt.Println("totalRows", totalRows)

	// execute the query
	queryBuilder.
		Limit(request.Limit).
		Offset(request.Offset).
		Find(&responses)

	// calculate the total pages
	totalPages := int(math.Ceil(float64(totalRows) / float64(request.Limit)))
	return responses, totalPages, totalRows, err
}

// GetNoPelaporan implements BriefingDefinition
func (briefing BriefingRepository) GetNoPelaporan(request *models.NoPelaporanRequest) (responses []models.NoPelaporanNullResponse, err error) {
	kode := "BR-"
	today := lib.GetTimeNow("date2")

	if request.ORGEH != "" {
		kode += "%" + request.ORGEH + "-" + today + "%"
	}

	query := `SELECT RIGHT(CONCAT("0000",(count(*) + 1)), 4) 'no_pelaporan' FROM briefing WHERE no_pelaporan LIKE ?`

	briefing.logger.Zap.Info(query)
	rows, err := briefing.dbRaw.DB.Query(query, kode)

	briefing.logger.Zap.Info("rows ", rows)
	if err != nil {
		return responses, err
	}

	response := models.NoPelaporanNullResponse{}
	for rows.Next() {
		_ = rows.Scan(
			&response.NoPelaporan,
		)

		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, err
	}

	return responses, err
}

func (repo BriefingRepository) BriefingReportFilter(request *models.BriefingFilterReport) (responses []models.BriefingFilterReportResponse, totalAktivitas int64, totalRows int64, err error) {
	db := repo.db.DB
	query := db
	sortQuery := ""
	var rows *sql.Rows
	var errPagination error

	// Set sort
	if request.Sort == "DESC" {
		sortQuery = "total DESC, id"
	} else {
		sortQuery = "total ASC, id"
	}

	// Apply filters
	// filter 1
	if request.ReportType == "aktivitas" &&
		request.Activity == "all" &&
		request.Product == "all" &&
		(request.Title == "all" || request.Title == "Semua") {
		fmt.Println("Filter 1")
		// Define subquery
		subquery := db.Table("briefing_materis").
			Select("COUNT(*) as total").
			Joins("JOIN briefing ON briefing.id = briefing_materis.briefing_id")

		totalAktivitasQuery := db.Table("briefing_materis").
			Select("COUNT(*) as totalAktivitas").
			Joins("JOIN briefing ON briefing.id = briefing_materis.briefing_id")

		if request.REGION != "all" {
			subquery = subquery.Where("briefing.REGION = ?", request.REGION)
			totalAktivitasQuery = totalAktivitasQuery.Where("briefing.REGION = ?", request.REGION)
		}

		if request.REGION != "all" && request.MAINBR != "all" {
			mainbrs := strings.Split(request.MAINBR, ",")
			subquery = subquery.Where("briefing.MAINBR in (?)", mainbrs)
			totalAktivitasQuery = totalAktivitasQuery.Where("briefing.MAINBR in (?)", mainbrs)
		}

		if request.REGION != "all" && request.MAINBR != "all" && request.BRANCH != "all" {
			branches := strings.Split(request.BRANCH, ",")
			subquery = subquery.Where("briefing.BRANCH in (?)", branches)
			totalAktivitasQuery = totalAktivitasQuery.Where("briefing.BRANCH in (?)", branches)
		}

		subquery = subquery.
			Where("briefing_materis.activity_id = activity.kode_activity").
			// Where("briefing_materis.created_at BETWEEN ? AND ?", request.StartDate, lib.FixEndDate(request.EndDate))
			Where("briefing_materis.created_at >= ? AND  briefing_materis.created_at<= ?", request.StartDate, lib.FixEndDate(request.EndDate))

		totalAktivitasQuery = totalAktivitasQuery.
			Where("briefing_materis.created_at >= ? AND  briefing_materis.created_at<= ?", request.StartDate, lib.FixEndDate(request.EndDate)).
			Find(&totalAktivitas)

		// Define query
		query := db.Table("activity").
			Select("activity.id, activity.kode_activity as code, activity.name, (?) total", subquery)

		query = query.Limit(request.Limit).
			Offset(request.Offset).
			Order(sortQuery).
			Find(&responses)

		// query pagination
		errPagination = db.Table("activity").Count(&totalRows).Error

		//filter 2
	} else if request.ReportType == "aktivitas" &&
		request.Activity != "all" &&
		request.Product == "all" &&
		(request.Title == "all" || request.Title == "Semua") {
		fmt.Println("Filter 2")
		// Define subquery
		subquery := db.Table("briefing_materis").
			Select("COUNT(*) as total").
			Joins("JOIN briefing ON briefing.id = briefing_materis.briefing_id")

		totalAktivitasQuery := db.Table("briefing_materis").
			Select("COUNT(*) as totalAktivitas").
			Joins("JOIN briefing ON briefing.id = briefing_materis.briefing_id")

		if request.REGION != "all" {
			subquery = subquery.Where("briefing.REGION = ?", request.REGION)
			totalAktivitasQuery = totalAktivitasQuery.Where("briefing.REGION = ?", request.REGION)
		}

		if request.REGION != "all" && request.MAINBR != "all" {
			mainbrs := strings.Split(request.MAINBR, ",")
			subquery = subquery.Where("briefing.MAINBR in (?)", mainbrs)
			totalAktivitasQuery = totalAktivitasQuery.Where("briefing.MAINBR in (?)", mainbrs)
		}

		if request.REGION != "all" && request.MAINBR != "all" && request.BRANCH != "all" {
			branches := strings.Split(request.BRANCH, ",")
			subquery = subquery.Where("briefing.BRANCH in (?)", branches)
			totalAktivitasQuery = totalAktivitasQuery.Where("briefing.BRANCH in (?)", branches)
		}

		subquery = subquery.
			Where("briefing_materis.activity_id = ?", request.Activity).
			Where("briefing_materis.product_id = product.id").
			// Where("briefing_materis.created_at BETWEEN ? AND ?", request.StartDate, lib.FixEndDate(request.EndDate))
			Where("briefing_materis.created_at >= ? AND  briefing_materis.created_at <= ?", request.StartDate, lib.FixEndDate(request.EndDate))

		totalAktivitasQuery = totalAktivitasQuery.
			Where("briefing_materis.activity_id = ?", request.Activity).
			// Where("briefing_materis.created_at BETWEEN ? AND ?", request.StartDate, lib.FixEndDate(request.EndDate))
			Where("briefing_materis.created_at >= ? AND  briefing_materis.created_at <= ?", request.StartDate, lib.FixEndDate(request.EndDate)).
			Find(&totalAktivitas)

		// Define query
		query := db.Table("product").
			Select("product.id, product.kode_product as code, product.product as name, (?) total", subquery).
			Where("product.activity_id = ?", request.Activity)

		errPagination = query.Count(&totalRows).Error

		// Apply bank wide filter
		query = query.Limit(request.Limit).
			Offset(request.Offset).
			Order(sortQuery).
			Find(&responses)

		// errPagination = db.Table("product").Where("product.activity_id", request.Activity).Count(&totalRows).Error

		//filter 3
	} else if request.ReportType == "aktivitas" &&
		request.Activity != "all" &&
		request.Product != "all" &&
		(request.Title == "all" || request.Title == "Semua") {
		fmt.Println("Filter 3")
		// Define query
		query := db.Table("briefing_materis").
			Select("briefing_materis.activity_id  as id, briefing_materis.product_id as code, briefing_materis.judul_materi as name, COUNT(*) as total").
			Where("briefing_materis.activity_id = ?", request.Activity).
			Where("briefing_materis.product_id = ?", request.Product).
			// Where("briefing_materis.created_at BETWEEN ? AND ?", request.StartDate, lib.FixEndDate(request.EndDate))
			Where("briefing_materis.created_at >= ? AND  briefing_materis.created_at<= ?", request.StartDate, lib.FixEndDate(request.EndDate)).
			Group("briefing_materis.judul_materi").
			Count(&totalRows)

		totalAktivitasQuery := db.Table("briefing_materis").
			Select("COUNT(*) as totalAktivitas").
			Joins("JOIN briefing ON briefing.id = briefing_materis.briefing_id")

		// queryPagination := db.Table("briefing_materis").
		// 	Select("SUM(COUNT(DISTINCT judul_materi)) OVER() AS pagination")

		if request.REGION != "all" {
			query = query.Joins("JOIN briefing ON briefing.id = briefing_materis.briefing_id")
			query = query.Where("briefing.REGION = ?", request.REGION)
			totalAktivitasQuery = totalAktivitasQuery.Where("briefing.REGION = ?", request.REGION)

			// queryPagination = queryPagination.Joins("JOIN briefing ON briefing.id = briefing_materis.briefing_id")
			// queryPagination = queryPagination.Where("briefing.REGION = ?", request.REGION)
		}

		if request.REGION != "all" && request.MAINBR != "all" {
			mainbrs := strings.Split(request.MAINBR, ",")
			query = query.Where("briefing.MAINBR in (?)", mainbrs)
			totalAktivitasQuery = totalAktivitasQuery.Where("briefing.MAINBR in (?)", mainbrs)
			// queryPagination = queryPagination.Where("briefing.MAINBR = ?", request.MAINBR)
		}

		if request.REGION != "all" && request.MAINBR != "all" && request.BRANCH != "all" {
			branches := strings.Split(request.BRANCH, ",")
			query = query.Where("briefing.BRANCH in (?)", branches)
			totalAktivitasQuery = totalAktivitasQuery.Where("briefing.MAINBR = ?", request.MAINBR)
			// queryPagination = queryPagination.Where("briefing.BRANCH in (?)", branches)
		}

		// query = query.

		totalAktivitasQuery = totalAktivitasQuery.
			Where("briefing_materis.activity_id = ?", request.Activity).
			Where("briefing_materis.product_id = ?", request.Product).
			// Where("briefing_materis.created_at BETWEEN ? AND ?", request.StartDate, lib.FixEndDate(request.EndDate))
			Where("briefing_materis.created_at >= ? AND  briefing_materis.created_at<= ?", request.StartDate, lib.FixEndDate(request.EndDate)).
			Find(&totalAktivitas)

		// Apply bank wide filter
		query = query.
			Limit(request.Limit).
			Offset(request.Offset).
			Order(sortQuery).
			Find(&responses)

		//query pagination
		// queryPagination = queryPagination.
		// 	Where("briefing_materis.activity_id", request.Activity).
		// 	Where("briefing_materis.product_id", request.Product).
		// 	Where("briefing_materis.created_at >= ? AND  briefing_materis.created_at<= ?", request.StartDate, lib.FixEndDate(request.EndDate)).
		// 	Scan(&totalRows)

		// errPagination = queryPagination.Error
	}

	//QUERY
	repo.logger.Zap.Info("briefing-query-activity-unknown", query)

	repo.logger.Zap.Info("briefing-rows-activity-unknown", rows)
	if err != nil {
		return responses, totalAktivitas, totalRows, err
	}

	if errPagination != nil {
		return responses, totalAktivitas, totalRows, errPagination
	}

	return responses, totalAktivitas, totalRows, err
}

// reportType unit kerja
func (repo BriefingRepository) BriefingReportByUkerFilter(request *models.BriefingFilterReportByUker) (responses []models.BriefingFilterReportByUkerResponse, totalRows int64, err error) {
	db := repo.db.DB
	query := ""
	sortQuery := ""
	var errPagination error

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

	if filter1 {
		fmt.Println("masuk filter1")
		if request.Sort == "DESC" {
			sortQuery = `PERCENTBRIEFING DESC, uker.REGION`
		} else {
			sortQuery = `PERCENTBRIEFING ASC, uker.REGION`
		}

		err = db.Table("dwh_branch uker").
			Select(`
						uker.REGION, 
						uker.RGDESC, 
						uker.MAINBR, 
						uker.MBDESC, 
						uker.BRANCH, 
						uker.BRDESC,
						COALESCE(total_briefing.TOTAL, 0) AS TOTALBRIEFING, 
						COALESCE(brc.TOTAL, 0) AS TOTALBRC,
						ROUND(COALESCE((total_briefing.TOTAL/brc.TOTAL)*100, 0), 0) AS PERCENTBRIEFING`).
			// Joins(
			// 	`LEFT JOIN (SELECT REGION, COUNT(id) AS TOTAL FROM briefing briefing_a
			// 			WHERE (created_at BETWEEN ? AND ?) GROUP BY REGION) total_briefing
			// 			ON total_briefing.REGION = uker.REGION
			// 			LEFT JOIN (SELECT kelolaan.REGION, COUNT(kelolaan.REGION) AS TOTAL FROM
			// 			(SELECT REGION, COUNT(pn) AS TOTAL FROM uker_kelolaan_user GROUP BY pn) kelolaan
			// 			GROUP BY kelolaan.REGION) brc ON brc.REGION = uker.REGION`, request.StartDate, lib.FixEndDate(request.EndDate)).
			Joins(
				`LEFT JOIN (SELECT REGION, COUNT(id) AS TOTAL FROM briefing briefing_a
						WHERE (created_at >= ? AND created_at <= ?) GROUP BY REGION) total_briefing
						ON total_briefing.REGION = uker.REGION
						LEFT JOIN (SELECT kelolaan.REGION, COUNT(kelolaan.REGION) AS TOTAL FROM
						(SELECT REGION, COUNT(pn) AS TOTAL FROM uker_kelolaan_user GROUP BY pn) kelolaan
						GROUP BY kelolaan.REGION) brc ON brc.REGION = uker.REGION`, request.StartDate, lib.FixEndDate(request.EndDate)).
			Where(
				"uker.BRUNIT = 'B' AND uker.MAINBR = uker.BRANCH AND (uker.BRDESC LIKE ? OR uker.BRDESC = ?)",
				"kanwil%", "Jkt KCK").
			Order(sortQuery).
			Limit(request.Limit).
			Offset(request.Offset).
			Find(&responses).Error

		errPagination = db.Table("dwh_branch").
			Where("BRUNIT = ? AND (BRDESC LIKE ? OR BRDESC = ?)", "B", "kanwil%", "Jkt KCK").
			Count(&totalRows).Error

	} else if filter2 {
		fmt.Println("masuk filter2")
		if request.Sort == "DESC" {
			sortQuery = `PERCENTBRIEFING DESC, uker.MAINBR`
		} else {
			sortQuery = `PERCENTBRIEFING ASC, uker.MAINBR`
		}

		err = db.Table("dwh_branch uker").
			Select(`uker.REGION, 
					uker.RGDESC, 
					uker.MAINBR, 
					uker.MBDESC, 
					uker.BRANCH, 
					uker.BRDESC, 
					COALESCE(total_briefing.TOTAL, 0) AS TOTALBRIEFING, 
					COALESCE(brc.TOTAL, 0) AS TOTALBRC, 
					ROUND(COALESCE((total_briefing.TOTAL / brc.TOTAL) * 100, 0), 0) AS PERCENTBRIEFING`).
			// Joins(`LEFT JOIN (
			// 		SELECT
			// 			REGION,
			// 			MAINBR,
			// 			COUNT(id) AS TOTAL
			// 		FROM briefing
			// 		WHERE
			// 			(created_at BETWEEN ? AND ?)
			// 		GROUP BY
			// 			MAINBR
			// 		) total_briefing ON total_briefing.REGION = uker.REGION AND total_briefing.MAINBR = uker.MAINBR`,
			// 	request.StartDate, lib.FixEndDate(request.EndDate)).
			Joins(`LEFT JOIN (
				SELECT 
					REGION, 
					MAINBR, 
					COUNT(id) AS TOTAL 
				FROM briefing 
				WHERE 
					(created_at >= ? AND created_at < ?)
				GROUP BY 
					MAINBR
				) total_briefing ON total_briefing.REGION = uker.REGION AND total_briefing.MAINBR = uker.MAINBR`,
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
			Find(&responses).Error

		errPagination = db.Table("dwh_branch").
			Select("COUNT(DISTINCT MAINBR) AS pagination").
			Where("REGION = ?", request.REGION).
			Scan(&totalRows).Error

	} else if filter3 {
		fmt.Println("masuk filter3")
		if request.Sort == "DESC" {
			sortQuery = `PERCENTBRIEFING DESC, uker.BRANCH`
		} else {
			sortQuery = `PERCENTBRIEFING ASC, uker.BRANCH`
		}

		mainbrs := strings.Split(request.MAINBR, ",")

		err = db.Table("dwh_branch uker").
			Select(`
				uker.REGION, 
				uker.RGDESC, 
				uker.MAINBR, 
				uker.MBDESC, 
				uker.BRANCH, 
				uker.BRDESC, 
				COALESCE(total_briefing.TOTAL,0) AS TOTALBRIEFING, 
				COALESCE(brc.TOTAL,0) AS TOTALBRC, 
				ROUND(COALESCE((total_briefing.TOTAL / brc.TOTAL)*100,0),0) AS PERCENTBRIEFING`).
			// Joins(`
			// 	LEFT JOIN (
			// 		SELECT
			// 			REGION,
			// 			MAINBR,
			// 			BRANCH,
			// 			COUNT(id) AS TOTAL
			// 		FROM briefing
			// 		WHERE created_at BETWEEN ? AND ?
			// 		GROUP BY BRANCH
			// 	) total_briefing ON total_briefing.REGION = uker.REGION AND total_briefing.MAINBR = uker.MAINBR AND total_briefing.BRANCH = uker.BRANCH`,
			// 	request.StartDate, lib.FixEndDate(request.EndDate)).
			Joins(`
				LEFT JOIN (
					SELECT 
						REGION, 
						MAINBR, 
						BRANCH, 
						COUNT(id) AS TOTAL 
					FROM briefing 
					WHERE (created_at >= ? AND  created_at <= ?) 
					GROUP BY BRANCH
				) total_briefing ON total_briefing.REGION = uker.REGION AND total_briefing.MAINBR = uker.MAINBR AND total_briefing.BRANCH = uker.BRANCH`,
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
			Scan(&responses).Error

		errPagination = db.Table("dwh_branch").
			Select("COUNT(DISTINCT BRANCH) AS pagination").
			Where("REGION = ? AND MAINBR in (?)", request.REGION, mainbrs).
			Scan(&totalRows).Error

	} else if filter4 { //add By Panji 24/11/2023
		fmt.Println("masuk filter4 UWU")
		if request.Sort == "DESC" {
			sortQuery = `PERCENTBRIEFING DESC, uker.BRANCH`
		} else {
			sortQuery = `PERCENTBRIEFING ASC, uker.BRANCH`
		}

		mainbrs := strings.Split(request.MAINBR, ",")
		branches := strings.Split(request.BRANCH, ",")

		err = db.Table("dwh_branch uker").
			Select(`
				uker.REGION, 
				uker.RGDESC, 
				uker.MAINBR, 
				uker.MBDESC, 
				uker.BRANCH, 
				uker.BRDESC, 
				COALESCE(total_briefing.TOTAL,0) AS TOTALBRIEFING, 
				COALESCE(brc.TOTAL,0) AS TOTALBRC, 
				ROUND(COALESCE((total_briefing.TOTAL / brc.TOTAL)*100,0),0) AS PERCENTBRIEFING`).
			Joins(`
				LEFT JOIN (
					SELECT 
						REGION, 
						MAINBR, 
						BRANCH, 
						COUNT(id) AS TOTAL 
					FROM briefing 
					WHERE (created_at >= ? AND created_at <= ?) 
					GROUP BY BRANCH
				) total_briefing ON total_briefing.REGION = uker.REGION AND total_briefing.MAINBR = uker.MAINBR AND total_briefing.BRANCH = uker.BRANCH`,
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
			Scan(&responses).Error

		errPagination = db.Table("dwh_branch").
			Select("COUNT(DISTINCT BRANCH) AS pagination").
			Where("REGION = ? AND MAINBR in (?) AND BRANCH in (?)", request.REGION, mainbrs, branches).
			Scan(&totalRows).Error
	}

	//QUERY PAGINATION

	if errPagination != nil {
		return responses, totalRows, err
	}

	repo.logger.Zap.Info("briefing-report-uker-query-activity-unknown", query)

	if err != nil {
		return responses, totalRows, err
	}

	return responses, totalRows, err
}

func (repo BriefingRepository) BriefingReportFilterComplete(request *models.BriefingFilterReport) (responses []models.BriefingFilterReportFinalResponse, totalRows int, err error) {
	db := repo.db.DB
	query := db
	queryPagination := db
	sortQuery := ""

	if request.Sort == "DESC" {
		sortQuery = `briefing_materis.created_at DESC, id`
	} else {
		sortQuery = `briefing_materis.created_at ASC, id`
	}

	query = db.Model(&responses).
		Select(`
				DISTINCT b.id as id, 
				b.created_at as date, 
				b.BRANCH, 
				b.BRDESC, 
				a.name as activity,
				p.product as product,
				briefing_materis.risk_issue_code as risk_issue`).
		Joins(`JOIN briefing b ON b.id = briefing_materis.briefing_id`).
		Joins(`JOIN activity a ON a.kode_activity = briefing_materis.activity_id`).
		Joins(`JOIN product p ON p.id = briefing_materis.product_id`).
		Where(`briefing_materis.activity_id = ?`, request.Activity).
		Where(`briefing_materis.product_id = ?`, request.Product).
		Where(`briefing_materis.judul_materi = ?`, request.Title).
		Where(`(briefing_materis.created_at >= ? AND briefing_materis.created_at <= ?)`, request.StartDate, lib.FixEndDate(request.EndDate))

	queryPagination = db.Table("briefing_materis").
		Select(`COUNT(*) as totalRows`).
		Joins(`JOIN briefing b ON b.id = briefing_materis.briefing_id`).
		Joins(`JOIN activity a ON a.kode_activity = briefing_materis.activity_id`).
		Joins(`JOIN product p ON p.id = briefing_materis.product_id`).
		Where(`briefing_materis.activity_id = ?`, request.Activity).
		Where(`briefing_materis.product_id = ?`, request.Product).
		Where(`briefing_materis.judul_materi = ?`, request.Title).
		Where(`(briefing_materis.created_at >= ? AND briefing_materis.created_at <= ?)`, request.StartDate, lib.FixEndDate(request.EndDate))

	if request.REGION != "all" {
		query = query.Joins("JOIN briefing ON briefing.id = briefing_materis.briefing_id")
		query = query.Where("briefing.REGION = ?", request.REGION)

		queryPagination = queryPagination.Joins("JOIN briefing ON briefing.id = briefing_materis.briefing_id")
		queryPagination = queryPagination.Where("briefing.REGION = ?", request.REGION)
	}

	if request.REGION != "all" && request.MAINBR != "all" {
		mainbrs := strings.Split(request.MAINBR, ",")

		query = query.Where("briefing.MAINBR in (?)", mainbrs)
		queryPagination = queryPagination.Where("briefing.MAINBR in (?)", mainbrs)
	}

	if request.REGION != "all" && request.MAINBR != "all" && request.BRANCH != "all" {
		branches := strings.Split(request.BRANCH, ",")
		query = query.Where("briefing.BRANCH in (?)", branches)
		queryPagination = queryPagination.Where("briefing.MAINBR in (?)", branches)
	}

	// if request.BRANCH != "all" {
	// 	branches := strings.Split(request.BRANCH, ",")
	// 	query = query.Where("briefing.BRANCH in (?)", branches)
	// 	queryPagination = queryPagination.Where("briefing.MAINBR in (?)", branches)
	// }

	query = query.Order(sortQuery).
		Limit(request.Limit).
		Offset(request.Offset).
		Find(&responses)

	queryPagination = queryPagination.Find(&totalRows)

	if err != nil {
		return responses, totalRows, err
	}

	return responses, totalRows, err
}

func (repo BriefingRepository) BriefingReportFilterByUkerComplete(request *models.BriefingFilterReportByUker) (responses []models.BriefingFilterReportFinalResponseNull, totalRows int, err error) {
	var rows *sql.Rows
	var errPagination error
	sortQuery := ""

	if request.Sort == "DESC" {
		sortQuery = `ORDER BY b.created_at DESC, id`
	} else {
		sortQuery = `ORDER BY b.created_at ASC, id`
	}

	query := `SELECT b.id as id, b.created_at as date, BRANCH, BRDESC, a.name as activity, p.product as product, bm.risk_issue_code as risk_issue
			FROM briefing_materis bm
			JOIN briefing b ON b.id = bm.briefing_id
			JOIN activity a ON a.kode_activity = bm.activity_id 
			JOIN product p ON p.id = bm.product_id 
			WHERE b.BRANCH = ?
			AND (b.created_at >= ? AND b.created_at <= ?)
			GROUP BY b.id
			` + sortQuery + `
			LIMIT ? OFFSET ?`

	queryPagination := `SELECT COUNT(*) as pagination FROM(
							SELECT bm.id
							FROM briefing_materis bm
							JOIN briefing b ON b.id = bm.briefing_id
							JOIN activity a ON a.kode_activity = bm.activity_id 
							JOIN product p ON p.id = bm.product_id 
							WHERE b.BRANCH = ?
							AND (b.created_at >= ? AND b.created_at <= ?)
							GROUP BY b.id
						) as Pagination`

	rows, err = repo.db.DB.Raw(query, request.BRANCH, request.StartDate, lib.FixEndDate(request.EndDate), strconv.Itoa(request.Limit), strconv.Itoa(request.Offset)).Rows()
	defer rows.Close()

	errPagination = repo.db.DB.Raw(queryPagination, request.BRANCH, request.StartDate, lib.FixEndDate(request.EndDate)).Scan(&totalRows).Error

	//QUERY PAGINATION
	repo.logger.Zap.Info("briefing-queryPagination-activity-unknown", queryPagination)

	if errPagination != nil {
		return responses, totalRows, err
	}

	repo.logger.Zap.Info("briefing-query-activity-unknown", query)

	repo.logger.Zap.Info("briefing-rows-activity-unknown", rows)
	if err != nil {
		return responses, totalRows, err
	}

	response := models.BriefingFilterReportFinalResponseNull{}
	for rows.Next() {
		_ = rows.Scan(
			&response.Id,
			&response.Date,
			&response.BRANCH,
			&response.BRDESC,
			&response.Activity,
			&response.Product,
			&response.RiskIssue,
		)
		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, totalRows, err
	}

	return responses, totalRows, err
}

func (repo BriefingRepository) BriefingReportDetail(request *models.BriefingReportDetailRequest) (responses models.BriefingReportDetailResponse, err error) {
	var rows1 *sql.Rows
	var rows2 *sql.Rows
	var err1 error
	var err2 error

	//get briefing detail
	query1 := `SELECT id, no_pelaporan, BRDESC as unit_kerja, jenis_peserta, jumlah_peserta FROM briefing WHERE id = ?`

	//get briefing materis
	query2 := `SELECT bm.id, a.name as activity, sa.name as sub_activity, p.product, bm.judul_materi, bm.rekomendasi_materi, bm.materi_tambahan
				FROM briefing_materis bm
				JOIN activity a ON a.kode_activity = bm.activity_id
				JOIN sub_activity sa ON sa.id = bm.sub_activity_id
				JOIN product p ON p.id = bm.product_id
				WHERE bm.briefing_id = ?`

	repo.logger.Zap.Info("briefing-query-detail", query1)
	repo.logger.Zap.Info("briefing-query-detail-materis", query2)

	rows1, err1 = repo.db.DB.Raw(query1, request.ID).Rows()
	rows2, err2 = repo.db.DB.Raw(query2, request.ID).Rows()
	defer rows1.Close()
	defer rows2.Close()

	repo.logger.Zap.Info("briefing-rows-detail", rows1)
	repo.logger.Zap.Info("briefing-rows-detail-materis", rows2)

	if err1 != nil {
		err = err1
		return responses, err
	}

	if err2 != nil {
		err = err2
		return responses, err
	}

	response1 := models.BriefingReportDetail{}
	for rows1.Next() {
		err = rows1.Scan(
			&response1.ID,
			&response1.NoPelaporan,
			&response1.UnitKerja,
			&response1.JenisPeserta,
			&response1.JumlahPeserta,
		)
		if err != nil {
			return responses, err
		}

		responses.BriefingDetail = response1
	}

	fmt.Println("")
	fmt.Println("========================== response1", response1)
	fmt.Println("")

	var responseMateris []models.BriefingReportDetailMateri
	response2 := models.BriefingReportDetailMateri{}
	for rows2.Next() {
		err = rows2.Scan(
			&response2.ID,
			&response2.Activity,
			&response2.SubActivity,
			&response2.Product,
			&response2.JudulMateri,
			&response2.RekomendasiMateri,
			&response2.MateriTambahan,
		)
		if err != nil {
			return responses, err
		}

		responseMateris = append(responseMateris, response2)
	}

	fmt.Println("")
	fmt.Println("========================== responseMateris", &responseMateris)
	fmt.Println("")

	responses.BriefingMateris = responseMateris

	if err1 = rows1.Err(); err1 != nil {
		return responses, err
	}

	if err2 = rows2.Err(); err2 != nil {
		return responses, err
	}

	return responses, err
}

func (repo BriefingRepository) BriefingReportMateriList(request *models.BriefingReportMateriRequest) (responses []models.BriefingDetailMateriResponseNull, err error) {
	var rows *sql.Rows

	fileId := strings.Split(request.ID, ",")
	query := `SELECT id, nama_lampiran, filename, path FROM risk_indicator_map_files rimf WHERE id IN ?`

	rows, err = repo.db.DB.Raw(query, fileId).Rows()

	// repo.logger.Zap.Info("briefing-query-activity-unknown", query)

	// repo.logger.Zap.Info("briefing-rows-activity-unknown", rows)

	if err != nil {
		return responses, err
	}

	response := models.BriefingDetailMateriResponseNull{}
	for rows.Next() {
		err = rows.Scan(
			&response.ID,
			&response.NamaLampiran,
			&response.Filename,
			&response.Path,
		)
		if err != nil {
			return responses, err
		}

		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, err
	}

	return responses, err
}

// BriefingReportList implements BriefingDefinition
func (repo BriefingRepository) BriefingReportList(request *models.BriefingReportListRequest) (responses []models.BriefingReportListResponse, totalRows int, err error) {
	db := repo.db.DB.Table("report_list_briefing rlb ")
	sorting := ""

	if request.Sort == "desc" {
		sorting = "rlb.id DESC"
	} else {
		sorting = "rlb.id ASC"
	}

	query := db.Select(`
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
	`).Order(sorting)

	if request.NoPelaporan != "" {
		query = query.Where("rlb.no_pelaporan = ?", request.NoPelaporan)
	}

	if request.REGION != "all" {
		query = query.Where("rlb.REGION = ?", request.REGION)
	}

	if request.MAINBR != "all" {
		mainbrs := strings.Split(request.MAINBR, ",")
		query = query.Where("rlb.MAINBR in (?)", mainbrs)
	}

	if request.BRANCH != "all" {
		branches := strings.Split(request.BRANCH, ",")
		query = query.Where("rlb.BRANCH in (?)", branches)
	}

	if request.ActivityID != "all" {
		query = query.Where("rlb.aktivity_id LIKE ?", fmt.Sprintf("%%%s%%", request.ActivityID))
	}

	if request.JudulMateri != "" {
		query = query.Where("rlb.judul_materi LIKE ?", fmt.Sprintf("%%%s%%", request.JudulMateri))
	}

	if request.Status != "" && request.Status != "Semua" {
		query = query.Where("rlb.status = ?", request.Status)
	}

	if request.StartDate != "" && lib.FixEndDate(request.EndDate) != "" {
		query = query.Where(`rlb.periode BETWEEN ? AND ?`, request.StartDate, lib.FixEndDate(request.EndDate))
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

// GetJudulMateri implements BriefingDefinition.
func (repo BriefingRepository) GetJudulMateri(id int64) (responses []models.JudulMateriBriefing, err error) {
	db := repo.db.DB.Table("briefing_materis").
		Select(`id`, `judul_materi`).Where(`briefing_id = ?`, id)

	err = db.Find(&responses).Error

	return responses, err
}

func (repo BriefingRepository) BriefingFrekuensiRpt(request *models.FrekuensiBriefingRequest) (responses []models.FrekuensiBriefingResponse, totalRows int, err error) {
	tableName := ""

	if request.Kegiatan == "briefing" {
		tableName = "briefing b"

		columnMappings := map[string]string{
			"activity":   "a.name AS 'aktivitas'",
			"product":    "p.product AS 'produk'",
			"risk_event": "bm.judul_materi AS 'risk_event'",
		}

		joinMappings := map[string]string{
			"activity": "JOIN activity a ON a.id = bm.activity_id",
			"product":  "JOIN product p ON p.id = bm.product_id",
		}

		groupMappings := map[string]string{
			"activity":   "bm.activity_id",
			"product":    "bm.product_id",
			"risk_event": "bm.risk_issue_code",
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

		query := repo.db.DB.Table(tableName).
			Select(strings.Join(selectParts, ", ")).
			Joins(`JOIN briefing_materis bm ON bm.briefing_id = b.id`).
			Joins(strings.Join(joins, " ")).
			Where("b.action IN (?, ?)", "Update", "Selesai").
			Group(strings.Join(groupBy, ", "))

		queryPaging := repo.db.DB.Table(tableName).
			Select(`Count(*)`).
			Joins(`JOIN briefing_materis bm ON bm.briefing_id=b.id`).
			Joins(strings.Join(joins, " ")).
			Where("b.action IN (?, ?)", "Update", "Selesai").
			Group(strings.Join(groupBy, ", "))

		if request.REGION != "all" && request.REGION != "" {
			query = query.Where("b.REGION = ?", request.REGION)
			queryPaging = queryPaging.Where("b.REGION = ?", request.REGION)
		}

		if request.MAINBR != "all" && request.MAINBR != "" {
			mainbrs := strings.Split(request.MAINBR, ",")
			query = query.Where("b.MAINBR in (?)", mainbrs)
			queryPaging = queryPaging.Where("b.MAINBR in (?)", mainbrs)
		}

		if request.BRANCH != "all" && request.BRANCH != "" {
			branches := strings.Split(request.BRANCH, ",")
			query = query.Where("b.BRANCH in (?)", branches)
			queryPaging = queryPaging.Where("b.BRANCH in (?)", branches)
		}

		if request.StartDate != "" && request.EndDate != "" {
			query = query.Order("jumlah DESC").Where("b.created_at BETWEEN ? AND ?", request.StartDate, lib.FixEndDate(request.EndDate))
			queryPaging = queryPaging.Where("b.created_at BETWEEN ? AND ?", request.StartDate, lib.FixEndDate((request.EndDate)))
		}

		var Count int64
		queryPaging = queryPaging.Count(&Count)

		totalRows = int(Count)

		query = query.Limit(request.Limit).Offset(request.Offset)

		responses = []models.FrekuensiBriefingResponse{}

		err = query.Scan(&responses).Error

		return responses, totalRows, err

	} else {
		return nil, 0, err
	}
}
