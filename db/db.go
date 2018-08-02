package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var dbPool *sql.DB

func ConnectDb(connStr string){
	db, err := sql.Open(
		"postgres",
		connStr,
		)

	if err != nil {
		log.Fatal(err)
	}

	dbPool = db

	dbPool.SetMaxOpenConns(20)

	conErr := IsConnected()

	if conErr != nil{
		log.Fatal(conErr)
	}
}

func IsConnected() error{
	return dbPool.Ping()
}

func GetConnection() *sql.DB{
	return dbPool
}