package monitoring

import (
	"riskmanagement/lib"
	models "riskmanagement/models/monitoring"
	service "riskmanagement/services/monitoring"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type MonitoringController struct {
	logger  logger.Logger
	service service.MonitoringServicesDefinition
}

func NewMonitoringController(MonitoringService service.MonitoringServicesDefinition, logger logger.Logger) MonitoringController {
	return MonitoringController{
		logger:  logger,
		service: MonitoringService,
	}
}

func (m MonitoringController) GetMonitoringTasklist(c *gin.Context) {
	data := models.MonitoringTasklistRequest{}

	if err := c.Bind(&data); err != nil {
		m.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	if data.JenisReport == "1" {
		result, totalRow, err := m.service.GetMonitoringPekerja(data)
		if err != nil {
			lib.ReturnToJson(c, 200, "400", "Error :"+err.Error(), false)
			return
		}

		lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery data berhasil", result, totalRow)
	} else if data.JenisReport == "2" {
		result, totalRow, err := m.service.GetMonitoringTasklistUker(data)
		if err != nil {
			m.logger.Zap.Error(err)
			lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
			return
		}

		// lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery data berhasil", data, 0)
		lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery data berhasil", result, totalRow)
	}

}
