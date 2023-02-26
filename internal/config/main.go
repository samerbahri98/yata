package config

import (
	"context"
	"log"

	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var configFile string

func Load(ctx context.Context) (*viper.Viper, error) {

	// Defaults
	v := viper.New()
	v.SetDefault("environment", "production")
	v.SetDefault("config", "/etc/yata/config.toml")

	// Parse Config
	flag.StringVarP(&configFile, "config", "c", "/etc/yata/config.toml", "Config File Path")
	flag.Parse()
	if err := viper.BindPFlags(flag.CommandLine); err != nil {
		return nil, err
	}

	configFile := v.GetString("config")
	if err := v.BindEnv("environment"); err != nil {
		return nil, err
	}

	v.SetConfigFile(configFile)
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalln("Config File not found")
		} else {
			return nil, err
		}
	}
	return v, nil
}
