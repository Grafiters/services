package controllersmajorproses

import (
	"database/sql"
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/submajorproses"
	services "riskmanagement/services/submajorproses"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type SubMajorProsesController struct {
	logger  logger.Logger
	service services.SubMajorProsesDefinition
}

func NewSubMajorProsesController(
	SUBMPService services.SubMajorProsesDefinition,
	logger logger.Logger,
) SubMajorProsesController {
	return SubMajorProsesController{
		logger:  logger,
		service: SUBMPService,
	}
}

func (submp SubMajorProsesController) GetAll(c *gin.Context) {
	datas, err := submp.service.GetAll()
	if err != nil {
		submp.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}

func (subMajor SubMajorProsesController) GetAllWithPaginate(c *gin.Context) {
	requests := models.Paginate{}

	if err := c.Bind(&requests); err != nil {
		subMajor.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := subMajor.service.GetAllWithPaginate(requests)
	if err != nil {
		subMajor.logger.Zap.Error(err)
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

func (submp SubMajorProsesController) GetOne(c *gin.Context) {
	requests := models.SubMajorProsesRequest{}

	if err := c.Bind(&requests); err != nil {
		submp.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := submp.service.GetOne(requests.ID)
	if err != nil {
		submp.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (submp SubMajorProsesController) Store(c *gin.Context) {
	data := models.SubMajorProsesRequest{}

	if err := c.Bind(&data); err != nil {
		submp.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}
	fmt.Println(data)
	response, err := submp.service.Store(&data)
	if err != nil {
		submp.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	if !response {
		lib.ReturnToJson(c, 200, "400", "Kode '"+data.KodeSubMajorProses+"' sudah terdaftar, silahkan masukkan kode lain.", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Input Data Berhasil", true)
}

func (submp SubMajorProsesController) Update(c *gin.Context) {
	data := models.SubMajorProsesRequest{}

	if err := c.Bind(&data); err != nil {
		submp.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}
	fmt.Println(data)
	if err := submp.service.Update(&data); err != nil {
		submp.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Update Data Berhasil", true)
}

func (submp SubMajorProsesController) Delete(c *gin.Context) {
	requests := models.SubMajorProsesRequest{}

	if err := c.Bind(&requests); err != nil {
		submp.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	if err := submp.service.Delete(requests.ID); err != nil {
		submp.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Data berhasil dihapus", true)
}

func (submp SubMajorProsesController) GetDataByID(c *gin.Context) {
	requests := models.KodeMajor{}

	if err := c.Bind(&requests); err != nil {
		submp.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai :"+err.Error(), "")
	}

	datas, err := submp.service.GetDataByID(requests)
	if err != nil {
		submp.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if len(datas) == 0 {
		submp.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data Tidak Ditemukan !", nil)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas)
}
