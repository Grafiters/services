package controller

import (
	"database/sql"
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/product"
	services "riskmanagement/services/product"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type ProductController struct {
	logger  logger.Logger
	service services.ProductDefinition
}

func NewProductController(
	ProductService services.ProductDefinition,
	logger logger.Logger,
) ProductController {
	return ProductController{
		service: ProductService,
		logger:  logger,
	}
}

func (product ProductController) GetAll(c *gin.Context) {
	datas, err := product.service.GetAll()
	if err != nil {
		product.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas)
}

func (product ProductController) GetOne(c *gin.Context) {
	requests := models.ProductRequest{}

	if err := c.Bind(&requests); err != nil {
		product.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := product.service.GetOne(requests.ID)
	if err != nil {
		product.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}

func (product ProductController) Store(c *gin.Context) {
	data := models.ProductRequest{}

	if err := c.Bind(&data); err != nil {
		product.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	fmt.Println(data)
	status, err := product.service.Store(&data)
	if err != nil {
		product.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	if !status {
		lib.ReturnToJson(c, 200, "400", "Produk '"+data.Product+"' sudah ada !", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Input data berhasil", true)
}

func (product ProductController) Update(c *gin.Context) {
	data := models.ProductRequest{}

	if err := c.Bind(&data); err != nil {
		product.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	if err := product.service.Update(&data); err != nil {
		product.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", data)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Update data berhasil", data)
}

func (product ProductController) Delete(c *gin.Context) {
	requests := models.ProductRequest{}

	if err := c.Bind(&requests); err != nil {
		product.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	if err := product.service.Delete(requests.ID); err != nil {
		product.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Data berhasil dihapus", true)
}

func (product ProductController) SearchProduct(c *gin.Context) {
	requests := models.KeywordRequest{}

	if err := c.Bind(&requests); err != nil {
		product.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := product.service.SearchProduct(requests)
	if err != nil {
		product.logger.Zap.Error(err)
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

func (product ProductController) GetKodeProduct(c *gin.Context) {
	datas, err := product.service.GetKodeProduct()
	if err != nil {
		product.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if len(datas) == 0 {
		product.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data Tidak Ditemukan !", nil)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas[0].KodeProduct)
}

func (product ProductController) GetAllWithPage(c *gin.Context) {
	requests := models.PageRequest{}

	if err := c.Bind(&requests); err != nil {
		product.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := product.service.GetAllWithPage(requests)
	if err != nil {
		product.logger.Zap.Error(err)
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

func (product ProductController) GetProductByActivity(c *gin.Context) {
	requests := models.GetProductByActivityRequest{}

	if err := c.Bind(&requests); err != nil {
		product.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := product.service.GetProductByActivity(requests)
	if err != nil {
		product.logger.Zap.Error(err)
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery data berhasil", datas, pagination)
}

func (product ProductController) GetProductBySegment(c *gin.Context) {
	requests := models.GetProductBySegmentRequest{}

	if err := c.Bind(&requests); err != nil {
		product.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := product.service.GetProductBySegment(requests)
	if err != nil {
		product.logger.Zap.Error(err)
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery data berhasil", datas, pagination)
}
