package model

import (
	"context"
	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"time"
)

type NotificationRepository struct {
	db *DB
}

func (repo *NotificationRepository) SetNotification(ctx context.Context, taskID uuid.UUID, notificationTime time.Time, notificationTag string) (alreadyTimeExists bool, err error) {
	normalTime := notificationTime.Add(-time.Duration(notificationTime.Second()) * time.Second)
	err = repo.db.GetDB(ctx).Create(&NotificationTime{Time: normalTime}).Error

	if err == nil {
	} else if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
		alreadyTimeExists = true
	} else {
		return
	}

	err = repo.db.GetDB(ctx).Create(&Notification{
		Time:   normalTime,
		TaskID: taskID,
		Tag:    notificationTag,
	}).Error
	if err != nil {
		return
	}

	return
}

func (repo *NotificationRepository) DeleteNotifications(ctx context.Context, taskID uuid.UUID) error {
	err := repo.db.GetDB(ctx).Where("task_id = ?", taskID).Delete(&Notification{}).Error
	return err
}

func (repo *NotificationRepository) GetLatestNotifications(ctx context.Context) ([]*Notification, error) {
	var latestTime time.Time
	err := repo.db.GetDB(ctx).
		Model(&NotificationTime{}).
		Where("noticed = false").Order("time asc").
		Limit(1).Select("time").Find(&latestTime).Error
	if err != nil {
		return nil, err
	}

	var notifications []*Notification
	err = repo.db.GetDB(ctx).Where("time = ?", latestTime).Find(&notifications).Error
	if err != nil {
		return nil, err
	}

	// Set noticed and task_count
	err = repo.db.GetDB(ctx).Model(&NotificationTime{Time: latestTime}).Updates(map[string]interface{}{
		"noticed":    true,
		"noticed_at": time.Now(),
		"task_count": len(notifications),
	}).Error
	if err != nil {
		return nil, err
	}

	return notifications, nil
}

func (repo *NotificationRepository) GetValidNotificationTimes(ctx context.Context) ([]time.Time, error) {
	var notificationTimes []time.Time
	err := repo.db.GetDB(ctx).
		Model(&NotificationTime{}).
		Where("noticed = false").Order("time asc").
		Select("time").Find(&notificationTimes).Error
	if err != nil {
		return nil, err
	}

	return notificationTimes, nil
}

func (repo *NotificationRepository) SetNotificationTimeNoticed(ctx context.Context, t time.Time) error {
	err := repo.db.GetDB(ctx).Model(&NotificationTime{Time: t}).Updates(map[string]interface{}{
		"noticed": true,
	}).Error
	return err
}
