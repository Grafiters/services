package controllers

import (
	"database/sql"
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/eventtypelv1"
	services "riskmanagement/services/eventtypelv1"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type EventTypeLv1Controller struct {
	logger  logger.Logger
	service services.EventTypeLv1Definition
}

func NewEventTypeLV1Controller(
	ET1Service services.EventTypeLv1Definition,
	logger logger.Logger,
) EventTypeLv1Controller {
	return EventTypeLv1Controller{
		logger:  logger,
		service: ET1Service,
	}
}

func (et1 EventTypeLv1Controller) GetAll(c *gin.Context) {
	datas, err := et1.service.GetAll()
	if err != nil {
		et1.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}

func (et1 EventTypeLv1Controller) GetAllWithPaginate(c *gin.Context) {
	requests := models.Paginate{}

	if err := c.Bind(&requests); err != nil {
		et1.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := et1.service.GetAllWithPaginate(requests)
	if err != nil {
		et1.logger.Zap.Error(err)
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

func (et1 EventTypeLv1Controller) GetOne(c *gin.Context) {
	requests := models.EventTypeLv1Request{}

	if err := c.Bind(&requests); err != nil {
		et1.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := et1.service.GetOne(requests.ID)
	if err != nil {
		et1.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (et1 EventTypeLv1Controller) Store(c *gin.Context) {
	data := models.EventTypeLv1Request{}

	if err := c.Bind(&data); err != nil {
		et1.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}
	fmt.Println(data)
	if err := et1.service.Store(&data); err != nil {
		et1.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Input Data Berhasil", true)
}

func (et1 EventTypeLv1Controller) Update(c *gin.Context) {
	data := models.EventTypeLv1Request{}

	if err := c.Bind(&data); err != nil {
		et1.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}
	fmt.Println(data)
	if err := et1.service.Update(&data); err != nil {
		et1.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Update Data Berhasil", true)
}

func (et1 EventTypeLv1Controller) Delete(c *gin.Context) {
	requests := models.EventTypeLv1Request{}

	if err := c.Bind(&requests); err != nil {
		et1.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	if err := et1.service.Delete(requests.ID); err != nil {
		et1.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Data berhasil dihapus", true)
}

func (et1 EventTypeLv1Controller) GetKodeEventType(c *gin.Context) {
	datas, err := et1.service.GetKodeEventType()
	if err != nil {
		et1.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if len(datas) == 0 {
		et1.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data Tidak Ditemukan !", "")
		return
	}

	KodeEventType := "ET1." + datas[0].KodeEventType

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", KodeEventType)
}
