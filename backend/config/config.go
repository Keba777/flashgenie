package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     int
	ServerPort int
	OpenAIKey  string
	JWTSecret  string
}

func Load() *Config {
	portStr := os.Getenv("DB_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Invalid DB_PORT: %s", portStr)
	}

	serverPortStr := os.Getenv("SERVER_PORT")
	serverPort := 8080 // default
	if serverPortStr != "" {
		sp, err := strconv.Atoi(serverPortStr)
		if err != nil {
			log.Fatalf("Invalid SERVER_PORT: %s", serverPortStr)
		}
		serverPort = sp
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable must be set")
	}

	return &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     port,
		ServerPort: serverPort,
		OpenAIKey:  os.Getenv("OPENAI_API_KEY"),
		JWTSecret:  jwtSecret,
	}
}
