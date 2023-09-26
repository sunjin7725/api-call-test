package main

import (
	"fmt"
	"path/filepath"

	"github.com/sunjin7725/api-call-test/configs"
)

func main() {
	config := configLoader("dev")
	fmt.Println(config.API_KEY)
}

func configLoader(env string) configs.Config {
	switch env {
	case "prod":
		config_path, _ := filepath.Abs("configs/config.prod.toml")
		return configs.LoadConfig(config_path)
	default:
		config_path, _ := filepath.Abs("configs/config.dev.toml")
		return configs.LoadConfig(config_path)
	}
}
