package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	port         int `mapstructure:"port" validate:"required"`
	DatabasePath DatabaseConfig `mapstructure:"database_path" validate:"required"`
}

type DatabaseConfig struct {
	Path string `mapstructure:"path" validate:"required"`
}

func LoadConfig() *Config {
viper.SetConfigName("config") 
    viper.SetConfigType("yaml")   // Config file type
    viper.AddConfigPath(".")      // Look for the config file in the current directory

    if err := viper.ReadInConfig(); err != nil {
        panic(err)
    }

    var config Config
    if err := viper.Unmarshal(&config); err != nil {
        panic(err)
    }

    return &config
}