package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	App struct {
		Host string `env:"CHAT_SERVICE_HOST" env-required:"true"`
		Port int    `env:"CHAT_SERVICE_PORT" env-required:"true"`
	}

	Logger struct {
		Level string `env:"LOGGER_LEVEL" env-required:"true"`
	}

	DB struct {
		User string `env:"DB_USER" env-required:"true"`
		Pass string `env:"DB_PASS" env-required:"true"`
		Host string `env:"DB_HOST" env-required:"true"`
		Port int    `env:"DB_PORT" env-required:"true"`
		Name string `env:"DB_NAME" env-required:"true"`
	}

	Redis struct {
		Host string `env:"REDIS_HOST" env-required:"true"`
		Port int    `env:"REDIS_PORT" env-required:"true"`
		Pass string `env:"REDIS_PASS" env-required:"true"`
		DB   int    `env:"REDIS_DB" env-required:"true"`
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
