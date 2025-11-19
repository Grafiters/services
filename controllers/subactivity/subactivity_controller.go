package controllers

import (
	"database/sql"
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/subactivity"
	services "riskmanagement/services/subactivity"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type SubActivityController struct {
	logger  logger.Logger
	service services.SubActivityDefinition
}

func NewSubActivityController(SubActivityService services.SubActivityDefinition, logger logger.Logger) SubActivityController {
	return SubActivityController{
		service: SubActivityService,
		logger:  logger,
	}
}

func (subactivity SubActivityController) GetAll(c *gin.Context) {
	datas, err := subactivity.service.GetAll()
	// fmt.Println(datas)
	if err != nil {
		subactivity.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}

func (subactivity SubActivityController) GetLastID(c *gin.Context) {
	// paramID := c.Param("id")
	// id, err := strconv.Atoi(paramID)
	requests := models.SubActivityRequest{}

	data, err := subactivity.service.GetLastID(requests.ID)

	// fmt.Println(data)
	if err != nil {
		subactivity.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", len(data))
}

func (subactivity SubActivityController) GetSubactivity(c *gin.Context) {
	requests := models.SubActivityRequest{}

	if err := c.Bind(&requests); err != nil {
		subactivity.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := subactivity.service.GetLastID(requests.ActivityID)

	// fmt.Println(data)
	if err != nil {
		subactivity.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (subactivity SubActivityController) GetAllWithPagination(c *gin.Context) {
	requests := models.Paginate{}

	if err := c.Bind(&requests); err != nil {
		subactivity.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, pagination, err := subactivity.service.GetAllWithPagination(requests)
	if err != nil {
		subactivity.logger.Zap.Error()
	}

	if pagination.Total == 0 {
		lib.ReturnToJson(c, 200, "404", "Data Kosong", nil)
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", data, pagination)
}

func (subactivity SubActivityController) GetOne(c *gin.Context) {
	requests := models.SubActivityRequest{}

	if err := c.Bind(&requests); err != nil {
		subactivity.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := subactivity.service.GetOne(requests.ID)
	if err != nil {
		subactivity.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhsil", data)
}

func (subactivity SubActivityController) Store(c *gin.Context) {
	data := models.SubActivityRequest{}

	if err := c.Bind(&data); err != nil {
		subactivity.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak sesuai : "+err.Error(), "")
		return
	}

	fmt.Println(data)
	if err := subactivity.service.Store(&data); err != nil {
		subactivity.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Input data berhasil", data)
}

func (subactivity SubActivityController) Update(c *gin.Context) {
	data := models.SubActivityRequest{}

	if err := c.Bind(&data); err != nil {
		subactivity.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak sesuai : "+err.Error(), "")
		return
	}

	if err := subactivity.service.Update(&data); err != nil {
		subactivity.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", data)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Update Data Berhasil", data)
}

func (subactivity SubActivityController) Delete(c *gin.Context) {
	requests := models.SubActivityRequest{}

	if err := c.Bind(&requests); err != nil {
		subactivity.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	if err := subactivity.service.Delete(requests.ID); err != nil {
		subactivity.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Data berhasil dihapus", true)
}

func (subactivity SubActivityController) GetKodeSubActivity(c *gin.Context) {
	requests := models.KodeSubActivityRequest{}

	if err := c.Bind(&requests); err != nil {
		subactivity.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	data, err := subactivity.service.GetKodeSubActivity(requests)

	if err != nil {
		subactivity.logger.Zap.Error(err)
	}

	if len(data) == 0 {
		subactivity.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data Tidak ditemukan", nil)
		return
	}

	counter := data[0].KodeSubActivity
	fmt.Println("kode_activity", data[0].KodeActivity)
	fmt.Println("kode_sub_activity", data[0].KodeSubActivity)

	KodeSub := data[0].KodeActivity + "." + counter
	// fmt.Print(KodeSub)

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", KodeSub)

}
