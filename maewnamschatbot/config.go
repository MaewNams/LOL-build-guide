package maewnamschatbot

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Config contains configured data read from .config.yml
type Config struct {
	Endpoints Endpoints     `yaml:"endpoints"`
	Message   MessageConfig `yaml:"message"`
}

// Endpoints is list of endpoints
type Endpoints struct {

}

// MessageConfig is set of message-related configuration
type MessageConfig struct {
	DefaultMsg       string `yaml:"default_msg"`
}

// ReadConfig reads configuration from the given file.
func ReadConfig(file string) (*Config, error) {
	configData, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		return nil, err
	}

	log.Println("config.message", config.Message)

	return &config, nil
}
