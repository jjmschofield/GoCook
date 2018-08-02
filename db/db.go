package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

func ConnectDb(connStr string){
	db, err := sql.Open(
		"postgres",
		connStr,
		)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

func GetConnection() *sql.DB{

	if db == nil{
		panic("Database is not connected")
	}

	return db
}
