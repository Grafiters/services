package controller

import (
	"database/sql"
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/msuker"
	services "riskmanagement/services/msuker"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type MsUkerController struct {
	logger  logger.Logger
	service services.MsUkerDefinition
}

func NewMsUkerController(
	msUkerService services.MsUkerDefinition,
	logger logger.Logger,
) MsUkerController {
	return MsUkerController{
		logger:  logger,
		service: msUkerService,
	}
}

func (msUker MsUkerController) GetAll(c *gin.Context) {
	datas, err := msUker.service.GetAll()

	if err != nil {
		msUker.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "200", "Internal Error", "")
		return
	}

	if len(datas) == 0 {
		msUker.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data tidak ditemukan", nil)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas)
}

func (msUker MsUkerController) GetUkerByBranch(c *gin.Context) {
	branchID := c.Param("branchid")
	id, err := strconv.Atoi(branchID)

	if err != nil {
		msUker.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai :"+err.Error(), "")
		return
	}

	data, err := msUker.service.GetUkerByBranch(int64(id))
	if err != nil {
		msUker.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}

func (msUker MsUkerController) GetUkerPerRegion(c *gin.Context) {
	requests := models.Region{}

	if err := c.Bind(&requests); err != nil {
		msUker.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai", "")
		return
	}

	data, err := msUker.service.GetUkerPerRegion(requests)
	if err != nil {
		msUker.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}

func (msUker MsUkerController) SearchUker(c *gin.Context) {
	requests := models.KeywordRequest{}

	if err := c.Bind(&requests); err != nil {
		msUker.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := msUker.service.SearchUker(requests)
	if err != nil {
		msUker.logger.Zap.Error(err)
	}

	if pagination.Total == 0 {
		lib.ReturnToJson(c, 200, "404", "Data Kosong", nil)
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}
	fmt.Println("Data =>", datas)
	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", datas, pagination)
}

func (msUker MsUkerController) SearchPeserta(c *gin.Context) {
	requests := models.KeywordRequest{}

	if err := c.Bind(&requests); err != nil {
		msUker.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := msUker.service.SearchPeserta(requests)
	if err != nil {
		msUker.logger.Zap.Error(err)
	}

	if pagination.Total == 0 {
		lib.ReturnToJson(c, 200, "404", "Data Kosong", nil)
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}
	// fmt.Println("Data =>", datas)
	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", datas, pagination)
}

func (msUker MsUkerController) GetPekerjaByBranch(c *gin.Context) {
	requests := models.BranchCodeInduk{}

	if err := c.Bind(&requests); err != nil {
		msUker.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai", "")
		return
	}

	data, err := msUker.service.GetPekerjaByBranch(requests)
	if err != nil {
		msUker.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}

func (msUker MsUkerController) GetPekerjaByRegion(c *gin.Context) {
	requests := models.SearchPNByRegionReq{}

	if err := c.Bind(&requests); err != nil {
		msUker.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai", "")
		return
	}

	data, err := msUker.service.GetPekerjaByRegion(&requests)
	if err != nil {
		msUker.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}

func (msUker MsUkerController) SearchJabatan(c *gin.Context) {
	requests := models.KeywordRequest{}

	if err := c.Bind(&requests); err != nil {
		msUker.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := msUker.service.SearchJabatan(requests)
	if err != nil {
		msUker.logger.Zap.Error(err)
	}

	if pagination.Total == 0 {
		lib.ReturnToJson(c, 200, "404", "Data Kosong", nil)
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}
	fmt.Println("Data =>", datas)
	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", datas, pagination)
}

func (msUker MsUkerController) SearchUkerByRegionPekerja(c *gin.Context) {
	requests := models.KeyRequest{}

	if err := c.Bind(&requests); err != nil {
		msUker.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := msUker.service.SearchUkerByRegionPekerja(requests)
	if err != nil {
		msUker.logger.Zap.Error(err)
	}

	if pagination.Total == 0 {
		lib.ReturnToJson(c, 200, "404", "Data Kosong", nil)
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}
	fmt.Println("Data =>", datas)
	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", datas, pagination)
}

func (msUker MsUkerController) SearchPesertaPerUker(c *gin.Context) {
	requests := models.KeywordRequest{}

	if err := c.Bind(&requests); err != nil {
		msUker.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := msUker.service.SearchPekerjaPerUker(requests)
	if err != nil {
		msUker.logger.Zap.Error(err)
	}

	if pagination.Total == 0 {
		lib.ReturnToJson(c, 200, "404", "Data Kosong", nil)
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}
	fmt.Println("Data =>", datas)
	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}

func (msUker MsUkerController) SearchSigner(c *gin.Context) {
	requests := models.KeywordRequest{}

	if err := c.Bind(&requests); err != nil {
		msUker.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := msUker.service.SearchSigner(requests)
	if err != nil {
		msUker.logger.Zap.Error(err)
	}

	if pagination.Total == 0 {
		lib.ReturnToJson(c, 200, "404", "Data Kosong", nil)
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}
	fmt.Println("Data =>", datas)
	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}

func (msUker MsUkerController) SearchRMC(c *gin.Context) {
	requests := models.KeywordRequest{}

	if err := c.Bind(&requests); err != nil {
		msUker.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := msUker.service.SearchRMC(requests)
	if err != nil {
		msUker.logger.Zap.Error(err)
	}

	if pagination.Total == 0 {
		lib.ReturnToJson(c, 200, "404", "Data Kosong", nil)
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}
	fmt.Println("Data =>", datas)
	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}

func (msUker MsUkerController) SearchPelakuFraud(c *gin.Context) {
	requests := models.KeywordRequest{}

	if err := c.Bind(&requests); err != nil {
		msUker.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := msUker.service.SearchPelakuFraud(requests)
	if err != nil {
		msUker.logger.Zap.Error(err)
	}

	if pagination.Total == 0 {
		lib.ReturnToJson(c, 200, "404", "Data Kosong", nil)
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}
	fmt.Println("Data =>", datas)
	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}

func (msUker MsUkerController) SearchBrcUrcPerRegion(c *gin.Context) {
	requests := models.KeywordRequest{}

	if err := c.Bind(&requests); err != nil {
		msUker.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := msUker.service.SearchBrcUrcPerRegion(requests)
	if err != nil {
		msUker.logger.Zap.Error(err)
	}

	if pagination.Total == 0 {
		lib.ReturnToJson(c, 200, "404", "Data Kosong", nil)
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}
	fmt.Println("Data =>", datas)
	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}

func (msUker MsUkerController) ListingJabatanPerUker(c *gin.Context) {
	requests := models.ListJabatanRequest{}

	if err := c.Bind(&requests); err != nil {
		msUker.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	datas, err := msUker.service.ListingJabatanPerUker(requests)
	if err != nil {
		msUker.logger.Zap.Error(err)
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}
	fmt.Println("Data =>", datas)
	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}

func (msUker MsUkerController) GetPekerjaBranchByHILFM(c *gin.Context) {
	requests := models.BranchByHilfmRequest{}

	if err := c.Bind(&requests); err != nil {
		msUker.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	datas, err := msUker.service.GetPekerjaBranchByHILFM(requests)
	if err != nil {
		msUker.logger.Zap.Error(err)
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}
	fmt.Println("Data =>", datas)
	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}

func (msUker MsUkerController) SearchRRMHead(c *gin.Context) {
	requests := models.KeywordRequest{}

	if err := c.Bind(&requests); err != nil {
		msUker.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := msUker.service.SearchRRMHead(requests)
	if err != nil {
		msUker.logger.Zap.Error(err)
	}

	if pagination.Total == 0 {
		lib.ReturnToJson(c, 200, "404", "Data Kosong", nil)
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}
	fmt.Println("Data =>", datas)
	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}

func (msUker MsUkerController) SearchPekerjaOrd(c *gin.Context) {
	requests := models.KeywordRequest{}

	if err := c.Bind(&requests); err != nil {
		msUker.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := msUker.service.SearchPekerjaOrd(requests)
	if err != nil {
		msUker.logger.Zap.Error(err)
	}

	if pagination.Total == 0 {
		lib.ReturnToJson(c, 200, "404", "Data Kosong", nil)
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}
	fmt.Println("Data =>", datas)
	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}
