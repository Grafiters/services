package controllers

import (
	"database/sql"
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/briefing"
	services "riskmanagement/services/briefing"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type BriefingController struct {
	logger  logger.Logger
	service services.BriefingDefinition
}

func NewBriefingController(
	BriefingService services.BriefingDefinition,
	logger logger.Logger,
) BriefingController {
	return BriefingController{
		service: BriefingService,
		logger:  logger,
	}
}

func (briefing BriefingController) GetAll(c *gin.Context) {
	datas, err := briefing.service.GetAll()
	if err != nil {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if len(datas) == 0 {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data Tidak ditemukan !", nil)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas)
}

func (briefing BriefingController) GetData(c *gin.Context) {
	datas, err := briefing.service.GetData()
	if err != nil {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if len(datas) == 0 {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data Tidak ditemukan !", "")
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas)
}

func (briefing BriefingController) GetOne(c *gin.Context) {
	// paramID := c.Param("id")
	// id, err := strconv.Atoi(paramID)

	request := models.BriefingGetOneRequest{}
	if err := c.Bind(&request); err != nil {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), request)
		return
	}

	data, status, err := briefing.service.GetOne(int64(request.ID))
	if err != nil {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	if !status {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data Tidak Ditemukan", nil)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", data)
}

func (briefing BriefingController) Store(c *gin.Context) {
	data := models.BriefingRequest{}
	if err := c.Bind(&data); err != nil {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), data)
		return
	}
	// fmt.Println(data.)
	status, message, err := briefing.service.Store(data)
	if err != nil {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", message, err.Error())
		return
	}

	if !status {
		// briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", message, false)
		return
	}
	lib.ReturnToJson(c, 200, "200", message, true)
}

func (briefing BriefingController) StoreDraft(c *gin.Context) {
	data := models.BriefingRequest{}
	if err := c.Bind(&data); err != nil {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), data)
		return
	}
	// fmt.Println(data.)
	status, err := briefing.service.StoreDraft(data)
	if err != nil {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if !status {
		// briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Input Data Berhasil", true)
}

func (briefing BriefingController) Delete(c *gin.Context) {
	data := models.BriefingRequestUpdate{}

	if err := c.Bind(&data); err != nil {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := briefing.service.Delete(&data)
	if err != nil {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if !status {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data Gagal Dihapus", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Hapus data berhasil", true)
}

func (briefing BriefingController) DeleteBriefingMateri(c *gin.Context) {
	data := models.BriefMateriRequest{}

	if err := c.Bind(&data); err != nil {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := briefing.service.DeleteBriefingMateri(&data)
	if err != nil {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if !status {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data Gagal Disimpan", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Delete data berhasil", true)
}

func (briefing BriefingController) UpdateAllBrief(c *gin.Context) {
	data := models.BriefingResponseMaintain{}

	if err := c.Bind(&data); err != nil {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	fmt.Println(data)

	status, message, err := briefing.service.UpdateAllBrief(&data)
	if err != nil {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", message, "")
		return
	}

	if !status {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", message, false)
		return
	}
	lib.ReturnToJson(c, 200, "200", message, true)
}

func (briefing BriefingController) UpdateDraft(c *gin.Context) {
	data := models.BriefingResponseMaintain{}

	if err := c.Bind(&data); err != nil {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	fmt.Println(data)

	status, err := briefing.service.UpdateDraft(&data)
	if err != nil {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if !status {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data Gagal diupdate", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Update data berhasil", true)
}

func (briefing BriefingController) FilterBriefing(c *gin.Context) {
	requests := models.BriefingFilterRequest{}

	if err := c.Bind(&requests); err != nil {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := briefing.service.FilterBriefing(requests)
	if err != nil {
		briefing.logger.Zap.Error(err)
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
	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery data berhasil", datas, pagination)
}

func (briefing BriefingController) GetDataWithPagination(c *gin.Context) {
	requests := models.BriefingPagination{}

	if err := c.Bind(&requests); err != nil {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := briefing.service.GetDataWithPagination(requests)
	if err != nil {
		briefing.logger.Zap.Error(err)
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
	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery data berhasil", datas, pagination)
}

func (briefing BriefingController) GetNoPelaporan(c *gin.Context) {
	request := models.NoPelaporanRequest{}
	today := lib.GetTimeNow("date2")

	if err := c.Bind(&request); err != nil {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	datas, err := briefing.service.GetNoPelaporan(request)

	if err != nil {
		briefing.logger.Zap.Error(err)
	}

	if len(datas) == 0 {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data Tidak Ditemukan", nil)
		return
	}

	counter := datas[0].NoPelaporan
	fmt.Println("ORGEH", datas[0].ORGEH)
	fmt.Println("DATE", today)
	fmt.Println("Counter", counter)

	NoPelaporan := "BR-" + datas[0].ORGEH + "-" + today + "-" + counter

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", NoPelaporan)
}

func (briefing BriefingController) BriefingReportFilter(c *gin.Context) {
	request := models.BriefingFilterReport{}

	if err := c.Bind(&request); err != nil {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := briefing.service.BriefingReportFilter(request)
	if err != nil {
		briefing.logger.Zap.Error(err)
	}

	// if pagination.Total == 0 {
	// 	lib.ReturnToJson(c, 200, "404", "Data Kosong", datas)
	// 	return
	// }

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	fmt.Println("Filter Data =>", datas)
	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", datas, pagination)
}

func (briefing BriefingController) DeleteMapPeserta(c *gin.Context) {
	data := models.BriefingMapPeserta{}

	if err := c.Bind(&data); err != nil {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := briefing.service.DeleteMapPeserta(&data)
	if err != nil {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if !status {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error status", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Delete data berhasil", true)
}

func (briefing BriefingController) BriefingReportFinalFilter(c *gin.Context) {
	request := models.BriefingFilterReport{}

	if err := c.Bind(&request); err != nil {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := briefing.service.BriefingReportFilterComplete(request)
	if err != nil {
		briefing.logger.Zap.Error(err)
	}

	// if pagination.Total == 0 {
	// 	lib.ReturnToJson(c, 200, "404", "Data Kosong", datas)
	// 	return
	// }

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	fmt.Println("Filter Data =>", datas)
	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", datas, pagination)
}

func (briefing BriefingController) BriefingReportDetail(c *gin.Context) {
	request := models.BriefingReportDetailRequest{}

	if err := c.Bind(&request); err != nil {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	datas, err := briefing.service.BriefingReportDetail(request)
	if err != nil {
		briefing.logger.Zap.Error(err)
	}

	// if pagination.Total == 0 {
	// 	lib.ReturnToJson(c, 200, "404", "Data Kosong", datas)
	// 	return
	// }

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	fmt.Println("Filter Data =>", datas)
	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}

func (briefing BriefingController) BriefingReportMateriList(c *gin.Context) {
	request := models.BriefingReportMateriRequest{}

	if err := c.Bind(&request); err != nil {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	datas, err := briefing.service.BriefingReportMateriList(request)
	if err != nil {
		briefing.logger.Zap.Error(err)
	}

	// if pagination.Total == 0 {
	// 	lib.ReturnToJson(c, 200, "404", "Data Kosong", datas)
	// 	return
	// }

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	fmt.Println("Filter Data =>", datas)
	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}

func (briefing BriefingController) BriefingReportByUkerFilter(c *gin.Context) {
	request := models.BriefingFilterReportByUker{}

	if err := c.Bind(&request); err != nil {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := briefing.service.BriefingReportByUkerFilter(request)
	if err != nil {
		briefing.logger.Zap.Error(err)
	}

	// if pagination.Total == 0 {
	// 	lib.ReturnToJson(c, 200, "404", "Data Kosong", datas)
	// 	return
	// }

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	fmt.Println("Filter Data =>", datas)
	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", datas, pagination)
}

func (briefing BriefingController) BriefingReportFilterByUkerComplete(c *gin.Context) {
	request := models.BriefingFilterReportByUker{}

	if err := c.Bind(&request); err != nil {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := briefing.service.BriefingReportFilterByUkerComplete(request)
	if err != nil {
		briefing.logger.Zap.Error(err)
	}

	// if pagination.Total == 0 {
	// 	lib.ReturnToJson(c, 200, "404", "Data Kosong", datas)
	// 	return
	// }

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	fmt.Println("Filter Data =>", datas)
	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", datas, pagination)
}

func (briefing BriefingController) BriefingReportList(c *gin.Context) {
	request := models.BriefingReportListRequest{}

	if err := c.Bind(&request); err != nil {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := briefing.service.BriefingReportList(request)
	if err != nil {
		briefing.logger.Zap.Error(err)
	}

	// if pagination.Total == 0 {
	// 	lib.ReturnToJson(c, 200, "404", "Data Kosong", datas)
	// 	return
	// }

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	// fmt.Println("Filter Data =>", datas)
	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", datas, pagination)
}

func (briefing BriefingController) BriefingFrekuensiRpt(c *gin.Context) {
	request := models.FrekuensiBriefingRequest{}

	if err := c.Bind(&request); err != nil {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), nil)
		return
	}

	datas, totalRows, err := briefing.service.BriefingFrekuensiRpt(request)
	if err != nil {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if len(datas) < 1 {
		lib.ReturnToJson(c, 200, "404", "Data tidak ditemukan", nil)
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquiry Data Berhasil", datas, totalRows)

}
