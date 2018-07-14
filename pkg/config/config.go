package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Entry struct {
	Name   string
	Params []string
}

type Config struct {
	// First element of pipeline
	Source map[string][]interface{} `yaml:"source"`
	// List of Sinks
	Sinks map[string][]interface{} `yaml:"sinks"`
}

func Read(file string) (Config, error) {

	config := Config{}

	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return config, err
	}

	if len(config.Source) != 1 {
		return config, fmt.Errorf("Config should have exactly one source defined")
	}

	return config, nil
}
