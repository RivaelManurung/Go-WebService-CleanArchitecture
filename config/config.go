package config

import "os"

type Config struct {
	DatabaseURL  string
	JWTSecretKey string
	ServerPort   string
}

func LoadConfig() (*Config, error) {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = ":8080" // Default port
	}

	return &Config{
		DatabaseURL:  os.Getenv("DATABASE_URL"),
		JWTSecretKey: os.Getenv("JWT_SECRET_KEY"),
		ServerPort:   port,
	}, nil
}