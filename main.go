package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"potatoDB/core"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dirName := os.Getenv("DIR_NAME")
	if startError := core.Init(dirName); startError != nil {
		log.Fatal("Error: directory doesnt exists")
	}
}
