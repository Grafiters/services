package mq

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"riskmanagement/lib"
	models "riskmanagement/models/mq"
	commonService "riskmanagement/services/common"
	services "riskmanagement/services/mq"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type MQController struct {
	logger        logger.Logger
	service       services.MQDefinition
	commonService commonService.CommonDefinition
}

func NewMQController(
	MQService services.MQDefinition,
	CommonService commonService.CommonDefinition,
	logger logger.Logger,
) MQController {
	return MQController{
		service:       MQService,
		commonService: CommonService,
		logger:        logger,
	}
}

func (mq MQController) GetAllMenu(c *gin.Context) {
	requests := models.MenuRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.GetAllMenu(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) GetAllMstMenu(c *gin.Context) {
	requests := models.MenuRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.GetAllMstMenu(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) GetTypeCode(c *gin.Context) {
	requests := models.TypeRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.GetTypeCode(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) GetAll(c *gin.Context) {
	requests := models.TypeRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.GetAll(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) GetAllSent(c *gin.Context) {
	requests := models.TypeRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.GetAllSent(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) Filter(c *gin.Context) {
	requests := models.TypeRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.Filter(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) FilterSent(c *gin.Context) {
	requests := models.TypeRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.FilterSent(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) GetTypeList(c *gin.Context) {
	requests := models.TypeRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.GetTypeList(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) GetTypeListAll(c *gin.Context) {
	requests := models.TypeRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.GetTypeListAll(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) GetTypeListAktif(c *gin.Context) {
	requests := models.TypeRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.GetTypeListAktif(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) GetOne(c *gin.Context) {
	requests := models.TypeRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.GetOne(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) Store(c *gin.Context) {
	requests := models.TypeRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.Store(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) Update(c *gin.Context) {
	requests := models.TypeRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.Update(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) Delete(c *gin.Context) {
	requests := models.TypeRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.Delete(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) SetStatusType(c *gin.Context) {
	requests := models.TypeRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.SetStatusType(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) GetAllORDMember(c *gin.Context) {
	data, err := mq.commonService.GetAllORDMember()

	if err != nil {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) GetDataQuestForm(c *gin.Context) {
	requests := models.TypeRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.GetDataQuestForm(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) SendQuest(c *gin.Context) {
	requests := models.TypeRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.SendQuest(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) ApproverUpdate(c *gin.Context) {
	requests := models.ApproverRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.ApproverUpdate(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) GenerateQuestionnaireRequest(c *gin.Context) {
	requests := models.GenerateQuestionnaireRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.GenerateQuestionnaireRequest(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) WeightTotal(c *gin.Context) {
	requests := models.GenerateRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.WeightTotal(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) UpdatePartWeight(c *gin.Context) {
	requests := models.PartWeightRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.UpdatePartWeight(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) GetPartCode(c *gin.Context) {
	requests := models.PartRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.GetPartCode(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) GetAllPart(c *gin.Context) {
	requests := models.PartRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.GetAllPart(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) FilterPart(c *gin.Context) {
	requests := models.PartRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.FilterPart(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) StorePart(c *gin.Context) {
	requests := models.PartRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.StorePart(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) UpdatePart(c *gin.Context) {
	requests := models.PartRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.UpdatePart(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) DeletePart(c *gin.Context) {
	requests := models.PartRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.DeletePart(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) GetOnePart(c *gin.Context) {
	requests := models.PartRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.GetOnePart(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) StoreSubPart(c *gin.Context) {
	requests := models.SubPart{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.StoreSubPart(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) UpdateSubPart(c *gin.Context) {
	requests := models.SubPartEditRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.UpdateSubPart(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) DeleteSubPart(c *gin.Context) {
	requests := models.SubPartListRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.DeleteSubPart(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) GetSubPartList(c *gin.Context) {
	requests := models.SubPartListRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.GetSubPartList(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) GetSubPartKode(c *gin.Context) {
	requests := models.SubPartListRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.GetSubPartKode(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) GetDetailSubPart(c *gin.Context) {
	requests := models.SubPartListRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.GetDetailSubPart(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) StoreQuestion(c *gin.Context) {
	requests := models.RequestQuestion{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.StoreQuestion(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) UpdateQuestion(c *gin.Context) {
	requests := models.RequestQuestionUpdate{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.UpdateQuestion(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) DeleteQuestion(c *gin.Context) {
	requests := models.RequestQuestion{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.DeleteQuestion(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) GetQuestionList(c *gin.Context) {
	requests := models.QuestionListRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.GetQuestionList(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) GetKodeQuestionnaire(c *gin.Context) {
	requests := models.QuestionListRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.GetKodeQuestionnaire(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) GetDetailQuestionnaire(c *gin.Context) {
	requests := models.QuestionListRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.GetDetailQuestionnaire(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) CommonGetType(c *gin.Context) {
	requests := models.CommonRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.CommonGetType(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) CommonGetPart(c *gin.Context) {
	requests := models.CommonRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.CommonGetPart(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) CommonGetSubPart(c *gin.Context) {
	requests := models.CommonRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.CommonGetSubPart(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) GetDetailType(c *gin.Context) {
	requests := models.CommonRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.GetDetailType(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) GetApprovalResponse(c *gin.Context) {
	requests := models.CommonRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.GetApprovalResponse(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) StoreLinkcage(c *gin.Context) {
	requests := models.LinkcageRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.StoreLinkcage(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) GetActive(c *gin.Context) {
	requests := models.LinkcageRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.GetActive(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) GetAllLinkcage(c *gin.Context) {
	requests := models.LinkcageRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.GetAllLinkcage(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) UpdateLinkcage(c *gin.Context) {
	requests := models.LinkcageRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.UpdateLinkcage(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) DeleteLinkcage(c *gin.Context) {
	requests := models.LinkcageRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.DeleteLinkcage(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) GetOneLinkcage(c *gin.Context) {
	requests := models.LinkcageRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.GetOneLinkcage(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) SetStatus(c *gin.Context) {
	requests := models.LinkcageRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.SetStatus(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

// =================response user=================
func (mq MQController) GetResponseUserList(c *gin.Context) {
	requests := models.RequestResponseUserList{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.GetResponseUserList(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) GetResponseApprovalList(c *gin.Context) {
	requests := models.RequestResponseApprovalList{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.GetResponseApprovalList(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) StoreResponseUser(c *gin.Context) {
	requests := models.RequestUserHistory{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.StoreResponseUser(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) GenerateQuestWithAnswer(c *gin.Context) {
	requests := models.GenerateRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.GenerateQuestWithAnswer(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) UpdateResponseUser(c *gin.Context) {
	requests := models.RequestUserHistory{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.UpdateResponseUser(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) ApproveResponseUser(c *gin.Context) {
	requests := models.ApprovalUpdate{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.ApproveResponseUser(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) RejectResponseUser(c *gin.Context) {
	requests := models.RejectedUpdate{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.RejectResponseUser(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) GetRejectResponse(c *gin.Context) {
	requests := models.RejectedUpdate{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.GetRejectResponse(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) GetNilaiAkhir(c *gin.Context) {
	requests := models.GenerateRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.GetNilaiAkhir(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) GenerateQuestPerPage(c *gin.Context) {
	requests := models.GenerateRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.GenerateQuestPerPage(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) GenerateQuestForApprover(c *gin.Context) {
	requests := models.GenerateRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.GenerateQuestForApprover(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) GenerateQuestPerPagePreview(c *gin.Context) {
	requests := models.GenerateRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.GenerateQuestPerPagePreview(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

// perbaikan DAST
func (mq MQController) GenerateQuestPerPagePreviewApprover(c *gin.Context) {
	requests := models.GenerateRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.GenerateQuestPerPagePreviewApprover(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) ProcessGeneratePagination(c *gin.Context) {
	requests := models.RequestUserHistory{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.ProcessGeneratePagination(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) CancelResponseUser(c *gin.Context) {
	requests := models.UpdateResponseUserHistory{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.CancelResponseUser(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) GetPartPagination(c *gin.Context) {
	requests := models.GenerateRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.GetPartPagination(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) GetPartPaginationDraft(c *gin.Context) {
	requests := models.GenerateRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.GetPartPaginationDraft(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) DisableQuestionByPart(c *gin.Context) {
	requests := models.RequestPartid{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.DisableQuestionByPart(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) GetReportList(c *gin.Context) {
	requests := models.RequestReportList{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.GetReportList(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) GenerateReportPerPage(c *gin.Context) {
	requests := models.GenerateRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.GenerateReportPerPage(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) GetNamaRespondenList(c *gin.Context) {
	requests := models.GenerateRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.GetNamaRespondenList(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) ResponseDownload(c *gin.Context) {
	requests := models.ReportListQuery{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.ResponseDownload(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	fmt.Println("responseeee download ====>", data)

	dataMap, ok := data.(map[string]interface{})
	if !ok {
		fmt.Println("Data tidak bertipe map[string]interface{}")
		return
	}

	nestedData, ok := dataMap["data"].(map[string]interface{})
	if !ok {
		fmt.Println("Key 'data' tidak ditemukan atau bukan map[string]interface{}")
		return
	}

	// Kemudian, akses nilai "file" dan "url" dari nested map
	fileName, _ := nestedData["file"].(string)
	url, _ := nestedData["url"].(string)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error retriveing file:", err)
	}
	defer resp.Body.Close()

	// Set the appropriate headers for file download
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

	_, err = io.Copy(c.Writer, resp.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	// lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mq MQController) GetSummary(c *gin.Context) {
	requests := models.GenerateRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		mq.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := mq.service.GetSummary(requests)
	if err != nil {
		mq.logger.Zap.Error()
	}
	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}
