package routes

import (
	controllers "riskmanagement/controllers/common"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type CommonRoutes struct {
	logger           logger.Logger
	handler          lib.RequestHandler
	CommonController controllers.CommonController
	authMiddleware   middlewares.JWTAuthMiddleware
}

func (s CommonRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	// api := s.handler.Gin.Group("api/v1/common").Use(s.authMiddleware.Handler())
	api := s.handler.Gin.Group("api/v1/common")
	{
		api.POST("/pnnama", s.CommonController.GetNpNamaFilter)
		api.POST("/kanwil", s.CommonController.GetKanwilFilter)
		api.POST("/kanca", s.CommonController.GetKancaFilter)
		api.POST("/uker", s.CommonController.GetUkerFilter)
		api.POST("/riskevent/activity/product", s.CommonController.FilterRiskEventByActifityAndPoduct)
		api.POST("/riskindicator/riskevent", s.CommonController.FilterRiskIndokatorByRiskEventID)

		api.POST("/rrmhead", s.CommonController.GetRRMHeadFilter)
		api.POST("/pimpinanuker", s.CommonController.GetPimpinanUker)

		api.POST("/pn-approval", s.CommonController.GetApprovalResponse)

		// Enhance MQ
		api.POST("/mst-data", s.CommonController.GetMstDataOption)

		api.POST("/search-brc", s.CommonController.SearchBrc)
	}
}

func NewCommonRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	CommonController controllers.CommonController,
	authMiddleware middlewares.JWTAuthMiddleware,
) CommonRoutes {
	return CommonRoutes{
		logger:           logger,
		handler:          handler,
		CommonController: CommonController,
		authMiddleware:   authMiddleware,
	}
}
