package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Configuration struct {
	DbHost              string `mapstructure:"DB_HOST"`
	DbName              string `mapstructure:"DB_NAME"`
	DbUser              string `mapstructure:"DB_USER"`
	DbPort              string `mapstructure:"DB_PORT"`
	DbPassword          string `mapstructure:"DB_PASSWORD"`
	FirebaseCredentials string `mapstructure:"FIREBASE_CREDENTIALS"`
	ServerPort          string `mapstructure:"SERVER_PORT"`
}

var envs = []string{
	"DB_HOST",
	"DB_NAME",
	"DB_USER",
	"DB_PORT",
	"DB_PASSWORD",
	"FIREBASE_CREDENTIALS",
	"SERVER_PORT",
}

func LoadConfiguration() (Configuration, error) {
	var configuration Configuration

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		return Configuration{}, err
	}

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return configuration, err
		}
	}

	if err := viper.Unmarshal(&configuration); err != nil {
		return configuration, err
	}

	if err := validator.New().Struct(&configuration); err != nil {
		return configuration, err
	}

	return configuration, nil
}
