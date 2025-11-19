package routes

import (
	controller "riskmanagement/controllers/pgsuser"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type PgsUserRoutes struct {
	logger            logger.Logger
	handler           lib.RequestHandler
	PgsUserController controller.PgsUserController
	authMiddleware    middlewares.JWTAuthMiddleware
}

func (s PgsUserRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/pgsuser").Use(s.authMiddleware.Handler())
	{
		api.POST("/getAll", s.PgsUserController.GetAll)
		api.POST("/getAllWithPaginate", s.PgsUserController.GetAllWithPaginate)
		api.POST("/getOne", s.PgsUserController.GetOne)
		api.POST("/store", s.PgsUserController.Store)
		api.POST("/getPgsApproval", s.PgsUserController.GetPgsApproval)
		api.POST("/update", s.PgsUserController.Update)
		api.POST("/approve", s.PgsUserController.ApprovePgsUser)
		api.POST("/reject", s.PgsUserController.RejectPgsUser)
		// api.POST("/login", s.PgsUserController.Login)
		api.POST("/delete", s.PgsUserController.Delete)
		api.POST("/searchPekerjaByPN", s.PgsUserController.SearchPekerjaByPn)
	}
}

func NewPgsUserRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	pgsUserController controller.PgsUserController,
	authMiddleware middlewares.JWTAuthMiddleware,
) PgsUserRoutes {
	return PgsUserRoutes{
		logger:            logger,
		handler:           handler,
		PgsUserController: pgsUserController,
		authMiddleware:    authMiddleware,
	}
}
