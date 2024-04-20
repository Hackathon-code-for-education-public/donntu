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
