package download

import (
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/download"
	"strings"

	"gitlab.com/golang-package-library/logger"
)

type DownloadDefinition interface {
	CheckIsExist(request *models.DownloadRequest) (responses models.DownloadResponse, err error)
	GetDownloadUrl(id string) (responses models.DownloadUrl, err error)
	GetListDownload(request *models.ListDownloadRequest) (responses []models.ListDownloadResponse, totalRows int64, err error)
	GetReportType() (responses []models.ReportTypeResponse, err error)
	CheckRptStatus(id int64) (responses models.ListDownloadResponse, err error)
}

type DownloadRepository struct {
	db     lib.Database
	dbRaw  lib.Databases
	logger logger.Logger
}

func NewDownloadRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) DownloadDefinition {
	return DownloadRepository{
		db:     db,
		dbRaw:  dbRaw,
		logger: logger,
	}
}

// Delete implements CoachingDefinition
func (download DownloadRepository) CheckIsExist(request *models.DownloadRequest) (responses models.DownloadResponse, err error) {
	db := download.db.DB
	isExist := db.Where("report_id = ?", request.ReportId).
		Where("json_params = ?", request.JSONPARAMS).
		Find(&responses)

	if isExist.Error != nil {
		return responses, isExist.Error
	}

	return responses, err
}

// Delete implements CoachingDefinition
func (download DownloadRepository) GetDownloadUrl(id string) (responses models.DownloadUrl, err error) {
	fmt.Println("masuk repo")

	db := download.db.DB

	result := db.Where("downloadUrl = ?", id).Find(&responses)
	if result.Error != nil {
		return responses, result.Error
	}

	return responses, err
}

// GetListDownload implements DownloadDefinition.
func (download DownloadRepository) GetListDownload(request *models.ListDownloadRequest) (responses []models.ListDownloadResponse, totalRows int64, err error) {
	query := download.db.DB.Table("tbl_export_xls tex").
		Select(`
			ROW_NUMBER() OVER (ORDER BY tex.id DESC) AS 'no',
			tex.id,
			rrt.id 'report_id',
			rrt.name 'nama_laporan',
			JSON_UNQUOTE(JSON_EXTRACT(tex.json_params, '$.REGION')) as 'kanwil',
			JSON_UNQUOTE(JSON_EXTRACT(tex.json_params, '$.MAINBR')) as 'kanca',
			JSON_UNQUOTE(JSON_EXTRACT(tex.json_params, '$.BRANCH')) as 'unit_kerja',
			tex.generate_date  'periode_data',
			tex.file_desc,
			tex.filename,
			tdl.filepath,
			tex.json_params,
			tex.rpt_status 'status'
		`).
		Joins("LEFT JOIN ref_report_type rrt ON tex.report_id = rrt.id ").
		Joins("LEFT JOIN tbl_download_link tdl ON tdl.downloadUrl = tex.downloadUrl AND tdl.filename = tex.filename").
		Where(`tex.maker_id = ?`, request.MakerId).Order("tex.id DESC")

	if request.NamaLaporan != "all" && request.NamaLaporan != "" {
		query = query.Where(`rrt.id = ?`, request.NamaLaporan)
	}

	if request.Kanwil != "" {
		query = query.Where(`JSON_UNQUOTE(JSON_EXTRACT(tex.json_params, '$.REGION')) = ?`, request.Kanwil)
	}

	if request.Kanca != "" {
		mainbrs := strings.Split(request.Kanca, ",")
		query = query.Where(`JSON_UNQUOTE(JSON_EXTRACT(tex.json_params, '$.MAINBR')) in (?)`, mainbrs)
	}

	if request.UnitKerja != "" {
		branches := strings.Split(request.UnitKerja, ",")
		query = query.Where(`JSON_UNQUOTE(JSON_EXTRACT(tex.json_params, '$.BRANCH')) in (?)`, branches)
	}

	if request.PeriodeData != "" {
		query = query.Where(`DATE(tex.generate_date) = ?`, request.PeriodeData)
	}

	query = query.Count(&totalRows)

	err = query.
		Limit(request.Limit).
		Offset(request.Offset).
		Scan(&responses).Error

	return responses, totalRows, err
}

// GetReportType implements DownloadDefinition.
func (download DownloadRepository) GetReportType() (responses []models.ReportTypeResponse, err error) {
	fmt.Println("masuk repo")

	db := download.db.DB.Table("ref_report_type")

	err = db.Select(`id, name`).Scan(&responses).Error

	return responses, err
}

// CheckRptStatus implements DownloadDefinition.
func (download DownloadRepository) CheckRptStatus(id int64) (responses models.ListDownloadResponse, err error) {
	query := download.db.DB.Table("tbl_export_xls tex").
		Select(`
			ROW_NUMBER() OVER (ORDER BY tex.id DESC) AS 'no',
			tex.id,
			rrt.id 'report_id',
			rrt.name 'nama_laporan',
			JSON_UNQUOTE(JSON_EXTRACT(tex.json_params, '$.REGION')) as 'kanwil',
			JSON_UNQUOTE(JSON_EXTRACT(tex.json_params, '$.MAINBR')) as 'kanca',
			JSON_UNQUOTE(JSON_EXTRACT(tex.json_params, '$.BRANCH')) as 'unit_kerja',
			tex.generate_date  'periode_data',
			tex.file_desc,
			tex.filename,
			tdl.filepath,
			tex.json_params,
			tex.rpt_status 'status'
		`).
		Joins("LEFT JOIN ref_report_type rrt ON tex.report_id = rrt.id ").
		Joins("LEFT JOIN tbl_download_link tdl ON tdl.downloadUrl = tex.downloadUrl AND tdl.filename = tex.filename").
		Where(`tex.id = ?`, id).Order("tex.id DESC")

	err = query.Scan(&responses).Error

	return responses, err
}
