package routes

import (
	controllers "riskmanagement/controllers/briefing"
	downloadControllers "riskmanagement/controllers/download"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type BriefingRoutes struct {
	logger             logger.Logger
	handler            lib.RequestHandler
	BriefingController controllers.BriefingController
	DownloadController downloadControllers.DownloadController
	authMiddleware     middlewares.JWTAuthMiddleware
	audiTrail          middlewares.AuditTrailMiddleware
}

func (s BriefingRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/briefing").Use(s.authMiddleware.Handler()).
		Use(s.authMiddleware.Handler()).
		Use(s.audiTrail.Handler())
	{
		api.POST("/getAll", s.BriefingController.GetAll)
		api.POST("/getData", s.BriefingController.GetData)
		api.POST("/getDataWithPagination", s.BriefingController.GetDataWithPagination)
		api.POST("/getOne", s.BriefingController.GetOne)
		api.POST("/store", s.BriefingController.Store)
		api.POST("/storeDraft", s.BriefingController.StoreDraft)
		api.POST("/deleteBriefingMateri", s.BriefingController.DeleteBriefingMateri)
		api.POST("/delete", s.BriefingController.Delete)
		api.POST("/update", s.BriefingController.UpdateAllBrief)
		api.POST("/updateDraft", s.BriefingController.UpdateDraft)
		api.POST("/filterBriefing", s.BriefingController.FilterBriefing)
		api.POST("/getNoPelaporan", s.BriefingController.GetNoPelaporan)
		api.POST("/BriefingReportFilter", s.BriefingController.BriefingReportFilter)
		api.POST("/deleteMapPeserta", s.BriefingController.DeleteMapPeserta)
		api.POST("/BriefingReportFinalFilter", s.BriefingController.BriefingReportFinalFilter)
		api.POST("/BriefingReportDetail", s.BriefingController.BriefingReportDetail)
		api.POST("/BriefingReportMateriList", s.BriefingController.BriefingReportMateriList)
		api.POST("/BriefingReportByUkerFilter", s.BriefingController.BriefingReportByUkerFilter)
		api.POST("/BriefingReportFilterByUkerComplete", s.BriefingController.BriefingReportFilterByUkerComplete)
		api.POST("/BriefingReportList", s.BriefingController.BriefingReportList)
		api.POST("/generate", s.DownloadController.Generate)
		api.POST("/download/:id", s.DownloadController.DownloadHandler)
		api.POST("/briefingfrekuensi", s.BriefingController.BriefingFrekuensiRpt)
	}
}

func NewBriefingRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	BriefingController controllers.BriefingController,
	DownloadController downloadControllers.DownloadController,
	authMiddleware middlewares.JWTAuthMiddleware,
	audiTrail middlewares.AuditTrailMiddleware,
) BriefingRoutes {
	return BriefingRoutes{
		logger:             logger,
		handler:            handler,
		BriefingController: BriefingController,
		DownloadController: DownloadController,
		authMiddleware:     authMiddleware,
		audiTrail:          audiTrail,
	}
}
