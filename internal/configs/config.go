package configs

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"
)

type Config struct {
	Debug bool `json:"debug"`
	Port  int  `json:"port"`

	Database struct {
		DBName   string `json:"db_name"`
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		SSL      struct {
			Enabled  bool   `json:"enabled"`
			Cert     string `json:"cert"`
			Key      string `json:"key2"`
			RootCert string `json:"root_cert"`
		}
	} `json:"database"`
}

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

func Load(p string) (*Config, error) {
	cfgs, err := ReadFromFile(p)
	if errors.Is(err, fs.ErrNotExist) {
		cfgs = &DefaultConfig
	} else if err != nil {
		return nil, err
	}
	return cfgs, nil
}
