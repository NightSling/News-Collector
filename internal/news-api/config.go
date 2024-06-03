package newsapi

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"pwd"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Database string `yaml:"database"`
		Table    string `yaml:"table"`
	} `yaml:"database"`
	ApiKey  string   `yaml:"api-key"`
	Sources []string `yaml:"sources"`
}

func NewConfig(configPath string) (*Config, error) {
	c := &Config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&c); err != nil {
		return nil, err
	}

	return c, nil
}
