package routes

import (
	controllers "riskmanagement/controllers/materi"
	"riskmanagement/lib"

	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type MateriRoutes struct {
	logger           logger.Logger
	handler          lib.RequestHandler
	MateriController controllers.MateriController
	authMiddleware   middlewares.JWTAuthMiddleware
}

func (s MateriRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("api/v1/materi").Use(s.authMiddleware.Handler())
	{
		api.POST("/getAll", s.MateriController.GetAll)
		api.POST("/store", s.MateriController.Store)
		api.POST("/getMateriByActivityAndProduct", s.MateriController.GetMateriByActivityAndProduct)
		api.POST("/getVerifikasiMateri", s.MateriController.GetVerifikasiMateri)
	}
}

func NewMateriRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	MaterController controllers.MateriController,
	authMiddleware middlewares.JWTAuthMiddleware,
) MateriRoutes {
	return MateriRoutes{
		logger:           logger,
		handler:          handler,
		MateriController: MaterController,
		authMiddleware:   authMiddleware,
	}
}
