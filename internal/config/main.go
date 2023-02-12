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
	if err := viper.BindPFlags(flag.CommandLine); err != nil {
		log.Fatalln(err)
	}
	configFile := viper.GetString("config")
	if err := viper.BindEnv("environment"); err != nil {
		log.Fatalln(err)
	}
	viper.SetConfigFile(configFile)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalln("Config File not found")
		} else {
			log.Panicln(err)
		}
	}
	return *viper.GetViper()
}
