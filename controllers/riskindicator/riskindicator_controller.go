package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/riskindicator"
	services "riskmanagement/services/riskindicator"

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
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai :"+err.Error(), requests)
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
		// riskIndicator.logger.Zap.Error(err)
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
