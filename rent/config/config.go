package config

import "github.com/spf13/viper"

type Config struct {
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
	}
	RabbitMQ struct {
		URL string
	}
	JWT struct {
		Secret string
	}
}

func LoadConfig() (*Config, error) {
	var config Config

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("config/")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
