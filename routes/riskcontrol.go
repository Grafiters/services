package routes

import (
	controller "riskmanagement/controllers/riskcontrol"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type RiskControlRoutes struct {
	logger                logger.Logger
	handler               lib.RequestHandler
	RiskControlController controller.RiskControlController
	authMiddleware        middlewares.JWTAuthMiddleware
}

func (s RiskControlRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("api/v1/riskcontrol").Use(s.authMiddleware.Handler())
	{
		api.POST("/getAll", s.RiskControlController.GetAll)
		api.POST("/getAllWithPaginate", s.RiskControlController.GetAllWithPaginate)
		api.POST("/getOne", s.RiskControlController.GetOne)
		api.POST("/store", s.RiskControlController.Store)
		api.POST("/update", s.RiskControlController.Update)
		api.POST("/delete", s.RiskControlController.Delete)
		api.POST("/getKode", s.RiskControlController.GetKodeRiskControl)
		api.POST("/preview", s.RiskControlController.PreviewImport)
		api.POST("/import", s.RiskControlController.ImportData)
		api.POST("/template", s.RiskControlController.Template)
		api.POST("/code", s.RiskControlController.GenCode)
		api.POST("/download/:format", s.RiskControlController.Download)
		api.PATCH("/status", s.RiskControlController.UpdateStatus)
		api.POST("/searchRiskControl", s.RiskControlController.SearchRiskIndicatorByIssue)
	}
}

func NewRiskControlRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	RiskControlConteroller controller.RiskControlController,
	authMiddleware middlewares.JWTAuthMiddleware,
) RiskControlRoutes {
	return RiskControlRoutes{
		logger:                logger,
		handler:               handler,
		RiskControlController: RiskControlConteroller,
		authMiddleware:        authMiddleware,
	}
}
