package controllers

import (
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/krid"
	services "riskmanagement/services/krid"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type KridController struct {
	logger  logger.Logger
	service services.KridDefinition
}

func NewKridController(
	kridService services.KridDefinition,
	logger logger.Logger,
) KridController {
	return KridController{
		logger:  logger,
		service: kridService,
	}
}

func (krid KridController) GetDetailIndikator(c *gin.Context) {
	data := models.HeaderRequest{}

	if err := c.Bind(&data); err != nil {
		krid.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}
	response, err := krid.service.GetDetailIndikator(&data)

	if err != nil {
		krid.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if len(response) == 0 {
		krid.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data Tidak Ditemukan", nil)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", response)
}

func (krid KridController) GetAllParameterIndikator(c *gin.Context) {
	response, err := krid.service.GetAllParameterIndikator()

	if err != nil {
		krid.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if len(response) == 0 {
		krid.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data Tidak Ditemukan", nil)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", response)
}

func (krid KridController) SearchIndicator(c *gin.Context) {
	request := models.KeywordSearch{}

	// fmt.Println("controller =>", request)

	if err := c.Bind(&request); err != nil {
		krid.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	response, err := krid.service.SearchIndikatorKRI(&request)

	if err != nil {
		krid.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	// if len(response) == 0 {
	// 	krid.logger.Zap.Error(err)
	// 	lib.ReturnToJson(c, 200, "404", "Data Tidak Ditemukan", nil)
	// 	return
	// }

	fmt.Println("isi response:", response)

	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", response)
}

func (krid KridController) SearchIndicatorEdit(c *gin.Context) {
	request := models.KeywordSearchEdit{}

	// fmt.Println("controller =>", request)

	if err := c.Bind(&request); err != nil {
		krid.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	response, err := krid.service.SearchIndikatorKRIEdit(&request)

	if err != nil {
		krid.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	// if len(response) == 0 {
	// 	krid.logger.Zap.Error(err)
	// 	lib.ReturnToJson(c, 200, "404", "Data Tidak Ditemukan", nil)
	// 	return
	// }

	fmt.Println("isi response:", response)

	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", response)
}
