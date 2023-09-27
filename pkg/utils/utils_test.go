package utils_test

import (
	"testing"

	"github.com/sunjin7725/api-call-test/pkg/utils"
)

func TestCallApi(t *testing.T) {
	url := "http://openapi.foodsafetykorea.go.kr/api/sample/I2790/json/1/5"
	header, data := utils.CallApi(url, "GET", nil)
	t.Log(header)
	t.Log(data)
}
