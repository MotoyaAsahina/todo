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
	err := GetDB(ctx).Order("name").Find(&tags).Error
	return tags, err
}

func PostTag(ctx context.Context, tag *Tag) error {
	return GetDB(ctx).Create(tag).Error
}

func PutTag(ctx context.Context, tag *Tag) error {
	return GetDB(ctx).Model(&Tag{ID: tag.ID}).Updates(map[string]interface{}{
		"name":  tag.Name,
		"color": tag.Color,
	}).Error
}

func GetTagMaps(ctx context.Context) (map[uuid.UUID][]uuid.UUID, error) {
	var tagMaps []TagMap
	err := GetDB(ctx).
		Select("tag_maps.id as id, tag_maps.task_id as task_id, tag_maps.tag_id as tag_id").
		Joins("right join tags on tag_maps.tag_id = tags.id").
		Order("tags.name").
		Find(&tagMaps).Error
	if err != nil {
		return nil, err
	}

	tagMap := make(map[uuid.UUID][]uuid.UUID)
	for _, t := range tagMaps {
		tagMap[t.TaskID] = append(tagMap[t.TaskID], t.TagID)
	}

	return tagMap, err
}

func GetTagMapsByTaskID(ctx context.Context, taskID uuid.UUID) ([]*TagMap, error) {
	var tagMaps []*TagMap
	err := GetDB(ctx).Where("task_id = ?", taskID).Find(&tagMaps).Error
	if err != nil {
		return nil, err
	}

	return tagMaps, err
}

func PostTagMaps(ctx context.Context, tagMap []*TagMap) error {
	return GetDB(ctx).Create(tagMap).Error
}

func DeleteTagMaps(ctx context.Context, tagMap []*TagMap) error {
	return GetDB(ctx).Delete(tagMap).Error
}

func DeleteTagMapsByTaskID(ctx context.Context, taskID uuid.UUID) error {
	return GetDB(ctx).Where("task_id = ?", taskID).Delete(TagMap{}).Error
}
