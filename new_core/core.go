package new_core

import "fmt"

type Core struct {
	DatabasePath string
	Clusters     []Cluster
}

var DatabaseCore *Core

func Init(databaseDirectoryName string) {
	DatabaseCore = &Core{DatabasePath: fmt.Sprintf("/%s/", databaseDirectoryName), Clusters: []Cluster{}}
	if IsDirectoryExists(DatabaseCore.DatabasePath) {

		return
	}

}
