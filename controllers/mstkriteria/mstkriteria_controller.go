package controller

import (
	"database/sql"
	"riskmanagement/lib"
	models "riskmanagement/models/mstkriteria"
	services "riskmanagement/services/mstkriteria"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type MstKriteriaController struct {
	logger  logger.Logger
	service services.MstKriteriaDefinition
}

func NewMstKriteriaController(
	mstKriteriaService services.MstKriteriaDefinition,
	logger logger.Logger,
) MstKriteriaController {
	return MstKriteriaController{
		logger:  logger,
		service: mstKriteriaService,
	}
}

func (mstKriteria MstKriteriaController) GetAll(c *gin.Context) {
	requests := models.FilterRequest{}
	if err := c.Bind(&requests); err != nil {
		mstKriteria.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, err := mstKriteria.service.GetAll(requests)

	if err != nil {
		mstKriteria.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "200", "Internal Error", "")
		return
	}

	if len(datas) == 0 {
		mstKriteria.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data tidak ditemukan", nil)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas)
}

func (mstKriteria MstKriteriaController) GetAllWithPaginate(c *gin.Context) {
	requests := models.FilterRequest{}
	if err := c.Bind(&requests); err != nil {
		mstKriteria.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := mstKriteria.service.GetAllWithPaginate(requests)
	if err != nil {
		mstKriteria.logger.Zap.Error(err)
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

func (mstKriteria MstKriteriaController) GetOne(c *gin.Context) {
	requests := models.MstKriteriaRequest{}

	if err := c.Bind(&requests); err != nil {
		mstKriteria.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mstKriteria.service.GetOne(requests.ID)
	if err != nil {
		mstKriteria.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mstKriteria MstKriteriaController) Store(c *gin.Context) {
	data := models.MstKriteriaRequest{}

	if err := c.Bind(&data); err != nil {
		mstKriteria.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, msg, err := mstKriteria.service.Store(&data)

	if err != nil {
		mstKriteria.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	if !status {
		lib.ReturnToJson(c, 200, "400", msg, false)
		return
	}

	lib.ReturnToJson(c, 200, "200", msg, true)
}

func (mstKriteria MstKriteriaController) Update(c *gin.Context) {
	data := models.MstKriteriaRequest{}

	if err := c.Bind(&data); err != nil {
		mstKriteria.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := mstKriteria.service.Update(&data)

	if err != nil {
		mstKriteria.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	if !status {
		lib.ReturnToJson(c, 200, "400", "Criteria '"+data.Criteria+"' sudah ada !", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Update Data Berhasil", true)
}

func (mstKriteria MstKriteriaController) Delete(c *gin.Context) {
	requests := models.MstKriteriaRequest{}

	if err := c.Bind(&requests); err != nil {
		mstKriteria.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	if err := mstKriteria.service.Delete(requests.ID); err != nil {
		mstKriteria.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Data berhasil dihapus", true)
}

func (mstKriteria MstKriteriaController) GetKodeCriteria(c *gin.Context) {
	data, err := mstKriteria.service.GetKodeCriteria()
	if err != nil {
		mstKriteria.logger.Zap.Error(err)
	}

	if len(data) == 0 {
		mstKriteria.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data Tidak ditemukan", nil)
		return
	}

	KodeCriteria := data[0].KodeCriteria

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", KodeCriteria)

}

func (mstKriteria MstKriteriaController) GetCriteriaById(c *gin.Context) {
	requests := models.CriteriaRequestById{}
	if err := c.Bind(&requests); err != nil {
		mstKriteria.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, err := mstKriteria.service.GetCriteriaById(requests)

	if err != nil {
		mstKriteria.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "200", "Internal Error", "")
		return
	}

	if len(datas) == 0 {
		mstKriteria.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data tidak ditemukan", nil)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas)
}

func (mstKriteria MstKriteriaController) GetCriteriaByPeriode(c *gin.Context) {
	requests := models.PeriodeRequest{}
	if err := c.Bind(&requests); err != nil {
		mstKriteria.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, err := mstKriteria.service.GetCriteriaByPeriode(requests)

	if err != nil {
		mstKriteria.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "200", "Internal Error", "")
		return
	}

	if len(datas) == 0 {
		mstKriteria.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data tidak ditemukan", nil)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas)
}
