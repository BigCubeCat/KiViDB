package core

var DBCore *Core

type Core struct {
	DirName      string
	Clusters     []string
	clusterNames map[string]bool
}

func Init(dirName string) error {
	value, err := FolderExists(dirName)
	if value {
		tables, er := GetTables(dirName)
		if er != nil {
			return er
		}
		clustersNames := make(map[string]bool)
		for _, clusterName := range tables {
			clustersNames[clusterName] = true
		}
		DBCore = &Core{
			DirName: dirName, Clusters: tables, clusterNames: clustersNames,
		}
	}
	return err
}
