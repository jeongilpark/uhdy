package config

import (
	"log"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
}

type Config struct {
	Database DatabaseConfig `mapstructure:"database"`

	Server struct {
		Port string `mapstructure:"port"`
	} `mapstructure:"server"`

	JWT struct {
		Secret string `mapstructure:"secret"`
	} `mapstructure:"jwt"`
}

var AppConfig Config

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// Set the path where viper will look for the config file
	viper.AddConfigPath(".")

	// Enable viper to read Environment Variables
	viper.AutomaticEnv()

	// Attempt to read the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// Unmarshal the config into the struct
	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
}
