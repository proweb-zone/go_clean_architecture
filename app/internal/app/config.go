package app

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `yaml:"env" env-default:"local"`
	HTTPServer `yaml:"http_server"`
	Kafka `yaml:"kafka"`
	Postgresql `yaml:"postgresql"`
}

type HTTPServer struct {
	Address string `yaml:"address" env-default:"localhost:8080"`
	Timeout time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type Kafka struct {
	Address string `yaml:"address" env-default:"localhost:29092"`
}

type Postgresql struct {
	HOST string `yaml:"host"`
	POSTGRES_DB string `yaml:"postgres_db"`
	POSTGRES_USER string `yaml:"postgres_user"`
	POSTGRES_PASSWORD string `yaml:"postgres_password"`
}

func MustLoad() *Config {
	var configPath string = "./local.yaml"

	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
