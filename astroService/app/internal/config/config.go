package config

import (
	"errors"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

// Config structure of project config
type Config struct {
	ApiKey  string `yaml:"apiToken"`
	ApiHost string `yaml:"apiHost"`
}

var configFilePath = "app/internal/data/config.yml"
var (
	configReadError  = errors.New("Error of reading config file")
	configParseError = errors.New("Error of parsing config file")
)

// New - is the constructor of Config
func New() (*Config, error) {

	rawYml, err := os.ReadFile(configFilePath)

	if err != nil {
		log.Println(err)
		return nil, err
		//return nil, configReadError
	}
	var config Config
	err = yaml.Unmarshal(rawYml, &config)
	log.Print(config.ApiKey)
	if err != nil {
		return nil, configParseError
	}
	return &config, nil

}
