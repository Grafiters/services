package notificationsGlobal

import (
	"time"

	"github.com/google/uuid"
)

type NotificationRequest struct {
	IdUser string `json:"id_user" form:"id_user" example:"00000000"`
	After  string `json:"after" form:"after" example:"2023-10-01T12:00:00Z"`
	IsRead *bool  `json:"is_read" form:"is_read" example:"false"`
	Limit  int    `json:"limit" form:"limit" example:"10"`
	Offset int    `json:"offset" form:"offset" example:"0"`
}

type NotificationGlobal struct {
	ID        uuid.UUID `json:"id" form:"id" example:"b2c3d4e5-f6a7-8b9c-a0b1-c2d3e4f5g6h7"`
	IdUser    string    `json:"id_user" form:"id_user" example:"00000000"`
	Type      string    `json:"type" form:"type" example:"info"`
	Title     string    `json:"title" form:"title" example:"Notification Title"`
	Message   string    `json:"message" form:"message" example:"This is a notification message."`
	URL       string    `json:"url" form:"url" example:"https://example.com"`
	IsRead    bool      `json:"is_read" form:"is_read" example:"false"`
	CreatedAt time.Time `json:"created_at" form:"created_at" example:"2023-10-01T12:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at" example:"2023-10-01T12:00:00Z"`
}

type NotificationGlobalCreate struct {
	IdUser  string `json:"id_user" form:"id_user" example:"00000000"`
	Type    string `json:"type" form:"type" example:"info"`
	Title   string `json:"title" form:"title" example:"Notification Title"`
	Message string `json:"message" form:"message" example:"This is a notification message."`
	URL     string `json:"url" form:"url" example:"https://example.com"`
	IsRead  bool   `json:"is_read" form:"is_read" example:"false"`
}

type NotificationGlobalUpdate struct {
	ID     uuid.UUID `json:"id" form:"id" example:"b2c3d4e5-f6a7-8b9c-a0b1-c2d3e4f5g6h7"`
	IdUser string    `json:"id_user" form:"id_user" example:"00000000"`
	IsRead bool      `json:"is_read" form:"is_read" example:"true"`
}

