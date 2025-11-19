package routes

import (
	pekerja "riskmanagement/controllers/pekerja"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type PekerjaRoutes struct {
	logger            logger.Logger
	handler           lib.RequestHandler
	PekerjaController pekerja.PekerjaController
	authMiddleware    middlewares.JWTAuthMiddleware
}

func (s PekerjaRoutes) Setup() {
	s.logger.Zap.Info("Setting Up Routes")
	apiPekerja := s.handler.Gin.Group("/api/v1/pekerja").Use(s.authMiddleware.Handler())
	{
		apiPekerja.POST("/getAllPekerjaBranch", s.PekerjaController.GetAllPekerjaBranch)
		apiPekerja.POST("/getApproval", s.PekerjaController.GetApproval)
	}
}

func NewPekerjaRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	PekerjaController pekerja.PekerjaController,
	authMiddleware middlewares.JWTAuthMiddleware,
) PekerjaRoutes {
	return PekerjaRoutes{
		logger:            logger,
		handler:           handler,
		PekerjaController: PekerjaController,
		authMiddleware:    authMiddleware,
	}
}
