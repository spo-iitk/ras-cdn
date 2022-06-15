package config

import (
	"log"

	"github.com/spf13/viper"
)

func viperConfig() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("../")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
		panic(err)
	}
	log.Println("Config file loaded successfully")
}
