package config

import (
	"encoding/json"
	"io/ioutil"
)

// Config is a configuration with all parameters to access OS infra
type Config struct {
	Cloud Cloud `json:"cloud"`
}

// Cloud defines the cloud providers distribution and properties
type Cloud struct {
	Plugin     string     `json:"plugin"`
	Properties Properties `json:"properties"`
}

// Properties defines properties of a cloud provider
type Properties struct {
	Openstack Openstack `json:"openstack"`
}

// Openstack defines data from the openstack cloud
type Openstack struct {
}

func readConfig(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

// ParseConfig returns a config object from a path to a config file
func ParseConfig(path string) (*Config, error) {
	var newConfig Config
	inputConfig, err := readConfig(path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(inputConfig, &newConfig)
	if err != nil {
		return nil, err
	}

	return &newConfig, nil
}
