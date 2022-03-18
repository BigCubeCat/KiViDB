package core

import (
	"encoding/json"
	"log"
	"regexp"
)

func (core *Core) GetAll(cluster string) ([]string, error) {
	var (
		err      error
		values   []Object
		output   []string
		jsonByte []byte
	)
	if err = core.ClusterExists(cluster); err != nil {
		return []string{}, err
	}
	values, err = core.ClusterValues(cluster)
	for _, doc := range values {
		jsonByte, err = json.Marshal(doc)
		if err != nil {
			log.Panicf("Invalid document: %v", err)
		}
		output = append(output, string(jsonByte))
	}
	return output, nil
}

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
	_, er := regexp.Compile(query)
	if er != nil {
		log.Println("Error in regex")
		return er
	}
	for _, value := range result {
		if matched, _ := regexp.MatchString(query, value.Value); matched {
			if err = core.Delete(cluster, value.Key); err != nil {
				log.Panicf("Cant delete document, %v", err)
			}
		}
	}
	return nil
}
