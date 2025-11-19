package routes

import (
	controllers "riskmanagement/controllers/eventtypelv3"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type EventType3Routes struct {
	logger         logger.Logger
	handler        lib.RequestHandler
	ET3Controller  controllers.EventTypeLv3Controller
	authMiddleware middlewares.JWTAuthMiddleware
}

func (s EventType3Routes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/eventtypelv3").Use(s.authMiddleware.Handler())
	{
		api.POST("/getAll", s.ET3Controller.GetAll)
		api.POST("/getAllWithPaginate", s.ET3Controller.GetAllWithPaginate)
		api.POST("/getOne", s.ET3Controller.GetOne)
		api.POST("/store", s.ET3Controller.Store)
		api.POST("/update", s.ET3Controller.Update)
		api.POST("/delete", s.ET3Controller.Delete)
		api.POST("/getKodeEventType", s.ET3Controller.GetKodeEventType)
		api.POST("/getEventlv3", s.ET3Controller.GetEventTypeById2)
	}
}

func NewEventTypeLV3Routes(
	logger logger.Logger,
	handle lib.RequestHandler,
	ET3Controller controllers.EventTypeLv3Controller,
	authMiddleware middlewares.JWTAuthMiddleware,
) EventType3Routes {
	return EventType3Routes{
		logger:         logger,
		handler:        handle,
		ET3Controller:  ET3Controller,
		authMiddleware: authMiddleware,
	}
}
