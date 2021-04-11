package main

import (
	"fmt"

	"github.com/jahangiranwari/cache-service/api"
)

func main() {
	fmt.Println("Starting service...")

	api.StartServer()
}
