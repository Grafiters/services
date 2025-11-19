package controllers

import (
	"database/sql"
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/linibisnislv2"
	services "riskmanagement/services/linibisnislv2"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type LiniBisnisLv2Controller struct {
	logger  logger.Logger
	service services.LiniBisnisLv2Definition
}

func NewLiniBisnisLV2Controller(
	LB2Service services.LiniBisnisLv2Definition,
	logger logger.Logger,
) LiniBisnisLv2Controller {
	return LiniBisnisLv2Controller{
		logger:  logger,
		service: LB2Service,
	}
}

func (lb2 LiniBisnisLv2Controller) GetAll(c *gin.Context) {
	datas, err := lb2.service.GetAll()
	if err != nil {
		lb2.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}

func (lb2 LiniBisnisLv2Controller) GetAllWithPaginate(c *gin.Context) {
	requests := models.Paginate{}

	if err := c.Bind(&requests); err != nil {
		lb2.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := lb2.service.GetAllWithPaginate(requests)
	if err != nil {
		lb2.logger.Zap.Error(err)
	}

	if pagination.Total == 0 {
		lib.ReturnToJson(c, 200, "404", "Data Kosong", nil)
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery data berhasil", datas, pagination)
}

func (lb2 LiniBisnisLv2Controller) GetOne(c *gin.Context) {
	requests := models.LiniBisnisLv2{}

	if err := c.Bind(&requests); err != nil {
		lb2.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := lb2.service.GetOne(requests.ID)
	if err != nil {
		lb2.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (lb2 LiniBisnisLv2Controller) Store(c *gin.Context) {
	data := models.LiniBisnisLv2Request{}

	if err := c.Bind(&data); err != nil {
		lb2.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}
	fmt.Println(data)
	if err := lb2.service.Store(&data); err != nil {
		lb2.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Input Data Berhasil", true)
}

func (lb2 LiniBisnisLv2Controller) Update(c *gin.Context) {
	data := models.LiniBisnisLv2Request{}

	if err := c.Bind(&data); err != nil {
		lb2.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}
	fmt.Println(data)
	if err := lb2.service.Update(&data); err != nil {
		lb2.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Update Data Berhasil", true)
}

func (lb2 LiniBisnisLv2Controller) Delete(c *gin.Context) {
	requests := models.LiniBisnisLv2{}

	if err := c.Bind(&requests); err != nil {
		lb2.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	if err := lb2.service.Delete(requests.ID); err != nil {
		lb2.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Data berhasil dihapus", true)
}

func (lb2 LiniBisnisLv2Controller) GetKodeLiniBisnis(c *gin.Context) {
	datas, err := lb2.service.GetKodeLiniBisnis()
	if err != nil {
		lb2.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if len(datas) == 0 {
		lb2.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data Tidak Ditemukan !", nil)
		return
	}

	KodeLiniBisnis := "LB2." + datas[0].KodeLiniBisnis

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", KodeLiniBisnis)
}

func (lb2 LiniBisnisLv2Controller) GetLBByID(c *gin.Context) {
	requests := models.KodeLB1{}

	if err := c.Bind(&requests); err != nil {
		lb2.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "input tidak sesuai :"+err.Error(), "")
		return
	}

	datas, err := lb2.service.GetLBByID(requests)
	if err != nil {
		lb2.logger.Zap.Error(err)
	}

	if len(datas) == 0 {
		lb2.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data tidak ditemukan", nil)
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	// fmt.Println("SubIncident =>", datas)
	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas)
}
