package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func Read() Config {
	jsonPath, err := getConfigFilePath()
	if err != nil {
		fmt.Println("Error getting config file path")
		return Config{}
	}

	jsonData, err := os.ReadFile(jsonPath)
	if err != nil {
		fmt.Println("Error reading config JSON file")
		return Config{}
	}

	config := Config{}
	if err := json.Unmarshal(jsonData, &config); err != nil {
		fmt.Println("Error parsing config JSON file")
		return Config{}
	}

	return config
}
