package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port              string
	DatabaseUrl       string
	KeycloakIssuerURL string
	KeycloakClientID  string
}

func Load() *Config {
	godotenv.Load()

	return &Config{
		Port:              getEnv("PORT", "8080"),
		DatabaseUrl:       mustGetEnv("DATABASE_URL"),
		KeycloakIssuerURL: mustGetEnv("KEYCLOAK_ISSUER_URL"),
		KeycloakClientID:  mustGetEnv("KEYCLOAK_CLIENT_ID"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func mustGetEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		panic(fmt.Sprintf("required environment variable %s not set", key))
	}
	return value
}
