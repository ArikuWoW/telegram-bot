package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const configFile = "data/config.yaml"

type Config struct {
	Token string `yaml:"token"`
}

type Service struct {
	config Config
}

func New() (*Service, error) {
	s := &Service{}

	rawYAML, err := os.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("Reading config file: %w", err)
	}

	err = yaml.Unmarshal(rawYAML, &s.config)
	if err != nil {
		return nil, fmt.Errorf("Parsing yaml: %w", err)
	}

	return s, nil
}

func (s *Service) Token() string {
	return s.config.Token
}
