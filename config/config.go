package config

import "os"

// Config menampung semua konfigurasi aplikasi.
type Config struct {
	DatabaseURL  string
	JWTSecretKey string // <-- TAMBAHKAN INI
}

// LoadConfig memuat konfigurasi dari environment variables.
func LoadConfig() (*Config, error) {
	return &Config{
		DatabaseURL:  os.Getenv("DATABASE_URL"),
		JWTSecretKey: os.Getenv("JWT_SECRET_KEY"), // <-- TAMBAHKAN INI
	}, nil
}