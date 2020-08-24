package config

import "os"

type DbConfig struct {
	Username string
	Password string
	Host     string
	Schema   string
}

type Config struct {
	Database DbConfig
}

// Return new config struct
func New() *Config {
	return &Config{
		Database:
		DbConfig{
			Username: getEnv("MYSQL_USERS_USERNAME", ""),
			Password: getEnv("MYSQL_USERS_PASSWORD", ""),
			Host:     getEnv("MYSQL_USERS_HOST", ""),
			Schema:   getEnv("MYSQL_USERS_SCHEMA", ""),
		},
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
