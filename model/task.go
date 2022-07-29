package model

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type Task struct {
	ID          uuid.UUID `json:"id" gorm:"type:varchar(36);not null;primaryKey"`
	GroupID     uuid.UUID `json:"group_id" gorm:"type:varchar(36);not null"`
	Title       string    `json:"title" gorm:"type:varchar(255);not null"`
	Description string    `json:"description" gorm:"type:text"`
	Done        bool      `json:"done" gorm:"type:boolean;not null"`
	DueDate     time.Time `json:"due_date" gorm:"type:timestamp"`
	CreatedAt   time.Time `json:"created_at" gorm:"type:timestamp;not null"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"type:timestamp;not null"`
	DoneAt      time.Time `json:"done_at" gorm:"type:timestamp"`
	Group       Group     `gorm:"foreignKey:GroupID"`
}

func GetTasks(ctx context.Context) ([]*Task, error) {
	var tasks []*Task
	err := GetDB(ctx).
		Model(&Task{}).
		Where("done = ?", false).
		Order("due_date, created_at").
		Find(&tasks).Error
	return tasks, err
}

func GetTask(ctx context.Context, id uuid.UUID) (*Task, error) {
	var task Task
	err := GetDB(ctx).Where("id = ?", id).First(&task).Error
	return &task, err
}

func PutTaskDone(ctx context.Context, id uuid.UUID) error {
	err := GetDB(ctx).Model(&Task{ID: id}).Updates(map[string]interface{}{
		"done":    true,
		"done_at": time.Now(),
	}).Error
	return err
}

func PutTask(ctx context.Context, task *Task) (*Task, error) {
	t := &Task{ID: task.ID}
	err := GetDB(ctx).Model(t).Updates(map[string]interface{}{
		"title":       task.Title,
		"description": task.Description,
		"due_date":    task.DueDate,
	}).Error
	return t, err
}

func PostTask(ctx context.Context, task *Task) (*Task, error) {
	err := GetDB(ctx).Create(task).Error
	if err != nil {
		return nil, err
	}
	return task, nil
}

func DeleteTask(ctx context.Context, id uuid.UUID) error {
	err := GetDB(ctx).Delete(&Task{ID: id}).Error
	return err
}
