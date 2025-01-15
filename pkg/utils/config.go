package utils

import (
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseUrl string `mapstructure:"DB_URL"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	viper.SetDefault("DB_URL", "pkg/database/database.db")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
