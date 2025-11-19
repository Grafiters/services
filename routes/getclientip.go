package routes

import (
	controllers "riskmanagement/controllers/getclientip"
	"riskmanagement/lib"

	// "riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type ClientIPRoutes struct {
	logger             logger.Logger
	handler            lib.RequestHandler
	ClientIPController controllers.ClientIPController
	// authMiddleware middlewares.JWTAuthMiddleware
}

func (s ClientIPRoutes) Setup() {
	s.logger.Zap.Info("Setting Up routes")
	api := s.handler.Gin.Group("/api/v1/clientip") //.Use(s.authMiddleware.Handler())
	{
		api.GET("/getClientIP", s.ClientIPController.ClientIP)

	}
}

func NewClientIPRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	ClientIPController controllers.ClientIPController,
	// authMiddleware middlewares.JWTAuthMiddleware,
) ClientIPRoutes {
	return ClientIPRoutes{
		logger:             logger,
		handler:            handler,
		ClientIPController: ClientIPController,
		// authMiddleware: authMiddleware,
	}
}
