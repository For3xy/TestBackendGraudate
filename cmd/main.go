package main

import (
	"github.com/spf13/viper"
	"log"
	"testbackendGraudate/internal/handler"
	"testbackendGraudate/internal/repository"
	"testbackendGraudate/internal/service"
	"testbackendGraudate/pkg/server"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing config: %s", err.Error())
	}

	rep := repository.NewRepository()
	services := service.NewService(rep)
	handlers := handler.NewHandler(services)

	server := new(server.Server)
	if err := server.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("server run error: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
