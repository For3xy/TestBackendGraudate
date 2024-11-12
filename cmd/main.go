package main

import (
	"log"
	"testbackendGraudate"
	"testbackendGraudate/pkg/handler"
)

func main() {
	server := new(testbackendGraudate.Server)
	handlers := new(handler.Handler)
	if err := server.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("server run error: %s", err.Error())
	}
}
