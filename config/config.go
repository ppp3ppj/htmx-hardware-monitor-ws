package config

import (
	"strings"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

var once sync.Once

type (
    AppInfo struct {
        Name        string `mapstructure:"name" validate:"required"`
        Version     string `mapstructure:"version" validate:"required"`
        Env string `mapstructure:"environtment" validate:"required"`
        SecretKey   string `mapstructure:"secretkey" validate:"required"`
    }

	Server struct {
		Port         int           `mapstructure:"port" validate:"required"`
		AllowOrigins []string      `mapstructure:"allowOrigins" validate:"required"`
		Timeout      time.Duration `mapstructure:"timeout" validate:"required"`
		BodyLimit    string        `mapstructure:"bodyLimit" validate:"required"`
	}

	Config struct {
		Server *Server   `mapstructure:"server" validate:"required"`
        AppInfo *AppInfo      `mapstructure:"appinfo" validate:"required"`
	}
)

var configInstance *Config

func ConfigGetting() *Config {
	once.Do(func() {
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		if err := viper.Unmarshal(&configInstance); err != nil {
			panic(err)
		}

		validate := validator.New()

		if err := validate.Struct(configInstance); err != nil {
			panic(err)
		}
	})

	return configInstance
}
