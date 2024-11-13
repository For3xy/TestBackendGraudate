package main

import (
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"testbackendGraudate/configs"
	"testbackendGraudate/internal/handler"
	"testbackendGraudate/internal/repository"
	"testbackendGraudate/internal/service"
	"testbackendGraudate/pkg/server"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}

	db, err := configs.NewPostgresDB(configs.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("error connecting to database: %s", err.Error())
	}

	rep := repository.NewRepository(db)
	services := service.NewService(rep)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if errServer := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); errServer != nil {
		logrus.Fatalf("server run error: %s", errServer.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
