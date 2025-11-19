package controllers

import (
	"database/sql"
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/ukerkelolaan"
	services "riskmanagement/services/ukerkelolaan"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type UkerKolalaanController struct {
	logger    logger.Logger
	ukService services.UkerKelolaanDefinition
}

func NewUkerKelolaanController(
	ukService services.UkerKelolaanDefinition,
	logger logger.Logger,
) UkerKolalaanController {
	return UkerKolalaanController{
		logger:    logger,
		ukService: ukService,
	}
}

func (uk UkerKolalaanController) GetAllWithPaginate(c *gin.Context) {
	requests := models.KeywordRequest{}

	if err := c.Bind(&requests); err != nil {
		uk.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := uk.ukService.GetAllWithPaginate(requests)
	if err != nil {
		uk.logger.Zap.Error(err)
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

func (uk UkerKolalaanController) FilterUkerKelolaan(c *gin.Context) {
	requests := models.KeywordRequest{}

	if err := c.Bind(&requests); err != nil {
		uk.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := uk.ukService.FilterUkerKelolaan(requests)
	if err != nil {
		uk.logger.Zap.Error(err)
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

func (uk UkerKolalaanController) Store(c *gin.Context) {
	requests := models.UkerKelolaanRequest{}

	if err := c.Bind(&requests); err != nil {
		uk.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai :"+err.Error(), "")
		return
	}

	status, err := uk.ukService.Store(requests)
	if err != nil {
		uk.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error status", false)
		return
	}

	if !status {
		// UkerKelolaan.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "PN '"+requests.Pn+"' dalam periode uker kelolaan aktif", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Input data berhasil", true)
}

func (uk UkerKolalaanController) GetOne(c *gin.Context) {
	requests := models.RequestOne{}

	if err := c.Bind(&requests); err != nil {
		uk.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai :"+err.Error(), "")
		return
	}

	data, status, err := uk.ukService.GetOne(requests.Id)
	if err != nil {
		uk.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	if !status {
		// uk.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data tidak ditemukan", nil)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}

func (uk UkerKolalaanController) Update(c *gin.Context) {
	requests := models.UkerKelolaanRequest{}

	if err := c.Bind(&requests); err != nil {
		uk.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai :"+err.Error(), "")
		return
	}

	status, err := uk.ukService.Update(&requests)
	if err != nil {
		uk.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error status", false)
		return
	}

	if !status {
		uk.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data Gagal diupdate", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Update data berhasil", true)
}

func (uk UkerKolalaanController) Delete(c *gin.Context) {
	request := models.UkerKelolaanRequest{}

	if err := c.Bind(&request); err != nil {
		uk.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := uk.ukService.Delete(request)
	if err != nil {
		uk.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if !status {
		uk.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error status", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Delete data berhasil", true)
}

func (uk UkerKolalaanController) GetListUkerKelolaan(c *gin.Context) {
	requests := models.PencarianUker{}

	fmt.Println(requests)

	if err := c.Bind(&requests); err != nil {
		uk.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	data, err := uk.ukService.GetListUkerKelolaan(&requests)
	if err != nil {
		uk.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}
