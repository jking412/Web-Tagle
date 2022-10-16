package viperlib

import (
	"fmt"
	"github.com/spf13/viper"
)

var Viper *viper.Viper

func InitViper() {
	Viper = viper.New()
	Viper.SetConfigType("yaml")
	Viper.AddConfigPath(".")
	Viper.SetConfigName("config")

	err := Viper.ReadInConfig()
	if err != nil {
		fmt.Println("读取配置文件失败")
		panic(err)
	}
	Viper.WatchConfig()
	Viper.AutomaticEnv()
}

func GetString(key string) string {
	return Viper.GetString(key)
}
