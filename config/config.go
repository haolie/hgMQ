package config

import (
	"sync"

	"github.com/spf13/viper"
)

const (
	con_config_path = "config.ini"
)

var (
	configMap = make(map[string]string, 16)
	loadOnce  = &sync.Once{}
)

func LoadConfig() error {
	var err error
	loadOnce.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("toml")
		viper.AddConfigPath(".")

		err = viper.ReadInConfig()
	})

	return err
}

func AddConfig(key string, v interface{}) {
	viper.Set(key, v)
}

func GetInt32(key string) int32 {
	return viper.GetInt32(key)
}

func GetInt64(key string) int64 {
	return viper.GetInt64(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}

func GetStr(key string) string {
	return viper.GetString(key)
}
