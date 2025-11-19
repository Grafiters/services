package controllers

import (
	"database/sql"
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/majorproses"
	services "riskmanagement/services/majorproses"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type MajorProsesController struct {
	logger  logger.Logger
	service services.MajorProsesDefinition
}

func NewMajorProsesController(
	MPService services.MajorProsesDefinition,
	logger logger.Logger,
) MajorProsesController {
	return MajorProsesController{
		logger:  logger,
		service: MPService,
	}
}

func (mp MajorProsesController) GetAll(c *gin.Context) {
	datas, err := mp.service.GetAll()
	if err != nil {
		mp.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}

func (major MajorProsesController) GetAllWithPaginate(c *gin.Context) {
	requests := models.Paginate{}

	if err := c.Bind(&requests); err != nil {
		major.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := major.service.GetAllWithPaginate(requests)
	if err != nil {
		major.logger.Zap.Error(err)
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

func (mp MajorProsesController) GetOne(c *gin.Context) {
	requests := models.MajorProsesRequest{}

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

func (mp MajorProsesController) Store(c *gin.Context) {
	data := models.MajorProsesRequest{}

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

func (mp MajorProsesController) Update(c *gin.Context) {
	data := models.MajorProsesRequest{}

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

func (mp MajorProsesController) Delete(c *gin.Context) {
	requests := models.MajorProsesRequest{}

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

func (mp MajorProsesController) GetKodeMajorProses(c *gin.Context) {
	requests := models.KodeMegaProses{}

	if err := c.Bind(&requests); err != nil {
		mp.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}
	datas, err := mp.service.GetKodeMajorProses(requests)
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

	counter := datas[0].KodeMajorProses
	KodeMajorProses := datas[0].KodeMegaProses + "." + counter

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", KodeMajorProses)
}

func (mp MajorProsesController) GetMajorByMegaProses(c *gin.Context) {
	requests := models.KodeMegaProses{}

	if err := c.Bind(&requests); err != nil {
		mp.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}
	datas, err := mp.service.GetMajorByMegaProses(requests)
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

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas)
}
