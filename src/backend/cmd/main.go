package main

import (
	"dog/api/middlewares"
	"dog/api/server"
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("%s", err)
		}
	}()

	server := server.NewServer()
	middlewares.AddDocumentationMiddleware(server)
	server.Run("80")
}
