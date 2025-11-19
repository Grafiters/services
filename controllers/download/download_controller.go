package controller

import (
	// "database/sql"
	"fmt"
	"io"
	"net/http"
	"riskmanagement/lib"
	models "riskmanagement/models/download"
	services "riskmanagement/services/download"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
	// "io/ioutil"
)

type DownloadController struct {
	logger  logger.Logger
	service services.DownloadDefinition
}

func NewDownloadController(
	DownloadService services.DownloadDefinition,
	logger logger.Logger,
) DownloadController {
	return DownloadController{
		service: DownloadService,
		logger:  logger,
	}
}

func (download DownloadController) Generate(c *gin.Context) {
	fmt.Println("masuk controller")
	data := models.DownloadRequest{}

	if err := c.Bind(&data); err != nil {
		download.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), data)
		return
	}

	datas, err := download.service.Generate(data)
	if err != nil {
		download.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}

func (download DownloadController) DownloadHandler(c *gin.Context) {
	fmt.Println("masuk controller")

	// get filename and filepath based on uid stored in db
	id := c.Param("id")
	url, filename, _ := download.service.DownloadHandler(id)
	fmt.Println("masuk controller - url", url)

	// Retrieve the file data from the URL = uid
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error retrieving file:", err)
	}
	defer resp.Body.Close()

	// Read the file data
	// fileData, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("Error reading file data:", err)
	// }

	// Set the appropriate headers for file download
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

	// Copy the file contents to the response writer
	_, err = io.Copy(c.Writer, resp.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
}

func (download DownloadController) GetListDownload(c *gin.Context) {
	data := models.ListDownloadRequest{}

	if err := c.Bind(&data); err != nil {
		download.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), data)
		return
	}

	datas, pagination, err := download.service.GetListDownload(data)
	if err != nil {
		download.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", datas, pagination)
}

func (download DownloadController) GetReportType(c *gin.Context) {
	fmt.Println("Masuk Controller")

	data, err := download.service.GetReportType()

	if err != nil {
		download.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if len(data) < 0 {
		// download.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data tidak ditemukan", nil)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}

func (download DownloadController) Retry(c *gin.Context) {
	fmt.Println("Masuk Retry Controller")
	data := models.RetryRequest{}

	if err := c.Bind(&data); err != nil {
		download.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), data)
		return
	}

	datas, err := download.service.Retry(data)
	if err != nil {
		download.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}

func (download DownloadController) FetchOneRows(c *gin.Context) {
	request := models.RetryRequest{}

	if err := c.Bind(&request); err != nil {
		download.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), request)
		return
	}

	datas, err := download.service.FetchOneRows(int64(request.InsertId))
	if err != nil {
		download.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}
