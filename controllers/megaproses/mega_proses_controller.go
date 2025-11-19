package controllers

import (
	"database/sql"
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/megaproses"
	services "riskmanagement/services/megaproses"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type MegaProsesController struct {
	logger  logger.Logger
	service services.MegaProsesDefinition
}

func NewMegaProsesController(
	MPService services.MegaProsesDefinition,
	logger logger.Logger,
) MegaProsesController {
	return MegaProsesController{
		logger:  logger,
		service: MPService,
	}
}

func (mp MegaProsesController) GetAll(c *gin.Context) {
	datas, err := mp.service.GetAll()
	if err != nil {
		mp.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}

func (mega MegaProsesController) GetAllWithPaginate(c *gin.Context) {
	requests := models.Paginate{}

	if err := c.Bind(&requests); err != nil {
		mega.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := mega.service.GetAllWithPaginate(requests)
	if err != nil {
		mega.logger.Zap.Error(err)
	}

	if pagination.Total == 0 {
		lib.ReturnToJson(c, 200, "404", "Data Kosong", nil)
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery data berhasil", datas, pagination)
}

func (mp MegaProsesController) GetOne(c *gin.Context) {
	requests := models.MegaProsesRequest{}

	if err := c.Bind(&requests); err != nil {
		mp.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mp.service.GetOne(requests.ID)
	if err != nil {
		mp.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mp MegaProsesController) Store(c *gin.Context) {
	data := models.MegaProsesRequest{}

	if err := c.Bind(&data); err != nil {
		mp.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}
	fmt.Println(data)
	if err := mp.service.Store(&data); err != nil {
		mp.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Input Data Berhasil", true)
}

func (mp MegaProsesController) Update(c *gin.Context) {
	data := models.MegaProsesRequest{}

	if err := c.Bind(&data); err != nil {
		mp.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}
	fmt.Println(data)
	if err := mp.service.Update(&data); err != nil {
		mp.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Update Data Berhasil", true)
}

func (mp MegaProsesController) Delete(c *gin.Context) {
	requests := models.MegaProsesRequest{}

	if err := c.Bind(&requests); err != nil {
		mp.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	if err := mp.service.Delete(requests.ID); err != nil {
		mp.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Data berhasil dihapus", true)
}

func (mp MegaProsesController) GetKodeMegaProses(c *gin.Context) {
	datas, err := mp.service.GetKodeMegaProses()
	if err != nil {
		mp.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if len(datas) == 0 {
		mp.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data Tidak Ditemukan !", nil)
		return
	}

	KodeMegaProses := datas[0].KodeMegaProses

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", KodeMegaProses)
}
