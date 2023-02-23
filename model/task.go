package model

import (
	"context"
	"time"

	"github.com/google/uuid"
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

type ITaskRepository interface {
	GetTasks(ctx context.Context) ([]*Task, error)
	GetTask(ctx context.Context, id uuid.UUID) (*Task, error)
	PutTaskDone(ctx context.Context, id uuid.UUID) error
	PutTask(ctx context.Context, task *Task) (*Task, error)
	PostTask(ctx context.Context, task *Task) (*Task, error)
	DeleteTask(ctx context.Context, id uuid.UUID) error
}
