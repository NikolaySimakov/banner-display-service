package config

import (
	"fmt"
	"log"
	"path"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		App    `yaml:"app"`
		HTTP   `yaml:"http"`
		Log    `yaml:"log"`
		PG     `yaml:"postgres"`
		Secure `yaml:"secure"`
	}

	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	Log struct {
		Level string `env-required:"true" yaml:"level" env:"LOG_LEVEL"`
	}

	PG struct {
		MaxPoolSize int    `env-required:"true" yaml:"max_pool_size" env:"PG_MAX_POOL_SIZE"`
		URL         string `env-required:"true"                      env:"PG_URL"`
	}

	Secure struct {
		Salt string `env-required:"true" env:"HASHER_SALT"`
	}
)

func NewConfig(configPath string) (*Config, error) {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("No .env file found")
	}
	
	cfg := &Config{}

	err = cleanenv.ReadConfig(path.Join("./", configPath), cfg)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	err = cleanenv.UpdateEnv(cfg)
	if err != nil {
		return nil, fmt.Errorf("error updating env: %w", err)
	}

	return cfg, nil
}