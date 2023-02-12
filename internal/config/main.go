package config

import (
	"log"

	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var configFile string

func Load() (c viper.Viper) {
	setDefault()
	flag.StringVarP(&configFile, "config", "c", "/etc/yata/config.toml", "Config File Path")
	flag.Parse()
	viper.BindPFlags(flag.CommandLine)
	configFile := viper.GetString("config")
	viper.BindEnv("environment")
	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalln("Config File not found")
		} else {
			log.Panicln(err)
		}
	}
	return *viper.GetViper()
}
