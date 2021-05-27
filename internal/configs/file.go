package configs

import (
	"encoding/json"
	"os"
)

func ReadFromFile(p string) (*Config, error) {
	raw, err := os.ReadFile(p)
	if err != nil {
		return nil, err
	}

	var cfg Config

	err = json.Unmarshal(raw, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
