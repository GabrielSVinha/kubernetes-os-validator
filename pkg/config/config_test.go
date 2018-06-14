package config

import (
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	path := os.Getenv("FIXTURES") + "/test/config/config.json"
	config, err := ParseConfig(path)
	if err != nil {
		t.Errorf("Error while parsing fiile: %s" + err.Error())
	}
	conf := &Config{}
	if config == conf {
		t.Errorf("Invalid config parsing")
	}
	cloud := Cloud{}
	if config.Cloud == cloud {
		t.Errorf("Invalid config parsing in 'cloud' parameter")
	}
	cloud.Plugin = "openstack"
	if config.Cloud.Plugin != cloud.Plugin {
		t.Errorf("Invalid config parsing in 'cloud.plugin' parameter")
	}
	cloud.Properties = Properties{}
	if config.Cloud.Properties != cloud.Properties {
		t.Errorf("Invalid config parsing in 'cloud.properties' parameter")
	}
}

func TestEmptyCloud(t *testing.T) {
	path := os.Getenv("FIXTURES") + "/test/config/empty_cloud.json"
	config, err := ParseConfig(path)
	if err != nil {
		t.Errorf("Error while parsing fiile: %s" + err.Error())
	}
	conf := &Config{}
	if config == conf {
		t.Errorf("Invalid config parsing")
	}
	cloud := Cloud{}
	if config.Cloud != cloud {
		t.Errorf("Invalid config parsing in 'cloud' parameter")
	}
}

func TestEmptyParse(t *testing.T) {
	path := os.Getenv("FIXTURES") + "/test/config/empty_config.json"
	config, err := ParseConfig(path)
	if err != nil {
		t.Errorf("Error while parsing fiile: %s" + err.Error())
	}
	conf := &Config{}
	if config == conf {
		t.Errorf("Invalid config parsing")
	}
	cloud := Cloud{}
	if config.Cloud != cloud {
		t.Errorf("Invalid config parsingin 'cloud' parameter")
	}
}
