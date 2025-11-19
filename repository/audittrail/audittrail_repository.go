package audittrail

import (
	"fmt"
	"math"
	"riskmanagement/lib"
	models "riskmanagement/models/audittrail"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type AuditTrailDefinition interface {
	GetLog(request models.FilterAudit) (responses []models.AuditTrailResponse, totalData int, totalRows int, err error)
	SaveLog(request *models.AuditTrail, tx *gorm.DB) (responses bool, err error)
	WithTrx(trxHandle *gorm.DB) AuditTrailRepository
}

type AuditTrailRepository struct {
	db      lib.Database
	logger  logger.Logger
	timeout time.Duration
}

func NewAuditTrailRepository(
	db lib.Database,
	logger logger.Logger,
) AuditTrailDefinition {
	return AuditTrailRepository{
		db:      db,
		logger:  logger,
		timeout: 0,
	}
}

// GetLog implements AuditTrailDefinition
func (audit AuditTrailRepository) GetLog(request models.FilterAudit) (responses []models.AuditTrailResponse, totalData int, totalRows int, err error) {
	db := audit.db.DB

	db = db.Table(`audit_trail`).
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
		`).Where("tanggal >= ? AND tanggal <= ?", request.StartDate, lib.FixEndDate(request.EndDate))

	if request.PERNR != "" && request.PERNR != "Semua" {
		db = db.Where(`pn = ?`, request.PERNR)
	}

	if request.Aktifitas != "all" {
		db = db.Where("aktifitas = ?", request.Aktifitas)
	}

	if request.REGION != "all" {
		db = db.Where("REGION = ?", request.REGION)
	}

	if request.MAINBR != "all" {
		db = db.Where("MAINBR = ?", request.MAINBR)
	}

	if request.BRANCH != "all" {
		db = db.Where("BRANCH = ?", request.BRANCH)
	}

	var count int64
	db.Count(&count)

	totalRows = int(count)
	fmt.Println("TotalRows =>", totalRows)

	err = db.
		Limit(request.Limit).
		Offset(request.Offset).
		Find(&responses).Error

	totalPages := int(math.Ceil(float64(totalRows) / float64(request.Limit)))
	return responses, totalPages, totalRows, err
}

// SaveLog implements AuditTrailDefinition
func (AuditTrailRepository) SaveLog(request *models.AuditTrail, tx *gorm.DB) (responses bool, err error) {
	return true, tx.Save(&request).Error
}

// WithTrx implements AuditTrailDefinition
func (audit AuditTrailRepository) WithTrx(trxHandle *gorm.DB) AuditTrailRepository {
	if trxHandle == nil {
		audit.logger.Zap.Error("transaction Database not found in gin context")
		return audit
	}
	audit.db.DB = trxHandle
	return audit
}
