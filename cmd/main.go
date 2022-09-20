package main

import (
	"WildFireTest/config"
	"WildFireTest/controller"
	"WildFireTest/server"
)

func main() {
	appConfiguration, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	server := server.NewServer(
		appConfiguration,
		controller.NewWildFireController(
			appConfiguration.App.Count,
			appConfiguration.App.Limit,
		),
	)

	if err = server.Run(); err != nil {
		panic(err)
	}
}
