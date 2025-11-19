package controller

import (
	"database/sql"
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/risktype"
	services "riskmanagement/services/risktype"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type RiskTypeController struct {
	logger  logger.Logger
	service services.RiskTypeDefinition
}

func NewRiskTypeController(
	RiskTypeService services.RiskTypeDefinition,
	logger logger.Logger,
) RiskTypeController {
	return RiskTypeController{
		service: RiskTypeService,
		logger:  logger,
	}
}

func (riskType RiskTypeController) GetAll(c *gin.Context) {
	datas, err := riskType.service.GetAll()
	if err != nil {
		riskType.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas)
}

func (rt RiskTypeController) GetAllWithPaginate(c *gin.Context) {
	requests := models.Paginate{}

	if err := c.Bind(&requests); err != nil {
		rt.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := rt.service.GetAllWithPaginate(requests)
	if err != nil {
		rt.logger.Zap.Error(err)
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
func (riskType RiskTypeController) GetOne(c *gin.Context) {
	requests := models.RiskTypeRequest{}

	if err := c.Bind(&requests); err != nil {
		riskType.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := riskType.service.GetOne(requests.ID)
	if err != nil {
		riskType.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}
func (riskType RiskTypeController) Store(c *gin.Context) {
	data := models.RiskTypeRequest{}

	if err := c.Bind(&data); err != nil {
		riskType.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	fmt.Println(data)
	if err := riskType.service.Store(&data); err != nil {
		riskType.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "200", "Internal error", data)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Input data berhasil", data)
}
func (riskType RiskTypeController) Update(c *gin.Context) {
	data := models.RiskTypeRequest{}

	if err := c.Bind(&data); err != nil {
		riskType.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	if err := riskType.service.Update(&data); err != nil {
		riskType.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "200", "Internal error", data)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Update data berhasil", data)
}
func (riskType RiskTypeController) Delete(c *gin.Context) {
	requests := models.RiskTypeRequest{}

	if err := c.Bind(&requests); err != nil {
		riskType.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	if err := riskType.service.Delete(requests.ID); err != nil {
		riskType.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Data berhasil dihapus", true)
}
