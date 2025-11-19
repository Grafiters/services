package datatematik

import (
	"riskmanagement/lib"
	models "riskmanagement/models/data_tematik"
	service "riskmanagement/services/data_tematik"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type DataTematikController struct {
	logger  logger.Logger
	service service.DataTematikServiceDefinition
}

func NewDataTematikController(
	logger logger.Logger,
	service service.DataTematikServiceDefinition,
) DataTematikController {
	return DataTematikController{
		logger:  logger,
		service: service,
	}
}

func (dt DataTematikController) GetSampleDataTematik(c *gin.Context) {
	request := models.DataTematikRequest{}

	if err := c.Bind(&request); err != nil {
		dt.logger.Zap.Error(err)
		lib.ReturnToJson(c, 400, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	data, totalData, err := dt.service.GetSampleDataTematik(request)

	if err != nil {
		dt.logger.Zap.Error(err)
		lib.ReturnToJson(c, 500, "500", "Internal Error", err)
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Berhasil", data, totalData)
}

func (dt DataTematikController) UpdateStatusDataSample(c *gin.Context) {
	requests := models.UpdaterData{}

	if err := c.Bind(&requests); err != nil {
		dt.logger.Zap.Error(err)
		lib.ReturnToJson(c, 400, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := dt.service.UpdateStatusDataSample(requests)
	if err != nil {
		dt.logger.Zap.Error("Error Update Data")
		lib.ReturnToJson(c, 500, "500", "Internal Error", status)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Update Berhasil", status)
}
