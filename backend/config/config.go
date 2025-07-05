package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	DBHost        string
	DBUser        string
	DBPassword    string
	DBName        string
	DBPort        int
	ServerPort    int
	TogetherAIKey string
	JWTSecret     string
}

func Load() *Config {
	dbPortStr := os.Getenv("DB_PORT")
	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		log.Fatalf("Invalid DB_PORT: %s", dbPortStr)
	}

	portStr := os.Getenv("PORT")
	serverPort := 8080
	if portStr != "" {
		p, err := strconv.Atoi(portStr)
		if err != nil {
			log.Fatalf("Invalid PORT: %s", portStr)
		}
		serverPort = p
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable must be set")
	}

	apiKey := os.Getenv("TOGETHER_API_KEY")
	if apiKey == "" {
		log.Fatal("TOGETHER_API_KEY environment variable must be set")
	}

	return &Config{
		DBHost:        os.Getenv("DB_HOST"),
		DBUser:        os.Getenv("DB_USER"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		DBName:        os.Getenv("DB_NAME"),
		DBPort:        dbPort,
		ServerPort:    serverPort,
		TogetherAIKey: apiKey,
		JWTSecret:     jwtSecret,
	}
}
