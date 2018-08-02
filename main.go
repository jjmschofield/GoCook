package main

import (
	"github.com/jjmschofield/GoCook/config"
	"github.com/jjmschofield/GoCook/server"
	"flag"
	"github.com/jjmschofield/GoCook/db"
)

type cliArgs struct{
	port string
	db string
}

func main() {
	args := getCommandLineArgs()

	config.LoadNonSensitiveConfig()

	db.ConnectDb(args.db)

	server.Start(args.port)
}

func getCommandLineArgs() cliArgs{
	port := flag.String("port", "8080", "The port to bind the webserver to")
	dbStr := flag.String("db", "postgres://postgres:password@localhost/gocook?sslmode=disable", "Database connection string")
	flag.Parse()

	return cliArgs{
		port: *port,
		db: *dbStr,
	}
}
