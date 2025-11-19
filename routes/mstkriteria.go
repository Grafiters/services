package routes

import (
	controller "riskmanagement/controllers/mstkriteria"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type MstKriteriaRoutes struct {
	logger                logger.Logger
	handler               lib.RequestHandler
	mstKriteriaController controller.MstKriteriaController
	authMiddleware        middlewares.JWTAuthMiddleware
}

func (s MstKriteriaRoutes) Setup() {
	s.logger.Zap.Info("Setting Up Routes")
	api := s.handler.Gin.Group("/api/v1/mstkriteria").Use(s.authMiddleware.Handler())
	{
		api.POST("/getAll", s.mstKriteriaController.GetAll)
		api.POST("/getAllWithPaginate", s.mstKriteriaController.GetAllWithPaginate)
		api.POST("/getOne", s.mstKriteriaController.GetOne)
		api.POST("/store", s.mstKriteriaController.Store)
		api.POST("/update", s.mstKriteriaController.Update)
		api.POST("/getKodeCriteria", s.mstKriteriaController.GetKodeCriteria)
		api.POST("/getCriteriaById", s.mstKriteriaController.GetCriteriaById)
		api.POST("/getCriteriaByPeriode", s.mstKriteriaController.GetCriteriaByPeriode)
	}
}

func NewMstKriteriaRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	mstKriteriaController controller.MstKriteriaController,
	authMiddleware middlewares.JWTAuthMiddleware,
) MstKriteriaRoutes {
	return MstKriteriaRoutes{
		logger:                logger,
		handler:               handler,
		mstKriteriaController: mstKriteriaController,
		authMiddleware:        authMiddleware,
	}
}
