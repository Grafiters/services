package routes

import (
	controllers "riskmanagement/controllers/penyebabkejadianlv3"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type PenyebabKejadianLv3Routes struct {
	logger                        logger.Logger
	handler                       lib.RequestHandler
	PenyebabKejadianLv3Controller controllers.PenyebabKejadianLv3Controller
	authMiddleware                middlewares.JWTAuthMiddleware
}

func (s PenyebabKejadianLv3Routes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/penyebabkejadianlv3").Use(s.authMiddleware.Handler())
	{
		api.POST("/getAll", s.PenyebabKejadianLv3Controller.GetAll)
		api.POST("/getAllWithPaginate", s.PenyebabKejadianLv3Controller.GetAllWithPaginate)
		api.POST("/getOne", s.PenyebabKejadianLv3Controller.GetOne)
		api.POST("/store", s.PenyebabKejadianLv3Controller.Store)
		api.POST("/update", s.PenyebabKejadianLv3Controller.Update)
		api.POST("/delete", s.PenyebabKejadianLv3Controller.Delete)
		api.POST("/getKode", s.PenyebabKejadianLv3Controller.GetKodePenyebabKejadian)
		api.POST("/getKejadianByID", s.PenyebabKejadianLv3Controller.GetKejadianByIDlv2)
		api.POST("/getKejadianByIDlv1", s.PenyebabKejadianLv3Controller.GetKejadianByIDlv1)
		api.POST("/getSubKejadian", s.PenyebabKejadianLv3Controller.GetSubKejadian)
	}
}

func NewPenyebabKejadianLv3Routes(
	logger logger.Logger,
	handler lib.RequestHandler,
	PenyebabKejadianLv3Controller controllers.PenyebabKejadianLv3Controller,
	authMiddleware middlewares.JWTAuthMiddleware,

) PenyebabKejadianLv3Routes {
	return PenyebabKejadianLv3Routes{
		handler:                       handler,
		logger:                        logger,
		PenyebabKejadianLv3Controller: PenyebabKejadianLv3Controller,
		authMiddleware:                authMiddleware,
	}
}
