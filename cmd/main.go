package main

import (
	"github.com/Keveray/VKinternship"
	"github.com/Keveray/VKinternship/pkg/handler"
	"github.com/Keveray/VKinternship/pkg/repository"
	"github.com/Keveray/VKinternship/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(VKinternship.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error acucured while running http server: %s", err.Error())
	}

}
