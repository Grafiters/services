package controllers

import (
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/aplikasi"
	services "riskmanagement/services/aplikasi"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type AplikasiController struct {
	logger  logger.Logger
	service services.AplikasiDefinition
}

func NewAplikasiController(AplikasiService services.AplikasiDefinition, logger logger.Logger) AplikasiController {
	return AplikasiController{
		service: AplikasiService,
		logger:  logger,
	}
}

func (aplikasi AplikasiController) GetAll(c *gin.Context) {
	datas, err := aplikasi.service.GetAll()
	if err != nil {
		aplikasi.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}

func (aplikasi AplikasiController) GetOne(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		aplikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := aplikasi.service.GetOne(int64(id))
	if err != nil {
		aplikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (aplikasi AplikasiController) Store(c *gin.Context) {
	data := models.AplikasiRequest{}

	if err := c.Bind(&data); err != nil {
		aplikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}
	fmt.Println(data)
	if err := aplikasi.service.Store(&data); err != nil {
		aplikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", true)
}

func (aplikasi AplikasiController) Update(c *gin.Context) {
	data := models.AplikasiRequest{}

	if err := c.Bind(&data); err != nil {
		aplikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	if err := aplikasi.service.Update(&data); err != nil {
		aplikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", data)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}

func (aplikasi AplikasiController) Delete(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)

	if err != nil {
		aplikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	if err := aplikasi.service.Delete(int64(id)); err != nil {
		aplikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Data berhasil dihapus", true)
}
