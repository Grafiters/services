package notificationglobal

import (
	models "riskmanagement/models/notification_global"
	repository "riskmanagement/repository/notification_global"

	"gitlab.com/golang-package-library/logger"
)

type NotificationGlobalDefinition interface {
	PushNotification(request models.NotificationGlobalCreate) (response bool, err error)
	GetNotificationByPernr(request models.NotificationRequest) (response []models.NotificationGlobal, totalData int64, err error)
	ViewAllNotificationByPernr(request models.NotificationRequest) (response []models.NotificationGlobal, totalData int64, err error)
	UpdateReadNotification(request models.NotificationGlobalUpdate) (response bool, err error)
	MarkAllAsRead(request models.NotificationGlobalUpdate) (response bool, err error)
}

type NotificationGlobalService struct {
	repository repository.NotificationGlobalDefinition
	logger     logger.Logger
}

func NewNotificationGlobalService(
	repository repository.NotificationGlobalDefinition,
	logger logger.Logger,
) NotificationGlobalDefinition {
	return NotificationGlobalService{
		repository: repository,
		logger:     logger,
	}
}

// GetNotificationByPernr implements NotificationGlobalDefinition.
func (n NotificationGlobalService) GetNotificationByPernr(request models.NotificationRequest) (response []models.NotificationGlobal, totalData int64, err error) {
	notification, totalData, err := n.repository.GetNotificationByPernr(request)

	if err != nil {
		n.logger.Zap.Error("Error getting notification: ", err)
		return nil, 0, err
	}

	return notification, totalData, nil
}

// ViewAllNotificationByPernr implements NotificationGlobalDefinition.
func (n NotificationGlobalService) ViewAllNotificationByPernr(request models.NotificationRequest) (response []models.NotificationGlobal, totalData int64, err error) {
	notification, totalData, err := n.repository.ViewAllNotificationByPernr(request)

	if err != nil {
		n.logger.Zap.Error("Error getting notification: ", err)
		return nil, 0, err
	}

	return notification, totalData, nil
}

// PushNotification implements NotificationGlobalDefinition.
func (n NotificationGlobalService) PushNotification(request models.NotificationGlobalCreate) (response bool, err error) {
	response, err = n.repository.PushNotification(request)
	if err != nil {
		n.logger.Zap.Error("Error pushing notification: ", err)
		return response, err
	}

	return response, nil
}

// UpdateReadNotification implements NotificationGlobalDefinition.
func (n NotificationGlobalService) UpdateReadNotification(request models.NotificationGlobalUpdate) (response bool, err error) {
	response, err = n.repository.UpdateReadNotification(request)
	if err != nil {
		n.logger.Zap.Error("Error updating notification: ", err)
		return response, err
	}

	return response, nil
}

// MarkAllAsRead implements NotificationGlobalDefinition.
func (n NotificationGlobalService) MarkAllAsRead(request models.NotificationGlobalUpdate) (response bool, err error) {
	response, err = n.repository.MarkAllAsRead(request)
	if err != nil {
		n.logger.Zap.Error("Error updating notification: ", err)
		return response, err
	}

	return response, nil
}
