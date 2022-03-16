package model

import (
	"context"
	"github.com/google/uuid"
)

type Group struct {
	Id    uuid.UUID `json:"id" gorm:"type:varchar(36);not null;primaryKey"`
	Name  string    `json:"name" gorm:"type:varchar(255);not null"`
	Order int       `json:"order" gorm:"type:int;not null"`
}

func GetGroups(ctx context.Context) ([]*Group, error) {
	var groups []*Group
	err := GetDB(ctx).Order("`order`").Find(&groups).Error
	return groups, err
}

func PostGroup(ctx context.Context, group *Group) error {
	return GetDB(ctx).Create(group).Error
}

func PutGroup(ctx context.Context, group *Group) error {
	return GetDB(ctx).Model(&Group{Id: group.Id}).Update("name", group.Name).Error
}

func PutGroupOrder(ctx context.Context, group *Group) error {
	return GetDB(ctx).Model(&Group{Id: group.Id}).Update("order", group.Order).Error
}
