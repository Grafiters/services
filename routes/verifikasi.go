package routes

import (
	downloadControllers "riskmanagement/controllers/download"
	controllers "riskmanagement/controllers/verifikasi"
	"riskmanagement/lib"

	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type VerifikasiRoutes struct {
	logger               logger.Logger
	handler              lib.RequestHandler
	VerifikasiController controllers.VerifikasiController
	DownloadController   downloadControllers.DownloadController
	authMiddleware       middlewares.JWTAuthMiddleware
	audiTrail            middlewares.AuditTrailMiddleware
}

func (s VerifikasiRoutes) Setup() {
	s.logger.Zap.Info("Setting Up Routes")
	api := s.handler.Gin.Group("/api/v1/verifikasi").
		Use(s.authMiddleware.Handler()).
		Use(s.audiTrail.Handler())
	{
		api.GET("/getAll", s.VerifikasiController.GetAll)
		api.POST("/getListData", s.VerifikasiController.GetListData)
		api.POST("/getDataWithPagination", s.VerifikasiController.GetDataWithPagination)
		api.POST("/storeDraft", s.VerifikasiController.StoreDraft)
		api.POST("/getOne", s.VerifikasiController.GetOne)
		api.POST("/deleteLampiran", s.VerifikasiController.DeleteLampiranVerifikasi)
		api.POST("/deleteRiskControl", s.VerifikasiController.DeleteRiskControl)
		api.POST("/filterVerifikasi", s.VerifikasiController.FilterVerifikasi)
		api.POST("/getNoPelaporan", s.VerifikasiController.GetNoPelaporan)
		api.POST("/delete", s.VerifikasiController.Delete)
		api.POST("/konfirm", s.VerifikasiController.KonfirmSave)
		api.POST("/update", s.VerifikasiController.UpdateAllVerifikasi)
		api.POST("/filterReport", s.VerifikasiController.FilterReport)
		api.POST("/storeSimpan", s.VerifikasiController.StoreSimpan)
		api.POST("/verifikasiReportFilter", s.VerifikasiController.VerifikasiReportFilter)
		api.POST("/verifikasiReportFilterComplete", s.VerifikasiController.VerifikasiReportFilterComplete)
		api.POST("/verifikasiReportDetail", s.VerifikasiController.VerifikasiReportDetail)
		api.POST("/verifikasiReportWithWeaknessOnlyFilter", s.VerifikasiController.VerifikasiReportWithWeaknessOnlyFilter)
		api.POST("/verifikasiReportWithNonWeaknessOnlyFilter", s.VerifikasiController.VerifikasiReportWithNonWeaknessOnlyFilter)
		api.POST("/riskControlByVerificationId", s.VerifikasiController.RiskControlByVerificationId)
		api.POST("/getRiskIndicatorAsMateri", s.VerifikasiController.GetRiskIndicatorAsMateri)
		api.POST("/verificationReportByUkerFilter", s.VerifikasiController.VerificationReportByUkerFilter)
		api.POST("/verificationReportFilterByUkerComplete", s.VerifikasiController.VerificationReportFilterByUkerComplete)
		api.POST("/verifikasiReportByFraudIndicatorFilter", s.VerifikasiController.VerifikasiReportByFraudIndicatorFilter)
		api.POST("/verificationReportFilterByFraudIndicatorComplete", s.VerifikasiController.VerificationReportFilterByFraudIndicatorComplete)
		api.POST("/VerifikasiReportMateriList", s.VerifikasiController.VerifikasiReportMateriList)
		api.POST("/verificationReportUkerByAllActivity", s.VerifikasiController.VerificationReportUkerByAllActivity)
		api.POST("/verificationReportUkerByAllActivityComplete", s.VerifikasiController.VerificationReportUkerByAllActivityComplete)
		api.POST("/verificationReportUkerByAllActivityCompleteWithRiskIssue", s.VerifikasiController.VerificationReportUkerByAllActivityCompleteWithRiskIssue)
		api.POST("/generate", s.DownloadController.Generate)
		api.POST("/download/:id", s.DownloadController.DownloadHandler)
		api.POST("/VerifikasiReportList", s.VerifikasiController.VerifikasiReportList)
		api.POST("/rptRekapitulasiBCV", s.VerifikasiController.RptRakapitulasiBCV)
		api.POST("/rptRekomendasiRisk", s.VerifikasiController.RptRekomendasiRisk)
		api.POST("/validasiVerifikasi", s.VerifikasiController.ValidasiVerifikasi)
		api.POST("/acceptValidasi", s.VerifikasiController.AcceptValidasi)
		api.POST("/updateStatusVerifikasi", s.VerifikasiController.UpdateStatusVerifikasi)
		api.POST("/rejectValidasi", s.VerifikasiController.RejectValidasi)

		api.POST("/getRtlIndikasiFraud", s.VerifikasiController.GetRtlIndikasiFraud)

		// Versioning 1.0.0.1 by panji 31/08/2023
		api.POST("/deleteAnomaliByID", s.VerifikasiController.DeleteAnomaliByID)
		api.POST("/validasiVerifikasiDetail", s.VerifikasiController.ValidasiVerifikasiDetailData)

		// Batch 3 add by panji 02/04/2024
		api.POST("/getHistoryRTL", s.VerifikasiController.GetRekomendasiTindakLanjut)
		api.POST("/deleteKejadian", s.VerifikasiController.DeletePenyebabKejadian)
		api.POST("/reportsummary", s.VerifikasiController.VerifikasiSummaryRpt)
		api.POST("/verifikasifrekuensi", s.VerifikasiController.VerifikasiFrekuensiRpt)
	}
}

func NewVerifikasiRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	VerifikasiController controllers.VerifikasiController,
	DownloadController downloadControllers.DownloadController,
	authMiddleware middlewares.JWTAuthMiddleware,
	audiTrail middlewares.AuditTrailMiddleware,
) VerifikasiRoutes {
	return VerifikasiRoutes{
		logger:               logger,
		handler:              handler,
		VerifikasiController: VerifikasiController,
		DownloadController:   DownloadController,
		authMiddleware:       authMiddleware,
		audiTrail:            audiTrail,
	}
}
