package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DB_URL            string `json:"db_url"`
	CURRENT_USER_NAME string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Error getting user home directory")
	}
	return homeDir + "/" + configFileName, nil
}

func writeConfigFile(c *Config) error {
	jsonData, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("Error marshalling config JSON")
	}
	jsonPath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	if err := os.WriteFile(jsonPath, jsonData, 0644); err != nil {
		return fmt.Errorf("Error writing config JSON file")
	}
	return nil
}

func (c *Config) SetUser(username string) error {
	c.CURRENT_USER_NAME = username
	if err := writeConfigFile(c); err != nil {
		return err
	}
	return nil
}
