package kriteriapinjaman

import (
	"database/sql"
	"riskmanagement/lib"
	models "riskmanagement/models/kriteriapinjaman"
	services "riskmanagement/services/kriteriapinjaman"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type KriteriaPinjamanController struct {
	logger  logger.Logger
	service services.KriteriaPinjamanDefinition
}

func NewKriteriaPinjamanController(
	service services.KriteriaPinjamanDefinition,
	logger logger.Logger,
) KriteriaPinjamanController {
	return KriteriaPinjamanController{
		logger:  logger,
		service: service,
	}
}

func (k KriteriaPinjamanController) GetAll(c *gin.Context) {
	request := models.Paginate{}

	if err := c.Bind(&request); err != nil {
		k.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	data, pagination, err := k.service.GetAll(request)
	if err != nil {
		k.logger.Zap.Error(err)
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery data berhasil", data, pagination)
}
