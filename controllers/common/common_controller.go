package common

import (
	"riskmanagement/lib"
	models "riskmanagement/models/common"
	services "riskmanagement/services/common"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type CommonController struct {
	logger  logger.Logger
	service services.CommonDefinition
}

func NewCommonController(
	CommonService services.CommonDefinition,
	logger logger.Logger,
) CommonController {
	return CommonController{
		service: CommonService,
		logger:  logger,
	}
}

func (common CommonController) GetNpNamaFilter(c *gin.Context) {

	data := models.KeywordRequest{}

	if err := c.Bind(&data); err != nil {
		common.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	result, err := common.service.FilterPnNama(data)
	if err != nil {
		common.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", result)
}

func (common CommonController) GetKanwilFilter(c *gin.Context) {

	data := models.KeywordRequest{}

	if err := c.Bind(&data); err != nil {
		common.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	result, err := common.service.FilterKanwil(data)
	if err != nil {
		common.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", result)
}

func (common CommonController) GetKancaFilter(c *gin.Context) {

	data := models.KeywordRequest{}

	if err := c.Bind(&data); err != nil {
		common.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	result, err := common.service.FilterKanca(data)
	if err != nil {
		common.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", result)
}

func (common CommonController) GetUkerFilter(c *gin.Context) {

	data := models.KeywordRequest{}

	if err := c.Bind(&data); err != nil {
		common.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	result, err := common.service.FilterUker(data)
	if err != nil {
		common.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", result)
}

func (common CommonController) FilterRiskEventByActifityAndPoduct(c *gin.Context) {

	data := models.RiskEventRequest{}

	if err := c.Bind(&data); err != nil {
		common.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	result, err := common.service.FilterRiskEventByActifityAndPoduct(data)
	if err != nil {
		common.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", result)
}

func (common CommonController) FilterRiskIndokatorByRiskEventID(c *gin.Context) {

	data := models.RiskIndikatorRequest{}

	if err := c.Bind(&data); err != nil {
		common.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	result, err := common.service.FilterRiskIndikatorByRiskEvent(data)
	if err != nil {
		common.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", result)
}

func (common CommonController) GetRRMHeadFilter(c *gin.Context) {
	data := models.RRMHeadRequest{}

	if err := c.Bind(&data); err != nil {
		common.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	result, err := common.service.FilterRRMHead(data)
	if err != nil {
		common.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Internal Error : "+err.Error(), false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", result)
}

func (common CommonController) GetPimpinanUker(c *gin.Context) {
	data := models.PimpinanUkerRequest{}

	if err := c.Bind(&data); err != nil {
		common.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	result, err := common.service.FilterPimpinanUker(data)
	if err != nil {
		common.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Internal Error : "+err.Error(), false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", result)
}

func (q CommonController) GetApprovalResponse(c *gin.Context) {
	request := models.CommonRequest{}

	if err := c.Bind(&request); err != nil {
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := q.service.GetApprovalResponse(request)
	if err != nil {
		lib.ReturnToJson(c, 200, "400", "Error :"+err.Error(), false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)

}

// Enhance MQ
func (q CommonController) GetMstDataOption(c *gin.Context) {
	data, err := q.service.GetMstDataOption()

	if err != nil {
		lib.ReturnToJson(c, 200, "400", "Error : "+err.Error(), false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}

func (common CommonController) SearchBrc(c *gin.Context) {
	requests := models.BrcKeywordRequest{}

	if err := c.Bind(&requests); err != nil {
		common.logger.Zap.Error("controller : ", err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, pagination, err := common.service.SearchBrc(requests)

	if err != nil {
		common.logger.Zap.Error("controller : ", err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if pagination.Total == 0 {
		common.logger.Zap.Error("controller : ", "Data Kosong")
		lib.ReturnToJson(c, 200, "404", "Data Kosong", nil)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}
