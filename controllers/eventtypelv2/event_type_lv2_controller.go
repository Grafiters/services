package controllers

import (
	"database/sql"
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/eventtypelv2"
	services "riskmanagement/services/eventtypelv2"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type EventTypeLv2Controller struct {
	logger  logger.Logger
	service services.EventTypeLv2Definition
}

func NewEventTypeLV2Controller(
	ET2Service services.EventTypeLv2Definition,
	logger logger.Logger,
) EventTypeLv2Controller {
	return EventTypeLv2Controller{
		logger:  logger,
		service: ET2Service,
	}
}

func (et2 EventTypeLv2Controller) GetAll(c *gin.Context) {
	datas, err := et2.service.GetAll()
	if err != nil {
		et2.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}

func (et2 EventTypeLv2Controller) GetOne(c *gin.Context) {
	requests := models.EventTypeLv2Request{}

	if err := c.Bind(&requests); err != nil {
		et2.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := et2.service.GetOne(requests.ID)
	if err != nil {
		et2.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (et2 EventTypeLv2Controller) GetAllWithPaginate(c *gin.Context) {
	requests := models.Paginate{}

	if err := c.Bind(&requests); err != nil {
		et2.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := et2.service.GetAllWithPaginate(requests)
	if err != nil {
		et2.logger.Zap.Error(err)
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

func (et2 EventTypeLv2Controller) Store(c *gin.Context) {
	data := models.EventTypeLv2Request{}

	if err := c.Bind(&data); err != nil {
		et2.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}
	fmt.Println(data)
	if err := et2.service.Store(&data); err != nil {
		et2.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Input Data Berhasil", true)
}

func (et2 EventTypeLv2Controller) Update(c *gin.Context) {
	data := models.EventTypeLv2Request{}

	if err := c.Bind(&data); err != nil {
		et2.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}
	fmt.Println(data)
	if err := et2.service.Update(&data); err != nil {
		et2.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Update Data Berhasil", true)
}

func (et2 EventTypeLv2Controller) Delete(c *gin.Context) {
	requests := models.EventTypeLv2Request{}

	if err := c.Bind(&requests); err != nil {
		et2.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	if err := et2.service.Delete(requests.ID); err != nil {
		et2.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Data berhasil dihapus", true)
}

func (et2 EventTypeLv2Controller) GetKodeEventType(c *gin.Context) {
	datas, err := et2.service.GetKodeEventType()
	if err != nil {
		et2.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if len(datas) == 0 {
		et2.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data Tidak Ditemukan !", nil)
		return
	}

	KodeEventType := "ET2.MOP." + datas[0].KodeEventType

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", KodeEventType)
}

func (et2 EventTypeLv2Controller) GetEventTypeById1(c *gin.Context) {
	requests := models.IDEventTypeLv1{}

	if err := c.Bind(&requests); err != nil {
		et2.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "input tidak sesuai :"+err.Error(), "")
		return
	}

	datas, err := et2.service.GetEventTypeById1(requests)
	if err != nil {
		et2.logger.Zap.Error(err)
	}

	if len(datas) == 0 {
		et2.logger.Zap.Error(err)
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
