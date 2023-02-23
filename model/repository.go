package model

type Repository interface {
	IGroupRepository
	INotificationRepository
	ITagRepository
	ITaskRepository
	ITokenRepository
}

type GormRepository struct {
	db *DB
}
