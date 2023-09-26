package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	API_KEY string
}

func LoadConfig(path string) Config {
	var config Config

	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	} else {
		config.API_KEY = viper.GetString("API_KEY")
	}
	return config
}
