package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"path/filepath"
	"riskmanagement/dto"
	"riskmanagement/lib"
	models "riskmanagement/models/riskcontrol"
	services "riskmanagement/services/riskcontrol"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type RiskControlController struct {
	logger  logger.Logger
	service services.RiskControlDefinition
}

func NewRiskControlController(
	RiskControlService services.RiskControlDefinition,
	logger logger.Logger,
) RiskControlController {
	return RiskControlController{
		logger:  logger,
		service: RiskControlService,
	}
}

func (riskControl RiskControlController) GetAll(c *gin.Context) {
	datas, err := riskControl.service.GetAll()
	if err != nil {
		riskControl.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas)
}

func (rc RiskControlController) GetAllWithPaginate(c *gin.Context) {
	requests := models.Paginate{}

	if err := c.Bind(&requests); err != nil {
		rc.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := rc.service.GetAllWithPaginate(requests)
	if err != nil {
		rc.logger.Zap.Error(err)
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

func (riskControl RiskControlController) GetOne(c *gin.Context) {
	requests := models.RiskControlRequest{}

	if err := c.Bind(&requests); err != nil {
		riskControl.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := riskControl.service.GetOne(requests.ID)
	if err != nil {
		riskControl.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}
func (riskControl RiskControlController) Store(c *gin.Context) {
	data := models.RiskControlRequest{}

	if err := c.Bind(&data); err != nil {
		riskControl.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	if err := riskControl.service.Store(&data); err != nil {
		riskControl.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "200", err.Error(), data)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Input data berhasil", data)
}
func (riskControl RiskControlController) Update(c *gin.Context) {
	data := models.RiskControlRequest{}

	if err := c.Bind(&data); err != nil {
		riskControl.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	if err := riskControl.service.Update(&data); err != nil {
		riskControl.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "200", "Internal error", data)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Update data berhasil", data)
}
func (riskControl RiskControlController) UpdateStatus(c *gin.Context) {
	requests := models.RiskControlRequest{}
	if err := c.Bind(&requests); err != nil {
		riskControl.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	if err := riskControl.service.UpdateStatus(requests.ID); err != nil {
		riskControl.logger.Zap.Error(err)
		lib.ReturnToJson(c, http.StatusUnprocessableEntity, strconv.Itoa(http.StatusUnprocessableEntity), err.Error(), nil)
		return
	}

	lib.ReturnToJson(c, http.StatusOK, strconv.Itoa(http.StatusOK), "Update data berhasil", nil)
}

func (riskControl RiskControlController) Delete(c *gin.Context) {
	requests := models.RiskControlRequest{}

	if err := c.Bind(&requests); err != nil {
		riskControl.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	if err := riskControl.service.Delete(requests.ID); err != nil {
		riskControl.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Data berhasil dihapus", true)
}
func (riskControl RiskControlController) GetKodeRiskControl(c *gin.Context) {
	datas, err := riskControl.service.GetKodeRiskControl()
	if err != nil {
		riskControl.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if len(datas) == 0 {
		riskControl.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data Tidak Ditemukan !", nil)
		return
	}

	KodeRiskControl := "C" + datas[0].KodeRiskControl

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", KodeRiskControl)
}

func (riskControl RiskControlController) GenCode(c *gin.Context) {
	code, err := riskControl.service.GenCode()
	if err != nil {
		riskControl.logger.Zap.Error(err)
		lib.ReturnToJson(c, http.StatusInternalServerError, strconv.Itoa(http.StatusInternalServerError), lib.InternalError, nil)
		return
	}

	response := dto.ResponseCode{
		Code: code,
	}

	lib.ReturnToJson(c, http.StatusOK, strconv.Itoa(http.StatusOK), lib.SuccessGetMessage, response)
}

func (riskControl RiskControlController) SearchRiskIndicatorByIssue(c *gin.Context) {
	requests := models.KeywordRequest{}

	if err := c.Bind(&requests); err != nil {
		riskControl.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := riskControl.service.SearchRiskControlByIssue(requests)
	if err != nil {
		riskControl.logger.Zap.Error(err)
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

func (riskControl RiskControlController) Template(c *gin.Context) {
	fileBytes, fileName, err := riskControl.service.Template()
	if err != nil {
		riskControl.logger.Zap.Error(err)
		lib.ReturnToJson(c, http.StatusInternalServerError, strconv.Itoa(http.StatusInternalServerError), "Generate Template error : "+err.Error(), "")
		return
	}

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", fileBytes)
}

func (riskControl RiskControlController) PreviewImport(c *gin.Context) {
	pernr := c.PostForm("pernr")
	if pernr == "" {
		lib.ReturnToJson(c, http.StatusUnprocessableEntity, strconv.Itoa(http.StatusUnprocessableEntity), lib.InvalidBody, nil)
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		riskControl.logger.Zap.Error(err)
		lib.ReturnToJson(c, http.StatusUnprocessableEntity, strconv.Itoa(http.StatusUnprocessableEntity), lib.InvalidBody, nil)
		return
	}
	src, err := lib.ExtractFile(file)
	if err != nil {
		riskControl.logger.Zap.Error(err)
		lib.ReturnToJson(c, http.StatusUnprocessableEntity, strconv.Itoa(http.StatusUnprocessableEntity), err.Error(), nil)
		return
	}

	ext := filepath.Ext(file.Filename)
	if ext != ".csv" && ext != ".xls" && ext != ".xlsx" {
		riskControl.logger.Zap.Error(err)
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
		riskControl.logger.Zap.Debug(extract)
		if err != nil {
			riskControl.logger.Zap.Error("Error parsing file excel: %s", err)
			lib.ReturnToJson(c, http.StatusUnprocessableEntity, strconv.Itoa(http.StatusUnprocessableEntity), lib.InternalError, nil)
			return
		}
	case ".csv":
		extract, err = lib.ParseCSVFile(src)
		if err != nil {
			riskControl.logger.Zap.Error("Error parsing file csv: %s", err)
			lib.ReturnToJson(c, http.StatusUnprocessableEntity, strconv.Itoa(http.StatusUnprocessableEntity), lib.InternalError, nil)
			return
		}
	}

	data, err := riskControl.service.Preview(pernr, extract)
	if err != nil {
		riskControl.logger.Zap.Error("Error validate data: %s ", err)
		lib.ReturnToJson(c, http.StatusUnprocessableEntity, strconv.Itoa(http.StatusUnprocessableEntity), err.Error(), nil)
		return
	}

	lib.ReturnToJson(c, http.StatusOK, strconv.Itoa(http.StatusOK), lib.SuccessGetMessage, data)
}

func (riskControl RiskControlController) ImportData(c *gin.Context) {
	pernr := c.PostForm("pernr")
	if pernr == "" {
		lib.ReturnToJson(c, http.StatusUnprocessableEntity, strconv.Itoa(http.StatusUnprocessableEntity), lib.InvalidBody, nil)
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		riskControl.logger.Zap.Error(err)
		lib.ReturnToJson(c, http.StatusUnprocessableEntity, strconv.Itoa(http.StatusUnprocessableEntity), "Invalid Body pernr is missing", nil)
		return
	}
	src, err := lib.ExtractFile(file)
	if err != nil {
		riskControl.logger.Zap.Error(err)
		lib.ReturnToJson(c, http.StatusUnprocessableEntity, strconv.Itoa(http.StatusUnprocessableEntity), err.Error(), nil)
		return
	}

	ext := filepath.Ext(file.Filename)
	if ext != ".csv" && ext != ".xls" && ext != ".xlsx" {
		riskControl.logger.Zap.Error(err)
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
		riskControl.logger.Zap.Debug(extract)
		if err != nil {
			riskControl.logger.Zap.Error("Error parsing file excel: %s", err)
			lib.ReturnToJson(c, http.StatusUnprocessableEntity, strconv.Itoa(http.StatusUnprocessableEntity), lib.InternalError, nil)
			return
		}
	case ".csv":
		extract, err = lib.ParseCSVFile(src)
		if err != nil {
			riskControl.logger.Zap.Error("Error parsing file csv: %s", err)
			lib.ReturnToJson(c, http.StatusUnprocessableEntity, strconv.Itoa(http.StatusUnprocessableEntity), lib.InternalError, nil)
			return
		}
	}

	err = riskControl.service.ImportData(pernr, extract)
	if err != nil {
		riskControl.logger.Zap.Error("Error saving data: %s", err)
		lib.ReturnToJson(c, http.StatusUnprocessableEntity, strconv.Itoa(http.StatusUnprocessableEntity), lib.InternalError, nil)
		return
	}

	lib.ReturnToJson(c, http.StatusCreated, strconv.Itoa(http.StatusCreated), "Success Import Data", nil)
}

func (riskControl RiskControlController) Download(c *gin.Context) {
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

	fileByte, fileName, err := riskControl.service.Download(pernr.Pernr, format)
	if err != nil {
		riskControl.logger.Zap.Error("Error generate file: %s", err)
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
