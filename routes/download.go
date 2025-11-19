package routes

import (
	downloadControllers "riskmanagement/controllers/download"
	"riskmanagement/lib"

	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type DownloadRoutes struct {
	logger             logger.Logger
	handler            lib.RequestHandler
	DownloadController downloadControllers.DownloadController
	authMiddleware     middlewares.JWTAuthMiddleware
}

func NewDownloadRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	DownloadController downloadControllers.DownloadController,
	authMiddleware middlewares.JWTAuthMiddleware,
) DownloadRoutes {
	return DownloadRoutes{
		logger:             logger,
		handler:            handler,
		DownloadController: DownloadController,
		authMiddleware:     authMiddleware,
	}
}

func (s DownloadRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/download").Use(s.authMiddleware.Handler())
	{
		api.POST("/getListDownload", s.DownloadController.GetListDownload)
		api.POST("/getReportType", s.DownloadController.GetReportType)
		api.POST("/retryQueue", s.DownloadController.Retry)
		api.POST("/fetchOneData", s.DownloadController.FetchOneRows)
	}
}
