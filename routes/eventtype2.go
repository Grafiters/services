package routes

import (
	controllers "riskmanagement/controllers/eventtypelv2"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type EventType2Routes struct {
	logger         logger.Logger
	handler        lib.RequestHandler
	ET2Controller  controllers.EventTypeLv2Controller
	authMiddleware middlewares.JWTAuthMiddleware
}

func (s EventType2Routes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/eventtypelv2").Use(s.authMiddleware.Handler())
	{
		api.POST("/getAll", s.ET2Controller.GetAll)
		api.POST("/getAllWithPaginate", s.ET2Controller.GetAllWithPaginate)
		api.POST("/getOne", s.ET2Controller.GetOne)
		api.POST("/store", s.ET2Controller.Store)
		api.POST("/update", s.ET2Controller.Update)
		api.POST("/delete", s.ET2Controller.Delete)
		api.POST("/getKodeEventType", s.ET2Controller.GetKodeEventType)
		api.POST("/getEventlv2", s.ET2Controller.GetEventTypeById1)
	}
}

func NewEventTypeLV2Routes(
	logger logger.Logger,
	handle lib.RequestHandler,
	ET2Controller controllers.EventTypeLv2Controller,
	authMiddleware middlewares.JWTAuthMiddleware,
) EventType2Routes {
	return EventType2Routes{
		logger:         logger,
		handler:        handle,
		ET2Controller:  ET2Controller,
		authMiddleware: authMiddleware,
	}
}
