package config

import (
	"github.com/spf13/viper"
	"log/slog"
	"sync"
)

type Config struct {
	Port      string `yaml:"port" mapstructure:"PORT"`
	RegSvcUrl string `yaml:"reg_svc_url" mapstructure:"REG_SVC_URL"`
	BotToken  string `yaml:"bot_token" mapstructure:"BOT_TOKEN"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	instance = &Config{}
	once.Do(func() {
		viper.AddConfigPath("./internal/config/envs")
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
