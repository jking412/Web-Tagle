package database

import (
	"fmt"
	"go-tagle/pkg/viperlib"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open(viperlib.GetString("database.dbname")), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库初始化失败")
		panic(err)
	}
}
