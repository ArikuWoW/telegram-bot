package config

import (
	"errors"
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
		return nil, errors.New("reading file")
	}

	err = yaml.Unmarshal(rawYAML, &s.config)
	if err != nil {
		return nil, errors.New("parsing yaml")
	}

	return s, nil
}

func (s *Service) Token() string {
	return s.config.Token
}
