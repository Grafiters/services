package routes

import (
	controllers "riskmanagement/controllers/incident"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type IncidentRoutes struct {
	logger             logger.Logger
	handler            lib.RequestHandler
	IncidentController controllers.IncidentController
	authMiddleware     middlewares.JWTAuthMiddleware
}

func (s IncidentRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/incident").Use(s.authMiddleware.Handler())
	{
		api.POST("/getAll", s.IncidentController.GetAll)
		api.POST("/getAllWithPaginate", s.IncidentController.GetAllWithPaginate)
		api.POST("/getOne", s.IncidentController.GetOne)
		api.POST("/store", s.IncidentController.Store)
		api.POST("/update", s.IncidentController.Update)
		api.POST("/delete", s.IncidentController.Delete)
		api.POST("/getKode", s.IncidentController.GetKodePenyebabKejadian)
	}
}

func NewIncidentRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	IncidentController controllers.IncidentController,
	authMiddleware middlewares.JWTAuthMiddleware,
) IncidentRoutes {
	return IncidentRoutes{
		handler:            handler,
		logger:             logger,
		IncidentController: IncidentController,
		authMiddleware:     authMiddleware,
	}
}
