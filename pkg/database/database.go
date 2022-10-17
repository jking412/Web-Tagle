package database

import (
	"go-tagle/pkg/logger"
	"go-tagle/pkg/viperlib"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open(viperlib.GetString("database.dbname")), &gorm.Config{})
	if err != nil {
		logger.ErrorString("database", "连接数据库失败", err.Error())
		panic(err)
	}
}
