package main

import (
	"github.com/jjmschofield/GoCook/config"
	"github.com/jjmschofield/GoCook/server"
	"flag"
)

type cliArgs struct{
	port string
}

func main() {
	args := getCommandLineArgs()

	config.LoadNonSensitiveConfig()

	server.Start(args.port)
}

func getCommandLineArgs() cliArgs{
	port := flag.String("port", "8080", "The port to bind the webserver too")
	flag.Parse()

	return cliArgs{
		port: *port,
	}
}