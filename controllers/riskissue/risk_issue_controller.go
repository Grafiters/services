package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"path/filepath"
	"riskmanagement/dto"
	"riskmanagement/lib"
	models "riskmanagement/models/riskissue"
	services "riskmanagement/services/riskissue"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type RiskIssueController struct {
	logger  logger.Logger
	service services.RiskIssueDefinition
}

func NewRiskIssueController(
	RiskIssueService services.RiskIssueDefinition,
	logger logger.Logger,
) RiskIssueController {
	return RiskIssueController{
		service: RiskIssueService,
		logger:  logger,
	}
}

func (riskIssue RiskIssueController) GetAll(c *gin.Context) {
	datas, err := riskIssue.service.GetAll()

	if err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "200", "Internal Error", "")
		return
	}

	if len(datas) == 0 {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data tidak ditemukan", "")
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas)
}
func (riskIssue RiskIssueController) GetOne(c *gin.Context) {
	requests := models.RiskIssueRequest{}

	if err := c.Bind(&requests); err != nil {
		riskIssue.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, status, err := riskIssue.service.GetOne(requests.ID)
	if err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	if !status {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data tidak ditemukan", nil)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}

func (riskIssue RiskIssueController) Store(c *gin.Context) {
	data := models.RiskIssueRequest{}

	if err := c.Bind(&data); err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai :"+err.Error(), "")
		return
	}

	status, err := riskIssue.service.Store(data)
	if err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if !status {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error status", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Input data berhasil", true)
}

func (riskIssue RiskIssueController) Update(c *gin.Context) {
	data := models.RiskIssueRequest{}

	if err := c.Bind(&data); err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := riskIssue.service.Update(&data)
	if err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if !status {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error status", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Update data berhasil", true)
}

func (riskIssue RiskIssueController) DeleteMapAktifitas(c *gin.Context) {
	data := models.MapAktifitas{}

	if err := c.Bind(&data); err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := riskIssue.service.DeleteMapAktifitas(&data)
	if err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if !status {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error status", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Delete data berhasil", true)
}

func (riskIssue RiskIssueController) DeleteMapEvent(c *gin.Context) {
	data := models.MapEvent{}

	if err := c.Bind(&data); err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := riskIssue.service.DeleteMapEvent(&data)
	if err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if !status {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error status", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Delete data berhasil", true)
}

func (riskIssue RiskIssueController) DeleteMapKejadian(c *gin.Context) {
	data := models.MapKejadian{}

	if err := c.Bind(&data); err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := riskIssue.service.DeleteMapKejadian(&data)
	if err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if !status {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error status", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Delete data berhasil", true)
}

func (riskIssue RiskIssueController) DeleteMapLiniBisnis(c *gin.Context) {
	data := models.MapLiniBisnis{}

	if err := c.Bind(&data); err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := riskIssue.service.DeleteMapLiniBisnis(&data)
	if err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if !status {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error status", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Delete data berhasil", true)
}

func (riskIssue RiskIssueController) DeleteMapProduct(c *gin.Context) {
	data := models.MapProduct{}

	if err := c.Bind(&data); err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := riskIssue.service.DeleteMapProduct(&data)
	if err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if !status {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error status", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Delete data berhasil", true)
}

func (riskIssue RiskIssueController) DeleteMapProses(c *gin.Context) {
	data := models.MapProses{}

	if err := c.Bind(&data); err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := riskIssue.service.DeleteMapProses(&data)
	if err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if !status {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error status", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Delete data berhasil", true)
}

func (riskIssue RiskIssueController) GetKode(c *gin.Context) {
	datas, err := riskIssue.service.GetKode()
	if err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if len(datas) == 0 {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data Tidak Ditemukan !", nil)
		return
	}

	Kode := "RE." + lib.GetTimeNow("date2") + "." + datas[0].Kode

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", Kode)
}

func (riskIssue RiskIssueController) MappingRiskControl(c *gin.Context) {
	data := models.MappingControlRequest{}

	if err := c.Bind(&data); err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := riskIssue.service.MappingRiskControl(data)
	if err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if !status {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error status", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Mapping data berhasil", true)
}

func (riskIssue RiskIssueController) GetMappingControlbyID(c *gin.Context) {
	requests := models.RiskIssueRequest{}

	if err := c.Bind(&requests); err != nil {
		riskIssue.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := riskIssue.service.GetMappingControlbyID(requests.ID)
	if err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}

func (riskIssue RiskIssueController) DeleteMapControl(c *gin.Context) {
	data := models.MapControl{}

	if err := c.Bind(&data); err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := riskIssue.service.DeleteMapControl(&data)
	if err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if !status {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error status", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Delete data berhasil", true)
}

func (riskIssue RiskIssueController) ListRiskIssue(c *gin.Context) {
	requests := models.ListRiskIssueRequest{}

	if err := c.Bind(&requests); err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}
	data, err := riskIssue.service.ListRiskIssue(requests)
	if err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}

func (riskIssue RiskIssueController) SearchRiskIssue(c *gin.Context) {
	requests := models.KeywordRequest{}

	if err := c.Bind(&requests); err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := riskIssue.service.SearchRiskIssue(requests)
	if err != nil {
		riskIssue.logger.Zap.Error(err)
	}

	if pagination.Total == 0 {
		lib.ReturnToJson(c, 200, "404", "Data Tidak Ditemukan", datas)
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", datas, pagination)
}

func (riskIssue RiskIssueController) SearchRiskIssueWithoutSub(c *gin.Context) {
	requests := models.RiskIssueWithoutSub{}

	if err := c.Bind(&requests); err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	fmt.Println("request Controller", requests)
	datas, pagination, err := riskIssue.service.SearchRiskIssueWithoutSub(requests)

	if err != nil {
		riskIssue.logger.Zap.Error(err)
	}

	// if pagination.Total == 0 {
	// 	// lib.ReturnToJson(c, 200, "404", "Data Tidak Ditemukan", datas)
	// 	lib.ReturnToJson(c, 200, "404", "Data Tidak Ditemukan", nil)
	// 	return
	// }

	// if err == sql.ErrNoRows {
	// 	lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	// 	return
	// }

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", datas, pagination)
}

func (riskIssue RiskIssueController) MappingRiskIndicator(c *gin.Context) {
	data := models.MappingIndicatorRequest{}

	if err := c.Bind(&data); err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := riskIssue.service.MappingRiskIndicator(data)
	if err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if !status {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error status", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Mapping data berhasil", true)
}

func (riskIssue RiskIssueController) GetMappingIndicatorbyID(c *gin.Context) {
	requests := models.RiskIssueRequest{}

	if err := c.Bind(&requests); err != nil {
		riskIssue.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := riskIssue.service.GetMappingIndicatorbyID(requests.ID)
	if err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}

func (riskIssue RiskIssueController) DeleteMapIndicator(c *gin.Context) {
	data := models.MapIndicator{}

	if err := c.Bind(&data); err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := riskIssue.service.DeleteMapIndicator(&data)
	if err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if !status {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error status", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Delete data berhasil", true)
}

func (riskIssue RiskIssueController) Delete(c *gin.Context) {
	data := models.RiskIssueDeleteRequest{}

	if err := c.Bind(&data); err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := riskIssue.service.Delete(&data)
	if err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if !status {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data Gagal Dihapus", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Hapus data berhasil", true)
}

func (riskIssue RiskIssueController) FilterRiskIssue(c *gin.Context) {
	requests := models.FilterRiskIssueRequest{}

	if err := c.Bind(&requests); err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := riskIssue.service.FilterRiskIssue(requests)
	if err != nil {
		riskIssue.logger.Zap.Error(err)
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

func (riskIssue RiskIssueController) GetRiskIssueByActivity(c *gin.Context) {
	requests := models.RiskIssueRequest{}

	if err := c.Bind(&requests); err != nil {
		riskIssue.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := riskIssue.service.GetRiskIssueByActivity(requests.ID)

	if err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	if len(data) == 0 {
		riskIssue.logger.Zap.Error(data)
		lib.ReturnToJson(c, 200, "404", "Data tidak ditemukan", nil)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (riskIssue RiskIssueController) GetRekomendasiMateri(c *gin.Context) {
	requests := models.RiskIssueRequest{}

	if err := c.Bind(&requests); err != nil {
		riskIssue.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := riskIssue.service.GetRekomendasiMateri(requests.ID)

	if err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	if len(data) == 0 {
		// riskIssue.logger.Zap.Error(data)
		lib.ReturnToJson(c, 200, "404", "Data tidak ditemukan", nil)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (riskIssue RiskIssueController) GetMateriByCode(c *gin.Context) {
	requests := models.RiskIssueCode{}

	if err := c.Bind(&requests); err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	data, err := riskIssue.service.GetMateriByCode(requests)

	if err != nil {
		riskIssue.logger.Zap.Error(err)
	}

	if len(data) == 0 {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data tidak ditemukan", nil)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}

func (riskIssue RiskIssueController) GetAllWithPaginate(c *gin.Context) {
	requests := models.Paginate{}

	if err := c.Bind(&requests); err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := riskIssue.service.GetAllWithPaginate(requests)
	if err != nil {
		riskIssue.logger.Zap.Error(err)
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

func (riskIssue RiskIssueController) GetRiskIssueByActivityID(c *gin.Context) {
	requests := models.RiskIssueRequest{}

	if err := c.Bind(&requests); err != nil {
		riskIssue.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := riskIssue.service.GetRiskIssueByActivityID(requests.ID)

	if err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	if len(data) == 0 {
		riskIssue.logger.Zap.Error(data)
		lib.ReturnToJson(c, 200, "404", "Data tidak ditemukan", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (riskIssue RiskIssueController) UpdateStatus(c *gin.Context) {
	request := models.RiskIssueRequest{}
	if err := c.Bind(&request); err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	if err := riskIssue.service.UpdateStatus(request.ID); err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, http.StatusUnprocessableEntity, strconv.Itoa(http.StatusUnprocessableEntity), err.Error(), nil)
		return
	}

	lib.ReturnToJson(c, http.StatusOK, strconv.Itoa(http.StatusOK), "Update data berhasil", nil)
}

func (riskIssue RiskIssueController) Template(c *gin.Context) {
	fileBytes, fileName, err := riskIssue.service.Template()
	if err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, http.StatusInternalServerError, strconv.Itoa(http.StatusInternalServerError), "Generate Template error : "+err.Error(), "")
		return
	}

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", fileBytes)
}

func (riskIssue RiskIssueController) PreviewData(c *gin.Context) {
	pernr := c.PostForm("pernr")
	if pernr == "" {
		lib.ReturnToJson(c, http.StatusUnprocessableEntity, strconv.Itoa(http.StatusUnprocessableEntity), lib.InvalidBody, nil)
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, http.StatusUnprocessableEntity, strconv.Itoa(http.StatusUnprocessableEntity), lib.InvalidBody, nil)
		return
	}
	src, err := lib.ExtractFile(file)
	if err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, http.StatusUnprocessableEntity, strconv.Itoa(http.StatusUnprocessableEntity), err.Error(), nil)
		return
	}

	ext := filepath.Ext(file.Filename)
	if ext != ".csv" && ext != ".xls" && ext != ".xlsx" {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, http.StatusUnprocessableEntity, strconv.Itoa(http.StatusUnprocessableEntity), lib.InvalidFormatFile, nil)
		return
	}
	// Logic Proces
	var (
		extract [][]string
	)

	switch ext {
	case ".xlsx":
		extract, err = lib.ParseExcelFile(src)
		riskIssue.logger.Zap.Debug(extract)
		if err != nil {
			riskIssue.logger.Zap.Error("Error parsing file excel: %s", err)
			lib.ReturnToJson(c, http.StatusUnprocessableEntity, strconv.Itoa(http.StatusUnprocessableEntity), lib.InternalError, nil)
			return
		}
	case ".csv":
		extract, err = lib.ParseCSVFile(src)
		if err != nil {
			riskIssue.logger.Zap.Error("Error parsing file csv: %s", err)
			lib.ReturnToJson(c, http.StatusUnprocessableEntity, strconv.Itoa(http.StatusUnprocessableEntity), lib.InternalError, nil)
			return
		}
	}

	data, err := riskIssue.service.PreviewData(pernr, extract)
	riskIssue.logger.Zap.Debug("============================	")
	if err != nil {
		riskIssue.logger.Zap.Error("Error validate data: %s ", err)
		lib.ReturnToJson(c, http.StatusUnprocessableEntity, strconv.Itoa(http.StatusUnprocessableEntity), err.Error(), nil)
		return
	}

	lib.ReturnToJson(c, http.StatusOK, strconv.Itoa(http.StatusOK), lib.SuccessGetMessage, data)
}

func (riskIssue RiskIssueController) ImportData(c *gin.Context) {
	pernr := c.PostForm("pernr")
	if pernr == "" {
		lib.ReturnToJson(c, http.StatusUnprocessableEntity, strconv.Itoa(http.StatusUnprocessableEntity), lib.InvalidBody, lib.InvalidBody)
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, http.StatusUnprocessableEntity, strconv.Itoa(http.StatusUnprocessableEntity), lib.InvalidBody, nil)
		return
	}
	src, err := lib.ExtractFile(file)
	if err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, http.StatusUnprocessableEntity, strconv.Itoa(http.StatusUnprocessableEntity), err.Error(), nil)
		return
	}

	ext := filepath.Ext(file.Filename)
	if ext != ".csv" && ext != ".xls" && ext != ".xlsx" {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, http.StatusUnprocessableEntity, strconv.Itoa(http.StatusUnprocessableEntity), lib.InvalidFormatFile, nil)
		return
	}
	// Logic Proces
	var (
		extract [][]string
	)

	switch ext {
	case ".xlsx":
		extract, err = lib.ParseExcelFile(src)
		riskIssue.logger.Zap.Debug(extract)
		if err != nil {
			riskIssue.logger.Zap.Error("Error parsing file excel: %s", err)
			lib.ReturnToJson(c, http.StatusUnprocessableEntity, strconv.Itoa(http.StatusUnprocessableEntity), lib.InternalError, nil)
			return
		}
	case ".csv":
		extract, err = lib.ParseCSVFile(src)
		if err != nil {
			riskIssue.logger.Zap.Error("Error parsing file csv: %s", err)
			lib.ReturnToJson(c, http.StatusUnprocessableEntity, strconv.Itoa(http.StatusUnprocessableEntity), lib.InternalError, nil)
			return
		}
	}

	err = riskIssue.service.ImportData(pernr, extract)
	riskIssue.logger.Zap.Debug("============================	")
	if err != nil {
		riskIssue.logger.Zap.Error("Error validate data: %s ", err)
		lib.ReturnToJson(c, http.StatusUnprocessableEntity, strconv.Itoa(http.StatusUnprocessableEntity), err.Error(), nil)
		return
	}

	lib.ReturnToJson(c, http.StatusOK, strconv.Itoa(http.StatusOK), "Success Import Data", nil)
}

func (riskIssue RiskIssueController) Download(c *gin.Context) {
	var pernr dto.PernrRequest
	if err := c.ShouldBindJSON(&pernr); err != nil {
		lib.ReturnToJson(c, http.StatusBadRequest, strconv.Itoa(http.StatusBadRequest), "Invalid JSON", nil)
		return
	}
	format := c.Param("format")
	if format == "" {
		lib.ReturnToJson(c, http.StatusUnprocessableEntity, strconv.Itoa(http.StatusUnprocessableEntity), lib.InvalidParam, nil)
		return
	}

	fileByte, fileName, err := riskIssue.service.Download(pernr.Pernr, format)
	if err != nil {
		riskIssue.logger.Zap.Error("Error generate file: %s", err)
		lib.ReturnToJson(c, http.StatusInternalServerError, strconv.Itoa(http.StatusInternalServerError), lib.InternalError, nil)
	}

	riskIssue.logger.Zap.Debug(fileByte)
	riskIssue.logger.Zap.Debug(fileName)
	// Logic Proces
	var contentType string
	switch format {
	case "csv":
		contentType = "text/csv"
	case "xlsx":
		contentType = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	case "pdf":
		contentType = "application/pdf"
	default:
		contentType = "application/octet-stream"
	}

	c.Header("Content-Type", contentType)
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Cache-Control", "no-cache")

	c.Data(http.StatusOK, format, fileByte)
}

func (riskIssue RiskIssueController) GetRiskCategories(c *gin.Context) {
	requests := models.RiskIssueIDsRequest{}

	if err := c.Bind(&requests); err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	data, err := riskIssue.service.GetRiskCategories(requests.RiskIssueIDs)
	if err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}
