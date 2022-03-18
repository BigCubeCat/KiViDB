package core

import (
	"os"
	"path"
)

func (core *Core) Get(cluster string, id string) ([]byte, error) {
	var (
		data []byte
		err  error
	)
	if err = core.ClusterExists(cluster); err != nil {
		return data, err
	}
	file := path.Join(core.DirName, cluster, id)
	if err = core.IdExists(file); err != nil {
		return data, err
	}
	data, err = os.ReadFile(id)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (core *Core) Set(cluster string, id string, data []byte) error {
	var err error
	if err = core.ClusterExists(cluster); err != nil {
		return err
	}
	file := path.Join(core.DirName, cluster, id)
	err = os.WriteFile(file, data, 0644)
	return err
}

func (core *Core) Add(cluster string, data []byte) (string, error) {
	var err error
	if err = core.ClusterExists(cluster); err != nil {
		return "", err
	}
	id := GenerateID()
	file := path.Join(core.DirName, cluster, id)
	err = os.WriteFile(file, data, 0644)
	return id, err
}
