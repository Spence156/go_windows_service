package tools

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port int `yaml:"port"`
}

// Function loads configuration from the config.yml file and converts to variables.
func LoadConfig() (*Config, error) {
	file, err := os.Open("../config/config.yml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
