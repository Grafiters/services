package organisasi

import (
	"riskmanagement/lib"
	models "riskmanagement/models/organisasi"
	service "riskmanagement/services/organisasi"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type OrganisasiController struct {
	logger  logger.Logger
	service service.OrganisasiServiceDefinition
}

func NewOrganisasiController(
	logger logger.Logger,
	service service.OrganisasiServiceDefinition,
) OrganisasiController {
	return OrganisasiController{
		logger:  logger,
		service: service,
	}
}

func (o OrganisasiController) GetCostCenter(c *gin.Context) {
	request := models.CostCenterRequest{}

	if err := c.Bind(&request); err != nil {
		o.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai", false)
		return
	}

	data, err := o.service.GetCostCenter(request)
	if err != nil {
		o.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}

func (o OrganisasiController) GetOrgUnit(c *gin.Context) {
	request := models.DepartmentRequest{}

	if err := c.Bind(&request); err != nil {
		o.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai", false)
		return
	}

	data, err := o.service.GetOrgUnit(request)
	if err != nil {
		o.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error : "+err.Error(), false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}

func (o OrganisasiController) GetHilfm(c *gin.Context) {
	request := models.JabatanRequest{}

	if err := c.Bind(&request); err != nil {
		o.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai", false)
		return
	}

	data, err := o.service.GetHilfm(request)
	if err != nil {
		o.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error : "+err.Error(), false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}
