package verifikasireportrealisasi

import (
	"riskmanagement/lib"
	models "riskmanagement/models/verifikasireportrealisasi"
	"strings"
	"time"

	"gitlab.com/golang-package-library/logger"
)

type VerifikasiReportRealisasiDefinition interface {
	ReportRealisasiKreditListFilter(request *models.ReportRealisasiKreditListRequest) (responses []models.ReportRealisasiKreditListResponseRaw, totalRows int64, err error)
	ReportRealisasiKreditListDownload(request *models.ReportRealisasiKreditListRequest) (responses []models.ReportRealisasiKreditListResponseRaw, totalRows int64, err error)
	ReportRealisasiKreditSummaryFilter(request *models.ReportRealisasiKreditSummaryRequest) (responses []models.ReportRealisasiKreditSummaryResponseRaw, totalRows int64, err error)
	ReportRealisasiKreditSummaryDownload(request *models.ReportRealisasiKreditSummaryRequest) (responses []models.ReportRealisasiKreditSummaryResponseRaw, totalRows int64, err error)
	GetAllSegmentRealisasiKredit(request *models.SegmentRealisasiKreditRequest) (responses []models.SegmentRealisasiKreditResponse, totalRows int64, err error)
}

type VerifikasiReportRealisasiRepository struct {
	db      lib.Database
	logger  logger.Logger
	timeout time.Duration
}

func NewVerfikasiReportRealisasiRepository(
	db lib.Database,
	logger logger.Logger,
) VerifikasiReportRealisasiDefinition {
	return VerifikasiReportRealisasiRepository{
		db:      db,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

func (repo VerifikasiReportRealisasiRepository) ReportRealisasiKreditListFilter(request *models.ReportRealisasiKreditListRequest) (responses []models.ReportRealisasiKreditListResponseRaw, totalRows int64, err error) {
	// totalType := request.TotalType
	db := repo.db.DB

	query := db.Table(`verifikasi_data_realisasi vdr`).
		Select(`vrk.id AS verifikasi_id,
				vrk.no_pelaporan AS no_pelaporan,
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
				vdr.status_verifikasi AS status_verifikasi`,
			`(SELECT filename FROM files WHERE id = (SELECT files_id FROM verifikasi_realisasi_lampiran WHERE verifikasi_id = vrk.id LIMIT 1) LIMIT 1) AS lampiran_name`,
			`(SELECT path FROM files WHERE id = (SELECT files_id FROM verifikasi_realisasi_lampiran WHERE verifikasi_id = vrk.id LIMIT 1) LIMIT 1) AS lampiran_path`,
		).
		Joins(`LEFT JOIN verifikasi_realisasi_kredit vrk ON vrk.id = vdr.verifikasi_id`).
		//filter status verifikasi selesai
		Where(`LOWER(vrk.action) = ?`, `selesai`)

	if request.Sort != "" {
		query = query.Order(`vdr.id ` + request.Sort)
	} else {
		query = query.Order(`vdr.id DESC`)
	}

	//Filter REGION
	if request.REGION != "" && request.REGION != "all" {
		query = query.Where(`vrk.REGION = ?`, request.REGION)
	}

	//Filter MAINBR
	if request.MAINBR != "" && request.MAINBR != "all" {
		mainbrs := strings.Split(request.MAINBR, ",")
		query = query.Where(`vrk.MAINBR in (?)`, mainbrs)
	}

	//Filter BRANCH
	if request.BRANCH != "" && request.BRANCH != "all" {
		branches := strings.Split(request.BRANCH, ",")

		query = query.Where(`vrk.BRANCH in (?)`, branches)
	}

	//Filter NoPelaporan
	if request.NoPelaporan != "" {
		query = query.Where(`vrk.no_pelaporan = ?`, request.NoPelaporan)
	}

	// Filter Produk
	if request.Product != "" && request.Product != "all" {
		query = query.Where("vrk.product_id = ?", request.Product)
	}

	//Filter Criteria
	if request.Criteria != "" && request.Criteria != "all" {
		query = query.Where(`JSON_CONTAINS(vrk.kriteria_data, ?)`, request.Criteria)
	}

	//Filter Segment
	if request.Segment != "" && request.Segment != "all" {
		query = query.Where(`JSON_UNQUOTE(JSON_EXTRACT(LOWER(vdr.data_realisasi), '$.segment')) = ?`, strings.ToLower(request.Segment))
	}

	//Filter Is Verified
	if request.StatusVerifikasi != "" && request.StatusVerifikasi != "all" {
		query = query.Where(`vdr.status_verifikasi = ?`, request.StatusVerifikasi)
	}

	if request.ButuhPerbaikan != "" && request.ButuhPerbaikan != "all" {
		query = query.Where(`vrk.butuh_perbaikan = ?`, request.ButuhPerbaikan)
	}

	//Filter Is Verified
	if request.IndikasiFraud != "" && request.IndikasiFraud != "all" {
		query = query.Where(`vrk.indikasi_fraud = ?`, request.IndikasiFraud)
	}

	if request.StartDate != "" && request.EndDate != "" {
		query = query.Where(`DATE(vrk.created_at) >= ?`, request.StartDate).Where(`DATE(vrk.created_at) <= ?`, lib.FixEndDate(request.EndDate))
	}

	query.Count(&totalRows)

	query = query.
		Group(`vdr.id`).
		Limit(request.Limit).
		Offset(request.Offset).
		Find(&responses)

	//QUERY PAGINATION
	repo.logger.Zap.Info("briefing-query-realisasi-kredit-unknown", totalRows)

	repo.logger.Zap.Info("verifikasi-query-realisasi-kredit-unknown", query)

	return responses, totalRows, err
}

func (repo VerifikasiReportRealisasiRepository) ReportRealisasiKreditListDownload(request *models.ReportRealisasiKreditListRequest) (responses []models.ReportRealisasiKreditListResponseRaw, totalRows int64, err error) {
	// totalType := request.TotalType
	db := repo.db.DB

	query := db.Table(`verifikasi_data_realisasi vdr`).
		Select(`vrk.id AS verifikasi_id,
				vrk.no_pelaporan AS no_pelaporan,
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
				vdr.status_verifikasi AS status_verifikasi`,
			`(SELECT filename FROM files WHERE id = (SELECT files_id FROM verifikasi_realisasi_lampiran WHERE verifikasi_id = vrk.id LIMIT 1) LIMIT 1) AS lampiran_name`,
			`(SELECT path FROM files WHERE id = (SELECT files_id FROM verifikasi_realisasi_lampiran WHERE verifikasi_id = vrk.id LIMIT 1) LIMIT 1) AS lampiran_path`,
		).
		Joins(`LEFT JOIN verifikasi_realisasi_kredit vrk ON vrk.id = vdr.verifikasi_id`).
		//filter status verifikasi selesai
		Where(`LOWER(vrk.action) = ?`, `selesai`)

	if request.Sort != "" {
		query = query.Order(`vdr.id ` + request.Sort)
	} else {
		query = query.Order(`vdr.id DESC`)
	}

	//Filter REGION
	if request.REGION != "" && request.REGION != "all" {
		query = query.Where(`vrk.REGION = ?`, request.REGION)
	}

	//Filter MAINBR
	if request.MAINBR != "" && request.MAINBR != "all" {
		query = query.Where(`vrk.MAINBR = ?`, request.MAINBR)
	}

	//Filter BRANCH
	if request.BRANCH != "" && request.BRANCH != "all" {
		query = query.Where(`vrk.BRANCH = ?`, request.BRANCH)
	}

	//Filter NoPelaporan
	if request.NoPelaporan != "" {
		query = query.Where(`vrk.no_pelaporan = ?`, request.NoPelaporan)
	}

	//Filter Criteria
	if request.Criteria != "" && request.Criteria != "all" {
		query = query.Where(`JSON_CONTAINS(vrk.kriteria_data, ?)`, request.Criteria)
	}

	//Filter Segment
	if request.Segment != "" && request.Segment != "all" {
		query = query.Where(`JSON_UNQUOTE(JSON_EXTRACT(LOWER(vdr.data_realisasi), '$.segment')) = ?`, strings.ToLower(request.Segment))
	}

	//Filter Is Verified
	if request.StatusVerifikasi != "" && request.StatusVerifikasi != "all" {
		query = query.Where(`vdr.status_verifikasi = ?`, request.StatusVerifikasi)
	}

	if request.ButuhPerbaikan != "" && request.ButuhPerbaikan != "all" {
		query = query.Where(`vrk.butuh_perbaikan = ?`, request.ButuhPerbaikan)
	}

	//Filter Is Verified
	if request.IndikasiFraud != "" && request.IndikasiFraud != "all" {
		query = query.Where(`vrk.indikasi_fraud = ?`, request.IndikasiFraud)
	}

	query.Count(&totalRows)

	query = query.
		Group(`vdr.id`).
		Offset(request.Offset).
		Find(&responses)

	//QUERY PAGINATION
	repo.logger.Zap.Info("briefing-query-realisasi-kredit-unknown", totalRows)

	repo.logger.Zap.Info("verifikasi-query-realisasi-kredit-unknown", query)

	return responses, totalRows, err
}

func (repo VerifikasiReportRealisasiRepository) ReportRealisasiKreditSummaryFilter(request *models.ReportRealisasiKreditSummaryRequest) (responses []models.ReportRealisasiKreditSummaryResponseRaw, totalRows int64, err error) {
	// totalType := request.TotalType
	db := repo.db.DB

	query := db.Table(`verifikasi_data_realisasi vdr`).
		Select(`vrk.product_id AS product_id,
		 		COUNT(vrk.id) AS total_verifikasi,
				vrk.product_name AS product_name, 
				vrk.created_id AS created_id, 
				vrk.created_desc AS created_desc, 
				vrk.REGION AS REGION, 
				vrk.RGDESC AS RGDESC, 
				vrk.MAINBR AS MAINBR, 
				vrk.MBDESC AS MBDESC, 
				vrk.BRANCH AS BRANCH, 
				vrk.BRDESC AS BRDESC, 
				vdr.status_verifikasi AS status_verifikasi,
				vdr.data_realisasi AS data_realisasi,
				COUNT(CASE WHEN vrk.butuh_perbaikan = 0 THEN 1 END) AS efektif,
				COUNT(CASE WHEN vrk.butuh_perbaikan = 1 THEN 1 END) AS non_efektif,
				CONCAT('[',GROUP_CONCAT(REPLACE(REPLACE(vrk.kriteria_data, ']', ''), '[', '') SEPARATOR ','),']') AS kriteria_data`).
		Joins(`LEFT JOIN verifikasi_realisasi_kredit vrk ON vdr.verifikasi_id = vrk.id`).
		//filter status verifikasi selesai
		Where(`LOWER(vrk.action) = ?`, `selesai`)

		//Sorting
	if request.Sort != "" {
		query = query.Order(`vdr.id ` + request.Sort)
	} else {
		query = query.Order(`vdr.id DESC`)
	}

	//Filter REGION
	if request.REGION != "" && request.REGION != "all" {
		query = query.Where(`vrk.REGION = ?`, request.REGION)
	}

	//Filter MAINBR
	if request.MAINBR != "" && request.MAINBR != "all" {
		mainbrs := strings.Split(request.MAINBR, ",")
		query = query.Where(`vrk.MAINBR in (?)`, mainbrs)
	}

	//Filter BRANCH
	if request.BRANCH != "" && request.BRANCH != "all" {
		branches := strings.Split(request.BRANCH, ",")

		query = query.Where(`vrk.BRANCH in (?)`, branches)
	}

	// //Filter Criteria
	// if request.Criteria != "" && request.Criteria != "all" {
	// 	query = query.Where(`JSON_CONTAINS(vrk.kriteria_data, ?)`, request.Criteria)
	// }

	if request.Product != nil {
		query = query.Where(`vrk.product_id IN (?)`, request.Product)
	}

	//Filter Is Verified
	if request.StartDate != "" && request.EndDate != "" {
		query = query.Where(`DATE(vrk.created_at) >= ?`, request.StartDate).Where(`DATE(vrk.created_at) <= ?`, request.EndDate)
	}

	if request.GroupBy != nil {
		//option for group by
		optionGroupBy := map[string]string{
			"regional-office":   "REGION",
			"branch-office":     "MAINBR",
			"unit-kerja":        "BRANCH",
			"produk":            "product_name",
			"pn-brc-urc":        "created_id",
			"status-verifikasi": "status_verifikasi",
		}

		for _, value := range request.GroupBy {
			if optionGroupBy[value] != "" {
				query = query.Group(optionGroupBy[value])
			}
		}
	}

	query.Count(&totalRows)

	query = query.
		Limit(request.Limit).
		Offset(request.Offset).
		Find(&responses)

	//QUERY PAGINATION
	repo.logger.Zap.Info("briefing-queryPagination-activity-unknown", totalRows)

	repo.logger.Zap.Info("verifikasi-query-activity-unknown", query)

	if err != nil {
		return responses, totalRows, err
	}

	return responses, totalRows, err
}

func (repo VerifikasiReportRealisasiRepository) GetAllSegmentRealisasiKredit(request *models.SegmentRealisasiKreditRequest) (responses []models.SegmentRealisasiKreditResponse, totalRows int64, err error) {
	// totalType := request.TotalType
	db := repo.db.DB

	query := db.Table(`segment_realisasi_kredit srk`).
		Select(`*`)

		//Sorting
	if request.IsActive != nil {
		query = query.Where(`srk.is_active`, request.IsActive)
	} else {
		query = query.Where(`srk.is_active`, 1)
	}

	if request.Segment != "" {
		query = query.Where(`srk.segment LIKE '%?%'`, request.Segment)
	}

	query.Count(&totalRows)

	query = query.
		Limit(request.Limit).
		Offset(request.Offset).
		Find(&responses)

	//QUERY PAGINATION
	repo.logger.Zap.Info("briefing-queryPagination-activity-unknown", totalRows)

	repo.logger.Zap.Info("verifikasi-query-activity-unknown", query)

	if query.Error != nil {
		return responses, totalRows, err
	}

	return responses, totalRows, err
}

func (repo VerifikasiReportRealisasiRepository) ReportRealisasiKreditSummaryDownload(request *models.ReportRealisasiKreditSummaryRequest) (responses []models.ReportRealisasiKreditSummaryResponseRaw, totalRows int64, err error) {
	// totalType := request.TotalType
	db := repo.db.DB

	query := db.Table(`verifikasi_data_realisasi vdr`).
		Select(`vrk.product_id AS product_id,
		 		COUNT(vrk.id) AS total_verifikasi,
				vrk.product_name AS product_name, 
				vrk.created_id AS created_id, 
				vrk.created_desc AS created_desc, 
				vrk.REGION AS REGION, 
				vrk.RGDESC AS RGDESC, 
				vrk.MAINBR AS MAINBR, 
				vrk.MBDESC AS MBDESC, 
				vrk.BRANCH AS BRANCH, 
				vrk.BRDESC AS BRDESC, 
				vdr.status_verifikasi AS status_verifikasi,
				vdr.data_realisasi AS data_realisasi,
				COUNT(CASE WHEN vrk.butuh_perbaikan = 0 THEN 1 END) AS efektif,
				COUNT(CASE WHEN vrk.butuh_perbaikan = 1 THEN 1 END) AS non_efektif,
				CONCAT('[',GROUP_CONCAT(REPLACE(REPLACE(vrk.kriteria_data, ']', ''), '[', '') SEPARATOR ','),']') AS kriteria_data`).
		Joins(`LEFT JOIN verifikasi_realisasi_kredit vrk ON vdr.verifikasi_id = vrk.id`).
		//filter status verifikasi selesai
		Where(`LOWER(vrk.action) = ?`, `selesai`)

		//Sorting
	if request.Sort != "" {
		query = query.Order(`vdr.id ` + request.Sort)
	} else {
		query = query.Order(`vdr.id DESC`)
	}

	//Filter REGION
	if request.REGION != "" && request.REGION != "all" {
		query = query.Where(`vrk.REGION = ?`, request.REGION)
	}

	//Filter MAINBR
	if request.MAINBR != "" && request.MAINBR != "all" {
		query = query.Where(`vrk.MAINBR = ?`, request.MAINBR)
	}

	//Filter BRANCH
	if request.BRANCH != "" && request.BRANCH != "all" {
		query = query.Where(`vrk.BRANCH = ?`, request.BRANCH)
	}

	// //Filter Criteria
	// if request.Criteria != "" && request.Criteria != "all" {
	// 	query = query.Where(`JSON_CONTAINS(vrk.kriteria_data, ?)`, request.Criteria)
	// }

	if request.Product != nil {

		query = query.Where(`vrk.product_id IN (?)`, request.Product)
	}

	//Filter Is Verified
	if request.StartDate != "" && request.EndDate != "" {
		query = query.Where(`DATE(vrk.created_at) >= ?`, request.StartDate).Where(`DATE(vrk.created_at) <= ?`, request.EndDate)
	}

	if request.GroupBy != nil {
		//option for group by
		optionGroupBy := map[string]string{
			"regional-office":   "REGION",
			"branch-office":     "MAINBR",
			"unit-kerja":        "BRANCH",
			"produk":            "product_name",
			"pn-brc-urc":        "created_id",
			"status-verifikasi": "status_verifikasi",
		}

		for _, value := range request.GroupBy {
			if optionGroupBy[value] != "" {
				query = query.Group(optionGroupBy[value])
			}
		}
	}

	query.Count(&totalRows)

	query = query.
		Limit(request.Limit).
		Offset(request.Offset).
		Find(&responses)

	//QUERY PAGINATION
	repo.logger.Zap.Info("briefing-queryPagination-activity-unknown", totalRows)

	repo.logger.Zap.Info("verifikasi-query-activity-unknown", query)

	if err != nil {
		return responses, totalRows, err
	}

	return responses, totalRows, err
}
