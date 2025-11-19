package routes

import (
	controllers "riskmanagement/controllers/taskassignment"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type TaskAssignmentsRoutes struct {
	logger                    logger.Logger
	handler                   lib.RequestHandler
	TaskAssignmentsController controllers.TaskAssignmentsController
	authMiddleware            middlewares.JWTAuthMiddleware
}

func (s TaskAssignmentsRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/task-assignment")
	{
		// api.POST("/storeData", s.TaskAssignmentsController.StoreData)
		api.POST("/checkTableExist", s.TaskAssignmentsController.CheckTableExist)

		// add by Panji
		api.POST("/getTaskData", s.TaskAssignmentsController.GetDataTask)
		api.POST("/getTaskApprovalList", s.TaskAssignmentsController.GetTaskApprovalList)
		api.POST("/getTaskDataDatail", s.TaskAssignmentsController.GetDataTaskDetail)
		api.POST("/getRejectionNotes", s.TaskAssignmentsController.GetRejectionNotes)
		api.POST("/getDetailTematik", s.TaskAssignmentsController.GetDetailTematik)
		api.POST("/deleteLampiran", s.TaskAssignmentsController.DeleteLampiran)
		api.POST("/myTasklist", s.TaskAssignmentsController.MyTasklist)
		api.POST("/myTasklistDetail", s.TaskAssignmentsController.MyTasklistDetail)

		api.GET("/getAll", s.TaskAssignmentsController.GetAllTasklist)
		api.POST("/generateNoTask", s.TaskAssignmentsController.GenerateNoTask)
		api.POST("/validate", s.TaskAssignmentsController.ValidateData)
		api.POST("/store", s.TaskAssignmentsController.StoreTasklist)
		api.POST("/approval", s.TaskAssignmentsController.Approval)
		api.PUT("/update/:id", s.TaskAssignmentsController.UpdateTasklist)
		api.POST("/deleteTasklist", s.TaskAssignmentsController.DeleteTasklist)

		api.POST("/getMyTasklistTotal", s.TaskAssignmentsController.GetMyTasklistTotal)
	}
}

func NewTaskAssignmentsRouter(
	logger logger.Logger,
	handler lib.RequestHandler,
	TaskAssignmentsController controllers.TaskAssignmentsController,
	authMiddleware middlewares.JWTAuthMiddleware,
) TaskAssignmentsRoutes {
	return TaskAssignmentsRoutes{
		handler:                   handler,
		logger:                    logger,
		TaskAssignmentsController: TaskAssignmentsController,
		authMiddleware:            authMiddleware,
	}
}
