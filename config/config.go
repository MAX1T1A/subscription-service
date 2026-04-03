package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	AppPort  string
	DB       DBConfig
	Worker   WorkerConfig
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

func (c DBConfig) DSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.User, c.Password, c.Host, c.Port, c.Name, c.SSLMode,
	)
}

type WorkerConfig struct {
	Interval         time.Duration
	RenewalThreshold time.Duration
}

func Load() *Config {
	return &Config{
		AppPort: getEnv("APP_PORT", "8080"),
		DB: DBConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			Name:     getEnv("DB_NAME", "subscriptions"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		Worker: WorkerConfig{
			Interval:         time.Duration(getEnvInt("WORKER_INTERVAL_SECONDS", 30)) * time.Second,
			RenewalThreshold: time.Duration(getEnvInt("WORKER_RENEWAL_THRESHOLD_SECONDS", 60)) * time.Second,
		},
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return fallback
	}
	return n
}
