package configuration

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
	"io"
)

const (
	AppPrefix = ""
)

// Config Main Configuration of a program containing server and database configuration
type Config struct {
	Server struct {
		Port    int `envconfig:"PORT" yaml:"port"`
		Timeout int `envconfig:"TIMEOUT" yaml:"timeout"`
	} `yaml:"server"`
	DataBaseURL string `envconfig:"DATABASE_URL" env-required:"true"`
}

// NewProgramConfig Constructor for program configuration
func NewProgramConfig(file io.Reader) (*Config, error) {
	cfg := Config{}
	err := envconfig.Process(AppPrefix, &cfg)
	if err != nil {
		return nil, fmt.Errorf("configuration error: %v", err)
	}

	if cfg.Server.Port == 0 || cfg.Server.Timeout == 0 {
		// Getting from other config, if not received from env
		yamlDecoder := yaml.NewDecoder(file)
		if err := yamlDecoder.Decode(&cfg); err != nil {
			return nil, fmt.Errorf("unkown yaml configuration recieved: %v", err)
		}
	}

	return &cfg, nil
}
