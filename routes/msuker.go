package routes

import (
	controller "riskmanagement/controllers/msuker"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type MsUkerRoutes struct {
	logger           logger.Logger
	handler          lib.RequestHandler
	msUkerController controller.MsUkerController
	authMiddleware   middlewares.JWTAuthMiddleware
}

func (s MsUkerRoutes) Setup() {
	s.logger.Zap.Info("Setting Up Routes")
	api := s.handler.Gin.Group("/api/v1/msuker").Use(s.authMiddleware.Handler())
	{
		api.POST("/getAll", s.msUkerController.GetAll)
		// api.POST("/getUkerByBranch/:branchid", s.msUkerController.GetUkerByBranch)
		api.POST("/searchUker", s.msUkerController.SearchUker)
		api.POST("/getUkerByRegion", s.msUkerController.GetUkerPerRegion)
		api.POST("/searchPekerja", s.msUkerController.SearchPeserta)
		api.POST("/getPekerjaByBranch", s.msUkerController.GetPekerjaByBranch)
		api.POST("/getPekerjaByRegion", s.msUkerController.GetPekerjaByRegion)
		api.POST("/searchJabatan", s.msUkerController.SearchJabatan)
		api.POST("/searchUkerByRegionPekerja", s.msUkerController.SearchUkerByRegionPekerja)
		api.POST("/searchPekerjaPerUker", s.msUkerController.SearchPesertaPerUker)
		api.POST("/searchSigner", s.msUkerController.SearchSigner)
		api.POST("/searchRMC", s.msUkerController.SearchRMC)
		api.POST("/searchPelaku", s.msUkerController.SearchPelakuFraud)
		api.POST("/searchBrcUrcPerRegion", s.msUkerController.SearchBrcUrcPerRegion)
		api.POST("/listingJabatanPerUker", s.msUkerController.ListingJabatanPerUker)
		api.POST("/getPekerjaUkerByHilfm", s.msUkerController.GetPekerjaBranchByHILFM)
		api.POST("/searchRRMHead", s.msUkerController.SearchRRMHead)
		api.POST("/searchPekerjaORD", s.msUkerController.SearchPekerjaOrd)
	}
}

func NewMsUkerRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	msUkerController controller.MsUkerController,
	authMiddleware middlewares.JWTAuthMiddleware,
) MsUkerRoutes {
	return MsUkerRoutes{
		logger:           logger,
		handler:          handler,
		msUkerController: msUkerController,
		authMiddleware:   authMiddleware,
	}
}
