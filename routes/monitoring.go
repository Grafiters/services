package routes

import (
	controllers "riskmanagement/controllers/monitoring"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type MonitoringRoutes struct {
	logger               logger.Logger
	handler              lib.RequestHandler
	MonitoringController controllers.MonitoringController
	authMiddleware       middlewares.JWTAuthMiddleware
}

func (s MonitoringRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("api/v1/monitoring")
	{
		api.POST("/tasklist", s.MonitoringController.GetMonitoringTasklist)
	}
}

func NewMonitoringRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	MonitoringController controllers.MonitoringController,
	authMiddleware middlewares.JWTAuthMiddleware,
) MonitoringRoutes {
	return MonitoringRoutes{
		logger:               logger,
		handler:              handler,
		MonitoringController: MonitoringController,
		authMiddleware:       authMiddleware,
	}
}
