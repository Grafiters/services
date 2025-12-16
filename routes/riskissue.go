package routes

import (
	controller "riskmanagement/controllers/riskissue"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type RiskIssueRoutes struct {
	logger              logger.Logger
	handler             lib.RequestHandler
	RiskIssueController controller.RiskIssueController
	authMiddleware      middlewares.JWTAuthMiddleware
}

func (s RiskIssueRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/riskissue")
	{
		api.POST("/getAll", s.RiskIssueController.GetAll)
		api.POST("/getAllWithPaginate", s.RiskIssueController.GetAllWithPaginate)
		api.POST("/getOne", s.RiskIssueController.GetOne)
		api.POST("/store", s.RiskIssueController.Store)
		api.POST("/update", s.RiskIssueController.Update)
		api.POST("/deleteMapAktifitas", s.RiskIssueController.DeleteMapAktifitas)
		api.POST("/deleteMapEvent", s.RiskIssueController.DeleteMapEvent)
		api.POST("/deleteMapKejadian", s.RiskIssueController.DeleteMapKejadian)
		api.POST("/deleteMapLiniBisnis", s.RiskIssueController.DeleteMapLiniBisnis)
		api.POST("/deleteMapProduct", s.RiskIssueController.DeleteMapProduct)
		api.POST("/deleteMapProses", s.RiskIssueController.DeleteMapProses)
		api.POST("/deleteMapControl", s.RiskIssueController.DeleteMapControl)
		api.POST("/getKode", s.RiskIssueController.GetKode)
		api.POST("/mappingControl", s.RiskIssueController.MappingRiskControl)
		api.POST("/getControlByID", s.RiskIssueController.GetMappingControlbyID)
		api.POST("/:id/getControlWithPaginate", s.RiskIssueController.GetMappingControlWithPaginate)
		api.POST("/:id/getIndicatorWithPaginate", s.RiskIssueController.GetMappingIndicatorWithPaginate)
		api.POST("/ListRiskIssue", s.RiskIssueController.ListRiskIssue)
		api.POST("/getByIndicator", s.RiskIssueController.ListRiskIssue)
		api.POST("/searchRiskIssue", s.RiskIssueController.SearchRiskIssue)
		api.POST("/searchRiskIssueWithoutSub", s.RiskIssueController.SearchRiskIssueWithoutSub)
		// api.GET("/getLastID/:id", s.RiskIssueController.GetLastID)
		api.POST("/delete", s.RiskIssueController.Delete)
		api.POST("/mappingIndicator", s.RiskIssueController.MappingRiskIndicator)
		api.POST("/getIndicatorByID", s.RiskIssueController.GetMappingIndicatorbyID)
		api.POST("/deleteMapIndicator", s.RiskIssueController.DeleteMapIndicator)
		api.POST("/filterRiskIssue", s.RiskIssueController.FilterRiskIssue)
		api.POST("/getRiskIssueByActivity", s.RiskIssueController.GetRiskIssueByActivity)
		api.POST("/getRekomendasiMateri", s.RiskIssueController.GetRekomendasiMateri)
		api.POST("/getMateriByCode", s.RiskIssueController.GetMateriByCode)
		api.POST("/getRiskIssueByActivityID", s.RiskIssueController.GetRiskIssueByActivityID)

		api.POST("/status", s.RiskIssueController.UpdateStatus)
		api.POST("/template", s.RiskIssueController.Template)
		api.POST("/preview", s.RiskIssueController.PreviewData)
		api.POST("/import", s.RiskIssueController.ImportData)
		api.POST("/download/:format", s.RiskIssueController.Download)
		api.POST("/getRiskCategory", s.RiskIssueController.GetRiskCategories)

	}
}

func NewRiskIssueRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	RiskIssueController controller.RiskIssueController,
	authMiddleware middlewares.JWTAuthMiddleware,
) RiskIssueRoutes {
	return RiskIssueRoutes{
		handler:             handler,
		logger:              logger,
		RiskIssueController: RiskIssueController,
		authMiddleware:      authMiddleware,
	}
}
