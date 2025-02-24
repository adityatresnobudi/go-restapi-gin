package main

import (
	"log"

	"github.com/adityatresnobudi/go-restapi-gin/config"
	"github.com/adityatresnobudi/go-restapi-gin/internal/server"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("err loading .env file: %s\n", err.Error())
	}
}

func main() {
	cfg := config.NewConfig()

	s := server.NewServer(cfg)

	s.Run()
}
