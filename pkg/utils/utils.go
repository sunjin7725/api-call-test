package utils

import "fmt"

func CallApi(url string, method string) string {
	switch method {
	case "post":
		goto API_CALL
	case "get":
		goto API_CALL
	default:
		panic("Check method parameter!")
	}
API_CALL:
	fmt.Println("testest")
	return url
}
