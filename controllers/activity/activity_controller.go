package controllers

import (
	"database/sql"
	"riskmanagement/lib"
	models "riskmanagement/models/activity"
	services "riskmanagement/services/activity"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type ActivityController struct {
	logger  logger.Logger
	service services.ActivityDefinition
}

func NewActivityController(ActivityService services.ActivityDefinition, logger logger.Logger) ActivityController {
	return ActivityController{
		service: ActivityService,
		logger:  logger,
	}
}

func (activity ActivityController) GetAll(c *gin.Context) {
	datas, err := activity.service.GetAll()
	if err != nil {
		activity.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}

func (activity ActivityController) GetAllWithPagination(c *gin.Context) {
	requests := models.Paginate{}

	if err := c.Bind(&requests); err != nil {
		activity.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, pagination, err := activity.service.GetAllWithPagination(requests)
	if err != nil {
		activity.logger.Zap.Error()
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

func (activity ActivityController) GetOne(c *gin.Context) {
	// paramID := c.Param("id")
	// id, err := strconv.Atoi(paramID)

	requests := models.ActivityRequest{}

	if err := c.Bind(&requests); err != nil {
		activity.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := activity.service.GetOne(requests.ID)
	if err != nil {
		activity.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (activity ActivityController) Store(c *gin.Context) {
	data := models.ActivityRequest{}

	if err := c.Bind(&data); err != nil {
		activity.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}
	if err := activity.service.Store(&data); err != nil {
		activity.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Input Data Berhasil", true)
}

func (activity ActivityController) Update(c *gin.Context) {
	data := models.ActivityRequest{}

	if err := c.Bind(&data); err != nil {
		activity.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	if err := activity.service.Update(&data); err != nil {
		activity.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Update data berhasil", data)
}

func (activity ActivityController) Delete(c *gin.Context) {
	requests := models.ActivityRequest{}

	if err := c.Bind(&requests); err != nil {
		activity.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	if err := activity.service.Delete(requests.ID); err != nil {
		activity.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Data berhasil dihapus", true)
}

func (activity ActivityController) GetKodeActivity(c *gin.Context) {
	datas, err := activity.service.GetKodeActivity()
	if err != nil {
		activity.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if len(datas) == 0 {
		activity.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data Tidak Ditemukan !", "")
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas[0].KodeActivity)
}
