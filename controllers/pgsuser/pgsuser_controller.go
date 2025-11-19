package controllers

import (
	"database/sql"
	"riskmanagement/lib"
	models "riskmanagement/models/pgsuser"

	services "riskmanagement/services/pgsuser"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type PgsUserController struct {
	logger  logger.Logger
	service services.PgsUserDefinition
}

func NewPgsUserController(
	logger logger.Logger,
	pgsUserService services.PgsUserDefinition,
) PgsUserController {
	return PgsUserController{
		logger:  logger,
		service: pgsUserService,
	}
}

func (pgsUser PgsUserController) Delete(c *gin.Context) {
	data := models.UpdateDelete{}

	if err := c.Bind(&data); err != nil {
		pgsUser.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := pgsUser.service.Delete(&data)
	if err != nil {
		pgsUser.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if !status {
		pgsUser.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data Gagal Dihapus", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Hapus data berhasil", true)
}

func (pgsUser PgsUserController) GetAll(c *gin.Context) {
	// paramID := c.Param("id")
	// makerID, err := strconv.Atoi(paramID)

	requests := models.PgsUser{}
	if err := c.Bind(&requests); err != nil {
		pgsUser.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai :"+err.Error(), "")
		return
	}

	datas, err := pgsUser.service.GetAll(requests.MakerID)
	if err != nil {
		pgsUser.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "200", "Internal Error", "")
		return
	}

	if len(datas) == 0 {
		pgsUser.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data tidak ditemukan", nil)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas)
}

func (pgsUser PgsUserController) GetAllWithPaginate(c *gin.Context) {
	requests := models.Paginate{}

	if err := c.Bind(&requests); err != nil {
		pgsUser.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := pgsUser.service.GetAllWithPaginate(requests)
	if err != nil {
		pgsUser.logger.Zap.Error(err)
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

func (pgsUser PgsUserController) Store(c *gin.Context) {
	requests := models.PgsUserRequest{}

	if err := c.Bind(&requests); err != nil {
		pgsUser.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai :"+err.Error(), "")
		return
	}

	status, err := pgsUser.service.Store(requests)
	if err != nil {
		pgsUser.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error status", false)
		return
	}

	if !status {
		// pgsUser.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Sedang dalam periode PGS aktif, atau sedang dalam proses request", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Input data berhasil", true)
}

func (pgsUser PgsUserController) GetPgsApproval(c *gin.Context) {
	requests := models.Paginate{}

	if err := c.Bind(&requests); err != nil {
		pgsUser.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai :"+err.Error(), "")
		return
	}

	datas, pagination, err := pgsUser.service.GetPgsApproval(requests)
	if err != nil {
		pgsUser.logger.Zap.Error(err)
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

func (pgsUser PgsUserController) GetOne(c *gin.Context) {
	requests := models.PgsUserRequest{}

	if err := c.Bind(&requests); err != nil {
		pgsUser.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai :"+err.Error(), "")
		return
	}

	data, status, err := pgsUser.service.GetOne(requests.ID)
	if err != nil {
		pgsUser.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	if !status {
		pgsUser.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data tidak ditemukan", nil)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}

func (pgsUser PgsUserController) Update(c *gin.Context) {
	data := models.PgsUserRequestUpdate{}

	if err := c.Bind(&data); err != nil {
		pgsUser.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := pgsUser.service.Update(&data)
	if err != nil {
		pgsUser.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if !status {
		pgsUser.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error status", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Update data berhasil", true)
}

func (pgsUser PgsUserController) ApprovePgsUser(c *gin.Context) {
	data := models.PgsUpdateApproval{}

	if err := c.Bind(&data); err != nil {
		pgsUser.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := pgsUser.service.ApprovePgsUser(&data)
	if err != nil {
		pgsUser.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if !status {
		pgsUser.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error status", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Update data berhasil", true)
}

func (pgsUser PgsUserController) RejectPgsUser(c *gin.Context) {
	data := models.PgsUpdateApproval{}

	if err := c.Bind(&data); err != nil {
		pgsUser.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := pgsUser.service.RejectPgsUser(&data)
	if err != nil {
		pgsUser.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if !status {
		pgsUser.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error status", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Update data berhasil", true)
}

func (pgsUser PgsUserController) SearchPekerjaByPn(c *gin.Context) {

	request := models.RequestPn{}

	if err := c.Bind(&request); err != nil {
		pgsUser.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai :"+err.Error(), "")
		return
	}

	data, err := pgsUser.service.SearchPekerjaByPn(request.PERNR)
	if err != nil {
		pgsUser.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	// if len(data) == 0 {
	// 	pgsUser.logger.Zap.Error(err)
	// 	lib.ReturnToJson(c, 200, "404", "Data tidak ditemukan", nil)
	// 	return
	// }

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}
