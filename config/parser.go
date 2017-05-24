package config

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func ParseJson(configDir string) (ConfigurationOptions, error) {
	options, readFileErr := ioutil.ReadFile(configDir)
	if readFileErr != nil {
		fmt.Println("Error reading file", readFileErr)
	}

	var opts ConfigurationOptions
	err := json.Unmarshal(options, &opts)

	return opts, err
}

func ParseYml(configDir string) (ConfigurationOptions, error) {
	options, readFileErr := ioutil.ReadFile(configDir)
	if readFileErr != nil {
		fmt.Println("Error reading file", readFileErr)
	}

	var opts ConfigurationOptions
	err := yaml.Unmarshal(options, &opts)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return opts, err
}
