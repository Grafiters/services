package controller

import (
	"riskmanagement/lib"
	service "riskmanagement/services/getclientip"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type ClientIPController struct {
	logger  logger.Logger
	service service.ClientIPDefinition
}

func NewClientIPController(
	service service.ClientIPDefinition,
	logger logger.Logger,
) ClientIPController {
	return ClientIPController{
		logger:  logger,
		service: service,
	}
}

func (CI ClientIPController) ClientIP(c *gin.Context) {
	response, err := CI.service.ClientIP()

	if err != nil {
		CI.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	// if len(response) == 0 {
	// 	CI.logger.Zap.Error(err)
	// 	lib.ReturnToJson(c, 200, "400", "Data Tidak Ditemukan", false)
	// 	return
	// }

	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", response)
}
