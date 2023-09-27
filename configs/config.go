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
	dir_path_slice := strings.Split(path.Dir(current_file_path), string(os.PathSeparator))
	dir_path_slice = append(dir_path_slice, "config."+os.Getenv("ENV")+".toml")
	config_file_path := path.Join(dir_path_slice...)
	var config Config

	viper.SetConfigFile(config_file_path)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	} else {
		config.API_KEY = viper.GetString("API_KEY")
	}
	return config
}
