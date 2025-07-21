package main

import (
	"github.com/Keveray/VKinternship"
	"github.com/Keveray/VKinternship/pkg/handler"
	"github.com/Keveray/VKinternship/pkg/repository"
	"github.com/Keveray/VKinternship/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing config: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(VKinternship.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error acucured while running http server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.SetEnvPrefix("APP")
	viper.AutomaticEnv()
	return viper.ReadInConfig()
}
