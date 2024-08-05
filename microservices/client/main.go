package main

import (
	"context"
	"log"
)

func main() {
	serviceName := "UserService"
	server, cleanup := InitializeServer(serviceName)
	defer cleanup()
	if err := server.Serve(context.Background()); err != nil {
		log.Panic(err)
	}
}
