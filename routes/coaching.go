package routes

import (
	controllers "riskmanagement/controllers/coaching"
	downloadControllers "riskmanagement/controllers/download"
	"riskmanagement/lib"

	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type CoachingRoutes struct {
	logger             logger.Logger
	handler            lib.RequestHandler
	CoachingController controllers.CoachingController
	authMiddleware     middlewares.JWTAuthMiddleware
	DownloadController downloadControllers.DownloadController
	audiTrail          middlewares.AuditTrailMiddleware
}

func (s CoachingRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/coaching").
		Use(s.authMiddleware.Handler()).
		Use(s.audiTrail.Handler())
	{
		api.POST("/getAll", s.CoachingController.GetAll)
		api.POST("/getOne", s.CoachingController.GetOne)
		api.POST("/store", s.CoachingController.Store)
		api.POST("/storeDraft", s.CoachingController.StoreDraft)
		api.POST("/deleteCoachingMateri", s.CoachingController.DeleteCoachingActivity)
		api.POST("/delete", s.CoachingController.Delete)
		api.POST("/update", s.CoachingController.UpdateAllCoaching)
		api.POST("/updateDraft", s.CoachingController.UpdateDraft)
		api.POST("/filterCoaching", s.CoachingController.FilterCoaching)
		api.POST("/getNoPelaporan", s.CoachingController.GetNoPelaporan)
		api.POST("/getData", s.CoachingController.GetData)
		api.POST("/getDataWithPagination", s.CoachingController.GetDataWithPagination)
		api.POST("/coachingReportFilter", s.CoachingController.CoachingReportFilter)
		api.POST("/coachingFinalReportFilter", s.CoachingController.CoachingFinalReportFilter)
		api.POST("/coachingReportDetail", s.CoachingController.CoachingReportDetail)
		api.POST("/coachingReportMateriList", s.CoachingController.CoachingReportMateriList)
		api.POST("/coachingReportFilterByUkerAllActivity", s.CoachingController.CoachingReportFilterByUkerAllActivity)
		api.POST("/coachingReportByUkerFilter", s.CoachingController.CoachingReportByUkerFilter)
		api.POST("/coachingReportFilterByUkerComplete", s.CoachingController.CoachingReportFilterByUkerComplete)
		api.POST("/coachingReportList", s.CoachingController.CoachingReportList)
		api.POST("/generate", s.DownloadController.Generate)
		api.POST("/download/:id", s.DownloadController.DownloadHandler)
		api.POST("/deleteMapPeserta", s.CoachingController.DeleteMapPeserta)
		api.POST("/coachingfrekuensi", s.CoachingController.CoachingFrekuensiRpt)
	}
}

func NewCoachingRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	CoachinngController controllers.CoachingController,
	DownloadController downloadControllers.DownloadController,
	authMiddleware middlewares.JWTAuthMiddleware,
	audiTrail middlewares.AuditTrailMiddleware,
) CoachingRoutes {
	return CoachingRoutes{
		logger:             logger,
		handler:            handler,
		CoachingController: CoachinngController,
		DownloadController: DownloadController,
		authMiddleware:     authMiddleware,
		audiTrail:          audiTrail,
	}
}
