package controllers

import (
	"database/sql"
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/coaching"
	services "riskmanagement/services/coaching"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type CoachingController struct {
	logger  logger.Logger
	service services.CoachingDefinition
}

func NewCoachingController(
	CoachingService services.CoachingDefinition,
	logger logger.Logger,
) CoachingController {
	return CoachingController{
		service: CoachingService,
		logger:  logger,
	}
}

func (coaching CoachingController) GetAll(c *gin.Context) {
	datas, err := coaching.service.GetAll()

	if err != nil {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if len(datas) == 0 {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data tidak ditemukan", nil)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas)
}

func (coaching CoachingController) Store(c *gin.Context) {
	data := models.CoachingRequest{}

	if err := c.Bind(&data); err != nil {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), data)
		return
	}

	status, message, err := coaching.service.Store(data)
	if err != nil {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", message, err.Error())
		return
	}

	if !status {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", message, false)
		return
	}
	lib.ReturnToJson(c, 200, "200", message, true)
}

func (coaching CoachingController) StoreDraft(c *gin.Context) {
	data := models.CoachingRequest{}

	if err := c.Bind(&data); err != nil {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), data)
		return
	}

	status, err := coaching.service.StoreDraft(data)
	if err != nil {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if !status {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error status", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Input data berhasil", true)
}

func (coaching CoachingController) GetOne(c *gin.Context) {
	// paramID := c.Param("id")
	// id, err := strconv.Atoi(paramID)
	request := models.CoachingGetOneRequest{}

	if err := c.Bind(&request); err != nil {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	data, status, err := coaching.service.GetOne(int64(request.ID))
	if err != nil {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if !status {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data tidak ditemukan", nil)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}

func (coaching CoachingController) DeleteCoachingActivity(c *gin.Context) {
	data := models.CoachingActRequest{}

	if err := c.Bind(&data); err != nil {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := coaching.service.DeleteCoachingActivity(&data)
	if err != nil {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if !status {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data gagal disimpan", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Delete data berhasil", true)
}

func (coaching CoachingController) Delete(c *gin.Context) {
	data := models.CoachingRequestUpdate{}

	if err := c.Bind(&data); err != nil {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := coaching.service.Delete(&data)
	if err != nil {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if !status {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data Gagal Dihapus", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Hapus data berhasil", true)
}

func (coaching CoachingController) UpdateAllCoaching(c *gin.Context) {
	data := models.CoachingResponseMaintain{}

	if err := c.Bind(&data); err != nil {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	fmt.Println(data)

	status, message, err := coaching.service.UpdateAllCoaching(&data)
	if err != nil {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", message, "")
		return
	}

	if !status {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", message, false)
		return
	}

	lib.ReturnToJson(c, 200, "200", message, true)
}

func (coaching CoachingController) UpdateDraft(c *gin.Context) {
	data := models.CoachingResponseMaintain{}

	if err := c.Bind(&data); err != nil {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	fmt.Println(data)

	status, err := coaching.service.UpdateDraft(&data)
	if err != nil {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if !status {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data Gagal Diupdate", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Update data berhasil", true)
}

func (coaching CoachingController) FilterCoaching(c *gin.Context) {
	requests := models.CoachingFilterRequest{}

	if err := c.Bind(&requests); err != nil {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := coaching.service.FilterCoaching(requests)
	if err != nil {
		coaching.logger.Zap.Error(err)
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
	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery data berhasil", datas, pagination)
}

func (coaching CoachingController) GetDataWithPagination(c *gin.Context) {
	requests := models.CoachingPagination{}

	if err := c.Bind(&requests); err != nil {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := coaching.service.GetDataWithPagination(requests)
	if err != nil {
		coaching.logger.Zap.Error(err)
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
	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery data berhasil", datas, pagination)
}

func (coaching CoachingController) GetNoPelaporan(c *gin.Context) {
	requests := models.NoPalaporanRequest{}
	today := lib.GetTimeNow("date2")

	if err := c.Bind(&requests); err != nil {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	datas, err := coaching.service.GetNoPelaporan(requests)

	if err != nil {
		coaching.logger.Zap.Error(err)
	}

	if len(datas) == 0 {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data Tidak Ditemukan", nil)
		return
	}

	counter := datas[0].NoPelaporan
	fmt.Println("ORGEH", datas[0].ORGEH)
	fmt.Println("DATE", today)
	fmt.Println("Counter", counter)

	NoPelaporan := "CO-" + datas[0].ORGEH + "-" + today + "-" + counter

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", NoPelaporan)
}

func (coaching CoachingController) GetData(c *gin.Context) {
	datas, err := coaching.service.GetData()
	if err != nil {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if len(datas) == 0 {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data Tidak Ditemukan !", "")
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", nil)
}

func (coaching CoachingController) DeleteMapPeserta(c *gin.Context) {
	data := models.CoachingMapPeserta{}

	if err := c.Bind(&data); err != nil {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := coaching.service.DeleteMapPeserta(&data)
	if err != nil {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if !status {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error status", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Delete data berhasil", true)
}

func (coaching CoachingController) CoachingReportFilter(c *gin.Context) {
	request := models.CoachingFilterReportRequest{}

	if err := c.Bind(&request); err != nil {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := coaching.service.CoachingReportFilter(&request)
	if err != nil {
		coaching.logger.Zap.Error(err)
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

func (coaching CoachingController) CoachingReportByUkerFilter(c *gin.Context) {
	request := models.CoachingFilterReportByUker{}

	if err := c.Bind(&request); err != nil {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := coaching.service.CoachingReportByUkerFilter(request)
	if err != nil {
		coaching.logger.Zap.Error(err)
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

func (coaching CoachingController) CoachingReportFilterByUkerComplete(c *gin.Context) {
	request := models.CoachingFilterReportByUker{}

	if err := c.Bind(&request); err != nil {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := coaching.service.CoachingReportFilterByUkerComplete(request)
	if err != nil {
		coaching.logger.Zap.Error(err)
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

func (coaching CoachingController) CoachingReportFilterByUkerAllActivity(c *gin.Context) {
	request := models.CoachingFilterReportRequest{}

	if err := c.Bind(&request); err != nil {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	datas, _, _, err := coaching.service.CoachingReportFilterByUkerAllActivity(request)
	if err != nil {
		coaching.logger.Zap.Error(err)
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
	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}

func (coaching CoachingController) CoachingFinalReportFilter(c *gin.Context) {
	request := models.CoachingFilterReportRequest{}

	if err := c.Bind(&request); err != nil {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := coaching.service.CoachingFinalReportFilter(&request)
	if err != nil {
		coaching.logger.Zap.Error(err)
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

func (coaching CoachingController) CoachingReportDetail(c *gin.Context) {
	request := models.CoachingReportDetailRequest{}

	if err := c.Bind(&request); err != nil {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	datas, err := coaching.service.CoachingReportDetail(request)
	if err != nil {
		coaching.logger.Zap.Error(err)
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
	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}

func (coaching CoachingController) CoachingReportMateriList(c *gin.Context) {
	request := models.CoachingReportMateriRequest{}

	if err := c.Bind(&request); err != nil {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	datas, err := coaching.service.CoachingReportMateriList(request)
	if err != nil {
		coaching.logger.Zap.Error(err)
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
	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}

func (coaching CoachingController) CoachingReportList(c *gin.Context) {
	request := models.CoachingReportListRequest{}

	if err := c.Bind(&request); err != nil {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := coaching.service.CoachingReportList(request)
	if err != nil {
		coaching.logger.Zap.Error(err)
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

func (coaching CoachingController) CoachingFrekuensiRpt(c *gin.Context) {
	request := models.FrekuensiCoachingRequest{}

	if err := c.Bind(&request); err != nil {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), nil)
		return
	}

	datas, totalRows, err := coaching.service.CoachingFrekuensiRpt(request)
	if err != nil {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", nil)
		return
	}

	if len(datas) < 1 {
		lib.ReturnToJson(c, 200, "404", "Data tidak ditemukan", nil)
		return
	}
	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquiry Data Berhasil", datas, totalRows)
}
