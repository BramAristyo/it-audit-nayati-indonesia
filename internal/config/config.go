package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Config struct {
	APIUrl string `json:"api_url"`
}

func GetConfig() (*Config, error) {
	file, err := os.ReadFile("config.json")
	if err != nil {
		return nil, fmt.Errorf("config.json not found: %v", err)
	}

	var cfg Config
	err = json.Unmarshal(file, &cfg)
	if err != nil {
		return nil, fmt.Errorf("invalid config.json format: %v", err)
	}

	if err = cfg.Validate(); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (cfg *Config) Validate() error {
	if cfg.APIUrl == "" {
		return errors.New("config: api_url is required")
	}

	return nil
}
