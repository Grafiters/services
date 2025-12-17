package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"riskmanagement/dto"
	"riskmanagement/lib"
	models "riskmanagement/models/riskindicator"
	services "riskmanagement/services/riskindicator"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type RiskIndicatorController struct {
	logger  logger.Logger
	service services.RiskIndicatorDefinition
}

func NewRiskIndicatorController(
	RiskIndicatorService services.RiskIndicatorDefinition,
	logger logger.Logger,
) RiskIndicatorController {
	return RiskIndicatorController{
		service: RiskIndicatorService,
		logger:  logger,
	}
}

func (riskIndicator RiskIndicatorController) GetAll(c *gin.Context) {
	datas, err := riskIndicator.service.GetAll()
	if err != nil {
		riskIndicator.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "200", "Internal Error", "")
		return
	}

	if len(datas) == 0 {
		riskIndicator.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data tidak ditemukan", nil)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas)
}

func (riskIndicator RiskIndicatorController) GetAllWithPaginate(c *gin.Context) {
	requests := models.Paginate{}

	if err := c.Bind(&requests); err != nil {
		riskIndicator.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := riskIndicator.service.GetAllWithPaginate(requests)
	if err != nil {
		riskIndicator.logger.Zap.Error(err)
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

func (riskIndicator RiskIndicatorController) Store(c *gin.Context) {
	requests := models.RiskIndicatorRequest{}

	if err := c.Bind(&requests); err != nil {
		riskIndicator.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", lib.InvalidBody, requests)
		return
	}

	status, err := riskIndicator.service.Store(requests)
	if err != nil {
		riskIndicator.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	// fmt.Println("status ===>", status)

	if !status {
		riskIndicator.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Kode '"+requests.RiskIndicatorCode+"' sudah ada, silahkan masukkan kode lain ", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Input data berhasil", true)
}

func (riskIndicator RiskIndicatorController) GetKode(c *gin.Context) {
	datas, err := riskIndicator.service.GetKode()
	if err != nil {
		riskIndicator.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if len(datas) == 0 {
		riskIndicator.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data Tidak Ditemukan !", nil)
		return
	}

	Kode := "RCL.MOP." + lib.GetTimeNow("date2") + "." + datas[0].Kode

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", Kode)
}

func (riskIndicator RiskIndicatorController) Delete(c *gin.Context) {
	data := models.UpdateDelete{}

	if err := c.Bind(&data); err != nil {
		riskIndicator.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := riskIndicator.service.Delete(&data)
	if err != nil {
		riskIndicator.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if !status {
		riskIndicator.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data Gagal Dihapus", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Hapus data berhasil", true)
}

func (riskIndicator RiskIndicatorController) GetOne(c *gin.Context) {
	requests := models.RiskIndicatorRequest{}

	if err := c.Bind(&requests); err != nil {
		riskIndicator.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, status, err := riskIndicator.service.GetOne(requests.ID)
	if err != nil {
		riskIndicator.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if !status {
		riskIndicator.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data Tidak Ditemukan", nil)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inqueri data berhasil", data)
}

func (riskIndicator RiskIndicatorController) Update(c *gin.Context) {
	data := models.RiskIndicatorRequest{}

	if err := c.Bind(&data); err != nil {
		riskIndicator.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := riskIndicator.service.Update(&data)
	if err != nil {
		riskIndicator.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error : ", "")
		return
	}

	if !status {
		riskIndicator.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data Gagal Diupdate : ", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Update data berhasil", true)
}

func (riskIndicator RiskIndicatorController) DeleteFilesByID(c *gin.Context) {
	requests := models.RiskIndicatorRequest{}

	if err := c.Bind(&requests); err != nil {
		riskIndicator.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := riskIndicator.service.DeleteFilesByID(requests.ID)
	if err != nil {
		riskIndicator.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Hapus data berhasil", data)
}

func (riskIndicator RiskIndicatorController) SearchRiskIndicatorByIssue(c *gin.Context) {
	requests := models.SearchRequest{}

	if err := c.Bind(&requests); err != nil {
		riskIndicator.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := riskIndicator.service.SearchRiskIndicatorByIssue(requests)
	if err != nil {
		riskIndicator.logger.Zap.Error(err)
	}

	if pagination.Total == 0 {
		lib.ReturnToJson(c, 200, "404", "Data Tidak Ditemukan", nil)
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", datas, pagination)
}

func (riskIndicator RiskIndicatorController) GetRekomendasiMateri(c *gin.Context) {
	requests := models.RiskIndicator{}

	if err := c.Bind(&requests); err != nil {
		riskIndicator.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	data, err := riskIndicator.service.GetRekomendasiMateri(int64(requests.ID))

	if err != nil {
		riskIndicator.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	if len(data) == 0 {
		// riskIndicator.logger.Zap.Error(data)
		lib.ReturnToJson(c, 200, "404", "Data tidak ditemukan", nil)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (riskIndicator RiskIndicatorController) SearchRiskIndicatorBySource(c *gin.Context) {
	request := models.KeyRiskBySourceRequest{}
	if err := c.Bind(&request); err != nil {
		riskIndicator.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}
	ri, pagination, err := riskIndicator.service.SearchRiskIndicatorBySource(request)
	if err != nil {
		riskIndicator.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}
	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", ri, pagination)

}

func (riskIndicator RiskIndicatorController) SearchRiskIndicatorKRID(c *gin.Context) {
	requests := models.KeyRiskRequest{}

	if err := c.Bind(&requests); err != nil {
		riskIndicator.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := riskIndicator.service.SearchRiskIndicatorKRID(requests)
	if err != nil {
		riskIndicator.logger.Zap.Error(err)
	}

	if pagination.Total == 0 {
		lib.ReturnToJson(c, 200, "404", "Data Tidak Ditemukan", nil)
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", datas, pagination)
}

func (riskIndicator RiskIndicatorController) FilterRiskIndicator(c *gin.Context) {
	requests := models.FilterRequest{}

	if err := c.Bind(&requests); err != nil {
		riskIndicator.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := riskIndicator.service.FilterRiskIndicator(requests)
	if err != nil {
		riskIndicator.logger.Zap.Error(err)
	}

	if pagination.Total == 0 {
		lib.ReturnToJson(c, 200, "404", "Data Kosong", datas)
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery data berhasil", datas, pagination)
}

func (ri RiskIndicatorController) SaveThreshold(c *gin.Context) {
	request := models.RiskIndicatorRequest{}

	if err := c.Bind(&request); err != nil {
		ri.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := ri.service.SaveThreshold(request)
	if err != nil {
		ri.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if !status {
		ri.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error status", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Threshold Berhasil Disimpan", true)
}

func (ri RiskIndicatorController) GetThreshold(c *gin.Context) {
	request := models.RiskIndicatorRequest{}

	if err := c.Bind(&request); err != nil {
		ri.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := ri.service.GetMappingThrehold(request.ID)
	if err != nil {
		ri.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}

func (ri RiskIndicatorController) GetMapRiskIssue(c *gin.Context) {
	request := models.RiskIndicatorRequest{}

	if err := c.Bind(&request); err != nil {
		ri.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := ri.service.GetMappingRiskIssue(request.ID)
	if err != nil {
		ri.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}

func (ri RiskIndicatorController) GetIndicatorByAktivityProduct(c *gin.Context) {
	request := models.IndicatorRequest{}

	if err := c.Bind(&request); err != nil {
		ri.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := ri.service.GetIndicatorByAktivityProduct(request)
	if err != nil {
		ri.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}

func (ri RiskIndicatorController) SearchRiskIndicatorTematik(c *gin.Context) {
	request := models.SearchRequest{}

	if err := c.Bind(&request); err != nil {
		ri.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := ri.service.SearchRiskIndicatorTematik(request)
	if err != nil {
		ri.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}

func (ri RiskIndicatorController) GetTematikData(c *gin.Context) {
	fmt.Println("masuk controller")
	request := models.TematikDataRequest{}

	if err := c.Bind(&request); err != nil {
		ri.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := ri.service.GetTematikData(request)
	if err != nil {
		ri.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	// var result interface{} // Define a variable to hold the decoded JSON data
	var raw json.RawMessage

	if err := json.Unmarshal(data, &raw); err != nil {
		ri.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Error decoding JSON data", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", raw)
}

func (ri RiskIndicatorController) GetMateriIfFinish(c *gin.Context) {
	request := models.RequestMateriIfFinish{}

	if err := c.Bind(&request); err != nil {
		ri.logger.Zap.Error("input tidak sesuai")
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	document, err := ri.service.GetMateriIfFinish(request)
	if err != nil {
		ri.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Berhasil", document)
}

func (ri RiskIndicatorController) UpdateStatus(c *gin.Context) {
	requests := models.RiskIndicatorRequest{}
	if err := c.Bind(&requests); err != nil {
		ri.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := ri.service.UpdateStatus(requests.ID)
	if err != nil {
		ri.logger.Zap.Error(err)
		lib.ReturnToJson(c, http.StatusUnprocessableEntity, strconv.Itoa(http.StatusUnprocessableEntity), err.Error(), nil)
		return
	}

	statusLabel := "Active"
	if status {
		statusLabel = "Inactive"
	}

	msg := fmt.Sprintf("Risk Indicator %s", statusLabel)

	lib.ReturnToJson(c, http.StatusOK, strconv.Itoa(http.StatusOK), msg, nil)
}

func (ri RiskIndicatorController) Template(c *gin.Context) {
	blob, fileName, err := ri.service.Template()
	if err != nil {
		ri.logger.Zap.Error(err)
		lib.ReturnToJson(c, http.StatusInternalServerError, strconv.Itoa(http.StatusInternalServerError), "Generate Template error : "+err.Error(), "")
		return
	}

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", blob)
}

func (ri RiskIndicatorController) Preview(c *gin.Context) {
	pernr := c.PostForm("pernr")
	if pernr == "" {
		lib.ReturnToJson(c, http.StatusUnprocessableEntity, strconv.Itoa(http.StatusUnprocessableEntity), lib.InvalidBody, nil)
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		ri.logger.Zap.Error(err)
		lib.ReturnToJson(c, http.StatusUnprocessableEntity, strconv.Itoa(http.StatusUnprocessableEntity), lib.InvalidBody, nil)
		return
	}
	src, err := lib.ExtractFile(file)
	if err != nil {
		ri.logger.Zap.Error(err)
		lib.ReturnToJson(c, http.StatusUnprocessableEntity, strconv.Itoa(http.StatusUnprocessableEntity), err.Error(), nil)
		return
	}

	ext := filepath.Ext(file.Filename)
	if ext != ".csv" && ext != ".xls" && ext != ".xlsx" {
		ri.logger.Zap.Error(err)
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
		ri.logger.Zap.Debug(extract)
		if err != nil {
			ri.logger.Zap.Error("Error parsing file excel: %s", err)
			lib.ReturnToJson(c, http.StatusUnprocessableEntity, strconv.Itoa(http.StatusUnprocessableEntity), lib.InternalError, nil)
			return
		}
	case ".csv":
		extract, err = lib.ParseCSVFile(src)
		if err != nil {
			ri.logger.Zap.Error("Error parsing file csv: %s", err)
			lib.ReturnToJson(c, http.StatusUnprocessableEntity, strconv.Itoa(http.StatusUnprocessableEntity), lib.InternalError, nil)
			return
		}
	}

	data, err := ri.service.Preview(pernr, extract)
	if err != nil {
		ri.logger.Zap.Error("Error validate data: %s ", err)
		lib.ReturnToJson(c, http.StatusUnprocessableEntity, strconv.Itoa(http.StatusUnprocessableEntity), err.Error(), nil)
		return
	}

	lib.ReturnToJson(c, http.StatusOK, strconv.Itoa(http.StatusOK), lib.SuccessGetMessage, data)
}

func (ri RiskIndicatorController) ImportData(c *gin.Context) {
	pernr := c.PostForm("pernr")
	if pernr == "" {
		lib.ReturnToJson(c, http.StatusUnprocessableEntity, strconv.Itoa(http.StatusUnprocessableEntity), lib.InvalidBody, nil)
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		ri.logger.Zap.Error(err)
		lib.ReturnToJson(c, http.StatusUnprocessableEntity, strconv.Itoa(http.StatusUnprocessableEntity), lib.InvalidBody, nil)
		return
	}
	src, err := lib.ExtractFile(file)
	if err != nil {
		ri.logger.Zap.Error(err)
		lib.ReturnToJson(c, http.StatusUnprocessableEntity, strconv.Itoa(http.StatusUnprocessableEntity), err.Error(), nil)
		return
	}

	ext := filepath.Ext(file.Filename)
	if ext != ".csv" && ext != ".xls" && ext != ".xlsx" {
		ri.logger.Zap.Error(err)
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
		ri.logger.Zap.Debug(extract)
		if err != nil {
			ri.logger.Zap.Error("Error parsing file excel: %s", err)
			lib.ReturnToJson(c, http.StatusUnprocessableEntity, strconv.Itoa(http.StatusUnprocessableEntity), lib.InternalError, nil)
			return
		}
	case ".csv":
		extract, err = lib.ParseCSVFile(src)
		if err != nil {
			ri.logger.Zap.Error("Error parsing file csv: %s", err)
			lib.ReturnToJson(c, http.StatusUnprocessableEntity, strconv.Itoa(http.StatusUnprocessableEntity), lib.InternalError, nil)
			return
		}
	}

	err = ri.service.ImportData(pernr, extract)
	if err != nil {
		ri.logger.Zap.Error("Error validate data: %s ", err)
		lib.ReturnToJson(c, http.StatusUnprocessableEntity, strconv.Itoa(http.StatusUnprocessableEntity), err.Error(), nil)
		return
	}

	lib.ReturnToJson(c, http.StatusOK, strconv.Itoa(http.StatusOK), "Success Import Data", nil)
}

func (ri RiskIndicatorController) Download(c *gin.Context) {
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

	fileByte, fileName, err := ri.service.Download(pernr.Pernr, format)
	if err != nil {
		ri.logger.Zap.Error("Error generate file: %s", err)
		lib.ReturnToJson(c, http.StatusInternalServerError, strconv.Itoa(http.StatusInternalServerError), lib.InternalError, nil)
	}

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
