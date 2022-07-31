package main

import (
	"fmt"
	"src/config"
	"src/router"
)

func main() {
	conf, err := config.Init()
	if err != nil {
		panic(fmt.Sprintf("Unable to load the config details: %v ", err.Error()))
	}
	server := router.Server{Gin: conf.Gin}
	server.Start()
}
