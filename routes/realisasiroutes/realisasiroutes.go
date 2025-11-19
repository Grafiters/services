package realisasiroutes

import (
	controllers "riskmanagement/controllers/realisasicontroller"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type RealisasiRoutes struct {
	logger    logger.Logger
	handler   lib.RequestHandler
	realisasi controllers.RealisasiController
	auth      middlewares.JWTAuthMiddleware
}

func (s RealisasiRoutes) Setup() {
	s.logger.Zap.Info("Setting Up routes")
	api := s.handler.Gin.Group("/api/v1/realisasi").Use(s.auth.Handler())
	{
		api.POST("/getDataParameter", s.realisasi.GetDataParameter)
		api.POST("/storeDataParameter", s.realisasi.StoreDataParameter)
		api.POST("/getDataRevisiUker", s.realisasi.GetDataRevisiUker)
		api.POST("/storeDataRevisiUker", s.realisasi.StoreDataRevisiUker)
		api.POST("/deleteDataRevisiUker", s.realisasi.DeleteDataRevisiUker)
		api.POST("/getDataRealisasi", s.realisasi.GetDataRealisasi)
		api.POST("/updateFlag", s.realisasi.UpdateFlagVerifikasi)
	}

}

func NewRealisasiRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	realisasi controllers.RealisasiController,
	auth middlewares.JWTAuthMiddleware,
) RealisasiRoutes {
	return RealisasiRoutes{
		logger:    logger,
		handler:   handler,
		realisasi: realisasi,
		auth:      auth,
	}
}
