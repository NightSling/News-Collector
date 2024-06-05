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
	ApiKey    string   `yaml:"api-key"`
	Sources   []string `yaml:"sources"`
	Scheduler struct {
		Interval int    `yaml:"interval"`
		LastRan  string `yaml:"last-ran"`
		Enabled  bool   `yaml:"enable"`
	} `yaml:"scheduler"`
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

func SaveConfig(configPath string, config *Config) error {
	file, err := os.OpenFile(configPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		// if the file does not exist, create it
		file, err = os.Create(configPath)
		if err != nil {
			return err
		}
	}
	defer file.Close()
	// write to a string stream
	d := yaml.NewEncoder(file)
	if err := d.Encode(config); err != nil {
		return err
	}
	return nil

}
