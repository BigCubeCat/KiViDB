package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"potatoDB/core"
	"potatoDB/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dirName := os.Getenv("DIR_NAME")
	address := os.Getenv("ADDRESS")
	if startError := core.Init(dirName); startError != nil {
		log.Fatal("Error: directory doesnt exists")
	}
	http.HandleFunc("/", server.HttpHandler)
	err = http.ListenAndServe(":"+address, nil)
	if err != nil {
		log.Fatal(err)
	}
}
