package routes

import (
	"riskmanagement/lib"

	controllers "riskmanagement/controllers/menu"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type MenuRoutes struct {
	logger         logger.Logger
	handler        lib.RequestHandler
	MenuController controllers.MenuController
	authMiddleware middlewares.JWTAuthMiddleware
}

func (s MenuRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("api/v1/menu")
	{
		api.POST("/getMenu", s.MenuController.GetMenuTree)
		api.POST("/getKuisioner", s.MenuController.GetKuisioner)

		api.POST("/getAll", s.MenuController.GetAll)
		api.POST("/getAllMstMenu", s.MenuController.GetAllMstMenu)

		api.POST("/deleteMenuRRM", s.MenuController.DeleteMenuRRM)
		api.POST("/storeMstRRM", s.MenuController.StoreMstRRM)
		api.POST("/deleteRole", s.MenuController.DeleteRole)
		api.POST("/storeRoleRRM", s.MenuController.StoreRoleRRM)
		api.POST("/setStatus", s.MenuController.SetStatus)

		api.POST("/getLastID", s.MenuController.GetLastID)
	}
}

func NewMenuRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	MenuController controllers.MenuController,
	authMiddleware middlewares.JWTAuthMiddleware,
) MenuRoutes {
	return MenuRoutes{
		logger:         logger,
		handler:        handler,
		MenuController: MenuController,
		authMiddleware: authMiddleware,
	}
}
