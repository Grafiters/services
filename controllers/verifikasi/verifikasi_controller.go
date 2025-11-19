package controller

import (
	"database/sql"
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/verifikasi"
	services "riskmanagement/services/verifikasi"

	"github.com/gin-gonic/gin"

	"gitlab.com/golang-package-library/logger"
)

type VerifikasiController struct {
	logger  logger.Logger
	service services.VerifikasiDefinition
}

func NewVerifikasiController(
	verifikasiService services.VerifikasiDefinition,
	logger logger.Logger,
) VerifikasiController {
	return VerifikasiController{
		service: verifikasiService,
		logger:  logger,
	}
}

func (verifikasi VerifikasiController) GetAll(c *gin.Context) {
	datas, err := verifikasi.service.GetAll()

	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "200", "Internal Error", "")
		return
	}

	if len(datas) == 0 {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data tidak ditemukan", nil)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas)
}

func (verifikasi VerifikasiController) GetListData(c *gin.Context) {
	datas, err := verifikasi.service.GetListData()

	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "200", "Internal Error", "")
		return
	}

	if len(datas) == 0 {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data tidak ditemukan", nil)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas)
}

func (verifikasi VerifikasiController) StoreDraft(c *gin.Context) {
	data := models.VerifikasiRequest{}

	if err := c.Bind(&data); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	status, err := verifikasi.service.StoreDraft(data)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if !status {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error status", err.Error())
		return
	}

	// getLastData, err := verifikasi.service.GetLastID()

	lib.ReturnToJson(c, 200, "200", "Berhasil disimpan sebagai draft", true)
}

func (verifikasi VerifikasiController) GetOne(c *gin.Context) {
	// paramID := c.Param("id")
	// id, err := strconv.Atoi(paramID)
	requests := models.VerfikasiRequestID{}

	if err := c.Bind(&requests); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	data, status, err := verifikasi.service.GetOne(int64(requests.ID))
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if !status {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data tidak ditemukan", nil)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inqueri data berhasil", data)

}

func (verifikasi VerifikasiController) DeleteLampiranVerifikasi(c *gin.Context) {
	data := models.VerifikasiFileRequest{}

	if err := c.Bind(&data); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	status, err := verifikasi.service.DeleteLampiranVerifikasi(&data)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if !status {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data Gagal dihapus", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Delete data berhasil", true)
}

func (verifikasi VerifikasiController) Delete(c *gin.Context) {
	data := models.VerifikasiRequestUpdateMaintain{}

	if err := c.Bind(&data); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	status, err := verifikasi.service.Delete(&data)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if !status {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data Gagal dihapus", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Delete data berhasil", true)
}

func (verifikasi VerifikasiController) DeleteRiskControl(c *gin.Context) {
	request := models.VerifikasiRiskControl{}

	if err := c.Bind(&request); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	status, err := verifikasi.service.DeleteRiskControl(request)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if !status {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data Gagal dihapus", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Delete data berhasil", true)
}

func (verifikasi VerifikasiController) KonfirmSave(c *gin.Context) {
	data := models.VerifikasiUpdateMaintain{}

	if err := c.Bind(&data); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	status, err := verifikasi.service.KonfirmSave(&data)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if !status {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data Gagal disimpan", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Berhasil menyimpan data", true)
}

func (verifikasi VerifikasiController) UpdateAllVerifikasi(c *gin.Context) {
	data := models.VerifikasiRequestMaintain{}

	if err := c.Bind(&data); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, message, err := verifikasi.service.UpdateAllVerifikasi(&data)
	if err != nil && !status {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", message, status)
		return
	}

	// if !status {
	// 	verifikasi.logger.Zap.Error(err)
	// 	lib.ReturnToJson(c, 200, "500", "Data Gagal Diupdate : ", false)
	// 	return
	// }

	lib.ReturnToJson(c, 200, "200", "Update data berhasil", true)
}

func (verifikasi VerifikasiController) FilterVerifikasi(c *gin.Context) {
	requests := models.VerifikasiFilterRequest{}

	if err := c.Bind(&requests); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := verifikasi.service.FilterVerifikasi(requests)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
	}

	if pagination.Total == 0 {
		lib.ReturnToJson(c, 200, "404", "Data Kosong", nil)
		return
	}

	// if len(datas) == 0 {
	// 	verifikasi.logger.Zap.Error(err)
	// 	lib.ReturnToJson(c, 200, "404", "Data Tidak Ditemukan", datas)
	// 	return
	// }

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	fmt.Println("Filter Data =>", datas)
	// lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", datas, pagination)
}

func (verifikasi VerifikasiController) GetDataWithPagination(c *gin.Context) {
	requests := models.VerifikasiPagination{}

	if err := c.Bind(&requests); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := verifikasi.service.GetDataWithPagination(requests)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
	}

	if pagination.Total == 0 {
		lib.ReturnToJson(c, 200, "404", "Data Kosong", nil)
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	// fmt.Println("Filter Data =>", datas)
	// lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", datas, pagination)
}

func (verifikasi VerifikasiController) GetNoPelaporan(c *gin.Context) {
	requests := models.NoPalaporanRequest{}
	today := lib.GetTimeNow("date2")
	// timeNow := lib.GetTimeNow("timestime")

	if err := c.Bind(&requests); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	datas, err := verifikasi.service.GetNoPelaporan(requests)

	if err != nil {
		verifikasi.logger.Zap.Error(err)
	}

	if len(datas) == 0 {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data Tidak Ditemukan", nil)
		return
	}

	counter := datas[0].NoPelaporan
	fmt.Println("ORGEH", datas[0].ORGEH)
	fmt.Println("DATE", today)
	fmt.Println("Counter", counter)

	NoPelaporan := "VER-" + datas[0].ORGEH + "-" + today + "-" + counter

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", NoPelaporan)
}

func (verifikasi VerifikasiController) FilterReport(c *gin.Context) {
	request := models.VerifikasiFilterReport{}

	if err := c.Bind(&request); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := verifikasi.service.FilterReport(request)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
	}

	if pagination.Total == 0 {
		lib.ReturnToJson(c, 200, "404", "Data Kosong", nil)
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	fmt.Println("Filter Data =>", datas)
	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", datas, pagination)

}

func (verifikasi VerifikasiController) StoreSimpan(c *gin.Context) {
	data := models.VerifikasiRequest{}

	fmt.Println("controller =>", data.UsulanPerbaikan)

	if err := c.Bind(&data); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	if data.NoPelaporan != "" {
		status, message, err := verifikasi.service.StoreSimpan(data)
		if !status {
			verifikasi.logger.Zap.Error(err)
			lib.ReturnToJson(c, 200, "500", message, status)
			return
		}

		// getLastData, err := verifikasi.service.GetLastID()
		var messageValid string
		if data.Action == "Draft" {
			messageValid = "Berhasil menyimpan draft"
		} else {
			messageValid = "Input data berhasil"
		}

		lib.ReturnToJson(c, 200, "200", messageValid, true)
	} else {
		lib.ReturnToJson(c, 200, "400", "Data Gagal disimpan, nomor pelaporan kosong!", false)
	}
}

func (verifikasi VerifikasiController) VerifikasiReportFilter(c *gin.Context) {
	data := models.VerifikasiFilterReportRequest{}

	if err := c.Bind(&data); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), data)
		return
	}

	datas, pagination, err := verifikasi.service.VerifikasiReportFilter(data)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", datas, pagination)
}

func (verifikasi VerifikasiController) VerifikasiReportWithWeaknessOnlyFilter(c *gin.Context) {
	data := models.VerifikasiFilterReportRequest{}

	if err := c.Bind(&data); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), data)
		return
	}

	datas, pagination, err := verifikasi.service.VerifikasiReportWithWeaknessOnlyFilter(data)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", datas, pagination)
}

func (verifikasi VerifikasiController) VerifikasiReportWithNonWeaknessOnlyFilter(c *gin.Context) {
	data := models.VerifikasiFilterReportRequest{}

	if err := c.Bind(&data); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), data)
		return
	}

	datas, pagination, err := verifikasi.service.VerifikasiReportWithNonWeaknessOnlyFilter(data)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", datas, pagination)
}

func (verifikasi VerifikasiController) VerifikasiReportFilterComplete(c *gin.Context) {
	data := models.VerifikasiFilterReportRequest{}

	if err := c.Bind(&data); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), data)
		return
	}

	datas, pagination, err := verifikasi.service.VerifikasiReportFilterComplete(data)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", datas, pagination)
}

func (verifikasi VerifikasiController) VerifikasiReportDetail(c *gin.Context) {
	data := models.VerifikasiReportDetailRequest{}

	if err := c.Bind(&data); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), data)
		return
	}

	datas, err := verifikasi.service.VerifikasiReportDetail(data)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	fmt.Println("")
	fmt.Println("response controller")
	fmt.Println(datas)

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}

func (verifikasi VerifikasiController) RiskControlByVerificationId(c *gin.Context) {
	data := models.DataRiskControlRequest{}

	if err := c.Bind(&data); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), data)
		return
	}

	datas, pagination, err := verifikasi.service.RiskControlByVerificationId(data)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", datas, pagination)
}

func (verifikasi VerifikasiController) GetRiskIndicatorAsMateri(c *gin.Context) {
	data := models.VerifikasiFilterReportRequest{}

	if err := c.Bind(&data); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), data)
		return
	}

	datas, err := verifikasi.service.GetRiskIndicatorAsMateri(data)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}

func (verifikasi VerifikasiController) VerificationReportByUkerFilter(c *gin.Context) {
	data := models.VerificationFilterReportByUkerRequest{}

	if err := c.Bind(&data); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), data)
		return
	}

	datas, pagination, err := verifikasi.service.VerificationReportByUkerFilter(data)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", datas, pagination)
}

func (verifikasi VerifikasiController) VerificationReportFilterByUkerComplete(c *gin.Context) {
	data := models.VerificationFilterReportByUkerRequest{}

	if err := c.Bind(&data); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), data)
		return
	}

	datas, pagination, err := verifikasi.service.VerificationReportFilterByUkerComplete(data)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", datas, pagination)
}

func (verifikasi VerifikasiController) VerifikasiReportByFraudIndicatorFilter(c *gin.Context) {
	data := models.VerificationFilterReportByUkerRequest{}

	if err := c.Bind(&data); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), data)
		return
	}

	datas, pagination, err := verifikasi.service.VerifikasiReportByFraudIndicatorFilter(data)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", datas, pagination)
}

func (verifikasi VerifikasiController) VerificationReportFilterByFraudIndicatorComplete(c *gin.Context) {
	data := models.VerificationFilterReportByUkerRequest{}

	if err := c.Bind(&data); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), data)
		return
	}

	datas, pagination, err := verifikasi.service.VerificationReportFilterByFraudIndicatorComplete(data)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", datas, pagination)
}

func (verifikasi VerifikasiController) VerifikasiReportMateriList(c *gin.Context) {
	request := models.VerifikasiMateriRequest{}

	if err := c.Bind(&request); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	datas, err := verifikasi.service.VerifikasiReportMateriList(request)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	fmt.Println("Filter Data =>", datas)
	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}

// Versioning 1.0.0.1 by panji 31/08/2023
func (verifikasi VerifikasiController) DeleteAnomaliByID(c *gin.Context) {
	request := models.VerifikasiAnomaliData{}

	if err := c.Bind(&request); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	status, err := verifikasi.service.DeleteAnomaliByID(&request)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if !status {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data Gagal dihapus", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Delete data berhasil", true)
}

func (verif VerifikasiController) VerifikasiReportList(c *gin.Context) {
	request := models.VerifikasiReportListRequest{}

	if err := c.Bind(&request); err != nil {
		verif.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := verif.service.VerifikasiReportList(request)
	if err != nil {
		verif.logger.Zap.Error(err)
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	fmt.Println("Filter Data =>", datas)
	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", datas, pagination)
}

func (verif VerifikasiController) RptRakapitulasiBCV(c *gin.Context) {
	request := models.RptRekapitulasiBCVRequest{}

	if err := c.Bind(&request); err != nil {
		verif.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := verif.service.RptRekapitulasiBCV(request)
	if err != nil {
		verif.logger.Zap.Error(err)
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	fmt.Println("Filter Data =>", datas)
	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", datas, pagination)
}

func (verif VerifikasiController) RptRekomendasiRisk(c *gin.Context) {
	request := models.RptRekomendasiRiskRequest{}

	if err := c.Bind(&request); err != nil {
		verif.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := verif.service.RptRekomendasiRisk(request)
	if err != nil {
		verif.logger.Zap.Error(err)
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	fmt.Println("Filter Data =>", datas)
	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", datas, pagination)
}

func (verifikasi VerifikasiController) VerificationReportUkerByAllActivity(c *gin.Context) {
	data := models.VerificationFilterReportByUkerRequest{}

	if err := c.Bind(&data); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), data)
		return
	}

	datas, pagination, err := verifikasi.service.VerificationReportUkerByAllActivity(data)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	fmt.Println("responses - controller - datas", datas)
	fmt.Println("responses - controller - err", err)

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", datas, pagination)
}

func (verifikasi VerifikasiController) VerificationReportUkerByAllActivityComplete(c *gin.Context) {
	data := models.VerificationFilterReportByUkerRequest{}

	if err := c.Bind(&data); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), data)
		return
	}

	datas, pagination, err := verifikasi.service.VerificationReportUkerByAllActivityComplete(data)

	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", datas, pagination)

}

func (verifikasi VerifikasiController) VerificationReportUkerByAllActivityCompleteWithRiskIssue(c *gin.Context) {
	data := models.VerificationFilterReportByUkerRequest{}

	if err := c.Bind(&data); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), data)
		return
	}

	datas, pagination, err := verifikasi.service.VerificationReportUkerByAllActivityCompleteWithRiskIssue(data)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", datas, pagination)
}

func (verif VerifikasiController) ValidasiVerifikasi(c *gin.Context) {
	request := models.ValidasiVerifikasiRequest{}

	if err := c.Bind(&request); err != nil {
		verif.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := verif.service.ValidasiVerifikasi(request)
	if err != nil {
		verif.logger.Zap.Error(err)
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	fmt.Println("Filter Data =>", datas)
	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", datas, pagination)
}

func (verifikasi VerifikasiController) AcceptValidasi(c *gin.Context) {
	data := models.AcceptValidasiRequest{}

	if err := c.Bind(&data); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := verifikasi.service.AcceptValidasi(&data)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error : ", "")
		return
	}

	if !status {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Gagal Divalidasi: ", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Berhasil Divalidasi", true)
}

func (verifikasi VerifikasiController) UpdateStatusVerifikasi(c *gin.Context) {
	data := models.UpdateStatusVerifikasi{}

	if err := c.Bind(&data); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := verifikasi.service.UpdateStatusVerifikasi(&data)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error : ", "")
		return
	}

	if !status {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Gagal Divalidasi: ", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Berhasil Divalidasi", true)
}

func (verifikasi VerifikasiController) RejectValidasi(c *gin.Context) {
	data := models.RejectValidasiRequest{}

	if err := c.Bind(&data); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := verifikasi.service.RejectValidasi(&data)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error : ", "")
		return
	}

	if !status {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Gagal Direject: ", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Reject data berhasil", true)
}

func (verifikasi VerifikasiController) GetRtlIndikasiFraud(c *gin.Context) {
	data := models.ReqRtlIndikasiFraud{}

	if err := c.Bind(&data); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), data)
		return
	}

	datas, pagination, err := verifikasi.service.GetRtlIndikasiFraud(data)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", datas, pagination)
}

func (verifikasi VerifikasiController) ValidasiVerifikasiDetailData(c *gin.Context) {
	request := models.VerifikasiReportDetailRequest{}

	if err := c.Bind(&request); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), nil)
		return
	}

	datas, err := verifikasi.service.ValidasiVerifikasiDetailData(request)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}

func (verifikasi VerifikasiController) GetRekomendasiTindakLanjut(c *gin.Context) {
	request := models.RTLRequest{}

	if err := c.Bind(&request); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), nil)
		return
	}

	datas, err := verifikasi.service.GetRekomendasiTindakLanjut(request)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if len(datas) < 1 {
		lib.ReturnToJson(c, 200, "404", "Data tidak ditemukan", nil)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}

func (verifikasi VerifikasiController) DeletePenyebabKejadian(c *gin.Context) {
	request := models.VerifikasiRiskControl{}

	if err := c.Bind(&request); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	status, err := verifikasi.service.DeletePenyebabKejadian(request.ID)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if !status {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data Gagal dihapus", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Delete data berhasil", true)
}

func (verifikasi VerifikasiController) VerifikasiSummaryRpt(c *gin.Context) {
	request := models.SummaryVerifikasiRequest{}

	if err := c.Bind(&request); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), nil)
		return
	}

	datas, totalRows, err := verifikasi.service.VerifikasiSummaryRpt(request)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if len(datas) < 1 {
		lib.ReturnToJson(c, 200, "404", "Data tidak ditemukan", nil)
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", datas, totalRows)
}

func (verifikasi VerifikasiController) VerifikasiFrekuensiRpt(c *gin.Context) {
	request := models.FrekuensiVerifikasiRequest{}

	if err := c.Bind(&request); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), nil)
		return
	}

	datas, totalRows, err := verifikasi.service.VerifikasiFrekuensiRpt(request)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if len(datas) < 1 {
		lib.ReturnToJson(c, 200, "404", "Data tidak ditemukan", nil)
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquiry Data Berhasil", datas, totalRows)
}
