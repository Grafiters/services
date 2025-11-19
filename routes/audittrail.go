package routes

import (
	controllers "riskmanagement/controllers/audittrail"
	downloadControllers "riskmanagement/controllers/download"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type AuditTrailRoutes struct {
	logger             logger.Logger
	handler            lib.RequestHandler
	controllers        controllers.AuditTrailController
	DownloadController downloadControllers.DownloadController
	authMiddleware     middlewares.JWTAuthMiddleware
}

func (s AuditTrailRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/audittrail").Use(s.authMiddleware.Handler())
	{
		api.POST("/store", s.controllers.Store)
		api.POST("/getLog", s.controllers.GetLog)
		api.POST("/generate", s.DownloadController.Generate)
		api.POST("/download/:id", s.DownloadController.DownloadHandler)
	}
}

func NewAuditTrailRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	controllers controllers.AuditTrailController,
	DownloadController downloadControllers.DownloadController,
	authMiddleware middlewares.JWTAuthMiddleware,
) AuditTrailRoutes {
	return AuditTrailRoutes{
		logger:             logger,
		handler:            handler,
		controllers:        controllers,
		DownloadController: DownloadController,
		authMiddleware:     authMiddleware,
	}
}
