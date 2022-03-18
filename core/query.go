package core

import (
	"encoding/json"
	"log"
	"regexp"
)

func (core *Core) GetAll(cluster string) ([]string, error) {
	var err error
	if err = core.ClusterExists(cluster); err != nil {
		return []string{}, err
	}
	return core.ClusterDocuments(cluster)
}

func (core *Core) Filter(cluster string, query string) ([]string, error) {
	var (
		err      error
		result   []Object
		row      Object
		jsonByte []byte
	)
	if err = core.ClusterExists(cluster); err != nil {
		return []string{}, err
	}
	result, err = core.ClusterValues(cluster)
	var accepted []string
	_, er := regexp.Compile(query)
	if er != nil {
		log.Println("Error in regex")
		return accepted, er
	}
	for _, value := range result {
		if matched, _ := regexp.MatchString(query, value.Value); matched {
			jsonByte, err = json.Marshal(row)
			if err != nil {
				log.Println(err)
			}
			accepted = append(accepted, string(jsonByte))
		}
	}
	return accepted, nil
}
