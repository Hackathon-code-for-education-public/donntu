package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	App struct {
		Host string `env:"APP_HOST" env-required:"true"`
		Port int    `env:"APP_PORT" env-required:"true"`
	}

	Logger struct {
		Level string `env:"LOGGER_LEVEL" env-required:"true"`
	}

	Minio struct {
		Endpoint  string `env:"MINIO_ENDPOINT" env-required:"true"`
		AccessKey string `env:"MINIO_ACCESS_KEY" env-required:"true"`
		SecretKey string `env:"MINIO_SECRET_KEY" env-required:"true"`
		Db        int    `env:"MINIO_DB" env-required:"true"`
		Bucket    string `env:"MINIO_BUCKET" env-required:"true"`
		Secure    bool   `env:"MINIO_SECURE" env-required:"true"`
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
