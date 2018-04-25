package main

import (
	"github.com/spf13/viper"
	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
)

func defineFlags() {
	flag.CommandLine.String("token", "", "Bot Token")
}

func loadConfig() {
	defineFlags()

	viper.BindPFlags(flag.CommandLine)

	viper.SetEnvPrefix("narcisse")
	viper.AutomaticEnv()

	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		log.WithField("error", err).Fatal("error reading the config file")
	}
}