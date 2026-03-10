package main

import (
	"log"
	"net/http"
	"os"
	"pulse/internal/db"

	"github.com/joho/godotenv"
)

func main() {
    if err := godotenv.Load("../.env"); err != nil {
        log.Fatal("Error loading .env file")
    }

    if err := db.Connect(); err != nil {
        log.Fatal(err)
    }
    log.Println("Connected to database")

    port := os.Getenv("SERVER_PORT")
    log.Printf("Server starting on port %s", port)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatal(err)
    }
}