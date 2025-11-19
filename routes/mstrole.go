package routes

import (
	controller "riskmanagement/controllers/mstrole"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type MstRoleRoutes struct {
	logger            logger.Logger
	handler           lib.RequestHandler
	mstRoleController controller.MstRoleController
	authMiddleware    middlewares.JWTAuthMiddleware
}

func (s MstRoleRoutes) Setup() {
	s.logger.Zap.Info("Setting Up Routes")
	api := s.handler.Gin.Group("/api/v1/mstrole").Use(s.authMiddleware.Handler())
	{
		api.POST("/getAll", s.mstRoleController.GetAll)
		api.POST("/getAllWithPaginate", s.mstRoleController.GetAllWithPaginate)
		api.POST("/getOne", s.mstRoleController.GetOne)
		api.POST("/store", s.mstRoleController.Store)
		api.POST("/update", s.mstRoleController.Update)
		api.POST("/delete", s.mstRoleController.Delete)
	}
}

func NewMstRoleRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	mstRoleController controller.MstRoleController,
	authMiddleware middlewares.JWTAuthMiddleware,
) MstRoleRoutes {
	return MstRoleRoutes{
		logger:            logger,
		handler:           handler,
		mstRoleController: mstRoleController,
		authMiddleware:    authMiddleware,
	}
}
