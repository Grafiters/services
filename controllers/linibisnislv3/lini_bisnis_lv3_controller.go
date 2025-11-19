package controllers

import (
	"database/sql"
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/linibisnislv3"
	services "riskmanagement/services/linibisnislv3"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type LiniBisnisLv3Controller struct {
	logger  logger.Logger
	service services.LiniBisnisLv3Definition
}

func NewLiniBisnisLV3Controller(
	LB3Service services.LiniBisnisLv3Definition,
	logger logger.Logger,
) LiniBisnisLv3Controller {
	return LiniBisnisLv3Controller{
		logger:  logger,
		service: LB3Service,
	}
}

func (lb3 LiniBisnisLv3Controller) GetAll(c *gin.Context) {
	datas, err := lb3.service.GetAll()
	if err != nil {
		lb3.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}

func (lb3 LiniBisnisLv3Controller) GetAllWithPaginate(c *gin.Context) {
	requests := models.Paginate{}

	if err := c.Bind(&requests); err != nil {
		lb3.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := lb3.service.GetAllWithPaginate(requests)
	if err != nil {
		lb3.logger.Zap.Error(err)
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

func (lb3 LiniBisnisLv3Controller) GetOne(c *gin.Context) {
	requests := models.LiniBisnisLv3{}

	if err := c.Bind(&requests); err != nil {
		lb3.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := lb3.service.GetOne(requests.ID)
	if err != nil {
		lb3.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (lb3 LiniBisnisLv3Controller) Store(c *gin.Context) {
	data := models.LiniBisnisLv3Request{}

	if err := c.Bind(&data); err != nil {
		lb3.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}
	fmt.Println(data)
	if err := lb3.service.Store(&data); err != nil {
		lb3.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Input Data Berhasil", true)
}

func (lb3 LiniBisnisLv3Controller) Update(c *gin.Context) {
	data := models.LiniBisnisLv3Request{}

	if err := c.Bind(&data); err != nil {
		lb3.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}
	fmt.Println(data)
	if err := lb3.service.Update(&data); err != nil {
		lb3.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Update Data Berhasil", true)
}

func (lb3 LiniBisnisLv3Controller) Delete(c *gin.Context) {
	requests := models.LiniBisnisLv3{}

	if err := c.Bind(&requests); err != nil {
		lb3.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	if err := lb3.service.Delete(requests.ID); err != nil {
		lb3.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Data berhasil dihapus", true)
}

func (lb3 LiniBisnisLv3Controller) GetKodeLiniBisnis(c *gin.Context) {
	datas, err := lb3.service.GetKodeLiniBisnis()
	if err != nil {
		lb3.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if len(datas) == 0 {
		lb3.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data Tidak Ditemukan !", nil)
		return
	}

	KodeLiniBisnis := "LB3." + datas[0].KodeLiniBisnis

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", KodeLiniBisnis)
}

func (lb3 LiniBisnisLv3Controller) GetLBByID(c *gin.Context) {
	requests := models.KodeLB2{}

	if err := c.Bind(&requests); err != nil {
		lb3.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "input tidak sesuai :"+err.Error(), "")
		return
	}

	datas, err := lb3.service.GetLBByID(requests)
	if err != nil {
		lb3.logger.Zap.Error(err)
	}

	if len(datas) == 0 {
		lb3.logger.Zap.Error(err)
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
