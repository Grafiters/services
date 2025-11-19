package verifikasi

import (
	"database/sql"
	"fmt"
	"math"
	"riskmanagement/lib"
	models "riskmanagement/models/verifikasi"
	"strconv"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"

	"strings"
)

type VerifikasiDefinition interface {
	WithTrx(trxHandle *gorm.DB) VerifikasiRepository
	GetAll() (responses []models.VerifikasiResponse, err error)
	GetListData() (responses []models.VerifikasiList, err error)
	GetOne(id int64) (responses models.VerifikasiResponse, err error)
	GetDataWithPagination(request *models.VerifikasiPagination) (responses []models.VerifikasiList, totalRows int, totalData int, err error)
	FilterVerifikasi(request *models.VerifikasiFilterRequest) (responses []models.VerifikasiList, totalRows int, totalData int, err error)
	Store(request *models.Verifikasi, tx *gorm.DB) (responses *models.Verifikasi, err error)
	Delete(request *models.VerifikasiUpdateDelete, include []string, tx *gorm.DB) (responses bool, err error)
	DeleteAnomaliData(id int64, tx *gorm.DB) (err error)
	DeleteLampiranVerifikasi(id int64, tx *gorm.DB) (err error)
	KonfirmSave(request *models.VerifikasiUpdateMaintain, include []string, tx *gorm.DB) (response bool, err error)
	UpdateAllVerifikasi(request *models.VerifikasiUpdateAll, include []string, tx *gorm.DB) (response bool, err error)
	GetNoPelaporan(request *models.NoPalaporanRequest) (responses []models.NoPelaporanNullResponse, err error)
	GetLastID() (responses []models.VerifikasiLastID, err error)
	FilterReport(request *models.VerifikasiFilterReport) (responses []models.VerifikasiReportResponseNull, totalRows int, totalData int, err error)

	VerifikasiReportFilter(request *models.VerifikasiFilterReportRequest) (responses []models.VerifikasiFilterReportResponsWithoutPercent, totalRows int64, err error)
	VerifikasiReportFilterComplete(request *models.VerifikasiFilterReportRequest) (responses []models.VerifikasiFilterReportCompleteResponse, totalRows int64, err error)

	VerifikasiReportDetail(request *models.VerifikasiReportDetailRequest) (responsesDetail models.VerifikasiReportDetailResponseWithoutDataAnomaliNull, err error)
	RiskControlByVerificationId(request *models.DataRiskControlRequest) (responses []models.DataRiskIndicatorResponseWithNoPercent, totalRows int64, err error)

	VerifikasiReportWithWeaknessOnlyFilter(request *models.VerifikasiFilterReportRequest) (responses []models.VerifikasiFilterReportWeaknessOnlyResponsWithoutPercent, totalRows int64, err error)
	VerifikasiReportWithNonWeaknessOnlyFilter(request *models.VerifikasiFilterReportRequest) (responses []models.VerifikasiFilterReportNonWeaknessOnlyResponsWithoutPercentNull, totalRows int64, err error)

	GetRiskIndicatorAsMateri(request *models.VerifikasiFilterReportRequest) (responses []models.GetRiskIndicatorAsMateriResponseNull, err error)

	VerificationReportByUkerFilter(request *models.VerificationFilterReportByUkerRequest) (responses []models.VerificationFilterReportByUkerResponseNull, SumData int64, totalRows int, err error)
	VerificationReportFilterByUkerComplete(request *models.VerificationFilterReportByUkerRequest) (responses []models.VerificationFilterByUkerReportCompleteResponse, totalRows int64, err error)

	VerifikasiReportByFraudIndicatorFilter(request *models.VerificationFilterReportByUkerRequest) (responses []models.VerificationFilterReportByFraudIndicatorResponse, totalRows int64, err error)
	VerificationReportFilterByFraudIndicatorComplete(request *models.VerificationFilterReportByUkerRequest) (responses []models.VerifikasiFilterReportCompleteResponse, totalRows int, err error)

	//add 23 Feb 2023 By Panji
	VerifikasiReportMateriList(request *models.VerifikasiMateriRequest) (responses []models.VerifikasiDetailMateriResponseNull, totalRows int, totalData int, err error)

	VerificationReportUkerByAllActivity(request *models.VerificationFilterReportByUkerRequest) (responses models.VerifikasiReportAllUker, totalRows int, err error)
	VerificationReportUkerByAllActivityComplete(request *models.VerificationFilterReportByUkerRequest) (responses []models.ResponsesAllActivityComplete, totalRows int, err error)
	VerificationReportUkerByAllActivityCompleteWithRiskIssue(request *models.VerificationFilterReportByUkerRequest) (responses []map[string]interface{}, totalRows int, err error)
	//
	VerifikasiReportList(request *models.VerifikasiReportListRequest) (responses []models.VerifikasiReportListResponse, totalRows int, err error)
	RptRakapitulasiBCV(request *models.RptRekapitulasiBCVRequest) (responses []models.RptRekapitulasiBCVResponse, totalRows int, err error)

	//RptRekomendasiRisk
	RptRekomendasiRiskFromVerifikasi(request *models.RptRekomendasiRiskRequest) (responses []models.RptRekomendasiRiskResponse, totalRows int, err error)
	RptRekomendasiRiskFromBriefing(request *models.RptRekomendasiRiskRequest) (responses []models.RptRekomendasiRiskResponse, totalRows int, err error)
	RptRekomendasiRiskFromCoaching(request *models.RptRekomendasiRiskRequest) (responses []models.RptRekomendasiRiskResponse, totalRows int, err error)

	RptRekomendasiRisk(request *models.RptRekomendasiRiskRequest) (responses []models.RptRekomendasiRiskResponse, totalRows int, err error)

	//ValidasiVerifikasi
	ValidasiVerifikasi(request *models.ValidasiVerifikasiRequest) (responses []models.ValidasiVerifikasiResponse, totalRows int, err error)
	ValidasiVerifikasiDetailData(request *models.VerifikasiReportDetailRequest) (responses models.ValidasiVerifikasiDetailResponse, err error)

	//RTL Indikasi Fraud
	GetRtlIndikasiFraud(request *models.ReqRtlIndikasiFraud) (responses models.RtlIndikasiFraudResponse, totalRows int, err error)

	// Batch 3
	GetRekomendasiTindakLanjut(request *models.RTLRequest) (responses []models.RTLResponses, err error)
	VerifikasiSummaryRpt(request *models.SummaryVerifikasiRequest) (responses []models.SummaryVerifikasiResponse, totalRows int, err error)
	VerifikasiFrekuensiRpt(request *models.FrekuensiVerifikasiRequest) (responses []models.FrekuensiVerifikasiResponse, totalRows int, err error)
}

type VerifikasiRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewVerfikasiRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) VerifikasiDefinition {
	return VerifikasiRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// Delete implements VerifikasiDefinition
func (verifikasi VerifikasiRepository) Delete(request *models.VerifikasiUpdateDelete, include []string, tx *gorm.DB) (responses bool, err error) {
	return true, tx.Save(&request).Error
}

// DeleteAnomaliData implements VerifikasiDefinition
func (verifikasi VerifikasiRepository) DeleteAnomaliData(id int64, tx *gorm.DB) (err error) {
	return tx.Where("id = ?", id).Delete(&models.VerifikasiAnomaliDataRequest{}).Error
}

// GetAll implements VerifikasiDefinition
func (verifikasi VerifikasiRepository) GetAll() (responses []models.VerifikasiResponse, err error) {
	return responses, verifikasi.db.DB.Find(&responses).Error
}

// GetListData implements VerifikasiDefinition
func (verifikasi VerifikasiRepository) GetListData() (responses []models.VerifikasiList, err error) {
	rows, err := verifikasi.db.DB.Raw(`
		SELECT 
			v.id ,
			v.no_pelaporan 'no_pelaporan',
			v.BRDESC 'unit_kerja',
			a.name 'aktifitas',
			CASE 
				WHEN v.status = "01a" AND v.action = "Draft" THEN "Draft"
				WHEN v.status = "02a" AND v.action = "Validasi RMC" THEN "Validasi RMC"
				WHEN v.status = "02b" AND v.action = "Reject RMC" THEN "Reject RMC"
				WHEN v.status = "03a" AND v.action = "Validasi ORD" THEN "Validasi ORD"
				WHEN v.status = "03b" AND v.action = "Reject ORD" THEN "Reject ORD"
				WHEN v.status = "04a" AND v.action = "Selesai" THEN "Selesai"
				WHEN v.status = "01b" THEN "Draft"
			END 'status_verif'
		FROM verifikasi v 
		JOIN activity a ON v.activity_id = a.id 
		WHERE v.deleted != 1
		GROUP BY v.id 
	`).Rows()

	defer rows.Close()
	var verif models.VerifikasiList
	for rows.Next() {
		verifikasi.db.DB.ScanRows(rows, &verif)
		responses = append(responses, verif)
	}

	return responses, err
}

// GetOne implements VerifikasiDefinition
func (verifikasi VerifikasiRepository) GetOne(id int64) (responses models.VerifikasiResponse, err error) {
	err = verifikasi.db.DB.Raw(`
		SELECT 
			verif.*
		FROM verifikasi verif 
		WHERE verif.id = ?`, id).Find(&responses).Error

	if err != nil {
		verifikasi.logger.Zap.Error(err)
		return responses, err
	}
	return responses, err
}

// Store implements VerifikasiDefinition
func (verifikasi VerifikasiRepository) Store(request *models.Verifikasi, tx *gorm.DB) (responses *models.Verifikasi, err error) {
	return request, tx.Save(&request).Error
}

// WithTrx implements VerifikasiDefinition
func (verifikasi VerifikasiRepository) WithTrx(trxHandle *gorm.DB) VerifikasiRepository {
	if trxHandle == nil {
		verifikasi.logger.Zap.Error("transaction Database not found in gin context.")
		return verifikasi
	}

	verifikasi.db.DB = trxHandle
	return verifikasi
}

// DeleteLampiranVerifikasi implements VerifikasiDefinition
func (verifikasi VerifikasiRepository) DeleteLampiranVerifikasi(id int64, tx *gorm.DB) (err error) {
	return tx.Where("id = ?", id).Delete(&models.VerifikasiFilesRequest{}).Error
}

// KonfirmSave implements VerifikasiDefinition
func (verifikasi VerifikasiRepository) KonfirmSave(request *models.VerifikasiUpdateMaintain, include []string, tx *gorm.DB) (response bool, err error) {
	return true, tx.Save(&request).Error
}

// UpdateAllVerifikasi implements VerifikasiDefinition
func (verifikasi VerifikasiRepository) UpdateAllVerifikasi(request *models.VerifikasiUpdateAll, include []string, tx *gorm.DB) (response bool, err error) {
	return true, tx.Save(&request).Error
}

// FilterVerifikasi implements VerifikasiDefinition
func (verifikasi VerifikasiRepository) FilterVerifikasi(request *models.VerifikasiFilterRequest) (responses []models.VerifikasiList, totalRows int, totalData int, err error) {
	branches := strings.Split(request.Branches, ",")

	db := verifikasi.db.DB

	queryBuilder := db.Model(&responses).
		Select(`
			verifikasi.id 'id',
			verifikasi.no_pelaporan 'no_pelaporan',
			verifikasi.BRDESC 'unit_kerja',
			act.name 'aktifitas',
			CASE
				WHEN verifikasi.indikasi_fraud = 1 THEN 'Ya'
				ELSE 'Tidak'
			END 'indikasi_fraud',
			CASE 
				WHEN verifikasi.status = "01a" AND lower(verifikasi.action) = "Draft" THEN "Draft"
				WHEN verifikasi.status = "02a" AND lower(verifikasi.action) = "validasi rmc" THEN "Validasi RMC"
				WHEN verifikasi.status = "02b" AND lower(verifikasi.action) = "reject rmc" THEN "Reject RMC"
				WHEN verifikasi.status = "03a" AND lower(verifikasi.action) = "validasi ord" THEN "Validasi ORD"
				WHEN verifikasi.status = "03b" AND lower(verifikasi.action) = "reject ord" THEN "Reject ORD"
				WHEN verifikasi.status = "04a" AND lower(verifikasi.action) = "Selesai" THEN "Selesai"
				WHEN verifikasi.status = "01b" THEN "Draft"
			END 'status_verif',
			CASE
				WHEN verifikasi.perbaikan = 1 THEN
					(
						SELECT 
							GROUP_CONCAT(
								CASE
									WHEN vptl.status = 0 THEN 'Backlog'
									WHEN vptl.status = 1 THEN 'On Progress'
									WHEN vptl.status = 2 Then 'Selesai'
								END
							) 
						FROM verifikasi_pic_tindak_lanjut vptl WHERE vptl.verifikasi_id = verifikasi.id
					)
				ELSE '-'
			END 'status_rtl',
			verifikasi.action_indikasi_fraud 'status_fraud'
		`).
		Joins("LEFT JOIN activity act on act.kode_activity = verifikasi.activity_id").
		// Joins("LEFT JOIN verifikasi_pic_tindak_lanjut vptl2 ON verifikasi.id = vptl2.verifikasi_id").
		Where("verifikasi.deleted != 1").
		// Where(`v.BRANCH in (?)`, branches).
		Order("verifikasi.id DESC")

	// if request.Branches != "" {
	// 	branches := strings.Split(request.Branches, ",")
	// 	db = db.Where(`v.BRANCH in (?)`, branches)
	// }

	if request.Kostl != "" {
		queryBuilder = queryBuilder.Where(`verifikasi.maker_id = ?`, request.Pernr)
	} else {
		queryBuilder = queryBuilder.Where(`verifikasi.BRANCH in (?)`, branches)
	}

	if request.NoPelaporan != "" {
		queryBuilder = queryBuilder.Where("verifikasi.no_pelaporan = ?", request.NoPelaporan)
	}

	if request.UnitKerja != "" {
		// queryBuilder = queryBuilder.Where("verifikasi.BRDESC like ?", fmt.Sprintf("%%%s%%", request.UnitKerja))
		queryBuilder = queryBuilder.Where("verifikasi.BRANCH = ?", request.UnitKerja)

	}

	if request.ActivityID != "" {
		queryBuilder = queryBuilder.Where("verifikasi.activity_id = ?", request.ActivityID)
	}

	if request.RiskIssueID != "" {
		queryBuilder = queryBuilder.Where("verifikasi.risk_issue_id = ?", request.RiskIssueID)
	}

	if request.Status != "" && request.Status != "Semua" && request.Status != "Selesai" {
		queryBuilder = queryBuilder.Where("verifikasi.action = ?", request.Status)
	}

	if request.Status == "Selesai" {
		queryBuilder = queryBuilder.Where("verifikasi.action IN(?, ?)", "Update", "Selesai")
	}

	if request.TglAwal != "" && request.TglAkhir != "" {
		queryBuilder = queryBuilder.Where("(verifikasi.created_at >= ? AND verifikasi.created_at <= ?)", request.TglAwal, lib.FixEndDate(request.TglAkhir))
	}

	if request.IndikasiFraud != "" {
		queryBuilder = queryBuilder.Where("verifikasi.indikasi_fraud = ?", request.IndikasiFraud)
	}

	if request.StatusRtl != "" {
		queryBuilder = queryBuilder.Where("vptl2.status = ?", request.StatusRtl)
	}

	var count int64
	queryBuilder.
		Group("verifikasi.id").
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

// GetDataWithPagination implements VerifikasiDefinition
func (verifikasi VerifikasiRepository) GetDataWithPagination(request *models.VerifikasiPagination) (responses []models.VerifikasiList, totalRows int, totalData int, err error) {
	branches := strings.Split(request.Branches, ",")

	db := verifikasi.db.DB.Table("verifikasi v")

	db = db.Select(`
				v.id ,
				v.no_pelaporan 'no_pelaporan',
				v.BRDESC 'unit_kerja',
				a.name 'aktifitas',
				CASE
					WHEN v.indikasi_fraud = 1 THEN 'Ya'
					ELSE 'Tidak'
				END 'indikasi_fraud',	
				CASE 
					WHEN v.status = "01a" AND lower(v.action) = "Draft" THEN "Draft"
					WHEN v.status = "02a" AND lower(v.action) = "validasi rmc" THEN "Validasi RMC"
					WHEN v.status = "02b" AND lower(v.action) = "reject rmc" THEN "Reject RMC"
					WHEN v.status = "03a" AND lower(v.action) = "validasi ord" THEN "Validasi ORD"
					WHEN v.status = "03b" AND lower(v.action) = "reject ord" THEN "Reject ORD"
					WHEN v.status = "04a" AND lower(v.action) = "Selesai" THEN "Selesai"
					WHEN v.status = "01b" THEN "Draft"
				END 'status_verif',
				CASE
					WHEN v.perbaikan = 1 THEN
						(
							SELECT 
								GROUP_CONCAT(
									CASE
										WHEN vptl.status = 0 THEN 'Backlog'
										WHEN vptl.status = 1 THEN 'On Progress'
										WHEN vptl.status = 2 Then 'Selesai'
									END
								) 
							FROM verifikasi_pic_tindak_lanjut vptl WHERE vptl.verifikasi_id = v.id
						)
					ELSE '-'
				END 'status_rtl',
				v.action_indikasi_fraud 'status_fraud'`).
		Joins("LEFT JOIN activity a on a.kode_activity = v.activity_id").
		Where(`v.deleted != 1`).
		// Where(`v.BRANCH in (?)`, branches).
		Order("v.id DESC")

	// if request.Branches != "" {
	// 	branches := strings.Split(request.Branches, ",")
	// 	db = db.Where(`v.BRANCH in (?)`, branches)
	// }

	if request.Kostl != "" {
		db = db.Where(`v.maker_id = ?`, request.Pernr)
	} else {
		db = db.Where(`v.BRANCH in (?)`, branches)
	}

	var count int64
	db.Group(`v.id`).Count(&count)

	totalRows = int(count)

	err = db.Limit(request.Limit).Offset(request.Offset).Find(&responses).Error

	totalPages := int(math.Ceil(float64(totalRows) / float64(request.Limit)))

	return responses, totalPages, totalRows, err
}

// GetNoPelaporan implements VerifikasiDefinition
func (verifikasi VerifikasiRepository) GetNoPelaporan(request *models.NoPalaporanRequest) (responses []models.NoPelaporanNullResponse, err error) {
	kode := "VER-"
	today := lib.GetTimeNow("date2")

	if request.ORGEH != "" {
		kode += request.ORGEH + "-" + today
	}

	query := `SELECT RIGHT(CONCAT("0000",(count(*) + 1)), 4) 'no_pelaporan' FROM verifikasi WHERE no_pelaporan like ?`

	verifikasi.logger.Zap.Info(query)
	rows, err := verifikasi.dbRaw.DB.Query(query, fmt.Sprintf("%%%s%%", kode))
	defer rows.Close()

	verifikasi.logger.Zap.Info("rows ", rows)
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

// GetLastID implements VerifikasiDefinition
func (verifikasi VerifikasiRepository) GetLastID() (responses []models.VerifikasiLastID, err error) {
	query := "SELECT id FROM verifikasi ORDER BY id DESC LIMIT 1"
	verifikasi.logger.Zap.Info(query)
	rows, err := verifikasi.dbRaw.DB.Query(query)
	defer rows.Close()

	verifikasi.logger.Zap.Info("rows ", rows)
	if err != nil {
		return responses, err
	}

	response := models.VerifikasiLastID{}
	for rows.Next() {
		_ = rows.Scan(
			&response.ID,
		)

		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, err
	}

	return responses, err
}

// FilterReport implements VerifikasiDefinition
func (verifikasi VerifikasiRepository) FilterReport(request *models.VerifikasiFilterReport) (responses []models.VerifikasiReportResponseNull, totalRows int, totalData int, err error) {
	where := " WHERE verif.deleted != 1"
	whereCount := " WHERE verif.deleted != 1"
	sorting := ""

	if request.UnitKerja != "" {
		where += " AND verif.unit_kerja = '" + request.UnitKerja + "'"
		whereCount += " AND verif.unit_kerja = '" + request.UnitKerja + "'"
	}

	if request.ActivityID != "" {
		where += " AND verif.activity_id = '" + request.ActivityID + "'"
		whereCount += " AND verif.activity_id = '" + request.ActivityID + "'"
	}

	if request.RiskIssueID != "" {
		where += " AND verif.risk_issue_id = '" + request.RiskIssueID + "'"
		whereCount += " AND verif.risk_issue_id = '" + request.RiskIssueID + "'"
	}

	if request.Perbaikan != "" {
		where += " AND verif.perbaikan = '" + request.Perbaikan + "'"
		whereCount += " AND verif.perbaikan = '" + request.Perbaikan + "'"
	}

	if request.IndikasiFraud != "" {
		where += " AND verif.indikasi_fraud = '" + request.IndikasiFraud + "'"
		whereCount += " AND verif.indikasi_fraud = '" + request.IndikasiFraud + "'"
	}

	if request.Status == "Selesai" {
		where += " AND verif.status = '02b' AND (verif.action = 'Update' OR verif.action = 'Selesai')"
		whereCount += " AND verif.status = '02b' AND (verif.action = 'Update' OR verif.action = 'Selesai')"
	}

	if request.Status != "" && request.Status != "Semua" && request.Status != "Selesai" {
		where += " AND verif.action = '" + request.Status + "'"
		whereCount += " AND verif.action = '" + request.Status + "'"
	}

	if request.TglAwal != "" && request.TglAkhir != "" {
		where += " AND CAST(created_at as date) >= '" + request.TglAwal + "' AND CAST(created_at as date) <= '" + lib.FixEndDate(request.TglAkhir) + "'"
		whereCount += " AND CAST(created_at as date) >= '" + request.TglAwal + "' AND CAST(created_at as date) <= '" + lib.FixEndDate(request.TglAkhir) + "'"
	}

	if request.Sort != "" {
		sorting += request.Sort
	}

	query := `SELECT 
				verif.id 'id',
				verif.created_at 'tanggal',
				verif.BRDESC 'unit_kerja',
				act.name 'aktifitas',
				prod.product 'produk',
				issue.risk_issue 'risk_issue',
				indicator.risk_indicator 'judul_materi'
			FROM verifikasi verif
			LEFT JOIN activity act on act.id = verif.activity_id
			LEFT JOIN product prod on prod.id = verif.product_id
			LEFT JOIN risk_issue issue on issue.id = verif.risk_issue_id
			LEFT JOIN risk_indicator indicator on indicator.id = verif.risk_indicator_id
			` + where + ` ORDER BY verif.created_at ` + sorting + ` LIMIT ? OFFSET ?`

	verifikasi.logger.Zap.Info(query)
	rows, err := verifikasi.dbRaw.DB.Query(query, request.Limit, request.Offset)
	defer rows.Close()

	verifikasi.logger.Zap.Info("rows ", rows)
	if err != nil {
		return responses, totalRows, totalData, err
	}

	response := models.VerifikasiReportResponseNull{}
	for rows.Next() {
		_ = rows.Scan(
			&response.ID,
			&response.Tanggal,
			&response.KodeBranch,
			&response.Aktifitas,
			&response.Produk,
			&response.RiskIssue,
			&response.JudulMateri,
		)
		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, totalRows, totalData, err
	}

	paginateQuery := `SELECT 
					count(*)
				FROM verifikasi verif
				LEFT JOIN activity act on act.id = verif.activity_id
				LEFT JOIN product prod on prod.id = verif.product_id
				LEFT JOIN risk_issue issue on issue.id = verif.risk_issue_id
				LEFT JOIN risk_indicator indicator on indicator.id = verif.risk_indicator_id` + whereCount + ` ORDER BY verif.created_at ` + sorting

	err = verifikasi.dbRaw.DB.QueryRow(paginateQuery).Scan(&totalRows)

	result := float64(totalRows) / float64(request.Limit)
	resultFinal := int(math.Ceil(result))

	return responses, resultFinal, totalRows, err
}

// filter report -- zan
// FraudIndication must be '1' or '0' or '1','0'

// by aktivitas
func (repo VerifikasiRepository) VerifikasiReportFilter(request *models.VerifikasiFilterReportRequest) (responses []models.VerifikasiFilterReportResponsWithoutPercent, totalRows int64, err error) {
	db := repo.db.DB
	query := db
	queryPagination := db
	sortQuery := ""

	subquery1 := db
	subquery2 := db

	var errPagination error

	//bank wide
	filter1 := request.ReportType == "aktivitas" &&
		request.Activity == "all" &&
		request.Product == "all" &&
		(request.RiskIssue == "all" || request.RiskIssue == "Semua") &&
		(request.Title == "all" || request.Title == "Semua")

	filter2 := request.ReportType == "aktivitas" &&
		request.Activity != "all" &&
		request.Product == "all" &&
		(request.RiskIssue == "all" || request.RiskIssue == "Semua") &&
		(request.Title == "all" || request.Title == "Semua")

	filter3 := request.ReportType == "aktivitas" &&
		request.Activity != "all" &&
		request.Product != "all" &&
		(request.RiskIssue == "all" || request.RiskIssue == "Semua") &&
		(request.Title == "all" || request.Title == "Semua")

	filter4 := request.ReportType == "aktivitas" &&
		request.Activity != "all" &&
		request.Product != "all" &&
		(request.RiskIssue != "all" && request.RiskIssue != "Semua") &&
		(request.Title == "all" || request.Title == "Semua")

	if filter1 {
		fmt.Println("====== query 1")
		fraudIndication := strings.Split(request.FraudIndication, ",")

		if request.Sort == "DESC" {
			sortQuery = `grand_total DESC, id`
		} else {
			sortQuery = `grand_total ASC, id`
		}

		// subquery
		subquery1 = db.Table("verifikasi").
			Select("activity_id, perbaikan, COUNT(*) as total").
			// Joins("JOIN verifikasi_risk_control vrc ON vrc.verifikasi_id = verifikasi.id").
			Where("perbaikan = ?", "1").
			Where("indikasi_fraud IN (?)", fraudIndication).
			Where("deleted != 1")

		subquery2 = db.Table("verifikasi").
			Select("activity_id, perbaikan, COUNT(*) as total").
			// Joins("JOIN verifikasi_risk_control vrc ON vrc.verifikasi_id = verifikasi.id").
			Where("perbaikan = ?", "0").
			Where("indikasi_fraud IN (?)", fraudIndication).
			Where("deleted != 1")

		//// uker on subquery
		if request.REGION != "all" {
			subquery1 = subquery1.Where("verifikasi.REGION = ?", request.REGION)
			subquery2 = subquery2.Where("verifikasi.REGION = ?", request.REGION)
		}

		if request.REGION != "all" && request.MAINBR != "all" {
			mainbrs := strings.Split(request.MAINBR, ",")

			subquery1 = subquery1.Where("verifikasi.MAINBR in (?)", mainbrs)
			subquery2 = subquery2.Where("verifikasi.MAINBR in (?)", mainbrs)
		}

		if request.REGION != "all" && request.MAINBR != "all" && request.BRANCH != "all" {
			branches := strings.Split(request.BRANCH, ",")
			if len(branches) > 1 {
				fmt.Println("Masuk 1")
				subquery1 = subquery1.Where("verifikasi.BRANCH in (?)", branches)
				subquery2 = subquery2.Where("verifikasi.BRANCH in (?)", branches)
			} else {
				fmt.Println("Masuk 2")
				subquery1 = subquery1.Where("verifikasi.BRANCH = ?", request.BRANCH)
				subquery2 = subquery2.Where("verifikasi.BRANCH = ?", request.BRANCH)
			}

		}

		//// status on subquery
		/*
			if request.Status == "draft" {
				subquery1 = subquery1.Where("(status = '01a' AND action = 'draft')")
				subquery2 = subquery2.Where("(status = '01a' AND action = 'draft')")
			} else if request.Status == "selesai" {
				subquery1 = subquery1.Where("((status = '02a' AND action = 'selesai') OR (status = '02b' AND (action = 'update' OR action = 'selesai')))")
				subquery2 = subquery2.Where("((status = '02a' AND action = 'selesai') OR (status = '02b' AND (action = 'update' OR action = 'selesai')))")
			}
		*/

		//edited By Panji #add status veladasi RMC & ORD
		if request.Status == "draft" {
			subquery1 = subquery1.Where("(status = '01a' AND lower(action) = 'draft') OR (status = '01b')")
			subquery2 = subquery2.Where("(status = '01a' AND lower(action) = 'draft') OR (status = '01b')")
		} else if request.Status == "selesai" {
			subquery1 = subquery1.Where("(status = '04a' AND lower(action) = 'selesai')")
			subquery2 = subquery2.Where("(status = '04a' AND lower(action) = 'selesai')")
		}

		// subquery1 = subquery1.Where("created_at BETWEEN ? AND ?", request.StartDate, lib.FixEndDate(request.EndDate)).
		subquery1 = subquery1.Where("(created_at >= ? AND created_at <= ?)", request.StartDate, lib.FixEndDate(lib.FixEndDate(request.EndDate))).
			Group("activity_id")

		subquery2 = subquery2.Where("(created_at >= ? AND created_at <= ?)", request.StartDate, lib.FixEndDate(request.EndDate)).
			Group("activity_id")

		// query
		query = db.Table("activity").
			Select(`
				id, 
				kode_activity as code, 
				name, 
				COALESCE(weakness.total, 0) as total_weakness, 
				COALESCE(non_weakness.total, 0) as total_non_weakness, 
				(COALESCE(weakness.total, 0) + COALESCE(non_weakness.total, 0)) as grand_total`).
			Joins(`
				LEFT JOIN (?) weakness ON weakness.activity_id = activity.id`, subquery1).
			Joins(`
				LEFT JOIN (?) non_weakness ON non_weakness.activity_id = activity.id`, subquery2).
			Order(sortQuery).
			Limit(request.Limit).
			Offset(request.Offset).
			Find(&responses)
		err = query.Error

		// query pagination
		queryPagination = db.Table("activity").Count(&totalRows)
		errPagination = queryPagination.Error

	} else if filter2 {
		fmt.Println("====== query 2")
		fraudIndication := strings.Split(request.FraudIndication, ",")

		if request.Sort == "DESC" {
			sortQuery = `grand_total DESC, id`
		} else {
			sortQuery = `grand_total ASC, id`
		}

		// subquery
		subquery1 = db.Table("verifikasi").
			Select("activity_id, product_id, perbaikan, COUNT(*) as total").
			// Joins("JOIN verifikasi_risk_control vrc ON vrc.verifikasi_id = verifikasi.id").
			Where("perbaikan = ?", "1").
			Where("indikasi_fraud IN (?)", fraudIndication).
			Where("deleted != 1")

		subquery2 = db.Table("verifikasi").
			Select("activity_id, product_id, perbaikan, COUNT(*) as total").
			// Joins("JOIN verifikasi_risk_control vrc ON vrc.verifikasi_id = verifikasi.id").
			Where("perbaikan = ?", "0").
			Where("indikasi_fraud IN (?)", fraudIndication).
			Where("deleted != 1")

		//// uker on subquery
		if request.REGION != "all" {
			subquery1 = subquery1.Where("verifikasi.REGION = ?", request.REGION)
			subquery2 = subquery2.Where("verifikasi.REGION = ?", request.REGION)
		}

		if request.REGION != "all" && request.MAINBR != "all" {
			mainbrs := strings.Split(request.MAINBR, ",")

			subquery1 = subquery1.Where("verifikasi.MAINBR in (?)", mainbrs)
			subquery2 = subquery2.Where("verifikasi.MAINBR in (?)", mainbrs)
		}

		if request.REGION != "all" && request.MAINBR != "all" && request.BRANCH != "all" {
			branches := strings.Split(request.BRANCH, ",")
			if len(branches) > 1 {
				fmt.Println("Masuk 1")
				subquery1 = subquery1.Where("verifikasi.BRANCH in (?)", branches)
				subquery2 = subquery2.Where("verifikasi.BRANCH in (?)", branches)
			} else {
				fmt.Println("Masuk 2")
				subquery1 = subquery1.Where("verifikasi.BRANCH = ?", request.BRANCH)
				subquery2 = subquery2.Where("verifikasi.BRANCH = ?", request.BRANCH)
			}

		}

		//// status on subquery
		// if request.Status == "draft" {
		// 	subquery1 = subquery1.Where("(status = '01a' AND action = 'draft')")
		// 	subquery2 = subquery2.Where("(status = '01a' AND action = 'draft')")
		// } else if request.Status == "selesai" {
		// 	subquery1 = subquery1.Where("((status = '02a' AND action = 'selesai') OR (status = '02b' AND (action = 'update' OR action = 'selesai')))")
		// 	subquery2 = subquery2.Where("((status = '02a' AND action = 'selesai') OR (status = '02b' AND (action = 'update' OR action = 'selesai')))")
		// }

		//edited By Panji #add status veladasi RMC & ORD
		if request.Status == "draft" {
			subquery1 = subquery1.Where("(status = '01a' AND lower(action) = 'draft') OR (status = '01b')")
			subquery2 = subquery2.Where("(status = '01a' AND lower(action) = 'draft') OR (status = '01b')")
		} else if request.Status == "selesai" {
			subquery1 = subquery1.Where("(status = '04a' AND lower(action) = 'selesai')")
			subquery2 = subquery2.Where("(status = '04a' AND lower(action) = 'selesai')")
		}

		subquery1 = subquery1.Where("(created_at >= ? AND created_at <= ?)", request.StartDate, lib.FixEndDate(request.EndDate)).
			Group("activity_id, product_id")

		subquery2 = subquery2.Where("(created_at >= ? AND created_at <= ?)", request.StartDate, lib.FixEndDate(request.EndDate)).
			Group("activity_id, product_id")

		// query
		query = db.Table("product p").
			Select(`
				id, 
				kode_product as code, 
				product as name, 
				COALESCE(weakness.total, 0) as total_weakness,
				COALESCE(non_weakness.total, 0) as total_non_weakness,
				(COALESCE(weakness.total, 0) + COALESCE(non_weakness.total, 0)) as grand_total`).
			Joins(`LEFT JOIN (?) weakness ON weakness.activity_id = ? AND weakness.product_id = p.id`, subquery1, request.Activity).
			Joins(`LEFT JOIN (?) non_weakness ON non_weakness.activity_id = ? AND non_weakness.product_id = p.id`, subquery2, request.Activity).
			Where(`p.activity_id = ?`, request.Activity).
			Order(sortQuery).
			Limit(request.Limit).
			Offset(request.Offset).
			Find(&responses)
		err = query.Error

		// query pagination
		queryPagination = db.Table("product").Where("activity_id = ?", request.Activity).Count(&totalRows)
		errPagination = queryPagination.Error
	} else if filter3 {
		fmt.Println("====== query 3")
		fraudIndication := strings.Split(request.FraudIndication, ",")

		if request.Sort == "DESC" {
			sortQuery = `grand_total DESC, id`
		} else {
			sortQuery = `grand_total ASC, id`
		}

		// subquery
		subquery1 = db.Table("verifikasi").
			Select("activity_id, product_id, risk_issue_id, risk_issue, perbaikan, COUNT(*) as total").
			// Joins("JOIN verifikasi_risk_control vrc ON vrc.verifikasi_id = verifikasi.id").
			Where("perbaikan = ?", "1").
			Where("indikasi_fraud IN (?)", fraudIndication).
			Where("deleted != 1").
			Where("activity_id = ?", request.Activity).
			Where("product_id = ?", request.Product)

		subquery2 = db.Table("verifikasi").
			Select("activity_id, product_id, risk_issue_id, risk_issue, perbaikan, COUNT(*) as total").
			// Joins("JOIN verifikasi_risk_control vrc ON vrc.verifikasi_id = verifikasi.id").
			Where("perbaikan = ?", "0").
			Where("indikasi_fraud IN (?)", fraudIndication).
			Where("deleted != 1").
			Where("activity_id = ?", request.Activity).
			Where("product_id = ?", request.Product)

		//// uker on subquery
		if request.REGION != "all" {
			subquery1 = subquery1.Where("verifikasi.REGION = ?", request.REGION)
			subquery2 = subquery2.Where("verifikasi.REGION = ?", request.REGION)
		}

		if request.REGION != "all" && request.MAINBR != "all" {
			mainbrs := strings.Split(request.MAINBR, ",")

			subquery1 = subquery1.Where("verifikasi.MAINBR in (?)", mainbrs)
			subquery2 = subquery2.Where("verifikasi.MAINBR in (?)", mainbrs)
		}

		if request.REGION != "all" && request.MAINBR != "all" && request.BRANCH != "all" {
			branches := strings.Split(request.BRANCH, ",")
			if len(branches) > 1 {
				fmt.Println("Masuk 1")
				subquery1 = subquery1.Where("verifikasi.BRANCH in (?)", branches)
				subquery2 = subquery2.Where("verifikasi.BRANCH in (?)", branches)
			} else {
				fmt.Println("Masuk 2")
				subquery1 = subquery1.Where("verifikasi.BRANCH = ?", request.BRANCH)
				subquery2 = subquery2.Where("verifikasi.BRANCH = ?", request.BRANCH)
			}

		}

		//// status on subquery
		// if request.Status == "draft" {
		// 	subquery1 = subquery1.Where("(status = '01a' AND action = 'draft')")
		// 	subquery2 = subquery2.Where("(status = '01a' AND action = 'draft')")
		// } else if request.Status == "selesai" {
		// 	subquery1 = subquery1.Where("((status = '02a' AND action = 'selesai') OR (status = '02b' AND (action = 'update' OR action = 'selesai')))")
		// 	subquery2 = subquery2.Where("((status = '02a' AND action = 'selesai') OR (status = '02b' AND (action = 'update' OR action = 'selesai')))")
		// }

		//edited By Panji #add status veladasi RMC & ORD
		if request.Status == "draft" {
			subquery1 = subquery1.Where("(status = '01a' AND lower(action) = 'draft') OR (status = '01b')")
			subquery2 = subquery2.Where("(status = '01a' AND lower(action) = 'draft') OR (status = '01b')")
		} else if request.Status == "selesai" {
			subquery1 = subquery1.Where("(status = '04a' AND lower(action) = 'selesai')")
			subquery2 = subquery2.Where("(status = '04a' AND lower(action) = 'selesai')")
		}

		subquery1 = subquery1.Where("(created_at >= ? AND created_at <= ?)", request.StartDate, lib.FixEndDate(request.EndDate)).
			Group("risk_issue")

		subquery2 = subquery2.Where("(created_at >= ? AND created_at <= ?)", request.StartDate, lib.FixEndDate(request.EndDate)).
			Group("risk_issue")

		// query
		query = db.Table("verifikasi v").
			Select(`
				v.risk_issue_id as id,
				v.product_id as code,
				v.risk_issue as name,
				COALESCE(weakness.total, 0) as total_weakness,
				COALESCE(non_weakness.total, 0) as total_non_weakness,
				(COALESCE(weakness.total, 0) + COALESCE(non_weakness.total, 0)) as grand_total`).
			Joins(`
				LEFT JOIN (?) weakness 
					ON weakness.activity_id = v.activity_id 
					AND weakness.product_id = v.product_id 
					AND weakness.risk_issue = v.risk_issue`, subquery1).
			Joins(`
				LEFT JOIN (?) non_weakness 
					ON non_weakness.activity_id = v.activity_id 
					AND non_weakness.product_id = v.product_id 
					AND non_weakness.risk_issue = v.risk_issue`, subquery2).
			Where(`v.activity_id = ?`, request.Activity).
			Where(`v.product_id = ?`, request.Product).
			Group(`name`).
			Order(sortQuery).
			Limit(request.Limit).
			Offset(request.Offset).
			Find(&responses)
		err = query.Error

		// query pagination
		queryPagination = db.Table("verifikasi").
			Select(`SUM(
				COUNT(
					DISTINCT risk_issue
				)
			) OVER() as pagination`).
			Where(`activity_id = ?`, request.Activity).
			Where(`product_id = ?`, request.Product).
			Scan(&totalRows)
		errPagination = queryPagination.Error

	} else if filter4 {
		fmt.Println("====== query 4")
		fraudIndication := strings.Split(request.FraudIndication, ",")

		if request.Sort == "DESC" {
			sortQuery = `grand_total DESC, id`
		} else {
			sortQuery = `grand_total ASC, id`
		}

		// subquery
		subquery1 = db.Table("verifikasi").
			Select("activity_id, product_id, risk_issue_id, risk_issue, perbaikan, risk_indicator, COUNT(*) as total").
			// Joins("JOIN verifikasi_risk_control vrc ON vrc.verifikasi_id = verifikasi.id").
			Where("perbaikan = ?", "1").
			Where("indikasi_fraud IN (?)", fraudIndication).
			Where("deleted != 1").
			Where("activity_id = ?", request.Activity).
			Where("product_id = ?", request.Product)

		subquery2 = db.Table("verifikasi").
			Select("activity_id, product_id, risk_issue_id, risk_issue, perbaikan, risk_indicator, COUNT(*) as total").
			// Joins("JOIN verifikasi_risk_control vrc ON vrc.verifikasi_id = verifikasi.id").
			Where("perbaikan = ?", "0").
			Where("indikasi_fraud IN (?)", fraudIndication).
			Where("deleted != 1").
			Where("activity_id = ?", request.Activity).
			Where("product_id = ?", request.Product)

		//// uker on subquery
		if request.REGION != "all" {
			subquery1 = subquery1.Where("verifikasi.REGION = ?", request.REGION)
			subquery2 = subquery2.Where("verifikasi.REGION = ?", request.REGION)
		}

		if request.REGION != "all" && request.MAINBR != "all" {
			mainbrs := strings.Split(request.MAINBR, ",")

			subquery1 = subquery1.Where("verifikasi.MAINBR in (?)", mainbrs)
			subquery2 = subquery2.Where("verifikasi.MAINBR in (?)", mainbrs)
		}

		if request.REGION != "all" && request.MAINBR != "all" && request.BRANCH != "all" {
			branches := strings.Split(request.BRANCH, ",")
			if len(branches) > 1 {
				fmt.Println("Masuk 1")
				subquery1 = subquery1.Where("verifikasi.BRANCH in (?)", branches)
				subquery2 = subquery2.Where("verifikasi.BRANCH in (?)", branches)
			} else {
				fmt.Println("Masuk 2")
				subquery1 = subquery1.Where("verifikasi.BRANCH = ?", request.BRANCH)
				subquery2 = subquery2.Where("verifikasi.BRANCH = ?", request.BRANCH)
			}

		}

		//// status on subquery
		// if request.Status == "draft" {
		// 	subquery1 = subquery1.Where("(status = '01a' AND action = 'draft')")
		// 	subquery2 = subquery2.Where("(status = '01a' AND action = 'draft')")
		// } else if request.Status == "selesai" {
		// 	subquery1 = subquery1.Where("((status = '02a' AND action = 'selesai') OR (status = '02b' AND (action = 'update' OR action = 'selesai')))")
		// 	subquery2 = subquery2.Where("((status = '02a' AND action = 'selesai') OR (status = '02b' AND (action = 'update' OR action = 'selesai')))")
		// }
		//edited By Panji #add status veladasi RMC & ORD
		if request.Status == "draft" {
			subquery1 = subquery1.Where("(status = '01a' AND lower(action) = 'draft') OR (status = '01b')")
			subquery2 = subquery2.Where("(status = '01a' AND lower(action) = 'draft') OR (status = '01b')")
		} else if request.Status == "selesai" {
			subquery1 = subquery1.Where("(status = '04a' AND lower(action) = 'selesai')")
			subquery2 = subquery2.Where("(status = '04a' AND lower(action) = 'selesai')")
		}

		subquery1 = subquery1.Where("(created_at >= ? AND created_at <= ?)", request.StartDate, lib.FixEndDate(request.EndDate)).
			Group("risk_indicator")

		subquery2 = subquery2.Where("(created_at >= ? AND created_at <= ?)", request.StartDate, lib.FixEndDate(request.EndDate)).
			Group("risk_indicator")

		// query
		query = db.Table("verifikasi v").
			Select(`
				v.risk_indicator_id as id,
				v.product_id as code,
				v.risk_indicator as name,
				COALESCE(weakness.total, 0) as total_weakness,
				COALESCE(non_weakness.total, 0) as total_non_weakness,
				(COALESCE(weakness.total, 0) + COALESCE(non_weakness.total, 0)) as grand_total`).
			Joins(`
				LEFT JOIN (?) weakness 
					ON weakness.activity_id = v.activity_id 
					AND weakness.product_id = v.product_id 
					AND weakness.risk_issue = v.risk_issue
					AND weakness.risk_indicator = v.risk_indicator`, subquery1).
			Joins(`
				LEFT JOIN (?) non_weakness 
					ON non_weakness.activity_id = v.activity_id 
					AND non_weakness.product_id = v.product_id 
					AND non_weakness.risk_issue = v.risk_issue
					AND non_weakness.risk_indicator = v.risk_indicator`, subquery2).
			Where(`v.activity_id = ?`, request.Activity).
			Where(`v.product_id = ?`, request.Product).
			Where(`v.risk_issue_id = ?`, request.RiskIssue).
			Group(`name`).
			Order(sortQuery).
			Limit(request.Limit).
			Offset(request.Offset).
			Find(&responses)
		err = query.Error

		// query pagination
		queryPagination = db.Table("verifikasi").
			Select(`SUM(
				COUNT(
					DISTINCT risk_indicator
				)
			) OVER() as pagination`).
			Where(`activity_id = ?`, request.Activity).
			Where(`product_id = ?`, request.Product).
			Where(`risk_issue_id = ?`, request.RiskIssue).
			Scan(&totalRows)
		errPagination = queryPagination.Error
	}

	//QUERY PAGINATION
	repo.logger.Zap.Info("verification-queryPagination-activity-reportfilter", queryPagination)

	if errPagination != nil {
		return responses, totalRows, err
	}

	repo.logger.Zap.Info("verification-query-activity-reportfilter", query)

	if err != nil {
		return responses, totalRows, err
	}

	fmt.Println("reponses repo")
	fmt.Println(responses)

	return responses, totalRows, err
}

// by aktivitas + weakness only
func (repo VerifikasiRepository) VerifikasiReportWithWeaknessOnlyFilter(request *models.VerifikasiFilterReportRequest) (responses []models.VerifikasiFilterReportWeaknessOnlyResponsWithoutPercent, totalRows int64, err error) {
	db := repo.db.DB
	query := db
	queryPagination := db
	sortQuery := ""

	subquery1 := db

	var errPagination error

	//bank wide
	filter1 := request.ReportType == "aktivitas" &&
		request.Activity == "all" &&
		request.Product == "all" &&
		(request.RiskIssue == "all" || request.RiskIssue == "Semua") &&
		(request.Title == "all" || request.Title == "Semua")

	filter2 := request.ReportType == "aktivitas" &&
		request.Activity != "all" &&
		request.Product == "all" &&
		(request.RiskIssue == "all" || request.RiskIssue == "Semua") &&
		(request.Title == "all" || request.Title == "Semua")

	filter3 := request.ReportType == "aktivitas" &&
		request.Activity != "all" &&
		request.Product != "all" &&
		(request.RiskIssue == "all" || request.RiskIssue == "Semua") &&
		(request.Title == "all" || request.Title == "Semua")

	filter4 := request.ReportType == "aktivitas" &&
		request.Activity != "all" &&
		request.Product != "all" &&
		(request.RiskIssue != "all" && request.RiskIssue != "Semua") &&
		(request.Title == "all" || request.Title == "Semua")

	if filter1 {
		fmt.Println("====== query 1")
		fraudIndication := strings.Split(request.FraudIndication, ",")

		if request.Sort == "DESC" {
			sortQuery = `total_weakness DESC, id`
		} else {
			sortQuery = `total_weakness ASC, id`
		}

		// subquery
		subquery1 = db.Table("verifikasi").
			Select("activity_id, perbaikan, COUNT(*) as total").
			Joins("JOIN verifikasi_risk_control vrc ON vrc.verifikasi_id = verifikasi.id").
			Where("perbaikan = ?", "1").
			Where("deleted != 1").
			Where("indikasi_fraud IN (?)", fraudIndication)

		//// uker on subquery
		if request.REGION != "all" {
			subquery1 = subquery1.Where("verifikasi.REGION = ?", request.REGION)
		}

		if request.REGION != "all" && request.MAINBR != "all" {
			mainbrs := strings.Split(request.MAINBR, ",")
			subquery1 = subquery1.Where("verifikasi.MAINBR in (?)", mainbrs)
		}

		if request.REGION != "all" && request.MAINBR != "all" && request.BRANCH != "all" {
			branches := strings.Split(request.BRANCH, ",")
			subquery1 = subquery1.Where("verifikasi.BRANCH in (?)", branches)
		}

		//// status on subquery
		// if request.Status == "draft" {
		// 	subquery1 = subquery1.Where("(status = '01a' AND action = 'draft')")
		// } else if request.Status == "selesai" {
		// 	subquery1 = subquery1.Where("((status = '02a' AND action = 'selesai') OR (status = '02b' AND (action = 'update' OR action = 'selesai')))")
		// }

		//edited By Panji #add status veladasi RMC & ORD
		if request.Status == "draft" {
			subquery1 = subquery1.Where("(status = '01a' AND lower(action) = 'draft') OR (status = '01b')")
		} else if request.Status == "selesai" {
			subquery1 = subquery1.Where("(status = '04a' AND lower(action) = 'selesai')")
		}

		subquery1 = subquery1.Where("(created_at >= ? AND created_at <= ?)", request.StartDate, lib.FixEndDate(request.EndDate)).
			Group("activity_id")

		// query
		query = db.Table("activity").
			Select(`
				id, 
				kode_activity as code, 
				name, 
				COALESCE(weakness.total, 0) as total_weakness`).
			Joins(`
				LEFT JOIN (?) weakness ON weakness.activity_id = activity.id`, subquery1).
			Order(sortQuery).
			Limit(request.Limit).
			Offset(request.Offset).
			Find(&responses)
		err = query.Error

		// query pagination
		queryPagination = db.Table("activity").Count(&totalRows)
		errPagination = queryPagination.Error

	} else if filter2 {
		fmt.Println("====== query 2")
		fraudIndication := strings.Split(request.FraudIndication, ",")

		if request.Sort == "DESC" {
			sortQuery = `total_weakness DESC, id`
		} else {
			sortQuery = `total_weakness ASC, id`
		}

		// subquery
		subquery1 = db.Table("verifikasi").
			Select("activity_id, product_id, perbaikan, COUNT(*) as total").
			Joins("JOIN verifikasi_risk_control vrc ON vrc.verifikasi_id = verifikasi.id").
			Where("perbaikan = ?", "1").
			Where("deleted != 1").
			Where("indikasi_fraud IN (?)", fraudIndication)

		//// uker on subquery
		if request.REGION != "all" {
			subquery1 = subquery1.Where("verifikasi.REGION = ?", request.REGION)
		}

		if request.REGION != "all" && request.MAINBR != "all" {
			mainbrs := strings.Split(request.MAINBR, ",")
			subquery1 = subquery1.Where("verifikasi.MAINBR in (?)", mainbrs)
		}

		if request.REGION != "all" && request.MAINBR != "all" && request.BRANCH != "all" {
			branches := strings.Split(request.BRANCH, ",")
			subquery1 = subquery1.Where("verifikasi.BRANCH in (?)", branches)
		}

		//// status on subquery
		// if request.Status == "draft" {
		// 	subquery1 = subquery1.Where("(status = '01a' AND action = 'draft')")
		// } else if request.Status == "selesai" {
		// 	subquery1 = subquery1.Where("((status = '02a' AND action = 'selesai') OR (status = '02b' AND (action = 'update' OR action = 'selesai')))")
		// }

		//edited By Panji #add status veladasi RMC & ORD
		if request.Status == "draft" {
			subquery1 = subquery1.Where("(status = '01a' AND lower(action) = 'draft') OR (status = '01b')")
		} else if request.Status == "selesai" {
			subquery1 = subquery1.Where("(status = '04a' AND lower(action) = 'selesai')")
		}

		subquery1 = subquery1.Where("(created_at >= ? AND created_at <= ?)", request.StartDate, lib.FixEndDate(request.EndDate)).
			Group("activity_id")

		// query
		query = db.Table("product p").
			Select(`
				id, 
				kode_product as code, 
				product as name, 
				COALESCE(weakness.total, 0) as total_weakness`).
			Joins(`
				LEFT JOIN (?) weakness ON weakness.activity_id = ? AND weakness.product_id = p.id`, subquery1, request.Activity).
			Where(`p.activity_id = ?`, request.Activity).
			Order(sortQuery).
			Limit(request.Limit).
			Offset(request.Offset).
			Find(&responses)
		err = query.Error

		// query pagination
		queryPagination = db.Table("product").Where("activity_id = ?", request.Activity).Count(&totalRows)
		errPagination = queryPagination.Error
	} else if filter3 {
		fmt.Println("====== query 3")
		fraudIndication := strings.Split(request.FraudIndication, ",")

		if request.Sort == "DESC" {
			sortQuery = `total_weakness DESC, id`
		} else {
			sortQuery = `total_weakness ASC, id`
		}

		// subquery
		subquery1 = db.Table("verifikasi").
			Select("activity_id, product_id, risk_issue_id, risk_issue, perbaikan, COUNT(*) as total").
			Joins("JOIN verifikasi_risk_control vrc ON vrc.verifikasi_id = verifikasi.id").
			Where("perbaikan = ?", "1").
			Where("deleted != 1").
			Where("indikasi_fraud IN (?)", fraudIndication).
			Where("activity_id = ?", request.Activity).
			Where("product_id = ?", request.Product)

		//// uker on subquery
		if request.REGION != "all" {
			subquery1 = subquery1.Where("verifikasi.REGION = ?", request.REGION)
		}

		if request.REGION != "all" && request.MAINBR != "all" {
			mainbrs := strings.Split(request.MAINBR, ",")
			subquery1 = subquery1.Where("verifikasi.MAINBR in (?)", mainbrs)
		}

		if request.REGION != "all" && request.MAINBR != "all" && request.BRANCH != "all" {
			branches := strings.Split(request.BRANCH, ",")
			subquery1 = subquery1.Where("verifikasi.BRANCH in (?)", branches)
		}

		//// status on subquery
		// if request.Status == "draft" {
		// 	subquery1 = subquery1.Where("(status = '01a' AND action = 'draft')")
		// } else if request.Status == "selesai" {
		// 	subquery1 = subquery1.Where("((status = '02a' AND action = 'selesai') OR (status = '02b' AND (action = 'update' OR action = 'selesai')))")
		// }

		//edited By Panji #add status veladasi RMC & ORD
		if request.Status == "draft" {
			subquery1 = subquery1.Where("(status = '01a' AND lower(action) = 'draft') OR (status = '01b')")
		} else if request.Status == "selesai" {
			subquery1 = subquery1.Where("(status = '04a' AND lower(action) = 'selesai')")
		}

		subquery1 = subquery1.Where("(created_at >= ? AND created_at <= ?)", request.StartDate, lib.FixEndDate(request.EndDate)).
			Group("risk_issue")

		// query
		query = db.Table("verifikasi v").
			Select(`
				v.risk_issue_id as id,
				v.product_id as code,
				v.risk_issue as name,
				COALESCE(weakness.total, 0) as total_weakness`).
			Joins(`
				LEFT JOIN (?) weakness 
					ON weakness.activity_id = v.activity_id 
					AND weakness.product_id = v.product_id 
					AND weakness.risk_issue = v.risk_issue`, subquery1).
			Where(`v.activity_id = ?`, request.Activity).
			Where(`v.product_id = ?`, request.Product).
			Group(`name`).
			Order(sortQuery).
			Limit(request.Limit).
			Offset(request.Offset).
			Find(&responses)
		err = query.Error

		// query pagination
		queryPagination = db.Table("verifikasi").
			Select(`SUM(
				COUNT(
					DISTINCT risk_issue
				)
			) OVER() as pagination`).
			Where(`activity_id = ?`, request.Activity).
			Where(`product_id = ?`, request.Product).
			Scan(&totalRows)
		errPagination = queryPagination.Error

	} else if filter4 {
		fmt.Println("====== query 4")
		fraudIndication := strings.Split(request.FraudIndication, ",")

		if request.Sort == "DESC" {
			sortQuery = `total_weakness DESC, id`
		} else {
			sortQuery = `total_weakness ASC, id`
		}

		// subquery
		subquery1 = db.Table("verifikasi").
			Select("activity_id, product_id, risk_issue_id, risk_issue, perbaikan, COUNT(*) as total").
			Joins("JOIN verifikasi_risk_control vrc ON vrc.verifikasi_id = verifikasi.id").
			Where("perbaikan = ?", "1").
			Where("deleted != 1").
			Where("indikasi_fraud IN (?)", fraudIndication).
			Where("activity_id = ?", request.Activity).
			Where("product_id = ?", request.Product)

		//// uker on subquery
		if request.REGION != "all" {
			subquery1 = subquery1.Where("verifikasi.REGION = ?", request.REGION)
		}

		if request.REGION != "all" && request.MAINBR != "all" {
			mainbrs := strings.Split(request.MAINBR, ",")
			subquery1 = subquery1.Where("verifikasi.MAINBR in (?)", mainbrs)
		}

		if request.REGION != "all" && request.MAINBR != "all" && request.BRANCH != "all" {
			branches := strings.Split(request.BRANCH, ",")
			subquery1 = subquery1.Where("verifikasi.BRANCH in (?)", branches)
		}

		//// status on subquery
		// if request.Status == "draft" {
		// 	subquery1 = subquery1.Where("(status = '01a' AND action = 'draft')")
		// } else if request.Status == "selesai" {
		// 	subquery1 = subquery1.Where("((status = '02a' AND action = 'selesai') OR (status = '02b' AND (action = 'update' OR action = 'selesai')))")
		// }

		//edited By Panji #add status veladasi RMC & ORD
		if request.Status == "draft" {
			subquery1 = subquery1.Where("(status = '01a' AND lower(action) = 'draft') OR (status = '01b')")
		} else if request.Status == "selesai" {
			subquery1 = subquery1.Where("(status = '04a' AND lower(action) = 'selesai')")
		}

		subquery1 = subquery1.Where("(created_at >= ? AND created_at <= ?)", request.StartDate, lib.FixEndDate(request.EndDate)).
			Group("risk_issue")

		// query
		query = db.Table("verifikasi v").
			Select(`
				v.risk_indicator_id as id,
				v.product_id as code,
				v.risk_indicator as name,
				COALESCE(weakness.total, 0) as total_weakness`).
			Joins(`
				LEFT JOIN (?) weakness 
					ON weakness.activity_id = v.activity_id 
					AND weakness.product_id = v.product_id 
					AND weakness.risk_issue = v.risk_issue`, subquery1).
			Where(`v.activity_id = ?`, request.Activity).
			Where(`v.product_id = ?`, request.Product).
			Where(`v.risk_issue_id = ?`, request.RiskIssue).
			Group(`name`).
			Order(sortQuery).
			Limit(request.Limit).
			Offset(request.Offset).
			Find(&responses)
		err = query.Error

		// query pagination
		queryPagination = db.Table("verifikasi").
			Select(`SUM(
				COUNT(
					DISTINCT risk_indicator
				)
			) OVER() as pagination`).
			Where(`activity_id = ?`, request.Activity).
			Where(`product_id = ?`, request.Product).
			Where(`risk_issue_id = ?`, request.RiskIssue).
			Scan(&totalRows)
		errPagination = queryPagination.Error
	}

	//QUERY PAGINATION
	repo.logger.Zap.Info("verification-queryPagination-activity-reportfilter", queryPagination)

	if errPagination != nil {
		return responses, totalRows, err
	}

	repo.logger.Zap.Info("verification-query-activity-reportfilter", query)

	if err != nil {
		return responses, totalRows, err
	}

	fmt.Println("reponses repo")
	fmt.Println(responses)

	return responses, totalRows, err
}

// by aktivitas + non weakness only
func (repo VerifikasiRepository) VerifikasiReportWithNonWeaknessOnlyFilter(request *models.VerifikasiFilterReportRequest) (responses []models.VerifikasiFilterReportNonWeaknessOnlyResponsWithoutPercentNull, totalRows int64, err error) {
	db := repo.db.DB
	query := db
	queryPagination := db
	sortQuery := ""

	subquery1 := db

	var errPagination error

	//bank wide
	filter1 := request.ReportType == "aktivitas" &&
		request.Activity == "all" &&
		request.Product == "all" &&
		(request.RiskIssue == "all" || request.RiskIssue == "Semua") &&
		(request.Title == "all" || request.Title == "Semua")

	filter2 := request.ReportType == "aktivitas" &&
		request.Activity != "all" &&
		request.Product == "all" &&
		(request.RiskIssue == "all" || request.RiskIssue == "Semua") &&
		(request.Title == "all" || request.Title == "Semua")

	filter3 := request.ReportType == "aktivitas" &&
		request.Activity != "all" &&
		request.Product != "all" &&
		(request.RiskIssue == "all" || request.RiskIssue == "Semua") &&
		(request.Title == "all" || request.Title == "Semua")

	filter4 := request.ReportType == "aktivitas" &&
		request.Activity != "all" &&
		request.Product != "all" &&
		(request.RiskIssue != "all" || request.RiskIssue != "Semua") &&
		(request.Title == "all" || request.Title == "Semua")

	if filter1 {
		fmt.Println("====== query 1")
		fraudIndication := strings.Split(request.FraudIndication, ",")

		if request.Sort == "DESC" {
			sortQuery = `total_weakness DESC, id`
		} else {
			sortQuery = `total_weakness ASC, id`
		}

		// subquery
		subquery1 = db.Table("verifikasi").
			Select("activity_id, perbaikan, COUNT(*) as total").
			Joins("JOIN verifikasi_risk_control vrc ON vrc.verifikasi_id = verifikasi.id").
			Where("perbaikan = ?", "0").
			Where("deleted != 1").
			Where("indikasi_fraud IN (?)", fraudIndication)

		//// uker on subquery
		if request.REGION != "all" {
			subquery1 = subquery1.Where("verifikasi.REGION = ?", request.REGION)
		}

		if request.REGION != "all" && request.MAINBR != "all" {
			mainbrs := strings.Split(request.MAINBR, ",")
			subquery1 = subquery1.Where("verifikasi.MAINBR in (?)", mainbrs)
		}

		if request.REGION != "all" && request.MAINBR != "all" && request.BRANCH != "all" {
			branches := strings.Split(request.BRANCH, ",")
			subquery1 = subquery1.Where("verifikasi.BRANCH in (?)", branches)
		}

		//// status on subquery
		// if request.Status == "draft" {
		// 	subquery1 = subquery1.Where("(status = '01a' AND action = 'draft')")
		// } else if request.Status == "selesai" {
		// 	subquery1 = subquery1.Where("((status = '02a' AND action = 'selesai') OR (status = '02b' AND (action = 'update' OR action = 'selesai')))")
		// }

		//edited By Panji #add status veladasi RMC & ORD
		if request.Status == "draft" {
			subquery1 = subquery1.Where("(status = '01a' AND lower(action) = 'draft') OR (status = '01b')")
		} else if request.Status == "selesai" {
			subquery1 = subquery1.Where("(status = '04a' AND lower(action) = 'selesai')")
		}

		subquery1 = subquery1.Where("(created_at >= ? AND created_at <= ?)", request.StartDate, lib.FixEndDate(request.EndDate)).
			Group("activity_id")

		// query
		query = db.Table("activity").
			Select(`
				id, 
				kode_activity as code, 
				name, 
				COALESCE(non_weakness.total, 0) as total_non_weakness`).
			Joins(`
				LEFT JOIN (?) non_weakness ON non_weakness.activity_id = activity.id`, subquery1).
			Order(sortQuery).
			Limit(request.Limit).
			Offset(request.Offset).
			Find(&responses)
		err = query.Error

		// query pagination
		queryPagination = db.Table("activity").Count(&totalRows)
		errPagination = queryPagination.Error

	} else if filter2 {
		fmt.Println("====== query 2")
		fraudIndication := strings.Split(request.FraudIndication, ",")

		if request.Sort == "DESC" {
			sortQuery = `total_weakness DESC, id`
		} else {
			sortQuery = `total_weakness ASC, id`
		}

		// subquery
		subquery1 = db.Table("verifikasi").
			Select("activity_id, product_id, perbaikan, COUNT(*) as total").
			Joins("JOIN verifikasi_risk_control vrc ON vrc.verifikasi_id = verifikasi.id").
			Where("perbaikan = ?", "0").
			Where("deleted != 1").
			Where("indikasi_fraud IN (?)", fraudIndication)

		//// uker on subquery
		if request.REGION != "all" {
			subquery1 = subquery1.Where("verifikasi.REGION = ?", request.REGION)
		}

		if request.REGION != "all" && request.MAINBR != "all" {
			mainbrs := strings.Split(request.MAINBR, ",")
			subquery1 = subquery1.Where("verifikasi.MAINBR in (?)", mainbrs)
		}

		if request.REGION != "all" && request.MAINBR != "all" && request.BRANCH != "all" {
			branches := strings.Split(request.BRANCH, ",")
			subquery1 = subquery1.Where("verifikasi.BRANCH in (?)", branches)
		}

		//// status on subquery
		// if request.Status == "draft" {
		// 	subquery1 = subquery1.Where("(status = '01a' AND action = 'draft')")
		// } else if request.Status == "selesai" {
		// 	subquery1 = subquery1.Where("((status = '02a' AND action = 'selesai') OR (status = '02b' AND (action = 'update' OR action = 'selesai')))")
		// }

		//edited By Panji #add status veladasi RMC & ORD
		if request.Status == "draft" {
			subquery1 = subquery1.Where("(status = '01a' AND lower(action) = 'draft') OR (status = '01b')")
		} else if request.Status == "selesai" {
			subquery1 = subquery1.Where("(status = '04a' AND lower(action) = 'selesai')")
		}

		subquery1 = subquery1.Where("(created_at >= ? AND created_at <= ?)", request.StartDate, lib.FixEndDate(request.EndDate)).
			Group("activity_id")

		// query
		query = db.Table("product p").
			Select(`
				id, 
				kode_product as code, 
				product as name, 
				COALESCE(non_weakness.total, 0) as total_non_weakness`).
			Joins(`
				LEFT JOIN (?) non_weakness ON non_weakness.activity_id = ? AND non_weakness.product_id = p.id`, subquery1, request.Activity).
			Where(`p.activity_id = ?`, request.Activity).
			Order(sortQuery).
			Limit(request.Limit).
			Offset(request.Offset).
			Find(&responses)
		err = query.Error

		// query pagination
		queryPagination = db.Table("product").Where("activity_id = ?", request.Activity).Count(&totalRows)
		errPagination = queryPagination.Error
	} else if filter3 {
		fmt.Println("====== query 3")
		fraudIndication := strings.Split(request.FraudIndication, ",")

		if request.Sort == "DESC" {
			sortQuery = `total_weakness DESC, id`
		} else {
			sortQuery = `total_weakness ASC, id`
		}

		// subquery
		subquery1 = db.Table("verifikasi").
			Select("activity_id, product_id, risk_issue_id, risk_issue, perbaikan, COUNT(*) as total").
			Joins("JOIN verifikasi_risk_control vrc ON vrc.verifikasi_id = verifikasi.id").
			Where("perbaikan = ?", "0").
			Where("deleted != 1").
			Where("indikasi_fraud IN (?)", fraudIndication).
			Where("activity_id = ?", request.Activity).
			Where("product_id = ?", request.Product)

		//// uker on subquery
		if request.REGION != "all" {
			subquery1 = subquery1.Where("verifikasi.REGION = ?", request.REGION)
		}

		if request.REGION != "all" && request.MAINBR != "all" {
			mainbrs := strings.Split(request.MAINBR, ",")
			subquery1 = subquery1.Where("verifikasi.MAINBR in (?)", mainbrs)
		}

		if request.REGION != "all" && request.MAINBR != "all" && request.BRANCH != "all" {
			branches := strings.Split(request.BRANCH, ",")
			subquery1 = subquery1.Where("verifikasi.BRANCH in (?)", branches)
		}

		//// status on subquery
		// if request.Status == "draft" {
		// 	subquery1 = subquery1.Where("(status = '01a' AND action = 'draft')")
		// } else if request.Status == "selesai" {
		// 	subquery1 = subquery1.Where("((status = '02a' AND action = 'selesai') OR (status = '02b' AND (action = 'update' OR action = 'selesai')))")
		// }

		//edited By Panji #add status veladasi RMC & ORD
		if request.Status == "draft" {
			subquery1 = subquery1.Where("(status = '01a' AND lower(action) = 'draft') OR (status = '01b')")
		} else if request.Status == "selesai" {
			subquery1 = subquery1.Where("(status = '04a' AND lower(action) = 'selesai')")
		}

		subquery1 = subquery1.Where("(created_at >= ? AND created_at <= ?)", request.StartDate, lib.FixEndDate(request.EndDate)).
			Group("risk_issue")

		// query
		query = db.Table("verifikasi v").
			Select(`
				v.risk_issue_id as id,
				v.product_id as code,
				v.risk_issue as name,
				COALESCE(non_weakness.total, 0) as total_non_weakness`).
			Joins(`
				LEFT JOIN (?) non_weakness 
					ON non_weakness.activity_id = v.activity_id 
					AND non_weakness.product_id = v.product_id 
					AND non_weakness.risk_issue = v.risk_issue`, subquery1).
			Where(`v.activity_id = ?`, request.Activity).
			Where(`v.product_id = ?`, request.Product).
			Group(`name`).
			Order(sortQuery).
			Limit(request.Limit).
			Offset(request.Offset).
			Find(&responses)
		err = query.Error

		// query pagination
		queryPagination = db.Table("verifikasi").
			Select(`SUM(
				COUNT(
					DISTINCT risk_issue
				)
			) OVER() as pagination`).
			Where(`activity_id = ?`, request.Activity).
			Where(`product_id = ?`, request.Product).
			Scan(&totalRows)
		errPagination = queryPagination.Error

	} else if filter4 {
		fmt.Println("====== query 4")
		fraudIndication := strings.Split(request.FraudIndication, ",")

		if request.Sort == "DESC" {
			sortQuery = `total_weakness DESC, id`
		} else {
			sortQuery = `total_weakness ASC, id`
		}

		// subquery
		subquery1 = db.Table("verifikasi").
			Select("activity_id, product_id, risk_issue_id, risk_issue, perbaikan, COUNT(*) as total").
			Joins("JOIN verifikasi_risk_control vrc ON vrc.verifikasi_id = verifikasi.id").
			Where("perbaikan = ?", "0").
			Where("deleted != 1").
			Where("indikasi_fraud IN (?)", fraudIndication).
			Where("activity_id = ?", request.Activity).
			Where("product_id = ?", request.Product)

		//// uker on subquery
		if request.REGION != "all" {
			subquery1 = subquery1.Where("verifikasi.REGION = ?", request.REGION)
		}

		if request.REGION != "all" && request.MAINBR != "all" {
			mainbrs := strings.Split(request.MAINBR, ",")
			subquery1 = subquery1.Where("verifikasi.MAINBR in (?)", mainbrs)
		}

		if request.REGION != "all" && request.MAINBR != "all" && request.BRANCH != "all" {
			branches := strings.Split(request.BRANCH, ",")
			subquery1 = subquery1.Where("verifikasi.BRANCH in (?)", branches)
		}

		//// status on subquery
		// if request.Status == "draft" {
		// 	subquery1 = subquery1.Where("(status = '01a' AND action = 'draft')")
		// } else if request.Status == "selesai" {
		// 	subquery1 = subquery1.Where("((status = '02a' AND action = 'selesai') OR (status = '02b' AND (action = 'update' OR action = 'selesai')))")
		// }

		//edited By Panji #add status veladasi RMC & ORD
		if request.Status == "draft" {
			subquery1 = subquery1.Where("(status = '01a' AND lower(action) = 'draft') OR (status = '01b')")
		} else if request.Status == "selesai" {
			subquery1 = subquery1.Where("(status = '04a' AND lower(action) = 'selesai')")
		}

		subquery1 = subquery1.Where("(created_at >= ? AND created_at <= ?)", request.StartDate, lib.FixEndDate(request.EndDate)).
			Group("risk_issue")

		// query
		query = db.Table("verifikasi v").
			Select(`
				v.risk_indicator_id as id,
				v.product_id as code,
				v.risk_indicator as name,
				COALESCE(non_weakness.total, 0) as total_non_weakness`).
			Joins(`
				LEFT JOIN (?) non_weakness 
					ON non_weakness.activity_id = v.activity_id 
					AND non_weakness.product_id = v.product_id 
					AND non_weakness.risk_issue = v.risk_issue`, subquery1).
			Where(`v.activity_id = ?`, request.Activity).
			Where(`v.product_id = ?`, request.Product).
			Where(`v.risk_issue_id = ?`, request.RiskIssue).
			Group(`name`).
			Order(sortQuery).
			Limit(request.Limit).
			Offset(request.Offset).
			Find(&responses)
		err = query.Error

		// query pagination
		queryPagination = db.Table("verifikasi").
			Select(`SUM(
				COUNT(
					DISTINCT risk_indicator
				)
			) OVER() as pagination`).
			Where(`activity_id = ?`, request.Activity).
			Where(`product_id = ?`, request.Product).
			Where(`risk_issue_id = ?`, request.RiskIssue).
			Scan(&totalRows)
		errPagination = queryPagination.Error
	}

	//QUERY PAGINATION
	repo.logger.Zap.Info("verification-queryPagination-activity-reportfilter", queryPagination)

	if errPagination != nil {
		return responses, totalRows, err
	}

	repo.logger.Zap.Info("verification-query-activity-reportfilter", query)

	if err != nil {
		return responses, totalRows, err
	}

	fmt.Println("reponses repo")
	fmt.Println(responses)

	return responses, totalRows, err
}

// report filter complete
func (repo VerifikasiRepository) VerifikasiReportFilterComplete(request *models.VerifikasiFilterReportRequest) (responses []models.VerifikasiFilterReportCompleteResponse, totalRows int64, err error) {
	reportType := request.Uker
	db := repo.db.DB
	uker := strings.Split(request.Uker, ",")
	var errPagination error

	query := db.Model(&responses).
		Select(`
			verifikasi.id 'id', 
			verifikasi.created_at 'date', 
			verifikasi.BRANCH 'BRANCH', 
			verifikasi.BRDESC 'BRDESC', 
			activity.name 'activity_name', 
			product.product 'product_name', 
			verifikasi.risk_issue 'risk_issue', 
			verifikasi.risk_indicator 'judul_materi'`).
		Joins("JOIN activity ON activity.kode_activity = verifikasi.activity_id").
		Joins("JOIN product ON product.id = verifikasi.product_id").
		Where("verifikasi.activity_id = ?", request.Activity).
		Where("verifikasi.product_id = ?", request.Product).
		Where("verifikasi.risk_issue_id = ?", request.RiskIssue).
		Where("verifikasi.deleted != 1").
		Where("verifikasi.risk_indicator_id = ?", request.RiskIndicator)

	if request.REGION != "all" {
		query = query.Where("verifikasi.REGION = ?", request.REGION)
	}

	if request.REGION != "all" && request.MAINBR != "all" {
		mainbrs := strings.Split(request.MAINBR, ",")
		query = query.Where("verifikasi.MAINBR in (?)", mainbrs)
	}

	if request.REGION != "all" && request.MAINBR != "all" && request.BRANCH != "all" {
		branches := strings.Split(request.BRANCH, ",")
		if len(branches) > 1 {
			fmt.Println("Masuk 1")
			query = query.Where("verifikasi.BRANCH in (?)", branches)
		} else {
			fmt.Println("Masuk 2")
			query = query.Where("verifikasi.BRANCH = ?", request.BRANCH)
		}

	}

	query = query.Where("(verifikasi.created_at >= ? AND verifikasi.created_at <= ?)", request.StartDate, lib.FixEndDate(request.EndDate)).
		Order("verifikasi.id").
		Limit(request.Limit).
		Offset(request.Offset).
		Find(&responses)

	queryPagination := db.Table("verifikasi").
		Joins("JOIN activity a ON a.kode_activity = verifikasi.activity_id").
		Joins("JOIN product p ON p.id = verifikasi.product_id").
		Where("verifikasi.activity_id = ?", request.Activity).
		Where("verifikasi.product_id = ?", request.Product).
		Where("verifikasi.risk_issue_id = ?", request.RiskIssue).
		Where("verifikasi.risk_indicator_id = ?", request.RiskIndicator)

	if request.REGION != "all" {
		queryPagination = queryPagination.Where("verifikasi.REGION = ?", request.REGION)
	}

	if request.REGION != "all" && request.MAINBR != "all" {
		mainbrs := strings.Split(request.MAINBR, ",")
		queryPagination = queryPagination.Where("verifikasi.MAINBR in (?)", mainbrs)
	}

	if request.REGION != "all" && request.MAINBR != "all" && request.BRANCH != "all" {
		branches := strings.Split(request.BRANCH, ",")
		if len(branches) > 1 {
			fmt.Println("Masuk 1")
			queryPagination = queryPagination.Where("verifikasi.BRANCH in (?)", branches)
		} else {
			fmt.Println("Masuk 2")
			queryPagination = queryPagination.Where("verifikasi.BRANCH = ?", request.BRANCH)
		}

	}

	queryPagination = queryPagination.Where("(verifikasi.created_at >= ? AND verifikasi.created_at <= ?)", request.StartDate, lib.FixEndDate(request.EndDate)).
		Count(&totalRows)

	if reportType != "bank_wide" {
		query = query.Where("AND verifikasi.BRANCH IN ?", uker)
		queryPagination = queryPagination.Where("AND verifikasi.BRANCH IN ?", uker)
	}

	//QUERY PAGINATION
	repo.logger.Zap.Info("verifikasi-queryPagination-activity-unknown", queryPagination)

	if errPagination != nil {
		return responses, totalRows, err
	}

	repo.logger.Zap.Info("verifikasi-query-activity-unknown", query)

	// repo.logger.Zap.Info("verifikasi-rows-activity-unknown", rows)
	if err != nil {
		return responses, totalRows, err
	}

	// response := models.VerifikasiFilterReportCompleteResponseNull{}
	// for rows.Next() {
	// 	_ = rows.Scan(
	// 		&response.Id,
	// 		&response.Date,
	// 		&response.BRANCH,
	// 		&response.BRDESC,
	// 		&response.ActivityName,
	// 		&response.ProductName,
	// 		&response.RiskIssue,
	// 		&response.JudulMateri,
	// 	)

	// 	responses = append(responses, response)
	// }

	// if err = rows.Err(); err != nil {
	// 	return responses, totalRows, err
	// }

	fmt.Println("reponses repo")
	fmt.Println(responses)

	return responses, totalRows, err
}

// detail
func (repo VerifikasiRepository) VerifikasiReportDetail(request *models.VerifikasiReportDetailRequest) (responsesDetail models.VerifikasiReportDetailResponseWithoutDataAnomaliNull, err error) {
	var rowsGetVerificationDetail *sql.Rows

	queryGetVerificationDetail := `
			SELECT 
				v.id,
				v.no_pelaporan,
				v.BRANCH,
				v.BRDESC,
				v.MAINBR,
				v.MBDESC,
				v.REGION,
				v.RGDESC,
				a.name as activity_name,
				sa.name as sub_activiy_name,
				p.product as product_name,
				v.risk_issue,
				v.risk_indicator,
				pkl1.kode_kejadian as incident_cause_code,
				pkl1.penyebab_kejadian as incident_cause,
				pkl3.kode_penyebab_kejadian_lv3 as sub_incident_cause_code,
				pkl3.penyebab_kejadian_lv3 as sub_incident_cause,
				v.hasil_verifikasi as verification_result,
				v.sumber_data as data_source,
				v.perbaikan,
				v.indikasi_fraud
			FROM 
				verifikasi v 
			LEFT JOIN
				activity a 
				ON a.kode_activity = v.activity_id 
			LEFT JOIN sub_activity sa 
				ON sa.id = v.sub_activity_id
			LEFT JOIN product p 
				ON p.id = v.product_id
			LEFT JOIN penyebab_kejadian_lv1 pkl1
				ON pkl1.id = v.incident_cause_id
			LEFT JOIN penyebab_kejadian_lv3 pkl3
				ON pkl3.id = v.sub_incident_cause_id 
			WHERE 
				v.id = ?
			`
	repo.logger.Zap.Info("verifikasi-query-activity-unknown", queryGetVerificationDetail)

	rowsGetVerificationDetail, err = repo.db.DB.Raw(queryGetVerificationDetail, request.Id).Rows()
	defer rowsGetVerificationDetail.Close()

	repo.logger.Zap.Info("verifikasi-rows-activity-unknown", rowsGetVerificationDetail)
	if err != nil {
		return responsesDetail, err
	}

	responseDetail := models.VerifikasiReportDetailResponseWithoutDataAnomaliNull{}
	for rowsGetVerificationDetail.Next() {
		_ = rowsGetVerificationDetail.Scan(
			&responseDetail.Id,
			&responseDetail.NoPelaporan,
			&responseDetail.BRANCH,
			&responseDetail.BRDESC,
			&responseDetail.MAINBR,
			&responseDetail.MBDESC,
			&responseDetail.REGION,
			&responseDetail.RGDESC,
			&responseDetail.ActivityName,
			&responseDetail.SubActivityName,
			&responseDetail.ProductName,
			&responseDetail.RiskIssue,
			&responseDetail.RiskIndicator,
			&responseDetail.IncidentCauseCode,
			&responseDetail.IncidentCauseName,
			&responseDetail.SubIncidentCauseCode,
			&responseDetail.SubIncidentCauseName,
			&responseDetail.VerificationResult,
			&responseDetail.DataSource,
			&responseDetail.Perbaikan,
			&responsesDetail.IndikasiFraud,
		)

		responsesDetail = responseDetail
	}

	if err = rowsGetVerificationDetail.Err(); err != nil {
		// return responsesDetail, responsesDataAnomali, err
		return responsesDetail, err
	}

	// get data anomali
	/*
		queryGetDataAnomali := ""

		if responsesDetail.DataSource.String == "KRID" {
			queryGetDataAnomali = `
				SELECT
					tanggal_kejadian as date,
					nomor_rekening as no_rek,
					nominal,
					keterangan
				FROM
					verifikasi_data_anomali_krid
				WHERE
					verifikasi_id = ` + request.Id + `
			`
		} else {
			queryGetDataAnomali = `
				SELECT
					tanggal_kejadian as date,
					nomor_rekening as no_rek,
					nominal,
					keterangan
				FROM
					verifikasi_data_anomali
				WHERE
					verifikasi_id = ` + request.Id + `
			`
		}

		repo.logger.Zap.Info("verifikasi-query-activity-unknown", queryGetDataAnomali)
		rowsGetDataAnomali, errGetDataAnomali := repo.dbRaw.DB.Query(queryGetDataAnomali)

		repo.logger.Zap.Info("verifikasi-rows-activity-unknown", rowsGetDataAnomali)
		if errGetDataAnomali != nil {
			return responsesDetail, responsesDataAnomali, errGetDataAnomali
		}

		responseDataAnomali := models.DataAnomaliNull{}
		for rowsGetDataAnomali.Next() {
			_ = rowsGetDataAnomali.Scan(
				&responseDataAnomali.Date,
				&responseDataAnomali.NoRek,
				&responseDataAnomali.Nominal,
				&responseDataAnomali.Keterangan,
			)

			responsesDataAnomali = append(responsesDataAnomali, responseDataAnomali)
		}

		if err = rowsGetDataAnomali.Err(); err != nil {
			return responsesDetail, responsesDataAnomali, err
		}
		fmt.Println("")
		fmt.Println("responsesDetail repo")
		fmt.Println(responsesDetail)
		fmt.Println("")
	*/

	// return responsesDetail, responsesDataAnomali, err
	return responsesDetail, err
}

// risk control
func (repo VerifikasiRepository) RiskControlByVerificationId(request *models.DataRiskControlRequest) (responses []models.DataRiskIndicatorResponseWithNoPercent, totalRows int64, err error) {
	db := repo.db.DB
	var errPagination error

	weakness := strings.Split(request.Weakness, ",")

	query := db.Model(&responses).
		Select(`
			COALESCE(vrc.risk_control, 'Other') AS risk_control,
			count(vrc.id) as total`).
		Joins("JOIN verifikasi_risk_control vrc ON vrc.verifikasi_id = verifikasi.id").
		Where("verifikasi.activity_id = ?", request.Activity).
		Where("verifikasi.perbaikan IN ?", weakness).
		Where("(verifikasi.created_at >= ? AND verifikas.created_at <= ?)", request.StartDate, lib.FixEndDate(request.EndDate))

	queryPagination := db.Table("verifikasi").
		Select(`SUM(COUNT(DISTINCT vrc.risk_control)) OVER() as totalRows`).
		Joins("JOIN verifikasi_risk_control vrc ON vrc.verifikasi_id = verifikasi.id").
		Where("verifikasi.activity_id = ?", request.Activity).
		Where("verifikasi.perbaikan IN ?", weakness).
		Where("(verifikasi.created_at >= ? AND verifikasi.created_at <= ?)", request.StartDate, lib.FixEndDate(request.EndDate))

	if request.Product != "" {
		query = query.Where("verifikasi.product_id = ?", request.Product)
		queryPagination = queryPagination.Where("verifikasi.product_id = ?", request.Product)
	}

	if request.RiskIssue != "" {
		query = query.Where("verifikasi.risk_issue_id = ?", request.RiskIssue)
		queryPagination = queryPagination.Where("verifikasi.product_id = ?", request.Product)
	}

	if request.RiskIndicator != "" {
		query = query.Where("verifikasi.risk_indicator_id = ?", request.RiskIndicator)
		queryPagination = queryPagination.Where("verifikasi.product_id = ?", request.Product)
	}

	if request.REGION != "all" {
		query = query.Where("verifikasi.REGION = ?", request.REGION)
		queryPagination = queryPagination.Where("verifikasi.REGION = ?", request.REGION)
	}

	if request.REGION != "all" && request.MAINBR != "all" {
		mainbrs := strings.Split(request.MAINBR, ",")
		query = query.Where("verifikasi.MAINBR in (?)", mainbrs)
		queryPagination = queryPagination.Where("verifikasi.MAINBR in (?)", mainbrs)
	}

	if request.REGION != "all" && request.MAINBR != "all" && request.BRANCH != "all" {
		branches := strings.Split(request.BRANCH, ",")
		query = query.Where("verifikasi.BRANCH in (?)", branches)
		queryPagination = queryPagination.Where("verifikasi.MAINBR in (?)", branches)
	}

	query.
		Group("vrc.risk_control").
		Limit(request.Limit).
		Offset(request.Offset).
		Find(&responses)

	queryPagination.Find(&totalRows)

	if errPagination != nil {
		return responses, totalRows, err
	}

	repo.logger.Zap.Info("verifikasi-query-activity-unknown", query)

	if err != nil {
		return responses, totalRows, err
	}

	fmt.Println("reponses repo")
	fmt.Println(responses)

	return responses, totalRows, err
}

// get risk indicator as materi
func (repo VerifikasiRepository) GetRiskIndicatorAsMateri(request *models.VerifikasiFilterReportRequest) (responses []models.GetRiskIndicatorAsMateriResponseNull, err error) {
	var rows *sql.Rows

	query := `
				SELECT
					v.risk_indicator_id as id,
					v.risk_indicator_id as code,
					v.risk_indicator as name
		 		FROM 
					verifikasi v 
				WHERE 
					v.activity_id = ?
				AND v.product_id = ?
				AND v.risk_issue_id = ?
				GROUP BY name
			`

	repo.logger.Zap.Info("verifikasi-query-activity-unknown", query, request.Activity, request.Product, request.RiskIssue)

	rows, err = repo.db.DB.Raw(query, request.Activity, request.Product, request.RiskIssue).Rows()
	if err != nil {
		return responses, err
	}
	defer rows.Close()

	repo.logger.Zap.Info("verifikasi-rows-activity-unknown", rows)
	if err != nil {
		return responses, err
	}

	response := models.GetRiskIndicatorAsMateriResponseNull{}
	for rows.Next() {
		_ = rows.Scan(
			&response.ID,
			&response.Code,
			&response.Name,
		)
		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, err
	}

	return responses, err
}

// REPORT BY UKER
func (repo VerifikasiRepository) VerificationReportByUkerFilter(request *models.VerificationFilterReportByUkerRequest) (responses []models.VerificationFilterReportByUkerResponseNull, SumData int64, totalRows int, err error) {
	query := ""
	queryStatus := ""
	queryPagination := ""
	var rows *sql.Rows
	sortQuery := ""
	var errPagination error

	if request.Status == "draft" {
		queryStatus = "AND ((v.status = '01a' AND lower(v.action) = 'draft') OR (v.status = '01b'))"
	} else if request.Status == "selesai" {
		queryStatus = "AND (status = '01a' AND lower(action) = 'draft')"
		queryStatus = "AND (v.status = '04a' AND lower(v.action) = 'selesai')"
	}

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

	// add by panji 23/11/2023 penyesuaian BRC
	filter4 := request.ReportType == "unitkerja" &&
		request.REGION != "all" &&
		request.MAINBR != "all" &&
		request.BRANCH != "all"

	if filter1 {
		fmt.Println("filter1")

		if request.Sort == "DESC" {
			sortQuery = `ORDER BY TOTALVERIFICATION DESC, uker.REGION`
		} else {
			sortQuery = `ORDER BY TOTALVERIFICATION ASC, uker.REGION`
		}

		fraudIndication := strings.Split(request.FraudIndication, ",")

		query = `SELECT
					uker.REGION,
					uker.RGDESC,
					uker.MAINBR,
					uker.MBDESC,
					uker.BRANCH,
					uker.BRDESC,
					COALESCE ( total_verification.TOTAL, 0 ) AS TOTALVERIFICATION,
					COALESCE ( brc.TOTAL, 0 ) AS TOTALBRC,
					COALESCE (total_non_weakness.TOTAL, 0) AS TOTALNONWEAKNESS,
					COALESCE (total_weakness.TOTAL, 0) AS TOTALNWEAKNESS,
					COALESCE (total_status_perbaikan_on_progress.TOTAL, 0) AS TOTALPERBAIKANONPROGRESS,
					COALESCE (total_status_perbaikan_done.TOTAL, 0) AS TOTALPERBAIKANDONE
				FROM
					dwh_branch uker
					LEFT JOIN (
						SELECT
							REGION,
							MAINBR,
							BRANCH,
							COUNT( id ) AS TOTAL 
						FROM
							verifikasi v
						WHERE
							(created_at >= ? AND created_at <= ?) 
							` + queryStatus + `
							AND v.indikasi_fraud IN ?
							AND v.deleted != 1
						GROUP BY REGION
					) total_verification ON total_verification.REGION = uker.REGION
					LEFT JOIN (
						SELECT
							kelolaan.REGION,
							kelolaan.MAINBR,
							kelolaan.BRANCH,
							COUNT( kelolaan.REGION ) AS TOTAL 
						FROM
							(
							SELECT
								REGION,
								MAINBR,
								BRANCH,
								COUNT( pn ) AS TOTAL 
							FROM
								uker_kelolaan_user 
							GROUP BY
								pn 
							) kelolaan 
						GROUP BY
							kelolaan.REGION
					) brc ON brc.REGION = uker.REGION 
					LEFT JOIN (
						SELECT 
							v.id,
							REGION,
							MAINBR,
							BRANCH,
							COUNT(v.id) as TOTAL
						FROM
							verifikasi v
						WHERE
							v.perbaikan = '0'
							` + queryStatus + `
						AND (created_at >= ? AND created_at <= ?)
						AND v.indikasi_fraud IN ?
						AND v.deleted != 1
						GROUP BY 
							REGION
					) total_non_weakness ON total_non_weakness.REGION = uker.REGION
					LEFT JOIN (
						SELECT 
							v.id,
							REGION,
							MAINBR,
							BRANCH,
							COUNT(v.id) as TOTAL
						FROM
							verifikasi v
						WHERE
							v.perbaikan = '1'
							` + queryStatus + `
						AND (created_at >= ? AND created_at <= ?)
						AND v.indikasi_fraud IN ?
						AND v.deleted != 1
						GROUP BY 
							REGION
					) total_weakness ON total_weakness.REGION = uker.REGION
					LEFT JOIN (
						SELECT 
							v.id,
							REGION,
							MAINBR,
							BRANCH,
							COUNT(v.id) as TOTAL
						FROM
							verifikasi v
						JOIN 
							verifikasi_pic_tindak_lanjut vptl 
							ON v.id = vptl.verifikasi_id 
						WHERE
							vptl.status = '1'
							` + queryStatus + `
						AND ( created_at >= ? AND created_at <= ? )
						AND v.indikasi_fraud IN ?
						AND v.deleted != 1
						GROUP BY 
							REGION
					) total_status_perbaikan_on_progress ON total_status_perbaikan_on_progress.REGION = uker.REGION
					LEFT JOIN (
						SELECT 
							v.id,
							REGION,
							MAINBR,
							BRANCH,
							COUNT(v.id) as TOTAL
						FROM
							verifikasi v
						JOIN 
							verifikasi_pic_tindak_lanjut vptl 
							ON v.id = vptl.verifikasi_id 
						WHERE
							vptl.status = '2'
							` + queryStatus + `
						AND (created_at >= ? AND created_at <= ?)
						AND v.indikasi_fraud IN ?
						AND v.deleted != 1
						GROUP BY 
							REGION
					) total_status_perbaikan_done ON total_status_perbaikan_done.REGION = uker.REGION
				WHERE
					uker.BRUNIT = 'B' 
				AND uker.MAINBR = uker.BRANCH 
				AND (
					uker.BRDESC LIKE "kanwil%" 
					OR uker.BRDESC = "Jkt KCK" 
				)
				GROUP BY
					REGION
				` + sortQuery + `
				LIMIT ? OFFSET ?`

		queryPagination = `
				SELECT
					COUNT(*) AS pagination 
				FROM 
					dwh_branch 
				WHERE 
					BRUNIT = 'B'
				AND (
					BRDESC LIKE "kanwil%" 
					OR BRDESC = "Jkt KCK" 
				)
			`

		repo.db.DB.Table("verifikasi").
			Select(`COUNT(*) 'SumData'`).
			Where(`(created_at >= ? AND created_at <= ?)`, request.StartDate, lib.FixEndDate(request.EndDate)).
			Where(`indikasi_fraud IN (?)`, fraudIndication).
			Find(&SumData)

		rows, err = repo.db.DB.Raw(query, request.StartDate, lib.FixEndDate(request.EndDate), fraudIndication, request.StartDate, lib.FixEndDate(request.EndDate), fraudIndication, request.StartDate, lib.FixEndDate(request.EndDate), fraudIndication, request.StartDate, lib.FixEndDate(request.EndDate), fraudIndication, request.StartDate, lib.FixEndDate(request.EndDate), fraudIndication, strconv.Itoa(request.Limit), strconv.Itoa(request.Offset)).Rows()
		defer rows.Close()

		errPagination = repo.db.DB.Raw(queryPagination).Scan(&totalRows).Error
	} else if filter2 {
		fmt.Println("filter2")

		if request.Sort == "DESC" {
			sortQuery = `ORDER BY TOTALVERIFICATION DESC, uker.MAINBR`
		} else {
			sortQuery = `ORDER BY TOTALVERIFICATION ASC, uker.MAINBR`
		}

		fraudIndication := strings.Split(request.FraudIndication, ",")

		query = `SELECT
					uker.REGION,
					uker.RGDESC,
					uker.MAINBR,
					uker.MBDESC,
					uker.BRANCH,
					uker.BRDESC,
					COALESCE ( total_verification.TOTAL, 0 ) AS TOTALVERIFICATION,
					COALESCE ( brc.TOTAL, 0 ) AS TOTALBRC,
					COALESCE (total_non_weakness.TOTAL, 0) AS TOTALNONWEAKNESS,
					COALESCE (total_weakness.TOTAL, 0) AS TOTALNWEAKNESS,
					COALESCE (total_status_perbaikan_on_progress.TOTAL, 0) AS TOTALPERBAIKANONPROGRESS,
					COALESCE (total_status_perbaikan_done.TOTAL, 0) AS TOTALPERBAIKANDONE
				FROM
					dwh_branch uker
					LEFT JOIN (
					SELECT
						REGION,
						MAINBR,
						BRANCH,
						COUNT( id ) AS TOTAL 
					FROM
						verifikasi v
					WHERE
						( created_at >= ? AND created_at <= ? ) 
						` + queryStatus + `
						AND v.indikasi_fraud IN ?
						AND v.deleted != 1
					GROUP BY
						MAINBR 
					) total_verification ON total_verification.REGION = uker.REGION 
					AND total_verification.MAINBR = uker.MAINBR
					LEFT JOIN (
					SELECT
						kelolaan.REGION,
						kelolaan.MAINBR,
						kelolaan.BRANCH,
						COUNT( kelolaan.REGION ) AS TOTAL 
					FROM
						(
						SELECT
							REGION,
							MAINBR,
							BRANCH,
							COUNT( pn ) AS TOTAL 
						FROM
							uker_kelolaan_user 
						GROUP BY
							pn 
						) kelolaan 
					GROUP BY
						kelolaan.REGION , kelolaan.MAINBR
					) brc ON brc.REGION = uker.REGION 
					AND brc.MAINBR = uker.MAINBR
					LEFT JOIN (
						SELECT 
							v.id,
							REGION,
							MAINBR,
							BRANCH,
							COUNT(v.id) as TOTAL
						FROM
							verifikasi v
						WHERE
							v.perbaikan = '0'
							` + queryStatus + `
						AND (created_at >= ? AND created_at <= ?)
						AND v.indikasi_fraud IN ?
						AND v.deleted != 1
						GROUP BY 
							MAINBR
					) total_non_weakness ON total_non_weakness.REGION = uker.REGION
					AND total_non_weakness.MAINBR = uker.MAINBR
					LEFT JOIN (
						SELECT 
							v.id,
							REGION,
							MAINBR,
							BRANCH,
							COUNT(v.id) as TOTAL
						FROM
							verifikasi v
						WHERE
							v.perbaikan = '1'
							` + queryStatus + `
						AND ( created_at >= ? AND created_at <= ? )
						AND v.indikasi_fraud IN ?
						AND v.deleted != 1
						GROUP BY 
							MAINBR
					) total_weakness ON total_weakness.REGION = uker.REGION
					AND total_weakness.MAINBR = uker.MAINBR
					LEFT JOIN (
						SELECT 
							v.id,
							REGION,
							MAINBR,
							BRANCH,
							COUNT(v.id) as TOTAL
						FROM
							verifikasi v
						JOIN 
							verifikasi_pic_tindak_lanjut vptl 
							ON v.id = vptl.verifikasi_id 
						WHERE
							vptl.status = '1'
							` + queryStatus + `
						AND (created_at >= ? AND created_at <= ?)
						AND v.indikasi_fraud IN ?
						AND v.deleted != 1
						GROUP BY 
							MAINBR
					) total_status_perbaikan_on_progress ON total_status_perbaikan_on_progress.REGION = uker.REGION
					AND total_status_perbaikan_on_progress.MAINBR = uker.MAINBR
					LEFT JOIN (
						SELECT 
							v.id,
							REGION,
							MAINBR,
							BRANCH,
							COUNT(v.id) as TOTAL
						FROM
							verifikasi v
						JOIN 
							verifikasi_pic_tindak_lanjut vptl 
							ON v.id = vptl.verifikasi_id 
						WHERE
							vptl.status = '2'
							` + queryStatus + `
						AND (created_at >= ? AND created_at <= ?)
						AND v.indikasi_fraud IN ?
						AND v.deleted != 1
						GROUP BY 
							MAINBR
					) total_status_perbaikan_done ON total_status_perbaikan_done.REGION = uker.REGION
					AND total_status_perbaikan_done.MAINBR = uker.MAINBR
				WHERE
					uker.REGION = ?
					AND uker.MBDESC = uker.BRDESC
				GROUP BY
					MAINBR
				` + sortQuery + `
				LIMIT ? OFFSET ?`

		repo.db.DB.Table("verifikasi").
			Select(`COUNT(*) 'SumData'`).
			Where(`REGION = ?`, request.REGION).
			Where(`(created_at >= ? AND created_at <= ?)`, request.StartDate, lib.FixEndDate(request.EndDate)).
			Where(`indikasi_fraud IN (?)`, fraudIndication).
			Find(&SumData)

		queryPagination = `
				SELECT
					COUNT(DISTINCT MAINBR) AS pagination 
				FROM 
					dwh_branch 
				WHERE 
					REGION = ?`

		rows, err = repo.db.DB.Raw(query, request.StartDate, lib.FixEndDate(request.EndDate), fraudIndication, request.StartDate, lib.FixEndDate(request.EndDate), fraudIndication, request.StartDate, lib.FixEndDate(request.EndDate), fraudIndication, request.StartDate, lib.FixEndDate(request.EndDate), fraudIndication, request.StartDate, lib.FixEndDate(request.EndDate), fraudIndication, request.REGION, strconv.Itoa(request.Limit), strconv.Itoa(request.Offset)).Rows()
		defer rows.Close()

		errPagination = repo.db.DB.Raw(queryPagination, request.REGION).Scan(&totalRows).Error

	} else if filter3 {
		fmt.Println("filter3")

		if request.Sort == "DESC" {
			sortQuery = `ORDER BY TOTALVERIFICATION DESC, uker.BRANCH`
		} else {
			sortQuery = `ORDER BY TOTALVERIFICATION ASC, uker.BRANCH`
		}

		fraudIndication := strings.Split(request.FraudIndication, ",")
		mainbrs := strings.Split(request.MAINBR, ",")

		query = `SELECT
					uker.REGION,
					uker.RGDESC,
					uker.MAINBR,
					uker.MBDESC,
					uker.BRANCH,
					uker.BRDESC,
					COALESCE ( total_verification.TOTAL, 0 ) AS TOTALVERIFICATION,
					COALESCE ( brc.TOTAL, 0 ) AS TOTALBRC,
					COALESCE (total_non_weakness.TOTAL, 0) AS TOTALNONWEAKNESS,
					COALESCE (total_weakness.TOTAL, 0) AS TOTALNWEAKNESS,
					COALESCE (total_status_perbaikan_on_progress.TOTAL, 0) AS TOTALPERBAIKANONPROGRESS,
					COALESCE (total_status_perbaikan_done.TOTAL, 0) AS TOTALPERBAIKANDONE
				FROM
					dwh_branch uker
					LEFT JOIN (
					SELECT
						REGION,
						MAINBR,
						BRANCH,
						COUNT( id ) AS TOTAL 
					FROM
						verifikasi v
					WHERE
						( created_at >= ? AND created_at <= ?) 
						` + queryStatus + `
						AND v.indikasi_fraud IN ?
						AND v.deleted != 1
					GROUP BY
						BRANCH 
					) total_verification ON total_verification.BRANCH = uker.BRANCH
					LEFT JOIN (
					SELECT
						kelolaan.REGION,
						kelolaan.MAINBR,
						kelolaan.BRANCH,
						COUNT( kelolaan.REGION ) AS TOTAL 
					FROM
						(
						SELECT
							REGION,
							MAINBR,
							BRANCH,
							COUNT( pn ) AS TOTAL 
						FROM
							uker_kelolaan_user 
						GROUP BY
							pn 
						) kelolaan 
					GROUP BY
						kelolaan.REGION, kelolaan.MAINBR, kelolaan.BRANCH
					) brc ON brc.BRANCH = uker.BRANCH
					LEFT JOIN (
						SELECT 
							v.id,
							REGION,
							MAINBR,
							BRANCH,
							COUNT(v.id) as TOTAL
						FROM
							verifikasi v
						WHERE
							v.perbaikan = '0'
							` + queryStatus + `
						AND (created_at >= ? AND created_at <= ?)
						AND v.indikasi_fraud IN ?
						AND v.deleted != 1
						GROUP BY 
							BRANCH
					) total_non_weakness ON total_non_weakness.BRANCH = uker.BRANCH
					LEFT JOIN (
						SELECT 
							v.id,
							REGION,
							MAINBR,
							BRANCH,
							COUNT(v.id) as TOTAL
						FROM
							verifikasi v
						WHERE
							v.perbaikan = '1'
							` + queryStatus + `
						AND (created_at >= ? AND created_at <= ?)
						AND v.indikasi_fraud IN ?
						GROUP BY 
							BRANCH
					) total_weakness ON total_weakness.BRANCH = uker.BRANCH
					LEFT JOIN (
						SELECT 
							v.id,
							REGION,
							MAINBR,
							BRANCH,
							COUNT(v.id) as TOTAL
						FROM
							verifikasi v
						JOIN 
							verifikasi_pic_tindak_lanjut vptl 
							ON v.id = vptl.verifikasi_id 
						WHERE
							vptl.status = '1'
							` + queryStatus + `
						AND (created_at >= ? AND created_at <= ?)
						AND v.indikasi_fraud IN ?
						AND v.deleted != 1
						GROUP BY 
							BRANCH
					) total_status_perbaikan_on_progress ON total_status_perbaikan_on_progress.BRANCH = uker.BRANCH
					LEFT JOIN (
						SELECT 
							v.id,
							REGION,
							MAINBR,
							BRANCH,
							COUNT(v.id) as TOTAL
						FROM
							verifikasi v
						JOIN 
							verifikasi_pic_tindak_lanjut vptl 
							ON v.id = vptl.verifikasi_id 
						WHERE
							vptl.status = '2'
							` + queryStatus + `
						AND (created_at >= ? AND created_at <= ?)
						AND v.indikasi_fraud IN ?
						AND v.deleted != 1
						GROUP BY 
							BRANCH
					) total_status_perbaikan_done ON total_status_perbaikan_done.BRANCH = uker.BRANCH
				WHERE
					uker.REGION = ?
				AND uker.MAINBR in (?)
				AND uker.MBDESC like '%kc%'
				GROUP BY
					BRANCH
				` + sortQuery + `
				LIMIT ? OFFSET ?`

		repo.db.DB.Table("verifikasi").
			Select(`COUNT(*) 'SumData'`).
			Where(`REGION = ?`, request.REGION).
			Where(`MAINBR in (?)`, mainbrs).
			Where(`(created_at >= ? AND created_at <= ?)`, request.StartDate, lib.FixEndDate(request.EndDate)).
			Where(`indikasi_fraud IN (?)`, fraudIndication).
			Find(&SumData)

		queryPagination = `
				SELECT
					COUNT(DISTINCT BRANCH) AS pagination 
				FROM 
					dwh_branch 
				WHERE 
					REGION = ?
				AND MAINBR in (?)
				AND MBDESC like '%kc%'`

		rows, err = repo.db.DB.Raw(query, request.StartDate, lib.FixEndDate(request.EndDate), fraudIndication, request.StartDate, lib.FixEndDate(request.EndDate), fraudIndication, request.StartDate, lib.FixEndDate(request.EndDate), fraudIndication, request.StartDate, lib.FixEndDate(request.EndDate), fraudIndication, request.StartDate, lib.FixEndDate(request.EndDate), fraudIndication, request.REGION, mainbrs, strconv.Itoa(request.Limit), strconv.Itoa(request.Offset)).Rows()
		defer rows.Close()

		errPagination = repo.db.DB.Raw(queryPagination, request.REGION, mainbrs).Scan(&totalRows).Error

	} else if filter4 {
		fmt.Println("filter4")

		if request.Sort == "DESC" {
			sortQuery = `ORDER BY TOTALVERIFICATION DESC, uker.BRANCH`
		} else {
			sortQuery = `ORDER BY TOTALVERIFICATION ASC, uker.BRANCH`
		}

		fraudIndication := strings.Split(request.FraudIndication, ",")

		// add by panji 23/11/2023
		mainbrs := strings.Split(request.MAINBR, ",")
		branches := strings.Split(request.BRANCH, ",")

		query = `SELECT
					uker.REGION,
					uker.RGDESC,
					uker.MAINBR,
					uker.MBDESC,
					uker.BRANCH,
					uker.BRDESC,
					COALESCE ( total_verification.TOTAL, 0 ) AS TOTALVERIFICATION,
					COALESCE ( brc.TOTAL, 0 ) AS TOTALBRC,
					COALESCE (total_non_weakness.TOTAL, 0) AS TOTALNONWEAKNESS,
					COALESCE (total_weakness.TOTAL, 0) AS TOTALNWEAKNESS,
					COALESCE (total_status_perbaikan_on_progress.TOTAL, 0) AS TOTALPERBAIKANONPROGRESS,
					COALESCE (total_status_perbaikan_done.TOTAL, 0) AS TOTALPERBAIKANDONE
				FROM
					dwh_branch uker
					LEFT JOIN (
					SELECT
						REGION,
						MAINBR,
						BRANCH,
						COUNT( id ) AS TOTAL 
					FROM
						verifikasi v
					WHERE
						(created_at >= ? AND created_at <= ?) 
						` + queryStatus + `
						AND v.indikasi_fraud IN ?
						AND v.deleted != 1
					GROUP BY
						BRANCH 
					) total_verification ON total_verification.BRANCH = uker.BRANCH
					LEFT JOIN (
					SELECT
						kelolaan.REGION,
						kelolaan.MAINBR,
						kelolaan.BRANCH,
						COUNT( kelolaan.REGION ) AS TOTAL 
					FROM
						(
						SELECT
							REGION,
							MAINBR,
							BRANCH,
							COUNT( pn ) AS TOTAL 
						FROM
							uker_kelolaan_user 
						GROUP BY
							BRANCH,
							pn 
						) kelolaan 
					GROUP BY
						kelolaan.BRANCH 
					) brc ON brc.BRANCH = uker.BRANCH
					LEFT JOIN (
						SELECT 
							v.id,
							REGION,
							MAINBR,
							BRANCH,
							COUNT(v.id) as TOTAL
						FROM
							verifikasi v
						WHERE
							v.perbaikan = '0'
							` + queryStatus + `
						AND (created_at >= ? AND created_at <= ?)
						AND v.indikasi_fraud IN ?
						AND v.deleted != 1
						GROUP BY 
							BRANCH
					) total_non_weakness ON total_non_weakness.BRANCH = uker.BRANCH
					LEFT JOIN (
						SELECT 
							v.id,
							REGION,
							MAINBR,
							BRANCH,
							COUNT(v.id) as TOTAL
						FROM
							verifikasi v
						WHERE
							v.perbaikan = '1'
							` + queryStatus + `
						AND (created_at >= ? AND created_at <= ?)
						AND v.indikasi_fraud IN ?
						AND v.deleted != 1
						GROUP BY BRANCH
					) total_weakness ON total_weakness.BRANCH = uker.BRANCH
					LEFT JOIN (
						SELECT 
							v.id,
							REGION,
							MAINBR,
							BRANCH,
							COUNT(v.id) as TOTAL
						FROM
							verifikasi v
						JOIN 
							verifikasi_pic_tindak_lanjut vptl 
							ON v.id = vptl.verifikasi_id 
						WHERE
							vptl.status = '1'
							` + queryStatus + `
						AND (created_at >= ? AND created_at <= ?)
						AND v.indikasi_fraud IN ?
						AND v.deleted != 1
						GROUP BY BRANCH
					) total_status_perbaikan_on_progress ON total_status_perbaikan_on_progress.BRANCH = uker.BRANCH
					LEFT JOIN (
						SELECT 
							v.id,
							REGION,
							MAINBR,
							BRANCH,
							COUNT(v.id) as TOTAL
						FROM
							verifikasi v
						JOIN 
							verifikasi_pic_tindak_lanjut vptl 
							ON v.id = vptl.verifikasi_id 
						WHERE
							vptl.status = '2'
							` + queryStatus + `
						AND (v.created_at >= ? AND v.created_at <= ?)
						AND v.indikasi_fraud IN ?
						AND v.deleted != 1
						GROUP BY 
							BRANCH
					) total_status_perbaikan_done ON total_status_perbaikan_done.BRANCH = uker.BRANCH
				WHERE
					uker.REGION = ?
				AND uker.MAINBR in (?)
				AND uker.BRANCH in (?)
				GROUP BY
					BRANCH
				` + sortQuery + `
				LIMIT ? OFFSET ?`

		repo.db.DB.Table("verifikasi").
			Select(`COUNT(*) 'SumData'`).
			Where(`REGION = ?`, request.REGION).
			Where(`MAINBR in (?)`, mainbrs).
			Where(`BRANCH in (?)`, branches).
			Where(`(created_at >= ? AND created_at <= ?)`, request.StartDate, lib.FixEndDate(request.EndDate)).
			Where(`indikasi_fraud IN (?)`, fraudIndication).
			Find(&SumData)

		queryPagination = `
				SELECT
					COUNT(DISTINCT BRANCH) AS pagination 
				FROM 
					dwh_branch 
				WHERE 
					REGION = ?
				AND MAINBR in (?)
				AND BRANCH in (?)`

		rows, err = repo.db.DB.Raw(query, request.StartDate, lib.FixEndDate(request.EndDate), fraudIndication, request.StartDate, lib.FixEndDate(request.EndDate), fraudIndication, request.StartDate, lib.FixEndDate(request.EndDate), fraudIndication, request.StartDate, lib.FixEndDate(request.EndDate), fraudIndication, request.StartDate, lib.FixEndDate(request.EndDate), fraudIndication, request.REGION, mainbrs, branches, strconv.Itoa(request.Limit), strconv.Itoa(request.Offset)).Rows()
		defer rows.Close()

		errPagination = repo.db.DB.Raw(queryPagination, request.REGION, mainbrs, branches).Scan(&totalRows).Error

	}

	//QUERY PAGINATION
	// repo.logger.Zap.Info("briefing-queryPagination-activity-unknown", queryPagination)

	if errPagination != nil {
		return responses, SumData, totalRows, err
	}

	// repo.logger.Zap.Info("verifikasi-report-uker-query-activity-unknown", query)

	// repo.logger.Zap.Info("verifikasi-report-uker-rows-activity-unknown", rows)

	if err != nil {
		return responses, SumData, totalRows, err
	}

	response := models.VerificationFilterReportByUkerResponseNull{}
	for rows.Next() {
		_ = rows.Scan(
			&response.REGION,
			&response.RGDESC,
			&response.MAINBR,
			&response.MBDESC,
			&response.BRANCH,
			&response.BRDESC,
			&response.TOTALVERIFICATION,
			&response.TOTALBRC,
			&response.TOTALNONWEAKNESS,
			&response.TOTALNWEAKNESS,
			&response.TOTALPERBAIKANONPROGRESS,
			&response.TOTALPERBAIKANDONE,
		)
		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, SumData, totalRows, err
	}

	return responses, SumData, totalRows, err
}

// uker complete
func (repo VerifikasiRepository) VerificationReportFilterByUkerComplete(request *models.VerificationFilterReportByUkerRequest) (responses []models.VerificationFilterByUkerReportCompleteResponse, totalRows int64, err error) {
	totalType := request.TotalType
	db := repo.db.DB

	query := db.Model(&responses).
		Select(`
			verifikasi.id 'id', 
			verifikasi.created_at 'date', 
			verifikasi.created_at 'date', 
			verifikasi.BRANCH 'BRANCH', 
			verifikasi.BRDESC 'BRDESC', 
			activity.name 'activity', 
			product.product 'product', 
			verifikasi.risk_issue 'risk_issue', 
			verifikasi.risk_indicator 'materi',
			verifikasi.perbaikan 'is_required_fixing',
			verifikasi.perbaikan 'is_required_fixing',
			COALESCE(verifikasi_pic_tindak_lanjut.status, '-') 'fixing_status'`).
		Joins("LEFT JOIN activity ON activity.kode_activity = verifikasi.activity_id").
		Joins("LEFT JOIN product ON product.id = verifikasi.product_id").
		Joins("LEFT JOIN verifikasi_pic_tindak_lanjut ON verifikasi_pic_tindak_lanjut.verifikasi_id = verifikasi.id").
		Where("(verifikasi.created_at >= ? AND verifikasi.created_at <= ?)", request.StartDate, lib.FixEndDate(request.EndDate)).
		Where("verifikasi.deleted != 1")

	queryPagination := db.Table("verifikasi").
		Select(`COUNT(*) as totalRows`).
		Joins("LEFT JOIN activity a ON a.kode_activity = verifikasi.activity_id").
		Joins("LEFT JOIN product p ON p.id = verifikasi.product_id").
		Joins("LEFT JOIN verifikasi_pic_tindak_lanjut ON verifikasi_pic_tindak_lanjut.verifikasi_id = verifikasi.id").
		Where("(verifikasi.created_at >= ? AND verifikasi.created_at <= ?)", request.StartDate, lib.FixEndDate(request.EndDate)).
		Where("verifikasi.deleted != 1")

	if totalType == "weakness" {
		query = query.Where(`verifikasi.perbaikan = 1`)
		queryPagination = queryPagination.Where(`verifikasi.perbaikan = 1`)
	} else if totalType == "non_weakness" {
		query = query.Where(`verifikasi.perbaikan = 0`)
		queryPagination = queryPagination.Where(`verifikasi.perbaikan = 0`)
	} else if totalType == "status_done" {
		// 0 = backlog, 1 = on progress, 2 = done
		query = query.Where(`verifikasi_pic_tindak_lanjut.status = 2`)
		queryPagination = queryPagination.Where(`verifikasi_pic_tindak_lanjut.status = 2`)
	} else if totalType == "status_onprogress" {
		// 0 = backlog, 1 = on progress, 2 = done
		query = query.Where(`verifikasi_pic_tindak_lanjut.status = 1`)
		queryPagination = queryPagination.Where(`verifikasi_pic_tindak_lanjut.status = 1`)
	}

	if request.REGION != "all" {
		query = query.Where(`verifikasi.REGION = ?`, request.REGION)
		queryPagination = queryPagination.Where(`verifikasi.REGION = ?`, request.REGION)
		// queryPagination = queryPagination.Where(`verifikasi.REGION = ?`, request.REGION)
	}

	if request.REGION != "all" && request.MAINBR != "all" {
		query = query.Where(`verifikasi.MAINBR = ?`, request.MAINBR)
		queryPagination = queryPagination.Where(`verifikasi.MAINBR = ?`, request.MAINBR)
		// queryPagination = queryPagination.Where(`verifikasi.MAINBR = ?`, request.MAINBR)
	}

	if request.REGION != "all" && request.MAINBR != "all" && request.BRANCH != "all" {
		// if request.REGION == "" || request.MAINBR == "" {
		// 	query = query.Where(`verifikasi.BRANCH = ?`, request.BRANCH)
		// }
		branches := strings.Split(request.BRANCH, ",")
		if len(branches) > 1 {
			fmt.Println("Masuk 1")
			query = query.Where("verifikasi.BRANCH in (?)", branches)
			queryPagination = queryPagination.Where("verifikasi.BRANCH in (?)", branches)
		} else {
			fmt.Println("Masuk 2")
			query = query.Where(`verifikasi.BRANCH = ?`, request.BRANCH)
			queryPagination = queryPagination.Where(`verifikasi.BRANCH = ?`, request.BRANCH)
		}
		// queryPagination = queryPagination.Where(`verifikasi.BRANCH = ?`, request.BRANCH)
	}

	if request.FraudIndication != "0,1" || request.FraudIndication != "" {
		fraudIndication := strings.Split(request.FraudIndication, ",")

		query = query.Where(`verifikasi.indikasi_fraud in (?)`, fraudIndication)
		queryPagination = queryPagination.Where(`verifikasi.indikasi_fraud in (?)`, fraudIndication)
	}

	query = query.
		Order("verifikasi.id").
		Limit(request.Limit).
		Offset(request.Offset).
		Find(&responses)

	queryPagination = queryPagination.Count(&totalRows)

	//QUERY PAGINATION
	repo.logger.Zap.Info("briefing-queryPagination-activity-unknown", queryPagination)

	repo.logger.Zap.Info("verifikasi-query-activity-unknown", query)

	if err != nil {
		return responses, totalRows, err
	}

	return responses, totalRows, err
}

// end of REPORT BY UKER

// indikator fraud
func (repo VerifikasiRepository) VerifikasiReportByFraudIndicatorFilter(request *models.VerificationFilterReportByUkerRequest) (responses []models.VerificationFilterReportByFraudIndicatorResponse, totalRows int64, err error) {
	sortQuery := ""
	db := repo.db.DB
	query := db
	queryPagination := db

	//bank wide
	filter1 := request.ReportType == "fraud_indication" &&
		request.REGION == "all" &&
		request.MAINBR == "all" && request.BRANCH == "all"

	filter2 := request.ReportType == "fraud_indication" &&
		request.REGION != "all" &&
		request.MAINBR == "all" && request.BRANCH == "all"

	filter3 := request.ReportType == "fraud_indication" &&
		request.REGION != "all" &&
		request.MAINBR != "all" && request.BRANCH == "all"

	filter4 := request.ReportType == "fraud_indication" &&
		request.REGION != "all" &&
		request.MAINBR != "all" && request.BRANCH != "all"

	if filter1 {
		fmt.Printf("Filter1")
		// var modelsSubQuery models.VerifikasiList
		subQuery := db.Table("verifikasi").
			Select(`
				verifikasi.REGION,
				COUNT(verifikasi.id) AS TOTAL
			`).
			Where("(verifikasi.created_at >= ? AND verifikasi.created_at <= ?)", request.StartDate, lib.FixEndDate(request.EndDate)).
			Where("verifikasi.deleted != 1").
			Where("verifikasi.indikasi_fraud = 1").
			Group("verifikasi.REGION")

		if request.Activity != "all" {
			subQuery = subQuery.Where(`verifikasi.activity_id = ?`, request.Activity)
		}

		if request.Product != "all" {
			subQuery = subQuery.Where(`verifikasi.product_id = ?`, request.Product)
		}

		if request.Sort == "DESC" {
			sortQuery = `TOTALFRAUD DESC, dwh_branch.REGION`
		} else {
			sortQuery = `TOTALFRAUD ASC, dwh_branch.REGION`
		}

		query = db.Model(&responses).
			Select(`
						dwh_branch.REGION,
						dwh_branch.RGDESC,
						dwh_branch.MAINBR,
						dwh_branch.MBDESC,
						dwh_branch.BRANCH,
						dwh_branch.BRDESC,
						COALESCE ( total_fraud.TOTAL, 0 ) AS 'TOTALFRAUD'`).
			Joins("LEFT JOIN (?) total_fraud ON total_fraud.REGION = dwh_branch.REGION", subQuery).
			Where(`dwh_branch.BRUNIT = 'B'`).
			Where(`dwh_branch.MAINBR = dwh_branch.BRANCH`).
			Where(`
						(
							dwh_branch.BRDESC LIKE "kanwil%" 
							OR dwh_branch.BRDESC = "Jkt KCK" 
						)`).
			Order(sortQuery).
			Limit(request.Limit).
			Offset(request.Offset).
			Find(&responses)

		queryPagination = db.Table("dwh_branch").
			Where("BRUNIT = 'B'").
			Where("MAINBR = BRANCH").
			Where(`(
				BRDESC LIKE "kanwil%" 
				OR BRDESC = "Jkt KCK" 
			)`).
			Count(&totalRows)

	} else if filter2 {
		fmt.Printf("Filter2")

		subQuery := db.Table("verifikasi").
			Select(`
						verifikasi.REGION,
						verifikasi.MAINBR,
						COUNT( verifikasi.id ) 'TOTAL'`).
			Where(`( verifikasi.created_at >= ? AND verifikasi.created_at <= ?)`, request.StartDate, lib.FixEndDate(request.EndDate)).
			Where("verifikasi.deleted != 1").
			Where(`verifikasi.indikasi_fraud = 1`).
			Group(`verifikasi.MAINBR`)

		if request.Activity != "all" {
			subQuery = subQuery.Where(`verifikasi.activity_id = ?`, request.Activity)
		}

		if request.Product != "all" {
			subQuery = subQuery.Where(`verifikasi.product_id = ?`, request.Product)
		}

		if request.Sort == "DESC" {
			sortQuery = `TOTALFRAUD DESC, dwh_branch.MAINBR`
		} else {
			sortQuery = `TOTALFRAUD ASC, dwh_branch.MAINBR`
		}

		query = db.Model(&responses).
			Select(`
						dwh_branch.REGION,
						dwh_branch.RGDESC,
						dwh_branch.MAINBR,
						dwh_branch.MBDESC,
						dwh_branch.BRANCH,
						dwh_branch.BRDESC,
						COALESCE ( total_fraud.TOTAL, 0 ) AS 'TOTALFRAUD'`).
			Joins("LEFT JOIN (?) total_fraud ON total_fraud.REGION = dwh_branch.REGION AND total_fraud.MAINBR = dwh_branch.MAINBR", subQuery).
			Where(`dwh_branch.REGION = ?`, request.REGION).
			Where(`dwh_branch.MBDESC NOT LIKE '%kanwil%'`).
			Group(`dwh_branch.MBDESC`).
			Order(sortQuery).
			Limit(request.Limit).
			Offset(request.Offset).
			Find(&responses)

		queryPagination = db.Table("dwh_branch").
			Where("REGION = ?", request.REGION).
			Where(`MBDESC NOT LIKE '%kanwil%'`).
			Group(`MBDESC`).
			Count(&totalRows)

	} else if filter3 {
		fmt.Printf("Filter3")

		// var modelsSubQuery models.VerifikasiList
		subQuery := db.Table("verifikasi").
			Select(`
						verifikasi.REGION,
						verifikasi.MAINBR,
						verifikasi.BRANCH,
						COUNT( verifikasi.id ) 'TOTAL'`).
			Where(`( verifikasi.created_at >= ? AND verifikasi.created_at <= ?)`, request.StartDate, lib.FixEndDate(request.EndDate)).
			Where("verifikasi.deleted != 1").
			Where(`verifikasi.indikasi_fraud = 1`).
			Group(`verifikasi.BRANCH`)

		if request.Activity != "all" {
			subQuery = subQuery.Where(`verifikasi.activity_id = ?`, request.Activity)
		}

		if request.Product != "all" {
			subQuery = subQuery.Where(`verifikasi.product_id = ?`, request.Product)
		}

		if request.Sort == "DESC" {
			sortQuery = `TOTALFRAUD DESC, dwh_branch.BRANCH`
		} else {
			sortQuery = `TOTALFRAUD ASC, dwh_branch.BRANCH`
		}

		mainbrs := strings.Split(request.BRANCH, ",")

		query = db.Model(&responses).
			Select(`
						dwh_branch.REGION,
						dwh_branch.RGDESC,
						dwh_branch.MAINBR,
						dwh_branch.MBDESC,
						dwh_branch.BRANCH,
						dwh_branch.BRDESC,
						COALESCE ( total_fraud.TOTAL, 0 ) AS 'TOTALFRAUD'`).
			Joins("LEFT JOIN (?) total_fraud ON total_fraud.REGION = dwh_branch.REGION AND total_fraud.MAINBR = dwh_branch.MAINBR AND total_fraud.BRANCH = dwh_branch.BRANCH", subQuery).
			Where(`dwh_branch.REGION = ?`, request.REGION).
			Where(`dwh_branch.MAINBR in (?)`, mainbrs).
			Order(sortQuery).
			Limit(request.Limit).
			Offset(request.Offset).
			Find(&responses)

		queryPagination = db.Table("dwh_branch").
			Where("REGION = ?", request.REGION).
			Where("MAINBR in (?)", mainbrs).
			Count(&totalRows)

	} else if filter4 {
		fmt.Println("Filter 4 uwu")
		// var modelsSubQuery models.VerifikasiList
		subQuery := db.Table("verifikasi").
			Select(`
						verifikasi.REGION,
						verifikasi.MAINBR,
						verifikasi.BRANCH,
						COUNT( verifikasi.id ) 'TOTAL'`).
			Where(`( verifikasi.created_at >= ? AND verifikasi.created_at <= ?)`, request.StartDate, lib.FixEndDate(request.EndDate)).
			Where("verifikasi.deleted != 1").
			Where(`verifikasi.indikasi_fraud = 1`).
			Group(`verifikasi.BRANCH`)

		if request.Activity != "all" {
			subQuery = subQuery.Where(`verifikasi.activity_id = ?`, request.Activity)
		}

		if request.Product != "all" {
			subQuery = subQuery.Where(`verifikasi.product_id = ?`, request.Product)
		}

		if request.Sort == "DESC" {
			sortQuery = `TOTALFRAUD DESC, dwh_branch.BRANCH`
		} else {
			sortQuery = `TOTALFRAUD ASC, dwh_branch.BRANCH`
		}

		mainbrs := strings.Split(request.BRANCH, ",")
		branches := strings.Split(request.BRANCH, ",")

		query = db.Model(&responses).
			Select(`
						dwh_branch.REGION,
						dwh_branch.RGDESC,
						dwh_branch.MAINBR,
						dwh_branch.MBDESC,
						dwh_branch.BRANCH,
						dwh_branch.BRDESC,
						COALESCE ( total_fraud.TOTAL, 0 ) AS 'TOTALFRAUD'`).
			Joins("LEFT JOIN (?) total_fraud ON total_fraud.REGION = dwh_branch.REGION AND total_fraud.MAINBR = dwh_branch.MAINBR AND total_fraud.BRANCH = dwh_branch.BRANCH", subQuery).
			Where(`dwh_branch.REGION = ?`, request.REGION).
			Where(`dwh_branch.MAINBR in (?)`, mainbrs).
			Where("dwh_branch.BRANCH in (?)", branches). //add by panji 24/11/2023
			Order(sortQuery).
			Limit(request.Limit).
			Offset(request.Offset).
			Find(&responses)

		queryPagination = db.Table("dwh_branch").
			Where("REGION = ?", request.REGION).
			Where("MAINBR in (?)", mainbrs).
			Where("BRANCH in (?)", branches).
			Count(&totalRows)
	}

	//QUERY PAGINATION
	repo.logger.Zap.Info("verifikasi-queryPagination-activity-unknown", queryPagination)

	repo.logger.Zap.Info("verifikasi-query-activity-unknown", query)

	if err != nil {
		return responses, totalRows, err
	}

	return responses, totalRows, err
}

// indikator fraud complete
func (repo VerifikasiRepository) VerificationReportFilterByFraudIndicatorComplete(request *models.VerificationFilterReportByUkerRequest) (responses []models.VerifikasiFilterReportCompleteResponse, totalRows int, err error) {
	// var rows *sql.Rows
	var errPagination error
	db := repo.db.DB

	query := db.Model(&responses).
		Select(`
			verifikasi.id, 
			verifikasi.created_at as date, 
			verifikasi.BRANCH, 
			verifikasi.BRDESC, 
			a.name as activity_name, 
			p.product as product_name, 
			verifikasi.risk_issue, 
			verifikasi.risk_indicator as judul_materi
		`).
		Joins("LEFT JOIN product p ON p.id = verifikasi.product_id").
		Joins("LEFT JOIN activity a ON a.kode_activity = verifikasi.activity_id").
		Where("verifikasi.deleted != 1").
		Where("verifikasi.indikasi_fraud = 1").
		Where("(verifikasi.created_at >= ? AND verifikasi.created_at <= ?)", request.StartDate, lib.FixEndDate(request.EndDate))

	queryPagination := db.Table("verifikasi").
		Select(`COUNT(verifikasi.id) as totalRows`).
		Joins("LEFT JOIN product p ON p.id = verifikasi.product_id").
		Joins("LEFT JOIN activity a ON a.kode_activity = verifikasi.activity_id").
		Where("verifikasi.deleted != 1").
		Where("verifikasi.indikasi_fraud = 1").
		Where("(verifikasi.created_at >= ? AND verifikasi.created_at <= ?)", request.StartDate, lib.FixEndDate(request.EndDate))

	if request.BRANCH != "all" {

		branches := strings.Split(request.BRANCH, ",")
		if len(branches) > 1 {
			fmt.Println("Masuk 1")
			query = query.Where("verifikasi.BRANCH in (?)", branches)
			queryPagination = queryPagination.Where("verifikasi.BRANCH in (?)", branches)
		} else {
			fmt.Println("Masuk 2")
			query = query.Where("verifikasi.BRANCH = ?", request.BRANCH)
			queryPagination = queryPagination.Where("verifikasi.BRANCH = ?", request.BRANCH)
		}
	}

	if request.Activity != "all" {
		query = query.Where("verifikasi.activity_id = ?", request.Activity)
		queryPagination = queryPagination.Where("verifikasi.activity_id = ?", request.Activity)
	}

	if request.Product != "all" {
		query = query.Where("verifikasi.product_id = ?", request.Product)
		queryPagination = queryPagination.Where("verifikasi.product_id = ?", request.Product)
	}

	query.Limit(request.Limit).
		Offset(request.Offset).
		Find(&responses)

	queryPagination.Find(&totalRows)

	repo.logger.Zap.Info("briefing-queryPagination-activity-unknown", queryPagination)

	repo.logger.Zap.Info("verifikasi-query-activity-unknown", query)

	if errPagination != nil {
		return responses, totalRows, err
	}

	if err != nil {
		return responses, totalRows, err
	}

	fmt.Println("reponses repo")
	fmt.Println(responses)

	return responses, totalRows, err

}

// add 23 Feb 2023 By Panji
func (repo VerifikasiRepository) VerifikasiReportMateriList(request *models.VerifikasiMateriRequest) (responses []models.VerifikasiDetailMateriResponseNull, totalRows int, totalData int, err error) {
	var rows *sql.Rows
	fileId := strings.Split(request.Id, ",")
	query := `SELECT id, filename,nama_lampiran, path FROM risk_indicator_map_files rimf WHERE id IN ?`

	rows, err = repo.db.DB.Raw(query, fileId).Rows()
	defer rows.Close()

	repo.logger.Zap.Info("verifikasi-query-activity-unknown", query)

	repo.logger.Zap.Info("verifikasi-rows-activity-unknown", rows)
	if err != nil {
		return responses, totalRows, totalData, err
	}

	response := models.VerifikasiDetailMateriResponseNull{}
	for rows.Next() {
		_ = rows.Scan(
			&response.ID,
			&response.Filename,
			&response.NamaLampiran,
			&response.Path,
		)
		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, totalRows, totalData, err
	}

	return responses, totalRows, totalData, err
}

func (repo VerifikasiRepository) VerificationReportUkerByAllActivity(request *models.VerificationFilterReportByUkerRequest) (responses models.VerifikasiReportAllUker, totalRows int, err error) {
	db := repo.db.DB
	query := db
	queryTotal := db
	queryTotalActivity := db
	queryPagination := db
	sortQuery := ""

	fraudIndication := strings.Split(request.FraudIndication, ",")

	branches := strings.Split(request.BRANCH, ",")

	filter1 := request.ReportType == "all_activity" &&
		request.REGION == "all" &&
		request.MAINBR == "all" &&
		request.BRANCH == "all"

	filter2 := request.ReportType == "all_activity" &&
		request.REGION != "all" &&
		request.MAINBR == "all" &&
		request.BRANCH == "all"

	filter3 := request.ReportType == "all_activity" &&
		request.REGION != "all" &&
		request.MAINBR != "all" &&
		request.BRANCH == "all"

	filter4 := request.ReportType == "all_activity" &&
		request.REGION != "all" &&
		request.MAINBR != "all" &&
		request.BRANCH != "all"

	if filter1 {
		fmt.Println("filter1")

		var tempVerificationAllActivity []models.TempVerificationAllActivity
		// tempMap := make(map[string]interface{})
		var tempArrayMap []map[string]interface{}

		if request.Sort == "DESC" {
			sortQuery = `uker.BRANCH, TOTAL DESC`
		} else {
			sortQuery = `uker.BRANCH, TOTAL`
		}

		var jumlahUker int64
		// get activity
		db.Table("activity").
			Select("id, kode_activity, name").Find(&responses.ActivityList)

		// get data per 1 activity
		for _, response := range responses.ActivityList {
			query =
				db.Table("dwh_branch uker").
					Select(`
					uker.REGION,
					uker.RGDESC,
					uker.MAINBR,
					uker.MBDESC,
					uker.BRANCH,
					uker.BRDESC,
					COALESCE(v.WEAKNESS,0) WEAKNESS,
					COALESCE(v1.TOTAL ,0) TOTAL
				`).
					Joins(`
					LEFT JOIN (
						SELECT COUNT(*) WEAKNESS, REGION
						FROM
							verifikasi v
						WHERE
							v.activity_id = ?
							AND v.indikasi_fraud IN (?)
							AND v.perbaikan = 1
							AND (v.created_at >= ? AND v.created_at <= ?)
							AND v.deleted != 1
						GROUP BY v.REGION 
					) v ON v.REGION = uker.REGION 
				`, response.ID, fraudIndication, request.StartDate, lib.FixEndDate(request.EndDate)).
					Joins(`
					LEFT JOIN (SELECT
						COUNT(*) TOTAL, REGION
					FROM
						verifikasi v
					WHERE
						v.activity_id = ?
						AND v.indikasi_fraud IN (?)
						AND (v.created_at >= ? AND v.created_at <= ?)
						AND v.deleted != 1
					GROUP BY v.REGION
					) v1 ON v1.REGION = uker.REGION
				`, response.ID, fraudIndication, request.StartDate, lib.FixEndDate(request.EndDate)).
					Where(`uker.BRUNIT = ? AND uker.MAINBR = uker.BRANCH AND (uker.BRDESC LIKE ? OR uker.BRDESC = ?)`, "B", "kanwil%", "Jkt KCK").
					Group(`
					uker.REGION,
					uker.RGDESC,
					uker.MAINBR,
					uker.MBDESC,
					uker.BRANCH,
					uker.BRDESC
				`).
					Order(sortQuery).
					Limit(request.Limit).
					Offset(request.Offset).
					Find(&tempVerificationAllActivity).
					Count(&jumlahUker)

			for i := range tempVerificationAllActivity {
				if len(tempArrayMap) < int(jumlahUker) {
					activityMap := map[string]interface{}{
						"REGION": tempVerificationAllActivity[i].REGION,
						"RGDESC": tempVerificationAllActivity[i].RGDESC,
						"MAINBR": tempVerificationAllActivity[i].MAINBR,
						"MBDESC": tempVerificationAllActivity[i].MBDESC,
						"BRANCH": tempVerificationAllActivity[i].BRANCH,
						"BRDESC": tempVerificationAllActivity[i].BRDESC,
					}

					tempArrayMap = append(tempArrayMap, activityMap)
				}

				// activityMap := map[string]interface{}{
				// 	"REGION": tempVerificationAllActivity[i].REGION,
				// 	"RGDESC": tempVerificationAllActivity[i].RGDESC,
				// 	"MAINBR": tempVerificationAllActivity[i].MAINBR,
				// 	"MBDESC": tempVerificationAllActivity[i].MBDESC,
				// 	"BRANCH": tempVerificationAllActivity[i].BRANCH,
				// 	"BRDESC": tempVerificationAllActivity[i].BRDESC,
				// }

				// tempArrayMap = append(tempArrayMap, activityMap)
			}

			// grand total for every Regional Office
			for i := range tempVerificationAllActivity {
				if i < len(tempArrayMap) {
					weakness := tempVerificationAllActivity[i].WEAKNESS
					total := tempVerificationAllActivity[i].TOTAL

					if i < len(tempVerificationAllActivity) {
						if total != 0 {
							percent := (float64(weakness) / float64(total)) * 100

							tempArrayMap[i][response.Name] = percent
						} else {
							tempArrayMap[i][response.Name] = 0.0
						}
					}
				}
			}

			//query nation wide
			var totalActivity int
			var total int

			fmt.Println("Jumlah Uker ==================================>", jumlahUker)
			// put nasional on the last index array
			if len(tempArrayMap) == int(jumlahUker) {
				activityMap := map[string]interface{}{
					"REGION": "NASIONAL",
					"RGDESC": "NASIONAL",
					"MAINBR": "NASIONAL",
					"MBDESC": "NASIONAL",
					"BRANCH": "NASIONAL",
					"BRDESC": "NASIONAL",
				}

				tempArrayMap = append(tempArrayMap, activityMap)
			}

			queryTotalActivity = db.Table("verifikasi v").
				Select(`COUNT(*) as count`).
				Where(`v.activity_id = ?
						AND v.indikasi_fraud IN (?)
						AND (v.created_at >= ? AND v.created_at <= ?)
						AND v.deleted != 1`, response.ID, fraudIndication, request.StartDate, lib.FixEndDate(request.EndDate)).
				Scan(&totalActivity)

			queryTotal = db.Table("verifikasi v").
				Select(`COUNT(*) as count`).
				Where(`v.indikasi_fraud IN (?)
						AND (v.created_at >= ? AND v.created_at <= ?)
						AND v.deleted != 1`, fraudIndication, request.StartDate, lib.FixEndDate(request.EndDate)).
				Scan(&total)

			fmt.Println("total - total", total)
			fmt.Println("total - Activity", totalActivity)

			if len(tempArrayMap) > int(jumlahUker) {
				if total != 0 {
					percent := (float64(totalActivity) / float64(total)) * 100
					tempArrayMap[jumlahUker][response.Name] = percent
				} else {
					tempArrayMap[jumlahUker][response.Name] = 0.0
				}
			}
		}

		// get data color for categorize number
		db.Table("ref_report_type_color_category ref").
			Select("ref.name, ref.condition, ref.bgcolor, ref.txcolor, ref.description").Where("ref.report_type_id = 1").Find(&responses.Colors)

		responses.Data = tempArrayMap

		queryPagination = db.Table(`dwh_branch`).
			Select(`COUNT(*) AS pagination`).
			Where(`BRUNIT = 'B'`).
			Where(`(
				BRDESC LIKE "kanwil%" 
				OR BRDESC = "Jkt KCK" 
			)`).Scan(&totalRows)

		repo.logger.Zap.Info("verifikasi-query-total-verifikasi_all_activity", queryTotal)
		repo.logger.Zap.Info("verifikasi-query-total_weakness-verifikasi_all_activity", queryTotalActivity)
	} else if filter2 {
		fmt.Println("filter2")

		var tempVerificationAllActivity []models.TempVerificationAllActivity
		// tempMap := make(map[string]interface{})
		var tempArrayMap []map[string]interface{}

		if request.Sort == "DESC" {
			sortQuery = `uker.BRANCH, TOTAL DESC`
		} else {
			sortQuery = `uker.BRANCH, TOTAL`
		}

		var jumlahUker int64

		// get activity
		db.Table("activity").
			Select("id, kode_activity, name").Find(&responses.ActivityList)

		// get data per 1 activity
		for _, response := range responses.ActivityList {
			query =
				db.Table("dwh_branch uker").
					Select(`
					uker.REGION,
					uker.RGDESC,
					uker.MAINBR,
					uker.MBDESC,
					uker.BRANCH,
					uker.BRDESC,
					COALESCE(v.WEAKNESS,0) WEAKNESS,
					COALESCE(v1.TOTAL ,0) TOTAL
				`).
					Joins(`
					LEFT JOIN (
						SELECT COUNT(*) WEAKNESS, REGION, MAINBR
						FROM
							verifikasi v
						WHERE
							v.activity_id = ?
							AND v.indikasi_fraud IN (?)
							AND v.perbaikan = 1
							AND (v.created_at >= ? AND v.created_at <= ?)
							AND v.deleted != 1
						GROUP BY v.REGION, v.MAINBR 
					) v ON v.REGION = uker.REGION AND v.MAINBR = uker.MAINBR
				`, response.ID, fraudIndication, request.StartDate, lib.FixEndDate(request.EndDate)).
					Joins(`
					LEFT JOIN (SELECT
						COUNT(*) TOTAL, REGION, MAINBR
					FROM
						verifikasi v
					WHERE
						v.activity_id = ?
						AND v.indikasi_fraud IN (?)
						AND (v.created_at >= ? AND v.created_at <= ?)
						AND v.deleted != 1
					GROUP BY v.REGION, v.MAINBR
					) v1 ON v1.REGION = uker.REGION AND v1.MAINBR
				`, response.ID, fraudIndication, request.StartDate, lib.FixEndDate(request.EndDate)).
					Where(`
					uker.REGION = ?
					AND uker.MBDESC = uker.BRDESC
				`, request.REGION).
					Group(`
					uker.REGION,
					uker.RGDESC,
					uker.MAINBR,
					uker.MBDESC,
					uker.BRANCH,
					uker.BRDESC
				`).
					Order(sortQuery).
					Limit(request.Limit).
					Offset(request.Offset).
					Find(&tempVerificationAllActivity).
					Count(&jumlahUker)

			fmt.Println("uker lenght =>", len(tempVerificationAllActivity))

			for i := range tempVerificationAllActivity {
				if len(tempArrayMap) < int(jumlahUker) {
					fmt.Println("masuk sini")
					activityMap := map[string]interface{}{
						"REGION": tempVerificationAllActivity[i].REGION,
						"RGDESC": tempVerificationAllActivity[i].RGDESC,
						"MAINBR": tempVerificationAllActivity[i].MAINBR,
						"MBDESC": tempVerificationAllActivity[i].MBDESC,
						"BRANCH": tempVerificationAllActivity[i].BRANCH,
						"BRDESC": tempVerificationAllActivity[i].BRDESC,
					}

					tempArrayMap = append(tempArrayMap, activityMap)
				}
			}

			// grand total for every Regional Office
			for i := range tempVerificationAllActivity {
				if i < len(tempArrayMap) {
					weakness := tempVerificationAllActivity[i].WEAKNESS
					total := tempVerificationAllActivity[i].TOTAL

					if i < len(tempVerificationAllActivity) {
						if total != 0 {
							percent := (float64(weakness) / float64(total)) * 100

							tempArrayMap[i][response.Name] = percent
						} else {
							tempArrayMap[i][response.Name] = 0.0
						}
					}
				}
			}

			//query nation wide
			var totalActivity int
			var total int

			// put nasional on the last index array
			if len(tempArrayMap) == int(jumlahUker) {
				activityMap := map[string]interface{}{
					"REGION": "NASIONAL",
					"RGDESC": "NASIONAL",
					"MAINBR": "NASIONAL",
					"MBDESC": "NASIONAL",
					"BRANCH": "NASIONAL",
					"BRDESC": "NASIONAL",
				}

				tempArrayMap = append(tempArrayMap, activityMap)
			}

			queryTotalActivity = db.Table("verifikasi v").
				Select(`COUNT(*) as count`).
				Where(`v.activity_id = ?
						AND v.indikasi_fraud IN (?)
						AND (v.created_at >= ? AND v.created_at <= ?)
						AND v.deleted != 1`, response.ID, fraudIndication, request.StartDate, lib.FixEndDate(request.EndDate)).
				Scan(&totalActivity)

			queryTotal = db.Table("verifikasi v").
				Select(`COUNT(*) as count`).
				Where(`v.indikasi_fraud IN (?)
						AND (v.created_at >= ? AND v.created_at <= ?)
						AND v.deleted != 1`, fraudIndication, request.StartDate, lib.FixEndDate(request.EndDate)).
				Scan(&total)

			fmt.Println("total - total", total)
			fmt.Println("total - Activity", totalActivity)

			if len(tempArrayMap) > int(jumlahUker) {
				if total != 0 {
					percent := (float64(totalActivity) / float64(total)) * 100
					tempArrayMap[jumlahUker][response.Name] = percent
				} else {
					tempArrayMap[jumlahUker][response.Name] = 0.0
				}
			}
		}

		// get data color for categorize number
		db.Table("ref_report_type_color_category ref").
			Select("ref.name, ref.condition, ref.bgcolor, ref.txcolor, ref.description").Where("ref.report_type_id = 1").Find(&responses.Colors)

		responses.Data = tempArrayMap

		queryPagination = db.Table(`dwh_branch`).
			Select(`COUNT(*) AS pagination`).
			// Where(`BRUNIT = 'B'`).
			Where(`
				REGION = ?
				AND MBDESC = BRDESC
			`, request.REGION).Scan(&totalRows)

		repo.logger.Zap.Info("verifikasi-query-total-verifikasi_all_activity", queryTotal)
		repo.logger.Zap.Info("verifikasi-query-total_weakness-verifikasi_all_activity", queryTotalActivity)
	} else if filter3 {
		fmt.Println("filter3")

		var tempVerificationAllActivity []models.TempVerificationAllActivity
		// tempMap := make(map[string]interface{})
		var tempArrayMap []map[string]interface{}

		if request.Sort == "DESC" {
			sortQuery = `uker.BRANCH, TOTAL DESC`
		} else {
			sortQuery = `uker.BRANCH, TOTAL`
		}

		var jumlahUker int64

		// get activity
		db.Table("activity").
			Select("id, kode_activity, name").Find(&responses.ActivityList)

		// get data per 1 activity
		for _, response := range responses.ActivityList {
			query =
				db.Table("dwh_branch uker").
					Select(`
					uker.REGION,
					uker.RGDESC,
					uker.MAINBR,
					uker.MBDESC,
					uker.BRANCH,
					uker.BRDESC,
					COALESCE(v.WEAKNESS,0) WEAKNESS,
					COALESCE(v1.TOTAL ,0) TOTAL
				`).
					Joins(`
				LEFT JOIN (
					SELECT COUNT(*) WEAKNESS, REGION, MAINBR, BRANCH
					FROM
						verifikasi v
					WHERE
						v.activity_id = ?
						AND v.indikasi_fraud IN (?)
						AND v.perbaikan = 1
						AND (v.created_at >= ? AND v.created_at <= ?)
						AND v.deleted != 1
					GROUP BY v.REGION, v.MAINBR, v.BRANCH
				) v ON v.REGION = uker.REGION AND v.MAINBR = uker.MAINBR AND v.BRANCH = uker.BRANCH`, response.ID, fraudIndication, request.StartDate, lib.FixEndDate(request.EndDate)).
					Joins(`
				LEFT JOIN (SELECT
					COUNT(*) TOTAL, REGION, MAINBR, BRANCH
				FROM
					verifikasi v
				WHERE
					v.activity_id = ?
					AND v.indikasi_fraud IN (?)
					AND (v.created_at >= ? AND v.created_at <= ?)
					AND v.deleted != 1
				GROUP BY v.REGION, v.MAINBR, v.BRANCH
				) v1 ON v1.REGION = uker.REGION AND v1.MAINBR AND v1.BRANCH = uker.BRANCH`, response.ID, fraudIndication, request.StartDate, lib.FixEndDate(request.EndDate)).
					Where(`
				uker.REGION = ?
				AND uker.MAINBR = ?`, request.REGION, request.MAINBR).
					Group(`
					uker.REGION,
					uker.RGDESC,
					uker.MAINBR,
					uker.MBDESC,
					uker.BRANCH,
					uker.BRDESC
				`).
					Order(sortQuery).
					Limit(request.Limit).
					Offset(request.Offset).
					Find(&tempVerificationAllActivity).
					Count(&jumlahUker)

			for i := range tempVerificationAllActivity {
				if len(tempArrayMap) < int(jumlahUker) {
					activityMap := map[string]interface{}{
						"REGION": tempVerificationAllActivity[i].REGION,
						"RGDESC": tempVerificationAllActivity[i].RGDESC,
						"MAINBR": tempVerificationAllActivity[i].MAINBR,
						"MBDESC": tempVerificationAllActivity[i].MBDESC,
						"BRANCH": tempVerificationAllActivity[i].BRANCH,
						"BRDESC": tempVerificationAllActivity[i].BRDESC,
					}

					tempArrayMap = append(tempArrayMap, activityMap)
				}
			}

			// grand total for every Regional Office
			for i := range tempVerificationAllActivity {
				if i < len(tempArrayMap) {
					weakness := tempVerificationAllActivity[i].WEAKNESS
					total := tempVerificationAllActivity[i].TOTAL

					if i < len(tempVerificationAllActivity) {
						if total != 0 {
							percent := (float64(weakness) / float64(total)) * 100

							tempArrayMap[i][response.Name] = percent
						} else {
							tempArrayMap[i][response.Name] = 0
						}
					}
				}
			}

			//query nation wide
			var totalActivity int
			var total int

			// put nasional on the last index array
			if len(tempArrayMap) == int(jumlahUker) {
				activityMap := map[string]interface{}{
					"REGION": "NASIONAL",
					"RGDESC": "NASIONAL",
					"MAINBR": "NASIONAL",
					"MBDESC": "NASIONAL",
					"BRANCH": "NASIONAL",
					"BRDESC": "NASIONAL",
				}

				tempArrayMap = append(tempArrayMap, activityMap)
			}

			queryTotalActivity = db.Table("verifikasi v").
				Select(`COUNT(*) as count`).
				Where(`v.activity_id = ?
						AND v.indikasi_fraud IN (?)
						AND (v.created_at >= ? AND v.created_at <= ?)
						AND v.deleted != 1`, response.ID, fraudIndication, request.StartDate, lib.FixEndDate(request.EndDate)).
				Scan(&totalActivity)

			queryTotal = db.Table("verifikasi v").
				Select(`COUNT(*) as count`).
				Where(`v.indikasi_fraud IN (?)
						AND (v.created_at >= ? AND v.created_at <= ?)
						AND v.deleted != 1`, fraudIndication, request.StartDate, lib.FixEndDate(request.EndDate)).
				Scan(&total)

			fmt.Println("total - total", total)
			fmt.Println("total - Activity", totalActivity)

			if len(tempArrayMap) > int(jumlahUker) {
				if total != 0 {
					percent := (float64(totalActivity) / float64(total)) * 100
					tempArrayMap[jumlahUker][response.Name] = percent
				} else {
					tempArrayMap[jumlahUker][response.Name] = 0.0
				}
			}
		}

		// get data color for categorize number
		db.Table("ref_report_type_color_category ref").
			Select("ref.name, ref.condition, ref.bgcolor, ref.txcolor, ref.description").Where("ref.report_type_id = 1").Find(&responses.Colors)

		responses.Data = tempArrayMap

		queryPagination = db.Table(`dwh_branch`).
			Select(`COUNT(*) AS pagination`).
			// Where(`BRUNIT = 'B'`).
			Where(`
				REGION = ?
				AND MAINBR = ?
			`, request.REGION, request.MAINBR).Scan(&totalRows)

		repo.logger.Zap.Info("verifikasi-query-total-verifikasi_all_activity", queryTotal)
		repo.logger.Zap.Info("verifikasi-query-total_weakness-verifikasi_all_activity", queryTotalActivity)
	} else if filter4 { //add by panji 24/11/2023 penyesuaian untuk rpt login brc
		fmt.Println("filter4 uwu")

		var tempVerificationAllActivity []models.TempVerificationAllActivity
		// tempMap := make(map[string]interface{})
		var tempArrayMap []map[string]interface{}

		if request.Sort == "DESC" {
			sortQuery = `uker.BRANCH, TOTAL DESC`
		} else {
			sortQuery = `uker.BRANCH, TOTAL`
		}

		var jumlahUker int64

		// get activity
		db.Table("activity").
			Select("id, kode_activity, name").Find(&responses.ActivityList)

		// get data per 1 activity
		for _, response := range responses.ActivityList {
			query =
				db.Table("dwh_branch uker").
					Select(`
					uker.REGION,
					uker.RGDESC,
					uker.MAINBR,
					uker.MBDESC,
					uker.BRANCH,
					uker.BRDESC,
					COALESCE(v.WEAKNESS,0) WEAKNESS,
					COALESCE(v1.TOTAL ,0) TOTAL
				`).
					Joins(`
					LEFT JOIN (
						SELECT COUNT(*) WEAKNESS, REGION, MAINBR, BRANCH
						FROM
							verifikasi v
						WHERE
							v.activity_id = ?
							AND v.indikasi_fraud IN (?)
							AND v.perbaikan = 1
							AND (v.created_at >= ? AND v.created_at <= ?)
							AND v.deleted != 1
						GROUP BY v.REGION, v.MAINBR, v.BRANCH
					) v ON v.REGION = uker.REGION AND v.MAINBR = uker.MAINBR AND v.BRANCH = uker.BRANCH
				`, response.ID, fraudIndication, request.StartDate, lib.FixEndDate(request.EndDate)).
					Joins(`
					LEFT JOIN (SELECT
						COUNT(*) TOTAL, REGION, MAINBR, BRANCH
					FROM
						verifikasi v
					WHERE
						v.activity_id = ?
						AND v.indikasi_fraud IN (?)
						AND (v.created_at >= ? AND v.created_at <= ?)
						AND v.deleted != 1
					GROUP BY v.REGION, v.MAINBR, v.BRANCH
					) v1 ON v1.REGION = uker.REGION AND v1.MAINBR AND v1.BRANCH = uker.BRANCH
				`, response.ID, fraudIndication, request.StartDate, lib.FixEndDate(request.EndDate)).
					Where(`
					uker.REGION = ?
					AND uker.MAINBR = ?
					AND uker.BRANCH in (?)
				`, request.REGION, request.MAINBR, branches).
					Group(`
					uker.REGION,
					uker.RGDESC,
					uker.MAINBR,
					uker.MBDESC,
					uker.BRANCH,
					uker.BRDESC
				`).
					Order(sortQuery).
					Limit(request.Limit).
					Offset(request.Offset).
					Find(&tempVerificationAllActivity).
					Count(&jumlahUker)

			for i := range tempVerificationAllActivity {
				if len(tempArrayMap) < int(jumlahUker) {
					activityMap := map[string]interface{}{
						"REGION": tempVerificationAllActivity[i].REGION,
						"RGDESC": tempVerificationAllActivity[i].RGDESC,
						"MAINBR": tempVerificationAllActivity[i].MAINBR,
						"MBDESC": tempVerificationAllActivity[i].MBDESC,
						"BRANCH": tempVerificationAllActivity[i].BRANCH,
						"BRDESC": tempVerificationAllActivity[i].BRDESC,
					}

					tempArrayMap = append(tempArrayMap, activityMap)
				}
			}

			// grand total for every Regional Office
			for i := range tempVerificationAllActivity {
				if i < len(tempArrayMap) {
					weakness := tempVerificationAllActivity[i].WEAKNESS
					total := tempVerificationAllActivity[i].TOTAL

					if i < len(tempVerificationAllActivity) {
						if total != 0 {
							percent := (float64(weakness) / float64(total)) * 100

							tempArrayMap[i][response.Name] = percent
						} else {
							tempArrayMap[i][response.Name] = 0
						}
					}
				}
			}

			//query nation wide
			var totalActivity int
			var total int

			// put nasional on the last index array
			if len(tempArrayMap) == int(jumlahUker) {
				activityMap := map[string]interface{}{
					"REGION": "NASIONAL",
					"RGDESC": "NASIONAL",
					"MAINBR": "NASIONAL",
					"MBDESC": "NASIONAL",
					"BRANCH": "NASIONAL",
					"BRDESC": "NASIONAL",
				}

				tempArrayMap = append(tempArrayMap, activityMap)
			}

			queryTotalActivity = db.Table("verifikasi v").
				Select(`COUNT(*) as count`).
				Where(`v.activity_id = ?
						AND v.indikasi_fraud IN (?)
						AND (v.created_at >= ? AND v.created_at <= ?)
						AND v.deleted != 1`, response.ID, fraudIndication, request.StartDate, lib.FixEndDate(request.EndDate)).
				Scan(&totalActivity)

			queryTotal = db.Table("verifikasi v").
				Select(`COUNT(*) as count`).
				Where(`v.indikasi_fraud IN (?)
						AND (v.created_at >= ? AND v.created_at <= ?)
						AND v.deleted != 1`, fraudIndication, request.StartDate, lib.FixEndDate(request.EndDate)).
				Scan(&total)

			fmt.Println("total - total", total)
			fmt.Println("total - Activity", totalActivity)

			if len(tempArrayMap) > int(jumlahUker) {
				if total != 0 {
					percent := (float64(totalActivity) / float64(total)) * 100
					tempArrayMap[jumlahUker][response.Name] = percent
				} else {
					tempArrayMap[jumlahUker][response.Name] = 0.0
				}
			}
		}

		// get data color for categorize number
		db.Table("ref_report_type_color_category ref").
			Select("ref.name, ref.condition, ref.bgcolor, ref.txcolor, ref.description").Where("ref.report_type_id = 1").Find(&responses.Colors)

		responses.Data = tempArrayMap

		queryPagination = db.Table(`dwh_branch`).
			Select(`COUNT(*) AS pagination`).
			Where(`BRUNIT = 'B'`).
			Where(`
				REGION = ?
				AND MAINBR = ?
				AND BRANCH in (?)
			`, request.REGION, request.MAINBR, branches).Scan(&totalRows)

		repo.logger.Zap.Info("verifikasi-query-total-verifikasi_all_activity", queryTotal)
		repo.logger.Zap.Info("verifikasi-query-total_weakness-verifikasi_all_activity", queryTotalActivity)
	}

	repo.logger.Zap.Info("briefing-queryPagination-activity-unknown", queryPagination)
	repo.logger.Zap.Info("verifikasi-query-activity-unknown", query)

	fmt.Println("responses - resp", responses)
	fmt.Println("responses - totalRows", totalRows)

	return responses, totalRows, err
}

func (repo VerifikasiRepository) VerificationReportUkerByAllActivityComplete(request *models.VerificationFilterReportByUkerRequest) (responses []models.ResponsesAllActivityComplete, totalRows int, err error) {
	db := repo.db.DB
	query := db
	queryPagination := db
	sortQuery := ""

	fraudIndication := strings.Split(request.FraudIndication, ",")

	if request.Sort == "DESC" {
		sortQuery = `TOTAL DESC`
	} else {
		sortQuery = `TOTAL`
	}

	query = db.Table("verifikasi v").
		Select(`
			v.risk_issue, 
			COUNT(CASE WHEN v.perbaikan = 1 THEN v.risk_issue END) AS WEAKNESS,
			COUNT(v.risk_issue) AS TOTAL
		`).
		Where(`v.REGION = ?`, request.REGION).
		Where(`v.MAINBR = ?`, request.MAINBR).
		Where(`v.BRANCH = ?`, request.BRANCH).
		Where(`v.activity_id = ?`, request.Activity).
		Where(`v.indikasi_fraud IN (?)`, fraudIndication).
		Where(`(v.created_at >= ? AND v.created_at <= ?)`, request.StartDate, lib.FixEndDate(request.EndDate)).
		Where(`v.deleted != 1`).
		Group(`v.risk_issue`).
		Order(sortQuery).
		Limit(request.Limit).
		Offset(request.Offset).
		Find(&responses)

	queryPagination = db.Table("verifikasi v").
		Select("COUNT(DISTINCT risk_issue) as pagination").
		Where(`v.REGION = ?`, request.REGION).
		Where(`v.MAINBR = ?`, request.MAINBR).
		Where(`v.BRANCH = ?`, request.BRANCH).
		Where(`v.activity_id = ?`, request.Activity).
		Where(`v.indikasi_fraud IN (?)`, fraudIndication).
		Where(`v.created_at >= ? AND v.created_at <= ?`, request.StartDate, lib.FixEndDate(request.EndDate)).
		Where(`v.deleted != 1`).
		Group(`v.risk_issue`).
		Scan(&totalRows)

	repo.logger.Zap.Info("verifikasi-queryPagination-activity-VerificationReportUkerByAllActivityComplete", queryPagination)

	repo.logger.Zap.Info("verifikasi-query-activity-VerificationReportUkerByAllActivityComplete", query)

	return responses, totalRows, err
}

func (repo VerifikasiRepository) VerificationReportUkerByAllActivityCompleteWithRiskIssue(request *models.VerificationFilterReportByUkerRequest) (responses []map[string]interface{}, totalRows int, err error) {
	db := repo.db.DB
	query := db
	queryPagination := db
	sortQuery := ""

	fraudIndication := strings.Split(request.FraudIndication, ",")

	if request.Sort == "DESC" {
		sortQuery = `REGION DESC`
	} else {
		sortQuery = `REGION`
	}

	query = db.Table("verifikasi v").
		Select(`
			v.id,
			v.created_at date,
			v.BRANCH kode_branch,
			v.BRDESC uker,
			a.name as aktivitas,
			p.product,
			v.risk_issue,
			v.risk_indicator title
		`).
		Joins(`JOIN activity a ON a.kode_activity = v.activity_id`).
		Joins(`JOIN product p ON p.id = v.product_id`).
		Where(`v.REGION = ?`, request.REGION).
		Where(`v.MAINBR = ?`, request.MAINBR).
		Where(`v.BRANCH = ?`, request.BRANCH).
		Where(`v.risk_issue = ?`, request.RiskIssue).
		Where(`v.activity_id = ?`, request.Activity).
		Where(`v.indikasi_fraud IN (?)`, fraudIndication).
		Where(`(v.created_at >= ? AND v.created_at <= ?)`, request.StartDate, lib.FixEndDate(request.EndDate)).
		Where(`v.deleted != 1`).
		Where(`v.perbaikan = 1`).
		Order(sortQuery).
		Limit(request.Limit).
		Offset(request.Offset).
		Find(&responses)

	queryPagination = db.Table("verifikasi v").
		Select("COUNT(risk_issue) as pagination").
		Where(`v.REGION = ?`, request.REGION).
		Where(`v.MAINBR = ?`, request.MAINBR).
		Where(`v.BRANCH = ?`, request.BRANCH).
		Where(`v.risk_issue = ?`, request.RiskIssue).
		Where(`v.activity_id = ?`, request.Activity).
		Where(`v.indikasi_fraud IN (?)`, fraudIndication).
		Where(`(v.created_at >= ? AND v.created_at <= ?)`, request.StartDate, lib.FixEndDate(request.EndDate)).
		Where(`v.deleted != 1`).
		Where(`v.perbaikan = 1`).
		Group(`v.risk_issue`).
		Scan(&totalRows)

	repo.logger.Zap.Info("briefing-queryPagination-activity-unknown", queryPagination)

	repo.logger.Zap.Info("verifikasi-query-activity-unknown", query)

	return responses, totalRows, err
}

func (repo VerifikasiRepository) VerifikasiReportList(request *models.VerifikasiReportListRequest) (responses []models.VerifikasiReportListResponse, totalRows int, err error) {
	fmt.Println("Masuk Repository ->", request)

	db := repo.db.DB.Table(`report_list_verifikasi rlv`)

	sortQuery := ""

	if request.Sort == "desc" {
		// sortQuery = `verifikasi.created_at DESC`
		sortQuery = `rlv.id DESC`
	} else {
		// sortQuery = `verifikasi.created_at ASC`
		sortQuery = `rlv.id ASC`
	}

	query := db.Select(`
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
		rlv.filepath 
	`).Order(sortQuery)

	if request.NoPelaporan != "" {
		query = query.Where("rlv.no_pelaporan = ?", request.NoPelaporan)
	}

	if request.BrcUrc != "Semua" && request.BrcUrc != "" {
		query = query.Where("rlv.maker LIKE ?", fmt.Sprintf("%%%s%%", request.BrcUrc))
	}

	if request.REGION != "all" {
		query = query.Where("rlv.REGION = ?", request.REGION)
	}

	if request.MAINBR != "all" {
		mainbrs := strings.Split(request.MAINBR, ",")
		query = query.Where("rlv.MAINBR in (?)", mainbrs)
	}

	if request.BRANCH != "all" {
		branches := strings.Split(request.BRANCH, ",")
		query = query.Where("rlv.BRANCH in (?)", branches)
	}

	if request.RiskIssueID != "all" {
		query = query.Where("rlv.risk_issue_id = ?", request.RiskIssueID)
	}

	if request.RiskIndicator != "all" {
		query = query.Where("rlv.risk_indicator = ?", request.RiskIndicator)
	}

	if request.IndikasiFraud != "all" {
		query = query.Where("rlv.indikasi_fraud = ?", request.IndikasiFraud)
	}

	if request.Status != "" && request.Status != "all" {
		query = query.Where("rlv.status = ?", request.Status)
	}

	if request.StartDate != "" && lib.FixEndDate(request.EndDate) != "" {
		query = query.Where("rlv.periode BETWEEN ? AND ?", request.StartDate, lib.FixEndDate(request.EndDate))
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

	err = query.Scan(&responses).Error

	return responses, totalRows, err
}

// RptRakapitulasiBCV implements VerifikasiDefinition
func (repo VerifikasiRepository) RptRakapitulasiBCV(request *models.RptRekapitulasiBCVRequest) (responses []models.RptRekapitulasiBCVResponse, totalRows int, err error) {
	db := repo.db.DB.Table(`rpt_rekapitulasi_bcv rrb`)

	query := db.Select(`
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

	if request.REGION != "all" {
		query = query.Where("rrb.REGION = ?", request.REGION)
	}

	if request.MAINBR != "all" {
		mainbrs := strings.Split(request.MAINBR, ",")
		query = query.Where("rrb.MAINBR in (?)", mainbrs)
	}

	if request.BRANCH != "all" {
		branches := strings.Split(request.BRANCH, ",")
		query.Where("rrb.BRANCH in (?)", branches)
	}

	if request.BRC != "Semua" {
		query = query.Where("rrb.pn = ?", request.BRC)
	}

	if request.StartDate != "" && request.EndDate != "" {
		query = query.Where("rrb.Tanggal BETWEEN ? AND ?", request.StartDate, lib.FixEndDate(request.EndDate))
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

	err = query.Scan(&responses).Error

	return responses, totalRows, err
}

// RptRekomendasiRiskFromBriefing implements VerifikasiDefinition
func (repo VerifikasiRepository) RptRekomendasiRiskFromBriefing(request *models.RptRekomendasiRiskRequest) (responses []models.RptRekomendasiRiskResponse, totalRows int, err error) {
	query := repo.db.DB

	if request.JenisData == "Risk Event" {
		query = query.Table("briefing_materis bm").
			Select(`
				bm.judul_materi 'risk_event',
				"Briefing" as 'module',
				COUNT(*) as 'count'`).
			Joins("JOIN briefing b ON bm.briefing_id = b.id").
			Where(`bm.risk_issue_code = "Other"`).
			Where("(b.created_at >= ? AND b.created_at <= ?)", request.StartDate, lib.FixEndDate(request.EndDate)).
			Group("bm.judul_materi")
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

	query.Scan(&responses)

	return responses, totalRows, err
}

// RptRekomendasiRiskFromCoaching implements VerifikasiDefinition
func (repo VerifikasiRepository) RptRekomendasiRiskFromCoaching(request *models.RptRekomendasiRiskRequest) (responses []models.RptRekomendasiRiskResponse, totalRows int, err error) {
	query := repo.db.DB

	if request.JenisData == "Risk Event" {
		query = query.Table("coaching_activity ca").
			Select(`
				ca.risk_issue 'risk_event',
				"Coaching" as 'module',
				COUNT(*) as 'count'`).
			Joins("JOIN coaching c ON ca.coaching_id = c.id").
			Where(`ca.risk_issue_code = "Other"`).
			Where("(c.created_at >= ? AND c.created_at <= ?)", request.StartDate, lib.FixEndDate(request.EndDate)).
			Group("ca.risk_issue")
	}

	if request.JenisData == "Risk Indicator" {
		query = query.Table("coaching_activity ca").
			Select(`
				ca.judul_materi 'risk_indicator',
				"Coaching" as 'module',
				COUNT(*) as 'count'`).
			Joins("JOIN coaching c ON ca.coaching_id = c.id").
			Where(`ca.risk_indicator_id = 0`).
			Where("(c.created_at >= ? AND c.created_at <= ?)", request.StartDate, lib.FixEndDate(request.EndDate)).
			Group("ca.judul_materi")
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

	query.Scan(&responses)

	return responses, totalRows, err
}

// RptRekomendasiRiskFromVerifikasi implements VerifikasiDefinition
func (repo VerifikasiRepository) RptRekomendasiRiskFromVerifikasi(request *models.RptRekomendasiRiskRequest) (responses []models.RptRekomendasiRiskResponse, totalRows int, err error) {
	query := repo.db.DB

	if request.JenisData == "Risk Event" {
		query = query.Table("verifikasi v").
			Select(`
				v.risk_issue_other 'risk_event',
				"Verifikasi" as 'module',
				COUNT(*) as 'count'`).
			Where(`v.risk_issue = "Other"`).
			Where("(v.created_at >= ? AND v.created_at <= ?)", request.StartDate, lib.FixEndDate(request.EndDate)).
			Group("v.risk_issue_other")
	}

	if request.JenisData == "Risk Indicator" {
		query = query.Table("verifikasi v").
			Select(`
				v.risk_indicator_other 'risk_indicator',
				"Verifikasi" as 'module',
				COUNT(*) as 'count'`).
			Where(`v.risk_issue = "Other"`).
			Where("(v.created_at >= ? AND v.created_at <= ?)", request.StartDate, lib.FixEndDate(request.EndDate)).
			Group("v.risk_issue_other")
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

	query.Scan(&responses)

	return responses, totalRows, err
}

// RptRekomendasiRisk implements VerifikasiDefinition
func (repo VerifikasiRepository) RptRekomendasiRisk(request *models.RptRekomendasiRiskRequest) (responses []models.RptRekomendasiRiskResponse, totalRows int, err error) {
	query := repo.db.DB.Table(`rpt_rekomendasi_risk rrr`)

	switch request.JenisData {
	case "Risk Event":
		query.Select(`
				rrr.nama_risk 'risk_event',
				rrr.modul 'module',
				SUM(rrr.jumlah) 'count'
			`).
			Where(`rrr.Tanggal BETWEEN ? AND ?`, request.StartDate, lib.FixEndDate(request.EndDate)).
			Where("rrr.jenis_risk = ?", strings.ToLower(request.JenisData)).
			Group(`rrr.nama_risk, rrr.modul`).
			Order(`rrr.jumlah DESC`)
	case "Risk Indicator":
		query.Select(`
				rrr.nama_risk 'risk_indicator',
				rrr.modul 'module,
				SUM(rrr.jumlah) 'count'
			`).
			Where(`rrr.Tanggal BETWEEN ? AND ?`, request.StartDate, lib.FixEndDate(request.EndDate)).
			Where("rrr.jenis_risk = ?", strings.ToLower(request.JenisData)).
			Group(`rrr.nama_risk, rrr.modul`).
			Order(`rrr.jumlah DESC`)
	default:
		err = fmt.Errorf("Jenis Risk Not Selected")
	}

	var count int64
	query.Count(&count)
	totalRows = int(count)

	err = query.Limit(request.Limit).Offset(request.Offset).Scan(&responses).Error

	return responses, totalRows, err
}

// ValidasiVerifikasi implements VerifikasiDefinition
func (repo VerifikasiRepository) ValidasiVerifikasi(request *models.ValidasiVerifikasiRequest) (responses []models.ValidasiVerifikasiResponse, totalRows int, err error) {
	query := repo.db.DB

	query = query.Table("verifikasi v").
		Select(`
			v.id 'verifikasi_id',
			v.no_pelaporan,
			v.BRDESC 'unit_kerja',
			a.name 'aktivitas',
			v.risk_issue,
			v.maker_desc,
			vq.id 'validasi_id',
			vq.checker 'validator_rmc',
			vq.status_validasi_rmc 'status_rmc',
			vq.signer 'validator_rrm',
			vq.status_validasi_signer 'status_signer',
			vq.approval_ord 'validator_ord',
			vq.status_validasi_ord 'status_ord'
		`).
		Joins(`LEFT JOIN activity a ON v.activity_id = a.kode_activity`).
		Joins(`JOIN verifikasi_questionner vq ON vq.verifikasi_id = v.id`).
		Where(`v.status != '01a' AND v.action != 'Draft'`).
		Where(`v.deleted != 1`).
		Order(`v.id DESC`)

	// if request.Checker != "" {
	// 	query = query.Where(`vq.checker LIKE ?`, fmt.Sprintf("%s%%", request.Checker))
	// }

	// if request.Signer != "" {
	// 	query = query.
	// 		Where(`vq.signer LIKE ?`, fmt.Sprintf("%s%%", request.Signer)).
	// 		Where(`vq.status_validasi_rmc = ?`, "01a")
	// }

	// if request.ApprovalOrd != "" {
	// 	query = query.
	// 		Where(`vq.approval_ord LIKE ?`, fmt.Sprintf("%s%%", request.ApprovalOrd)).
	// 		Where(`vq.status_validasi_rmc = ?`, "01a").
	// 		Where(`vq.status_validasi_signer = ?`, "01a")
	// }

	if request.Validator != "" {
		validator := fmt.Sprintf("%s%%", request.Validator)
		query = query.Where(`vq.checker LIKE ? OR vq.signer LIKE ? OR vq.approval_ord LIKE ?`, validator, validator, validator)
	}

	if request.Pn != "Semua" && request.Pn != "" {
		query = query.Where(`v.maker_id = ?`, request.Pn)
	}

	if request.REGION != "all" && request.REGION != "" {
		query = query.Where("v.REGION = ?", request.REGION)
	}

	if request.MAINBR != "all" && request.MAINBR != "" {
		mainbrs := strings.Split(request.MAINBR, ",")
		query = query.Where("v.MAINBR in (?)", mainbrs)
	}

	if request.BRANCH != "all" && request.BRANCH != "" {
		branches := strings.Split(request.BRANCH, ",")
		query = query.Where("v.BRANCH in (?)", branches)
	}

	if request.StartDate != "" && lib.FixEndDate(request.EndDate) != "" {
		query = query.Where("(v.created_at >= ? AND v.created_at <= ?)", request.StartDate, lib.FixEndDate(request.EndDate))
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

	query.Scan(&responses)

	return responses, totalRows, err
}

func (verif VerifikasiRepository) GetRtlIndikasiFraud(request *models.ReqRtlIndikasiFraud) (responses models.RtlIndikasiFraudResponse, totalRows int, err error) {
	query := verif.db.DB
	queryTotal := verif.db.DB
	queryTotalInputBRCURC := verif.db.DB

	query =
		query.Table("dwh_branch uker").
			Select(`
		uker.REGION,
		uker.RGDESC,
		uker.MAINBR,
		uker.MBDESC,
		uker.BRANCH,
		uker.BRDESC,
		SUM( CASE WHEN v.REGION = uker.REGION THEN 1 ELSE 0 END ) INPUTBRCURC,
		SUM( CASE WHEN v.indikasi_fraud = 1 THEN 1 ELSE 0 END) INDIKASIFRAUD,
		COUNT( vq.id ) VALIDASIORD,
		SUM( CASE WHEN vq.tindak_lanjut_ord = 'SA AIW' THEN 1 ELSE 0 END ) SAAIW,
		SUM( CASE WHEN vq.tindak_lanjut_ord = 'Hukuman Disiplin' THEN 1 ELSE 0 END ) KOORDINASIRCEO
	`).
			Joins(`
		LEFT JOIN (
			SELECT *
		FROM
			verifikasi v
		WHERE
			(v.created_at >= ? AND v.created_at <= ?) 
			AND v.deleted = 0
		) v ON v.REGION = uker.REGION
	`, request.StartDate, lib.FixEndDate(request.EndDate)).
			Joins(`
		LEFT JOIN (
			SELECT *
		FROM
			verifikasi_questionner vq
		WHERE
			vq.tindak_lanjut_ord != ''
		) vq ON vq.verifikasi_id = v.id
	`).
			Where(`uker.BRUNIT = ? AND uker.MAINBR = uker.BRANCH AND (uker.BRDESC LIKE ? OR uker.BRDESC = ?)`, "B", "kanwil%", "Jkt KCK").
			Group(`
		uker.REGION,
		uker.RGDESC,
		uker.MAINBR,
		uker.MBDESC,
		uker.BRANCH,
		uker.BRDESC
	`)

	queryTotal =
		queryTotal.Table("dwh_branch uker").
			Select(`
		uker.REGION,
		uker.RGDESC,
		uker.MAINBR,
		uker.MBDESC,
		uker.BRANCH,
		uker.BRDESC
	`).
			Where(`uker.BRUNIT = ? AND uker.MAINBR = uker.BRANCH AND (uker.BRDESC LIKE ? OR uker.BRDESC = ?)`, "B", "kanwil%", "Jkt KCK").
			Group(`
		uker.REGION,
		uker.RGDESC,
		uker.MAINBR,
		uker.MBDESC,
		uker.BRANCH,
		uker.BRDESC
	`)

	var count int64
	queryTotal.Count(&count)

	fmt.Println("count", count)

	totalRows = int(count)

	if request.Limit != 0 {
		query = query.Limit(request.Limit)
	}

	query = query.Offset(request.Offset)

	query.Scan(&responses.RtlIndikasiFraud)

	// // query get all total
	var totalVerifikasi models.TotalVerifikasi
	queryTotalInputBRCURC =
		queryTotalInputBRCURC.Table("verifikasi v").
			Select(`
					COUNT( v.id ) TOTALINPUTBRCURC,
					SUM( CASE WHEN v.indikasi_fraud = 1 THEN 1 ELSE 0 END) TOTALINDIKASIFRAUD,
					COUNT( vq.id ) TOTALVALIDASIORD,
					SUM( CASE WHEN vq.tindak_lanjut_ord = 'SA AIW' THEN 1 ELSE 0 END ) TOTALSAAIW,
					SUM( CASE WHEN vq.tindak_lanjut_ord = 'Hukuman Disiplin' THEN 1 ELSE 0 END ) TOTALKOORDINASIRCEO`).
			Where("(v.created_at >= ? AND v.created_at <= ?)", request.StartDate, lib.FixEndDate(request.EndDate)).
			Where("v.deleted = 0").
			Where("v.BRDESC != 'Kas Kanpus'").
			Joins(`
					LEFT JOIN (
						SELECT *
					FROM
						verifikasi_questionner vq
					WHERE
						vq.tindak_lanjut_ord != ''
					) vq ON vq.verifikasi_id = v.id
				`).
			Scan(&totalVerifikasi)

	verif.logger.Zap.Info("verifikasi-query-activity-GetRtlIndikasiFraud", query)

	responses.TotalVerifikasi = totalVerifikasi

	return responses, totalRows, err
}

// ValidasiVerifikasiDetailData implements VerifikasiDefinition.
func (repo VerifikasiRepository) ValidasiVerifikasiDetailData(request *models.VerifikasiReportDetailRequest) (responses models.ValidasiVerifikasiDetailResponse, err error) {
	db := repo.db.DB.Table("verifikasi v")

	err = db.Select(`
		v.id,
		v.no_pelaporan,
		v.BRANCH,
		v.BRDESC,
		v.MAINBR,
		v.MBDESC,
		v.REGION,
		v.RGDESC,
		a.name as activity_name,
		sa.name as sub_activity_name,
		p.product as product_name,
		v.risk_issue,
		v.risk_indicator,
		v.hasil_verifikasi as verification_result,
		v.sumber_data as data_source,
		v.perbaikan,
		v.indikasi_fraud,
		v.terdapat_kerugian_finansial,
		v.kunjungan_nasabah,
		v.jenis_kerugian_finansial,
		v.jumlah_perkiraan_kerugian,
		v.jenis_rekomendasi,
		v.rekomendasi_tindak_lanjut,
		v.rencana_tindak_lanjut,
		rt.risk_type,
		v.tanggal_ditemukan,
		v.tanggal_mulai_rtl,
		v.tanggal_target_selesai,
		v.ada_usulan_perbaikan`).
		Joins(`LEFT JOIN activity a ON a.kode_activity = v.activity_id`).
		Joins(`LEFT JOIN sub_activity sa ON sa.id = v.sub_activity_id`).
		Joins(`LEFT JOIN product p ON p.id = v.product_id`).
		Joins(`LEFT JOIN penyebab_kejadian_lv1 pkl1 ON pkl1.id = v.incident_cause_id`).
		Joins(`LEFT JOIN penyebab_kejadian_lv3 pkl3 ON pkl3.id = v.sub_incident_cause_id`).
		Joins(`LEFT JOIN risk_type rt ON v.risk_type_id = rt.id`).
		Where(`v.id = ?`, request.Id).
		Scan(&responses).Error

	return responses, err
}

// GetRekomendasiTindakLanjut implements VerifikasiDefinition.
func (repo VerifikasiRepository) GetRekomendasiTindakLanjut(request *models.RTLRequest) (responses []models.RTLResponses, err error) {
	db := repo.db.DB.Table("verifikasi v").
		Select(`
			ROW_NUMBER() OVER (ORDER BY v.id) AS 'no',
			v.RGDESC 'kanwil',
			v.MBDESC 'kanca',
			v.BRDESC 'uker',
			p.product 'produk',
			v.risk_issue 'risk_event',
			v.risk_indicator,
			vrc.risk_control 'kelemahan_kontrol',
			pkl.penyebab_kejadian_lv3,
			v.rencana_tindak_lanjut`).
		Joins(`LEFT JOIN verifikasi_risk_control vrc ON vrc.verifikasi_id = v.id`).
		Joins(`LEFT JOIN product p ON p.id = v.product_id`).
		Joins(`LEFT JOIN penyebab_kejadian_lv3 pkl ON pkl.id = v.sub_incident_cause_id `).
		Where(`deleted != 1`).
		Where(`v.perbaikan = 1`).
		Where(`v.product_id = ?`, request.Produk).
		Where(`v.risk_issue_id = ?`, request.RiskEvent).
		Where(`v.status = '04a' AND action = 'Selesai'`).
		Where(`v.created_at >= DATE_SUB(CURDATE(), INTERVAL 36 MONTH)`).
		Group(`v.id`)

	if request.RiskIndicator != "" || request.RiskIndicator == "0" {
		db = db.Where("v.risk_indicator_id = ?", request.RiskIndicator)
	}

	if request.RiskControl != "" {
		riskControl := strings.Split(request.RiskControl, ",")

		db = db.Where("vrc.risk_control_id in (?)", riskControl)
	}

	err = db.Scan(&responses).Error

	return responses, err
}

// VerifikasiSummaryRpt implements VerifikasiDefinition.
func (verifikasi VerifikasiRepository) VerifikasiSummaryRpt(request *models.SummaryVerifikasiRequest) (responses []models.SummaryVerifikasiResponse, totalRows int, err error) {
	tableName := ""

	if request.Kegiatan == "verifikasi" {
		tableName = "verifikasi v"

		columnMappings := map[string]string{
			"activity":       "a.name AS 'aktivitas'",
			"product":        "p.product AS 'produk'",
			"risk_event":     "v.risk_issue AS 'risk_event'",
			"risk_indicator": "v.risk_indicator AS 'risk_indicator'",
			"risk_control":   "vrc.risk_control AS 'risk_control'",
		}

		joinMappings := map[string]string{
			"activity":     "JOIN activity a ON a.id = v.activity_id",
			"product":      "JOIN product p ON p.id = v.product_id",
			"risk_control": "JOIN verifikasi_risk_control vrc ON vrc.verifikasi_id = v.id",
		}

		groupMapping := map[string]string{
			"activity":       "v.activity_id",
			"product":        "v.product_id",
			"risk_event":     "v.risk_issue",
			"risk_indicator": "v.risk_indicator",
			"risk_control":   "vrc.risk_control",
		}

		selectParts := []string{"SUM(v.jumlah_perkiraan_kerugian) AS 'jumlah'"}
		joinParts := make(map[string]bool) // To ensure joins are not repeated
		groupByParts := make(map[string]bool)

		for _, gd := range request.GroupData {
			if column, exists := columnMappings[gd.Column]; exists {
				selectParts = append(selectParts, column)
				if join, ok := joinMappings[gd.Column]; ok {
					joinParts[join] = true
				}

				if groupBy, ok := groupMapping[gd.Column]; ok {
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

		query := verifikasi.db.DB.Table(tableName).
			Select(strings.Join(selectParts, ", ")).
			Where("v.perbaikan = ?", 1).
			Where("v.terdapat_kerugian_finansial = ?", 1).
			Where("v.status = ?", "04a").
			Joins(strings.Join(joins, " ")).
			Group(strings.Join(groupBy, ", "))

		if request.REGION != "all" && request.REGION != "" {
			query = query.Where("v.REGION = ?", request.REGION)
		}

		if request.MAINBR != "all" && request.MAINBR != "" {
			mainbrs := strings.Split(request.MAINBR, ",")
			query = query.Where("v.MAINBR in (?)", mainbrs)
		}

		if request.BRANCH != "all" && request.BRANCH != "" {
			branches := strings.Split(request.BRANCH, ",")
			query = query.Where("v.BRANCH in (?)", branches)
		}

		if request.JenisKerugian != "all" {
			query = query.Where(`v.jenis_kerugian_finansial = ?`, request.JenisKerugian)
		}

		if request.StartDate != "" && request.EndDate != "" {
			query = query.Where("v.created_at BETWEEN ? AND ?", request.StartDate, lib.FixEndDate(request.EndDate))
		}

		var count int64
		query.Count(&count)

		totalRows = int(count)

		query = query.Order("jumlah DESC").Limit(request.Limit).Offset(request.Offset)

		responses = []models.SummaryVerifikasiResponse{}

		err = query.Scan(&responses).Error

		return responses, totalRows, err

	} else {
		return nil, 0, err
	}

}

func (verifikasi VerifikasiRepository) VerifikasiFrekuensiRpt(request *models.FrekuensiVerifikasiRequest) (responses []models.FrekuensiVerifikasiResponse, totalRows int, err error) {
	tableName := ""

	if request.Kegiatan == "verifikasi" {
		tableName = "verifikasi v"

		columnMappings := map[string]string{
			"activity":       "a.name AS 'aktivitas'",
			"product":        "p.product AS 'produk'",
			"risk_event":     "v.risk_issue AS 'risk_event'",
			"risk_indicator": "v.risk_indicator AS 'risk_indicator'",
			"risk_control":   "vrc.risk_control AS 'risk_control'",
		}

		joinMappings := map[string]string{
			"activity":     "JOIN activity a ON a.id = v.activity_id",
			"product":      "JOIN product p ON p.id = v.product_id",
			"risk_control": "JOIN verifikasi_risk_control vrc ON vrc.verifikasi_id = v.id",
		}

		groupMappings := map[string]string{
			"activity":       "v.activity_id",
			"product":        "v.product_id",
			"risk_event":     "v.risk_issue",
			"risk_indicator": "v.risk_indicator",
			"risk_control":   "vrc.risk_control",
		}

		selectParts := []string{"COUNT(v.id) AS 'jumlah'"}
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

		query := verifikasi.db.DB.Table(tableName).
			Select(strings.Join(selectParts, ", ")).
			Where("v.status = ? AND lower(action) = ?", "04a", "selesai").
			Where("v.deleted != ?", 1).
			Joins(strings.Join(joins, " ")).
			Group(strings.Join(groupBy, ", "))

		queryPaging := verifikasi.db.DB.Table(tableName).
			Select(`Count(*)`).
			Where("v.status = ? AND lower(action) = ?", "04a", "selesai").
			Where("v.status != ?", 1).
			Joins(strings.Join(joins, " ")).
			Group(strings.Join(groupBy, ", "))

		if request.REGION != "all" && request.REGION != "" {
			query = query.Where("v.REGION = ?", request.REGION)
			queryPaging = queryPaging.Where("v.REGION = ?", request.REGION)
		}

		if request.MAINBR != "all" && request.MAINBR != "" {
			mainbrs := strings.Split(request.MAINBR, ",")

			query = query.Where("v.MAINBR in (?)", mainbrs)
			queryPaging = queryPaging.Where("v.MAINBR in (?)", mainbrs)
		}

		if request.BRANCH != "all" && request.BRANCH != "" {
			branches := strings.Split(request.BRANCH, ",")
			query = query.Where("v.BRANCH in (?)", branches)
			queryPaging = queryPaging.Where("v.BRANCH in (?)", branches)
		}

		if request.StartDate != "" && request.EndDate != "" {
			query = query.Where("v.created_at BETWEEN ? AND ?", request.StartDate, lib.FixEndDate(request.EndDate))
			queryPaging = queryPaging.Where("v.created_at BETWEEN ? AND ?", request.StartDate, lib.FixEndDate((request.EndDate)))
		}

		var Count int64
		queryPaging.Count(&Count)

		totalRows = int(Count)

		query = query.Order("jumlah DESC").Limit(request.Limit).Offset(request.Offset)

		responses = []models.FrekuensiVerifikasiResponse{}

		fmt.Println(query)

		err = query.Scan(&responses).Error

		return responses, totalRows, err

	} else {
		return nil, 0, err
	}

}
