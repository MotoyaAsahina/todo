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

func (repo *GormRepository) InitDB() error {
	db, err := initDB()
	if err != nil {
		return err
	}
	repo.db = db
	return nil
}
