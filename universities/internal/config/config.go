package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type GRPCConfig struct {
	Host string `env:"UNIVERSITIES_SERVICE_HOST"`
	Port string `env:"UNIVERSITIES_SERVICE_PORT"`
}

type DatabaseConfig struct {
	Host string `env:"UNIVERSITIES_DATABASE_HOST"`
	Port string `env:"UNIVERSITIES_DATABASE_PORT"`
	User string `env:"UNIVERSITIES_DATABASE_USER"`
	Pass string `env:"UNIVERSITIES_DATABASE_PASS"`
	Name string `env:"UNIVERSITIES_DATABASE_NAME"`
}

type Config struct {
	GRPC     GRPCConfig
	Database DatabaseConfig
}

func New() *Config {
	cfg := &Config{}
	if err := cleanenv.ReadEnv(cfg); err != nil {
		header := "universities"
		f := cleanenv.FUsage(os.Stdout, cfg, &header)
		f()
		panic(err)
	}

	return cfg
}
