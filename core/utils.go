package core

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"os"
	"path"
)

type Object struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func FolderExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return err
	}
	return nil
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
	if err := FolderExists(path.Join(core.DirName, cluster)); err != nil {
		return errors.New(fmt.Sprintf("cluster [%v] doesn't exists", cluster))
	}
	return nil
}

func (core *Core) IdExists(id string) error {
	if _, err := os.Stat(id); errors.Is(err, os.ErrNotExist) {
		return errors.New(fmt.Sprintf("document [%v] doesn't exists", id))
	}
	return nil
}

func (core *Core) ClusterDocuments(cluster string) ([]string, error) {
	var (
		result []string
	)
	if err := core.ClusterExists(cluster); err != nil {
		return result, err
	}
	files, err := ioutil.ReadDir(path.Join(core.DirName, cluster))
	if err != nil {
		return result, err
	}
	for _, file := range files {
		result = append(result, file.Name())
	}
	return result, nil
}

func (core *Core) ClusterValues(cluster string) ([]Object, error) {
	var (
		result []Object
	)
	if err := core.ClusterExists(cluster); err != nil {
		return result, err
	}
	files, err := ioutil.ReadDir(path.Join(core.DirName, cluster))
	if err != nil {
		return result, err
	}
	var data []byte
	for _, file := range files {
		data, err = core.DocumentData(cluster, file.Name())
		if err != nil {
			log.Panicf("unable to get cluster values: %v", err)
		}
		result = append(result, Object{file.Name(), string(data)})
	}
	return result, nil
}

func (core *Core) DocumentData(cluster string, id string) ([]byte, error) {
	return os.ReadFile(path.Join(core.DirName, cluster, id))
}
