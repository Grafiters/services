package routes

import (
	controllers "riskmanagement/controllers/krid"
	"riskmanagement/lib"

	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type KridRoutes struct {
	logger         logger.Logger
	handler        lib.RequestHandler
	KridController controllers.KridController
	authMiddleware middlewares.JWTAuthMiddleware
}

func (s KridRoutes) Setup() {
	s.logger.Zap.Info("Setting Up routes")
	api := s.handler.Gin.Group("/api/v1/krid").Use(s.authMiddleware.Handler())
	{
		api.POST("/getDetailIndikator", s.KridController.GetDetailIndikator)
		api.POST("/getAllParameterIndikator", s.KridController.GetAllParameterIndikator)
		api.POST("/searchIndikator", s.KridController.SearchIndicator)
		api.POST("/searchIndikatorEdit", s.KridController.SearchIndicatorEdit)
	}
}

func NewKridRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	KridController controllers.KridController,
	authMiddleware middlewares.JWTAuthMiddleware,
) KridRoutes {
	return KridRoutes{
		logger:         logger,
		handler:        handler,
		KridController: KridController,
		authMiddleware: authMiddleware,
	}
}
