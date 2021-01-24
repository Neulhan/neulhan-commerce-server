package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

func GetEnv(key string) string {
	if os.Getenv("GIN_MODE") == "release" {
		viper.SetConfigFile("prd.env")
	} else {
		viper.SetConfigFile("dev.env")
	}

	err := viper.ReadInConfig()

	if err != nil {
		log.Printf("Error while reading config file %s", err)
	}

	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatalf("Invalid type assertion %s", value)
	}

	return value
}
