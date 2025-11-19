package routes

import (
	datatematik "riskmanagement/controllers/data_tematik"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type DataTematikRoutes struct {
	logger         logger.Logger
	handler        lib.RequestHandler
	datatematik    datatematik.DataTematikController
	authMiddleware middlewares.JWTAuthMiddleware
}

func NewDataTematikRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	datatematik datatematik.DataTematikController,
	authMiddleware middlewares.JWTAuthMiddleware,
) DataTematikRoutes {
	return DataTematikRoutes{
		logger:         logger,
		handler:        handler,
		datatematik:    datatematik,
		authMiddleware: authMiddleware,
	}
}

func (s DataTematikRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/datatematik").Use(s.authMiddleware.Handler())
	{
		api.POST("/getSampleData", s.datatematik.GetSampleDataTematik)
		api.POST("/updateStatusVerif", s.datatematik.UpdateStatusDataSample)
	}
}
