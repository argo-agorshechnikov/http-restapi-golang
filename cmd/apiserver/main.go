package main

import (
	"log"

	"github.com/argo-agorshechnikov/http-restapi-golang/internal/app/apiserver"
)

func main() {
	// Initial config
	config := apiserver.NewConfig()

	// Start to apiserver
	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
