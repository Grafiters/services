package routes

import (
	controllers "riskmanagement/controllers/kriteriapinjaman"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type KriteriaPinjamanRoutes struct {
	logger           logger.Logger
	handler          lib.RequestHandler
	kriteriapinjaman controllers.KriteriaPinjamanController
	authMiddleware   middlewares.JWTAuthMiddleware
}

func NewKriteriaPinjamanRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	kriteriapinjaman controllers.KriteriaPinjamanController,
	authMiddleware middlewares.JWTAuthMiddleware,
) KriteriaPinjamanRoutes {
	return KriteriaPinjamanRoutes{
		logger:           logger,
		handler:          handler,
		kriteriapinjaman: kriteriapinjaman,
		authMiddleware:   authMiddleware,
	}
}

func (s KriteriaPinjamanRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/kriteriapinjaman").Use(s.authMiddleware.Handler())
	{
		api.POST("/getAll", s.kriteriapinjaman.GetAll)
	}
}
