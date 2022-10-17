package boot

import (
	"go-tagle/model"
	"go-tagle/pkg/database"
)

func initDB() {
	database.ConnectDB()
	database.DB.AutoMigrate(&model.User{},
		&model.Habit{},
		&model.Task{})
}
