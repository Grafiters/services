package notificationglobal

import (
	"riskmanagement/lib"
	models "riskmanagement/models/notification_global"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type NotificationGlobalDefinition interface {
	WithTrx(trxHandle *gorm.DB) NotificationGlobalRepository
	PushNotification(request models.NotificationGlobalCreate) (response bool, err error)
	GetNotificationByPernr(request models.NotificationRequest) (response []models.NotificationGlobal, totalData int64, err error)
	ViewAllNotificationByPernr(request models.NotificationRequest) (response []models.NotificationGlobal, totalData int64, err error)
	UpdateReadNotification(request models.NotificationGlobalUpdate) (response bool, err error)
	MarkAllAsRead(request models.NotificationGlobalUpdate) (response bool, err error)
}

type NotificationGlobalRepository struct {
	db      lib.Database
	logger  logger.Logger
	timeout time.Duration
}

func NewNotificationGlobalRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) NotificationGlobalDefinition {
	return NotificationGlobalRepository{
		db:      db,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

func (notif NotificationGlobalRepository) PushNotification(request models.NotificationGlobalCreate) (response bool, err error) {
	err = notif.db.DB.Table("notifications_global").Create(&request).Error
	if err != nil {
		notif.logger.Zap.Error("Error creating notification: ", err)
		return false, err
	}

	return true, nil
}

func (notif NotificationGlobalRepository) GetNotificationByPernr(request models.NotificationRequest) (response []models.NotificationGlobal, totalData int64, err error) {
	db := notif.db.DB.Table("notifications_global")

	var sinceTime time.Time
	if request.After != "" {
		parsed, err := time.Parse(time.RFC3339, request.After)
		if err == nil {
			sinceTime = parsed
		}
	}

	query := db.Where("id_user = ?", request.IdUser).Order("created_at DESC")

	// Only filter by is_read if the request specifies it (e.g., pointer or has a flag)
	if request.IsRead != nil {
		query = query.Where("is_read = ?", *request.IsRead)
	}

	if !sinceTime.IsZero() {
		query = query.Where("created_at > ?", sinceTime)
	}

	if err = query.Count(&totalData).Error; err != nil {
		notif.logger.Zap.Error("Error counting records:", err)
		return
	}

	if request.Limit != 0 {
		query = query.Limit(int(request.Limit))
	}

	if request.Offset != 0 {
		query = query.Offset(int(request.Offset))
	}

	err = query.Scan(&response).Error

	if err != nil {
		notif.logger.Zap.Error("Error fetching notifications: ", err)
		return nil, totalData, err
	}

	return response, totalData, nil
}

// ViewAllNotificationByPernr implements NotificationGlobalDefinition.
func (notif NotificationGlobalRepository) ViewAllNotificationByPernr(request models.NotificationRequest) (response []models.NotificationGlobal, totalData int64, err error) {
	db := notif.db.DB.Table("notifications_global")

	query := db.Where("id_user = ?", request.IdUser).Order("created_at DESC")

	if err = query.Count(&totalData).Error; err != nil {
		notif.logger.Zap.Error("Error counting records:", err)
		return
	}

	if request.Limit != 0 {
		query = query.Limit(int(request.Limit))
	}

	if request.Offset != 0 {
		query = query.Offset(int(request.Offset))
	}

	err = query.Scan(&response).Error

	if err != nil {
		notif.logger.Zap.Error("Error fetching notifications: ", err)
		return nil, totalData, err
	}

	return response, totalData, nil
}

// UpdateReadNotification implements NotificationGlobalDefinition.
func (notif NotificationGlobalRepository) UpdateReadNotification(request models.NotificationGlobalUpdate) (response bool, err error) {
	err = notif.db.DB.Table("notifications_global").Save(&request).Where(`id = ?`, request.ID).Error
	if err != nil {
		notif.logger.Zap.Error("Error creating notification: ", err)
		return false, err
	}

	return true, nil
}

// MarkAllAsRead implements NotificationGlobalDefinition.
func (notif NotificationGlobalRepository) MarkAllAsRead(request models.NotificationGlobalUpdate) (response bool, err error) {
	err = notif.db.DB.Table("notifications_global").Save("is_read = true").Where(`id_user = ?`, request.IdUser).Error
	if err != nil {
		notif.logger.Zap.Error("Error creating notification: ", err)
		return false, err
	}

	return true, nil
}

// WithTrx implements NotificationGlobalDefinition.
func (notif NotificationGlobalRepository) WithTrx(trxHandle *gorm.DB) NotificationGlobalRepository {
	panic("unimplemented")
}
