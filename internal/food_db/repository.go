package food_db

import (
	"encoding/json"
	"io"
	"strconv"

	"github.com/sunjin7725/api-call-test/configs"
	"github.com/sunjin7725/api-call-test/pkg/utils"
)

func GetData(page int, perPage int, searchKeyword string) Response {
	config := configs.LoadConfig()

	startRowNum := strconv.FormatInt(int64(((page-1)*perPage)+1), 10)
	endRowNum := strconv.FormatInt(int64(page*perPage), 10)

	basicUrl := "http://openapi.foodsafetykorea.go.kr/api"
	url := basicUrl + "/" + config.API_KEY + "/I2790/json/" + startRowNum + "/" + endRowNum
	// test sample
	// url := "http://openapi.foodsafetykorea.go.kr/api/sample/I2790/json/1/5"
	_, responseBody := utils.CallApi(url, "GET", nil)

	var response Response
	data, err := io.ReadAll(responseBody)
	if err != nil {
		panic(err)
	}
	defer responseBody.Close()

	json.Unmarshal(data, &response)
	return response
}
