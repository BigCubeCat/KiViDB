package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

func HttpHandler(w http.ResponseWriter, r *http.Request) {
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
		fmt.Println("Wrong method! Use only POST, GET and DELETE")
	}
}
