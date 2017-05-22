package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func ParseConfig(configDir string) (ConfigurationOptions, error) {
	configOptions, readFileErr := ioutil.ReadFile(configDir)
	if readFileErr != nil {
		fmt.Println("Error reading file", readFileErr)
	}

	var opts ConfigurationOptions
	err := json.Unmarshal(configOptions, &opts)

	return opts, err
}


