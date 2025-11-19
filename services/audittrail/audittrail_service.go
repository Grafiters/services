package audittrail

import (
	"riskmanagement/lib"
	models "riskmanagement/models/audittrail"
	audittrailrepo "riskmanagement/repository/audittrail"

	"github.com/google/uuid"
	"gitlab.com/golang-package-library/logger"
)

var (
	UUID = uuid.NewString()
)

type AuditTrailDefinition interface {
	Getaudit(request models.FilterAudit) (responses []models.AuditTrailResponse, pagination lib.Pagination, err error)
	Store(request models.AuditTrail) (responses bool, err error)
}

type AuditTrailService struct {
	db             lib.Database
	logger         logger.Logger
	audittrailrepo audittrailrepo.AuditTrailDefinition
}

func NewAuditTrailService(
	db lib.Database,
	logger logger.Logger,
	audittrailrepo audittrailrepo.AuditTrailDefinition,
) AuditTrailDefinition {
	return AuditTrailService{
		db:             db,
		logger:         logger,
		audittrailrepo: audittrailrepo,
	}
}

// Getaudit implements AuditTrailDefinition
func (audit AuditTrailService) Getaudit(request models.FilterAudit) (responses []models.AuditTrailResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	dataAudit, totalRows, totalData, err := audit.audittrailrepo.GetLog(request)

	if err != nil {
		audit.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataAudit {
		responses = append(responses, models.AuditTrailResponse{
			ID:          response.ID,
			Tanggal:     response.Tanggal,
			PN:          response.PN,
			NamaBrcUrc:  response.NamaBrcUrc,
			Kanwil:      response.Kanwil,
			Kanca:       response.Kanca,
			Uker:        response.Uker,
			NoPelaporan: response.NoPelaporan,
			Aktifitas:   response.Aktifitas,
			IpAddress:   response.IpAddress,
			Lokasi:      response.Lokasi,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err

}

// Store implements AuditTrailDefinition
func (audit AuditTrailService) Store(request models.AuditTrail) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	tx := audit.db.DB.Begin()

	insert := &models.AuditTrail{
		Tanggal:     timeNow,
		PN:          request.PN,
		NamaBrcUrc:  request.NamaBrcUrc,
		REGION:      request.REGION,
		RGDESC:      request.RGDESC,
		MAINBR:      request.MAINBR,
		MBDESC:      request.MBDESC,
		BRANCH:      request.BRANCH,
		BRDESC:      request.BRDESC,
		NoPelaporan: request.NoPelaporan,
		Aktifitas:   request.Aktifitas,
		IpAddress:   request.IpAddress,
		Lokasi:      request.Lokasi,
	}

	save, err := audit.audittrailrepo.SaveLog(insert, tx)

	if err != nil {
		tx.Rollback()
		audit.logger.Zap.Error(err)
		return false, err
	}

	tx.Commit()
	return save, err
}
