package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// GlobalConfig global config
type GlobalConfig struct {
	IfMergeHeader bool `yaml:"ifMergeHeader"`
}

// GConfig global configuration
var GConfig = &GlobalConfig{}

func Init(configPath string) {
	if configPath == "" {
		configPath = "config.yaml"
	}
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Printf("ReadConfig Error, errorMsg is %s\n", err.Error())
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, GConfig)
	if err != nil {
		fmt.Printf("Unmarshal Error, error message is %s\n", err.Error())
		panic(err)
	}
}
