package routes

import (
	controllers "riskmanagement/controllers/aplikasi"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type AplikasiRoutes struct {
	logger             logger.Logger
	handler            lib.RequestHandler
	aplikasiController controllers.AplikasiController
	authMiddleware     middlewares.JWTAuthMiddleware
}

func (s AplikasiRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/aplikasi").Use(s.authMiddleware.Handler())
	{
		api.POST("/getAll", s.aplikasiController.GetAll)
		api.POST("/getOne/:id", s.aplikasiController.GetOne)
		api.POST("/store", s.aplikasiController.Store)
		api.POST("/update", s.aplikasiController.Update)
		api.POST("/delete/:id", s.aplikasiController.Delete)
	}
}

func NewaplikasiRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	aplikasiController controllers.AplikasiController,
	authMiddleware middlewares.JWTAuthMiddleware,
) AplikasiRoutes {
	return AplikasiRoutes{
		handler:            handler,
		logger:             logger,
		aplikasiController: aplikasiController,
		authMiddleware:     authMiddleware,
	}
}
