package routes

import (
	controller "riskmanagement/controllers/uploaddata"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type UploadDataRoutes struct {
	logger               logger.Logger
	handler              lib.RequestHandler
	authMiddleware       middlewares.JWTAuthMiddleware
	UploadDataController controller.UploadDataController
}

func (s UploadDataRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("api/v1/upload").Use(s.authMiddleware.Handler())
	{
		api.POST("/riskcontrol", s.UploadDataController.UploadRiskControl)
		api.POST("/riskindicator", s.UploadDataController.UploadRiskIndicator)
		api.POST("/riskevent", s.UploadDataController.UploadRiskEvent)
	}
}

func NewUploadDataRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	authMiddleware middlewares.JWTAuthMiddleware,
	UploadDataController controller.UploadDataController,
) UploadDataRoutes {
	return UploadDataRoutes{
		logger:               logger,
		handler:              handler,
		authMiddleware:       authMiddleware,
		UploadDataController: UploadDataController,
	}
}
