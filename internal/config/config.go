package config

import (
	"github.com/spf13/viper"
)

// Config holds application config
type Config struct {
	Port int    `mapstructure:"port"`
	DSN  string `mapstructure:"dsn"`
}

// Load configuration
func Load() (*Config, error) {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}

	return &c, nil
}
