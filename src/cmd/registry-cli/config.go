package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func configDefaults() {
}

func configInit() {
	logrus.SetOutput(os.Stderr)
	if viper.GetBool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	logrus.Debug("+++configInit()")
	viper.SetConfigName("registry-cli")
	viper.SetConfigType("toml")
	viper.AddConfigPath("$HOME/.config")
	viper.AddConfigPath("/etc/registry-cli")
	viper.AddConfigPath("/etc/docker")
	viper.AddConfigPath(".")

	configDefaults()

	viper.ReadInConfig()
}
