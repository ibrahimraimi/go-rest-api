package main

import (
	"log"

	"github.com/ibrahimraimi/go-rest-api/cmd/api"
)

func main() {
	server := api.NewAPIServer(":8000", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
