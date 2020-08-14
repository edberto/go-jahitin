package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type IConfig interface {
	GetString(key string) string
	GetInt(key string) int
}

type Config struct {
	*viper.Viper
}

func NewConfig(filepath string) IConfig {
	viper.SetConfigFile(filepath)
	if err := viper.ReadInConfig(); err != nil {
		panic(errors.Wrap(err, "Failed to read config"))
	}
	return &Config{
		viper.GetViper(),
	}
}

func (c *Config) GetString(key string) string {
	return c.Viper.GetString(key)
}

func (c *Config) GetInt(key string) int {
	return c.Viper.GetInt(key)
}
