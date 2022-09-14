package config

import (
	"os"
)

// Config holds all the application configurations
type Config struct {
	FileReader *JSONReaderConfig
	Redis      *RedisConfig
}

// JSONReaderConfig holds the configurations of the json file reader
type JSONReaderConfig struct {
	FilePath string
}

// RedisConfig holds the redis database configuration
type RedisConfig struct {
	Host     string
	Password string
	Port     string
}

// NewConfig returns the application configuration filled
func NewConfig() *Config {
	return &Config{
		FileReader: &JSONReaderConfig{
			FilePath: os.Getenv("JSON_FILE_PATH"),
		},
		Redis: &RedisConfig{
			Host:     os.Getenv("REDIS_HOST"),
			Password: os.Getenv("REDIS_PASSWORD"),
			Port:     os.Getenv("REDIS_PORT"),
		},
	}
}
