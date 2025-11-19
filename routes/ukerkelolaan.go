package routes

import (
	controller "riskmanagement/controllers/ukerkelolaan"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type UkerKelolaanRoutes struct {
	logger         logger.Logger
	handler        lib.RequestHandler
	UKController   controller.UkerKolalaanController
	authMiddleware middlewares.JWTAuthMiddleware
}

func (s UkerKelolaanRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/ukerkelolaan").Use(s.authMiddleware.Handler())
	{
		api.POST("/getAllWithPaginate", s.UKController.GetAllWithPaginate)
		api.POST("/filterUkerKelolaan", s.UKController.FilterUkerKelolaan)
		api.POST("/store", s.UKController.Store)
		api.POST("/getOne", s.UKController.GetOne)
		api.POST("/update", s.UKController.Update)
		api.POST("/delete", s.UKController.Delete)
		api.POST("/getListUkerKelolaan", s.UKController.GetListUkerKelolaan)
	}
}

func NewUkerKelolaanRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	UKConttroller controller.UkerKolalaanController,
	authMiddleware middlewares.JWTAuthMiddleware,
) UkerKelolaanRoutes {
	return UkerKelolaanRoutes{
		logger:         logger,
		handler:        handler,
		UKController:   UKConttroller,
		authMiddleware: authMiddleware,
	}
}
