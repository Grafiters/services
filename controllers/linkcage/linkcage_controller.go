package Linkcage

import (
	"database/sql"
	"riskmanagement/lib"
	models "riskmanagement/models/linkcage"
	services "riskmanagement/services/linkcage"

	"github.com/gin-gonic/gin"
)

type LinkcageController struct {
	service services.LinkcageDefinition
}

func NewLinkcageController(LinkcageService services.LinkcageDefinition) LinkcageController {
	return LinkcageController{
		service: LinkcageService,
	}
}

func (link LinkcageController) GetAll(c *gin.Context) {
	requests := models.LinkcageRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, pagination, err := link.service.GetAll(requests)

	if err != nil {
		lib.ReturnToJson(c, 200, "500", "Internal Error", err)
		return
	}

	if pagination.Total == 0 {
		lib.ReturnToJson(c, 200, "404", "Data Kosong", data)
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", err)
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", data, pagination)
}

func (link LinkcageController) Store(c *gin.Context) {
	data := models.LinkcageRequest{}

	if err := c.Bind(&data); err != nil {
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), false)
		return
	}

	status, err := link.service.Store(data)
	if err != nil {
		lib.ReturnToJson(c, 200, "500", err.Error(), false)
		return
	}

	if !status {
		lib.ReturnToJson(c, 200, "500", "Internal Error Status", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Input Data Berhasil", true)
}

func (link LinkcageController) SetStatus(c *gin.Context) {
	data := models.LinkcageRequest{}

	if err := c.Bind(&data); err != nil {
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	status, err := link.service.SetStatus(&data)
	if err != nil {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if !status {
		lib.ReturnToJson(c, 200, "500", "Ubah Status Gagal", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Ubah Status Berhasil", true)
}

func (link LinkcageController) GetActive(c *gin.Context) {
	data, err := link.service.GetActive()

	if err != nil {
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (link LinkcageController) Delete(c *gin.Context) {
	data := models.LinkcageRequest{}

	if err := c.Bind(&data); err != nil {
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	status, err := link.service.Delete(&data)
	if err != nil {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if !status {
		lib.ReturnToJson(c, 200, "500", "Data Gagal Dihapus", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Delete Data Berhasil", true)
}

func (link LinkcageController) GetOne(c *gin.Context) {
	requests := models.LinkcageRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, status, err := link.service.GetOne(requests)
	if err != nil {

	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	if !status {
		lib.ReturnToJson(c, 200, "500", "Internal Error status", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (link LinkcageController) Update(c *gin.Context) {
	data := models.LinkcageRequest{}

	if err := c.Bind(&data); err != nil {
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	status, err := link.service.Update(&data)
	if err != nil {
		lib.ReturnToJson(c, 200, "500", err.Error(), false)
		return
	}

	if !status {
		lib.ReturnToJson(c, 200, "500", "Internal Error status", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Update data berhasil", true)
}
