package config

import (
	"os"

	"metric-collection-service/internal/model"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Agent   Agent   `yaml:"agent"`
	Server  Server  `yaml:"server"`
	Storage Storage `yaml:"storage"`
}
type Agent struct {
	PollInterval   string `yaml:"pollInterval"`
	ReportInterval string `yaml:"reportInterval"`
}

type Server struct {
	Address string `yaml:"address"`
	Port    string `yaml:"port"`
}

type Storage struct {
	Type model.StorageType `yaml:"type"`
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
