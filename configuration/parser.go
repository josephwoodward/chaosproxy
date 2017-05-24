package configuration

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func ParseJson(configDir string) (ConfigurationOptions, error) {
	configOptions, readFileErr := ioutil.ReadFile(configDir)
	if readFileErr != nil {
		fmt.Println("Error reading file", readFileErr)
	}

	var opts ConfigurationOptions
	err := json.Unmarshal(configOptions, &opts)

	return opts, err
}

func ParseYml(configDir string) (ConfigurationOptions, error) {
	configOptions, readFileErr := ioutil.ReadFile(configDir)
	if readFileErr != nil {
		fmt.Println("Error reading file", readFileErr)
	}

	var opts ConfigurationOptions
	err := yaml.Unmarshal(configOptions, &opts)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return opts, err
}
