package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"
)

type Config struct {
	NumberOfBackups int    `json:"number_of_backups"`
	BackupDir       string `json:"backup_dir"`
}

func GetConfigPath() string {
	configDir, _ := os.UserConfigDir()
	return filepath.Join(configDir, "autosync", "config.json")
}

func InitConfigDir() error {
	return os.MkdirAll(filepath.Dir(GetConfigPath()), 0755)
}

func Load() (Config, error) {
	configPath := GetConfigPath()
	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return Config{NumberOfBackups: 5}, nil // Default value
		}
		return Config{}, err
	}

	var cfg Config
	return cfg, json.Unmarshal(data, &cfg)
}

func Save(cfg Config) error {
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(GetConfigPath(), data, 0644)
}

func ParseBackups(text string) int {
	num, err := strconv.Atoi(text)
	if err != nil {
		return 5 // Default if conversion fails
	}
	return num
}
