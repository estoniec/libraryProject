package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type PostgreSQL struct {
	Username string `yaml:"psql_username" env:"PSQL_USERNAME"`
	Password string `yaml:"psql_password" env:"PSQL_PASSWORD"`
	Host     string `yaml:"psql_host" env:"PSQL_HOST"`
	Port     string `yaml:"psql_port" env:"PSQL_PORT"`
	Database string `yaml:"psql_database" env:"PSQL_DATABASE"`
}

type Config struct {
	Port     string     `yaml:"grpc_port" env:"GRPC_PORT"`
	Database PostgreSQL `yaml:"database"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}

		if err := cleanenv.ReadConfig("./config.yaml", instance); err != nil {
			var text string
			help, _ := cleanenv.GetDescription(instance, &text)
			log.Print(help)
			log.Fatal(err)
		}
	})
	return instance
}
