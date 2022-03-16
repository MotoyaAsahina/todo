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

func GetTags(ctx context.Context) ([]Tag, error) {
	var tags []Tag
	err := GetDB(ctx).Find(&tags).Error
	return tags, err
}

func PostTag(ctx context.Context, tag *Tag) error {
	return GetDB(ctx).Create(tag).Error
}

func UpdateTag(ctx context.Context, tag *Tag) error {
	return GetDB(ctx).Save(tag).Error
}

func GetTagMaps(ctx context.Context) (map[uuid.UUID][]uuid.UUID, error) {
	var tagMaps []TagMap
	err := GetDB(ctx).Find(&tagMaps).Error
	if err != nil {
		return nil, err
	}

	tagMap := make(map[uuid.UUID][]uuid.UUID)
	for _, t := range tagMaps {
		tagMap[t.TaskID] = append(tagMap[t.TaskID], t.TagID)
	}

	return tagMap, err
}

func PostTagMaps(ctx context.Context, tagMap []*TagMap) error {
	return GetDB(ctx).Create(tagMap).Error
}
