package config

import (
	"fmt"
	"github.com/jjmschofield/GoCook/common/logger"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func LoadNonSensitiveConfig() {
	logger.Info("Loading non-sensitive config from ./cook.json")

	setDefaultConfig()

	viper.SetConfigName("cook")
	viper.AddConfigPath("./")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		logger.Fatal(fmt.Sprintf("Fatal error config file: %s \n", err), zap.Error(err))
	}
}

func setDefaultConfig() {
	viper.SetDefault("AUTH_JWKS_ENDPOINT", "AUTH_JWKS_ENDPOINT_NOT_SET")
	viper.SetDefault("AUTH_ISSUER", "AUTH_ISSUER_NOT_SET")
	viper.SetDefault("AUTH_AUDIENCE", "AUTH_AUDIENCE_NOT_SET")
}
