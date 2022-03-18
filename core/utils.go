package core

import (
	"errors"
	"github.com/google/uuid"
	"io/ioutil"
	"os"
)

func FolderExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetTables(path string) ([]string, error) {
	var names []string
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return names, err
	}

	for _, f := range files {
		names = append(names, f.Name())
	}
	return names, nil
}

func GenerateID() string {
	id := uuid.New()
	return id.String()
}

func (core *Core) ClusterExists(cluster string) error {
	_, ok := core.clusterNames[cluster]
	if !ok {
		return errors.New("cluster doesnt exists")
	}
	return nil
}

func (core *Core) IdExists(id string) error {
	if _, err := os.Stat(id); errors.Is(err, os.ErrNotExist) {
		return err
	}
	return nil
}
