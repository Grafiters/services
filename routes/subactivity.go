package routes

import (
	controllers "riskmanagement/controllers/subactivity"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type SubActivityRoutes struct {
	logger                logger.Logger
	handler               lib.RequestHandler
	SubActivityController controllers.SubActivityController
	authMiddleware        middlewares.JWTAuthMiddleware
}

func (s SubActivityRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/subactivity").Use(s.authMiddleware.Handler())
	{
		api.POST("/getAll", s.SubActivityController.GetAll)
		api.POST("/getAllWithPagination", s.SubActivityController.GetAllWithPagination)
		api.POST("/getOne", s.SubActivityController.GetOne)
		api.POST("/getLastID", s.SubActivityController.GetLastID)
		api.POST("/getSubactivity", s.SubActivityController.GetSubactivity)
		api.POST("/store", s.SubActivityController.Store)
		api.POST("/update", s.SubActivityController.Update)
		api.POST("/delete", s.SubActivityController.Delete)
		api.POST("/getKodeSub", s.SubActivityController.GetKodeSubActivity)
	}
}

func NewSubActivityRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	SubActivityController controllers.SubActivityController,
	authMiddleware middlewares.JWTAuthMiddleware,
) SubActivityRoutes {
	return SubActivityRoutes{
		handler:               handler,
		logger:                logger,
		SubActivityController: SubActivityController,
		authMiddleware:        authMiddleware,
	}
}
