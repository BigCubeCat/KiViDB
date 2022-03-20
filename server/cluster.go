package server

import (
	"encoding/json"
	"kiviDB/core"
	"log"
	"net/http"
	"strings"
)

type ClusterJSON struct {
	Cluster string
}

func ClusterHandler(w http.ResponseWriter, r *http.Request) {
	var (
		arg string
		err error
	)
	arg = strings.TrimPrefix(r.URL.Path, "/cluster/")
	if arg == "" {
		log.Printf("Argument is empty!")
	} else {
		switch r.Method {
		case "GET":
			var values []string
			values, err = core.DBCore.GetAll(arg)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				log.Printf("API error: %v\n", err)
			}
			w.Header().Set("Content-Type", "application/json")
			err = json.NewEncoder(w).Encode(values)
			if err != nil {
				log.Printf("Encoding error: %v\n", err)
			}
		case "DELETE":
			if err = core.DBCore.DropCluster(arg); err != nil {
				http.Error(w, "API error: "+err.Error(), http.StatusBadRequest)
				log.Printf("API error: %v\n", err)
			}
			w.Header().Set("Content-Type", "application/json")
		case "POST":
			if err = core.DBCore.CreateCluster(arg); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				log.Printf("API error: %v\n", err)
			}
			w.Header().Set("Content-Type", "application/json")
		default:
			log.Printf("Wrong method: %v\n", r.Method)
		}
	}
}
