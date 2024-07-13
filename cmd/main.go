package main

import (
	"log"

	"github.com/rohithrajasekharan/go-ecom/cmd/api"
)

func main() {
	server := api.NewAPIServer(":8090", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
