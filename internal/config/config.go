package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
	CurrentUserID   string `json:"current_user_id"`
}

const (
	filename = ".gatorconfig.json"
)

func Read() (Config, error) {
	cfg := Config{}

	cfgFile, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	cfgData, err := os.ReadFile(cfgFile)
	if err != nil {
		return Config{}, err
	}

	if err := json.Unmarshal(cfgData, &cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func SetUser(name string, id string) error {
	cfg, err := Read()
	if err != nil {
		return err
	}

	cfg.CurrentUserName = name
	cfg.CurrentUserID = id
	return write(cfg)
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, filename), nil
}

func write(cfg Config) error {
	cfgFile, err := getConfigFilePath()
	if err != nil {
		return err
	}

	cfgData, err := json.Marshal(&cfg)
	if err != nil {
		return err
	}

	if err := os.WriteFile(cfgFile, cfgData, 0644); err != nil {
		return err
	}

	return nil
}
