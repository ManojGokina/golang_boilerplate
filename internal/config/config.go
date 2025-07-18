package config

import (
	"log"
	"os"
	"strconv"
	"github.com/joho/godotenv"
)

type Config struct {
	Environment  string `mapstructure:"ENVIRONMENT"`
	Port         string `mapstructure:"PORT"`
	MongoURL     string `mapstructure:"MONGO_URL"`
	MongoDB      string `mapstructure:"MONGO_DB"`
	RedisURL     string `mapstructure:"REDIS_URL"`
	JWTSecret    string `mapstructure:"JWT_SECRET"`
	LogLevel     string `mapstructure:"LOG_LEVEL"`
	RateLimitRPS int    `mapstructure:"RATE_LIMIT_RPS"`
}

func Load() *Config {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	config := &Config{
		Environment:  getEnv("ENVIRONMENT", "development"),
		Port:         getEnv("PORT", "8080"),
		MongoURL:     getEnv("MONGO_URL", ""),
		MongoDB:      getEnv("MONGO_DB", "backend"),
		RedisURL:     getEnv("REDIS_URL", ""),
		JWTSecret:    getEnv("JWT_SECRET", ""),
		LogLevel:     getEnv("LOG_LEVEL", "info"),
		RateLimitRPS: getEnvAsInt("RATE_LIMIT_RPS", 100),
	}

	// Validate required fields
	if config.MongoURL == "" {
		log.Fatal("MONGO_URL is required")
	}
	if config.JWTSecret == "" {
		log.Fatal("JWT_SECRET is required")
	}

	return config
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}