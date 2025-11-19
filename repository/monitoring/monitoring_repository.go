package monitoring

import (
	"math"
	"riskmanagement/lib"
	models "riskmanagement/models/monitoring"
	"strings"
	"time"

	"gitlab.com/golang-package-library/logger"
)

var (
	bulan = lib.GetTimeNow("month")
	tahun = lib.GetTimeNow("year")
)

type MonitoringDefinition interface {
	GetUnitKerjaTasklist(requests models.MonitoringTasklistRequest) (responses []models.MonitoringTasklistUnitKerjaResponse, totalRow int, totalData int, err error)
	GetPekerjaTasklist(request models.MonitoringTasklistRequest) (responses []models.MonitoringTasklistPekerjaResponse, totalRow int, totalData int, err error)
}

type MonitoringRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewMonitoringRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) MonitoringDefinition {
	return MonitoringRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: 0,
	}
}

func (m MonitoringRepository) GetUnitKerjaTasklist(request models.MonitoringTasklistRequest) (responses []models.MonitoringTasklistUnitKerjaResponse, totalRow int, totalData int, err error) {
	db := m.db.DB

	subQuery := db.Table("tasklists_uker").
		Select(`tasklists_uker.BRANCH`).
		Joins(`join tasklists on tasklists_uker.tasklist_id = tasklists.id`).
		Where(`tasklists.status = "Aktif"`).
		Where(`tasklists.approval_status = "Disetujui"`).
		Where("YEAR(tasklists.created_at) = " + tahun)

	if strings.ToLower(request.Periode) == "month" {
		subQuery = subQuery.Where("MONTH(tasklists.created_at) = ? ", bulan)
	}

	if request.JenisTask != "" {
		subQuery = subQuery.Where("tasklists.task_type = ? ", request.JenisTask)
	}

	query := db.Table("uker_kelolaan_user").
		Select(`RGDESC "Kanwil",
				MBDESC "Kanca",
				BRANCH "KodeBranch",
				BRDESC "unit_kerja",
				SNAME "Pengelola",
				pn`).
		Where(`BRANCH NOT IN (?) `, subQuery)

	if request.Kanwil != "" && request.Kanwil != "all" {
		query = query.Where(`uker_kelolaan_user.REGION = ? `, request.Kanwil)
	}
	// Where(`uker_kelolaan_user.REGION = ? `, request.Kanwil)
	// Group(`pn`).Order("SNAME asc")
	query.Order("REGION asc")

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
		m.logger.Zap.Error(err)
		return responses, 0, 0, err
	}

	return responses, resultFinal, totalRow, err
}

func (m MonitoringRepository) GetPekerjaTasklist(request models.MonitoringTasklistRequest) (responses []models.MonitoringTasklistPekerjaResponse, totalRow int, totalData int, err error) {
	db := m.db.DB

	subQuery := db.Table("tasklists_uker").
		Select(`tasklists_uker.BRANCH`).
		Joins(`join tasklists on tasklists_uker.tasklist_id = tasklists.id`).
		Where(`tasklists.status = "Aktif"`).
		Where(`tasklists.approval_status = "Disetujui"`).
		Where("YEAR(tasklists.created_at) = " + tahun)

	if strings.ToLower(request.Periode) == "month" {
		subQuery = subQuery.Where("MONTH(tasklists.created_at) = ? ", bulan)
	}

	if request.JenisTask != "" {
		subQuery = subQuery.Where("tasklists.task_type = ? ", request.JenisTask)
	}

	query := db.Table("uker_kelolaan_user").
		Select(`pn,
		SNAME "Nama",
		RGDESC "Kanwil",
		GROUP_CONCAT(DISTINCT BRDESC SEPARATOR '; ') "unit_kerja"`).
		Where(`BRANCH NOT IN (?) `, subQuery)

	// Where(`uker_kelolaan_user.REGION = ? `, request.Kanwil)
	if request.Kanwil != "" && request.Kanwil != "all" {
		query = query.Where(`uker_kelolaan_user.REGION = ? `, request.Kanwil)
	}

	query = query.Group(`pn`).Group(`RGDESC`).Order("SNAME asc")

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
		m.logger.Zap.Error(err)
		return responses, 0, 0, err
	}

	return responses, resultFinal, totalRow, err
	// return responses, 0, 0, err
}
