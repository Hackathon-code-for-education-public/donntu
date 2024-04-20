package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	App struct {
		Host string `env:"VERIFICATION_SERVICE_HOST" env-required:"true"`
		Port int    `env:"VERIFICATION_SERVICE_PORT" env-required:"true"`
	}

	Logger struct {
		Level string `env:"LOGGER_LEVEL" env-required:"true"`
	}

	DB struct {
		User string `env:"VERIFICATION_SERVICE_DB_USER" env-required:"true"`
		Pass string `env:"VERIFICATION_SERVICE_DB_PASS" env-required:"true"`
		Host string `env:"DB_HOST" env-required:"true"`
		Port int    `env:"DB_PORT" env-required:"true"`
		Name string `env:"DB_NAME" env-required:"true"`
	}

	AuthService struct {
		Host string `env:"AUTH_SERVICE_HOST" env-required:"true"`
		Port int    `env:"AUTH_SERVICE_PORT" env-required:"true"`
	}
}

func New() *Config {
	config := &Config{}

	if err := cleanenv.ReadEnv(config); err != nil {
		header := "FILE SERVICE ENVs"
		f := cleanenv.FUsage(os.Stdout, config, &header)
		f()
		panic(err)
	}

	log.Println(config)

	return config
}
