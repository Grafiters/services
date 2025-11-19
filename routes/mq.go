package routes

import (
	controllers "riskmanagement/controllers/mq"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type MQRoutes struct {
	logger         logger.Logger
	handler        lib.RequestHandler
	MQController   controllers.MQController
	authMiddleware middlewares.JWTAuthMiddleware
}

func (s MQRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/mq").Use(s.authMiddleware.Handler())
	{
		api.POST("/menu/getAll", s.MQController.GetAllMenu)
		api.POST("/menu/getAllMstMenu", s.MQController.GetAllMstMenu)

		api.POST("/type/getTypeCode", s.MQController.GetTypeCode)
		api.POST("/type/getAll", s.MQController.GetAll)
		api.POST("/type/getAllSent", s.MQController.GetAllSent)
		api.POST("/type/filter", s.MQController.Filter)
		api.POST("/type/filterSent", s.MQController.FilterSent)
		api.POST("/type/getTypeList", s.MQController.GetTypeList)
		api.POST("/type/getTypeListAll", s.MQController.GetTypeListAll)
		api.POST("/type/getTypeListAktif", s.MQController.GetTypeListAktif)
		api.POST("/type/getOne", s.MQController.GetOne)
		api.POST("/type/store", s.MQController.Store)
		api.POST("/type/update", s.MQController.Update)
		api.POST("/type/delete", s.MQController.Delete)
		api.POST("/type/setStatusType", s.MQController.SetStatusType)
		api.POST("/type/getAllORDMember", s.MQController.GetAllORDMember)
		api.POST("/type/getDataQuestForm", s.MQController.GetDataQuestForm)
		api.POST("/type/sendQuest", s.MQController.SendQuest)
		api.POST("/type/approverUpdate", s.MQController.ApproverUpdate)
		api.POST("/type/generateQuestionnaireRequest", s.MQController.GenerateQuestionnaireRequest)
		api.POST("/type/generate-perpage", s.MQController.GenerateQuestPerPagePreview)
		api.POST("/type/weightTotal", s.MQController.WeightTotal)
		api.POST("/type/updatePartWeight", s.MQController.UpdatePartWeight)
		api.POST("/type/generate-perpage-approval", s.MQController.GenerateQuestPerPagePreviewApprover) // Perbaikan DAST

		api.POST("/part/getPartCode", s.MQController.GetPartCode)
		api.POST("/part/getAllPart", s.MQController.GetAllPart)
		api.POST("/part/filterPart", s.MQController.FilterPart)
		api.POST("/part/getOnePart", s.MQController.GetOnePart)
		api.POST("/part/storePart", s.MQController.StorePart)
		api.POST("/part/deletePart", s.MQController.DeletePart)
		api.POST("/part/updatePart", s.MQController.UpdatePart)

		api.POST("/subpart/store", s.MQController.StoreSubPart)
		api.POST("/subpart/update", s.MQController.UpdateSubPart)
		api.POST("/subpart/delete", s.MQController.DeleteSubPart)
		api.POST("/subpart/getSubPartList", s.MQController.GetSubPartList)
		api.POST("/subpart/getSubPartKode", s.MQController.GetSubPartKode)
		api.POST("/subpart/getDetailSubPart", s.MQController.GetDetailSubPart)

		api.POST("/questionnaire/store", s.MQController.StoreQuestion)
		api.POST("/questionnaire/update", s.MQController.UpdateQuestion)
		api.POST("/questionnaire/delete", s.MQController.DeleteQuestion)
		api.POST("/questionnaire/getQuestionnaireList", s.MQController.GetQuestionList)
		api.POST("/questionnaire/getKodeQuestionnaire", s.MQController.GetKodeQuestionnaire)
		api.POST("/questionnaire/getDetailQuestionnaire", s.MQController.GetDetailQuestionnaire)

		api.POST("/common/type", s.MQController.CommonGetType)
		api.POST("/common/part", s.MQController.CommonGetPart)
		api.POST("/common/sub-part", s.MQController.CommonGetSubPart)
		api.POST("/common/detail-type", s.MQController.GetDetailType)
		api.POST("/common/pn-approval", s.MQController.GetApprovalResponse)

		api.POST("/linkcage/storeLinkcage", s.MQController.StoreLinkcage)
		api.POST("/linkcage/getActive", s.MQController.GetActive)
		api.POST("/linkcage/getAll", s.MQController.GetAllLinkcage)
		api.POST("/linkcage/getOne", s.MQController.GetOneLinkcage)
		api.POST("/linkcage/setStatus", s.MQController.SetStatus)
		api.POST("/linkcage/updateLinkcage", s.MQController.UpdateLinkcage)
		api.POST("/linkcage/deleteLinkcage", s.MQController.DeleteLinkcage)

		api.POST("/response-user/list", s.MQController.GetResponseUserList)
		api.POST("/response-user/approval-list", s.MQController.GetResponseApprovalList)
		api.POST("/response-user/store", s.MQController.StoreResponseUser)
		api.POST("/response-user/generate", s.MQController.GenerateQuestWithAnswer)
		api.POST("/response-user/update", s.MQController.UpdateResponseUser)
		api.POST("/response-user/approve", s.MQController.ApproveResponseUser)
		api.POST("/response-user/reject", s.MQController.RejectResponseUser)
		api.POST("/response-user/reject-list", s.MQController.GetRejectResponse)
		api.POST("/response-user/nilai-akhir", s.MQController.GetNilaiAkhir)
		api.POST("/response-user/generate-perpage", s.MQController.GenerateQuestPerPage)
		api.POST("/response-user/pagination-process", s.MQController.ProcessGeneratePagination)
		api.POST("/response-user/cancel", s.MQController.CancelResponseUser)
		api.POST("/response-user/get-part-pagination", s.MQController.GetPartPagination)
		api.POST("/response-user/get-part-draft-pagination", s.MQController.GetPartPaginationDraft)
		api.POST("/response-user/disable-quest", s.MQController.DisableQuestionByPart)
		api.POST("/response-user/generate-perpage-approval", s.MQController.GenerateQuestForApprover) // Perbaikan DAST

		api.POST("/report/list", s.MQController.GetReportList)
		api.POST("/report/generate", s.MQController.GenerateReportPerPage)
		api.POST("/report/responden", s.MQController.GetNamaRespondenList)
		api.POST("/report/download", s.MQController.ResponseDownload)
		api.POST("/report/getSummary", s.MQController.GetSummary)
	}
}

func NewMQRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	mqController controllers.MQController,
	authMiddleware middlewares.JWTAuthMiddleware,
) MQRoutes {
	return MQRoutes{
		handler:        handler,
		logger:         logger,
		MQController:   mqController,
		authMiddleware: authMiddleware,
	}
}
