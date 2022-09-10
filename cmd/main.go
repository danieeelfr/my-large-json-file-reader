package main

import (
	"fmt"

	"github.com/danieeelfr/my-large-json-file-reader/internal/config"
	"github.com/joho/godotenv"
)

func main() {
	_ = config.NewConfig()
}

func init() {
	fmt.Println("starting the application...")
	err := godotenv.Load("./../local.env")
	if err != nil {
		_ = godotenv.Load("./local.env")
	}
}
