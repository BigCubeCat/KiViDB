package core

var DBCore *Core

type Core struct {
	DirName  string
	Clusters []string
}

func Init(dirName string) error {
	value, err := FolderExists(dirName)
	if value {
		tables, er := GetTables(dirName)
		if er != nil {
			return er
		}
		DBCore = &Core{DirName: dirName, Clusters: tables}
	} else {
		return err
	}
	return nil
}
