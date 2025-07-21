package main

import (
	VKinternship "github.com/Keveray/vk"
	"log"
)

func main() {
	srv := new(VKinternship.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error acucured while running http server: %s", err.Error())
	}
}
