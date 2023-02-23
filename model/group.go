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

type IGroupRepository interface {
	GetGroups(ctx context.Context) ([]*Group, error)
	PostGroup(ctx context.Context, group *Group) error
	PutGroup(ctx context.Context, group *Group) error
	PutGroupOrder(ctx context.Context, group *Group) error
}
