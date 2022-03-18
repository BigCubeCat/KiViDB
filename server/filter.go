package server

import (
	"encoding/json"
	"log"
	"net/http"
	"potatoDB/core"
)

type FilterJSON struct {
	Cluster string
	Regex   string
}

func FilterHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var (
			data   FilterJSON
			values []string
		)

		// Decoding get request data
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Panicf("Decoding error: %v\n", err)
		}
		// Getting value with API
		values, err = core.DBCore.Filter(data.Cluster, data.Regex)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Panicf("API error: %v\n", err)
		}
		// Sending value back
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(values)
		if err != nil {
			log.Panicf("Encoding error: %v\n", err)
		}
	case "DELETE":
		var data GetAndDeleteJSON

		// Decoding get request data
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Panicf("Decoding error: %v\n", err)
		}
		// Deleting key and value
		err = core.DBCore.Delete(data.Cluster, data.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Panicf("API error: %v\n", err)
		}
	default:
		log.Panicf("Wrong method: %v\n", r.Method)
	}
}
