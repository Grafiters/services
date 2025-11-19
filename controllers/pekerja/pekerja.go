package pekerja

import (
	"riskmanagement/lib"
	"riskmanagement/models/pekerja"
	service "riskmanagement/services/pekerja"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type PekerjaController struct {
	logger  logger.Logger
	service service.PekerjaDefinition
}

func NewPekerjaController(
	logger logger.Logger,
	service service.PekerjaDefinition,
) PekerjaController {
	return PekerjaController{
		logger:  logger,
		service: service,
	}
}

func (p PekerjaController) GetAllPekerjaBranch(c *gin.Context) {
	request := pekerja.PekerjaUkerRequest{}

	if err := c.Bind(&request); err != nil {
		p.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai", false)
		return
	}

	data, err := p.service.GetAllPekerjaBranch(&request)
	if err != nil {
		p.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}

func (p PekerjaController) GetApproval(c *gin.Context) {
	request := pekerja.RequestApproval{}

	if err := c.Bind(&request); err != nil {
		// p.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai", err.Error())
		return
	}

	data, err := p.service.GetApproval(&request)
	if err != nil {
		// p.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}
