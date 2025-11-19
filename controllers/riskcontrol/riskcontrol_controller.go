package controller

import (
	"database/sql"
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/riskcontrol"
	services "riskmanagement/services/riskcontrol"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type RiskControlController struct {
	logger  logger.Logger
	service services.RiskControlDefinition
}

func NewRiskControlController(
	RiskControlService services.RiskControlDefinition,
	logger logger.Logger,
) RiskControlController {
	return RiskControlController{
		logger:  logger,
		service: RiskControlService,
	}
}

func (riskControl RiskControlController) GetAll(c *gin.Context) {
	datas, err := riskControl.service.GetAll()
	if err != nil {
		riskControl.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas)
}

func (rc RiskControlController) GetAllWithPaginate(c *gin.Context) {
	requests := models.Paginate{}

	if err := c.Bind(&requests); err != nil {
		rc.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := rc.service.GetAllWithPaginate(requests)
	if err != nil {
		rc.logger.Zap.Error(err)
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

func (riskControl RiskControlController) GetOne(c *gin.Context) {
	requests := models.RiskControlRequest{}

	if err := c.Bind(&requests); err != nil {
		riskControl.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := riskControl.service.GetOne(requests.ID)
	if err != nil {
		riskControl.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}
func (riskControl RiskControlController) Store(c *gin.Context) {
	data := models.RiskControlRequest{}

	if err := c.Bind(&data); err != nil {
		riskControl.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	fmt.Println(data)
	if err := riskControl.service.Store(&data); err != nil {
		riskControl.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "200", "Internal error", data)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Input data berhasil", data)
}
func (riskControl RiskControlController) Update(c *gin.Context) {
	data := models.RiskControlRequest{}

	if err := c.Bind(&data); err != nil {
		riskControl.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	if err := riskControl.service.Update(&data); err != nil {
		riskControl.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "200", "Internal error", data)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Update data berhasil", data)
}
func (riskControl RiskControlController) Delete(c *gin.Context) {
	requests := models.RiskControlRequest{}

	if err := c.Bind(&requests); err != nil {
		riskControl.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	if err := riskControl.service.Delete(requests.ID); err != nil {
		riskControl.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Data berhasil dihapus", true)
}
func (riskControl RiskControlController) GetKodeRiskControl(c *gin.Context) {
	datas, err := riskControl.service.GetKodeRiskControl()
	if err != nil {
		riskControl.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if len(datas) == 0 {
		riskControl.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data Tidak Ditemukan !", nil)
		return
	}

	KodeRiskControl := "C" + datas[0].KodeRiskControl

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", KodeRiskControl)
}

func (riskControl RiskControlController) SearchRiskIndicatorByIssue(c *gin.Context) {
	requests := models.KeywordRequest{}

	if err := c.Bind(&requests); err != nil {
		riskControl.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := riskControl.service.SearchRiskControlByIssue(requests)
	if err != nil {
		riskControl.logger.Zap.Error(err)
	}

	if pagination.Total == 0 {
		lib.ReturnToJson(c, 200, "404", "Data Tidak Ditemukan", nil)
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", datas, pagination)
}
