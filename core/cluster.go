package core

import (
	"errors"
	"os"
	"path"
	"regexp"
)

const ClusterRegex = "[a-z0-9_]"

func (core *Core) CreateCluster(cluster string) error {
	if err := core.ClusterExists(cluster); err == nil {
		return errors.New("cluster with this name already exists")
	}
	if match, _ := regexp.MatchString(ClusterRegex, cluster); !match {
		return errors.New("cluster name must contain only lowercase letters and numbers")
	}
	core.Clusters = append(core.Clusters, cluster)
	if err := os.MkdirAll(path.Join(core.DirName, cluster), os.ModePerm); err != nil {
		return err
	}
	return nil
}

func (core *Core) DropCluster(cluster string) error {
	// Removing cluster from clusters slice
	var remove = func(slice []string, deleteValue string) []string {
		var newClusters []string
		for _, cluster := range slice {
			if deleteValue != cluster {
				newClusters = append(newClusters, cluster)
			}
		}
		return newClusters
	}
	core.Clusters = remove(core.Clusters, cluster)
	err := os.RemoveAll(path.Join(core.DirName, cluster))
	return err
}
