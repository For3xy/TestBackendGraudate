package main

import (
	"log"
	"testbackendGraudate/internal/handler"
	"testbackendGraudate/internal/repository"
	"testbackendGraudate/internal/service"
	"testbackendGraudate/pkg/server"
)

func main() {
	rep := repository.NewRepository()
	services := service.NewService(rep)
	handlers := handler.NewHandler(services)

	server := new(server.Server)
	if err := server.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("server run error: %s", err.Error())
	}
}
