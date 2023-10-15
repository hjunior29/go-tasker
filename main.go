package main

import (
	"log"
	"os"

	"github.com/hjunior29/go-tasker/api"
	"github.com/hjunior29/go-tasker/store"
	"github.com/hjunior29/go-tasker/worker"
	"github.com/joho/godotenv"
)

func loadenv() {
	if os.Getenv("DOCKER_ENV") != "" {
		log.Println("Running in Docker, skipping .env file loading")
		return
	}

	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	loadenv()

	if err := store.InitDatabase(); err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	go worker.StartWorker()
	go worker.StartWorker()
	go worker.StartWorker()

	port := os.Getenv("PORT")
	if port == "" {
		port = "6143"
	}

	api.SetupRoutes().Run("0.0.0.0:" + port)
}
