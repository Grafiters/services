package controllers

import (
	"database/sql"
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/incident"
	services "riskmanagement/services/incident"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type IncidentController struct {
	logger  logger.Logger
	service services.IncidentDefinition
}

func NewIncidentController(
	IncidentService services.IncidentDefinition,
	logger logger.Logger,
) IncidentController {
	return IncidentController{
		service: IncidentService,
		logger:  logger,
	}
}

func (incident IncidentController) GetAll(c *gin.Context) {
	datas, err := incident.service.GetAll()
	if err != nil {
		incident.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas)
}

func (incident IncidentController) GetOne(c *gin.Context) {
	requests := models.IncidentRequest{}

	if err := c.Bind(&requests); err != nil {
		incident.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := incident.service.GetOne(requests.ID)
	if err != nil {
		incident.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}

func (pk1 IncidentController) GetAllWithPaginate(c *gin.Context) {
	requests := models.Paginate{}

	if err := c.Bind(&requests); err != nil {
		pk1.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := pk1.service.GetAllWithPaginate(requests)
	if err != nil {
		pk1.logger.Zap.Error(err)
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

func (incident IncidentController) Store(c *gin.Context) {
	data := models.IncidentRequest{}

	if err := c.Bind(&data); err != nil {
		incident.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	fmt.Println(data)
	if err := incident.service.Store(&data); err != nil {
		incident.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Input data berhasil", data)
}

func (incident IncidentController) Update(c *gin.Context) {
	data := models.IncidentRequest{}

	if err := c.Bind(&data); err != nil {
		incident.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	if err := incident.service.Update(&data); err != nil {
		incident.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Update data berhasil", data)
}

func (incident IncidentController) Delete(c *gin.Context) {
	requests := models.IncidentRequest{}

	if err := c.Bind(&requests); err != nil {
		incident.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	if err := incident.service.Delete(requests.ID); err != nil {
		incident.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Data berhasil dihapus", true)
}

func (incident IncidentController) GetKodePenyebabKejadian(c *gin.Context) {
	datas, err := incident.service.GetKodePenyebabKejadian()
	if err != nil {
		incident.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if len(datas) == 0 {
		incident.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data Tidak Ditemukan !", nil)
		return
	}

	KodePenyebabKejadian := "PK1.MOP." + datas[0].KodePenyebabKejadian

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", KodePenyebabKejadian)
}
