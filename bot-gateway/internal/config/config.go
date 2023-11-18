package config

import (
	"github.com/spf13/viper"
	"log/slog"
	"sync"
)

type Config struct {
	Port        string `yaml:"port" mapstructure:"PORT"`
	RegSvcUrl   string `yaml:"reg_svc_url" mapstructure:"REG_SVC_URL"`
	BooksSvcUrl string `yaml:"books_svc_url" mapstructure:"BOOKS_SVC_URL"`
	RentSvcUrl  string `yaml:"rent_svc_url" mapstructure:"RENT_SVC_URL"`
	Redis       string `yaml:"redis" mapstructure:"REDIS"`
	BotToken    string `yaml:"bot_token" mapstructure:"BOT_TOKEN"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	instance = &Config{}
	once.Do(func() {
		viper.AddConfigPath("../app/internal/config/envs")
		viper.SetConfigName("dev")
		viper.SetConfigType("env")

		viper.AutomaticEnv()

		err := viper.ReadInConfig()

		if err != nil {
			slog.Error(err.Error())
		}

		err = viper.Unmarshal(&instance)
	})
	return instance
}
