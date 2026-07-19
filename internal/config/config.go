package config

import (
	"path/filepath"
	"os"
	"encoding/json"
)

type Config struct {
	DbUrl string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

//Function to get the filepath
func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	fullPath := filepath.Join(home, ".gatorconfig.json")
	
	return fullPath, nil
}

//Function to read the filepath
func Read() (Config, error) {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(fullPath)
	if err != nil {
		return Config{}, err
	}

	var configuration Config
	err = json.Unmarshal(data, &configuration)
	if err != nil {
		return Config{}, err
	}

	return configuration, nil
}

//Set user in the Config struct
func (s *Config) SetUser(username string) error {
	s.CurrentUserName = username
	fullPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(s)
	if err != nil {
		return err
	}

	err = os.WriteFile(fullPath, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}