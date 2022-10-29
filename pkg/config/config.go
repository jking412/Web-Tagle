package config

import "github.com/spf13/cast"

type MapData map[string]interface{}

func LoadEnv(key string, defaultValue ...interface{}) interface{} {
	if !internalViper.IsSet(key) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return internalViper.Get(key)
}

func LoadFloat64(key string, defaultValue ...interface{}) float64 {
	value := LoadEnv(key, defaultValue...)
	return cast.ToFloat64(value)
}

func LoadInt(key string, defaultValue ...interface{}) int {
	value := LoadEnv(key, defaultValue...)
	return cast.ToInt(value)
}

func LoadString(key string, defaultValue ...interface{}) string {
	value := LoadEnv(key, defaultValue...)
	return cast.ToString(value)
}

func LoadBool(key string, defaultValue ...interface{}) bool {
	value := LoadEnv(key, defaultValue...)
	return cast.ToBool(value)
}
