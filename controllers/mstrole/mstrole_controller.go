package controllers

import (
	"database/sql"
	"riskmanagement/lib"
	models "riskmanagement/models/mstrole"
	services "riskmanagement/services/mstrole"

	"github.com/gin-gonic/gin"

	"gitlab.com/golang-package-library/logger"
)

type MstRoleController struct {
	logger  logger.Logger
	service services.MstRoleDefinition
}

func NewMstRoleController(
	MstRoleService services.MstRoleDefinition,
	logger logger.Logger,
) MstRoleController {
	return MstRoleController{
		service: MstRoleService,
		logger:  logger,
	}
}

func (mstRole MstRoleController) GetAll(c *gin.Context) {
	datas, err := mstRole.service.GetAll()

	if err != nil {
		mstRole.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "200", "Internal Error", "")
		return
	}

	if len(datas) == 0 {
		mstRole.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data tidak ditemukan", nil)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas)
}

func (msRole MstRoleController) GetAllWithPaginate(c *gin.Context) {
	requests := models.Paginate{}

	if err := c.Bind(&requests); err != nil {
		msRole.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := msRole.service.GetAllWithPaginate(requests)
	if err != nil {
		msRole.logger.Zap.Error(err)
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

func (mstRole MstRoleController) GetOne(c *gin.Context) {
	request := models.MstRoleRequest{}

	if err := c.Bind(&request); err != nil {
		mstRole.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai :"+err.Error(), "")
		return
	}

	data, status, err := mstRole.service.GetOne(request.ID)
	if err != nil {
		mstRole.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	if !status {
		mstRole.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data tidak ditemukan", nil)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}

func (mstRole MstRoleController) Store(c *gin.Context) {
	data := models.MstRoleRequest{}

	if err := c.Bind(&data); err != nil {
		mstRole.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai :"+err.Error(), "")
		return
	}

	status, err := mstRole.service.Store(data)
	if err != nil {
		mstRole.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if !status {
		mstRole.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error status", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Input data berhasil", true)
}

func (mstRole MstRoleController) Update(c *gin.Context) {
	data := models.MstRoleRequest{}

	if err := c.Bind(&data); err != nil {
		mstRole.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai :"+err.Error(), "")
		return
	}

	status, err := mstRole.service.Update(&data)
	if err != nil {
		mstRole.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if !status {
		mstRole.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error status", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Update data berhasil", true)
}

func (mstRole MstRoleController) Delete(c *gin.Context) {
	data := models.MstRoleRequestDelete{}

	if err := c.Bind(&data); err != nil {
		mstRole.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai :"+err.Error(), "")
		return
	}

	status, err := mstRole.service.Delete(&data)
	if err != nil {
		mstRole.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if !status {
		mstRole.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error status", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Delete data berhasil", true)
}
