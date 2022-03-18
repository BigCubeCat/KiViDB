package server

import (
	"encoding/json"
	"log"
	"net/http"
	"potatoDB/core"
)

type ClusterJSON struct {
	Cluster string
}

func ClusterHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var data ClusterJSON
		var values []string

		// Decoding get request data
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Panicf("Decoding error: %v\n", err)
		}
		// Getting values with API
		values, err = core.DBCore.GetAll(data.Cluster)
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
	default:
		log.Panicf("Wrong method: %v\n", r.Method)
	}
}
