package routes

import (
	controllers "riskmanagement/controllers/activity"
	"riskmanagement/lib"

	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type ActivityRoutes struct {
	logger             logger.Logger
	handler            lib.RequestHandler
	ActivityController controllers.ActivityController
	authMiddleware     middlewares.JWTAuthMiddleware
}

func (s ActivityRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/activity").Use(s.authMiddleware.Handler())
	{
		api.POST("/getAll", s.ActivityController.GetAll)
		api.POST("/getAllWithPagination", s.ActivityController.GetAllWithPagination)
		api.POST("/getOne", s.ActivityController.GetOne)
		api.POST("/store", s.ActivityController.Store)
		api.POST("/update", s.ActivityController.Update)
		api.POST("/delete", s.ActivityController.Delete)
		api.POST("/getKodeActivity", s.ActivityController.GetKodeActivity)
	}
}

func NewActivityRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	ActivityController controllers.ActivityController,
	authMiddleware middlewares.JWTAuthMiddleware,
) ActivityRoutes {
	return ActivityRoutes{
		handler:            handler,
		logger:             logger,
		ActivityController: ActivityController,
		authMiddleware:     authMiddleware,
	}
}
