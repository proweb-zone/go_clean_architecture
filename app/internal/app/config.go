package app

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env string `yaml:"env" env-default:"local"`
	HTTPServer `yaml:"http_server"`
	Kafka `yaml:"kafka"`
	Db `yaml:"DB"`
}

type HTTPServer struct {
	Address string `yaml:"address" env-default:"localhost:8080"`
	Timeout time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type Kafka struct {
	Address string `yaml:"address" env-default:"localhost:29092"`
}

type Db struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	NameDb string `yaml:"name_db"`
	UserDb string `yaml:"user_db"`
	PasswordDb string `yaml:"password_db"`
}

func InitConfig() *Config {
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

func (h *HTTPServer) GetHTTPServer() *HTTPServer {
	return h
}

func (h *HTTPServer) GetAddressHttpServer() string {
	return h.Address
}

func (h *HTTPServer) GetTimeoutHttpServer() time.Duration {
	return h.Timeout
}

func (h *HTTPServer) GetIdleTimeoutHttpServer() time.Duration {
	return h.IdleTimeout
}

func (d *Db) GetConfigDb() *Db {
	return d
}

func (p *Db) GetHostDb() string {
	return p.Host
}

func (p *Db) GetPortDb() string {
	return p.Port
}

func (p *Db) GetNameDb() string {
	return p.NameDb
}

func (p *Db) GetUserDb() string {
	return p.UserDb
}

func (p *Db) GetPasswordDb() string {
	return p.GetPasswordDb()
}

func (k *Kafka) GetAddressKafka() string {
	return k.Address
}

func (k *Kafka) GetConfigKafka() *Kafka {
	return k
}

type IHTTPServer interface {
	GetAddressHttpServer() string
	GetTimeoutHttpServer() time.Duration
	GetIdleTimeoutHttpServer() time.Duration
}

type IDb interface {
	GetHostDb() string
	GetPortDb() string
	GetNameDb() string
	GetUserDb() string
	GetPasswordDb() string
}

type IKafka interface {
	GetAddressKafka() string
}

type IConfig interface {
	GetHTTPServer() IHTTPServer
	GetConfigDb() IDb
	GetConfigKafka() IKafka
}
