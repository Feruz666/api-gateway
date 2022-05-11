package main

import (
	"log"

	"github.com/Feruz666/api-gateway/pkg"
	"github.com/Feruz666/api-gateway/util"
)

const (
	serverAddress = "0.0.0.0:8000"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	server, err := pkg.ServiceRoutes(config)
	if err != nil {
		log.Fatal("cannot create server", err)
	}

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}

}
