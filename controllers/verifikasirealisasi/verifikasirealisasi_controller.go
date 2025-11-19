package verifikasirealisasi

import (
	"riskmanagement/lib"
	models "riskmanagement/models/verifikasirealisasi"
	service "riskmanagement/services/verifikasirealisasi"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type VerifikasiRealisasiController struct {
	logger  logger.Logger
	service service.VerifikasiRealisasiServiceDefinition
}

func NewVerifikasiRealisasiController(
	logger logger.Logger,
	service service.VerifikasiRealisasiServiceDefinition,
) VerifikasiRealisasiController {
	return VerifikasiRealisasiController{
		logger:  logger,
		service: service,
	}
}

func (v VerifikasiRealisasiController) GetData(c *gin.Context) {
	request := models.VerifikasiRealisasiFilterRequest{}

	if err := c.Bind(&request); err != nil {
		v.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	data, totalData, err := v.service.GetData(request)

	if err != nil {
		v.logger.Zap.Error(err.Error())
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if totalData < 0 {
		lib.ReturnToJson(c, 200, "404", "Data Not Found", nil)
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Berhasil", data, totalData)
}

func (v VerifikasiRealisasiController) StoreVerifikasi(c *gin.Context) {
	request := models.VerifikasiRealisasiRequest{}

	if err := c.Bind(&request); err != nil {
		v.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	status, err := v.service.StoreVerifikasi(&request)

	if err != nil && !status {
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Input data berhasil", status)
}

func (v VerifikasiRealisasiController) GetOne(c *gin.Context) {
	requests := models.VerifikasiRealisasiRequestID{}

	if err := c.Bind(&requests); err != nil {
		v.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	data, status, err := v.service.GetOne(int64(requests.ID))
	if err != nil {
		v.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if !status {
		v.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data tidak ditemukan", nil)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", data)
}

func (v VerifikasiRealisasiController) Delete(c *gin.Context) {
	data := models.VerifikasiRealisasiRequest{}

	if err := c.Bind(&data); err != nil {
		v.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, sampledata, err := v.service.Delete(&data)
	if err != nil {
		v.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if !status {
		v.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data Gagal dihapus", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Delete data berhasil", sampledata)
}

func (v VerifikasiRealisasiController) Update(c *gin.Context) {
	data := models.VerifikasiRealisasiRequest{}

	if err := c.Bind(&data); err != nil {
		v.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai :"+err.Error(), "")
		return
	}

	status, err := v.service.Update(&data)
	if err != nil {
		v.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error : ", "")
		return
	}

	if !status {
		v.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data gagal diupdate :", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Update data berhasil", true)
}

func (v VerifikasiRealisasiController) GetNoPelaporan(c *gin.Context) {
	request := &models.NoPalaporanVerifikasiRealisasiRequest{}

	if err := c.Bind(request); err != nil {
		v.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	no_pelaporan, err := v.service.GetNoPelaporan(request)

	if err != nil {
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery berhasil", no_pelaporan)
}
