package routes

import (
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	controllers "riskmanagement/controllers/notification_global"

	"gitlab.com/golang-package-library/logger"
)

type NotificationGlobalRoutes struct {
	logger                       logger.Logger
	handler                      lib.RequestHandler
	NotificationGlobalController controllers.NotificaitonGlobalController
	authMiddleware               middlewares.JWTAuthMiddleware
}

func (s NotificationGlobalRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/notification")
	{
		api.POST("/getNotification", s.NotificationGlobalController.GetNotification)
		api.POST("/viewAllNotification", s.NotificationGlobalController.ViewAllNotificationByPernr)
		api.POST("/pushNotification", s.NotificationGlobalController.PushNotification)
		api.POST("/updateReadNotification", s.NotificationGlobalController.UpdateReadNotification)
		api.POST("/markAllAsRead", s.NotificationGlobalController.MarkAllAsRead)
	}
}

func NewNotificationGlobalRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	NotificationGlobalController controllers.NotificaitonGlobalController,
	authMiddleware middlewares.JWTAuthMiddleware,
) NotificationGlobalRoutes {
	return NotificationGlobalRoutes{
		handler:                      handler,
		logger:                       logger,
		NotificationGlobalController: NotificationGlobalController,
		authMiddleware:               authMiddleware,
	}
}
