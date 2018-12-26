package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var dbPool *sql.DB

func ConnectDb(connStr string) {
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

	if conErr != nil {
		log.Fatal(conErr)
	}
}

func IsConnected() error {
	return dbPool.Ping()
}

func GetConnection() *sql.DB {
	return dbPool
}

func SingleRowQuery(query string) (result []byte, err error) {
	err = GetConnection().QueryRow(query).Scan(&result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func MultiRowQuery(query string) (results [][]byte, err error) {
	rows, queryErr := GetConnection().Query(query)

	if queryErr != nil {
		return nil, queryErr
	}

	defer rows.Close()

	for rows.Next() {

		var thisResult []byte

		err = rows.Scan(&thisResult)

		if err != nil {
			return nil, err
		}

		results = append(results, thisResult)
	}

	rowsErr := rows.Err()

	if rowsErr != nil {
		return nil, err
	}
	return results, err
}
