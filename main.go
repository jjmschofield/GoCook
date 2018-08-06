package main

import (
	"github.com/jjmschofield/GoCook/common/config"
	"github.com/jjmschofield/GoCook/common/db"
	"github.com/jjmschofield/GoCook/api"
	"flag"
)

type cliArgs struct {
	port string
	db   string
}

func main() {
	args := getCommandLineArgs()

	config.LoadNonSensitiveConfig()

	db.ConnectDb(args.db)

	api.Start(args.port)
}

func getCommandLineArgs() cliArgs {
	port := flag.String("port", "8080", "The port to bind the webserver to")
	dbStr := flag.String("db", "postgres://postgres:password@localhost/gocook?sslmode=disable", "Database connection string")
	flag.Parse()

	return cliArgs{
		port: *port,
		db:   *dbStr,
	}
}

// @title GoCook API Reference
// @version 1.0
// @description An API for working for collaborating on recipes and sharing meal plans.
// @termsOfService https://not-implemented.com

// @contact.name API Support
// @contact.url https://github.com/jjmschofield/GoCook

// @license.name No Licence
// @license.url https://github.com/jjmschofield/GoCook

// @schemes http
// @host go-cook.herokuapp.com
// @BasePath /
