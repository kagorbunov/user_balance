package main

import (
	"github.com/kagorbunov/user_balance"
	"github.com/kagorbunov/user_balance/pkg/handler"
	"github.com/kagorbunov/user_balance/pkg/repository"
	"github.com/kagorbunov/user_balance/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(avitoAPI.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("configs")
	return viper.ReadInConfig()
}
