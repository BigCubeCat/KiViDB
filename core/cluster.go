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
	core.Clusters = append(core.Clusters, cluster)
	core.clusterNames[cluster] = true
	return os.MkdirAll(path.Join(core.DirName, cluster), os.ModePerm)
}

func (core *Core) DropCluster(cluster string) error {
	var (
		err error
	)
	if err = core.ClusterExists(cluster); err != nil {
		return errors.New("cluster already exists")
	}
	delete(core.clusterNames, cluster)
	var remove = func(slice []string, deleteValue string) []string {
		var newClusters []string
		for _, clstr := range slice {
			if deleteValue != clstr {
				newClusters = append(newClusters, clstr)
			}
		}
		return newClusters
	}
	core.Clusters = remove(core.Clusters, cluster)
	return os.RemoveAll(path.Join(core.DirName, cluster))
}
