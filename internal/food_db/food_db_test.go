package food_db_test

import (
	"testing"

	"github.com/sunjin7725/api-call-test/internal/food_db"
)

func TestGetData(t *testing.T) {
	data := food_db.GetData(1, 5, "")
	if data.I2790.Result.Code != "INFO-000" {
		t.Error("API 결과를 가져오지 못했습니다.")
	}
}
