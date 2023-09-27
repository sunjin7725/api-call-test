package food_api

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type Response struct {
	I2790 I2790 `json:"I2790"`
}

type I2790 struct {
	TotalCount stringToInt  `json:"total_count"`
	Result     StatusResult `json:"RESULT"`
	Row        []RowItem    `json:"row"`
}

type StatusResult struct {
	Msg  string `json:"MSG"`
	Code string `json:"CODE"`
}

type RowItem struct {
	Num                stringToInt     `json:"NUM"`
	FoodCd             string          `json:"FOOD_CD"`
	GroupName          string          `json:"GROUP_NAME"`
	DescKor            string          `json:"DESC_KOR"`
	ResearchYear       string          `json:"RESEARCH_YEAR"`
	MakerName          string          `json:"MAKER_NAME"`
	ServingSize        stringToInt     `json:"SERVING_SIZE"`
	SubRefName         string          `json:"SUB_REF_NAME"`
	SamplingRegionCd   string          `json:"SAMPLING_REGION_CD"`
	SamplingRegionName string          `json:"SAMPLING_REGION_NAME"`
	SamplingMonthCd    string          `json:"SAMPLING_MONTH_CD"`
	SamplingMonthName  string          `json:"SAMPLING_MONTH_NAME"`
	NutrCont1          stringToFloat32 `json:"NUTR_CONT1"`
	NutrCont2          stringToFloat32 `json:"NUTR_CONT2"`
	NutrCont3          stringToFloat32 `json:"NUTR_CONT3"`
	NutrCont4          stringToFloat32 `json:"NUTR_CONT4"`
	NutrCont5          stringToFloat32 `json:"NUTR_CONT5"`
	NutrCont6          stringToFloat32 `json:"NUTR_CONT6"`
	NutrCont7          stringToFloat32 `json:"NUTR_CONT7"`
	NutrCont8          stringToFloat32 `json:"NUTR_CONT8"`
	NutrCont9          stringToFloat32 `json:"NUTR_CONT9"`
}

func (resp Response) PrintJson() {
	json_data, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(json_data))
}

type stringToFloat32 float32

type stringToInt int

func (v *stringToFloat32) UnmarshalJSON(data []byte) error {
	return_val, err := strconv.ParseFloat(strings.ReplaceAll(string(data), "\"", ""), 32)
	if err != nil {
		panic(err)
	}
	*v = stringToFloat32(return_val)
	return err
}

func (v *stringToInt) UnmarshalJSON(data []byte) error {
	return_val, err := strconv.Atoi(strings.ReplaceAll(string(data), "\"", ""))
	if err != nil {
		panic(err)
	}
	*v = stringToInt(return_val)
	return err
}
