package core

import (
	"encoding/json"
	"log"
	"regexp"
)

func (core *Core) Filter(cluster string, query string) ([]string, error) {
	var (
		err      error
		result   []Object
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
			jsonByte, err = json.Marshal(value)
			if err != nil {
				log.Println(err)
			}
			accepted = append(accepted, string(jsonByte))
		}
	}
	return accepted, nil
}

func (core *Core) DeleteByRegex(cluster string, query string) error {
	var (
		err    error
		result []Object
	)
	if err = core.ClusterExists(cluster); err != nil {
		return err
	}
	result, err = core.ClusterValues(cluster)
	_, err = regexp.Compile(query)
	if err != nil {
		log.Println("error in regex")
		return err
	}
	for _, value := range result {
		if matched, _ := regexp.MatchString(query, value.Value); matched {
			if err = core.Delete(cluster, value.Key); err != nil {
				log.Printf("cannot delete document: %v", err)
				return err
			}
		}
	}
	return nil
}
