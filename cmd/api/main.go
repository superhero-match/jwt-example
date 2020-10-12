package main

import (
	"github.com/jwt-example/cmd/api/controller"
	"github.com/jwt-example/internal/config"
	"log"
)


func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	ctrl, err := controller.NewController(cfg)
	if err != nil {
		panic(err)
	}

	r := ctrl.RegisterRoutes()

	log.Fatal(r.Run(":8080"))
}
