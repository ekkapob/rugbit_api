package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var (
	user   = "rugbit"
	pwd    = "rugbit_123"
	dbname = "rugbitdb"
)

func GetDb() *sql.DB {
	db, err := sql.Open("postgres", "user="+user+" dbname="+pwd+" sslmode="+dbname)
	if err != nil {
		panic(err)
	}
	return db
}
