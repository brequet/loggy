package config

import (
	"embed"
	"fmt"
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

//go:embed resources/default-conf.yml
var defaultConfigFile embed.FS

type (
	Parser struct {
		Formats    []ParserLogFormat `yaml:"formats"`
		AppLogDirs []AppLogDir       `yaml:"app-log-dirs"`
	}

	ParserLogFormat struct {
		Name        string `yaml:"Name"`
		DateFormat  string `yaml:"DateFormat"`
		RegexParser string `yaml:"RegexParser"`
	}

	AppLogDir struct {
		Name string `yaml:"name"`
		Path string `yaml:"path"`
	}

	Config struct {
		Server struct {
			Port int `yaml:"port"`
		} `yaml:"server"`

		Parser Parser `yaml:"parser"`
	}
)

func LoadConfig(overloadConfigFile string, logger *slog.Logger) (*Config, error) {
	var cfg Config

	var data []byte
	var err error
	if overloadConfigFile != "" {
		logger.Debug("Overloading config file", "overloadConfigFile", overloadConfigFile)
		data, err = os.ReadFile(overloadConfigFile)
	} else {
		data, err = defaultConfigFile.ReadFile("resources/default-conf.yml")
	}
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	return &cfg, nil
}
