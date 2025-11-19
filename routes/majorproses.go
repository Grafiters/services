package routes

import (
	controllers "riskmanagement/controllers/majorproses"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type MajorProsesRoutes struct {
	logger         logger.Logger
	handler        lib.RequestHandler
	MJPController  controllers.MajorProsesController
	authMiddleware middlewares.JWTAuthMiddleware
}

func (s MajorProsesRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/majorproses").Use(s.authMiddleware.Handler())
	{
		api.POST("/getAll", s.MJPController.GetAll)
		api.POST("/getAllWithPaginate", s.MJPController.GetAllWithPaginate)
		api.POST("/getOne", s.MJPController.GetOne)
		api.POST("/store", s.MJPController.Store)
		api.POST("/update", s.MJPController.Update)
		api.POST("/delete", s.MJPController.Delete)
		api.POST("/getKodeMajorProses", s.MJPController.GetKodeMajorProses)
		api.POST("/getMajorByMegaProses", s.MJPController.GetMajorByMegaProses)
	}
}

func NewMajorProsesRoutes(
	logger logger.Logger,
	handle lib.RequestHandler,
	MJPController controllers.MajorProsesController,
	authMiddleware middlewares.JWTAuthMiddleware,
) MajorProsesRoutes {
	return MajorProsesRoutes{
		logger:         logger,
		handler:        handle,
		MJPController:  MJPController,
		authMiddleware: authMiddleware,
	}
}
