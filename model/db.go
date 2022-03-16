package model

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

var (
	db     *gorm.DB
	tables = []interface{}{
		&Group{},
		&Tag{},
		&Task{},
		&TagMap{},
	}
)

func InitDB() {
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

	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true&loc=Asia%%2FTokyo&charset=utf8mb4", user, pass, host, dbname)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		panic(err)
	}

	db = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci")

	err = db.AutoMigrate(tables...)
	if err != nil {
		panic(err)
	}
}

func GetDB(ctx context.Context) *gorm.DB {
	return db.WithContext(ctx)
}
