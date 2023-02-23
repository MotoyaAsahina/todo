package model

import (
	"context"
	"github.com/google/uuid"
)

type Tag struct {
	ID    uuid.UUID `json:"id" gorm:"type:varchar(36);not null;primaryKey"`
	Name  string    `json:"name" gorm:"type:varchar(255);not null"`
	Color string    `json:"color" gorm:"type:varchar(255);not null"`
}

type TagMap struct {
	ID     uuid.UUID `json:"id" gorm:"type:varchar(36);not null;primaryKey"`
	TaskID uuid.UUID `json:"task_id" gorm:"type:varchar(36);not null"`
	TagID  uuid.UUID `json:"tag_id" gorm:"type:varchar(36);not null"`
	Task   Task      `gorm:"foreignKey:TaskID"`
	Tag    Tag       `gorm:"foreignKey:TagID"`
}

type ITagRepository interface {
	GetTags(ctx context.Context) ([]Tag, error)
	PostTag(ctx context.Context, tag *Tag) error
	PutTag(ctx context.Context, tag *Tag) error
	GetTagMaps(ctx context.Context) (map[uuid.UUID][]uuid.UUID, error)
	GetTagMapsByTaskID(ctx context.Context, taskID uuid.UUID) ([]*TagMap, error)
	GetTagNamesByTaskID(ctx context.Context, taskID uuid.UUID) ([]string, error)
	PostTagMaps(ctx context.Context, tagMap []*TagMap) error
	DeleteTagMaps(ctx context.Context, tagMap []*TagMap) error
	DeleteTagMapsByTaskID(ctx context.Context, taskID uuid.UUID) error
}
