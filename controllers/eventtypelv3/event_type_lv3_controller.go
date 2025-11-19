package controllers

import (
	"database/sql"
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/eventtypelv3"
	services "riskmanagement/services/eventtypelv3"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type EventTypeLv3Controller struct {
	logger  logger.Logger
	service services.EventTypeLv3Definition
}

func NewEventTypeLV3Controller(
	ET3Service services.EventTypeLv3Definition,
	logger logger.Logger,
) EventTypeLv3Controller {
	return EventTypeLv3Controller{
		logger:  logger,
		service: ET3Service,
	}
}

func (et3 EventTypeLv3Controller) GetAll(c *gin.Context) {
	datas, err := et3.service.GetAll()
	if err != nil {
		et3.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}

func (et3 EventTypeLv3Controller) GetAllWithPaginate(c *gin.Context) {
	requests := models.Paginate{}

	if err := c.Bind(&requests); err != nil {
		et3.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := et3.service.GetAllWithPaginate(requests)
	if err != nil {
		et3.logger.Zap.Error(err)
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

func (et3 EventTypeLv3Controller) GetOne(c *gin.Context) {
	requests := models.EventTypeLv3Request{}

	if err := c.Bind(&requests); err != nil {
		et3.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := et3.service.GetOne(requests.ID)
	if err != nil {
		et3.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (et3 EventTypeLv3Controller) Store(c *gin.Context) {
	data := models.EventTypeLv3Request{}

	if err := c.Bind(&data); err != nil {
		et3.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}
	fmt.Println(data)
	if err := et3.service.Store(&data); err != nil {
		et3.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Input Data Berhasil", true)
}

func (et3 EventTypeLv3Controller) Update(c *gin.Context) {
	data := models.EventTypeLv3Request{}

	if err := c.Bind(&data); err != nil {
		et3.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}
	fmt.Println(data)
	if err := et3.service.Update(&data); err != nil {
		et3.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Update Data Berhasil", true)
}

func (et3 EventTypeLv3Controller) Delete(c *gin.Context) {
	requests := models.EventTypeLv3Request{}

	if err := c.Bind(&requests); err != nil {
		et3.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	if err := et3.service.Delete(requests.ID); err != nil {
		et3.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Data berhasil dihapus", true)
}

func (et3 EventTypeLv3Controller) GetKodeEventType(c *gin.Context) {
	datas, err := et3.service.GetKodeEventType()
	if err != nil {
		et3.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if len(datas) == 0 {
		et3.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data Tidak Ditemukan !", nil)
		return
	}

	KodeEventType := "ET3.MOP." + datas[0].KodeEventType

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", KodeEventType)
}

func (et3 EventTypeLv3Controller) GetEventTypeById2(c *gin.Context) {
	requests := models.IDEventTypeLv2{}

	if err := c.Bind(&requests); err != nil {
		et3.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "input tidak sesuai :"+err.Error(), "")
		return
	}

	datas, err := et3.service.GetEventTypeById2(requests)
	if err != nil {
		et3.logger.Zap.Error(err)
	}

	if len(datas) == 0 {
		et3.logger.Zap.Error(err)
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
