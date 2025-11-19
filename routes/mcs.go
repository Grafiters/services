package routes

import (
	controllers "riskmanagement/controllers/mcs"
	"riskmanagement/lib"

	"gitlab.com/golang-package-library/logger"
)

type McsRoutes struct {
	logger        logger.Logger
	handler       lib.RequestHandler
	McsController controllers.McsController
}

func (s McsRoutes) Setup() {
	s.logger.Zap.Info("Setting Up routes")
	api := s.handler.Gin.Group("/api/v1/mcs")
	{
		api.POST("/getUker", s.McsController.GetUker)
		api.POST("/getPIC", s.McsController.GetPIC)
		api.POST("/getJabatan", s.McsController.GetJabatan)
		api.POST("/getOrgeh", s.McsController.GetOrgeh)
	}
}

func NewMcsRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	McsController controllers.McsController,
) McsRoutes {
	return McsRoutes{
		logger:        logger,
		handler:       handler,
		McsController: McsController,
	}
}
