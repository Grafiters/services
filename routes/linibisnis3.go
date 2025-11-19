package routes

import (
	controllers "riskmanagement/controllers/linibisnislv3"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type LiniBisnis3Routes struct {
	logger         logger.Logger
	handler        lib.RequestHandler
	LB3Controller  controllers.LiniBisnisLv3Controller
	authMiddleware middlewares.JWTAuthMiddleware
}

func (s LiniBisnis3Routes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/linibisnislv3").Use(s.authMiddleware.Handler())
	{
		api.POST("/getAll", s.LB3Controller.GetAll)
		api.POST("/getAllWithPaginate", s.LB3Controller.GetAllWithPaginate)
		api.POST("/getOne", s.LB3Controller.GetOne)
		api.POST("/store", s.LB3Controller.Store)
		api.POST("/update", s.LB3Controller.Update)
		api.POST("/delete", s.LB3Controller.Delete)
		api.POST("/getKodeLiniBisnis", s.LB3Controller.GetKodeLiniBisnis)
		api.POST("/getLbById", s.LB3Controller.GetLBByID)
	}
}

func NewLiniBisnisLV3Routes(
	logger logger.Logger,
	handle lib.RequestHandler,
	LB3Controller controllers.LiniBisnisLv3Controller,
	authMiddleware middlewares.JWTAuthMiddleware,
) LiniBisnis3Routes {
	return LiniBisnis3Routes{
		logger:         logger,
		handler:        handle,
		LB3Controller:  LB3Controller,
		authMiddleware: authMiddleware,
	}
}
