package model

import (
	"context"
)

type GroupRepository struct {
	db *DB
}

func (repo *GroupRepository) GetGroups(ctx context.Context) ([]*Group, error) {
	var groups []*Group
	err := repo.db.GetDB(ctx).Order("`order`").Find(&groups).Error
	return groups, err
}

func (repo *GroupRepository) PostGroup(ctx context.Context, group *Group) error {
	return repo.db.GetDB(ctx).Create(group).Error
}

func (repo *GroupRepository) PutGroup(ctx context.Context, group *Group) error {
	return repo.db.GetDB(ctx).Model(&Group{Id: group.Id}).Update("name", group.Name).Error
}

func (repo *GroupRepository) PutGroupOrder(ctx context.Context, group *Group) error {
	return repo.db.GetDB(ctx).Model(&Group{Id: group.Id}).Update("order", group.Order).Error
}
