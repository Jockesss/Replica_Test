package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type (
	Config struct {
		HTTP HTTP
	}

	HTTP struct {
		Host           string
		Port           string
		ReadTimeout    time.Duration
		WriteTimeout   time.Duration
		MaxHeaderBytes int
	}
)

func InitConfig() (*Config, error) {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	pathToEnv := filepath.Join(basepath, "../../.env")

	if err := godotenv.Load(pathToEnv); err != nil {
		log.Printf("Error loading .env file: %v", err)
		return nil, err
	}

	var cfg Config
	setFromEnv(&cfg)
	return &cfg, nil
}

func setFromEnv(cfg *Config) {
	cfg.HTTP.Host = os.Getenv("HTTP_HOST")
	cfg.HTTP.Port = os.Getenv("HTTP_PORT")
}
