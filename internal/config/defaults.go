package config

import "github.com/spf13/viper"

func setDefault() {
	viper.SetDefault("environment", "production")
	viper.SetDefault("config", "/etc/yata/config.toml")
}
