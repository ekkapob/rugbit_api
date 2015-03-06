package authen

import (
	"database/sql"
)

type User struct {
	id        int
	username  string
	firstname string
	lastname  string
}

func GetUser(db *sql.DB, username string) (err error, user User) {
	var (
		id                  int
		firstname, lastname string
	)
	err = db.QueryRow("select id, firstname, lastname from users where username = $1", username).Scan(&id, &firstname, &lastname)
	if err == nil {
		user = User{id, username, firstname, lastname}
	}
	return err, user
}

func getHashPassword(db *sql.DB, username string) (err error, hashPassword string) {
	err = db.QueryRow("select password from users where username = $1", username).Scan(&hashPassword)
	return err, hashPassword
}
