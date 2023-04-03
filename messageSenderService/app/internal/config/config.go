package config

import (
	"errors"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	MessageSenderMail     string `yaml:"notificationsEmail"`
	MessageSenderPassword string `yaml:"notificationsEmailPassword"`
}

var (
	InitConfigError       = errors.New("error when init the config")
	UnmarshallConfigError = errors.New("error when unmarshall the config")
	configPath            = "messageSenderService/app/internal/data/config.yml"
)

func Init() (*Config, error) {

	rawYml, err := os.ReadFile(configPath)

	if err != nil {
		return nil, InitConfigError
	}
	var config Config

	err = yaml.Unmarshal(rawYml, &config)
	if err != nil {
		return nil, UnmarshallConfigError
	}
	return &config, nil
}

func ChangeSenderUser(senderMail, senderPassword string) error {
	//TODO finish
	return nil
}
