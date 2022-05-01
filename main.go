package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"kiviDB/api"
	"kiviDB/core"
	"log"
	"os"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dirName := os.Getenv("DIR_NAME")
	if dirName == "" {
		dirName = "KiViDataBase"
	}
	port := os.Getenv("PORT")
	host := os.Getenv("HOST")
	if startError := core.Init(dirName); startError != nil {
		_ = os.MkdirAll(dirName, os.ModePerm)
		if startError = core.Init(dirName); startError != nil {
			log.Fatal("Error: directory doesn't exists")
		}
	}
	logFileName := "logs/" + time.Now().Format("01-02-2006 15-04-05 Mon") + ".log"
	if _, err := os.Stat("./logs"); os.IsNotExist(err) {
		_ = os.MkdirAll("./logs", os.ModePerm)
	}
	f, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Unable to open log file: %v", err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Panicf("Unable to close log file: %v", err)
		}
	}(f)
	log.SetOutput(f)
	app := fiber.New(fiber.Config{})
	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Post("/cluster/:id", api.PostClusterHandler)
	app.Delete("/cluster/:id", api.DeleteClusterHandler)
	app.Get("/cluster/:id", api.GetClusterHandler)

	app.Get("/doc/:cluster/:id", api.GetDocumentHandler)
	app.Post("/doc/:cluster/:id", api.PostDocumentHandler)
	app.Post("/doc/:cluster", api.CreateDocumentHandler)
	app.Delete("/doc/:cluster/:id", api.DeleteDocumentHandler)

	log.Println("Starting...")
	log.Printf("Listening %v:%v\n", host, port)
	err = app.Listen(host + ":" + port)
	log.Fatal(err)

}
