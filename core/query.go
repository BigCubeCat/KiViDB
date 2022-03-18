package core

func (core *Core) GetAll(cluster string) ([]string, error) {
	var err error
	if err = core.ClusterExists(cluster); err != nil {
		return []string{}, err
	}
	return core.ClusterDocuments(cluster)
}
