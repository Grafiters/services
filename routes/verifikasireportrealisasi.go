package routes

import (
	downloadControllers "riskmanagement/controllers/download"
	controllers "riskmanagement/controllers/verifikasireportrealisasi"
	"riskmanagement/lib"

	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type VerifikasiReportRealisasiRoutes struct {
	logger                              logger.Logger
	handler                             lib.RequestHandler
	VerifikasiReportRealisasiController controllers.VerifikasiReportRealisasiController
	DownloadController                  downloadControllers.DownloadController
	authMiddleware                      middlewares.JWTAuthMiddleware
}

func (s VerifikasiReportRealisasiRoutes) Setup() {
	s.logger.Zap.Info("Setting Up Routes")
	api := s.handler.Gin.Group("/api/v1/verifikasireportrealisasi").Use(s.authMiddleware.Handler())
	{
		// Report Realisasi Kredit 11/07/2024
		api.POST("/reportRealisasiKreditListFilter", s.VerifikasiReportRealisasiController.ReportRealisasiKreditListFilter)
		api.POST("/reportRealisasiKreditSummaryFilter", s.VerifikasiReportRealisasiController.ReportRealisasiKreditSummaryFilter)
		api.POST("/segmentRealisasiKredit", s.VerifikasiReportRealisasiController.GetAllSegmentRealisasiKredit)
		api.POST("/generate", s.DownloadController.Generate)
		api.POST("/download/:id", s.DownloadController.DownloadHandler)
	}
}

func NewVerifikasiReportRealisasiRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	VerifikasiReportRealisasiController controllers.VerifikasiReportRealisasiController,
	DownloadController downloadControllers.DownloadController,
	authMiddleware middlewares.JWTAuthMiddleware,
) VerifikasiReportRealisasiRoutes {
	return VerifikasiReportRealisasiRoutes{
		logger:                              logger,
		handler:                             handler,
		VerifikasiReportRealisasiController: VerifikasiReportRealisasiController,
		DownloadController:                  DownloadController,
		authMiddleware:                      authMiddleware,
	}
}
