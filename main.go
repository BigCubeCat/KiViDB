package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"kiviDB/api"
	"kiviDB/core"
	"kiviDB/logger"
	"log"
	"net/http"
	"os"
)

func runFTPServer(directory string, port string) {
	// TODO: setup routing for many ftp servers
	// fmt.Println("here");
	http.Handle("/", http.FileServer(http.Dir(directory)))

	log.Printf("Serving %s on HTTP port: %s\n", directory, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// getting database folder name
	dirName := os.Getenv("DIR_NAME")
	if dirName == "" {
		dirName = "KiViDataBase"
	}
	// http server vars
	port := os.Getenv("PORT")
	host := os.Getenv("HOST")
	// ftp server vars
	ftpPort := os.Getenv("FTP_PORT")
	ftpDir := os.Getenv("FTP_DIR")

	logger.Init()        // setting up logger
	defer logger.Close() // defer closing log file

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
