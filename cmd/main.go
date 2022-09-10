package main

import (
	"fmt"
	"log"

	"github.com/danieeelfr/my-large-json-file-reader/internal/config"
	"github.com/danieeelfr/my-large-json-file-reader/internal/repository"
	"github.com/danieeelfr/my-large-json-file-reader/internal/service"
	"github.com/joho/godotenv"
)

func main() {
	cfg := config.NewConfig()

	repo := repository.NewRepository(cfg)
	srv := service.NewPortService(cfg)

	ports, err := srv.ReadAndDecode()
	if err != nil {
		log.Fatal(err)
	}

	if err := repo.Set(ports.Records); err != nil {
		log.Fatal(err)
	}

	// TODO
	// result, err := repo.Get("", 100)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Println("finished without errors!")
}

func init() {
	fmt.Println("starting the application...")
	err := godotenv.Load("./../local.env")
	if err != nil {
		_ = godotenv.Load("./local.env")
	}
}
