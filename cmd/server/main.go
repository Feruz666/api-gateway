package main

import (
	"fmt"
	"log"

	"github.com/Feruz666/api-gateway/pkg"
	"github.com/Feruz666/api-gateway/util"
)

func main() {
	config, err := util.LoadConfig("./")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	fmt.Println(config.AUTH_SYSTEM_ADDRESS)

	server, err := pkg.ServiceRoutes(config)
	if err != nil {
		log.Fatal("cannot create server", err)
	}

	err = server.Start(config.SERVER_ADDRESS)
	if err != nil {
		log.Fatal("cannot start server", err)
	}

}
