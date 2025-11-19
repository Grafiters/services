package routes

import (
	controllers "riskmanagement/controllers/subincident"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type SubIncidentRoutes struct {
	logger                logger.Logger
	handler               lib.RequestHandler
	SubIncidentController controllers.SubIncidentController
	authMiddleware        middlewares.JWTAuthMiddleware
}

func (s SubIncidentRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/subincident").Use(s.authMiddleware.Handler())
	{
		api.POST("/getAll", s.SubIncidentController.GetAll)
		api.POST("/getAllWithPaginate", s.SubIncidentController.GetAllWithPaginate)
		api.POST("/getOne", s.SubIncidentController.GetOne)
		api.POST("/getSubIncidentByID", s.SubIncidentController.GetSubIncidentByID)
		api.POST("/store", s.SubIncidentController.Store)
		api.POST("/update", s.SubIncidentController.Update)
		api.POST("/delete", s.SubIncidentController.Delete)
		api.POST("/getKode", s.SubIncidentController.GetKodePenyebabKejadian)
	}
}

func NewSubIncidentRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	SubIncidentController controllers.SubIncidentController,
	authMiddleware middlewares.JWTAuthMiddleware,
) SubIncidentRoutes {
	return SubIncidentRoutes{
		handler:               handler,
		logger:                logger,
		SubIncidentController: SubIncidentController,
		authMiddleware:        authMiddleware,
	}
}
