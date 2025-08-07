package application

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	Environment string `mapstructure:"environment"`
	Service     struct {
		Name  string         `mapstructure:"name"`
		Ports map[string]int `mapstructure:"ports"`
	} `mapstructure:"service"`
}

func LoadConfig() (*AppConfig, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config AppConfig
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
