package model

import (
	"context"
	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"time"
)

type NotificationTime struct {
	Time      time.Time `json:"time" gorm:"type:timestamp;not null;primaryKey;default:0"`
	Noticed   bool      `json:"noticed" gorm:"type:boolean;not null;default:false"`
	TaskCount int       `json:"task_count" gorm:"type:int;not null;default:0"`
}

type Notification struct {
	Time             time.Time        `json:"time" gorm:"type:timestamp;not null;primaryKey;default:0"`
	TaskID           uuid.UUID        `json:"task_id" gorm:"type:varchar(36);not null;primaryKey"`
	Tag              string           `json:"tag" gorm:"type:varchar(255);not null"`
	NotificationTime NotificationTime `gorm:"foreignKey:Time"`
	Task             Task             `gorm:"foreignKey:TaskID"`
}

func SetNotification(ctx context.Context, taskID uuid.UUID, notificationTime time.Time, notificationTag string) (alreadyTimeExists bool, err error) {
	normalTime := notificationTime.Add(-time.Duration(notificationTime.Second()) * time.Second)
	err = GetDB(ctx).Create(&NotificationTime{Time: normalTime}).Error

	if err == nil {
	} else if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
		alreadyTimeExists = true
	} else {
		return
	}

	err = GetDB(ctx).Create(&Notification{
		Time:   normalTime,
		TaskID: taskID,
		Tag:    notificationTag,
	}).Error
	if err != nil {
		return
	}

	return
}

func DeleteNotifications(ctx context.Context, taskID uuid.UUID) error {
	err := GetDB(ctx).Where("task_id = ?", taskID).Delete(&Notification{}).Error
	return err
}

func GetLatestNotifications(ctx context.Context) ([]*Notification, error) {
	var latestTime time.Time
	err := GetDB(ctx).
		Model(&NotificationTime{}).
		Where("noticed = false").Order("time asc").
		Limit(1).Select("time").Find(&latestTime).Error
	if err != nil {
		return nil, err
	}

	var notifications []*Notification
	err = GetDB(ctx).Where("time = ?", latestTime).Find(&notifications).Error
	if err != nil {
		return nil, err
	}

	// Set noticed and task_count
	err = GetDB(ctx).Model(&NotificationTime{Time: latestTime}).Updates(map[string]interface{}{
		"noticed":    true,
		"task_count": len(notifications),
	}).Error
	if err != nil {
		return nil, err
	}

	return notifications, nil
}
