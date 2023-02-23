package model

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type NotificationTime struct {
	Time      time.Time    `json:"time" gorm:"type:timestamp;not null;primaryKey;default:0"`
	Noticed   bool         `json:"noticed" gorm:"type:boolean;not null;default:false"`
	NoticedAt sql.NullTime `json:"noticed_at" gorm:"type:timestamp null"`
	TaskCount int          `json:"task_count" gorm:"type:int;not null;default:0"`
}

type Notification struct {
	Time             time.Time        `json:"time" gorm:"type:timestamp;not null;primaryKey;default:0"`
	TaskID           uuid.UUID        `json:"task_id" gorm:"type:varchar(36);not null;primaryKey"`
	Tag              string           `json:"tag" gorm:"type:varchar(255);not null"`
	NotificationTime NotificationTime `gorm:"foreignKey:Time"`
	Task             Task             `gorm:"foreignKey:TaskID"`
}

type INotificationRepository interface {
	SetNotification(ctx context.Context, taskID uuid.UUID, notificationTime time.Time, notificationTag string) (alreadyTimeExists bool, err error)
	DeleteNotifications(ctx context.Context, taskID uuid.UUID) error
	GetLatestNotifications(ctx context.Context) ([]*Notification, error)
	GetValidNotificationTimes(ctx context.Context) ([]time.Time, error)
	SetNotificationTimeNoticed(ctx context.Context, t time.Time) error
}
