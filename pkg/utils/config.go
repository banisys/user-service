package utils

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DatabaseUrl string `mapstructure:"DB_URL"`
}

func LoadConfig(path string) (config Config, err error) {

	configFile := ".env"
	if os.Getenv("GO_ENV") == "test" {
		configFile = "../test.env"
	}

	viper.AddConfigPath(path)
	viper.SetConfigFile(configFile)

	viper.SetDefault("DB_URL", "pkg/database/database.db")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
