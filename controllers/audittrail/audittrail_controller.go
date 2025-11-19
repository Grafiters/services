package controller

import (
	"database/sql"
	"riskmanagement/lib"
	models "riskmanagement/models/audittrail"
	services "riskmanagement/services/audittrail"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type AuditTrailController struct {
	logger   logger.Logger
	services services.AuditTrailDefinition
}

func NewAuditTrailController(
	services services.AuditTrailDefinition,
	logger logger.Logger,
) AuditTrailController {
	return AuditTrailController{
		logger:   logger,
		services: services,
	}
}

func (audit AuditTrailController) Store(c *gin.Context) {
	request := models.AuditTrail{}

	if err := c.Bind(&request); err != nil {
		audit.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai :"+err.Error(), request)
		return
	}

	status, err := audit.services.Store(models.AuditTrail{
		ID:          request.ID,
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
		IpAddress:   lib.GetIPClient(c.ClientIP()),
		Lokasi:      request.Lokasi,
	})

	if err != nil {
		audit.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	if !status {
		lib.ReturnToJson(c, 200, "400", "Data gagal diinput", false)
	}

	lib.ReturnToJson(c, 200, "200", "Input data berhasil", true)
}

func (audit AuditTrailController) GetLog(c *gin.Context) {
	requests := models.FilterAudit{}

	if err := c.Bind(&requests); err != nil {
		audit.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := audit.services.Getaudit(requests)
	if err != nil {
		audit.logger.Zap.Error(err)
	}

	if pagination.Total == 0 {
		lib.ReturnToJson(c, 200, "404", "Data Kosong", datas)
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery data berhasil", datas, pagination)
}
