package jenistask

import (
	"riskmanagement/lib"
	models "riskmanagement/models/jenistask"
	services "riskmanagement/services/jenistask"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type JenisTaskController struct {
	logger  logger.Logger
	service services.JenisTaskDefinition
}

func NewJenisTaskController(
	JenisTaskService services.JenisTaskDefinition,
	logger logger.Logger,
) JenisTaskController {
	return JenisTaskController{
		service: JenisTaskService,
		logger:  logger,
	}
}

func (task JenisTaskController) GetData(c *gin.Context) {
	requests := models.TaskRequest{}

	if err := c.Bind(&requests); err != nil {
		task.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := task.service.GetData(&requests)
	if err != nil {
		task.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if len(data) == 0 {
		lib.ReturnToJson(c, 200, "404", "Data Kosong", "")
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}
