package routes

import (
	controllers "riskmanagement/controllers/submajorproses"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type SubMajorProsesRoutes struct {
	logger          logger.Logger
	handler         lib.RequestHandler
	SUBMPController controllers.SubMajorProsesController
	authMiddleware  middlewares.JWTAuthMiddleware
}

func (s SubMajorProsesRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/submajorproses").Use(s.authMiddleware.Handler())
	{
		api.POST("/getAll", s.SUBMPController.GetAll)
		api.POST("/getAllWithPaginate", s.SUBMPController.GetAllWithPaginate)
		api.POST("/getOne", s.SUBMPController.GetOne)
		api.POST("/store", s.SUBMPController.Store)
		api.POST("/update", s.SUBMPController.Update)
		api.POST("/delete", s.SUBMPController.Delete)
		api.POST("/getDataByID", s.SUBMPController.GetDataByID)
	}
}

func NewSubMajorProsesRoutes(
	logger logger.Logger,
	handle lib.RequestHandler,
	SUBMPController controllers.SubMajorProsesController,
	authMiddleware middlewares.JWTAuthMiddleware,
) SubMajorProsesRoutes {
	return SubMajorProsesRoutes{
		logger:          logger,
		handler:         handle,
		SUBMPController: SUBMPController,
		authMiddleware:  authMiddleware,
	}
}
