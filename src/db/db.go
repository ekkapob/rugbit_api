package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	DB_USER = "rugbit"
	DB_PWD  = "rugbit_123"
	DB_NAME = "rugbitdb"
)

type User struct {
	id        int
	username  string
	firstname string
	lastname  string
}

// func GetUserPwd(db *sql.DB, email string) (password string) {
// 	_ = db.QueryRow("select password from users where username = $1 ", email).Scan(&password)
// 	return password
// }

// func GetUserId(db *sql.DB, username string) int {
// 	var id int
// 	err := db.QueryRow("select id from users where username = $1", username).Scan(&id)
// 	if err != sql.ErrNoRows {
// 		return id
// 	}
// 	return 0
// }

func GetDb() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PWD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
	return db
}
