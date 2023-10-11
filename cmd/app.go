package main

import (
	"github.com/sunjin7725/api-call-test/internal"
)

func main() {
	r := internal.SetupRouter()
	r.Run(":8080")
}
