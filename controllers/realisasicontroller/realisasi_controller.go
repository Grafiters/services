package realisasicontroller

import (
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/RealisasiModels"
	service "riskmanagement/services/realisasiservice"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type RealisasiController struct {
	logger           logger.Logger
	realisasiService service.RealisasiDefinition
}

func NewRealisasiController(
	logger logger.Logger,
	realisasiService service.RealisasiDefinition,
) RealisasiController {
	return RealisasiController{
		logger:           logger,
		realisasiService: realisasiService,
	}
}

func (r RealisasiController) GetDataParameter(c *gin.Context) {
	request := models.ParameterGetHeaderRequest{}
	if err := c.Bind(&request); err != nil {
		r.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai :"+err.Error(), "")
		return
	}

	response, err := r.realisasiService.GetDataParameter(&request)
	fmt.Println("data =>", response.Data)
	fmt.Println("pagintion =>", response.Pagination)

	// pagination := response.Pagination

	if err != nil {
		r.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal error", err.Error())
		return
	}

	if response.Pagination == 0 {
		// r.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data tidak ditemukan", "")
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery berhasil", response.Data, response.Pagination)
}

func (r RealisasiController) StoreDataParameter(c *gin.Context) {
	request := models.ParameterStoreHeaderRequest{}

	if err := c.Bind(&request); err != nil {
		r.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai :"+err.Error(), "")
		return
	}

	status, err := r.realisasiService.StoreDataParameter(&request)

	if err != nil {
		lib.ReturnToJson(c, 200, "500", "Input data gagal", status.Data)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Input data berhasil", status.Data)
}

func (r RealisasiController) GetDataRevisiUker(c *gin.Context) {
	request := models.RevisiUkerGetHeaderRequest{}
	if err := c.Bind(&request); err != nil {
		r.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai :"+err.Error(), "")
		return
	}

	response, err := r.realisasiService.GetDataRevisiUker(&request)
	fmt.Println("data =>", response.Data)
	fmt.Println("pagintion =>", response.Pagination)

	// pagination := response.Pagination

	if err != nil {
		r.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal error", err.Error())
		return
	}

	if response.Pagination == 0 {
		// r.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data tidak ditemukan", "")
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery berhasil", response.Data, response.Pagination)
}

func (r RealisasiController) StoreDataRevisiUker(c *gin.Context) {
	request := models.RevisiUkerStoreHeaderRequest{}

	if err := c.Bind(&request); err != nil {
		r.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai :"+err.Error(), "")
		return
	}

	status, err := r.realisasiService.StoreDataRevisiUker(&request)

	if err != nil || status.Status == "500" {
		lib.ReturnToJson(c, 200, "500", "Input data gagal", status.Data)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Input data berhasil", status.Data)
}

func (r RealisasiController) DeleteDataRevisiUker(c *gin.Context) {
	request := models.RevisiUkerDeleteHeaderRequest{}

	if err := c.Bind(&request); err != nil {
		r.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai :"+err.Error(), "")
		return
	}

	status, err := r.realisasiService.DeleteDataRevisiUker(&request)

	if err != nil {
		lib.ReturnToJson(c, 200, "500", "Delete data gagal", status.Data)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Delete data berhasil", status.Data)
}

func (r RealisasiController) GetDataRealisasi(c *gin.Context) {
	request := models.RealisasiGetHeaderRequest{}
	if err := c.Bind(&request); err != nil {
		r.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai :"+err.Error(), "")
		return
	}

	response, err := r.realisasiService.GetDataRealisasi(&request)
	fmt.Println("data =>", response.Data)
	fmt.Println("pagintion =>", response.Pagination)

	// pagination := response.Pagination

	if err != nil {
		r.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal error", err.Error())
		return
	}

	if response.Pagination == 0 {
		// r.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data tidak ditemukan", "")
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery berhasil", response.Data, response.Pagination)
}

func (r RealisasiController) UpdateFlagVerifikasi(c *gin.Context) {
	request := models.RealisasiUpdateFlagRequest{}

	if err := c.Bind(&request); err != nil {
		r.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai :"+err.Error(), "")
		return
	}

	status, err := r.realisasiService.UpdateFlagVerifikasi(&request)

	if err != nil {
		lib.ReturnToJson(c, 200, "500", "Update data gagal", status.Data)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Update data berhasil", status.Data)
}
