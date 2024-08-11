package config

import (
	"embed"
	"fmt"

	"gopkg.in/yaml.v3"
)

//go:embed resources/conf.yml
var defaultConfigFile embed.FS

type ParserFormat struct {
	Name        string `yaml:"Name"`
	DateFormat  string `yaml:"DateFormat"`
	RegexParser string `yaml:"RegexParser"`
}

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`

	Parser struct {
		Formats []ParserFormat `yaml:"formats"`
	} `yaml:"parser"`
}

func LoadConfig() (*Config, error) {
	var cfg Config

	data, err := defaultConfigFile.ReadFile("resources/conf.yml")
	if err != nil {
		return nil, fmt.Errorf("error reading embedded config file: %w", err)
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	return &cfg, nil
}
