package configs

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	API_KEY string
}

// func (c Config) init()

func LoadConfig() Config {
	os.Setenv("ENV", "dev")
	_, current_file_path, _, _ := runtime.Caller(0)
	configFilePath := strings.Join([]string{path.Dir(current_file_path), "config." + os.Getenv("ENV") + ".toml"}, string(os.PathSeparator))
	var config Config

	viper.SetConfigFile(configFilePath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	} else {
		config.API_KEY = viper.GetString("API_KEY")
	}
	return config
}
