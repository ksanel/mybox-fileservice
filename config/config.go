package config

import (
	"os"

	"github.com/spf13/viper"
)

var config *viper.Viper

func Init(env string) (*viper.Viper, error) {
	config := viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("./config/")
	err := config.ReadInConfig()
	if err != nil {
		return nil, err
	}

	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", config.GetString("app.gcpcredpath"))

	return config, nil
}

func GetConfig(key string) (string, error) {
	return config.GetString(key), nil
}
