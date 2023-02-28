package main

import (
	"R2/app/config"
	"R2/app/rest"
	"fmt"
	"log"
)

func main() {
	config.Load()

	server := rest.Router()
	port := fmt.Sprintf(":%s", config.ENV.Port)

	if err := server.Run(port); err != nil {
		log.Fatalln("error when server is initializing")
	}
}
