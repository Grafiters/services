package routes

import (
	controller "riskmanagement/controllers/verifikasirealisasi"
	"riskmanagement/lib"

	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type VerifikasiRealisasiRoutes struct {
	logger                   logger.Logger
	handler                  lib.RequestHandler
	verifRealisasiController controller.VerifikasiRealisasiController
	authMiddleware           middlewares.JWTAuthMiddleware
}

func (vrk VerifikasiRealisasiRoutes) Setup() {
	vrk.logger.Zap.Info("Setting Up Routes")
	api := vrk.handler.Gin.Group("/api/v1/verifikasirealisasi").Use(vrk.authMiddleware.Handler())
	{
		api.POST("/getNoPelaporan", vrk.verifRealisasiController.GetNoPelaporan)
		api.POST("/getData", vrk.verifRealisasiController.GetData)
		api.POST("/storeVerifikasi", vrk.verifRealisasiController.StoreVerifikasi)
		api.POST("/getOne", vrk.verifRealisasiController.GetOne)
		api.POST("/delete", vrk.verifRealisasiController.Delete)
		api.POST("/update", vrk.verifRealisasiController.Update)
	}
}

func NewVerifikasiRealisasiRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	verifRealisasiController controller.VerifikasiRealisasiController,
	authMiddleware middlewares.JWTAuthMiddleware,
) VerifikasiRealisasiRoutes {
	return VerifikasiRealisasiRoutes{
		logger:                   logger,
		handler:                  handler,
		verifRealisasiController: verifRealisasiController,
		authMiddleware:           authMiddleware,
	}
}
