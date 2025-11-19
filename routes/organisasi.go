package routes

import (
	"riskmanagement/controllers/organisasi"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type OrganisasiRoutes struct {
	logger         logger.Logger
	handler        lib.RequestHandler
	Controller     organisasi.OrganisasiController
	authMiddleware middlewares.JWTAuthMiddleware
}

func (r OrganisasiRoutes) Setup() {
	r.logger.Zap.Info("Setting up routes")
	api := r.handler.Gin.Group("/api/v1/organisasi").Use(r.authMiddleware.Handler())
	{
		api.POST("/getCostCenter", r.Controller.GetCostCenter)
		api.POST("/getOrgUnit", r.Controller.GetOrgUnit)
		api.POST("/getHilfm", r.Controller.GetHilfm)
	}
}

func NewOrganisasiRoute(
	logger logger.Logger,
	handler lib.RequestHandler,
	controller organisasi.OrganisasiController,
	authMiddleware middlewares.JWTAuthMiddleware,
) OrganisasiRoutes {
	return OrganisasiRoutes{
		logger:         logger,
		handler:        handler,
		Controller:     controller,
		authMiddleware: authMiddleware,
	}
}
