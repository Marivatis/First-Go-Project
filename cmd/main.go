package main

import (
	"First-Go-Project/internal/config"
	"First-Go-Project/internal/server"
	"log"
)

func main() {
	cfg, err := config.LoadConfig("config/config.yaml")

	if err != nil {
		log.Fatal("Error loading config", err)
	}

	e := server.NewServer()

	if err := server.Start(e, cfg.Server.Port); err != nil {
		log.Fatal("Error starting server", err)
	}
}
