package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type GRPCServicesConfig struct {
	AuthService struct {
		Host string `env:"AUTH_SERVICE_HOST"`
		Port int    `env:"AUTH_SERVICE_PORT"`
	}
	UniversityService struct {
		Host string `env:"UNIVERSITY_SERVICE_HOST"`
		Port int    `env:"UNIVERSITY_SERVICE_PORT"`
	}
	FileService struct {
		Host string `env:"FILE_SERVICE_HOST"`
		Port int    `env:"FILE_SERVICE_PORT"`
	}
}
type HTTPConfig struct {
	Port int `env:"HTTP_PORT"`
}

type Config struct {
	HTTP     HTTPConfig
	Services GRPCServicesConfig
}

func New() *Config {
	cfg := &Config{}
	if err := cleanenv.ReadEnv(cfg); err != nil {
		header := "gateway"
		f := cleanenv.FUsage(os.Stdout, cfg, &header)
		f()
		panic(err)
	}

	return cfg
}
