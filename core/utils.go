package core

import (
	"errors"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"os"
	"path"
)

type Object struct {
	Key   string
	Value string
}

func FolderExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
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
	if _, err := os.Stat(path.Join(core.DirName, cluster)); os.IsNotExist(err) {
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

func (core *Core) ClusterDocuments(cluster string) ([]string, error) {
	var (
		result []string
		err    error
	)
	if err = core.ClusterExists(cluster); err != nil {
		return result, err
	}
	file, e := ioutil.ReadDir(path.Join(core.DirName, cluster))
	if e != nil {
		return result, e
	}
	for _, f := range file {
		result = append(result, f.Name())
	}
	return result, nil
}

func (core *Core) ClusterValues(cluster string) ([]Object, error) {
	var (
		result []Object
		err    error
	)
	if err = core.ClusterExists(cluster); err != nil {
		return result, err
	}
	file, e := ioutil.ReadDir(path.Join(core.DirName, cluster))
	if e != nil {
		return result, e
	}
	var dat []byte
	for _, f := range file {
		dat, err = core.DocumentData(cluster, f.Name())
		if err != nil {
			log.Panicf("Error in get cluster values: %v", err)
		}
		result = append(result, Object{f.Name(), string(dat)})
	}
	return result, nil
}

func (core *Core) DocumentData(cluster string, id string) ([]byte, error) {
	return os.ReadFile(path.Join(core.DirName, cluster, id))
}
