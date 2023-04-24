package config

import (
	"log"
	"os"
)

type Config struct {
	Port       string
	DBUser     string
	DBPassword string
	DBName     string
}

var config Config

func LoadConfig() {
	config = Config{
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_USER", "user1"),
		DBPassword: getEnv("DB_PASSWORD", "password1"),
		DBName:     getEnv("DB_NAME", "thesaurus"),
	}
}

func GetConfig() Config {
	return config
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Printf("Environment variable %s not found, using default value: %s", key, fallback)
		return fallback
	}
	return value
}
