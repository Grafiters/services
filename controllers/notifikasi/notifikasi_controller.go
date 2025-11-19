package notifikasi

import (
	"database/sql"
	"riskmanagement/lib"
	models "riskmanagement/models/notifikasi"
	services "riskmanagement/services/notifikasi"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type NotifikasiController struct {
	logger  logger.Logger
	service services.NotifikasiServicesDefinition
}

func NewNotifikasiController(
	NotifikasiService services.NotifikasiServicesDefinition,
	logger logger.Logger,
) NotifikasiController {
	return NotifikasiController{
		service: NotifikasiService,
		logger:  logger,
	}
}

func (n NotifikasiController) CreateNotifikasi(c *gin.Context) {

	request := models.TasklistNotifikasiRequest{}

	if err := c.Bind(&request); err != nil {
		n.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	data, err := n.service.Store(request)
	if err != nil {
		n.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Create data berhasil", data)
}

func (n NotifikasiController) GetListNotifikasi(c *gin.Context) {

	data := models.NotifikasiRequest{}

	if err := c.Bind(&data); err != nil {
		n.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	result, pagination, err := n.service.GetNotifikasi(data)
	if err != nil {
		n.logger.Zap.Error(err)
	}

	if pagination.Total == 0 {
		lib.ReturnToJson(c, 200, "404", "Data Kosong", data)
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", result, pagination)
}

func (n NotifikasiController) GetTotalNotifikasi(c *gin.Context) {
	dataRequest := models.NotifikasiTotalRequest{}

	if err := c.Bind(&dataRequest); err != nil {
		n.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	data, totalRow, err := n.service.GetTotalNotifikasi(dataRequest)
	if err != nil {
		n.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery data berhasil", data, totalRow)
}

func (n NotifikasiController) UpdateStatusNotifikasi(c *gin.Context) {

	request := models.NotifikasiUpdateStatus{}

	if err := c.Bind(&request); err != nil {
		n.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	data, err := n.service.UpdateStatus(request)
	if err != nil {
		n.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Update data berhasil", data)
}

func (n NotifikasiController) DeleteStatusNotifikasi(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		n.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	request := models.NotifikasiUpdateStatus{}

	if err := c.Bind(&request); err != nil {
		n.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	data, err := n.service.DeleteStatus(id)
	if err != nil {
		n.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Update data berhasil", data)
}

func (n NotifikasiController) Delete(c *gin.Context) {
	data := models.NotifikasiUpdateStatus{}

	if err := c.Bind(&data); err != nil {
		n.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai:"+err.Error(), "")
		return
	}

	status, err := n.service.DeleteStatus(data.ID)
	if err != nil {
		n.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if !status {
		n.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data Gagal dihapus", false)
	}

	lib.ReturnToJson(c, 200, "200", "Data berhasil dihapus", true)
}
