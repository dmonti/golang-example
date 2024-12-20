package main

import (
	"example/src/infra/http"
	"fmt"
)

func main() {
	server := http.Run()
	fmt.Printf("server started: %v\n", server)
}
