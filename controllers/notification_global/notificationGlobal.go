package notificationglobal

import (
	"riskmanagement/lib"
	models "riskmanagement/models/notification_global"
	services "riskmanagement/services/notification_global"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type NotificaitonGlobalController struct {
	service services.NotificationGlobalDefinition
	logger  logger.Logger
}

func NewNotificationGlobalController(
	service services.NotificationGlobalDefinition,
	logger logger.Logger,
) NotificaitonGlobalController {
	return NotificaitonGlobalController{
		service: service,
		logger:  logger,
	}
}

func (n NotificaitonGlobalController) GetNotification(c *gin.Context) {
	request := models.NotificationRequest{}

	if err := c.Bind(&request); err != nil {
		n.logger.Zap.Error("Error binding request: ", err)
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	notification, totalData, err := n.service.GetNotificationByPernr(request)

	if err != nil {
		n.logger.Zap.Error("Error getting notification: ", err)
		lib.ReturnToJson(c, 200, "500", "Error getting notification", nil)
		return
	}

	if len(notification) == 0 {
		lib.ReturnToJson(c, 200, "404", "No notification found", nil)
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", notification, totalData)
}

func (n NotificaitonGlobalController) ViewAllNotificationByPernr(c *gin.Context) {
	request := models.NotificationRequest{}

	if err := c.Bind(&request); err != nil {
		n.logger.Zap.Error("Error binding request: ", err)
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	notification, totalData, err := n.service.ViewAllNotificationByPernr(request)

	if err != nil {
		n.logger.Zap.Error("Error getting notification: ", err)
		lib.ReturnToJson(c, 200, "500", "Error getting notification", nil)
		return
	}

	if len(notification) == 0 {
		lib.ReturnToJson(c, 200, "404", "No notification found", nil)
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", notification, totalData)
}

func (n NotificaitonGlobalController) PushNotification(c *gin.Context) {
	request := models.NotificationGlobalCreate{}

	if err := c.Bind(&request); err != nil {
		n.logger.Zap.Error("Error binding request: ", err)
		lib.ReturnToJson(c, 200, "400", "Invalid Request", err)
		return
	}

	status, err := n.service.PushNotification(request)

	if err != nil {
		n.logger.Zap.Error("Error getting notification: ", err)
		lib.ReturnToJson(c, 200, "500", "Error push notification", nil)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Push Notifikasi Berhasil", status)
}

func (n NotificaitonGlobalController) UpdateReadNotification(c *gin.Context) {
	request := models.NotificationGlobalUpdate{}

	if err := c.Bind(&request); err != nil {
		n.logger.Zap.Error("Error binding request: ", err)
		lib.ReturnToJson(c, 200, "400", "Invalid Request", err)
		return
	}

	status, err := n.service.UpdateReadNotification(request)

	if err != nil {
		n.logger.Zap.Error("Error getting notification: ", err)
		lib.ReturnToJson(c, 200, "500", "Error push notification", nil)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Update Notifikasi Berhasil", status)
}

func (n NotificaitonGlobalController) MarkAllAsRead(c *gin.Context) {
	request := models.NotificationGlobalUpdate{}

	if err := c.Bind(&request); err != nil {
		n.logger.Zap.Error("Error binding request: ", err)
		lib.ReturnToJson(c, 200, "400", "Invalid Request", err)
		return
	}

	status, err := n.service.MarkAllAsRead(request)

	if err != nil {
		n.logger.Zap.Error("Error getting notification: ", err)
		lib.ReturnToJson(c, 200, "500", "Error push notification", nil)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Update Notifikasi Berhasil", status)
}
