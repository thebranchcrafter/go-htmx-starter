package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds the application's configuration variables.
type Config struct {
	AddressPort string
	AppEnv      string
}

// LoadConfig loads environment variables from a .env file and populates the Config struct.
func LoadConfig() *Config {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return load()
}

// LoadTestConfig loads environment variables from a .env file and populates the Config struct.
func LoadTestConfig() *Config {
	// Load the .env file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return load()
}

func load() *Config {
	return &Config{
		AddressPort: getEnv("ADDRESS_PORT", ":8081"),
		AppEnv:      getEnv("APP_ENV", "test"),
	}
}

// getEnv gets an environment variable or returns a default value if not set.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
