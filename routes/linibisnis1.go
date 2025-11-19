package routes

import (
	controllers "riskmanagement/controllers/linibisnislv1"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type LiniBisnis1Routes struct {
	logger         logger.Logger
	handler        lib.RequestHandler
	LB1Controller  controllers.LiniBisnisLv1Controller
	authMiddleware middlewares.JWTAuthMiddleware
}

func (s LiniBisnis1Routes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/linibisnislv1").Use(s.authMiddleware.Handler())
	{
		api.POST("/getAll", s.LB1Controller.GetAll)
		api.POST("/getAllWithPaginate", s.LB1Controller.GetAllWithPaginate)
		api.POST("/getOne", s.LB1Controller.GetOne)
		api.POST("/store", s.LB1Controller.Store)
		api.POST("/update", s.LB1Controller.Update)
		api.POST("/delete", s.LB1Controller.Delete)
		api.POST("/getKodeLiniBisnis", s.LB1Controller.GetKodeLiniBisnis)
	}
}

func NewLiniBisnisLV1Routes(
	logger logger.Logger,
	handle lib.RequestHandler,
	LB1Controller controllers.LiniBisnisLv1Controller,
	authMiddleware middlewares.JWTAuthMiddleware,
) LiniBisnis1Routes {
	return LiniBisnis1Routes{
		logger:         logger,
		handler:        handle,
		LB1Controller:  LB1Controller,
		authMiddleware: authMiddleware,
	}
}
