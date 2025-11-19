package controllers

import (
	"database/sql"
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/subincident"
	services "riskmanagement/services/subincident"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type SubIncidentController struct {
	logger  logger.Logger
	service services.SubIncidentDefinition
}

func NewSubIncidentController(
	SubIncidentService services.SubIncidentDefinition,
	logger logger.Logger,
) SubIncidentController {
	return SubIncidentController{
		service: SubIncidentService,
		logger:  logger,
	}
}

func (subIncident SubIncidentController) GetAll(c *gin.Context) {
	datas, err := subIncident.service.GetAll()
	if err != nil {
		subIncident.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}

func (pk2 SubIncidentController) GetAllWithPaginate(c *gin.Context) {
	requests := models.Paginate{}

	if err := c.Bind(&requests); err != nil {
		pk2.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := pk2.service.GetAllWithPaginate(requests)
	if err != nil {
		pk2.logger.Zap.Error(err)
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

func (subIncident SubIncidentController) GetOne(c *gin.Context) {
	requests := models.SubIncidentRequest{}

	if err := c.Bind(&requests); err != nil {
		subIncident.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := subIncident.service.GetOne(requests.ID)
	if err != nil {
		subIncident.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhsil", data)
}

func (subIncident SubIncidentController) GetSubIncidentByID(c *gin.Context) {
	requests := models.SubIncidentFilterRequest{}

	if err := c.Bind(&requests); err != nil {
		subIncident.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai :"+err.Error(), "")
		return
	}

	datas, err := subIncident.service.GetSubIncidentByID(requests)
	if err != nil {
		subIncident.logger.Zap.Error(err)
	}

	if len(datas) == 0 {
		subIncident.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data tidak ditemukan", nil)
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	fmt.Println("SubIncident =>", datas)
	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas)
}

func (subIncident SubIncidentController) Store(c *gin.Context) {
	data := models.SubIncidentRequest{}

	if err := c.Bind(&data); err != nil {
		subIncident.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak sesuai : "+err.Error(), "")
		return
	}

	fmt.Println(data)
	if err := subIncident.service.Store(&data); err != nil {
		subIncident.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "200", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Input data berhasil", data)
}

func (subIncident SubIncidentController) Update(c *gin.Context) {
	data := models.SubIncidentRequest{}

	if err := c.Bind(&data); err != nil {
		subIncident.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak sesuai : "+err.Error(), "")
		return
	}

	if err := subIncident.service.Update(&data); err != nil {
		subIncident.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", data)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Update Data Berhasil", data)
}

func (subIncident SubIncidentController) Delete(c *gin.Context) {
	requests := models.SubIncidentRequest{}

	if err := c.Bind(&requests); err != nil {
		subIncident.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	if err := subIncident.service.Delete(requests.ID); err != nil {
		subIncident.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Data berhasil dihapus", true)
}

func (subIncident SubIncidentController) GetKodePenyebabKejadian(c *gin.Context) {
	datas, err := subIncident.service.GetKodePenyebabKejadian()
	if err != nil {
		subIncident.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if len(datas) == 0 {
		subIncident.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data Tidak Ditemukan !", nil)
		return
	}

	KodePenyebabKejadian := "PK2.MOP." + datas[0].KodePenyebabKejadian

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", KodePenyebabKejadian)
}
