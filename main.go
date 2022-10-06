package main

import (
	"fmt"
	"kiviDB/api"
	"kiviDB/core"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func runFTPServer(directory string, port string) {
	// TODO: setup routing for many ftp servers
	fmt.Println("here");
	http.Handle("/", http.FileServer(http.Dir(directory)))

	log.Printf("Serving %s on HTTP port: %s\n", directory, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

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

	ftpPort := os.Getenv("FTP_PORT")
	ftpDir := os.Getenv("FTP_DIR")

	logFileName := "logs/" + time.Now().Format("01-02-2006 15-04-05") + ".log"
	if _, err = os.Stat("./logs"); os.IsNotExist(err) {
		_ = os.MkdirAll("./logs", os.ModePerm)
	}
	var f *os.File
	f, err = os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Unable to open log file: %v", err)
	}
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			log.Panicf("Unable to close log file: %v", err)
		}
	}(f)
	log.SetOutput(f)

	if startError := core.Init(dirName); startError != nil {
		log.Printf("Creating database folder with name: %v\n", dirName)
		if err = os.MkdirAll(dirName, os.ModePerm); err != nil {
			log.Fatalf("Unable to create database folder: %v\n", err)
		}
		if startError = core.Init(dirName); startError != nil {
			log.Fatalln(err)
		}
	}
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

	// calling ftp server MUST BE before run main server
	go runFTPServer(ftpDir, ftpPort)

	err = app.Listen(host + ":" + port)
	log.Fatal(err)

}
