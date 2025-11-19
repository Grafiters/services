package controllers

import (
	"riskmanagement/lib"
	models "riskmanagement/models/mcs"

	services "riskmanagement/services/mcs"

	"github.com/gin-gonic/gin"

	"gitlab.com/golang-package-library/logger"
)

type McsController struct {
	logger  logger.Logger
	service services.McsDefinition
}

func NewMcsController(McsService services.McsDefinition, logger logger.Logger) McsController {
	return McsController{
		logger:  logger,
		service: McsService,
	}
}

func (Mcs McsController) GetUker(c *gin.Context) {
	data := models.McsRequest{}

	if err := c.Bind(&data); err != nil {
		Mcs.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}
	response, err := Mcs.service.GetUker(&data)
	if err != nil {
		Mcs.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}
	if len(response) == 0 {
		Mcs.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data Tidak Ditemukan", nil)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", response)
}

func (Mcs McsController) GetPIC(c *gin.Context) {
	data := models.McsRequest{}

	if err := c.Bind(&data); err != nil {
		Mcs.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai: "+err.Error(), "")
		return
	}

	response, err := Mcs.service.GetPIC(&data)

	if err != nil {
		Mcs.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if len(response) == 0 {
		Mcs.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data Tidak Ditemukan", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", response)
}

func (Mcs McsController) GetJabatan(c *gin.Context) {
	data := models.McsRequest{}

	if err := c.Bind(&data); err != nil {
		Mcs.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai: "+err.Error(), "")
		return
	}

	response, err := Mcs.service.GetJabatan(&data)

	if err != nil {
		Mcs.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if len(response) == 0 {
		Mcs.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data Tidak Ditemukan", nil)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", response)
}

func (Mcs McsController) GetOrgeh(c *gin.Context) {
	data := models.McsRequest{}

	if err := c.Bind(&data); err != nil {
		Mcs.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai: "+err.Error(), "")
		return
	}

	response, err := Mcs.service.GetOrgeh(&data)

	if err != nil {
		Mcs.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if len(response) == 0 {
		Mcs.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data Tidak Ditemukan", nil)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", response)
}
