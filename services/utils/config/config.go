package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type DefaultConfig struct {
	Database struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Name     string `mapstructure:"name"`
	} `mapstructure:"database"`

	Server struct {
		Port string `mapstructure:"port"`
	} `mapstructure:"server"`
}

var DefaultCfg DefaultConfig

func ReadConfig(cfgName string, cfgType string, cfgPath string, rawVal any) {
	viper.SetConfigName(cfgName)
	viper.SetConfigType(cfgType)

	viper.AddConfigPath(cfgPath)

	// Enable viper to read Environment Variables
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Attempt to read the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// Unmarshal the config into the struct
	if err := viper.Unmarshal(rawVal); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	log.Println(rawVal)
}

func init() {
	ReadConfig("default", "yaml", "/etc/uhdy", &DefaultCfg)
}
