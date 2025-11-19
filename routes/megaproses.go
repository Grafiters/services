package routes

import (
	controllers "riskmanagement/controllers/megaproses"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type MegaProsesRoutes struct {
	logger         logger.Logger
	handler        lib.RequestHandler
	MPController   controllers.MegaProsesController
	authMiddleware middlewares.JWTAuthMiddleware
}

func (s MegaProsesRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/megaproses").Use(s.authMiddleware.Handler())
	{
		api.POST("/getAll", s.MPController.GetAll)
		api.POST("/getAllWithPaginate", s.MPController.GetAllWithPaginate)
		api.POST("/getOne", s.MPController.GetOne)
		api.POST("/store", s.MPController.Store)
		api.POST("/update", s.MPController.Update)
		api.POST("/delete", s.MPController.Delete)
		api.POST("/getKodeMegaProses", s.MPController.GetKodeMegaProses)
	}
}

func NewMegaProsesRoutes(
	logger logger.Logger,
	handle lib.RequestHandler,
	MPController controllers.MegaProsesController,
	authMiddleware middlewares.JWTAuthMiddleware,
) MegaProsesRoutes {
	return MegaProsesRoutes{
		logger:         logger,
		handler:        handle,
		MPController:   MPController,
		authMiddleware: authMiddleware,
	}
}
