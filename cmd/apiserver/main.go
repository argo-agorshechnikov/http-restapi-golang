package main

import (
	"log"

	"github.com/argo-agorshechnikov/http-restapi-golang/internal/app/apiserver"
)

func main() {

	// Start to apiserver
	s := apiserver.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
