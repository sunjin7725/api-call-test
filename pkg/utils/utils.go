package utils

import (
	"io"
	"net/http"
)

func CallApi(url string, method string, body io.Reader) (http.Header, io.ReadCloser) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// header := resp.Header
	// data, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }
	return resp.Header, resp.Body
}
