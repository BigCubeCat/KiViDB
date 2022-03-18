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
	logFileName := os.Getenv("LOG_FILE")
	if startError := core.Init(dirName); startError != nil {
		log.Fatal("Error: directory doesnt exists")
	}
	path := logFileName
	_ = os.Remove(path)
	f, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Unable to open log file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println("This is a test log entry")
	http.HandleFunc("/", server.HttpHandler)
	err = http.ListenAndServe(":"+address, nil)
	if err != nil {
		log.Fatal(err)
	}
}
