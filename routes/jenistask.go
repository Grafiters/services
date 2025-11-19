package routes

import (
	controller "riskmanagement/controllers/jenistask"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type JenisTaskRoutes struct {
	logger              logger.Logger
	handler             lib.RequestHandler
	JenisTaskController controller.JenisTaskController
	authMiddleware      middlewares.JWTAuthMiddleware
}

func (s JenisTaskRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/jenistask").Use(s.authMiddleware.Handler())
	{
		api.POST("/getData", s.JenisTaskController.GetData)
	}
}

func NewJenisTaskRouter(
	logger logger.Logger,
	handler lib.RequestHandler,
	JenisTaskController controller.JenisTaskController,
	authMiddleware middlewares.JWTAuthMiddleware,
) JenisTaskRoutes {
	return JenisTaskRoutes{
		handler:             handler,
		logger:              logger,
		JenisTaskController: JenisTaskController,
		authMiddleware:      authMiddleware,
	}
}
