package config

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

var once sync.Once

var internalViper *viper.Viper

func InitConfig() {
	once.Do(func() {
		internalViper = newViper()
	})
}

func newViper() *viper.Viper {
	_viper := viper.New()
	_viper.SetConfigType("yaml")
	_viper.AddConfigPath(".")
	_viper.SetConfigName("config")

	err := _viper.ReadInConfig()
	if err != nil {
		fmt.Println("读取配置文件失败")
		panic(err)
	}
	_viper.WatchConfig()
	_viper.AutomaticEnv()
	return _viper
}

func GetString(key string) string {
	return internalViper.GetString(key)
}

func GetInt(key string) int {
	return internalViper.GetInt(key)
}

func GetBool(key string) bool {
	return internalViper.GetBool(key)
}
