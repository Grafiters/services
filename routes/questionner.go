package routes

import (
	controllers "riskmanagement/controllers/questionner"
	"riskmanagement/lib"

	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type QuestionnerRoutes struct {
	logger             logger.Logger
	handler            lib.RequestHandler
	questionController controllers.QuestionnerController
	authMiddleware     middlewares.JWTAuthMiddleware
}

func (s QuestionnerRoutes) Setup() {
	s.logger.Zap.Info("Setting Up Routes")
	api := s.handler.Gin.Group("/api/v1/questionner").Use(s.authMiddleware.Handler())
	{
		api.POST("/getQuestionnerList", s.questionController.GetQuestionnerList)
	}
}

func NewQuestionnerRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	questionController controllers.QuestionnerController,
	authMiddleware middlewares.JWTAuthMiddleware,
) QuestionnerRoutes {
	return QuestionnerRoutes{
		logger:             logger,
		handler:            handler,
		questionController: questionController,
		authMiddleware:     authMiddleware,
	}
}
