package configs_test

import (
	"reflect"
	"testing"

	"github.com/sunjin7725/api-call-test/configs"
)

func TestLoadConfig(t *testing.T) {
	config := configs.LoadConfig("config.dev.toml")
	if reflect.TypeOf(config.API_KEY).Kind() != reflect.String {
		t.Log(config)
		t.Error("Cannot read config")
	}
}
