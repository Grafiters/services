package routes

import (
	controller "riskmanagement/controllers/risktype"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type RiskTypeRoutes struct {
	logger             logger.Logger
	handler            lib.RequestHandler
	RiskTypeController controller.RiskTypeController
	authMiddleware     middlewares.JWTAuthMiddleware
}

func (s RiskTypeRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/risktype").Use(s.authMiddleware.Handler())
	{
		api.POST("/getAll", s.RiskTypeController.GetAll)
		api.POST("/getAllWithPaginate", s.RiskTypeController.GetAllWithPaginate)
		api.POST("/getOne", s.RiskTypeController.GetOne)
		api.POST("/store", s.RiskTypeController.Store)
		api.POST("/update", s.RiskTypeController.Update)
		api.POST("/delete", s.RiskTypeController.Delete)
	}
}

func NewRiskTypeRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	RiskTypeController controller.RiskTypeController,
	authMiddleware middlewares.JWTAuthMiddleware,
) RiskTypeRoutes {
	return RiskTypeRoutes{
		handler:            handler,
		logger:             logger,
		RiskTypeController: RiskTypeController,
		authMiddleware:     authMiddleware,
	}
}
