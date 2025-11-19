package verifikasirealisasi

import (
	"riskmanagement/lib"
	models "riskmanagement/models/verifikasireportrealisasi"
	service "riskmanagement/services/verifikasireportrealisasi"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type VerifikasiReportRealisasiController struct {
	logger  logger.Logger
	service service.VerifikasiReportRealisasiServiceDefinition
}

func NewVerifikasiReportRealisasiController(
	logger logger.Logger,
	service service.VerifikasiReportRealisasiServiceDefinition,
) VerifikasiReportRealisasiController {
	return VerifikasiReportRealisasiController{
		logger:  logger,
		service: service,
	}
}

func (verifikasi VerifikasiReportRealisasiController) ReportRealisasiKreditListFilter(c *gin.Context) {
	data := models.ReportRealisasiKreditListRequest{}

	if err := c.Bind(&data); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), data)
		return
	}

	datas, pagination, err := verifikasi.service.ReportRealisasiKreditListFilter(data)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", datas, pagination)
}

func (verifikasi VerifikasiReportRealisasiController) ReportRealisasiKreditSummaryFilter(c *gin.Context) {
	data := models.ReportRealisasiKreditSummaryRequest{}

	if err := c.Bind(&data); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), data)
		return
	}

	datas, pagination, err := verifikasi.service.ReportRealisasiKreditSummaryFilter(data)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", datas, pagination)
}

func (verifikasi VerifikasiReportRealisasiController) GetAllSegmentRealisasiKredit(c *gin.Context) {
	data := models.SegmentRealisasiKreditRequest{}

	if err := c.Bind(&data); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), data)
		return
	}

	datas, pagination, err := verifikasi.service.GetAllSegmentRealisasiKredit(data)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", datas, pagination)
}
