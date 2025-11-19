package routes

import (
	controllers "riskmanagement/controllers/eventtypelv1"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type EventType1Routes struct {
	logger         logger.Logger
	handler        lib.RequestHandler
	ET1Controller  controllers.EventTypeLv1Controller
	authMiddleware middlewares.JWTAuthMiddleware
}

func (s EventType1Routes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/eventtypelv1").Use(s.authMiddleware.Handler())
	{
		api.POST("/getAll", s.ET1Controller.GetAll)
		api.POST("/getAllWithPaginate", s.ET1Controller.GetAllWithPaginate)
		api.POST("/getOne", s.ET1Controller.GetOne)
		api.POST("/store", s.ET1Controller.Store)
		api.POST("/update", s.ET1Controller.Update)
		api.POST("/delete", s.ET1Controller.Delete)
		api.POST("/getKodeEventType", s.ET1Controller.GetKodeEventType)
	}
}

func NewEventTypeLV1Routes(
	logger logger.Logger,
	handle lib.RequestHandler,
	ET1Controller controllers.EventTypeLv1Controller,
	authMiddleware middlewares.JWTAuthMiddleware,
) EventType1Routes {
	return EventType1Routes{
		logger:         logger,
		handler:        handle,
		ET1Controller:  ET1Controller,
		authMiddleware: authMiddleware,
	}
}
