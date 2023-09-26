package main

import (
	"encoding/json"
	"io"
	"net/http"
	"path/filepath"

	"github.com/davecgh/go-spew/spew"
	"github.com/sunjin7725/api-call-test/configs"
)

type Response struct {
	I2790 I2790 `json:"I2790"`
}

type I2790 struct {
	TotalCount string       `json:"total_count"`
	Result     StatusResult `json:"RESULT"`
	Row        []RowItem    `json:"row"`
}

type StatusResult struct {
	Msg  string `json:"MSG"`
	Code string `json:"CODE"`
}

type RowItem struct {
	Num                string `json:"NUM"`
	FoodCd             string `json:"FOOD_CD"`
	GroupName          string `json:"GROUP_NAME"`
	DescKor            string `json:"DESC_KOR"`
	ResearchYear       string `json:"RESEARCH_YEAR"`
	MakerName          string `json:"MAKER_NAME"`
	ServingSize        string `json:"SERVING_SIZE"`
	SubRefName         string `json:"SUB_REF_NAME"`
	SamplingRegionCd   string `json:"SAMPLING_REGION_CD"`
	SamplingRegionName string `json:"SAMPLING_REGION_NAME"`
	SamplingMonthCd    string `json:"SAMPLING_MONTH_CD"`
	SamplingMonthName  string `json:"SAMPLING_MONTH_NAME"`
	NutrCont1          string `json:"NUTR_CONT1"`
	NutrCont2          string `json:"NUTR_CONT2"`
	NutrCont3          string `json:"NUTR_CONT3"`
	NutrCont4          string `json:"NUTR_CONT4"`
	NutrCont5          string `json:"NUTR_CONT5"`
	NutrCont6          string `json:"NUTR_CONT6"`
	NutrCont7          string `json:"NUTR_CONT7"`
	NutrCont8          string `json:"NUTR_CONT8"`
	NutrCont9          string `json:"NUTR_CONT9"`
}

func main() {
	config := configLoader("dev")
	basic_url := "http://openapi.foodsafetykorea.go.kr/api"
	url := basic_url + "/" + config.API_KEY + "/I2790/json/1/5"
	// test sample
	// url := "http://openapi.foodsafetykorea.go.kr/api/sample/I2790/json/1/5"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var response Response
	json.Unmarshal(data, &response)

	spew.Dump(response)
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
