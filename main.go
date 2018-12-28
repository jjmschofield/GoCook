package main

import (
	"flag"
	"github.com/jjmschofield/GoCook/api"
	"github.com/jjmschofield/GoCook/common/config"
	"github.com/jjmschofield/GoCook/common/db"
	"github.com/jjmschofield/GoCook/common/logger"
)

type cliArgs struct {
	port    string
	profile string
	db      string
}

func main() {
	args := getCommandLineArgs()

	logger.Init(args.profile)

	logger.Info("Server starting...")

	config.LoadNonSensitiveConfig()

	db.ConnectDb(args.db)

	logger.Info("Server listening via HTTP on " + args.port)

	api.Start(args.port)
}

func getCommandLineArgs() cliArgs {
	profile := flag.String("profile", "prod", "A profile specifying the type of environment dev/prod")
	port := flag.String("port", "8080", "The port to bind the webserver to")
	dbStr := flag.String("db", "postgres://postgres:password@localhost/gocook?sslmode=disable", "Database connection string")
	flag.Parse()

	return cliArgs{
		profile: *profile,
		port:    *port,
		db:      *dbStr,
	}
}
