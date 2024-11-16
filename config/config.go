package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/InosRahul/student-crud-api/utils"
	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	Port       string
}

func LoadConfig() (*Config, error) {
	var configPath string

	// Find the current file path
	_, filename, _, _ := runtime.Caller(0)
	currentFileDir := filepath.Dir(filename)

	// Walk up the directory tree to find the .env file
	for {
		configPath = filepath.Join(currentFileDir, ".env")
		if _, err := os.Stat(configPath); !os.IsNotExist(err) {
			break
		}
		parentDir := filepath.Dir(currentFileDir)
		if parentDir == currentFileDir {
			return nil, fmt.Errorf(".env file not found")
		}
		currentFileDir = parentDir
	}

	err := godotenv.Load(configPath)
	if err != nil {
		utils.Logger.Fatalf("Error loading .env file")
		return nil, err
	}

	return &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		Port:       os.Getenv("PORT"),
	}, nil
}

func (c *Config) GetDBConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)
}
