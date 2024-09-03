package env

import (
	"github.com/spf13/viper"
)

func Setup() {
	viper.AutomaticEnv()
}

func GetStringEnv(key string, def string) string {
	viper.SetDefault(key, def)
	return viper.GetString(key)
}

func GetIntEnv(key string, def int) int {
	viper.SetDefault(key, def)
	return viper.GetInt(key)
}
