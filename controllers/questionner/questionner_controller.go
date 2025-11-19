package controller

import (
	"riskmanagement/lib"
	services "riskmanagement/services/questionner"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type QuestionnerController struct {
	logger   logger.Logger
	services services.QuestionnerDefinition
}

func NewQuestionnerController(
	logger logger.Logger,
	services services.QuestionnerDefinition,
) QuestionnerController {
	return QuestionnerController{
		logger:   logger,
		services: services,
	}
}

func (q QuestionnerController) GetQuestionnerList(c *gin.Context) {
	datas, err := q.services.GetQuestionner()

	if err != nil {
		q.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "200", "Internal Error", "")
		return
	}

	if len(datas) == 0 {
		q.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data tidak ditemukan", "")
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas)
}
