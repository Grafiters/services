package controllers

import (
	"database/sql"
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/penyebabkejadianlv3"
	services "riskmanagement/services/penyebabkejadianlv3"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type PenyebabKejadianLv3Controller struct {
	logger  logger.Logger
	service services.PenyebabKejadianLv3Definition
}

func NewPenyebabKejadianLv3Controller(
	PenyebabKejadianLv3Service services.PenyebabKejadianLv3Definition,
	logger logger.Logger,
) PenyebabKejadianLv3Controller {
	return PenyebabKejadianLv3Controller{
		service: PenyebabKejadianLv3Service,
		logger:  logger,
	}
}

func (PK3 PenyebabKejadianLv3Controller) GetAll(c *gin.Context) {
	datas, err := PK3.service.GetAll()
	if err != nil {
		PK3.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}

func (pk3 PenyebabKejadianLv3Controller) GetAllWithPaginate(c *gin.Context) {
	requests := models.Paginate{}

	if err := c.Bind(&requests); err != nil {
		pk3.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := pk3.service.GetAllWithPaginate(requests)
	if err != nil {
		pk3.logger.Zap.Error(err)
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

func (PK3 PenyebabKejadianLv3Controller) GetOne(c *gin.Context) {
	requests := models.PenyebabKejadianLv3Response{}

	if err := c.Bind(&requests); err != nil {
		PK3.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := PK3.service.GetOne(requests.ID)
	if err != nil {
		PK3.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (PK3 PenyebabKejadianLv3Controller) Store(c *gin.Context) {
	data := models.PenyebabKejadianLv3Request{}

	if err := c.Bind(&data); err != nil {
		PK3.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak sesuai : "+err.Error(), "")
		return
	}

	fmt.Println(data)
	if err := PK3.service.Store(&data); err != nil {
		PK3.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "200", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Input data berhasil", data)
}

func (PK3 PenyebabKejadianLv3Controller) Update(c *gin.Context) {
	data := models.PenyebabKejadianLv3Request{}

	if err := c.Bind(&data); err != nil {
		PK3.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak sesuai : "+err.Error(), "")
		return
	}

	if err := PK3.service.Update(&data); err != nil {
		PK3.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", data)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Update Data Berhasil", data)
}

func (PK3 PenyebabKejadianLv3Controller) Delete(c *gin.Context) {
	requests := models.PenyebabKejadianLv3Response{}

	if err := c.Bind(&requests); err != nil {
		PK3.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	if err := PK3.service.Delete(requests.ID); err != nil {
		PK3.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Data berhasil dihapus", true)
}

func (PK3 PenyebabKejadianLv3Controller) GetKodePenyebabKejadian(c *gin.Context) {
	datas, err := PK3.service.GetKodePenyebabKejadian()
	if err != nil {
		PK3.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if len(datas) == 0 {
		PK3.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data Tidak Ditemukan !", nil)
		return
	}

	KodePenyebabKejadian := "PK3.MOP." + datas[0].KodePenyebabKejadian

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", KodePenyebabKejadian)
}

func (PK3 PenyebabKejadianLv3Controller) GetKejadianByIDlv2(c *gin.Context) {
	requests := models.KodeSubKejadianRequest{}

	if err := c.Bind(&requests); err != nil {
		PK3.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai :"+err.Error(), "")
		return
	}

	datas, err := PK3.service.GetKejadianByIDlv2(requests)
	if err != nil {
		PK3.logger.Zap.Error(err)
	}

	if len(datas) == 0 {
		PK3.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data tidak ditemukan", nil)
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	// fmt.Println("SubIncident =>", datas)
	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas)

}

func (PK3 PenyebabKejadianLv3Controller) GetKejadianByIDlv1(c *gin.Context) {
	requests := models.KodePenyebabKejadian{}

	if err := c.Bind(&requests); err != nil {
		PK3.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai :"+err.Error(), "")
		return
	}

	datas, err := PK3.service.GetKejadianByIDlv1(requests)
	if err != nil {
		PK3.logger.Zap.Error(err)
	}

	if len(datas) == 0 {
		PK3.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data tidak ditemukan", nil)
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	// fmt.Println("SubIncident =>", datas)
	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas)

}

func (PK3 PenyebabKejadianLv3Controller) GetSubKejadian(c *gin.Context) {
	requests := models.PenyebabKejadianLv3Request{}

	if err := c.Bind(&requests); err != nil {
		PK3.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai :"+err.Error(), "")
		return
	}

	datas, err := PK3.service.GetSubKejadian(requests.ID)
	if err != nil {
		PK3.logger.Zap.Error(err)
	}

	if len(datas) == 0 {
		PK3.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data tidak ditemukan", nil)
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	// fmt.Println("SubIncident =>", datas)
	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas)

}
