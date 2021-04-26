package config

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/go-yaml/yaml"
)

type ServiceConfig struct {
	IPPort   string `json:"ip_port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func LoadServiceConfig(fileName string) (*ServiceConfig, error) {
	var cfg ServiceConfig

	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	err = dec.Decode(&cfg)

	if cfg.IPPort == "" {
		return nil, errors.New("ip_port не может быть пустым")
	}
	if cfg.User == "" {
		return nil, errors.New("user не может быть пустым")
	}
	if cfg.Password == "" {
		return nil, errors.New("password не может быть пустым")
	}

	return &cfg, err
}

type DBConfig struct {
	Host     string `yaml:"host"`
	DB       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func LoadDBConfig(fileName string) (*DBConfig, error) {
	var cfg DBConfig

	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dec := yaml.NewDecoder(f)
	err = dec.Decode(&cfg)
	if err != nil {
		return nil, err
	}

	if cfg.Host == "" {
		return nil, errors.New("host не может быть пустым")
	}
	if cfg.DB == "" {
		return nil, errors.New("db не может быть пустым")
	}
	if cfg.User == "" {
		return nil, errors.New("user не может быть пустым")
	}
	if cfg.Password == "" {
		return nil, errors.New("password не может быть пустым")
	}

	return &cfg, nil
}
