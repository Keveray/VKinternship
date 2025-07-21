package main

import (
	"github.com/Keveray/VKinternship"
	"github.com/Keveray/VKinternship/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(VKinternship.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error acucured while running http server: %s", err.Error())
	}

}
