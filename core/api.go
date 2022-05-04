package core

import (
	"os"
	"path"
)

// Get returns document data by id from cluster
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
	data, err = os.ReadFile(file)
	if err != nil {
		return data, err
	}
	return data, nil
}

// Set changes value by id in cluster on data
// if document with id doesn't exist, create document
func (core *Core) Set(cluster string, id string, data []byte) error {
	var err error
	if err = core.ClusterExists(cluster); err != nil {
		return err
	}
	file := path.Join(core.DirName, cluster, id)
	err = os.WriteFile(file, data, 0644)
	return err
}

// Add creates document with auto-gen id and save data to cluster
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

// Delete removes document with id in cluster
func (core *Core) Delete(cluster string, id string) error {
	var err error
	if err = core.ClusterExists(cluster); err != nil {
		return err
	}
	file := path.Join(core.DirName, cluster, id)
	err = os.Remove(file)
	return err
}
