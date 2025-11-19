package routes

import (
	controllers "riskmanagement/controllers/managementuser"
	"riskmanagement/lib"

	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type ManagementUserRoutes struct {
	logger         logger.Logger
	handler        lib.RequestHandler
	MUController   controllers.ManagementUserController
	authMiddleware middlewares.JWTAuthMiddleware
}

func (s ManagementUserRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/managementuser").Use(s.authMiddleware.Handler())
	{
		api.POST("/getAll", s.MUController.GetAll)
		api.POST("/getAllWithPaginate", s.MUController.GetAllWithPaginate)
		api.POST("/getOne", s.MUController.GetOne)
		api.POST("/store", s.MUController.Store)
		api.POST("/update", s.MUController.Update)
		api.POST("/delete", s.MUController.Delete)
		api.POST("/getMappingMenu", s.MUController.GetMappingMenu)
		api.POST("/mappingMenu", s.MUController.MappingMenu)
		api.POST("/getAllMenu", s.MUController.GetAllMenu)
		api.POST("/deleteMapMenu", s.MUController.DeleteMapControl)
		// api.POST("/getMenu", s.MUController.GetMenu)
		api.POST("/getUkerKelolaan", s.MUController.GetUkerKelolaan)
		api.POST("/getTreeMenu", s.MUController.GetTreeMenu)
		api.POST("/getLevelUker", s.MUController.GetLevelUker)
		api.POST("/getJabatanRole", s.MUController.GetJabatanRole)

		// Batch3
		api.POST("/getAdditionalMenu", s.MUController.GetAdditionalMenu)
	}
}

func NewManagementUserRoutes(
	logger logger.Logger,
	handle lib.RequestHandler,
	MUController controllers.ManagementUserController,
	authMiddleware middlewares.JWTAuthMiddleware,
) ManagementUserRoutes {
	return ManagementUserRoutes{
		logger:         logger,
		handler:        handle,
		MUController:   MUController,
		authMiddleware: authMiddleware,
	}
}
