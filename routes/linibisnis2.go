package routes

import (
	controllers "riskmanagement/controllers/linibisnislv2"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type LiniBisnis2Routes struct {
	logger         logger.Logger
	handler        lib.RequestHandler
	LB2Controller  controllers.LiniBisnisLv2Controller
	authMiddleware middlewares.JWTAuthMiddleware
}

func (s LiniBisnis2Routes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/linibisnislv2").Use(s.authMiddleware.Handler())
	{
		api.POST("/getAll", s.LB2Controller.GetAll)
		api.POST("/getAllWithPaginate", s.LB2Controller.GetAllWithPaginate)
		api.POST("/getOne", s.LB2Controller.GetOne)
		api.POST("/store", s.LB2Controller.Store)
		api.POST("/update", s.LB2Controller.Update)
		api.POST("/delete", s.LB2Controller.Delete)
		api.POST("/getKodeLiniBisnis", s.LB2Controller.GetKodeLiniBisnis)
		api.POST("/getLbById", s.LB2Controller.GetLBByID)
	}
}

func NewLiniBisnisLV2Routes(
	logger logger.Logger,
	handle lib.RequestHandler,
	LB2Controller controllers.LiniBisnisLv2Controller,
	authMiddleware middlewares.JWTAuthMiddleware,
) LiniBisnis2Routes {
	return LiniBisnis2Routes{
		logger:         logger,
		handler:        handle,
		LB2Controller:  LB2Controller,
		authMiddleware: authMiddleware,
	}
}
