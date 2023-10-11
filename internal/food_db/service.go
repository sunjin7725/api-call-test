package food_db

import (
	"encoding/json"
	"io"
	"strconv"

	"github.com/sunjin7725/api-call-test/configs"
	"github.com/sunjin7725/api-call-test/pkg/utils"
)

func GetData(page int, per_page int, search_keyword string) Response {
	config := configs.LoadConfig()

	start_row_num := strconv.FormatInt(int64(((page-1)*per_page)+1), 10)
	end_row_num := strconv.FormatInt(int64(page*per_page), 10)

	basic_url := "http://openapi.foodsafetykorea.go.kr/api"
	url := basic_url + "/" + config.API_KEY + "/I2790/json/" + start_row_num + "/" + end_row_num
	// test sample
	// url := "http://openapi.foodsafetykorea.go.kr/api/sample/I2790/json/1/5"
	_, response_body := utils.CallApi(url, "GET", nil)

	var response Response
	data, err := io.ReadAll(response_body)
	if err != nil {
		panic(err)
	}
	defer response_body.Close()

	json.Unmarshal(data, &response)
	return response
}
