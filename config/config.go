package config

import (
	"errors"

	"github.com/spf13/viper"
)

var (
	ErrHTTPPortNotSet      = errors.New("HTTP_PORT is not set")
	ErrWeatherAPIKeyNotSet = errors.New("WEATHER_API_KEY is not set")
)

type Config struct {
	HTTPPort      string `mapstructure:"HTTP_PORT"`
	WeatherAPIKey string `mapstructure:"WEATHER_API_KEY"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	if config.HTTPPort == "" {
		return nil, ErrHTTPPortNotSet
	}

	if config.WeatherAPIKey == "" {
		return nil, ErrWeatherAPIKeyNotSet
	}

	return &config, nil
}
