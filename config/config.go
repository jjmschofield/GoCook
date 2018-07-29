package config

import (
	"github.com/spf13/viper"
	"fmt"
)

func LoadNonSensitiveConfig(){
	setDefaultConfig()

	viper.SetConfigName("cook")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func setDefaultConfig(){
	viper.SetDefault("HTTP_PORT", "8080")
	viper.SetDefault("AUTH_JWKS_ENDPOINT", "AUTH_JWKS_ENDPOINT_NOT_SET")
	viper.SetDefault("AUTH_ISSUER", "AUTH_ISSUER_NOT_SET")
	viper.SetDefault("AUTH_AUDIENCE", "AUTH_AUDIENCE_NOT_SET")
}