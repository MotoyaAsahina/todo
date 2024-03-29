package model

import (
	"context"

	"github.com/google/uuid"
)

func (repo *GormRepository) GetTags(ctx context.Context) ([]Tag, error) {
	var tags []Tag
	err := repo.db.GetDB(ctx).Order("name").Find(&tags).Error
	return tags, err
}

func (repo *GormRepository) PostTag(ctx context.Context, tag *Tag) error {
	return repo.db.GetDB(ctx).Create(tag).Error
}

func (repo *GormRepository) PutTag(ctx context.Context, tag *Tag) error {
	return repo.db.GetDB(ctx).Model(&Tag{ID: tag.ID}).Updates(map[string]interface{}{
		"name":  tag.Name,
		"color": tag.Color,
	}).Error
}

func (repo *GormRepository) GetTagMaps(ctx context.Context) (map[uuid.UUID][]uuid.UUID, error) {
	var tagMaps []TagMap
	err := repo.db.GetDB(ctx).
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

func (repo *GormRepository) GetTagMapsByTaskID(ctx context.Context, taskID uuid.UUID) ([]*TagMap, error) {
	var tagMaps []*TagMap
	err := repo.db.GetDB(ctx).Where("task_id = ?", taskID).Find(&tagMaps).Error
	if err != nil {
		return nil, err
	}

	return tagMaps, err
}

func (repo *GormRepository) GetTagNamesByTaskID(ctx context.Context, taskID uuid.UUID) ([]string, error) {
	var tagNames []string
	err := repo.db.GetDB(ctx).
		Model(&Tag{}).
		Select("tags.name as name").
		Joins("right join tag_maps on tag_maps.tag_id = tags.id").
		Where("tag_maps.task_id = ?", taskID).
		Order("tags.name").
		Pluck("tags.name", &tagNames).Error
	if err != nil {
		return nil, err
	}

	return tagNames, err
}

func (repo *GormRepository) PostTagMaps(ctx context.Context, tagMap []*TagMap) error {
	return repo.db.GetDB(ctx).Create(tagMap).Error
}

func (repo *GormRepository) DeleteTagMaps(ctx context.Context, tagMap []*TagMap) error {
	return repo.db.GetDB(ctx).Delete(tagMap).Error
}

func (repo *GormRepository) DeleteTagMapsByTaskID(ctx context.Context, taskID uuid.UUID) error {
	return repo.db.GetDB(ctx).Where("task_id = ?", taskID).Delete(TagMap{}).Error
}
