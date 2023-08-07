package main

import (
	"github.com/mendesx3/microservice-golang/config"
	"github.com/mendesx3/microservice-golang/infra/router"
	"log"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatal("Error initializing config")
	}
	router.Initialize()
}
