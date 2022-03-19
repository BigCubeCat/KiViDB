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
	var data ClusterJSON
	// Decoding request's data
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Decoding error: %v\n", err)
	}
	switch r.Method {
	case "GET":
		var values []string
		values, err = core.DBCore.GetAll(data.Cluster)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Printf("API error: %v\n", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(values)
		if err != nil {
			log.Printf("Encoding error: %v\n", err)
		}
	case "DELETE":
		if err = core.DBCore.DropCluster(data.Cluster); err != nil {
			http.Error(w, "API error: "+err.Error(), http.StatusBadRequest)
			log.Printf("API error: %v\n", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	case "POST":
		if err = core.DBCore.CreateCluster(data.Cluster); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Printf("API error: %v\n", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	default:
		log.Printf("Wrong method: %v\n", r.Method)
	}
}
