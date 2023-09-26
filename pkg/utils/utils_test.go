package utils_test

import (
	"testing"

	"github.com/sunjin7725/api-call-test/pkg/utils"
)

func TestCallApi(t *testing.T) {
	const url string = "url test"
	return_val := utils.CallApi(url, "post")
	t.Log(return_val)
}
