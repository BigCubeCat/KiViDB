package core

import (
	"errors"
	"os"
	"path"
	"regexp"
)

const ClusterRegex = "[a-z0-9_]"

func (core *Core) CreateCluster(cluster string) error {
	var (
		err error
	)
	if err = core.ClusterExists(cluster); err == nil {
		return errors.New("cluster already exists")
	}
	if match, _ := regexp.MatchString(ClusterRegex, cluster); !match {
		return errors.New("cluster name must contain only lowercase and numbers")
	}
	return os.MkdirAll(path.Join(core.DirName, cluster), os.ModePerm)
}

func (core *Core) DropCluster(cluster string) error {
	var (
		err error
	)
	if err = core.ClusterExists(cluster); err != nil {
		return errors.New("cluster already exists")
	}
	return os.RemoveAll(path.Join(core.DirName, cluster))
}
