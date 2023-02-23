package model

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

type DB struct {
	db *gorm.DB
}

var tables = []interface{}{
	&Group{},
	&Tag{},
	&Task{},
	&TagMap{},
	&Token{},
	&NotificationTime{},
	&Notification{},
}

func InitDB() (*DB, error) {
	user := os.Getenv("MARIADB_USERNAME")
	pass := os.Getenv("MARIADB_PASSWORD")
	host := os.Getenv("MARIADB_HOSTNAME")
	dbname := os.Getenv("MARIADB_DATABASE")
	isProduction := os.Getenv("PRODUCTION") == "true"

	var logLevel logger.LogLevel
	if isProduction {
		logLevel = logger.Silent
	} else {
		logLevel = logger.Info
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true&loc=Asia%%2FTokyo&charset=utf8mb4", user, pass, host, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	db = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci")

	err = db.AutoMigrate(tables...)
	if err != nil {
		return nil, fmt.Errorf("failed in table's migration: %w", err)
	}

	return &DB{db}, nil
}

func (db *DB) GetDB(ctx context.Context) *gorm.DB {
	return db.db.WithContext(ctx)
}
