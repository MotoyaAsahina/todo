package model

import (
	"context"
)

func (repo *GormRepository) GetGroups(ctx context.Context) ([]*Group, error) {
	var groups []*Group
	err := repo.db.GetDB(ctx).Order("`order`").Find(&groups).Error
	return groups, err
}

func (repo *GormRepository) PostGroup(ctx context.Context, group *Group) error {
	return repo.db.GetDB(ctx).Create(group).Error
}

func (repo *GormRepository) PutGroup(ctx context.Context, group *Group) error {
	return repo.db.GetDB(ctx).Model(&Group{Id: group.Id}).Update("name", group.Name).Error
}

func (repo *GormRepository) PutGroupOrder(ctx context.Context, group *Group) error {
	return repo.db.GetDB(ctx).Model(&Group{Id: group.Id}).Update("order", group.Order).Error
}
