package main

import "github.com/sunjin7725/api-call-test/internal/food_api"

func main() {
	food_data := food_api.GetData(2, 5, "")
	food_data.PrintJson()
}
