package configuration

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
)

// Config Main Configuration of a program containing server and database configuration
// TODO: Default configuration support
// TODO: Docker env configuration
// TODO: Check for necessary configuration, if not replace with default
type Config struct {
	Server struct {
		Port    int `yaml:"port"`
		Timeout int `yaml:"timeout"`
	} `yaml:"server"`
	Database struct {
		URL      string `yaml:"url"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"database"`
}

// NewProgramConfig Constructor for program configuration
func NewProgramConfig(file io.Reader) (*Config, error) {
	cfg := Config{}
	yamlDecoder := yaml.NewDecoder(file)

	if err := yamlDecoder.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("unkown yaml configuration recieved: %v", err)
	}

	return &cfg, nil
}
