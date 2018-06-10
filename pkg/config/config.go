package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Config is a configuration with all parameters to access OS infra
type Config struct {
	Cloud Cloud `yaml:"cloud"`
}

// Cloud defines the cloud providers distribution and properties
type Cloud struct {
	Plugin string `yaml:"plugin"`
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

	err = yaml.Unmarshal(inputConfig, &newConfig)
	if err != nil {
		return nil, err
	}

	return &newConfig, nil
}
