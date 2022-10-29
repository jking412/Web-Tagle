package database

import (
	"go-tagle/pkg/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(dbname string) {
	var err error
	DB, err = gorm.Open(sqlite.Open(dbname), &gorm.Config{})
	if err != nil {
		logger.ErrorString("database", "连接数据库失败", err.Error())
		panic(err)
	}
}
