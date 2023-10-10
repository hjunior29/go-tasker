package main

import (
	"log"
	"os"

	"github.com/hjunior29/go-tasker/routes"
	"github.com/hjunior29/go-tasker/store"
	"github.com/hjunior29/go-tasker/worker"
	_ "github.com/hjunior29/go-tasker/worker"
	"github.com/joho/godotenv"
)

func loadenv() {

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
		port = "5000"
	}

	routes.SetupRoutes().Run()
}
