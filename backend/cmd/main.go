package main

import (
	"github.com/sebastiaopamplona/sherpa2/api/handler"
	"log"

	sw "github.com/sebastiaopamplona/sherpa2/api/openapi"
)

func main() {
	apiHandler := handler.New()
	log.Printf("Server started")
	log.Fatal(sw.NewRouter(apiHandler.Routes).Run(":8080"))
}
