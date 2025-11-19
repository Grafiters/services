package routes

import (
	controllers "riskmanagement/controllers/tasklists"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type TasklistsRoutes struct {
	logger              logger.Logger
	handler             lib.RequestHandler
	TasklistsController controllers.TasklistsController
	authMiddleware      middlewares.JWTAuthMiddleware
}

func (s TasklistsRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/tasklists")
	{
		api.POST("/getAll", s.TasklistsController.GetAll)
		api.POST("/getTaskByID", s.TasklistsController.GetTaskByID)
		api.POST("/getOne", s.TasklistsController.GetOne)
		api.POST("/getAllOfficer", s.TasklistsController.GetAllOfficer)
		api.POST("/filter", s.TasklistsController.Filter)
		api.POST("/filterByID", s.TasklistsController.FilterByID)
		api.POST("/filterOfficer", s.TasklistsController.FilterOfficer)
		api.POST("/store", s.TasklistsController.Store)
		api.POST("/update", s.TasklistsController.Update)
		api.POST("/updateEndDate", s.TasklistsController.UpdateEndDate)
		api.POST("/delete", s.TasklistsController.Delete)
		api.POST("/approval", s.TasklistsController.Approval)
		api.POST("/validation", s.TasklistsController.Validation)
		api.POST("/countTask", s.TasklistsController.CountTask)
		api.POST("/checkAvailability", s.TasklistsController.CheckAvailability)
		api.POST("/storeDone", s.TasklistsController.StoreDone)
		api.GET("/getDone", s.TasklistsController.GetDone)
		api.POST("/leadAutocompleteVal", s.TasklistsController.LeadAutocompleteVal)
		api.POST("/leadAutocompleteApr", s.TasklistsController.LeadAutocompleteApr)
		api.POST("/userRegion", s.TasklistsController.UserRegion)
		api.POST("/getDataAnomali", s.TasklistsController.GetDataAnomali)
		api.POST("/getDataVerifikasi", s.TasklistsController.GetDataVerifikasi)
		api.POST("/getLampiranIndicator", s.TasklistsController.GetLampiranIndikator)
		api.POST("/downloadLampiranIndikator", s.TasklistsController.DownloadLampiranIndikatorTemplate)
		api.POST("/insertTasklistRejected", s.TasklistsController.InsertTasklistRejected)
		api.POST("/getTasklistData", s.TasklistsController.GetTasklistData)
		api.POST("/storeDaily", s.TasklistsController.StoreDaily)
		api.POST("/updateDaily", s.TasklistsController.UpdateDaily)
		api.POST("/anomaliHeader", s.TasklistsController.GetAnomaliHeader)
		api.POST("/anomaliValue", s.TasklistsController.GetAnomaliValue)
		api.POST("/getFirstLampiran", s.TasklistsController.GetFirstLampiran)
	}
}

func NewTasklistsRouter(
	logger logger.Logger,
	handler lib.RequestHandler,
	TasklistsController controllers.TasklistsController,
	authMiddleware middlewares.JWTAuthMiddleware,
) TasklistsRoutes {
	return TasklistsRoutes{
		handler:             handler,
		logger:              logger,
		TasklistsController: TasklistsController,
		authMiddleware:      authMiddleware,
	}
}
