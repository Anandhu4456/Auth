package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBURL string

	JWTAccessSecrect          string
	JWTAccessExpiryTimeMinute int
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, fmt.Errorf("env file loading failed : %w", err)
	}

	cfg := &Config{
		DBURL:                     getFromEnv("DB_URL"),
		JWTAccessSecrect:          getFromEnv("JWT_ACCESS_SECRET"),
		JWTAccessExpiryTimeMinute: getEnvAsInt("JWT_ACCESS_EXPIRY_TIME_MINUTE"),
	}

	return cfg, nil
}

func getFromEnv(s string) string {
	v := os.Getenv(s)
	return v
}

func getEnvAsInt(s string) int {
	strV := os.Getenv(s)
	intEnv, err := strconv.Atoi(strV)
	if err != nil {
		log.Printf("env get as 'int' failed : %v", err)
	}

	return intEnv
}
