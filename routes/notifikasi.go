package routes

import (
	controllers "riskmanagement/controllers/notifikasi"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type NotifikasiRoutes struct {
	logger               logger.Logger
	handler              lib.RequestHandler
	NotifikasiController controllers.NotifikasiController
	authMiddleware       middlewares.JWTAuthMiddleware
}

func (s NotifikasiRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("api/v1/notifikasi").Use(s.authMiddleware.Handler())
	// api := s.handler.Gin.Group("api/v1/notifikasi")
	{
		api.POST("/list", s.NotifikasiController.GetListNotifikasi)
		api.POST("/total", s.NotifikasiController.GetTotalNotifikasi)
		api.POST("/update/status", s.NotifikasiController.UpdateStatusNotifikasi)
		// api.DELETE("/delete/:id", s.NotifikasiController.DeleteStatusNotifikasi)
		api.POST("/create", s.NotifikasiController.CreateNotifikasi)
		api.POST("/delete", s.NotifikasiController.Delete)
	}
}

func NewNotifikasiRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	NotifikasiController controllers.NotifikasiController,
	authMiddleware middlewares.JWTAuthMiddleware,
) NotifikasiRoutes {
	return NotifikasiRoutes{
		logger:               logger,
		handler:              handler,
		NotifikasiController: NotifikasiController,
		authMiddleware:       authMiddleware,
	}
}
