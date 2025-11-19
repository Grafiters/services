package routes

import (
	controllers "riskmanagement/controllers/laporan"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type LaporanRoutes struct {
	logger            logger.Logger
	handler           lib.RequestHandler
	LaporanController controllers.LaporanController
	authMiddleware    middlewares.JWTAuthMiddleware
}

func (s LaporanRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	//api := s.handler.Gin.Group("api/v1/laporan").Use(s.authMiddleware.Handler())
	api := s.handler.Gin.Group("api/v1/laporan")
	{
		api.POST("/historyTaskDataVerifikasi", s.LaporanController.GetLaporanHistoriTaskDataVerifikasi)
		api.POST("/historyTaskDataVerifikasi/", s.LaporanController.GetLaporanHistoriTaskDataVerifikasiDetail)
		api.POST("/perhitunganPersentasePenyelesaian", s.LaporanController.GetLaporanPerhitunganPersentasePenyelesaianBaru)
		api.POST("/perhitunganPersentasePenyelesaian/download", s.LaporanController.GetLaporanPerhitunganPersentasePenyelesaianBaruDownload)
		api.POST("/historyTaskDataVerifikasi/download", s.LaporanController.GetLaporanHistoriTaskDataVerifikasiDownload)
		api.GET("/historyTaskDataVerifikasi/riskEvent", s.LaporanController.GetRiskEventOnTaskList)
		api.POST("/getMonitoringJob", s.LaporanController.GetMonitoringJob)
		api.POST("/getNamaJob", s.LaporanController.GetNamaJob)
		api.POST("/getActivityDaily", s.LaporanController.GetActivityDaily)
		api.POST("/getActivityDailyDetail", s.LaporanController.GetActivityDailyDetail)
	}
}

func NewLaporanRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	LaporanController controllers.LaporanController,
	authMiddleware middlewares.JWTAuthMiddleware,
) LaporanRoutes {
	return LaporanRoutes{
		logger:            logger,
		handler:           handler,
		LaporanController: LaporanController,
		authMiddleware:    authMiddleware,
	}
}
