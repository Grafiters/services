package admin_setting

import (
	"database/sql"
	"riskmanagement/lib"
	models "riskmanagement/models/admin_setting"
	services "riskmanagement/services/admin_setting"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type AdminSettingController struct {
	logger  logger.Logger
	service services.AdminSettingDefinition
}

func NewAdminSettingController(
	AdminSettingService services.AdminSettingDefinition,
	logger logger.Logger,
) AdminSettingController {
	return AdminSettingController{
		service: AdminSettingService,
		logger:  logger,
	}
}

func (setting AdminSettingController) GetAll(c *gin.Context) {
	requests := models.Paginate{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		setting.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, pagination, err := setting.service.GetAll(requests)
	if err != nil {
		setting.logger.Zap.Error()
	}

	if pagination.Total == 0 {
		lib.ReturnToJson(c, 200, "404", "Data Kosong", data)
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", data, pagination)
}

func (setting AdminSettingController) Show(c *gin.Context) {
	requests := models.Paginate{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		setting.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, pagination, err := setting.service.Show(requests)
	if err != nil {
		setting.logger.Zap.Error()
	}

	if pagination.Total == 0 {
		lib.ReturnToJson(c, 200, "404", "Data Kosong", data)
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", data, pagination)
}

func (setting AdminSettingController) Store(c *gin.Context) {
	data := models.AdminSettingRequest{}

	if err := c.Bind(&data); err != nil {
		setting.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	status, err := setting.service.Store(data)
	if err != nil {
		setting.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if !status {
		setting.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error status", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Input data berhasil", true)
}

func (setting AdminSettingController) Update(c *gin.Context) {
	data := models.AdminSettingUpdateRequest{}

	if err := c.Bind(&data); err != nil {
		setting.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := setting.service.Update(data)
	if err != nil {
		setting.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error : ", "")
		return
	}

	if !status {
		setting.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data Gagal Diupdate : ", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Update data berhasil", true)
}

func (setting AdminSettingController) Delete(c *gin.Context) {
	data := models.AdminSettingDelete{}

	if err := c.Bind(&data); err != nil {
		setting.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := setting.service.Delete(&data)
	if err != nil {
		setting.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error : ", "")
		return
	}

	if !status {
		setting.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data Gagal Didelete : ", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Hapus data berhasil", true)
}

func (setting AdminSettingController) SearchTaskType(c *gin.Context) {
	requests := models.KeywordRequest{}

	if err := c.Bind(&requests); err != nil {
		setting.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := setting.service.SearchTaskType(requests)
	if err != nil {
		setting.logger.Zap.Error(err)
	}

	if pagination.Total == 0 {
		lib.ReturnToJson(c, 200, "404", "Data Tidak Ditemukan", datas)
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", datas, pagination)
}

func (setting AdminSettingController) SearchTaskTypeInput(c *gin.Context) {
	requests := models.KeywordRequest{}

	if err := c.Bind(&requests); err != nil {
		setting.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := setting.service.SearchTaskTypeInput(requests)
	if err != nil {
		setting.logger.Zap.Error(err)
	}

	if pagination.Total == 0 {
		lib.ReturnToJson(c, 200, "404", "Data Tidak Ditemukan", datas)
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", datas, pagination)
}

func (setting AdminSettingController) SearchTaskTypeInputByKegiatan(c *gin.Context) {
	requests := models.KeywordRequest{}

	if err := c.Bind(&requests); err != nil {
		setting.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := setting.service.SearchTaskTypeInputByKegiatan(requests)
	if err != nil {
		setting.logger.Zap.Error(err)
	}

	if pagination.Total == 0 {
		lib.ReturnToJson(c, 200, "404", "Data Tidak Ditemukan", datas)
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", datas, pagination)
}

func (setting AdminSettingController) GetOne(c *gin.Context) {
	requests := models.TaskTypeRequestOne{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		setting.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}
	data, pagination, err := setting.service.GetOne(requests)
	if err != nil {
		setting.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", data, pagination)
}
