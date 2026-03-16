package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBURL string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, fmt.Errorf("env file loading failed : %w", err)
	}

	cfg := &Config{
		DBURL: getFromEnv("DB_URL"),
	}

	return cfg, nil
}

func getFromEnv(s string) string {
	v := os.Getenv(s)
	return v
}
