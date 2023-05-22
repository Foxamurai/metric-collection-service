package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Agent  Agent  `yaml:"agent"`
	Server Server `yaml:"server"`
}
type Agent struct {
	PollInterval   string `yaml:"pollInterval"`
	ReportInterval string `yaml:"reportInterval"`
}

type Server struct {
	Address string `yaml:"address"`
	Port    string `yaml:"port"`
}

var config *Config

func Init(cfgPath string) (*Config, error) {
	if config != nil {
		return config, nil
	}

	yamlFile, err := os.ReadFile(cfgPath)
	if err != nil {
		return nil, err
	}

	config = &Config{}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func Get() *Config {
	if config == nil {
		config = &Config{}
	}
	return config
}
