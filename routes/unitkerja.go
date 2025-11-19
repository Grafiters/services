package routes

import (
	controllers "riskmanagement/controllers/unitkerja"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type UnitKerjaRoutes struct {
	logger              logger.Logger
	handler             lib.RequestHandler
	UnitKerjaController controllers.UnitKerjaController
	authMiddleware      middlewares.JWTAuthMiddleware
}

func (s UnitKerjaRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/unitkerja").Use(s.authMiddleware.Handler())
	{
		api.GET("/getAll", s.UnitKerjaController.GetAll)
		api.GET("/getOne/:id", s.UnitKerjaController.GetOne)
		api.POST("/store", s.UnitKerjaController.Store)
		api.POST("/update", s.UnitKerjaController.Update)
		api.DELETE("/delete/:id", s.UnitKerjaController.Delete)
		api.POST("/getRegionList", s.UnitKerjaController.GetRegionList)
		api.POST("/getMainbrList", s.UnitKerjaController.GetMainbrList)
		api.POST("/getMainbrListKW", s.UnitKerjaController.GetMainbrListKW)
		api.POST("/getBranchList", s.UnitKerjaController.GetBranchList)
		api.POST("/getEmployeeRegion", s.UnitKerjaController.GetEmployeeRegion)

		// DisasterMaps
		api.POST("/getMapRegionList", s.UnitKerjaController.GetMapRegionList)
		api.POST("/getMapBranchList", s.UnitKerjaController.GetMapBranchList)
		api.POST("/getMapUnitList", s.UnitKerjaController.GetMapUnitList)
	}
}

func NewUnitKerjaRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	UnitKerjaController controllers.UnitKerjaController,
	authMiddleware middlewares.JWTAuthMiddleware,
) UnitKerjaRoutes {
	return UnitKerjaRoutes{
		handler:             handler,
		logger:              logger,
		UnitKerjaController: UnitKerjaController,
		authMiddleware:      authMiddleware,
	}
}
