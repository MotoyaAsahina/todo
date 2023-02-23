package model

import (
	"context"
	"time"

	"github.com/google/uuid"
)

func (repo *GormRepository) GetTasks(ctx context.Context) ([]*Task, error) {
	var tasks []*Task
	err := repo.db.GetDB(ctx).
		Model(&Task{}).
		Where("done = ?", false).
		Order("due_date, created_at").
		Find(&tasks).Error
	return tasks, err
}

func (repo *GormRepository) GetTask(ctx context.Context, id uuid.UUID) (*Task, error) {
	var task Task
	err := repo.db.GetDB(ctx).Where("id = ?", id).First(&task).Error
	return &task, err
}

func (repo *GormRepository) PutTaskDone(ctx context.Context, id uuid.UUID) error {
	err := repo.db.GetDB(ctx).Model(&Task{ID: id}).Updates(map[string]interface{}{
		"done":    true,
		"done_at": time.Now(),
	}).Error
	return err
}

func (repo *GormRepository) PutTask(ctx context.Context, task *Task) (*Task, error) {
	t := &Task{ID: task.ID}
	err := repo.db.GetDB(ctx).Model(t).Updates(map[string]interface{}{
		"title":       task.Title,
		"description": task.Description,
		"due_date":    task.DueDate,
	}).Error
	return t, err
}

func (repo *GormRepository) PostTask(ctx context.Context, task *Task) (*Task, error) {
	err := repo.db.GetDB(ctx).Create(task).Error
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (repo *GormRepository) DeleteTask(ctx context.Context, id uuid.UUID) error {
	err := repo.db.GetDB(ctx).Delete(&Task{ID: id}).Error
	return err
}
