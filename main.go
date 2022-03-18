package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"potatoDB/core"
)

type PostJSON struct {
	Cluster string
	ID      string
	Data    string
}

type GetJSON struct {
	Cluster string
	ID      string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dirName := os.Getenv("DIR_NAME")
	if startError := core.Init(dirName); startError != nil {
		log.Fatal("Error: directory doesnt exists")
	}
	http.HandleFunc("/", handler)
	// URL: http://127.0.0.1:8080/
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var data GetJSON
		var value []byte

		// Decoding get request data
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Fatal(err)
		}
		// Getting value with API
		value, err = core.DBCore.Get(data.Cluster, data.ID)
		if err != nil {
			log.Fatalf("API error: %s", err)
		}
		// Sending value back
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(string(value))
		if err != nil {
			log.Fatal(err)
		}
	case "POST":
		var data PostJSON
		// Decoding post request data
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println(err)
		}
		// Using API
		if data.ID == "" {
			var id string
			id, err = core.DBCore.Add(data.Cluster, []byte(data.Data))
			if err != nil {
				log.Fatalf("API error: %s", err)
			}
			fmt.Printf("New ID: %s\n", id)
		} else {
			err = core.DBCore.Set(data.Cluster, data.ID, []byte(data.Data))
			if err != nil {
				log.Fatalf("API error: %s", err)
			}
		}
	case "DELETE":
		fmt.Println(r.Method)
	default:
		fmt.Println("Wrong method!")
	}
}
