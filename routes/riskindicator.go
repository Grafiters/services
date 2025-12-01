package routes

import (
	downloadControllers "riskmanagement/controllers/download"
	controller "riskmanagement/controllers/riskindicator"
	"riskmanagement/middlewares"

	"riskmanagement/lib"

	"gitlab.com/golang-package-library/logger"
)

type RiskIndicatorRoutes struct {
	logger                  logger.Logger
	handler                 lib.RequestHandler
	RiskIndicatorController controller.RiskIndicatorController
	DownloadController      downloadControllers.DownloadController

	authMiddleware middlewares.JWTAuthMiddleware
}

func (s RiskIndicatorRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/riskindicator").Use(s.authMiddleware.Handler())
	{
		api.POST("/getAll", s.RiskIndicatorController.GetAll)
		api.POST("/GetAllWithPaginate", s.RiskIndicatorController.GetAllWithPaginate)
		api.POST("/getOne", s.RiskIndicatorController.GetOne)
		api.POST("/store", s.RiskIndicatorController.Store)
		api.POST("/update", s.RiskIndicatorController.Update)
		api.POST("/delete", s.RiskIndicatorController.Delete)
		api.POST("/deleteFileByID", s.RiskIndicatorController.DeleteFilesByID)
		api.POST("/searchRiskIndicatorByIssue", s.RiskIndicatorController.SearchRiskIndicatorByIssue)
		api.POST("/getRekomendasiMateri", s.RiskIndicatorController.GetRekomendasiMateri)
		api.POST("/searchRiskIndicatorKRID", s.RiskIndicatorController.SearchRiskIndicatorKRID)
		api.POST("/getKode", s.RiskIndicatorController.GetKode)
		api.POST("/filterRiskIndicator", s.RiskIndicatorController.FilterRiskIndicator)
		api.POST("/saveThreshold", s.RiskIndicatorController.SaveThreshold)
		api.POST("/getThreshold", s.RiskIndicatorController.GetThreshold)
		api.POST("/getMapRiskIssue", s.RiskIndicatorController.GetMapRiskIssue)
		api.POST("/generate", s.DownloadController.Generate)
		api.POST("/download/:id", s.DownloadController.DownloadHandler)
		api.POST("/getIndicatorByAktivityProduct", s.RiskIndicatorController.GetIndicatorByAktivityProduct)

		// Batch3
		api.POST("/searchRiskIndicatorTematik", s.RiskIndicatorController.SearchRiskIndicatorTematik)
		api.POST("/getDataTematik", s.RiskIndicatorController.GetTematikData)
		api.POST("/getDocumentList", s.RiskIndicatorController.GetMateriIfFinish)
		api.POST("/template", s.RiskIndicatorController.Template)
		api.POST("/preview", s.RiskIndicatorController.Preview)
		api.POST("/import", s.RiskIndicatorController.ImportData)
		api.POST("/export/:format", s.RiskIndicatorController.Download)
	}
}

func NewRiskIndicatorRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	RiskIndicatorController controller.RiskIndicatorController,
	DownloadController downloadControllers.DownloadController,
	authMiddleware middlewares.JWTAuthMiddleware,
) RiskIndicatorRoutes {
	return RiskIndicatorRoutes{
		handler:                 handler,
		logger:                  logger,
		RiskIndicatorController: RiskIndicatorController,
		DownloadController:      DownloadController,
		authMiddleware:          authMiddleware,
	}
}
