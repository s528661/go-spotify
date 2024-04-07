package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	id := os.Getenv("SPOTIFY_ID")
	secret := os.Getenv("SPOTIFY_SECRET")

	fmt.Println(id)
	fmt.Println(secret)
}
