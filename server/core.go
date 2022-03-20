package server

import (
	"encoding/json"
	"fmt"
	"kiviDB/core"
	"log"
	"net/http"
	"strings"
)

type PostJSON struct {
	Cluster string
	Id      string
	Value   string
}

type GetAndDeleteJSON struct {
	Cluster string
	Id      string
}

func CoreHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var (
			value []byte
			err   error
			data  GetAndDeleteJSON
		)

		args := strings.Split(strings.TrimPrefix(r.URL.Path, "/cluster/"), "/")
		if len(args) == 2 {
			data = GetAndDeleteJSON{Cluster: args[0], Id: args[0]}
			// Getting value with API
			value, err = core.DBCore.Get(data.Cluster, data.Id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Printf("API error: %v\n", err)
			}
			// Sending value back
			w.Header().Set("Content-Type", "application/json")
			err = json.NewEncoder(w).Encode(string(value))
			if err != nil {
				log.Printf("Encoding error: %v\n", err)
			}
		} else {
			fmt.Printf("Not enough arguments: expected 2, but got: %d", len(args))
		}
	case "POST":
		var data PostJSON
		// Decoding post request data
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Printf("Decoding error: %v\n", err)
		}
		// Using API
		if data.Id == "" {
			var id string
			id, err = core.DBCore.Add(data.Cluster, []byte(data.Value))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Printf("API error: %v\n", err)
			}
			fmt.Println(id)
		} else {
			err = core.DBCore.Set(data.Cluster, data.Id, []byte(data.Value))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Printf("API error: %v\n", err)
			}
		}
	case "DELETE":
		var data GetAndDeleteJSON

		// Decoding get request data
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Printf("Decoding error: %v\n", err)
		}
		// Deleting key and value
		err = core.DBCore.Delete(data.Cluster, data.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("API error: %v\n", err)
		}
	default:
		log.Printf("Wrong method: %v\n", r.Method)
	}
}
