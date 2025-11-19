package routes

import (
	controllers "riskmanagement/controllers/admin_setting"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type AdminSettingRoutes struct {
	logger                 logger.Logger
	handler                lib.RequestHandler
	AdminSettingController controllers.AdminSettingController
	authMiddleware         middlewares.JWTAuthMiddleware
}

func (s AdminSettingRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/admin_setting").Use(s.authMiddleware.Handler())
	{
		api.POST("/show", s.AdminSettingController.Show)
		api.POST("/getAll", s.AdminSettingController.GetAll)
		api.POST("/store", s.AdminSettingController.Store)
		api.POST("/update", s.AdminSettingController.Update)
		api.POST("/delete", s.AdminSettingController.Delete)
		api.POST("/searchTaskType", s.AdminSettingController.SearchTaskType)
		api.POST("/searchTaskTypeInput", s.AdminSettingController.SearchTaskTypeInput)
		api.POST("/searchTaskTypeInputByKegiatan", s.AdminSettingController.SearchTaskTypeInputByKegiatan)
		api.POST("/getOne", s.AdminSettingController.GetOne)
	}
}

func NewAdminSettingRouter(
	logger logger.Logger,
	handler lib.RequestHandler,
	AdminSettingController controllers.AdminSettingController,
	authMiddleware middlewares.JWTAuthMiddleware,
) AdminSettingRoutes {
	return AdminSettingRoutes{
		handler:                handler,
		logger:                 logger,
		AdminSettingController: AdminSettingController,
		authMiddleware:         authMiddleware,
	}
}
