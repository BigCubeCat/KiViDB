package core

import "log"

var DBCore *Core

type Core struct {
	DirName  string
	Clusters []string
}

// Init create Core struct object in core package
func Init(dirName string) error {
	err := FolderExists(dirName)
	if err != nil {
		log.Panic(err)
		return err
	}
	tables, err := GetTables(dirName)
	if err != nil {
		log.Panic(err)
		return err
	}
	clustersNames := make(map[string]bool)
	for _, clusterName := range tables {
		clustersNames[clusterName] = true
	}
	DBCore = &Core{
		DirName: dirName, Clusters: tables,
	}
	return nil
}
