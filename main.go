package main

import (
	"github.com/jjmschofield/GoCook/config"
	"github.com/jjmschofield/GoCook/server"
)

func main() {
	config.LoadNonSensitiveConfig()
	server.Start()
}
