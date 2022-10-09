package main

import (
	"dog/api/middlewares"
	"dog/api/server"
	"dog/configs"
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("%s", err)
		}
	}()

	server := server.NewServer("./configs/app_config.yml")
	middlewares.AddDocumentationMiddleware(server)
	server.Run(configs.AppConfig.Application.Port)
}
