package controller

import (
	"riskmanagement/lib"
	models "riskmanagement/models/uploaddata"
	services "riskmanagement/services/uploaddata"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type UploadDataController struct {
	logger  logger.Logger
	service services.UploadDataDefinition
}

func NewUploadDataController(
	logger logger.Logger,
	service services.UploadDataDefinition,
) UploadDataController {
	return UploadDataController{
		logger:  logger,
		service: service,
	}
}

func (upload UploadDataController) UploadRiskControl(c *gin.Context) {
	request := models.UploadControlRequest{}

	if err := c.Bind(&request); err != nil {
		upload.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai :"+err.Error(), "")
		return
	}

	status, message, _ := upload.service.UploadRiskControl(request)
	if !status {
		// upload.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", message, false)
		return
	}

	lib.ReturnToJson(c, 200, "200", message, true)
}

func (upload UploadDataController) UploadRiskIndicator(c *gin.Context) {
	request := models.UploadIndicatorRequest{}

	if err := c.Bind(&request); err != nil {
		upload.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai :"+err.Error(), "")
		return
	}

	status, message, _ := upload.service.UploadRisknIndicator(request)
	if !status {
		// upload.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", message, false)
		return
	}

	lib.ReturnToJson(c, 200, "200", message, true)
}

func (upload UploadDataController) UploadRiskEvent(c *gin.Context) {
	request := models.UploadRiskIssueRequest{}

	if err := c.Bind(&request); err != nil {
		upload.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai :"+err.Error(), "")
		return
	}

	status, message, _ := upload.service.UploadRiskEvent(request)
	if !status {
		// upload.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", message, false)
		return
	}

	lib.ReturnToJson(c, 200, "200", message, true)
}
